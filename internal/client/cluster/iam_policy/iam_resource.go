/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iamclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	clusteriammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy/cluster"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters:iam"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
)

// New creates a new cluster IAM policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster IAM policy resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1ClusterIAMPolicy Client methods.
type ClientService interface {
	ManageV1alpha1ClusterIAMPolicyGet(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) (*clusteriammodel.VmwareTanzuManageV1alpha1ClusterGetClusterIAMPolicyResponse, error)

	ManageV1alpha1ClusterIAMPolicyPatch(request *clusteriammodel.VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest) (*clusteriammodel.VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyResponse, error)

	ManageV1alpha1ClusterIAMPolicyUpdate(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*clusteriammodel.VmwareTanzuManageV1alpha1ClusterUpdateClusterIAMPolicyResponse, error)
}

/*
  ManageV1alpha1ClusterIAMPolicyGet gets all iam policies scoped to a cluster.
*/

func (c *Client) ManageV1alpha1ClusterIAMPolicyGet(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) (*clusteriammodel.VmwareTanzuManageV1alpha1ClusterGetClusterIAMPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).AppendQueryParams(queryParams).String()
	response := &clusteriammodel.VmwareTanzuManageV1alpha1ClusterGetClusterIAMPolicyResponse{}
	err := c.Get(requestURL, response)

	return response, err
}

/*
  ManageV1alpha1ClusterIAMPolicyPatch patches all iam policies scoped to a cluster.
*/

func (c *Client) ManageV1alpha1ClusterIAMPolicyPatch(request *clusteriammodel.VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest) (*clusteriammodel.VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyResponse, error) {
	response := &clusteriammodel.VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.FullName.Name).String()
	err := c.Patch(requestURL, request, response)

	return response, err
}

/*
  ManageV1alpha1ClusterIAMPolicyUpdate updates overwrites all iam policies scoped to a cluster, deletes if body is empty.
*/

func (c *Client) ManageV1alpha1ClusterIAMPolicyUpdate(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*clusteriammodel.VmwareTanzuManageV1alpha1ClusterUpdateClusterIAMPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).AppendQueryParams(queryParams).String()
	response := &clusteriammodel.VmwareTanzuManageV1alpha1ClusterUpdateClusterIAMPolicyResponse{}
	err := c.Update(requestURL, request, response)

	return response, err
}
