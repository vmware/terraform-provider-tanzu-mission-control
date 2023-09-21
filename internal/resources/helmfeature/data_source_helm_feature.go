/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmfeature

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmfeature/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmfeature/status"
)

func DataSourceHelm() *schema.Resource {
	return &schema.Resource{
		Schema: helmSchema,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceHelmRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
	}
}

type dataFromServer struct {
	UID                     string
	meta                    *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
	clusterScopeStatus      *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus
	clusterGroupScopeStatus *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmStatus
}

func dataSourceHelmRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	scopedFullnameData := scope.ConstructScope(d)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control helm feature entry; Scope full name is empty")
	}

	helmDataFromServer, err := retrieveHelmFeatureDataFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return
		}

		return diag.FromErr(err)
	}

	// always run
	d.SetId(helmDataFromServer.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(helmDataFromServer.meta)); err != nil {
		return diag.FromErr(err)
	}

	var (
		flattenedStatus interface{}
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		flattenedStatus = status.FlattenStatusForClusterScope(helmDataFromServer.clusterScopeStatus)
	case commonscope.ClusterGroupScope:
		flattenedStatus = status.FlattenStatusForClusterGroupScope(helmDataFromServer.clusterGroupScopeStatus)
	}

	if err := d.Set(statusKey, flattenedStatus); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func retrieveHelmFeatureDataFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (*dataFromServer, error) {
	var helmDataFromServer = &dataFromServer{}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
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
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return helmDataFromServer, err
				}

				return helmDataFromServer, errors.Wrap(err, "Unable to get Tanzu Mission Control cluster helm feature entry")
			}

			if len(resp.Helms) == 0 {
				return helmDataFromServer, errors.New("Tanzu Mission Control no helm feature entry found")
			}

			helmRes := resp.Helms[0]

			scopedFullnameData.FullnameCluster = helmRes.FullName
			helmDataFromServer.UID = helmRes.Meta.UID
			helmDataFromServer.meta = helmRes.Meta
			helmDataFromServer.clusterScopeStatus = helmRes.Status
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupHelmResourceService.VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceList(
				&helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupHelmListHelmRequestParameters{
					SearchScope: &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmSearchScope{
						ClusterGroupName: scopedFullnameData.FullnameClusterGroup.ClusterGroupName,
					},
				},
			)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return helmDataFromServer, err
				}

				return helmDataFromServer, errors.Wrap(err, "Unable to get Tanzu Mission Control cluster helm feature entry")
			}

			if len(resp.Helms) == 0 {
				return helmDataFromServer, errors.New("Tanzu Mission Control no helm feature entry found")
			}

			helmRes := resp.Helms[0]

			scopedFullnameData.FullnameClusterGroup = helmRes.FullName
			helmDataFromServer.UID = helmRes.Meta.UID
			helmDataFromServer.meta = helmRes.Meta
			helmDataFromServer.clusterGroupScopeStatus = helmRes.Status
		}
	case commonscope.UnknownScope:
		return helmDataFromServer, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	fullName := scope.FlattenScope(scopedFullnameData)

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return helmDataFromServer, err
	}

	return helmDataFromServer, nil
}
