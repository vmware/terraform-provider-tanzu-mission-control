/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindmutation

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/mutation/recipe"
)

const (
	ResourceName             = "tanzu-mission-control_mutation_policy"
	typePolicy               = "mutation-policy" // Type of Policy as defined in API
	UnknownRecipe     Recipe = policy.UnknownRecipe
	PodSecurityRecipe Recipe = recipe.PodSecurityKey
	LabelRecipe       Recipe = recipe.LabelKey
	AnnotationRecipe  Recipe = recipe.AnnotationKey
)
