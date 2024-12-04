// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackagerepository

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
)

func initTestProvider(t *testing.T) *schema.Provider {
	testProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceName:         ResourcePackageRepository(),
			cluster.ResourceName: cluster.ResourceTMCCluster(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			ResourceName:         DataSourcePackageRepository(),
			cluster.ResourceName: cluster.DataSourceTMCCluster(),
		},
		ConfigureContextFunc: getConfigureContextFunc(),
	}
	if err := testProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testProvider
}

func MetaDataSourceAttributeCheck(dataSourceName, resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.description", resourceName, "meta.0.description"),
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key1", resourceName, "meta.0.labels.key1"),
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key2", resourceName, "meta.0.labels.key2"),
		resource.TestCheckResourceAttrSet(dataSourceName, "meta.0.uid"),
	}
}
