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
	ResourceName     = "tanzu-mission-control_custom_policy"
	typeDefaultValue = "custom-policy"
)

// Allowed input recipes.
const (
	unknownRecipe                     recipe = policy.UnknownRecipe
	tmcBlockNodeportServiceRecipe     recipe = reciperesource.TMCBlockNodeportServiceKey
	tmcBlockResourcesRecipe           recipe = reciperesource.TMCBlockResourcesKey
	tmcBlockRolebindingSubjectsRecipe recipe = reciperesource.TMCBlockRolebindingSubjectsKey
	tmcExternalIPSRecipe              recipe = reciperesource.TMCExternalIPSKey
	tmcHTTPSIngressRecipe             recipe = reciperesource.TMCHTTPSIngressKey
	tmcRequireLabelsRecipe            recipe = reciperesource.TMCRequireLabelsKey
)
