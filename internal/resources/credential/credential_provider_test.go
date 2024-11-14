// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package credential

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
)

const (
	credentialResource = ResourceName
	// nolint:gosec
	credentialResourceVar = "test_credential"
	// nolint:gosec
	credentialDataSourceVar = "test_data_credential"
)

func initTestProvider(t *testing.T) *schema.Provider {
	resource := ResourceCredential()
	resource.UpdateContext = schema.NoopContext

	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceName: resource,
		},
		DataSourcesMap: map[string]*schema.Resource{
			ResourceName: DataSourceCredential(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
