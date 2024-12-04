// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package nodepools

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	nodepoolsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceClusterNodePool() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceClusterNodePoolRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: nodePoolSchema,
	}
}

func dataSourceClusterNodePoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var (
		resp *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse
		err  error
	)

	getNodepoolResourceRetryableFn := func() (retry bool, err error) {
		resp, err = config.TMCConnection.NodePoolResourceService.ManageV1alpha1ClusterNodePoolResourceServiceGet(constructFullName(d))
		if err != nil {
			if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
				_ = schema.RemoveFromState(d, m)

				if helper.IsRefreshState(ctx) {
					return false, nil
				}

				return true, nil
			}

			// refresh auth bearer token if it expired
			authctx.RefreshUserAuthContext(&config, clienterrors.IsUnauthorizedError, err)

			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control nodepool entry, name : %s", d.Get(nodePoolNameKey))
		}

		if resp.Nodepool.Status == nil || resp.Nodepool.Status.Phase == nil || resp.Nodepool.Status.Conditions == nil {
			return true, nil
		}

		// always run
		d.SetId(resp.Nodepool.FullName.Name + ":" + resp.Nodepool.FullName.ClusterName)

		return false, nil
	}
	timeoutValueData, _ := d.Get(waitKey).(string)

	timeoutDuration, _ := time.ParseDuration(timeoutValueData)

	_, err = helper.RetryUntilTimeout(getNodepoolResourceRetryableFn, 10*time.Second, timeoutDuration)

	if err != nil || resp == nil || resp.Nodepool == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get tanzu cluster node pool entry"))
	}

	var readyConditon nodepoolsmodel.VmwareTanzuCoreV1alpha1StatusCondition

	if resp.Nodepool.Status.Conditions != nil {
		if value, ok := resp.Nodepool.Status.Conditions[ready]; ok {
			readyConditon = value
		}
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
