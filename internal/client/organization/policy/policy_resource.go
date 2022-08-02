/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyorganizationclient

import (
	"net/url"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/helper"
	policyorganizationmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
)

const (
	apiVersionGroupAndKind = "v1alpha1/organization/policies"
	queryParamKeyOrgID     = "fullName.orgId"
)

// New creates a new organization policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for organization policy resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1OrganizationPolicyResourceService Client methods.
type ClientService interface {
	ManageV1alpha1OrganizationPolicyResourceServiceCreate(request *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse, error)

	ManageV1alpha1OrganizationPolicyResourceServiceDelete(fn *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) error

	ManageV1alpha1OrganizationPolicyResourceServiceGet(fn *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse, error)

	ManageV1alpha1OrganizationPolicyResourceServiceUpdate(request *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse, error)
}

/*
ManageV1alpha1OrganizationPolicyResourceServiceCreate creates a policy scoped to an organization resource.
*/
func (p *Client) ManageV1alpha1OrganizationPolicyResourceServiceCreate(request *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse, error) {
	policyOrganizationResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse{}
	err := p.Create(apiVersionGroupAndKind, request, policyOrganizationResponse)

	return policyOrganizationResponse, err
}

/*
ManageV1alpha1OrganizationPolicyResourceServiceDelete deletes a policy scoped to an organization resource.
*/
func (p *Client) ManageV1alpha1OrganizationPolicyResourceServiceDelete(fn *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) error {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionGroupAndKind, fn.Name).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
ManageV1alpha1OrganizationPolicyResourceServiceGet gets a policy scoped to an organization resource.
*/
func (p *Client) ManageV1alpha1OrganizationPolicyResourceServiceGet(fn *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionGroupAndKind, fn.Name).AppendQueryParams(queryParams).String()
	policyOrganizationResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse{}
	err := p.Get(requestURL, policyOrganizationResponse)

	return policyOrganizationResponse, err
}

/*
ManageV1alpha1OrganizationPolicyResourceServiceUpdate updates overwrite a policy scoped to an organization resource.
*/
func (p *Client) ManageV1alpha1OrganizationPolicyResourceServiceUpdate(request *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionGroupAndKind, request.Policy.FullName.Name).String()
	policyOrganizationResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse{}
	err := p.Update(requestURL, request, policyOrganizationResponse)

	return policyOrganizationResponse, err
}
