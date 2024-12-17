// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kustomization

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization/status"
)

type dataFromServer struct {
	UID                     string
	meta                    *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
	atomicSpec              *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec
	clusterScopeStatus      *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationStatus
	clusterGroupScopeStatus *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationStatus
}

func ResourceKustomization() *schema.Resource {
	return &schema.Resource{
		Schema:        kustomizationSchema,
		CreateContext: resourceKustomizationCreate,
		ReadContext:   resourceKustomizationRead,
		UpdateContext: resourceKustomizationInPlaceUpdate,
		DeleteContext: resourceKustomizationDelete,
		CustomizeDiff: schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey, commonscope.ClusterGroupKey})),
	}
}

var kustomizationSchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the Kustomization.",
		Required:    true,
		ForceNew:    true,
	},
	namespaceNameKey: {
		Type:        schema.TypeString,
		Description: "Name of Namespace.",
		Required:    true,
		ForceNew:    true,
	},
	commonscope.ScopeKey: scope.ScopeSchema,
	common.MetaKey:       common.Meta,
	spec.SpecKey:         spec.SpecSchema,
	statusKey:            status.StatusSchema,
}

func resourceKustomizationRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	kustomizationName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read kustomization name")
	}

	kustomizationNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read kustomization namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, kustomizationName, kustomizationNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control kustomization entry; Scope full name is empty")
	}

	kustomizationDataFromServer, err := retrieveKustomizationDataFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			_ = schema.RemoveFromState(d, m)
			return diags
		}

		return diag.FromErr(err)
	}

	// always run
	d.SetId(kustomizationDataFromServer.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(kustomizationDataFromServer.meta)); err != nil {
		return diag.FromErr(err)
	}

	var (
		flattenedSpec   []interface{}
		flattenedStatus interface{}
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		flattenedSpec = spec.FlattenSpecForClusterScope(kustomizationDataFromServer.atomicSpec)
		flattenedStatus = status.FlattenStatusForClusterScope(kustomizationDataFromServer.clusterScopeStatus)
	case commonscope.ClusterGroupScope:
		clusterGroupScopeSpec := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec{
			AtomicSpec: kustomizationDataFromServer.atomicSpec,
		}
		flattenedSpec = spec.FlattenSpecForClusterGroupScope(clusterGroupScopeSpec)
		flattenedStatus = status.FlattenStatusForClusterGroupScope(kustomizationDataFromServer.clusterGroupScopeStatus)
	}

	if err := d.Set(spec.SpecKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(statusKey, flattenedStatus); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

// nolint: gocognit
func retrieveKustomizationDataFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (*dataFromServer, error) {
	var kustomizationDataFromServer = &dataFromServer{}

	// nolint: dupl
	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			resp, err := config.TMCConnection.ClusterKustomizationResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceGet(scopedFullnameData.FullnameCluster)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return kustomizationDataFromServer, err
				}

				return kustomizationDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster kustomization entry, name : %s", scopedFullnameData.FullnameCluster.Name)
			}

			scopedFullnameData.FullnameCluster = resp.Kustomization.FullName
			kustomizationDataFromServer.UID = resp.Kustomization.Meta.UID
			kustomizationDataFromServer.meta = resp.Kustomization.Meta
			kustomizationDataFromServer.atomicSpec = resp.Kustomization.Spec
			kustomizationDataFromServer.clusterScopeStatus = resp.Kustomization.Status
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupKustomizationResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceGet(scopedFullnameData.FullnameClusterGroup)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return kustomizationDataFromServer, err
				}

				return kustomizationDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group kustomization entry, name : %s", scopedFullnameData.FullnameClusterGroup.Name)
			}

			scopedFullnameData.FullnameClusterGroup = resp.Kustomization.FullName
			kustomizationDataFromServer.UID = resp.Kustomization.Meta.UID
			kustomizationDataFromServer.meta = resp.Kustomization.Meta
			kustomizationDataFromServer.atomicSpec = resp.Kustomization.Spec.AtomicSpec
			kustomizationDataFromServer.clusterGroupScopeStatus = resp.Kustomization.Status
		}
	case commonscope.UnknownScope:
		return kustomizationDataFromServer, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	fullName, name, namespace := scope.FlattenScope(scopedFullnameData)

	if err := d.Set(nameKey, name); err != nil {
		return kustomizationDataFromServer, err
	}

	if err := d.Set(namespaceNameKey, namespace); err != nil {
		return kustomizationDataFromServer, err
	}

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return kustomizationDataFromServer, err
	}

	return kustomizationDataFromServer, nil
}

func resourceKustomizationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	kustomizationName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read kustomization name")
	}

	kustomizationNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read kustomization namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, kustomizationName, kustomizationNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control kustomization entry; Scope full name is empty")
	}

	var (
		UID  string
		meta = common.ConstructMeta(d)
	)

	err := enableContinuousDelivery(&config, scopedFullnameData, meta)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control kustomization entry, name : %s", kustomizationName))
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			kustomizationReq := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest{
				Kustomization: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
					Spec:     spec.ConstructSpecForClusterScope(d),
				},
			}

			kustomizationResponse, err := config.TMCConnection.ClusterKustomizationResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceCreate(kustomizationReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster kustomization entry, name : %s", kustomizationName))
			}

			UID = kustomizationResponse.Kustomization.Meta.UID
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			kustomizationReq := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest{
				Kustomization: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
					Spec:     spec.ConstructSpecForClusterGroupScope(d),
				},
			}

			kustomizationResponse, err := config.TMCConnection.ClusterGroupKustomizationResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceCreate(kustomizationReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster group kustomization entry, name : %s", kustomizationName))
			}

			UID = kustomizationResponse.Kustomization.Meta.UID
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	// always run
	d.SetId(UID)

	return resourceKustomizationRead(ctx, d, m)
}

func resourceKustomizationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	kustomizationName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read kustomization name")
	}

	kustomizationNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read kustomization namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, kustomizationName, kustomizationNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to delete Tanzu Mission Control kustomization entry; Scope full name is empty")
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			err := config.TMCConnection.ClusterKustomizationResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceDelete(scopedFullnameData.FullnameCluster)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster kustomization entry, name : %s", kustomizationName))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			err := config.TMCConnection.ClusterGroupKustomizationResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceDelete(scopedFullnameData.FullnameClusterGroup)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group kustomization entry, name : %s", kustomizationName))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func resourceKustomizationInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	kustomizationName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read kustomization name")
	}

	kustomizationNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read kustomization namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, kustomizationName, kustomizationNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to update Tanzu Mission Control kustomization entry; Scope full name is empty")
	}

	// nolint: dogsled
	kustomizationDataFromServer, err := retrieveKustomizationDataFromServer(config, scopedFullnameData, d)
	if err != nil {
		return diag.FromErr(err)
	}

	var updateAvailable bool

	if updateCheckForMeta(d, kustomizationDataFromServer.meta) {
		updateAvailable = true
	}

	if updateCheckForSpec(d, kustomizationDataFromServer.atomicSpec, scopedFullnameData.Scope) {
		updateAvailable = true
	}

	if !updateAvailable {
		log.Printf("[INFO] kustomization update is not required")
		return diags
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			kustomizationReq := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest{
				Kustomization: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     kustomizationDataFromServer.meta,
					Spec:     kustomizationDataFromServer.atomicSpec,
				},
			}

			_, err = config.TMCConnection.ClusterKustomizationResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceUpdate(kustomizationReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster kustomization entry, name : %s", kustomizationName))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			kustomizationReq := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest{
				Kustomization: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     kustomizationDataFromServer.meta,
					Spec: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec{
						AtomicSpec: kustomizationDataFromServer.atomicSpec,
					},
				},
			}

			_, err = config.TMCConnection.ClusterGroupKustomizationResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceUpdate(kustomizationReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster group kustomization entry, name : %s", kustomizationName))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	log.Printf("[INFO] kustomization update successful")

	return resourceKustomizationRead(ctx, d, m)
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

	log.Printf("[INFO] updating kustomization meta data")

	return true
}

func updateCheckForSpec(d *schema.ResourceData, atomicSpec *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec, scope commonscope.Scope) bool {
	if !spec.HasSpecChanged(d) {
		return false
	}

	var kustomizationSpec *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec

	switch scope {
	case commonscope.ClusterScope:
		kustomizationSpec = spec.ConstructSpecForClusterScope(d)
	case commonscope.ClusterGroupScope:
		clusterGroupScopeSpec := spec.ConstructSpecForClusterGroupScope(d)
		kustomizationSpec = clusterGroupScopeSpec.AtomicSpec
	}

	atomicSpec.Source = kustomizationSpec.Source
	atomicSpec.Path = kustomizationSpec.Path
	atomicSpec.Prune = kustomizationSpec.Prune
	atomicSpec.Interval = kustomizationSpec.Interval
	atomicSpec.TargetNamespace = kustomizationSpec.TargetNamespace

	log.Printf("[INFO] updating kustomization spec")

	return true
}
