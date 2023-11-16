//go:build targetlocation
// +build targetlocation

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationtests

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	credentialres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/credential"
	targetlocationres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/targetlocation"
)

const (
	tmcManagedCredentialsEnv = "TMC_MANAGED_CREDENTIALS_NAME" // #nosec G101
	azureCredentialsNameEnv  = "AZURE_CREDENTIALS_NAME"       // #nosec G101
)

var (
	testScopeHelper = commonscope.NewScopeHelperResources()
)

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			targetlocationres.ResourceName: targetlocationres.ResourceTargetLocation(),
			dataprotection.ResourceName:    dataprotection.ResourceEnableDataProtection(),
			cluster.ResourceName:           cluster.ResourceTMCCluster(),
			clustergroup.ResourceName:      clustergroup.ResourceClusterGroup(),
			credentialres.ResourceName:     credentialres.ResourceCredential(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			targetlocationres.ResourceName: targetlocationres.DataSourceTargetLocations(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}

	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
