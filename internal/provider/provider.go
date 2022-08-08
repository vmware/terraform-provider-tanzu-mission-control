/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/nodepools"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/namespace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

// Provider for Tanzu Mission Control resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			cluster.ResourceName:      cluster.ResourceTMCCluster(),
			workspace.ResourceName:    workspace.ResourceWorkspace(),
			namespace.ResourceName:    namespace.ResourceNamespace(),
			clustergroup.ResourceName: clustergroup.ResourceClusterGroup(),
			nodepools.ResourceName:    nodepools.ResourceNodePool(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			cluster.ResourceName:      cluster.DataSourceTMCCluster(),
			workspace.ResourceName:    workspace.DataSourceWorkspace(),
			namespace.ResourceName:    namespace.DataSourceNamespace(),
			clustergroup.ResourceName: clustergroup.DataSourceClusterGroup(),
			nodepools.ResourceName:    nodepools.DataSourceClusterNodePool(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
}
