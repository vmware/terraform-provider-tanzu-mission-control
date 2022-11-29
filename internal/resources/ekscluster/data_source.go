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
		ReadContext: dataSourceTMCEKSClusterRead,
		Schema:      clusterSchema,
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

	getEksClusterResourceRetryableFn := func() (retry bool, err error) {
		resp, err = config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceGet(constructFullname(d))
		if err != nil {
			if clienterrors.IsNotFoundError(err) {
				d.SetId("")
				return false, nil
			}

			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey))
		}

		d.SetId(resp.EksCluster.Meta.UID)

		if eksmodel.NewVmwareTanzuManageV1alpha1EksclusterPhase(eksmodel.VmwareTanzuManageV1alpha1EksclusterPhaseREADY) != resp.EksCluster.Status.Phase {
			log.Printf("[DEBUG] waiting for cluster(%s) to be in READY phase", constructFullname(d).ToString())
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
		timeoutValueData = "3m"

		fallthrough
	default:
		timeoutDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr != nil {
			log.Printf("[INFO] unable to prase the duration value for the key %s. Defaulting to 3 minutes(3m)"+
				" Please refer to 'https://pkg.go.dev/time#ParseDuration' for providing the right value", waitKey)

			timeoutDuration = 3 * time.Minute
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

	clusterSpec := constructSpec(d)

	resp.EksCluster.Spec.NodePools = clusterSpec.NodePools

	if err := d.Set(specKey, flattenClusterSpec(resp.EksCluster.Spec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
