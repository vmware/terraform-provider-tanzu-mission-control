/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipeclient

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"

	recipemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/recipe"
)

const (
	policyAPIPath = "v1alpha1/policy/types"
	recipeAPIPath = "recipes"
)

// New creates a new recipe resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for recipe resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	RecipeResourceServiceGet(fn *recipemodels.VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName) (*recipemodels.VmwareTanzuManageV1alpha1PolicyTypeRecipeData, error)
}

/*
RecipeResourceServiceGet gets a recipe.
*/
func (c *Client) RecipeResourceServiceGet(fn *recipemodels.VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName) (*recipemodels.VmwareTanzuManageV1alpha1PolicyTypeRecipeData, error) {
	requestURL := helper.ConstructRequestURL(policyAPIPath, fn.TypeName, recipeAPIPath, fn.Name).String()
	resp := &recipemodels.VmwareTanzuManageV1alpha1PolicyTypeRecipeData{}
	err := c.Get(requestURL, resp)

	return resp, err
}
