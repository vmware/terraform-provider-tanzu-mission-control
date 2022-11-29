/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindimage

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image/recipe"
)

const (
	ResourceName     = "tanzu-mission-control_image_registry_policy"
	typeDefaultValue = "image-registry-policy"
)

// Allowed input recipes.
const (
	UnknownRecipe        Recipe = policy.UnknownRecipe
	AllowedNameTagRecipe Recipe = reciperesource.AllowedNameTagKey
	CustomRecipe         Recipe = reciperesource.CustomKey
	BlockLatestTagRecipe Recipe = reciperesource.BlockLatestTagKey
	RequireDigestRecipe  Recipe = reciperesource.RequireDigestKey
)
