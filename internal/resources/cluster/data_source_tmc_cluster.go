/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tmccluster

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
	clustermodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/cluster"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/common"
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
	var diags diag.Diagnostics

	managementClusterName := d.Get(managementClusterNameKey).(string)
	provisionerName := d.Get(provisionerNameKey).(string)
	clusterName := d.Get(clusterNameKey).(string)

	fn := &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
		ManagementClusterName: managementClusterName,
		ProvisionerName:       provisionerName,
		Name:                  clusterName,
	}

	resp, err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(fn)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get tanzu TMC cluster entry, name : %s", clusterName))
	}

	// always run
	d.SetId(resp.Cluster.Meta.UID)

	switch managementClusterName {
	case "aws-hosted":
		// todo
	case "attached":
		attachData := []interface{}{
			map[string]interface{}{
				"execution_cmd": fmt.Sprintf("kubectl create -f '%s'", resp.Cluster.Status.InstallerLink),
			},
		}
		if err := d.Set("attach", attachData); err != nil {
			return diag.FromErr(err)
		}
	}

	status := map[string]interface{}{
		"phase":                 resp.Cluster.Status.Phase,
		"health":                resp.Cluster.Status.Health,
		"infra_provider":        resp.Cluster.Status.InfrastructureProvider,
		"infra_provider_region": resp.Cluster.Status.InfrastructureProviderRegion,
		"k8s_version":           resp.Cluster.Status.KubeServerVersion,
		"k8s_provider_type":     resp.Cluster.Status.KubernetesProvider.Type,
		"k8s_provider_version":  resp.Cluster.Status.KubernetesProvider.Version,
		"installer_link":        resp.Cluster.Status.InstallerLink,
	}

	if err := d.Set("status", status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Cluster.Meta)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
