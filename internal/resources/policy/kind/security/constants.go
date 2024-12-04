// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policykindsecurity

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security/recipe"
)

const (
	ResourceName = "tanzu-mission-control_security_policy"
	typePolicy   = "security-policy" // Type of Policy as defined in API
)

// Allowed input recipes.
const (
	UnknownRecipe  Recipe = policy.UnknownRecipe
	BaselineRecipe Recipe = reciperesource.BaselineKey
	CustomRecipe   Recipe = reciperesource.CustomKey
	StrictRecipe   Recipe = reciperesource.StrictKey
)
