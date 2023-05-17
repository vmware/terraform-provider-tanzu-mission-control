/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	pkgrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/status"
)

type contextMethodKey struct{}

func ResourceSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePackageRepositoryCreate,
		DeleteContext: resourcePackageRepositoryDelete,
		UpdateContext: resourcePackageRepositoryInPlaceUpdate,
		ReadContext:   dataPackageRepositoryRead,
		Schema:        getResourceSchema(),
		CustomizeDiff: customdiff.All(
			schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey})),
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
	var packageRepositorySchema = map[string]*schema.Schema{
		nameKey: {
			Type:        schema.TypeString,
			Description: "Name of the package repository resource.",
			Required:    true,
			ForceNew:    true,
		},
		NamespaceKey: {
			Type:        schema.TypeString,
			Description: "Name of Namespace where package repository will be created.",
			Computed:    true,
		},
		disabledKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
			Default:     false,
		},
		commonscope.ScopeKey: scope.ScopeSchema,
		status.StateKey:      status.StatusSchema,
	}

	innerMap := map[string]*schema.Schema{
		spec.SpecKey: spec.SpecSchema,
	}

	for key, value := range innerMap {
		if isDataSource {
			packageRepositorySchema[key] = helper.UpdateDataSourceSchema(value)
		} else {
			packageRepositorySchema[key] = value
		}
	}

	return packageRepositorySchema
}

// var packageRepositorySchema = map[string]*schema.Schema{
// 	nameKey: {
// 		Type:        schema.TypeString,
// 		Description: "Name of the package repository resource.",
// 		Required:    true,
// 		ForceNew:    true,
// 	},
// 	NamespaceKey: {
// 		Type:        schema.TypeString,
// 		Description: "Name of Namespace where package repository will be created.",
// 		Computed:    true,
// 	},
// 	disabledKey: {
// 		Type:        schema.TypeBool,
// 		Description: "",
// 		Optional:    true,
// 		Default:     false,
// 	},
// 	commonscope.ScopeKey: scope.ScopeSchema,
// 	spec.SpecKey:         spec.SpecSchema,
// 	status.StateKey:      status.StatusSchema,
// }

func resourcePackageRepositoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read package repository name")
	}

	packageRepositoryNamespace, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("unable to read package repository name")
	}

	scopedFullnameData := scope.ConstructScope(d, packageRepositoryName, packageRepositoryNamespace)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package repository entry; Scope full name is empty")
	}

	var (
		UID  string
		meta = common.ConstructMeta(d)
	)

	packageRepositoryReq := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRepositoryRequest{
		Repository: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
			FullName: scopedFullnameData.FullnameCluster,
			Meta:     meta,
			Spec:     spec.ConstructSpecForClusterScope(d),
		},
	}

	packageRepositoryResponse, err := config.TMCConnection.ClusterPackageRepositoryService.RepositoryResourceServiceCreate(packageRepositoryReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster package repository entry, name : %s", packageRepositoryName))
	}

	UID = packageRepositoryResponse.Repository.Meta.UID

	// always run
	d.SetId(UID)

	if d.Get(disabledKey).(bool) {
		setAvailabilityRequest := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest{
			Disabled: true,
			FullName: scopedFullnameData.FullnameCluster,
		}

		resp, err := config.TMCConnection.ClusterPackageRepositoryAvailabilityService.SetRepositoryAvailability(setAvailabilityRequest)

		if err != nil || resp == nil {
			if clienterrors.IsNotFoundError(err) {
				log.Println("??? while disabling the pkg repo")
				return diag.FromErr(errors.Wrapf(err, "Tanzu Mission Control cluster package repository  not found, name : %s", packageRepositoryName))
			}
			return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control package repository entry, name : %s", packageRepositoryName))
		}
	}

	return dataPackageRepositoryRead(ctx, d, m)
}

func resourcePackageRepositoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read package repository name")
	}

	packageRepositoryNamespace, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("unable to read package repository name")
	}

	scopedFullnameData := scope.ConstructScope(d, packageRepositoryName, packageRepositoryNamespace)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package repository entry; Scope full name is empty")
	}

	err := config.TMCConnection.ClusterPackageRepositoryService.RepositoryResourceServiceDelete(scopedFullnameData.FullnameCluster)
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster package repository entry, name : %s", packageRepositoryName))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func resourcePackageRepositoryInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read package repository name")
	}

	packageRepositoryNamespace, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("unable to read package repository name")
	}

	scopedFullnameData := scope.ConstructScope(d, packageRepositoryName, packageRepositoryNamespace)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package repository entry; Scope full name is empty")
	}

	// nolint: dogsled
	_, meta, repoSpec, _, err := retrievePackageRepositoryUIDMetaAndSpecFromServer(config, scopedFullnameData, d)
	if err != nil {
		return diag.FromErr(err)
	}

	var updateAvailable bool

	if updateCheckForMeta(d, meta) {
		updateAvailable = true
	}

	if updateCheckForSpec(d, repoSpec) {
		updateAvailable = true
	}

	if d.HasChange(disabledKey) {
		updateAvailable = true
	}

	if !updateAvailable {
		return
	}

	pkgRepoReq := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRepositoryRequest{
		Repository: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
			FullName: scopedFullnameData.FullnameCluster,
			Meta:     meta,
			Spec:     repoSpec,
		},
	}

	_, err = config.TMCConnection.ClusterPackageRepositoryService.RepositoryResourceServiceUpdate(pkgRepoReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster package repository entry, name : %s", packageRepositoryName))
	}

	if d.HasChange(disabledKey) {
		pkgrepoavailabilityReq := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest{
			Disabled: d.Get(disabledKey).(bool),
			FullName: scopedFullnameData.FullnameCluster,
		}

		_, err = config.TMCConnection.ClusterPackageRepositoryAvailabilityService.SetRepositoryAvailability(pkgrepoavailabilityReq)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster package repository entry, name : %s", packageRepositoryName))
		}
	}

	return dataPackageRepositoryRead(ctx, d, m)
}

func retrievePackageRepositoryUIDMetaAndSpecFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (
	string,
	*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta,
	*pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec,
	*pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus,
	error) {
	var (
		UID                string
		meta               *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
		spec               *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec
		clusterScopeStatus *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus
	)

	resp, err := config.TMCConnection.ClusterPackageRepositoryService.RepositoryResourceServiceGet(scopedFullnameData.FullnameCluster)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			d.SetId("")
			return "", nil, nil, nil, err
		}

		return "", nil, nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster source secret entry, name : %s", scopedFullnameData.FullnameCluster.Name)
	}

	scopedFullnameData = &scope.ScopedFullname{
		Scope:           commonscope.ClusterScope,
		FullnameCluster: resp.Repository.FullName,
	}

	fullName, name, namespace := scope.FlattenScope(scopedFullnameData) // scope.FlattenScope(scopedFullnameData)

	if err := d.Set(nameKey, name); err != nil {
		return "", nil, nil, nil, err
	}

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return "", nil, nil, nil, err
	}

	if err := d.Set(NamespaceKey, namespace); err != nil {
		return "", nil, nil, nil, err
	}

	UID = resp.Repository.Meta.UID
	meta = resp.Repository.Meta
	spec = resp.Repository.Spec
	clusterScopeStatus = resp.Repository.Status

	return UID, meta, spec, clusterScopeStatus, nil
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

func updateCheckForSpec(d *schema.ResourceData, repoSpec *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec) bool {
	if !d.HasChange(helper.GetFirstElementOf(spec.SpecKey, spec.ImgpkgBundleKey, spec.ImageKey)) {
		return false
	}

	var updatedRepoSpec *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec

	updatedRepoSpec = spec.ConstructSpecForClusterScope(d)

	repoSpec.ImgpkgBundle.Image = updatedRepoSpec.ImgpkgBundle.Image

	return true
}
