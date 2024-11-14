// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackagerepository

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	tanzupakageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackage"
	pkgrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/status"
)

type dataFromServer struct {
	UID    string
	meta   *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
	spec   *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec
	status *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus
}

type contextMethodKey struct{}

func ResourcePackageRepository() *schema.Resource {
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
	return getPkgRepoSchema(false)
}

func getDataSourceSchema() map[string]*schema.Schema {
	return getPkgRepoSchema(true)
}

func getPkgRepoSchema(isDataSource bool) map[string]*schema.Schema {
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
			Description: "If true, Package Repository is disabled for cluster.",
			Optional:    true,
			Default:     false,
		},
		commonscope.ScopeKey: scope.ScopeSchema,
		common.MetaKey:       common.Meta,
		status.StatusKey:     status.StatusSchema,
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

func resourcePackageRepositoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package repository name")
	}

	scopedFullnameData := scope.ConstructScope(d, packageRepositoryName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package repository entry; Scope full name is empty")
	}

	_, err := GetGlobalNamespace(config, scopedFullnameData, d)
	if err != nil {
		return diag.Errorf("failed to get package repository global namespace for cluster: %v", err)
	}

	packageRepositoryReq := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest{
		Repository: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
			FullName: scopedFullnameData.FullnameCluster,
			Meta:     common.ConstructMeta(d),
			Spec:     spec.ConstructSpecForClusterScope(d),
		},
	}

	packageRepositoryResponse, err := config.TMCConnection.ClusterPackageRepositoryService.RepositoryResourceServiceCreate(packageRepositoryReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster package repository entry, name : %s", packageRepositoryName))
	}

	UID := packageRepositoryResponse.Repository.Meta.UID

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
				return diag.FromErr(errors.Wrapf(err, "Tanzu Mission Control cluster package repository  not found, name : %s", packageRepositoryName))
			}

			return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control package repository entry, name : %s", packageRepositoryName))
		}
	}

	return dataPackageRepositoryRead(ctx, d, m)
}

func resourcePackageRepositoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package repository name")
	}

	scopedFullnameData := scope.ConstructScope(d, packageRepositoryName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package repository entry; Scope full name is empty")
	}

	packageRepositoryNamespacename, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package repository name")
	}

	scopedFullnameData.FullnameCluster.NamespaceName = packageRepositoryNamespacename

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
		return diag.Errorf("Unable to read package repository name")
	}

	scopedFullnameData := scope.ConstructScope(d, packageRepositoryName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package repository entry; Scope full name is empty")
	}

	packageRepositoryNamespacename, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package repository name")
	}

	scopedFullnameData.FullnameCluster.NamespaceName = packageRepositoryNamespacename

	pkgRepoDataFromServer, err := retrievePackageRepositoryUIDMetaAndSpecFromServer(config, scopedFullnameData, d)
	if err != nil {
		return diag.FromErr(err)
	}

	if updateCheckForMeta(d, pkgRepoDataFromServer.meta) || updateCheckForSpec(d, pkgRepoDataFromServer.spec) {
		pkgRepoReq := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest{
			Repository: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
				FullName: scopedFullnameData.FullnameCluster,
				Meta:     pkgRepoDataFromServer.meta,
				Spec:     pkgRepoDataFromServer.spec,
			},
		}

		_, err = config.TMCConnection.ClusterPackageRepositoryService.RepositoryResourceServiceUpdate(pkgRepoReq)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster package repository entry, name : %s", packageRepositoryName))
		}
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

	if !updateCheckForMeta(d, pkgRepoDataFromServer.meta) && !updateCheckForSpec(d, pkgRepoDataFromServer.spec) && !d.HasChange(disabledKey) {
		return
	}

	return dataPackageRepositoryRead(ctx, d, m)
}

func retrievePackageRepositoryUIDMetaAndSpecFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (
	*dataFromServer, error) {
	var pkgRepoDataFromServer = &dataFromServer{}

	resp, err := config.TMCConnection.ClusterPackageRepositoryService.RepositoryResourceServiceGet(scopedFullnameData.FullnameCluster)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			d.SetId("")
			return pkgRepoDataFromServer, err
		}

		return pkgRepoDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster package repository entry, name : %s", scopedFullnameData.FullnameCluster.Name)
	}

	scopedFullnameData.FullnameCluster = resp.Repository.FullName
	pkgRepoDataFromServer.UID = resp.Repository.Meta.UID
	pkgRepoDataFromServer.meta = resp.Repository.Meta
	pkgRepoDataFromServer.spec = resp.Repository.Spec
	pkgRepoDataFromServer.status = resp.Repository.Status

	scopedFullnameData = &scope.ScopedFullname{
		Scope:           commonscope.ClusterScope,
		FullnameCluster: resp.Repository.FullName,
	}

	fullName, name, namespace := scope.FlattenScope(scopedFullnameData)

	if err := d.Set(nameKey, name); err != nil {
		return pkgRepoDataFromServer, err
	}

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return pkgRepoDataFromServer, err
	}

	if err := d.Set(NamespaceKey, namespace); err != nil {
		return pkgRepoDataFromServer, err
	}

	return pkgRepoDataFromServer, nil
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

	updatedRepoSpec := spec.ConstructSpecForClusterScope(d)

	repoSpec.ImgpkgBundle.Image = updatedRepoSpec.ImgpkgBundle.Image

	return true
}

func GetGlobalNamespace(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (string, error) {
	ss := &tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageSearchScope{
		ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
		ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
		ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
	}

	response, err := config.TMCConnection.ClusterTanzuPackageService.TanzuPackageResourceServiceList(ss)
	if err != nil {
		return "", err
	}

	if len(response.TanzuPackages) == 0 {
		return "", fmt.Errorf("cluster not found")
	}

	globalNs := (response.TanzuPackages[0]).Status.PackageRepositoryGlobalNamespace

	if globalNs == "" {
		return "", fmt.Errorf("global namespace not set for cluster")
	}

	scopedFullnameData.FullnameCluster.NamespaceName = globalNs

	if err := d.Set(NamespaceKey, globalNs); err != nil {
		return "", err
	}

	return globalNs, nil
}
