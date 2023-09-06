/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackageclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tanzupackage "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackage"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiKind                            = "tanzupackage"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
)

// New creates a new repository resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for repository resource service API.
*/
type Client struct {
	*transport.Client
}

type ClientService interface {
	TanzuPackageResourceServiceList(request *tanzupackage.VmwareTanzuManageV1alpha1ClusterTanzupackageSearchScope) (*tanzupackage.VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse, error)
}

/*
RepositoryResourceServiceGet gets a repository.
*/
func (c *Client) TanzuPackageResourceServiceList(fn *tanzupackage.VmwareTanzuManageV1alpha1ClusterTanzupackageSearchScope) (*tanzupackage.VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind).AppendQueryParams(queryParams).String()
	pkgListResponse := &tanzupackage.VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse{}
	err := c.Get(requestURL, pkgListResponse)

	return pkgListResponse, err
}
