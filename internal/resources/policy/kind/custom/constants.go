/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindcustom

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom/recipe"
)

const (
	ResourceName = "tanzu-mission-control_custom_policy"
	typePolicy   = "custom-policy" // Type of Policy as defined in API
)

// Allowed input recipes.
const (
	UnknownRecipe                     Recipe = policy.UnknownRecipe
	TMCBlockNodeportServiceRecipe     Recipe = reciperesource.TMCBlockNodeportServiceKey
	TMCBlockResourcesRecipe           Recipe = reciperesource.TMCBlockResourcesKey
	TMCBlockRolebindingSubjectsRecipe Recipe = reciperesource.TMCBlockRolebindingSubjectsKey
	TMCExternalIPSRecipe              Recipe = reciperesource.TMCExternalIPSKey
	TMCHTTPSIngressRecipe             Recipe = reciperesource.TMCHTTPSIngressKey
	TMCRequireLabelsRecipe            Recipe = reciperesource.TMCRequireLabelsKey
	TMCCustomRecipe                   Recipe = reciperesource.TMCCustomKey
)
