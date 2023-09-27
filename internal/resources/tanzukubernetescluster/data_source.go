/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkc

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
	tkcmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	tkcstatus "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/status"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceTMCTKCCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceTMCTKCClusterRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: clusterSchema,
	}
}

func dataSourceTMCTKCClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var (
		diags diag.Diagnostics
		resp  *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse
		err   error
	)

	clusterFn := constructFullname(d)
	getTkcClusterResourceRetryableFn := func() (retry bool, err error) {
		resp, err = config.TMCConnection.TKCClusterResourceService.TanzuKubernetesClusterResourceServiceGet(clusterFn)
		if err != nil {
			if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
				_ = schema.RemoveFromState(d, m)
				return false, nil
			}

			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control TKC cluster entry, name : %s", d.Get(NameKey))
		}

		d.SetId(resp.TanzuKubernetesCluster.Meta.UID)

		if ctx.Value(contextMethodKey{}) == "create" &&
			resp.TanzuKubernetesCluster.Status.Phase != nil &&
			*resp.TanzuKubernetesCluster.Status.Phase != tkcstatus.VmwareTanzuManageV1alpha1CommonClusterStatusPhaseREADY {
			if c, ok := resp.TanzuKubernetesCluster.Status.Conditions[readyCondition]; ok &&
				c.Severity != nil &&
				*c.Severity == tkcstatus.VmwareTanzuCoreV1alpha1StatusConditionSeverityERROR {
				return false, errors.Errorf("Cluster %s creation failed due to %s, %s", d.Get(NameKey), c.Reason, c.Message)
			}

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
		_, err = getTkcClusterResourceRetryableFn()
	case "":
		fallthrough
	case "default":
		timeoutValueData = "5m"

		fallthrough
	default:
		timeoutDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr != nil {
			defaultTimeoutInMinutes := defaultTimeout / time.Minute
			log.Printf("[INFO] unable to parse the duration value for the key %s. Defaulting to %d minutes(%dm)"+
				" Please refer to 'https://pkg.go.dev/time#ParseDuration' for providing the right value", waitKey, defaultTimeoutInMinutes, defaultTimeoutInMinutes)

			timeoutDuration = defaultTimeout
		}

		_, err = helper.RetryUntilTimeout(getTkcClusterResourceRetryableFn, 10*time.Second, timeoutDuration)
	}

	if err != nil || resp == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control TKC cluster entry, name : %s", d.Get(NameKey)))
	}

	// always run
	d.SetId(resp.TanzuKubernetesCluster.Meta.UID)

	err = setResourceData(d, resp.TanzuKubernetesCluster)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to set resource data for cluster read"))
	}

	return diags
}

func setResourceData(d *schema.ResourceData, tkcCluster *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster) error {
	status := map[string]interface{}{
		// TODO: add condition
		"phase": tkcCluster.Status.Phase,
	}

	if err := d.Set(StatusKey, status); err != nil {
		return errors.Wrapf(err, "Failed to set status for the cluster %s", tkcCluster.FullName.Name)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(tkcCluster.Meta)); err != nil {
		return errors.Wrap(err, "Failed to set meta for the cluster")
	}

	constructTkcClusterSpec(d)

	if err := d.Set(specKey, flattenClusterSpec(tkcCluster.Spec)); err != nil {
		return errors.Wrapf(err, "Failed to set the spec for cluster %s", tkcCluster.FullName.Name)
	}

	return nil
}
