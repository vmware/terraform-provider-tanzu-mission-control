/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepools

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/authctx"
	nodepoolsmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceClusterNodePool() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClusterNodePoolRead,
		Schema:      nodePoolSchema,
	}
}

func dataSourceClusterNodePoolRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var (
		diags diag.Diagnostics
		resp  *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse
		err   error
	)

	resp, err = config.TMCConnection.NodePoolResourceService.ManageV1alpha1ClusterNodePoolResourceServiceGet(constructFullName(d))

	if err != nil || resp == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get tanzu cluster node pool entry"))
	}

	// always run
	d.SetId(resp.Nodepool.FullName.Name + ":" + resp.Nodepool.FullName.ClusterName)

	var readyConditon nodepoolsmodel.VmwareTanzuCoreV1alpha1StatusCondition

	if value, ok := resp.Nodepool.Status.Conditions[ready]; ok {
		readyConditon = value
	}

	status := map[string]interface{}{
		"phase":    resp.Nodepool.Status.Phase,
		"type":     readyConditon.Type,
		"status":   readyConditon.Status,
		"severity": readyConditon.Severity,
	}

	if err := d.Set(statusKey, status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Nodepool.Meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(specKey, flattenSpec(resp.Nodepool.Spec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
