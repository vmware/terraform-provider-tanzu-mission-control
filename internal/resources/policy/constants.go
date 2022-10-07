/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policy

const (
	NamespaceSelectorKey      = "namespace_selector"
	MatchExpressionsKey       = "match_expressions"
	KeyKey                    = "key"
	OperatorKey               = "operator"
	ValuesKey                 = "values"
	ScopeKey                  = "scope"
	clusterKey                = "cluster"
	clusterGroupKey           = "cluster_group"
	organizationKey           = "organization"
	SpecKey                   = "spec"
	NameKey                   = "name"
	InputKey                  = "input"
	RecipeVersionDefaultValue = "v1"
	UnknownRecipe             = ""
)

// Allowed scopes.
const (
	UnknownScope Scope = iota
	ClusterScope
	ClusterGroupScope
	OrganizationScope
)
