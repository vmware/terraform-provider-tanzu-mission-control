/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package mutationpolicyresource

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	policykindmutation "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/mutation"
)

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			policykindmutation.ResourceName: ResourceMutationPolicy(),
			cluster.ResourceName:            cluster.ResourceTMCCluster(),
			clustergroup.ResourceName:       clustergroup.ResourceClusterGroup(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
