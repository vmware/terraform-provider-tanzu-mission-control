package managementclusterregistration

import (
	"testing"

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
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
	if err := testProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testProvider
}
