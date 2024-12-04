// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policykindquota

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota/recipe"
)

const (
	ResourceName = "tanzu-mission-control_namespace_quota_policy"
	typePolicy   = "namespace-quota-policy" // Type of Policy as defined in API
)

// Allowed input recipes.
const (
	UnknownRecipe Recipe = policy.UnknownRecipe
	CustomRecipe  Recipe = reciperesource.CustomKey
	SmallRecipe   Recipe = reciperesource.SmallKey
	MediumRecipe  Recipe = reciperesource.MediumKey
	LargeRecipe   Recipe = reciperesource.LargeKey
)
