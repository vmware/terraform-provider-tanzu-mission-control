// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kubernetessecret

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
)

func initTestProvider(t *testing.T) *schema.Provider {
	testProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceName:              ResourceSecret(),
			cluster.ResourceName:      cluster.ResourceTMCCluster(),
			clustergroup.ResourceName: clustergroup.ResourceClusterGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			ResourceName:              DataSourceSecret(),
			cluster.ResourceName:      cluster.DataSourceTMCCluster(),
			clustergroup.ResourceName: clustergroup.DataSourceClusterGroup(),
		},
		ConfigureContextFunc: getConfigureContextFunc(),
	}
	if err := testProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testProvider
}

func MetaResourceAttributeCheck(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.uid"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.resource_version"),
	}
}

func MetaDataSourceAttributeCheck(dataSourceName, resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.description", resourceName, "meta.0.description"),
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key1", resourceName, "meta.0.labels.key1"),
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key2", resourceName, "meta.0.labels.key2"),
		resource.TestCheckResourceAttrSet(dataSourceName, "meta.0.uid"),
	}
}
