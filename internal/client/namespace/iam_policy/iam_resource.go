/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iamnamespaceclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	namespaceiammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy/namespace"
	namespacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/namespace"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiKind                            = "namespaces:iam"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
)

// New creates a new namespace IAM policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for namespace IAM policy resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1ClusterNamespaceIAMPolicy Client methods.
type ClientService interface {
	ManageV1alpha1ClusterNamespaceIAMPolicyGet(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) (*namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceIAMPolicyResponse, error)

	ManageV1alpha1ClusterNamespaceIAMPolicyPatch(request *namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest) (*namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyResponse, error)

	ManageV1alpha1ClusterNamespaceIAMPolicyUpdate(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespaceUpdateNamespaceIAMPolicyResponse, error)
}

/*
  ManageV1alpha1ClusterNamespaceIAMPolicyGet gets all iam policies scoped to a namespace.
*/

func (c *Client) ManageV1alpha1ClusterNamespaceIAMPolicyGet(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) (*namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceIAMPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	response := &namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceIAMPolicyResponse{}
	err := c.Get(requestURL, response)

	return response, err
}

/*
  ManageV1alpha1ClusterNamespaceIAMPolicyPatch patches all iam policies scoped to a namespace.
*/

func (c *Client) ManageV1alpha1ClusterNamespaceIAMPolicyPatch(request *namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest) (*namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyResponse, error) {
	response := &namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.FullName.ClusterName, apiKind, request.FullName.Name).String()

	err := c.Patch(requestURL, request, response)

	return response, err
}

/*
  ManageV1alpha1ClusterNamespaceIAMPolicyUpdate updates all iam policies scoped to a namespace, deletes if body is empty.
*/

func (c *Client) ManageV1alpha1ClusterNamespaceIAMPolicyUpdate(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespaceUpdateNamespaceIAMPolicyResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	response := &namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespaceUpdateNamespaceIAMPolicyResponse{}
	err := c.Update(requestURL, request, response)

	return response, err
}
