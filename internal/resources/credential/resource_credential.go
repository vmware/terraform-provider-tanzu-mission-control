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
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
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
		Type:        schema.TypeString,
		Description: "Name of this credential",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey: common.Meta,
	specKey:        credSpec,
	statusKey: {
		Type:        schema.TypeMap,
		Description: "Status of credential resource",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	waitKey: {
		Type:        schema.TypeString,
		Description: "Wait timeout duration until credential resource reaches VALID state. Accepted timeout duration values like 5s, 5m, or 1h, higher than zero.",
		Default:     defaultWaitTimeout.String(),
		Optional:    true,
	},
}

var credSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec of credential resource",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			capabilityKey: {
				Type:        schema.TypeString,
				Description: "The Tanzu capability for which the credential shall be used. Value must be in list [DATA_PROTECTION TANZU_OBSERVABILITY TANZU_SERVICE_MESH PROXY_CONFIG MANAGED_K8S_PROVIDER IMAGE_REGISTRY]",
				Optional:    true,
			},
			providerKey: {
				Type:        schema.TypeString,
				Description: "The Tanzu provider for which describes credential data type. Value must be in list [PROVIDER_UNSPECIFIED,AWS_EC2,GENERIC_S3,AZURE_AD,AWS_EKS,AZURE_AKS,GENERIC_KEY_VALUE]",
				Default:     string(credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderPROVIDERUNSPECIFIED),
				Optional:    true,
			},
			dataKey: dataSpec,
		},
	},
}

var dataSpec = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	MaxItems:    1,
	Description: "Holds credentials sensitive data",
	Sensitive:   true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			genericCredentialKey: {
				Type:        schema.TypeString,
				Description: "Generic credential data type used to hold a blob of data represented as string",
				Optional:    true,
			},
			awsCredentialKey: awsCredSpec,
			keyValueKey:      keyValueSpec,
		},
	},
}

func resourceCredentialCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(d, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control credential."))
	}

	request := &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialCreateCredentialRequest{
		Credential: model,
	}

	response, err := config.TMCConnection.CredentialResourceService.CredentialResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control credential entry, name : %s", NameKey))
	}

	d.SetId(response.Credential.Meta.UID)

	return dataSourceCredentialRead(helper.GetContextWithCaller(ctx, helper.CreateState), d, m)
}

func resourceCredentialDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(d, []string{NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control credential."))
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err = config.TMCConnection.CredentialResourceService.CredentialResourceServiceDelete(model.FullName)

	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control credential entry, name : %s", model.FullName.Name))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	_ = schema.RemoveFromState(d, m)

	return diags
}

func resourceCredentialUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return diag.FromErr(errors.New("update of Tanzu Mission Control credential is not supported"))
}
