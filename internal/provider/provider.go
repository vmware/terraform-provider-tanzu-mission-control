/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/cluster"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/clustergroup"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/namespace"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/workspace"
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
		},
		DataSourcesMap: map[string]*schema.Resource{
			cluster.ResourceName:      cluster.DataSourceTMCCluster(),
			workspace.ResourceName:    workspace.DataSourceWorkspace(),
			namespace.ResourceName:    namespace.DataSourceNamespace(),
			clustergroup.ResourceName: clustergroup.DataSourceClusterGroup(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
}
