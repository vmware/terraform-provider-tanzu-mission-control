/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceTMCEKSCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceTMCEKSClusterRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: clusterSchema,
	}
}

func DataSourceTMCEKSNodepool() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceTMCEKSNodepoolRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: nodepoolSchema,
	}
}

func dataSourceTMCEKSClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var (
		diags diag.Diagnostics
		resp  *eksmodel.VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse
		err   error
	)

	clusterFn := constructClusterFullname(d)
	getEksClusterResourceRetryableFn := func() (retry bool, err error) {
		resp, err = config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceGet(clusterFn)
		if err != nil {
			if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
				_ = schema.RemoveFromState(d, m)
				return false, nil
			}

			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey))
		}

		d.SetId(resp.EksCluster.Meta.UID)

		if ctx.Value(contextMethodKey{}) == "create" &&
			resp.EksCluster.Status.Phase != nil &&
			*resp.EksCluster.Status.Phase != eksmodel.VmwareTanzuManageV1alpha1EksclusterPhaseREADY {
			if c, ok := resp.EksCluster.Status.Conditions[readyCondition]; ok &&
				c.Severity != nil &&
				*c.Severity == eksmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityERROR {
				return false, errors.Errorf("Cluster %s creation failed due to %s, %s", d.Get(NameKey), c.Reason, c.Message)
			}

			log.Printf("[DEBUG] waiting for cluster(%s) to be in READY phase", constructClusterFullname(d).ToString())

			return true, nil
		}

		return false, nil
	}

	timeoutValueData, _ := d.Get(waitKey).(string)

	if ctx.Value(contextMethodKey{}) != "create" {
		timeoutValueData = "do_not_retry"
	}

	switch timeoutValueData {
	case "do_not_retry":
		_, err = getEksClusterResourceRetryableFn()
	case "":
		fallthrough
	case "default":
		timeoutValueData = "5m"

		fallthrough
	default:
		timeoutDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr != nil {
			log.Printf("[INFO] unable to prase the duration value for the key %s. Defaulting to 5 minutes(5m)"+
				" Please refer to 'https://pkg.go.dev/time#ParseDuration' for providing the right value", waitKey)

			timeoutDuration = defaultTimeout
		}

		_, err = helper.RetryUntilTimeout(getEksClusterResourceRetryableFn, 10*time.Second, timeoutDuration)
	}

	if err != nil || resp == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	// always run
	d.SetId(resp.EksCluster.Meta.UID)

	status := map[string]interface{}{
		// TODO: add condition
		"platform_version": resp.EksCluster.Status.PlatformVersion,
		"phase":            resp.EksCluster.Status.Phase,
	}

	if err := d.Set(StatusKey, status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.EksCluster.Meta)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func dataSourceTMCEKSNodepoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var (
		diags  diag.Diagnostics
		resp   *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse
		npresp *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse
		err    error
	)

	nodepoolFn := constructNodepoolFullname(d)
	getEksNodepoolResourceRetryableFn := func() (retry bool, err error) {
		resp, err = config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceGet(nodepoolFn)
		if err != nil {
			if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
				_ = schema.RemoveFromState(d, m)
				return false, nil
			}

			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS nodepool entry, name : %s", d.Get(NameKey))
		}

		d.SetId(resp.Nodepool.Meta.UID)

		if ctx.Value(contextMethodKey{}) == "create" &&
			resp.Nodepool.Status.Phase != nil &&
			*resp.Nodepool.Status.Phase != eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseREADY {
			if c, ok := resp.Nodepool.Status.Conditions[readyCondition]; ok &&
				c.Severity != nil &&
				*c.Severity == eksmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityERROR {
				return false, errors.Errorf("Nodepool %s creation failed due to %s, %s", d.Get(NameKey), c.Reason, c.Message)
			}

			log.Printf("[DEBUG] waiting for nodepool(%s) to be in READY phase", constructNodepoolFullname(d).ToString())

			return true, nil
		}

		return false, nil
	}

	timeoutValueData, _ := d.Get(waitKey).(string)

	if ctx.Value(contextMethodKey{}) != "create" {
		timeoutValueData = "do_not_retry"
	}

	switch timeoutValueData {
	case "do_not_retry":
		_, err = getEksNodepoolResourceRetryableFn()
	case "":
		fallthrough
	case "default":
		timeoutValueData = "5m"

		fallthrough
	default:
		timeoutDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr != nil {
			log.Printf("[INFO] unable to prase the duration value for the key %s. Defaulting to 5 minutes(5m)"+
				" Please refer to 'https://pkg.go.dev/time#ParseDuration' for providing the right value", waitKey)

			timeoutDuration = defaultTimeout
		}

		_, err = helper.RetryUntilTimeout(getEksNodepoolResourceRetryableFn, 10*time.Second, timeoutDuration)
	}

	if err != nil || resp == nil || npresp == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	// always run
	d.SetId(resp.Nodepool.Meta.UID)

	status := map[string]interface{}{
		// TODO: add condition
		"phase": resp.Nodepool.Status.Phase,
	}

	if err := d.Set(StatusKey, status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Nodepool.Meta)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

// Returns mapping of nodepool names to their positions in the array.
// This is needed because, we need to put the nodepools we receive from the
// API at the same location so that terraform can compute the diff properly.
//
// Note: setting nodepools as TypeSet won't work as it will use the passed
// hash function to check for change. This won't render much helpful diff.
func nodepoolPosMap(nps []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) map[string]int {
	ret := map[string]int{}
	for i, np := range nps {
		ret[np.Info.Name] = i
	}

	return ret
}
