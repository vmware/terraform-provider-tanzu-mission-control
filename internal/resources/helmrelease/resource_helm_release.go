/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrelease

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
	releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
	releaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/status"
)

const (
	ResourceName = "tanzu-mission-control_helm_release"

	nameKey          = "name"
	namespaceNameKey = "namespace_name"
	statusKey        = "status"
	featureRefKey    = "feature_ref"
	disabledKey      = "Disabled"
)

func ResourceHelmRelease() *schema.Resource {
	return &schema.Resource{
		Schema:        getHelmReleaseSchema(false),
		CreateContext: resourceHelmReleaseCreate,
		ReadContext:   dataSourceHelmReleaseRead,
		UpdateContext: resourceHelmReleaseInPlaceUpdate,
		DeleteContext: resourceHelmReleaseDelete,
		CustomizeDiff: schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey, commonscope.ClusterGroupKey})),
	}
}

func getHelmReleaseSchema(isDataSource bool) map[string]*schema.Schema {
	var helmReleaseSchema = map[string]*schema.Schema{
		nameKey: {
			Type:        schema.TypeString,
			Description: "Name of the Repository.",
			Required:    true,
			ForceNew:    true,
		},
		namespaceNameKey: {
			Type:        schema.TypeString,
			Description: "Name of Namespace.",
			Required:    true,
			ForceNew:    true,
		},
		featureRefKey: {
			Type:        schema.TypeString,
			Description: "when specified, ensures clean up of this Terraform resource from the state file by creating a dependency on the Helm feature when the Helm feature is disabled",
			Optional:    true,
		},
		commonscope.ScopeKey: scope.ScopeSchema,
		common.MetaKey:       common.Meta,
		statusKey:            status.StatusSchema,
	}

	innerMap := map[string]*schema.Schema{
		spec.SpecKey: spec.SpecSchema,
	}

	for key, value := range innerMap {
		if isDataSource {
			helmReleaseSchema[key] = helper.UpdateDataSourceSchema(value)
		} else {
			helmReleaseSchema[key] = value
		}
	}

	return helmReleaseSchema
}

func resourceHelmReleaseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	helmReleaseName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read helm release name")
	}

	helmReleaseNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read helm release namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, helmReleaseName, helmReleaseNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control helm release entry; Scope full name is empty")
	}

	err := checkHelmFeature(config, scopedFullnameData)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster helm release entry, name : %s", helmReleaseName))
	}

	var (
		UID  string
		meta = common.ConstructMeta(d)
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			specVal, err := spec.ConstructSpecForClusterScope(d)

			if err != nil {
				return diag.FromErr(err)
			}

			helmReleaseReq := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest{
				Release: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
					Spec:     specVal,
				},
			}

			helmReleaseResponse, err := config.TMCConnection.ClusterHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceCreate(helmReleaseReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster helm release entry, name : %s", helmReleaseName))
			}

			UID = helmReleaseResponse.Release.Meta.UID
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			specVal, err := spec.ConstructSpecForClusterGroupScope(d)

			if err != nil {
				return diag.FromErr(err)
			}

			helmReleaseReq := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest{
				Release: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
					Spec:     specVal,
				},
			}

			helmReleaseResponse, err := config.TMCConnection.ClusterGroupHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceCreate(helmReleaseReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster group helm release entry, name : %s", helmReleaseName))
			}

			UID = helmReleaseResponse.Release.Meta.UID
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	// always run
	d.SetId(UID)

	return dataSourceHelmReleaseRead(ctx, d, m)
}

func resourceHelmReleaseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	helmReleaseName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read helm release name")
	}

	helmReleaseNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read helm release namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, helmReleaseName, helmReleaseNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to delete Tanzu Mission Control helm release entry; Scope full name is empty")
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			err := config.TMCConnection.ClusterHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceDelete(scopedFullnameData.FullnameCluster)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster helm release entry, name : %s", helmReleaseName))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			err := config.TMCConnection.ClusterGroupHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceDelete(scopedFullnameData.FullnameClusterGroup)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group helm release entry, name : %s", helmReleaseName))
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

func resourceHelmReleaseInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	helmReleaseName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read helm release name")
	}

	helmReleaseNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read helm release namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, helmReleaseName, helmReleaseNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to update Tanzu Mission Control helm release entry; Scope full name is empty")
	}

	err := checkHelmFeature(config, scopedFullnameData)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to Update, Tanzu Mission Control cluster helm release entry, name : %s", helmReleaseName))
	}

	helmReleaseDataFromServer, err := retrieveHelmReleaseDataFromServer(config, scopedFullnameData, d)
	if err != nil {
		log.Println("[ERROR] Unable to get data from server.")
		return diag.FromErr(err)
	}

	var updateAvailable bool

	if updateCheckForMeta(d, helmReleaseDataFromServer.meta) {
		updateAvailable = true
	}

	specCheck, err := updateCheckForSpec(d, helmReleaseDataFromServer.atomicSpec, scopedFullnameData.Scope)
	if err != nil {
		log.Println("[ERROR] Unable to check spec has been updated.")
		diag.FromErr(err)
	}

	if specCheck {
		updateAvailable = true
	}

	if !updateAvailable {
		log.Printf("[INFO] helm release update is not required")
		return
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			helmReleaseReq := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest{
				Release: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     helmReleaseDataFromServer.meta,
					Spec:     helmReleaseDataFromServer.atomicSpec,
				},
			}

			_, err = config.TMCConnection.ClusterHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceUpdate(helmReleaseReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster helm release entry, name : %s", helmReleaseName))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			helmReleaseReq := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest{
				Release: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     helmReleaseDataFromServer.meta,
					Spec: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec{
						AtomicSpec: helmReleaseDataFromServer.atomicSpec,
					},
				},
			}

			_, err = config.TMCConnection.ClusterGroupHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceUpdate(helmReleaseReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster group helm release entry, name : %s", helmReleaseName))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	log.Printf("[INFO] helm release update successful")

	return dataSourceHelmReleaseRead(ctx, d, m)
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

	log.Printf("[INFO] updating helm release meta data")

	return true
}

func updateCheckForSpec(d *schema.ResourceData, atomicSpec *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec, scope commonscope.Scope) (bool, error) {
	if !spec.HasSpecChanged(d) {
		return false, nil
	}

	var helmreleaseSpec *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec

	switch scope {
	case commonscope.ClusterScope:
		specVal, err := spec.ConstructSpecForClusterScope(d)
		if err != nil {
			return false, err
		}

		helmreleaseSpec = specVal
	case commonscope.ClusterGroupScope:
		specVal, err := spec.ConstructSpecForClusterGroupScope(d)
		if err != nil {
			return false, err
		}

		clusterGroupScopeSpec := specVal
		helmreleaseSpec = clusterGroupScopeSpec.AtomicSpec
	}

	atomicSpec.InlineConfiguration = helmreleaseSpec.InlineConfiguration
	atomicSpec.Interval = helmreleaseSpec.Interval
	atomicSpec.TargetNamespace = helmreleaseSpec.TargetNamespace
	atomicSpec.ChartRef.Chart = helmreleaseSpec.ChartRef.Chart
	atomicSpec.ChartRef.RepositoryName = helmreleaseSpec.ChartRef.RepositoryName
	atomicSpec.ChartRef.RepositoryNamespace = helmreleaseSpec.ChartRef.RepositoryNamespace
	atomicSpec.ChartRef.RepositoryType = helmreleaseSpec.ChartRef.RepositoryType
	atomicSpec.ChartRef.Version = helmreleaseSpec.ChartRef.Version

	log.Printf("[INFO] updating helm release spec")

	return true, nil
}

func checkHelmFeature(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname) error {
	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		resp, err := config.TMCConnection.ClusterHelmResourceService.VmwareTanzuManageV1alpha1ClusterHelmResourceServiceList(
			&helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmRequestParameters{
				SearchScope: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmSearchScope{
					ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
					ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
					ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
				},
			},
		)

		if err != nil {
			return err
		}

		if len(resp.Helms) == 0 {
			return errors.Errorf("Tanzu mission control helm feature is disable on cluster, name: %s", scopedFullnameData.FullnameCluster.ClusterName)
		}

		if _, ok := resp.Helms[0].Status.Conditions[disabledKey]; ok {
			return errors.Errorf("Tanzu mission control helm feature is disable on cluster, name: %s", scopedFullnameData.FullnameCluster.ClusterName)
		}
	case commonscope.ClusterGroupScope:
		resp, err := config.TMCConnection.ClusterGroupHelmResourceService.VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceList(
			&helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupHelmListHelmRequestParameters{
				SearchScope: &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmSearchScope{
					ClusterGroupName: scopedFullnameData.FullnameClusterGroup.ClusterGroupName,
				},
			},
		)

		if err != nil {
			return err
		}

		if len(resp.Helms) == 0 {
			return errors.Errorf("Tanzu mission control helm feature is disable on cluster group, name: %s", scopedFullnameData.FullnameCluster.ClusterName)
		}

		if resp.Helms[0].Status.Phase == nil {
			return errors.Errorf("Tanzu mission control helm feature is disable on cluster group, name: %s", scopedFullnameData.FullnameCluster.ClusterName)
		}
	}

	return nil
}
