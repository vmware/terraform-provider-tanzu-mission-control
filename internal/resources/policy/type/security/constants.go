/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package security

import reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/type/security/recipe"

const (
	ResourceName              = "tanzu-mission-control_security_policy"
	nameKey                   = "name"
	scopeKey                  = "scope"
	clusterKey                = "cluster"
	clusterGroupKey           = "cluster_group"
	organizationKey           = "organization"
	specKey                   = "spec"
	inputKey                  = "input"
	typeDefaultValue          = "security-policy"
	recipeVersionDefaultValue = "v1"
)

// Allowed scopes.
const (
	clusterScope scope = iota
	clusterGroupScope
	organizationScope
)

// Allowed input recipes.
const (
	baselineRecipe recipe = reciperesource.BaselineKey
	customRecipe   recipe = reciperesource.CustomKey
	strictRecipe   recipe = reciperesource.StrictKey
)
