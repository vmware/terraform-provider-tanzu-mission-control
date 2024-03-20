/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"context"
	"log"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustersecretmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	secretexportclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster/secretexport"
	clustergroupsecretmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
	secretexportclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup/secretexport"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/status"
)

func ResourceSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecretCreate,
		DeleteContext: resourceSecretDelete,
		UpdateContext: resourceSecretInPlaceUpdate,
		ReadContext:   dataSourceSecretRead,
		Schema:        getResourceSchema(),
		CustomizeDiff: customdiff.All(
			schema.CustomizeDiffFunc(commonscope.ValidateScope(scope.ScopesAllowed[:])),
			spec.ValidateInput,
		),
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
		commonscope.ScopeKey: scope.ScopeSchema,
		statusKey:            status.StatusSchema,
		common.MetaKey:       common.Meta,
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

	secretName, ok := d.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read secret name")
	}

	secretNamespaceName, ok := d.Get(NamespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read secret namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, secretName, secretNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control secret entry; Scope full name is empty")
	}

	var (
		UID  string
		meta = common.ConstructMeta(d)
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			secretReq := &clustersecretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest{
				Secret: &clustersecretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecret{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
					Spec:     spec.ConstructSpecForClusterScope(d),
				},
			}

			secretResponse, err := config.TMCConnection.SecretResourceService.SecretResourceServiceCreate(secretReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster secret entry, name : %s", secretName))
			}

			UID = secretResponse.Secret.Meta.UID
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			secretReq := &clustergroupsecretmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest{
				Secret: &clustergroupsecretmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
					Spec:     spec.ConstructSpecForClusterGroupScope(d),
				},
			}

			secretResponse, err := config.TMCConnection.ClusterGroupSecretResourceService.SecretResourceServiceCreate(secretReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster group secret entry, name : %s", secretName))
			}

			UID = secretResponse.Secret.Meta.UID
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	d.SetId(UID)

	if d.Get(ExportKey).(bool) {
		err := createDeleteSecretExport(true, scopedFullnameData, config, d)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return dataSourceSecretRead(ctx, d, m)
}

func resourceSecretDelete(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	secretName, ok := d.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read secret name")
	}

	secretNamespaceName, ok := d.Get(NamespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read secret namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, secretName, secretNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to delete Tanzu Mission Control secret entry; Scope full name is empty")
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			err := config.TMCConnection.SecretExportResourceService.SecretExportResourceServiceDelete(
				&secretexportclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName{
					Name:                  scopedFullnameData.FullnameCluster.Name,
					ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
					ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
					ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
					NamespaceName:         scopedFullnameData.FullnameCluster.NamespaceName,
					OrgID:                 scopedFullnameData.FullnameCluster.OrgID,
				},
			)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster secret export entry, name : %s", secretName))
			}

			err = config.TMCConnection.SecretResourceService.SecretResourceServiceDelete(scopedFullnameData.FullnameCluster)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control secret entry, name : %s", secretName))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			err := config.TMCConnection.ClusterGroupSecretExportResourceService.SecretExportResourceServiceDelete(
				&secretexportclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportFullName{
					Name:             scopedFullnameData.FullnameClusterGroup.Name,
					ClusterGroupName: scopedFullnameData.FullnameClusterGroup.ClusterGroupName,
					NamespaceName:    scopedFullnameData.FullnameClusterGroup.NamespaceName,
					OrgID:            scopedFullnameData.FullnameClusterGroup.OrgID,
				},
			)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group secret export entry, name : %s", secretName))
			}

			err = config.TMCConnection.ClusterGroupSecretResourceService.SecretResourceServiceDelete(scopedFullnameData.FullnameClusterGroup)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group secret entry, name : %s", secretName))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	_ = schema.RemoveFromState(d, m)

	return diags
}

func resourceSecretInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	secretName, ok := d.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read secret name")
	}

	secretNamespaceName, ok := d.Get(NamespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read secret namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, secretName, secretNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to update Tanzu Mission Control secret entry; Scope full name is empty")
	}

	secretDataFromServer, err := retrieveSecretDataFromServer(config, scopedFullnameData, d)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange(helper.GetFirstElementOf(spec.SpecKey, spec.DockerConfigjsonKey, spec.ImageRegistryURLKey)) {
		return diag.Errorf("updating %v is not possible", spec.ImageRegistryURLKey)
	}

	updateRequiredForSepc := updateCheckForSpec(d, secretDataFromServer.atomicSpec, scopedFullnameData.Scope)
	updateRequiredForMeta := updateCheckForMeta(d, secretDataFromServer.meta)

	if updateRequiredForSepc || updateRequiredForMeta {
		switch scopedFullnameData.Scope {
		case commonscope.ClusterScope:
			if scopedFullnameData.FullnameCluster != nil {
				secretReq := &clustersecretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest{
					Secret: &clustersecretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecret{
						FullName: scopedFullnameData.FullnameCluster,
						Meta:     secretDataFromServer.meta,
						Spec:     secretDataFromServer.atomicSpec,
					},
				}

				_, err = config.TMCConnection.SecretResourceService.SecretResourceServiceUpdate(secretReq)
				if err != nil {
					return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster secret entry, name : %s", secretName))
				}
			}
		case commonscope.ClusterGroupScope:
			if scopedFullnameData.FullnameClusterGroup != nil {
				secretReq := &clustergroupsecretmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest{
					Secret: &clustergroupsecretmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret{
						FullName: scopedFullnameData.FullnameClusterGroup,
						Meta:     secretDataFromServer.meta,
						Spec: &clustergroupsecretmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec{
							AtomicSpec: secretDataFromServer.atomicSpec,
						},
					},
				}

				_, err = config.TMCConnection.ClusterGroupSecretResourceService.SecretResourceServiceUpdate(secretReq)
				if err != nil {
					return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster group secret entry, name : %s", secretName))
				}
			}
		case commonscope.UnknownScope:
			return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
		}
	}

	if d.HasChange(ExportKey) {
		err := createDeleteSecretExport(d.Get(ExportKey).(bool), scopedFullnameData, config, d)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return dataSourceSecretRead(ctx, d, m)
}

func updateCheckForSpec(d *schema.ResourceData, atomicSpec *clustersecretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec, scope commonscope.Scope) bool {
	if !(spec.HasSpecChanged(d)) {
		if atomicSpec.SecretType == clustersecretmodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(clustersecretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEDOCKERCONFIGJSON) {
			username := d.Get(helper.GetFirstElementOf(spec.SpecKey, spec.DockerConfigjsonKey, spec.UsernameKey))
			password := d.Get(helper.GetFirstElementOf(spec.SpecKey, spec.DockerConfigjsonKey, spec.PasswordKey))
			url := d.Get(helper.GetFirstElementOf(spec.SpecKey, spec.DockerConfigjsonKey, spec.ImageRegistryURLKey))

			secretSpecData, _ := spec.GetEncodedSpecData(url.(string), username.(string), password.(string))

			atomicSpec.Data = map[string]strfmt.Base64{
				spec.DockerconfigKey: secretSpecData,
			}
		}

		if atomicSpec.SecretType == clustersecretmodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(clustersecretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEOPAQUE) {
			kv := d.Get(helper.GetFirstElementOf(spec.SpecKey, spec.OpaqueKey))
			atomicSpec.Data = spec.GetEncodedOpaqueData(kv.(map[string]string))
		}

		return false
	}

	var secretSpec *clustersecretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec

	switch scope {
	case commonscope.ClusterScope:
		secretSpec = spec.ConstructSpecForClusterScope(d)
	case commonscope.ClusterGroupScope:
		clusterGroupScopeSpec := spec.ConstructSpecForClusterGroupScope(d)
		secretSpec = clusterGroupScopeSpec.AtomicSpec
	}

	atomicSpec.SecretType = secretSpec.SecretType
	atomicSpec.Data = secretSpec.Data

	log.Printf("[INFO] updating secret spec")

	return true
}

func updateCheckForMeta(d *schema.ResourceData, meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta) bool {
	if !common.HasMetaChanged(d) {
		return false
	}

	objectMeta := common.ConstructMeta(d)

	if value, ok := meta.Labels[common.CreatorLabelKey]; ok {
		objectMeta.Labels[common.CreatorLabelKey] = value
	}

	meta.Labels = objectMeta.Labels
	meta.Description = objectMeta.Description

	log.Printf("[INFO] updating secret meta data")

	return true
}

func createDeleteSecretExport(createKey bool, scopedFullnameData *scope.ScopedFullname, config authctx.TanzuContext, d *schema.ResourceData) error {
	if createKey {
		switch scopedFullnameData.Scope {
		case commonscope.ClusterScope:
			if scopedFullnameData.FullnameCluster != nil {
				secretReq := &secretexportclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest{
					SecretExport: &secretexportclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExport{
						FullName: &secretexportclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName{
							Name:                  scopedFullnameData.FullnameCluster.Name,
							ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
							ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
							ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
							NamespaceName:         scopedFullnameData.FullnameCluster.NamespaceName,
							OrgID:                 scopedFullnameData.FullnameCluster.OrgID,
						},
						Meta: common.ConstructMeta(d),
					},
				}

				_, err := config.TMCConnection.SecretExportResourceService.SecretExportResourceServiceCreate(secretReq)
				if err != nil {
					return errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster secret export entry, name : %s", scopedFullnameData.FullnameCluster.Name)
				}
			}
		case commonscope.ClusterGroupScope:
			if scopedFullnameData.FullnameClusterGroup != nil {
				secretReq := &secretexportclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportRequest{
					SecretExport: &secretexportclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportSecretExport{
						FullName: &secretexportclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportFullName{
							Name:             scopedFullnameData.FullnameClusterGroup.Name,
							ClusterGroupName: scopedFullnameData.FullnameClusterGroup.ClusterGroupName,
							NamespaceName:    scopedFullnameData.FullnameClusterGroup.NamespaceName,
							OrgID:            scopedFullnameData.FullnameClusterGroup.OrgID,
						},
						Meta: common.ConstructMeta(d),
					},
				}

				_, err := config.TMCConnection.ClusterGroupSecretExportResourceService.SecretExportResourceServiceCreate(secretReq)
				if err != nil {
					return errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster group secret export entry, name : %s", scopedFullnameData.FullnameClusterGroup.Name)
				}
			}
		case commonscope.UnknownScope:
			return errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
		}
	} else {
		switch scopedFullnameData.Scope {
		case commonscope.ClusterScope:
			if scopedFullnameData.FullnameCluster != nil {
				err := config.TMCConnection.SecretExportResourceService.SecretExportResourceServiceDelete(
					&secretexportclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName{
						Name:                  scopedFullnameData.FullnameCluster.Name,
						ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
						ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
						ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
						NamespaceName:         scopedFullnameData.FullnameCluster.NamespaceName,
						OrgID:                 scopedFullnameData.FullnameCluster.OrgID,
					},
				)
				if err != nil && !clienterrors.IsNotFoundError(err) {
					return errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster secret export entry, name : %s", scopedFullnameData.FullnameCluster.Name)
				}
			}
		case commonscope.ClusterGroupScope:
			if scopedFullnameData.FullnameClusterGroup != nil {
				err := config.TMCConnection.ClusterGroupSecretExportResourceService.SecretExportResourceServiceDelete(
					&secretexportclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportFullName{
						Name:             scopedFullnameData.FullnameClusterGroup.Name,
						ClusterGroupName: scopedFullnameData.FullnameClusterGroup.ClusterGroupName,
						NamespaceName:    scopedFullnameData.FullnameClusterGroup.NamespaceName,
						OrgID:            scopedFullnameData.FullnameClusterGroup.OrgID,
					},
				)
				if err != nil && !clienterrors.IsNotFoundError(err) {
					return errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group secret export entry, name : %s", scopedFullnameData.FullnameClusterGroup.Name)
				}
			}
		case commonscope.UnknownScope:
			return errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
		}
	}

	return nil
}
