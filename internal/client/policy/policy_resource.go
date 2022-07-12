/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyclient

import (
	"fmt"
	"net/url"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/client/transport"
	policyclustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	policyworkspacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
)

// New creates a new policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for policy resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1ClusterPolicyResourceServiceCreate(request *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse, error)

	ManageV1alpha1ClusterPolicyResourceServiceDelete(fn *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) error

	ManageV1alpha1ClusterPolicyResourceServiceGet(fn *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyGetPolicyResponse, error)

	ManageV1alpha1ClusterPolicyResourceServiceUpdate(request *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse, error)

	ManageV1alpha1ClustergroupPolicyResourceServiceCreate(request *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse, error)

	ManageV1alpha1ClustergroupPolicyResourceServiceDelete(fn *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) error

	ManageV1alpha1ClustergroupPolicyResourceServiceGet(fn *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse, error)

	ManageV1alpha1ClustergroupPolicyResourceServiceUpdate(request *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse, error)

	ManageV1alpha1OrganizationPolicyResourceServiceCreate(request *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse, error)

	ManageV1alpha1OrganizationPolicyResourceServiceDelete(fn *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) error

	ManageV1alpha1OrganizationPolicyResourceServiceGet(fn *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse, error)

	ManageV1alpha1OrganizationPolicyResourceServiceUpdate(request *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse, error)

	ManageV1alpha1WorkspacePolicyResourceServiceCreate(request *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse, error)

	ManageV1alpha1WorkspacePolicyResourceServiceDelete(fn *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) error

	ManageV1alpha1WorkspacePolicyResourceServiceGet(fn *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse, error)

	ManageV1alpha1WorkspacePolicyResourceServiceUpdate(request *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse, error)
}

/*
ManageV1alpha1ClusterPolicyResourceServiceCreate creates a policy.
*/
func (p *Client) ManageV1alpha1ClusterPolicyResourceServiceCreate(request *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s", "v1alpha1/clusters", request.Policy.FullName.ClusterName, "policies")
	policyClusterResponse := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse{}
	err := p.Create(requestURL, request, policyClusterResponse)

	return policyClusterResponse, err
}

/*
ManageV1alpha1ClusterPolicyResourceServiceDelete deletes a policy.
*/
func (p *Client) ManageV1alpha1ClusterPolicyResourceServiceDelete(fn *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s?%s", "v1alpha1/clusters", fn.ClusterName, "policies", fn.Name, queryParams.Encode())

	return p.Delete(requestURL)
}

/*
ManageV1alpha1ClusterPolicyResourceServiceGet gets a policy.
*/
func (p *Client) ManageV1alpha1ClusterPolicyResourceServiceGet(fn *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyGetPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s?%s", "v1alpha1/clusters", fn.ClusterName, "policies", fn.Name, queryParams.Encode())
	policyClusterResponse := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyGetPolicyResponse{}
	err := p.Get(requestURL, policyClusterResponse)

	return policyClusterResponse, err
}

/*
ManageV1alpha1ClusterPolicyResourceServiceUpdate updates overwrite a policy.
*/
func (p *Client) ManageV1alpha1ClusterPolicyResourceServiceUpdate(request *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s/%s", "v1alpha1/clusters", request.Policy.FullName.ClusterName, "policies", request.Policy.FullName.Name)
	policyClusterResponse := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse{}
	err := p.Update(requestURL, request, policyClusterResponse)

	return policyClusterResponse, err
}

/*
ManageV1alpha1ClustergroupPolicyResourceServiceCreate creates a policy.
*/
func (p *Client) ManageV1alpha1ClustergroupPolicyResourceServiceCreate(request *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s", "v1alpha1/clustergroups", request.Policy.FullName.ClusterGroupName, "policies")
	policyClusterGroupResponse := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse{}
	err := p.Create(requestURL, request, policyClusterGroupResponse)

	return policyClusterGroupResponse, err
}

/*
ManageV1alpha1ClustergroupPolicyResourceServiceDelete deletes a policy.
*/
func (p *Client) ManageV1alpha1ClustergroupPolicyResourceServiceDelete(fn *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) error {
	requestURL := fmt.Sprintf("%s/%s/%s/%s", "v1alpha1/clustergroups", fn.ClusterGroupName, "policies", fn.Name)

	return p.Delete(requestURL)
}

/*
ManageV1alpha1ClustergroupPolicyResourceServiceGet gets a policy.
*/
func (p *Client) ManageV1alpha1ClustergroupPolicyResourceServiceGet(fn *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s/%s", "v1alpha1/clustergroups", fn.ClusterGroupName, "policies", fn.Name)
	policyClusterGroupResponse := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse{}
	err := p.Get(requestURL, policyClusterGroupResponse)

	return policyClusterGroupResponse, err
}

/*
ManageV1alpha1ClustergroupPolicyResourceServiceUpdate updates overwrite a policy.
*/
func (p *Client) ManageV1alpha1ClustergroupPolicyResourceServiceUpdate(request *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s/%s", "v1alpha1/clustergroups", request.Policy.FullName.ClusterGroupName, "policies", request.Policy.FullName.Name)
	policyClusterGroupResponse := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse{}
	err := p.Update(requestURL, request, policyClusterGroupResponse)

	return policyClusterGroupResponse, err
}

/*
ManageV1alpha1OrganizationPolicyResourceServiceCreate creates a policy.
*/
func (p *Client) ManageV1alpha1OrganizationPolicyResourceServiceCreate(request *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse, error) {
	policyOrganizationResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse{}
	err := p.Create("v1alpha1/organization/policies", request, policyOrganizationResponse)

	return policyOrganizationResponse, err
}

/*
ManageV1alpha1OrganizationPolicyResourceServiceDelete deletes a policy.
*/
func (p *Client) ManageV1alpha1OrganizationPolicyResourceServiceDelete(fn *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) error {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams["fullName.orgId"] = []string{fn.OrgID}
	}

	requestURL := fmt.Sprintf("%s/%s?%s", "v1alpha1/organization/policies", fn.Name, queryParams.Encode())

	return p.Delete(requestURL)
}

/*
ManageV1alpha1OrganizationPolicyResourceServiceGet gets a policy.
*/
func (p *Client) ManageV1alpha1OrganizationPolicyResourceServiceGet(fn *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams["fullName.orgId"] = []string{fn.OrgID}
	}

	requestURL := fmt.Sprintf("%s/%s?%s", "v1alpha1/organization/policies", fn.Name, queryParams.Encode())
	policyOrganizationResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse{}
	err := p.Get(requestURL, policyOrganizationResponse)

	return policyOrganizationResponse, err
}

/*
ManageV1alpha1OrganizationPolicyResourceServiceUpdate updates overwrite a policy.
*/
func (p *Client) ManageV1alpha1OrganizationPolicyResourceServiceUpdate(request *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) (*policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse, error) {
	requestURL := fmt.Sprintf("%s/%s", "v1alpha1/organization/policies", request.Policy.FullName.Name)
	policyOrganizationResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse{}
	err := p.Update(requestURL, request, policyOrganizationResponse)

	return policyOrganizationResponse, err
}

/*
ManageV1alpha1WorkspacePolicyResourceServiceCreate creates a policy.
*/
func (p *Client) ManageV1alpha1WorkspacePolicyResourceServiceCreate(request *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s", "v1alpha1/workspaces", request.Policy.FullName.WorkspaceName, "policies")
	policyWorkspaceResponse := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse{}
	err := p.Create(requestURL, request, policyWorkspaceResponse)

	return policyWorkspaceResponse, err
}

/*
ManageV1alpha1WorkspacePolicyResourceServiceDelete deletes a policy.
*/
func (p *Client) ManageV1alpha1WorkspacePolicyResourceServiceDelete(fn *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) error {
	requestURL := fmt.Sprintf("%s/%s/%s/%s", "v1alpha1/workspaces", fn.WorkspaceName, "policies", fn.Name)

	return p.Delete(requestURL)
}

/*
ManageV1alpha1WorkspacePolicyResourceServiceGet gets a policy.
*/
func (p *Client) ManageV1alpha1WorkspacePolicyResourceServiceGet(fn *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s/%s", "v1alpha1/workspaces", fn.WorkspaceName, "policies", fn.Name)
	policyWorkspaceResponse := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse{}
	err := p.Get(requestURL, policyWorkspaceResponse)

	return policyWorkspaceResponse, err
}

/*
ManageV1alpha1WorkspacePolicyResourceServiceUpdate updates overwrite a policy.
*/
func (p *Client) ManageV1alpha1WorkspacePolicyResourceServiceUpdate(request *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s/%s", "v1alpha1/workspaces", request.Policy.FullName.WorkspaceName, "policies", request.Policy.FullName.Name)
	policyWorkspaceResponse := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse{}
	err := p.Update(requestURL, request, policyWorkspaceResponse)

	return policyWorkspaceResponse, err
}
