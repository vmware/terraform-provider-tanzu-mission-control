/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package respositorycredentialclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	repositorycredentialclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/repositorycredential/cluster"
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
	SourceSecretResourceServiceCreate(request *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretRequest) (*repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse, error)

	SourceSecretResourceServiceDelete(fn *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) error

	SourceSecretResourceServiceGet(fn *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) (*repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretGetSourceSecretResponse, error)

	SourceSecretResourceServiceUpdate(request *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretRequest) (*repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse, error)
}

/*
SourceSecretResourceServiceCreate creates a repository credential.
*/
func (c *Client) SourceSecretResourceServiceCreate(request *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretRequest) (*repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SourceSecret.FullName.ClusterName, apiKind).String()
	repocredResponse := &repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse{}
	err := c.Create(requestURL, request, repocredResponse)

	return repocredResponse, err
}

/*
RepositorycredentialResourceServiceDelete deletes a repository credential.
*/
func (c *Client) SourceSecretResourceServiceDelete(fn *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) error {
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
func (c *Client) SourceSecretResourceServiceGet(fn *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) (*repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretGetSourceSecretResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams[queryParamKeyFullNameManagementClusterName] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams[queryParamKeyFullNameManagementClusterName] = []string{fn.ProvisionerName}
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	repocredResponse := &repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretGetSourceSecretResponse{}
	err := c.Get(requestURL, repocredResponse)

	return repocredResponse, err
}

/*
RepositorycredentialResourceServiceUpdate updates overwrite a repository credential.
*/
func (c *Client) SourceSecretResourceServiceUpdate(request *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretRequest) (*repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SourceSecret.FullName.ClusterName, apiKind, request.SourceSecret.FullName.Name).String()
	repocredResponse := &repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse{}
	err := c.Update(requestURL, request, repocredResponse)

	return repocredResponse, err
}
