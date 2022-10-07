/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/integration"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/nodepools"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/credential"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/iampolicy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/namespace"
	custompolicy "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom"
	custompolicyresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom/resource"
	securitypolicy "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security"
	securitypolicyresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security/resource"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

// Provider for Tanzu Mission Control resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			cluster.ResourceName:        cluster.ResourceTMCCluster(),
			workspace.ResourceName:      workspace.ResourceWorkspace(),
			namespace.ResourceName:      namespace.ResourceNamespace(),
			clustergroup.ResourceName:   clustergroup.ResourceClusterGroup(),
			nodepools.ResourceName:      nodepools.ResourceNodePool(),
			iampolicy.ResourceName:      iampolicy.ResourceIAMPolicy(),
			custompolicy.ResourceName:   custompolicyresource.ResourceCustomPolicy(),
			securitypolicy.ResourceName: securitypolicyresource.ResourceSecurityPolicy(),
			credential.ResourceName:     credential.ResourceCredential(),
			integration.ResourceName:    integration.ResourceIntegration(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			cluster.ResourceName:      cluster.DataSourceTMCCluster(),
			workspace.ResourceName:    workspace.DataSourceWorkspace(),
			namespace.ResourceName:    namespace.DataSourceNamespace(),
			clustergroup.ResourceName: clustergroup.DataSourceClusterGroup(),
			nodepools.ResourceName:    nodepools.DataSourceClusterNodePool(),
			credential.ResourceName:   credential.DataSourceCredential(),
			integration.ResourceName:  integration.DataSourceIntegration(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
}
