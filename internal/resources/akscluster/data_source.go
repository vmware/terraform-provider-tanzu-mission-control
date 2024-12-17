// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package akscluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
)

// getDataSourceSchema creates a data source version of the resource schema.
func getDataSourceSchema() map[string]*schema.Schema {
	ds := make(map[string]*schema.Schema, len(ClusterSchema))

	for k, v := range ClusterSchema {
		dv := v
		// make cluster 'spec' field optional
		if k == clusterSpecKey {
			// Create a new copy of the struct
			newSpec := &schema.Schema{
				Type:        v.Type,
				Description: v.Description,
				Required:    v.Required,
				MaxItems:    v.MaxItems,
				Elem:        v.Elem,
			}
			newSpec.Required = false
			newSpec.Optional = true

			dv = newSpec
		}

		ds[k] = dv
	}

	return ds
}

func DataSourceTMCAKSCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTMCAKSClusterRead,
		Schema:      getDataSourceSchema(),
	}
}

func dataSourceTMCAKSClusterRead(ctx context.Context, data *schema.ResourceData, config interface{}) diag.Diagnostics {
	tc, ok := config.(authctx.TanzuContext)
	if !ok {
		return diag.Errorf("error while retrieving Tanzu auth config")
	}

	clusterResp, nodepoolResp, err := getClusterAndNodepools(ctx, data, tc.TMCConnection)

	// The cluster does not exist it will be removed from any state.
	if clienterrors.IsNotFoundError(err) {
		_ = schema.RemoveFromState(data, nil)
		return diag.Diagnostics{}
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if clusterResp == nil || nodepoolResp == nil {
		return diag.FromErr(errors.Errorf("Unable to get Tanzu Mission Control AKS cluster entry, name : %s", data.Get(NameKey)))
	}

	if stateErr := setResourceState(data, clusterResp.AksCluster, nodepoolResp.Nodepools); stateErr != nil {
		return diag.FromErr(stateErr)
	}

	// load kubeconfig data
	if err := pollForKubeConfig(ctx, data, tc.TMCConnection, getPollInterval(ctx)); err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}

func getClusterAndNodepools(_ context.Context, data *schema.ResourceData, client *client.TanzuMissionControl) (*models.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse, *models.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse, error) {
	fn := extractClusterFullName(data)
	clusterResp, err := client.AKSClusterResourceService.AksClusterResourceServiceGet(fn)

	if clienterrors.IsNotFoundError(err) {
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control AKS cluster entry, name : %s", data.Get(NameKey))
	}

	nodepoolResp, err := client.AKSNodePoolResourceService.AksNodePoolResourceServiceList(fn)
	if clienterrors.IsNotFoundError(err) {
		return clusterResp, nodepoolResp, nil
	}

	if err != nil {
		return nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control AKS nodepools for cluster %s", data.Get(NameKey))
	}

	return clusterResp, nodepoolResp, err
}
