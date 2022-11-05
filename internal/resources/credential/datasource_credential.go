/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package credential

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceCredential() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCredentialRead,
		Schema:      credentialSchema,
	}
}

func dataSourceCredentialRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	var (
		diags diag.Diagnostics
		resp  *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialGetCredentialResponse
		err   error
	)

	name, _ := d.Get(NameKey).(string)

	resp, err = config.TMCConnection.CredentialResourceService.CredentialResourceServiceGet(constructFullname(d))
	if err != nil || resp == nil {
		if clienterrors.IsNotFoundError(err) {
			d.SetId("")
			return diags
		}

		return diag.FromErr(errors.Wrapf(err, "unable to get Tanzu Mission Control credential entry, name : %s", name))
	}

	d.SetId(resp.Credential.Meta.UID)

	status := map[string]interface{}{
		"phase":      resp.Credential.Status.Phase,
		"phase_info": resp.Credential.Status.PhaseInfo,
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Credential.Meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(statusKey, status); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
