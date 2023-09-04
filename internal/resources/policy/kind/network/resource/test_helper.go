/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package networkpolicyresource

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindNetwork "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network"
)

const (
	networkPolicyResource    = policykindNetwork.ResourceName
	networkPolicyResourceVar = "test_network_policy"
	networkPolicyNamePrefix  = "tf-np-test"
)

type testAcceptanceConfig struct {
	Provider                  *schema.Provider
	NetworkPolicyResource     string
	NetworkPolicyResourceVar  string
	NetworkPolicyResourceName string
	NetworkPolicyName         string
	ScopeHelperResources      *policy.ScopeHelperResources
}

// Function to set context containing different env variables based on the type of testing- Mock/Actual environment.
func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_POLICY_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
