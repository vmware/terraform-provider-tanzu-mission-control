/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicytemplateclient

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	iampolicytemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iampolicytemplate"
)

const (
	iamPolicyTemplateAPIVersionAndGroup = "v1alpha1/policy/templates"
)

// New creates a new cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for iam policy template resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	IAMPolicyTemplateResourceServiceCreate(request *iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData) (*iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error)

	IAMPolicyTemplateResourceServiceUpdate(request *iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData) (*iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error)

	IAMPolicyTemplateResourceServiceDelete(fn *iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName) error

	IAMPolicyTemplateResourceServiceGet(fn *iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName) (*iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error)
}

/*
IAMPolicyTemplateResourceServiceGet gets an iam policy template.
*/
func (c *Client) IAMPolicyTemplateResourceServiceGet(fullName *iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName) (*iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error) {
	requestURL := helper.ConstructRequestURL(iamPolicyTemplateAPIVersionAndGroup, fullName.Name).String()
	resp := &iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{}
	err := c.Get(requestURL, resp)

	return resp, err
}

/*
IAMPolicyTemplateResourceServiceCreate creates an iam policy template.
*/
func (c *Client) IAMPolicyTemplateResourceServiceCreate(request *iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData) (*iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error) {
	response := &iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{}
	requestURL := helper.ConstructRequestURL(iamPolicyTemplateAPIVersionAndGroup).String()
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
IAMPolicyTemplateResourceServiceUpdate updates an iam policy template.
*/
func (c *Client) IAMPolicyTemplateResourceServiceUpdate(request *iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData) (*iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error) {
	response := &iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{}
	requestURL := helper.ConstructRequestURL(iamPolicyTemplateAPIVersionAndGroup, request.Template.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}

/*
IAMPolicyTemplateResourceServiceDelete deletes an iam policy template.
*/
func (c *Client) IAMPolicyTemplateResourceServiceDelete(fullName *iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName) error {
	requestURL := helper.ConstructRequestURL(iamPolicyTemplateAPIVersionAndGroup, fullName.Name).String()

	return c.Delete(requestURL)
}
