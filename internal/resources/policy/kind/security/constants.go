/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindsecurity

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security/recipe"
)

const (
	ResourceName     = "tanzu-mission-control_security_policy"
	typeDefaultValue = "security-policy"
)

// Allowed input recipes.
const (
	UnknownRecipe  Recipe = policy.UnknownRecipe
	BaselineRecipe Recipe = reciperesource.BaselineKey
	CustomRecipe   Recipe = reciperesource.CustomKey
	StrictRecipe   Recipe = reciperesource.StrictKey
)
