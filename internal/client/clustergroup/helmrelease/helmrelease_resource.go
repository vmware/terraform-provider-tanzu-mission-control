/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmreleaseclustergroupclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	helmreleaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/clustergroup"
)

const (
	apiVersionAndGroup         = "v1alpha1/clustergroups"
	apiSubGroup                = "namespace"
	apiKind                    = "fluxcd/helm/releases"
	queryParamKeyNamespaceName = "fullName.namespaceName"
	queryParamKeyOrgID         = "fullName.orgID"
)

// New creates a new cluster Flux CD helm release resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster group Flux CD helm release resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmReleaseResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceCreate(request *helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest) (*helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceDelete(fn *helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName) error

	VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceGet(fn *helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName) (*helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseGetResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceUpdate(request *helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest) (*helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceCreate creates a Flux CD helm release scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceCreate(request *helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest) (*helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Release.FullName.ClusterGroupName, apiSubGroup, apiKind).String()
	releaseClusterGroupResponse := &helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse{}
	err := p.Create(requestURL, request, releaseClusterGroupResponse)

	return releaseClusterGroupResponse, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceDelete deletes a Flux CD helm release scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceDelete(fn *helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName) error {
	queryParams := url.Values{}

	if fn.NamespaceName != "" {
		queryParams.Add(queryParamKeyNamespaceName, fn.NamespaceName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiSubGroup, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceGet gets a Flux CD helm release scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceGet(fn *helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName) (*helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseGetResponse, error) {
	queryParams := url.Values{}

	if fn.NamespaceName != "" {
		queryParams.Add(queryParamKeyNamespaceName, fn.NamespaceName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiSubGroup, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	releaseClusterGroupResponse := &helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseGetResponse{}
	err := p.Get(requestURL, releaseClusterGroupResponse)

	return releaseClusterGroupResponse, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceUpdate updates overwrite a Flux CD helm release scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceUpdate(request *helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest) (*helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Release.FullName.ClusterGroupName, apiSubGroup, apiKind, request.Release.FullName.Name).String()
	releaseClusterGroupResponse := &helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse{}
	err := p.Update(requestURL, request, releaseClusterGroupResponse)

	return releaseClusterGroupResponse, err
}
