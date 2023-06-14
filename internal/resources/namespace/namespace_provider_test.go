/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespace

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
)

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
