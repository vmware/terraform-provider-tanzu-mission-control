/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindnetwork

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network/recipe"
)

const (
	ResourceName = "tanzu-mission-control_network_policy"
	typePolicy   = "network-policy" // Type of Policy as defined in API

	// Allowed input recipes.
	UnknownRecipe        Recipe = policy.UnknownRecipe
	AllowAllRecipe       Recipe = reciperesource.AllowAllKey
	AllowAllToPodsRecipe Recipe = reciperesource.AllowAllToPodsKey
	AllowAllEgressRecipe Recipe = reciperesource.AllowAllEgressKey
	DenyAllRecipe        Recipe = reciperesource.DenyAllKey
	DenyAllToPodsRecipe  Recipe = reciperesource.DenyAllToPodsKey
	DenyAllEgressRecipe  Recipe = reciperesource.DenyAllEgressKey
	CustomEgressRecipe   Recipe = reciperesource.CustomEgressKey
	CustomIngressRecipe  Recipe = reciperesource.CustomIngressKey
)
