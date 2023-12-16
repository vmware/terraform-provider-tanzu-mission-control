/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	configModels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubeconfig"
)

func DataSourceTMCEKSCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceTMCEKSClusterRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: clusterSchema,
	}
}

func dataSourceTMCEKSClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var (
		diags  diag.Diagnostics
		resp   *eksmodel.VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse
		npresp *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse
		err    error
	)

	clusterFn := constructFullname(d)
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

		npresp, err = config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceList(clusterFn)
		if err != nil {
			if clienterrors.IsNotFoundError(err) {
				return false, nil
			}

			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS nodepools for cluster %s", d.Get(NameKey))
		}

		if ctx.Value(contextMethodKey{}) == "create" &&
			resp.EksCluster.Status.Phase != nil &&
			*resp.EksCluster.Status.Phase != eksmodel.VmwareTanzuManageV1alpha1EksclusterPhaseREADY {
			if c, ok := resp.EksCluster.Status.Conditions[readyCondition]; ok &&
				c.Severity != nil &&
				*c.Severity == eksmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityERROR {
				return false, errors.Errorf("Cluster %s creation failed due to %s, %s", d.Get(NameKey), c.Reason, c.Message)
			}

			log.Printf("[DEBUG] waiting for cluster(%s) to be in READY phase", constructFullname(d).ToString())

			return true, nil
		}

		if isWaitForKubeconfig(d) {
			clusFullName := &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
				Name:                  resp.EksCluster.Spec.AgentName,
				OrgID:                 clusterFn.OrgID,
				ManagementClusterName: "eks",
				ProvisionerName:       "eks",
			}
			clusterResp, err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(clusFullName)
			// nolint: wsl
			if err != nil {
				log.Printf("Unable to get Tanzu Mission Control cluster entry, name : %s, error :  %s", clusterFn.Name, err.Error())
				return true, err
			}

			mgmtClusterHealthy, err := isManagemetClusterHealthy(clusterResp)
			if err != nil {
				log.Printf("[DEBUG] waiting for cluster(%s) to be in Healthy status", clusterFn.Name)
				return true, nil
			} else if !mgmtClusterHealthy {
				log.Printf("[DEBUG] waiting for cluster(%s) to be in Healthy status", clusterFn.Name)
				return true, nil
			}

			fn := &configModels.VmwareTanzuManageV1alpha1ClusterFullName{
				ManagementClusterName: "eks",
				ProvisionerName:       "eks",
				Name:                  resp.EksCluster.Spec.AgentName,
			}
			resp, err := config.TMCConnection.KubeConfigResourceService.KubeconfigServiceGet(fn)
			// nolint: wsl
			if err != nil {
				log.Printf("Unable to get Tanzu Mission Control Kubeconfig entry, name : %s, error :  %s", fn.Name, err.Error())
				return true, err
			}

			if kubeConfigReady(err, resp) {
				if err = d.Set(kubeconfigKey, resp.Kubeconfig); err != nil {
					log.Printf("Failed to set Kubeconfig for cluster %s, error : %s", clusterFn.Name, err.Error())
					return false, err
				}
			} else {
				log.Printf("[DEBUG] waiting for cluster(%s)'s Kubeconfig to be in Ready status", clusterFn.Name)
				return true, nil
			}
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
		timeoutValueData = strconv.Itoa(minutesBasedDefaultTimeout) + "m"

		fallthrough
	default:
		timeoutDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr != nil {
			log.Printf("[INFO] unable to parse the duration value for the key %s. Defaulting to %d minutes(%dm)"+
				" Please refer to 'https://pkg.go.dev/time#ParseDuration' for providing the right value", waitKey, minutesBasedDefaultTimeout, minutesBasedDefaultTimeout)

			timeoutDuration = nanoSecondsBasedDefaultTimeout
		}

		_, err = helper.RetryUntilTimeout(getEksClusterResourceRetryableFn, 10*time.Second, timeoutDuration)
	}

	if err != nil || resp == nil || npresp == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	// always run
	d.SetId(resp.EksCluster.Meta.UID)

	err = setResourceData(d, resp.EksCluster, npresp.Nodepools)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to set resource data for cluster read"))
	}

	return diags
}

func isManagemetClusterHealthy(cluster *clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse) (bool, error) {
	if cluster == nil || cluster.Cluster == nil || cluster.Cluster.Status == nil || cluster.Cluster.Status.Health == nil {
		return false, errors.New("cluster data is invalid or nil")
	}

	if *cluster.Cluster.Status.Health == clustermodel.VmwareTanzuManageV1alpha1CommonClusterHealthHEALTHY {
		return true, nil
	}

	return false, nil
}

func kubeConfigReady(err error, resp *configModels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse) bool {
	return err == nil && *resp.Status == configModels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusREADY
}

func setResourceData(d *schema.ResourceData, eksCluster *eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster, remoteNodepools []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool) error {
	status := map[string]interface{}{
		// TODO: add condition
		"platform_version": eksCluster.Status.PlatformVersion,
		"phase":            eksCluster.Status.Phase,
	}

	if err := d.Set(StatusKey, status); err != nil {
		return errors.Wrapf(err, "Failed to set status for the cluster %s", eksCluster.FullName.Name)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(eksCluster.Meta)); err != nil {
		return errors.Wrap(err, "Failed to set meta for the cluster")
	}

	_, tfNodepools := constructEksClusterSpec(d)

	// see the explanation of this in the func doc of nodepoolPosMap
	npPosMap := nodepoolPosMap(tfNodepools)

	nodepools := make([]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition, len(tfNodepools))

	for _, np := range remoteNodepools {
		npDef := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
			Info: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{
				Description: np.Meta.Description,
				Name:        np.FullName.Name,
			},
			Spec: np.Spec,
		}

		if pos, ok := npPosMap[np.FullName.Name]; ok {
			nodepools[pos] = npDef
		} else {
			nodepools = append(nodepools, npDef)
		}
	}

	// check for deleted nodepools
	for i := range nodepools {
		if nodepools[i] == nil {
			nodepools[i] = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{}
		}
	}

	spec := flattenClusterSpec(eksCluster.Spec, nodepools)
	if err := d.Set(specKey, spec); err != nil {
		return errors.Wrapf(err, "Failed to set the spec for cluster %s", eksCluster.FullName.Name)
	}

	return nil
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
