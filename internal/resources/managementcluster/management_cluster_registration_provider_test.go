package managementcluster

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
)

func initTestProvider(t *testing.T) *schema.Provider {
	testProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceName: ResourceManagementClusterRegistration(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			ResourceName: DataSourceManagementClusterRegistration(),
		},
		ConfigureContextFunc: getConfigureContextFunc(),
	}
	if err := testProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testProvider
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_POLICY_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
