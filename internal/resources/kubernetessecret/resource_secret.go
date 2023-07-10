/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	secretmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	secretexportmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster/secretexport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/spec"
)

type contextMethodKey struct{}

func ResourceSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecretCreate,
		DeleteContext: resourceSecretDelete,
		UpdateContext: resourceSecretInPlaceUpdate,
		ReadContext:   dataSourceSecretRead,
		Schema:        getResourceSchema(),
	}
}

func getResourceSchema() map[string]*schema.Schema {
	return getSecretSchema(false)
}

func getDataSourceSchema() map[string]*schema.Schema {
	return getSecretSchema(true)
}

func getSecretSchema(isDataSource bool) map[string]*schema.Schema {
	var secretSchema = map[string]*schema.Schema{
		NameKey: {
			Type:        schema.TypeString,
			Description: "Name of the secret resource.",
			Required:    true,
			ForceNew:    true,
			ValidateFunc: validation.All(
				validation.StringLenBetween(1, 126),
				validation.StringIsNotEmpty,
				validation.StringIsNotWhiteSpace,
			),
		},
		NamespaceNameKey: {
			Type:        schema.TypeString,
			Description: "Name of Namespace where secret will be created.",
			Required:    true,
			ForceNew:    true,
			ValidateFunc: validation.All(
				validation.StringIsNotEmpty,
				validation.StringIsNotWhiteSpace,
			),
		},
		OrgIDKey: {
			Type:        schema.TypeString,
			Description: "ID of Organization.",
			Optional:    true,
		},
		scope.ScopeKey: scope.ScopeSchema,
		statusKey: {
			Type:        schema.TypeMap,
			Description: "Status for the Secret Export.",
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		common.MetaKey: common.Meta,
	}

	innerMap := map[string]*schema.Schema{
		spec.SpecKey: spec.SecretSpec,
		ExportKey: {
			Type:        schema.TypeBool,
			Description: "Export the secret to all namespaces.",
			Optional:    true,
			Default:     false,
		},
	}

	for key, value := range innerMap {
		if isDataSource {
			secretSchema[key] = helper.UpdateDataSourceSchema(value)
		} else {
			secretSchema[key] = value
		}
	}

	return secretSchema
}

func resourceSecretCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	secretRequest := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest{
		Secret: &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecret{
			FullName: constructFullname(d),
			Meta:     common.ConstructMeta(d),
			Spec:     spec.ConstructSpec(d),
		},
	}

	secretResponse, err := config.TMCConnection.SecretResourceService.SecretResourceServiceCreate(secretRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control secret entry, name : %s", secretRequest.Secret.FullName.Name))
	}

	d.SetId(secretResponse.Secret.Meta.UID)

	if d.Get(ExportKey).(bool) {
		secretexportRequest := &secretexportmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest{
			SecretExport: &secretexportmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExport{
				FullName: constructFullnameSecetExport(d),
				Meta:     common.ConstructMeta(d),
			},
		}

		_, err = config.TMCConnection.SecretExportResourceService.SecretExportResourceServiceCreate(secretexportRequest)

		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control secret export entry, name : %s", NameKey))
		}
	}

	return dataSourceSecretRead(ctx, d, m)
}

func resourceSecretDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	secretName, _ := d.Get(NameKey).(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.SecretExportResourceService.SecretExportResourceServiceDelete(constructFullnameSecetExport(d))
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control secret export entry, name : %s", secretName))
	}

	err = config.TMCConnection.SecretResourceService.SecretResourceServiceDelete(constructFullname(d))
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control secret entry, name : %s", secretName))
	}

	_ = schema.RemoveFromState(d, m)

	return diags
}

func resourceSecretInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(spec.SpecKey, spec.DockerConfigjsonKey, spec.ImageRegistryURLKey)):
		return diag.Errorf("updating %v is not possible", spec.ImageRegistryURLKey)
	case common.HasMetaChanged(d):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(spec.SpecKey, spec.DockerConfigjsonKey, spec.UsernameKey)) || d.HasChange(helper.GetFirstElementOf(spec.SpecKey, spec.DockerConfigjsonKey, spec.PasswordKey)) || d.HasChange(ExportKey):
		updateRequired = true
	}

	if !updateRequired {
		return diags
	}

	getResp, err := config.TMCConnection.SecretResourceService.SecretResourceServiceGet(constructFullname(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to get Tanzu Mission Control secret entry, name : %s", d.Get(scope.ClusterNameKey)))
	}

	if common.HasMetaChanged(d) {
		meta := common.ConstructMeta(d)

		if value, ok := getResp.Secret.Meta.Labels[common.CreatorLabelKey]; ok {
			meta.Labels[common.CreatorLabelKey] = value
		}

		getResp.Secret.Meta.Labels = meta.Labels
		getResp.Secret.Meta.Description = meta.Description
	}

	getResp.Secret.Spec = spec.ConstructSpec(d)

	_, err = config.TMCConnection.SecretResourceService.SecretResourceServiceUpdate(
		&secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest{
			Secret: getResp.Secret,
		},
	)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to update Tanzu Mission Control secret entry, name : %s", d.Get(scope.ClusterNameKey)))
	}

	if d.HasChange(ExportKey) {
		secretexportName := d.Get(NameKey).(string)

		if d.Get(ExportKey).(bool) {
			secretexportRequest := &secretexportmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest{
				SecretExport: &secretexportmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExport{
					FullName: constructFullnameSecetExport(d),
					Meta:     common.ConstructMeta(d),
				},
			}

			_, err = config.TMCConnection.SecretExportResourceService.SecretExportResourceServiceCreate(secretexportRequest)

			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control secret export entry, name : %s", secretexportName))
			}
		} else {
			err = config.TMCConnection.SecretExportResourceService.SecretExportResourceServiceDelete(constructFullnameSecetExport(d))
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control secret export entry, name : %s", secretexportName))
			}
		}
	}

	return dataSourceSecretRead(ctx, d, m)
}
