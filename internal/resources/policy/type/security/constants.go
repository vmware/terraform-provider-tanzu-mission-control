/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package security

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/type/security/recipe"
)

const (
	ResourceName     = "tanzu-mission-control_security_policy"
	typeDefaultValue = "security-policy"
)

// Allowed input recipes.
const (
	unknownRecipe  recipe = policy.UnknownRecipe
	baselineRecipe recipe = reciperesource.BaselineKey
	customRecipe   recipe = reciperesource.CustomKey
	strictRecipe   recipe = reciperesource.StrictKey
)
