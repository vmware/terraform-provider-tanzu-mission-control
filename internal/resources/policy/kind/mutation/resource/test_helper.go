/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package mutationpolicyresource

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindmutation "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/mutation"
)

const (
	mutationPolicyResource    = policykindmutation.ResourceName
	mutationPolicyResourceVar = "test_mutation_policy"
	mutationPolicyNamePrefix  = "tf-mp-test"
	annotation                = "annotation"
	label                     = "label"
	podSecurity               = "pod-security"
)

type testAcceptanceConfig struct {
	Provider                   *schema.Provider
	MutationPolicyResource     string
	MutationPolicyResourceVar  string
	MutationPolicyResourceName string
	MutationPolicyName         string
	ScopeHelperResources       *policy.ScopeHelperResources
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_POLICY_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}
