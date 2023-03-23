/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package respositorycredentialclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	respoistorycredential "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/repositorycredential/cluster"
)

const (
	apiVersionAndGroup                         = "v1alpha1/clusters"
	apiKind                                    = "fluxcd/sourcesecrets"
	queryParamKeyFullNameManagementClusterName = "fullName.managementClusterName"
	queryParamKeyFullNameProvisionerName       = "fullName.provisionerName"
)

// New creates a new repository credentials resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for repository credentials resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	RepositorycredentialResourceServiceCreate(request *respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRequest) (*respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse, error)

	RepositorycredentialResourceServiceDelete(fn *respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialFullName) error

	RepositorycredentialResourceServiceGet(fn *respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialFullName) (*respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdGetRepositorycredentialResponse, error)

	RepositorycredentialResourceServiceUpdate(request *respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRequest) (*respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse, error)
}

/*
RepositorycredentialResourceServiceCreate creates a repository credential.
*/
func (c *Client) RepositorycredentialResourceServiceCreate(request *respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRequest) (*respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Repositorycredential.FullName.ClusterName, apiKind).String()
	repocredResponse := &respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse{}
	err := c.Create(requestURL, request, repocredResponse)

	return repocredResponse, err
}

/*
RepositorycredentialResourceServiceDelete deletes a repository credential.
*/
func (c *Client) RepositorycredentialResourceServiceDelete(fn *respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams[queryParamKeyFullNameManagementClusterName] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams[queryParamKeyFullNameProvisionerName] = []string{fn.ProvisionerName}
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}

/*
RepositorycredentialResourceServiceGet gets a repository credential.
*/
func (c *Client) RepositorycredentialResourceServiceGet(fn *respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialFullName) (*respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdGetRepositorycredentialResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams[queryParamKeyFullNameManagementClusterName] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams[queryParamKeyFullNameManagementClusterName] = []string{fn.ProvisionerName}
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	repocredResponse := &respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdGetRepositorycredentialResponse{}
	err := c.Get(requestURL, repocredResponse)

	return repocredResponse, err
}

/*
RepositorycredentialResourceServiceUpdate updates overwrite a repository credential.
*/
func (c *Client) RepositorycredentialResourceServiceUpdate(request *respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRequest) (*respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Repositorycredential.FullName.ClusterName, apiKind, request.Repositorycredential.FullName.Name).String()
	repocredResponse := &respoistorycredential.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse{}
	err := c.Update(requestURL, request, repocredResponse)

	return repocredResponse, err
}
