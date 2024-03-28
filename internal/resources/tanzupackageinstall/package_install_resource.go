/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackageinstall

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	tanzupackage "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
	tanzupakageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackage"
	pkginstallclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/status"
)

type contextMethodKey struct{}

func ResourcePackageInstall() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePackageInstallCreate,
		DeleteContext: resourcePackageInstallDelete,
		UpdateContext: resourcePackageInstallInPlaceUpdate,
		ReadContext:   dataPackageInstallRead,
		Schema:        getResourceSchema(),
		CustomizeDiff: customdiff.All(
			schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey})),
			schema.CustomizeDiffFunc(spec.ValidateInlineValues()),
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
	var packageInstallSchema = map[string]*schema.Schema{
		nameKey: {
			Type:        schema.TypeString,
			Description: "Name of the package install resource.",
			Required:    true,
			ForceNew:    true,
		},
		NamespaceKey: {
			Type:        schema.TypeString,
			Description: "Name of Namespace where package install will be created.",
			Required:    true,
			ForceNew:    true,
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
			packageInstallSchema[key] = helper.UpdateDataSourceSchema(value)
		} else {
			packageInstallSchema[key] = value
		}
	}

	return packageInstallSchema
}

func resourcePackageInstallCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageInstallName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package install name")
	}

	packageInstallNamespace, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package install namespace name")
	}

	scopedFullnameData, _ := scope.ConstructScope(d, packageInstallName, packageInstallNamespace)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package install entry; Scope full name is empty")
	}

	specVal, err := spec.ConstructSpecForClusterScope(d)
	if err != nil {
		return diag.FromErr(err)
	}

	packageInstallReq := &pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest{
		Install: &pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall{
			FullName: scopedFullnameData.FullnameCluster,
			Meta:     common.ConstructMeta(d),
			Spec:     specVal,
		},
	}

	packageInstallResponse, err := config.TMCConnection.PackageInstallResourceService.InstallResourceServiceCreate(packageInstallReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster package install entry, name : %s", packageInstallName))
	}

	UID := packageInstallResponse.Install.Meta.UID

	// always run
	d.SetId(UID)

	return dataPackageInstallRead(ctx, d, m)
}

func resourcePackageInstallDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageInstallName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package install name")
	}

	packageInstallNamespace, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package install namespace name")
	}

	scopedFullnameData, _ := scope.ConstructScope(d, packageInstallName, packageInstallNamespace)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package install entry; Scope full name is empty")
	}

	err := config.TMCConnection.PackageInstallResourceService.InstallResourceServiceDelete(scopedFullnameData.FullnameCluster)
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster package install entry, name : %s", packageInstallName))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func resourcePackageInstallInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageInstallName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package install name")
	}

	packageInstallNamespace, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package install namespace name")
	}

	scopedFullnameData, _ := scope.ConstructScope(d, packageInstallName, packageInstallNamespace)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package install entry; Scope full name is empty")
	}

	_, meta, installSpec, _, err := retrievePackageInstallUIDMetaAndSpecFromServer(config, scopedFullnameData, d)
	if err != nil {
		return diag.FromErr(err)
	}

	var updateAvailable bool

	if updateCheckForMeta(d, meta) {
		updateAvailable = true
	}

	specCheck, err := updateCheckForSpec(d, installSpec)
	if err != nil {
		log.Println("[ERROR] Unable to check spec has been updated.")
		return diag.FromErr(err)
	}

	if specCheck {
		err = CheckForUpdatedPackage(config, scopedFullnameData, installSpec)
		if err != nil {
			return diag.FromErr(err)
		}

		updateAvailable = true
	}

	if !updateAvailable {
		log.Printf("[INFO] package install update is not required")
		return diags
	}

	pkgInstallReq := &pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest{
		Install: &pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall{
			FullName: scopedFullnameData.FullnameCluster,
			Meta:     meta,
			Spec:     installSpec,
		},
	}

	_, err = config.TMCConnection.PackageInstallResourceService.InstallResourceServiceUpdate(pkgInstallReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster package install entry, name : %s", packageInstallName))
	}

	return dataPackageInstallRead(ctx, d, m)
}

func retrievePackageInstallUIDMetaAndSpecFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (
	string,
	*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta,
	*pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec,
	*pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus,
	error) {
	var (
		UID                string
		meta               *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
		spec               *pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec
		clusterScopeStatus *pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus
	)

	resp, err := config.TMCConnection.PackageInstallResourceService.InstallResourceServiceGet(scopedFullnameData.FullnameCluster)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			d.SetId("")
			return "", nil, nil, nil, err
		}

		return "", nil, nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster package install entry, name : %s", scopedFullnameData.FullnameCluster.Name)
	}

	scopedFullnameData = &scope.ScopedFullname{
		Scope:           commonscope.ClusterScope,
		FullnameCluster: resp.Install.FullName,
	}

	fullName, name, namespace := scope.FlattenScope(scopedFullnameData)

	if err := d.Set(nameKey, name); err != nil {
		return "", nil, nil, nil, err
	}

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return "", nil, nil, nil, err
	}

	if err := d.Set(NamespaceKey, namespace); err != nil {
		return "", nil, nil, nil, err
	}

	UID = resp.Install.Meta.UID
	meta = resp.Install.Meta
	spec = resp.Install.Spec
	clusterScopeStatus = resp.Install.Status

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

func updateCheckForSpec(d *schema.ResourceData, installSpec *pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec) (bool, error) {
	if !spec.HasSpecChanged(d) {
		return false, nil
	}

	specVal, err := spec.ConstructSpecForClusterScope(d)
	if err != nil {
		return false, err
	}

	installSpec.PackageRef.PackageMetadataName = specVal.PackageRef.PackageMetadataName
	installSpec.PackageRef.VersionSelection.Constraints = specVal.PackageRef.VersionSelection.Constraints
	installSpec.RoleBindingScope = specVal.RoleBindingScope
	installSpec.InlineValues = specVal.InlineValues

	log.Printf("[INFO] updating package install spec")

	return true, nil
}

func GetGlobalNamespace(config authctx.TanzuContext, searchscope *tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageSearchScope) (string, error) {
	response, err := config.TMCConnection.ClusterTanzuPackageService.TanzuPackageResourceServiceList(searchscope)
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

	return globalNs, nil
}

func CheckForUpdatedPackage(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, spec *pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec) error {
	globalNs, err := GetGlobalNamespace(config, &tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageSearchScope{
		ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
		ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
		ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
	})
	if err != nil {
		return err
	}

	resp, err := config.TMCConnection.TanzupackageResourceService.ManageV1alpha1ClusterPackageResourceServiceGet(&tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageFullName{
		ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
		ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
		ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
		MetadataName:          spec.PackageRef.PackageMetadataName,
		Name:                  spec.PackageRef.VersionSelection.Constraints,
		NamespaceName:         globalNs,
		OrgID:                 scopedFullnameData.FullnameCluster.OrgID,
	})

	if resp == nil || err != nil {
		return fmt.Errorf("unable to get package metatdata for update: %v", err.Error())
	}

	return nil
}
