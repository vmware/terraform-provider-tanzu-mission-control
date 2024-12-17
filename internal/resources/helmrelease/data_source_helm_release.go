// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmrelease

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
	releaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/status"
)

type dataFromServer struct {
	UID                     string
	meta                    *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
	atomicSpec              *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec
	clusterScopeStatus      *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus
	clusterGroupScopeStatus *releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus
}

func DataSourceHelmRelease() *schema.Resource {
	return &schema.Resource{
		Schema: getHelmReleaseSchema(true),
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceHelmReleaseRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
	}
}

func dataSourceHelmReleaseRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
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
		return diag.Errorf("Unable to get Tanzu Mission Control helm release entry; Scope full name is empty")
	}

	helmReleaseDataFromServer, err := retrieveHelmReleaseDataFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return diags
		}

		return diag.FromErr(err)
	}

	// always run
	d.SetId(helmReleaseDataFromServer.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(helmReleaseDataFromServer.meta)); err != nil {
		return diag.FromErr(err)
	}

	var (
		flattenedSpec   []interface{}
		flattenedStatus interface{}
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		flattenedSpec = spec.FlattenSpecForClusterScope(helmReleaseDataFromServer.atomicSpec)
		flattenedStatus = status.FlattenStatusForClusterScope(helmReleaseDataFromServer.clusterScopeStatus)
	case commonscope.ClusterGroupScope:
		clusterGroupScopeSpec := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec{
			AtomicSpec: helmReleaseDataFromServer.atomicSpec,
		}
		flattenedSpec = spec.FlattenSpecForClusterGroupScope(clusterGroupScopeSpec)
		flattenedStatus = status.FlattenStatusForClusterGroupScope(helmReleaseDataFromServer.clusterGroupScopeStatus)
	}

	if err := d.Set(spec.SpecKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(statusKey, flattenedStatus); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func retrieveHelmReleaseDataFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (*dataFromServer, error) {
	var helmReleaseDataFromServer = &dataFromServer{}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			resp, err := config.TMCConnection.ClusterHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceGet(scopedFullnameData.FullnameCluster)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return helmReleaseDataFromServer, err
				}

				return helmReleaseDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster helm release entry, name : %s", scopedFullnameData.FullnameCluster.Name)
			}

			scopedFullnameData.FullnameCluster = resp.Release.FullName
			helmReleaseDataFromServer.UID = resp.Release.Meta.UID
			helmReleaseDataFromServer.meta = resp.Release.Meta
			helmReleaseDataFromServer.atomicSpec = resp.Release.Spec
			helmReleaseDataFromServer.clusterScopeStatus = resp.Release.Status
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceGet(scopedFullnameData.FullnameClusterGroup)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return helmReleaseDataFromServer, err
				}

				return helmReleaseDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group helm release entry, name : %s", scopedFullnameData.FullnameClusterGroup.Name)
			}

			scopedFullnameData.FullnameClusterGroup = resp.Release.FullName
			helmReleaseDataFromServer.UID = resp.Release.Meta.UID
			helmReleaseDataFromServer.meta = resp.Release.Meta
			helmReleaseDataFromServer.atomicSpec = resp.Release.Spec.AtomicSpec
			helmReleaseDataFromServer.clusterGroupScopeStatus = resp.Release.Status
		}
	case commonscope.UnknownScope:
		return helmReleaseDataFromServer, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	fullName, name, namespace := scope.FlattenScope(scopedFullnameData)

	if err := d.Set(nameKey, name); err != nil {
		return helmReleaseDataFromServer, err
	}

	if err := d.Set(namespaceNameKey, namespace); err != nil {
		return helmReleaseDataFromServer, err
	}

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return helmReleaseDataFromServer, err
	}

	return helmReleaseDataFromServer, nil
}
