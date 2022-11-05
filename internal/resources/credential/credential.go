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

func ResourceCredential() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCredentialCreate,
		ReadContext:   dataSourceCredentialRead,
		UpdateContext: resourceCredentialUpdate,
		DeleteContext: resourceCredentialDelete,
		Schema:        credentialSchema,
	}
}

var credentialSchema = map[string]*schema.Schema{
	NameKey: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	common.MetaKey: common.Meta,
	specKey:        credSpec,
	statusKey: {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
}

var credSpec = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			capabilityKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			providerKey: {
				Type:     schema.TypeString,
				Default:  "PROVIDER_UNSPECIFIED",
				Optional: true,
			},
			dataKey: dataSpec,
		},
	},
}

var dataSpec = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			genericCredentialKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			awsCredentialKey: awsCredSpec,
			keyValueKey:      keyValueSpec,
		},
	},
}

func resourceCredentialCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	request := &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialCreateCredentialRequest{
		Credential: &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialCredential{
			FullName: constructFullname(d),
			Meta:     common.ConstructMeta(d),
			Spec:     constructSpec(d),
		},
	}

	response, err := config.TMCConnection.CredentialResourceService.CredentialResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control credential entry, name : %s", NameKey))
	}

	d.SetId(response.Credential.Meta.UID)

	return dataSourceCredentialRead(ctx, d, m)
}

func resourceCredentialDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	namespaceName, _ := d.Get(NameKey).(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.CredentialResourceService.CredentialResourceServiceDelete(constructFullname(d))
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control credential entry, name : %s", namespaceName))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	_ = schema.RemoveFromState(d, m)

	return diags
}

func resourceCredentialUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return diag.FromErr(errors.New("update of Tanzu Mission Control credential is not supported"))
}
