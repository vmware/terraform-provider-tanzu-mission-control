/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"
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
)

type contextMethodKey struct{}

func ResourceSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecretCreate,
		DeleteContext: resourceSecretDelete,
		UpdateContext: resourceSecretInPlaceUpdate,
		ReadContext:   dataSourceSecretRead,
		Schema:        getSecretSchema(false),
	}
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
		specKey: secretSpec,
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

type dockerConfigJSON struct {
	Auths map[string]map[string]interface{} `json:"auths,omitempty"`
}

var secretSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the kubernetes secret",
	Required:    true,
	// Optional:    false,
	// Default:     nil,
	MaxItems: 1,
	MinItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			DockerConfigjsonKey: {
				Type:        schema.TypeList,
				Required:    true,
				Description: "SecretType definition - SECRET_TYPE_DOCKERCONFIGJSON, Kubernetes secrets type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						UsernameKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Username of the registry.",
							Required:    true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
						PasswordKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Password of the registry.",
							Required:    true,
							Sensitive:   true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
						ImageRegistryURLKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Server URL of the registry.",
							Required:    true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
					},
				},
			},
		},
	},
}

func constructSpec(d *schema.ResourceData) (spec *secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec) {
	spec = &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	spec.SecretType = secretmodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEDOCKERCONFIGJSON)

	if v, ok := specData[DockerConfigjsonKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			auth := v1[0].(map[string]interface{})

			var serverURL, username, password string

			if v, ok := auth[ImageRegistryURLKey]; ok {
				serverURL = v.(string)
			}

			if v, ok := auth[UsernameKey]; ok {
				username = v.(string)
			}

			if v, ok := auth[PasswordKey]; ok {
				password = v.(string)
			}

			secretSpecData, err := getEncodedSpecData(serverURL, username, password)
			if err != nil {
				return spec
			}

			spec.Data = map[string]strfmt.Base64{
				dockerConfigJSONSecretDataKey: secretSpecData,
			}
		}
	}

	return spec
}

func flattenSpec(spec *secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec, pass string) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	secretData, ok := spec.Data[dockerConfigJSONSecretDataKey]
	if !ok {
		return data
	}

	dockerConfigJSON, err := getDecodedSpecData(secretData)
	if err != nil {
		return data
	}

	var dockerConfigJSONData = make(map[string]interface{})

	for serverURL, creds := range dockerConfigJSON.Auths {
		for attribute, value := range creds {
			dockerConfigJSONData[ImageRegistryURLKey] = serverURL

			if attribute == UsernameKey {
				stringValue, ok := value.(string)
				if !ok {
					return data
				}

				dockerConfigJSONData[UsernameKey] = stringValue
			}
		}
	}

	dockerConfigJSONData[PasswordKey] = pass

	flattenSpecData[DockerConfigjsonKey] = []interface{}{dockerConfigJSONData}

	return []interface{}{flattenSpecData}
}

func getEncodedSpecData(serverURL, username, password string) (strfmt.Base64, error) {
	var secretspecdata strfmt.Base64

	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	err := secretspecdata.Scan(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"auths":{"%s":{"username":"%s","password":"%s","auth":"%s"}}}`, serverURL, username, password, auth))))
	if err != nil {
		return nil, err
	}

	return secretspecdata, nil
}

func getDecodedSpecData(data strfmt.Base64) (*dockerConfigJSON, error) {
	rawData, err := base64.StdEncoding.DecodeString(data.String())
	if err != nil {
		return nil, err
	}

	dockerConfigJSON := &dockerConfigJSON{}

	err = json.Unmarshal(rawData, dockerConfigJSON)
	if err != nil {
		return nil, err
	}

	return dockerConfigJSON, nil
}

func resourceSecretCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	secretRequest := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest{
		Secret: &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecret{
			FullName: constructFullname(d),
			Meta:     common.ConstructMeta(d),
			Spec:     constructSpec(d),
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

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	_ = schema.RemoveFromState(d, m)

	return diags
}

func resourceSecretInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(specKey, DockerConfigjsonKey, ImageRegistryURLKey)):
		return diag.Errorf("updating %v is not possible", ImageRegistryURLKey)
	case common.HasMetaChanged(d):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(specKey, DockerConfigjsonKey, UsernameKey)) || d.HasChange(helper.GetFirstElementOf(specKey, DockerConfigjsonKey, PasswordKey)) || d.HasChange(ExportKey):
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

	getResp.Secret.Spec = constructSpec(d)

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
