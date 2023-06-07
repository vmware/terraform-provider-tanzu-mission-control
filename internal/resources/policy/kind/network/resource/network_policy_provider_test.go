/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package networkpolicyresource

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	policykindnetwork "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			policykindnetwork.ResourceName: ResourceNetworkPolicy(),
			workspace.ResourceName:         workspace.ResourceWorkspace(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
