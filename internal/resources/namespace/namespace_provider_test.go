/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespace

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/cluster"
)

const providerName = "tmc"

func initTestProvider(t *testing.T) *schema.Provider {
	testProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceName:         ResourceNamespace(),
			cluster.ResourceName: cluster.ResourceTMCCluster(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			ResourceName:         DataSourceNamespace(),
			cluster.ResourceName: cluster.DataSourceTMCCluster(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
	if err := testProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testProvider
}

func testPreCheck(t *testing.T) func() {
	return func() {
		for _, env := range []string{authctx.ServerEndpointEnvVar, authctx.CSPTokenEnvVar, authctx.CSPEndpointEnvVar} {
			require.NotEmpty(t, os.Getenv(env))
		}
	}
}

func getTestProviderFactories(provider *schema.Provider) map[string]func() (*schema.Provider, error) {
	//nolint:unparam
	return map[string]func() (*schema.Provider, error){
		providerName: func() (*schema.Provider, error) { return provider, nil },
	}
}
