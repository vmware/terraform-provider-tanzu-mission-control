/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package sourcesecret

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
)

func initTestProvider(t *testing.T) *schema.Provider {
	testProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceName:              ResourceSourceSecret(),
			cluster.ResourceName:      cluster.ResourceTMCCluster(),
			clustergroup.ResourceName: clustergroup.ResourceClusterGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			ResourceName:              DataSourceSourcesecret(),
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
