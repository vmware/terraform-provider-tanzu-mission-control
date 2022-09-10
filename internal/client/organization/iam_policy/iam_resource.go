/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iamorganizationclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	organizationiammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy/organization"
	organizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/organization"
)

const (
	apiVersionGroupAndKind = "v1alpha1/organization:iam"
	queryParamKeyOrgID     = "fullName.orgId"
)

// New creates a new organization IAM policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
  Client for organization IAM policy resource service API.
*/

type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1OrganizationIAMPolicy Client methods.
type ClientService interface {
	ManageV1alpha1OrganizationIAMPolicyGet(fn *organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName) (*organizationiammodel.VmwareTanzuManageV1alpha1OrganizationGetOrganizationIAMPolicyResponse, error)

	ManageV1alpha1OrganizationIAMPolicyPatch(request *organizationiammodel.VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest) (*organizationiammodel.VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyResponse, error)

	ManageV1alpha1OrganizationIAMPolicyUpdate(fn *organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*organizationiammodel.VmwareTanzuManageV1alpha1OrganizationUpdateOrganizationIAMPolicyResponse, error)
}

/*
  ManageV1alpha1OrganizationIAMPolicyGet gets all iam policies scoped to an organization.
*/

func (c *Client) ManageV1alpha1OrganizationIAMPolicyGet(fn *organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName) (*organizationiammodel.VmwareTanzuManageV1alpha1OrganizationGetOrganizationIAMPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionGroupAndKind).AppendQueryParams(queryParams).String()
	response := &organizationiammodel.VmwareTanzuManageV1alpha1OrganizationGetOrganizationIAMPolicyResponse{}
	err := c.Get(requestURL, response)

	return response, err
}

/*
  ManageV1alpha1OrganizationIAMPolicyPatch patches all iam policies scoped to an organization.
*/

func (c *Client) ManageV1alpha1OrganizationIAMPolicyPatch(request *organizationiammodel.VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest) (*organizationiammodel.VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyResponse, error) {
	response := &organizationiammodel.VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionGroupAndKind).String()
	err := c.Patch(requestURL, request, response)

	return response, err
}

/*
  ManageV1alpha1OrganizationIAMPolicyUpdate updates all iam policies scoped to an organization, deletes if body is empty.
*/

func (c *Client) ManageV1alpha1OrganizationIAMPolicyUpdate(fn *organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*organizationiammodel.VmwareTanzuManageV1alpha1OrganizationUpdateOrganizationIAMPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionGroupAndKind).AppendQueryParams(queryParams).String()
	response := &organizationiammodel.VmwareTanzuManageV1alpha1OrganizationUpdateOrganizationIAMPolicyResponse{}
	err := c.Update(requestURL, request, response)

	return response, err
}
