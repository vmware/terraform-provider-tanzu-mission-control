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
	"github.com/pkg/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

func ProviderAuthSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		endpoint: {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc(ServerEndpointEnvVar, ""),
		},
		vmwCloudEndpoint: {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc(VMWCloudEndpointEnvVar, defaultCSPEndpoint),
		},

		vmwCloudAPIToken: {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			DefaultFunc: schema.EnvDefaultFunc(VMWCloudAPITokenEnvVar, ""),
		},

		selfManaged: selfManagedAuthSchema,

		insecureAllowUnverifiedSSL: {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			DefaultFunc: schema.EnvDefaultFunc(InsecureAllowUnverifiedSSLEnvVar, false),
		},
		clientAuthCertFile: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc(ClientAuthCertFileEnvVar, nil),
		},
		clientAuthKeyFile: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc(ClientAuthKeyFileEnvVar, nil),
		},
		caFile: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc(CAFileEnvVar, nil),
		},
		clientAuthCert: {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			DefaultFunc: schema.EnvDefaultFunc(ClientAuthCertEnvVar, nil),
		},
		clientAuthKey: {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			DefaultFunc: schema.EnvDefaultFunc(ClientAuthKeyEnvVar, nil),
		},
		caCert: {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			DefaultFunc: schema.EnvDefaultFunc(CACertEnvVar, nil),
		},
	}
}

var selfManagedAuthSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			oidcIssuer: {
				Type:     schema.TypeString,
				Optional: true,
			},
			smUsername: {
				Type:     schema.TypeString,
				Optional: true,
			},
			smPassword: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}

func ProviderConfigureContext(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := TanzuContext{
		TLSConfig:   &proxy.TLSConfig{},
		SelfManaged: false,
	}

	_, saasAuth := d.GetOk(vmwCloudAPIToken)
	_, smAuth := d.GetOk(selfManaged)

	config.ServerEndpoint, _ = d.Get(endpoint).(string)
	config.VMWCloudEndPoint, _ = d.Get(vmwCloudEndpoint).(string)
	config.Token, _ = d.Get(vmwCloudAPIToken).(string)

	smOIDCIssuer, _ := d.Get(helper.GetFirstElementOf(selfManaged, oidcIssuer)).(string)
	smUsrname, _ := d.Get(helper.GetFirstElementOf(selfManaged, smUsername)).(string)
	smPwd, _ := d.Get(helper.GetFirstElementOf(selfManaged, smPassword)).(string)

	config.TLSConfig.Insecure, _ = d.Get(insecureAllowUnverifiedSSL).(bool)
	config.TLSConfig.ClientAuthCertFile, _ = d.Get(clientAuthCertFile).(string)
	config.TLSConfig.ClientAuthKeyFile, _ = d.Get(clientAuthKeyFile).(string)
	config.TLSConfig.CaFile, _ = d.Get(caFile).(string)
	config.TLSConfig.ClientAuthCert, _ = d.Get(clientAuthCert).(string)
	config.TLSConfig.ClientAuthKey, _ = d.Get(clientAuthKey).(string)
	config.TLSConfig.CaCert, _ = d.Get(caCert).(string)

	switch {
	case saasAuth && smAuth:
		return nil, diag.FromErr(errors.New("Please configure authentication info either for SaaS or Self-Managed TMC flavour."))
	case saasAuth:
		if config.Token == "" {
			return nil, diag.FromErr(errors.Errorf("Please set %s", vmwCloudAPIToken))
		}
	case smAuth:
		if smOIDCIssuer == "" || smUsrname == "" || smPwd == "" {
			return nil, diag.FromErr(errors.New("Please set all the attributes under self_managed block"))
		}

		config.SelfManaged = smAuth
		config.VMWCloudEndPoint = smOIDCIssuer
		config.SMUsername = smUsrname
		config.Token = smPwd
	}

	return setContext(&config)
}

// The default transport is needed for mocking. The http mocking library used in testing
// can only intercept calls if they're made with the default transport.
func ProviderConfigureContextWithDefaultTransportForTesting(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := TanzuContext{
		TLSConfig: &proxy.TLSConfig{},
	}

	config.ServerEndpoint, _ = d.Get(endpoint).(string)
	config.VMWCloudEndPoint, _ = d.Get(vmwCloudEndpoint).(string)
	config.Token, _ = d.Get(vmwCloudAPIToken).(string)
	config.TLSConfig.Insecure, _ = d.Get(insecureAllowUnverifiedSSL).(bool)
	config.TLSConfig.ClientAuthCertFile, _ = d.Get(clientAuthCertFile).(string)
	config.TLSConfig.ClientAuthKeyFile, _ = d.Get(clientAuthKeyFile).(string)
	config.TLSConfig.CaFile, _ = d.Get(caFile).(string)
	config.TLSConfig.ClientAuthCert, _ = d.Get(clientAuthCert).(string)
	config.TLSConfig.ClientAuthKey, _ = d.Get(clientAuthKey).(string)
	config.TLSConfig.CaCert, _ = d.Get(caCert).(string)

	var err error
	config.TMCConnection, err = client.NewHTTPClient(config.TLSConfig)

	if err != nil {
		panic(err)
	}

	return setContextWithDefaultTransport(&config)
}

func setContext(config *TanzuContext) (TanzuContext, diag.Diagnostics) {
	var diags diag.Diagnostics

	if (config.ServerEndpoint == "") || (config.Token == "") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Tanzu Mission Control credentials environment is not set",
			Detail:   fmt.Sprintf("Please set %s & %s to authenticate to Tanzu Mission Control provider", ServerEndpointEnvVar, VMWCloudAPITokenEnvVar),
		})

		return *config, diags
	}

	err := config.Setup()

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create connection to Tanzu Mission Control",
			Detail:   fmt.Sprintf("Detailed error message: %s", err.Error()),
		})

		return *config, diags
	}

	return *config, diags
}

func setContextWithDefaultTransport(config *TanzuContext) (TanzuContext, diag.Diagnostics) {
	var diags diag.Diagnostics

	if (config.ServerEndpoint == "") || (config.Token == "") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Tanzu Mission Control credentials environment is not set",
			Detail:   fmt.Sprintf("Please set %s, %s & %s to authenticate to Tanzu Mission Control provider", ServerEndpointEnvVar, VMWCloudEndpointEnvVar, VMWCloudAPITokenEnvVar),
		})

		return *config, diags
	}

	err := config.SetupWithDefaultTransportForTesting()

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create connection to Tanzu Mission Control",
			Detail:   fmt.Sprintf("Detailed error message: %s", err.Error()),
		})

		return *config, diags
	}

	return *config, diags
}
