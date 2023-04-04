/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	namespacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/namespace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceNamespace() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceNamespaceRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: namespaceSchema,
	}
}

func dataSourceNamespaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	var (
		diags diag.Diagnostics
		resp  *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse
		err   error
	)

	namespaceName, _ := d.Get(NameKey).(string)

	resp, err = config.TMCConnection.NamespaceResourceService.ManageV1alpha1NamespaceResourceServiceGet(constructFullname(d))
	if err != nil || resp == nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return diags
		}

		return diag.FromErr(errors.Wrapf(err, "unable to get Tanzu Mission Control namespace entry, name : %s", namespaceName))
	}

	d.SetId(resp.Namespace.Meta.UID)

	status := map[string]interface{}{
		"phase":      resp.Namespace.Status.Phase,
		"phase_info": resp.Namespace.Status.PhaseInfo,
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Namespace.Meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(statusKey, status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(specKey, flattenSpec(resp.Namespace.Spec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
