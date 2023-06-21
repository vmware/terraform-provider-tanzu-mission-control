/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package sourcesecret

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	sourcesecretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
	sourcesecretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret/spec"
)

type contextMethodKey struct{}

func ResourceSourceSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSourcesecretCreate,
		DeleteContext: resourceSourcesecretDelete,
		UpdateContext: resourceSourcesecretInPlaceUpdate,
		ReadContext:   resourceSourcesecretRead,
		Schema:        getResourceSchema(),
		CustomizeDiff: customdiff.All(
			spec.ValidateSpec,
			schema.CustomizeDiffFunc(commonscope.ValidateScope(scope.CredentialTypesAllowed[:])),
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
	var sourcesecretSchema = map[string]*schema.Schema{
		nameKey: {
			Type:        schema.TypeString,
			Description: "Name of the source secret.",
			Required:    true,
			ForceNew:    true,
		},
		orgIDKey: {
			Type:        schema.TypeString,
			Description: "ID of Organization.",
			Optional:    true,
		},
		commonscope.ScopeKey: scope.ScopeSchema,
		common.MetaKey:       common.Meta,
	}

	innerMap := map[string]*schema.Schema{
		spec.SpecKey: spec.SpecSchema,
	}

	for key, value := range innerMap {
		if isDataSource {
			sourcesecretSchema[key] = helper.UpdateDataSourceSchema(value)
		} else {
			sourcesecretSchema[key] = value
		}
	}

	return sourcesecretSchema
}

func retrieveSourcesecretUIDMetaAndSpecFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (
	string,
	*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta,
	*sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec,
	error) {
	var (
		UID  string
		meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
		spec *sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			resp, err := config.TMCConnection.ClusterSourcesecretResourceService.ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceGet(scopedFullnameData.FullnameCluster)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, err
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster source secret entry, name : %s", scopedFullnameData.FullnameCluster.Name)
			}

			scopedFullnameData = &scope.ScopedFullname{
				Scope:           commonscope.ClusterScope,
				FullnameCluster: resp.SourceSecret.FullName,
			}

			fullName, name := scope.FlattenScope(scopedFullnameData)

			if err := d.Set(nameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.SourceSecret.Meta.UID
			meta = resp.SourceSecret.Meta
			spec = resp.SourceSecret.Spec
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupSourcesecretResourceService.ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceGet(scopedFullnameData.FullnameClusterGroup)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, err
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group source secret entry, name : %s", scopedFullnameData.FullnameClusterGroup.Name)
			}

			scopedFullnameData = &scope.ScopedFullname{
				Scope:                commonscope.ClusterGroupScope,
				FullnameClusterGroup: resp.SourceSecret.FullName,
			}

			fullName, name := scope.FlattenScope(scopedFullnameData)

			if err := d.Set(nameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.SourceSecret.Meta.UID
			meta = resp.SourceSecret.Meta
			spec = resp.SourceSecret.Spec.AtomicSpec
		}
	case commonscope.UnknownScope:
		return "", nil, nil, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.CredentialTypesAllowed[:], `, `))
	}

	return UID, meta, spec, nil
}

func resourceSourcesecretCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	sourcesecretName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read source secret name")
	}

	scopedFullnameData, _ := scope.ConstructScope(d, sourcesecretName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control source secret entry; Scope full name is empty")
	}

	var (
		UID  string
		meta = common.ConstructMeta(d)
	)

	err := enableContinuousDelivery(&config, scopedFullnameData, meta)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control source secret entry, name : %s", sourcesecretName))
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			sourcesecretReq := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest{
				SourceSecret: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
					Spec:     spec.ConstructSpecForClusterScope(d),
				},
			}

			sourcesecretResponse, err := config.TMCConnection.ClusterSourcesecretResourceService.ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceCreate(sourcesecretReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster source secret entry, name : %s", sourcesecretName))
			}

			UID = sourcesecretResponse.SourceSecret.Meta.UID
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			sourcesecretReq := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest{
				SourceSecret: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
					Spec:     spec.ConstructSpecForClusterGroupScope(d),
				},
			}

			sourcesecretResponse, err := config.TMCConnection.ClusterGroupSourcesecretResourceService.ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceCreate(sourcesecretReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster group sourcesecret entry, name : %s", sourcesecretName))
			}

			UID = sourcesecretResponse.SourceSecret.Meta.UID
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.CredentialTypesAllowed[:], `, `))
	}

	// always run
	d.SetId(UID)

	return resourceSourcesecretRead(ctx, d, m)
}

func resourceSourcesecretInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	sourcesecretName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read source secret name")
	}

	scopedFullnameData, _ := scope.ConstructScope(d, sourcesecretName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to update Tanzu Mission Control source secret entry; Scope full name is empty")
	}

	_, meta, atomicSpec, err := retrieveSourcesecretUIDMetaAndSpecFromServer(config, scopedFullnameData, d)
	if err != nil {
		return diag.FromErr(err)
	}

	var updateAvailable bool

	if updateCheckForMeta(d, meta) {
		updateAvailable = true
	}

	if updateCheckForSpec(d, atomicSpec, scopedFullnameData.Scope) {
		updateAvailable = true
	}

	if !updateAvailable {
		return
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			sourcesecretReq := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest{
				SourceSecret: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
					Spec:     atomicSpec,
				},
			}

			_, err := config.TMCConnection.ClusterSourcesecretResourceService.ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceUpdate(sourcesecretReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster source secret entry, name : %s", sourcesecretName))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			sourcesecretReq := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest{
				SourceSecret: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
					Spec: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec{
						AtomicSpec: atomicSpec,
					},
				},
			}

			_, err := config.TMCConnection.ClusterGroupSourcesecretResourceService.ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceUpdate(sourcesecretReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster group source secret entry, name : %s", sourcesecretName))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.CredentialTypesAllowed[:], `, `))
	}

	log.Printf("[INFO] source secret update successful")

	return resourceSourcesecretRead(ctx, d, m)
}

func resourceSourcesecretDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	sourcesecretName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read source secret name")
	}

	scopedFullnameData, _ := scope.ConstructScope(d, sourcesecretName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to delete Tanzu Mission Control source secret entry; Scope full name is empty")
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			err := config.TMCConnection.ClusterSourcesecretResourceService.ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceDelete(scopedFullnameData.FullnameCluster)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster source secret entry, name : %s", sourcesecretName))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			err := config.TMCConnection.ClusterGroupSourcesecretResourceService.ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceDelete(scopedFullnameData.FullnameClusterGroup)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group source secret entry, name : %s", sourcesecretName))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.CredentialTypesAllowed[:], `, `))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
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

	return true
}

func updateCheckForSpec(d *schema.ResourceData, atomicSpec *sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec, scope commonscope.Scope) bool {
	if !spec.HasSpecChanged(d) {
		return false
	}

	var sourcesecretSpec *sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec

	switch scope {
	case commonscope.ClusterScope:
		sourcesecretSpec = spec.ConstructSpecForClusterScope(d)
	case commonscope.ClusterGroupScope:
		clusterGroupScopeSpec := spec.ConstructSpecForClusterGroupScope(d)
		sourcesecretSpec = clusterGroupScopeSpec.AtomicSpec
	}

	atomicSpec.Data.Data = sourcesecretSpec.Data.Data
	atomicSpec.Data.Type = sourcesecretSpec.Data.Type
	atomicSpec.SourceSecretType = sourcesecretSpec.SourceSecretType

	return true
}
