/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package authctx

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ProviderAuthSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		endpoint: {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc(ServerEndpointEnvVar, nil),
		},
		vmwCloudEndpoint: {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc(VMWCloudEndpointEnvVar, defaultCSPEndpoint),
		},
		vmwCloudAPIToken: {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			DefaultFunc: schema.EnvDefaultFunc(VMWCloudAPITokenEnvVar, nil),
		},
	}
}

func ProviderConfigureContext(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := TanzuContext{}

	config.ServerEndpoint, _ = d.Get(endpoint).(string)
	config.VMWCloudEndPoint, _ = d.Get(vmwCloudEndpoint).(string)
	config.Token, _ = d.Get(vmwCloudAPIToken).(string)

	return setContext(config)
}

func setContext(config TanzuContext) (TanzuContext, diag.Diagnostics) {
	var diags diag.Diagnostics

	if (config.ServerEndpoint == "") || (config.Token == "") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Tanzu Mission Control credentials environment is not set",
			Detail:   fmt.Sprintf("Please set %s, %s & %s to authenticate to Tanzu Mission Control provider", ServerEndpointEnvVar, VMWCloudEndpointEnvVar, VMWCloudAPITokenEnvVar),
		})

		return config, diags
	}

	err := config.Setup()

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create connection to Tanzu Mission Control",
			Detail:   fmt.Sprintf("Detailed error message: %s", err.Error()),
		})

		return config, diags
	}

	return config, diags
}
