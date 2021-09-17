/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespaceclient

import (
	"fmt"
	"net/url"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"
	namespacemodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/namespace"
)

// New creates a new namespace resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for namespace resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1NamespaceResourceServiceCreate(request *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse, error)

	ManageV1alpha1NamespaceResourceServiceDelete(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) error

	ManageV1alpha1NamespaceResourceServiceGet(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse, error)

	ManageV1alpha1NamespaceResourceServiceUpdate(request *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse, error)
}

/*
ManageV1alpha1NamespaceResourceServiceCreate creates a Namespace.
*/
func (c *Client) ManageV1alpha1NamespaceResourceServiceCreate(request *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s", "v1alpha1/clusters", request.Namespace.FullName.ClusterName, "namespaces")
	namespaceResponse := &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse{}
	err := c.Create(requestURL, request, namespaceResponse)

	return namespaceResponse, err
}

/*
ManageV1alpha1NamespaceResourceServiceUpdate updates a Namespace.
*/
func (c *Client) ManageV1alpha1NamespaceResourceServiceUpdate(request *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse, error) {
	requestURL := fmt.Sprintf("%s/%s/%s/%s", "v1alpha1/clusters", request.Namespace.FullName.ClusterName, "namespaces", request.Namespace.FullName.Name)
	namespaceResponse := &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse{}
	err := c.Update(requestURL, request, namespaceResponse)

	return namespaceResponse, err
}

/*
ManageV1alpha1NamespaceResourceServiceDelete deletes a Namespace.
*/
func (c *Client) ManageV1alpha1NamespaceResourceServiceDelete(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s?%s", "v1alpha1/clusters", fn.ClusterName, "namespaces", fn.Name, queryParams.Encode())

	return c.Delete(requestURL)
}

/*
ManageV1alpha1NamespaceResourceServiceGet gets a namespace.
*/
func (c *Client) ManageV1alpha1NamespaceResourceServiceGet(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s?%s", "v1alpha1/clusters", fn.ClusterName, "namespaces", fn.Name, queryParams.Encode())
	namespaceResponse := &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse{}
	err := c.Get(requestURL, namespaceResponse)

	return namespaceResponse, err
}
