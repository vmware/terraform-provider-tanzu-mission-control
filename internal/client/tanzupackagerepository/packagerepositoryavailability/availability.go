/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepositoryavailabilityclient

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	pkgrepository "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
)

const (
	apiVersionAndGroup = "v1alpha1/clusters"
	apiKind            = "tanzupackage/repositories:setavailability"
	namespaces         = "namespaces"
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

// ClientService is the interface for Client methods.
type ClientService interface {
	SetRepositoryAvailability(request *pkgrepository.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest) (*pkgrepository.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse, error)
}

// RepositoryHelperSetRepositoryAvailability enables or disable package repository for a cluster.
func (c *Client) SetRepositoryAvailability(request *pkgrepository.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest) (*pkgrepository.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.FullName.ClusterName, namespaces, request.FullName.NamespaceName, apiKind, request.FullName.Name).String()
	pkgRepositorySetAvailabilityResponse := &pkgrepository.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse{}
	err := c.Create(requestURL, request, pkgRepositorySetAvailabilityResponse)

	return pkgRepositorySetAvailabilityResponse, err
}
