/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package sourcesecretclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	sourcesecretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
)

const (
	apiVersionAndGroup         = "v1alpha1/clustergroups"
	apiKind                    = "fluxcd/sourcesecrets"
	queryParamKeyNamespaceName = "fullName.namespaceName"
	queryParamKeyOrgID         = "fullName.orgID"
)

// New creates a new cluster Flux CD source secret resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster group Flux CD source secret resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1ClustergroupFluxcdSourcesecretResourceService Client methods.
type ClientService interface {
	ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceCreate(request *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest) (*sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse, error)

	ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceDelete(fn *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName) error

	ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceGet(fn *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName) (*sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdGetSourceSecretResponse, error)

	ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceUpdate(request *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest) (*sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse, error)
}

/*
ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceCreate creates a Flux CD source secret scoped to a cluster group resource.
*/
func (p *Client) ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceCreate(request *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest) (*sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SourceSecret.FullName.ClusterGroupName, apiKind).String()
	fluxCDSourcesecretClusterGroupResponse := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse{}
	err := p.Create(requestURL, request, fluxCDSourcesecretClusterGroupResponse)

	return fluxCDSourcesecretClusterGroupResponse, err
}

/*
ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceDelete deletes a Flux CD source secret scoped to a cluster group resource.
*/
func (p *Client) ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceDelete(fn *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName) error {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceGet gets a Flux CD source secret scoped to a cluster group resource.
*/
func (p *Client) ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceGet(fn *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName) (*sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdGetSourceSecretResponse, error) {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	fluxCDSourcesecretClusterGroupResponse := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdGetSourceSecretResponse{}
	err := p.Get(requestURL, fluxCDSourcesecretClusterGroupResponse)

	return fluxCDSourcesecretClusterGroupResponse, err
}

/*
ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceUpdate updates overwrite a Flux CD source secret scoped to a cluster group resource.
*/
func (p *Client) ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceUpdate(request *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest) (*sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SourceSecret.FullName.ClusterGroupName, apiKind, request.SourceSecret.FullName.Name).String()
	fluxCDSourcesecretClusterGroupResponse := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse{}
	err := p.Update(requestURL, request, fluxCDSourcesecretClusterGroupResponse)

	return fluxCDSourcesecretClusterGroupResponse, err
}
