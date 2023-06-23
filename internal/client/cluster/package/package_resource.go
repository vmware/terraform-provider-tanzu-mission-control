/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package packageclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tanzupackage "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiKind                            = "tanzupackage/metadatas"
	namespaces                         = "namespaces"
	packages                           = "packages"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
	queryParamKeyOrgID                 = "fullName.orgID"
)

// New creates a new secret resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for secret resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1ClusterPackageResourceServiceGet(fn *tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageFullName) (*tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse, error)

	ManageV1alpha1ClusterPackageResourceServiceList(req *tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSearchScope) (*tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageListPackagesResponse, error)
}

/*
ManageV1alpha1ClusterPackageResourceServiceGet gets a source secret.
*/
func (c *Client) ManageV1alpha1ClusterPackageResourceServiceGet(fn *tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageFullName) (*tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, namespaces, fn.NamespaceName, apiKind, fn.MetadataName, packages, fn.Name).AppendQueryParams(queryParams).String()
	packageResponse := &tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse{}
	err := c.Get(requestURL, packageResponse)

	return packageResponse, err
}

/*
ManageV1alpha1ClusterPackageResourceServiceList gets a source secret.
*/
func (c *Client) ManageV1alpha1ClusterPackageResourceServiceList(req *tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSearchScope) (*tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageListPackagesResponse, error) {
	queryParams := url.Values{}

	if req.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, req.ManagementClusterName)
	}

	if req.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, req.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, req.ClusterName, namespaces, req.NamespaceName, apiKind, req.MetadataName, packages).AppendQueryParams(queryParams).String()
	listPackageResponse := &tanzupackage.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageListPackagesResponse{}
	err := c.Get(requestURL, listPackageResponse)

	return listPackageResponse, err
}
