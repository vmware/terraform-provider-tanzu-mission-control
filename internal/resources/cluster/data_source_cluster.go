/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceTMCCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTMCClusterRead,
		Schema:      clusterSchema,
	}
}

func dataSourceTMCClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var (
		diags diag.Diagnostics
		resp  *clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse
		err   error
	)

	getClusterResourceRetryableFn := func() (retry bool, err error) {
		resp, err = config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(constructFullname(d))
		if err != nil {
			if clienterrors.IsNotFoundError(err) {
				d.SetId("")
				return false, nil
			}

			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster entry, name : %s", d.Get(clusterNameKey))
		}

		d.SetId(resp.Cluster.Meta.UID)

		if clustermodel.NewVmwareTanzuManageV1alpha1ClusterPhase(clustermodel.VmwareTanzuManageV1alpha1ClusterPhaseREADY) != resp.Cluster.Status.Phase {
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
		_, err = getClusterResourceRetryableFn()
	case "":
		fallthrough
	case "default":
		_, attachClusterWithKubeconfig := d.GetOk(attachClusterKey)
		fn := constructFullname(d)

		if fn.ManagementClusterName == attachedValue && !attachClusterWithKubeconfig {
			_, err = getClusterResourceRetryableFn()
			break
		}

		timeoutValueData = "3m"

		fallthrough
	default:
		timeoutDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr != nil {
			log.Printf("[INFO] unable to prase the duration value for the key %s. Defaulting to 3 minutes(3m)"+
				" Please refer to 'https://pkg.go.dev/time#ParseDuration' for providing the right value", waitKey)

			timeoutDuration = 3 * time.Minute
		}

		_, err = helper.RetryUntilTimeout(getClusterResourceRetryableFn, 10*time.Second, timeoutDuration)
	}

	if err != nil || resp == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster entry, name : %s", d.Get(clusterNameKey)))
	}

	// always run
	d.SetId(resp.Cluster.Meta.UID)

	status := map[string]interface{}{
		"type":                  resp.Cluster.Status.Type,
		"phase":                 resp.Cluster.Status.Phase,
		"health":                resp.Cluster.Status.Health,
		"k8s_version":           resp.Cluster.Status.KubeServerVersion,
		"node_count":            resp.Cluster.Status.NodeCount,
		"k8s_provider_type":     resp.Cluster.Status.KubernetesProvider.Type,
		"k8s_provider_version":  resp.Cluster.Status.KubernetesProvider.Version,
		"infra_provider":        resp.Cluster.Status.InfrastructureProvider,
		"infra_provider_region": resp.Cluster.Status.InfrastructureProviderRegion,
	}

	if resp.Cluster.FullName.ManagementClusterName == attachedValue && resp.Cluster.Status.InstallerLink != "" {
		status["installer_link"] = resp.Cluster.Status.InstallerLink
		status["execution_cmd"] = fmt.Sprintf("kubectl create -f '%s'", resp.Cluster.Status.InstallerLink)
	}

	if err := d.Set(StatusKey, status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Cluster.Meta)); err != nil {
		return diag.FromErr(err)
	}

	clusterSpec := constructSpec(d)

	switch {
	case resp.Cluster.Spec.TkgAws != nil:
		resp.Cluster.Spec.TkgAws.Topology.NodePools = clusterSpec.TkgAws.Topology.NodePools
	case resp.Cluster.Spec.TkgServiceVsphere != nil:
		resp.Cluster.Spec.TkgServiceVsphere.Topology.NodePools = clusterSpec.TkgServiceVsphere.Topology.NodePools
	case resp.Cluster.Spec.TkgVsphere != nil:
		resp.Cluster.Spec.TkgVsphere.Topology.NodePools = clusterSpec.TkgVsphere.Topology.NodePools
	}

	if err := d.Set(SpecKey, flattenSpec(resp.Cluster.Spec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
