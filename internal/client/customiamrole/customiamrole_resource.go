/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrole

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	customiamrolemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/customiamrole"
)

const (
	iamRoleAPIVersionAndGroup = "v1alpha1/iam/roles"
)

// New creates a new cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for credentials resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	CustomIARoleResourceServiceCreate(request *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData) (*customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData, error)

	CustomIARoleResourceServiceUpdate(request *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData) (*customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData, error)

	CustomIARoleResourceServiceDelete(fn *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleFullName) error

	CustomIARoleResourceServiceGet(fn *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleFullName) (*customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData, error)
}

/*
CustomIARoleResourceServiceList gets a target location.
*/

func (c *Client) CustomIARoleResourceServiceGet(fullName *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleFullName) (*customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData, error) {
	requestURL := helper.ConstructRequestURL(iamRoleAPIVersionAndGroup, fullName.Name).String()
	resp := &customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData{}
	err := c.Get(requestURL, resp)

	return resp, err
}

/*
CustomIARoleResourceServiceCreate creates a target location.
*/
func (c *Client) CustomIARoleResourceServiceCreate(request *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData) (*customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData, error) {
	response := &customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData{}
	requestURL := helper.ConstructRequestURL(iamRoleAPIVersionAndGroup).String()
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
CustomIARoleResourceServiceUpdate updates a target location.
*/
func (c *Client) CustomIARoleResourceServiceUpdate(request *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData) (*customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData, error) {
	response := &customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData{}
	requestURL := helper.ConstructRequestURL(iamRoleAPIVersionAndGroup, request.Role.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}

/*
CustomIARoleResourceServiceDelete deletes a target location.
*/
func (c *Client) CustomIARoleResourceServiceDelete(fullName *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleFullName) error {
	requestURL := helper.ConstructRequestURL(iamRoleAPIVersionAndGroup, fullName.Name).String()

	return c.Delete(requestURL)
}
