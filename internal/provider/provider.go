/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/cluster"
	clustergroup "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/clustergroup"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/namespace"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/workspace"
)

// Provider for Tanzu Mission Control resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			endpoint: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(authctx.ServerEndpointEnvVar, nil),
			},
			token: {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(authctx.CSPTokenEnvVar, nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			tmcCluster:      cluster.ResourceTMCCluster(),
			tmcWorkspace:    workspace.ResourceWorkspace(),
			tmcNamespace:    namespace.ResourceNamespace(),
			tmcClusterGroup: clustergroup.ResourceClusterGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			tmcCluster:   cluster.DataSourceTMCCluster(),
			tmcWorkspace: workspace.DataSourceTMCWorkspace(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := authctx.TanzuContext{CSPEndPoint: "console-stg.cloud.vmware.com"}

	config.ServerEndpoint = d.Get(endpoint).(string)
	config.Token = d.Get(token).(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (config.ServerEndpoint == "") || (config.Token == "") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "TANZU TMC credentials environment is not set",
			Detail:   fmt.Sprintf("Please set %s & %s to authenticate to TANZU TMC provider", authctx.ServerEndpointEnvVar, authctx.CSPTokenEnvVar),
		})

		return nil, diags
	}

	err := config.Setup()

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create connection to TMC",
			Detail:   fmt.Sprintf("Detailed error message: %s", err.Error()),
		})

		return nil, diags
	}

	return config, diags
}
