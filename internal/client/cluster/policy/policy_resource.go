// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policyclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiKind                            = "policies"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
)

// New creates a new cluster policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster policy resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1ClusterPolicyResourceService Client methods.
type ClientService interface {
	ManageV1alpha1ClusterPolicyResourceServiceCreate(request *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse, error)

	ManageV1alpha1ClusterPolicyResourceServiceDelete(fn *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) error

	ManageV1alpha1ClusterPolicyResourceServiceGet(fn *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyGetPolicyResponse, error)

	ManageV1alpha1ClusterPolicyResourceServiceUpdate(request *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse, error)
}

/*
ManageV1alpha1ClusterPolicyResourceServiceCreate creates a policy scoped to a cluster resource.
*/
func (p *Client) ManageV1alpha1ClusterPolicyResourceServiceCreate(request *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Policy.FullName.ClusterName, apiKind).String()
	policyClusterResponse := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse{}
	err := p.Create(requestURL, request, policyClusterResponse)

	return policyClusterResponse, err
}

/*
ManageV1alpha1ClusterPolicyResourceServiceDelete deletes a policy scoped to a cluster resource.
*/
func (p *Client) ManageV1alpha1ClusterPolicyResourceServiceDelete(fn *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
ManageV1alpha1ClusterPolicyResourceServiceGet gets a policy scoped to a cluster resource.
*/
func (p *Client) ManageV1alpha1ClusterPolicyResourceServiceGet(fn *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyGetPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	policyClusterResponse := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyGetPolicyResponse{}
	err := p.Get(requestURL, policyClusterResponse)

	return policyClusterResponse, err
}

/*
ManageV1alpha1ClusterPolicyResourceServiceUpdate updates overwrite a policy scoped to a cluster resource.
*/
func (p *Client) ManageV1alpha1ClusterPolicyResourceServiceUpdate(request *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) (*policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Policy.FullName.ClusterName, apiKind, request.Policy.FullName.Name).String()
	policyClusterResponse := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse{}
	err := p.Update(requestURL, request, policyClusterResponse)

	return policyClusterResponse, err
}
