// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kustomizationclustergroupclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
)

const (
	apiVersionAndGroup         = "v1alpha1/clustergroups"
	apiSubGroup                = "namespace"
	apiKind                    = "fluxcd/kustomizations"
	queryParamKeyNamespaceName = "fullName.namespaceName"
	queryParamKeyOrgID         = "fullName.orgID"
)

// New creates a new cluster Flux CD kustomization resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster group Flux CD kustomization resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceCreate(request *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest) (*kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceDelete(fn *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName) error

	VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceGet(fn *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName) (*kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceUpdate(request *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest) (*kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceCreate creates a Flux CD kustomization scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceCreate(request *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest) (*kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Kustomization.FullName.ClusterGroupName, apiSubGroup, apiKind).String()
	fluxCDKustomizationClusterGroupResponse := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse{}
	err := p.Create(requestURL, request, fluxCDKustomizationClusterGroupResponse)

	return fluxCDKustomizationClusterGroupResponse, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceDelete deletes a Flux CD kustomization scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceDelete(fn *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName) error {
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
VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceGet gets a Flux CD kustomization scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceGet(fn *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName) (*kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse, error) {
	queryParams := url.Values{}

	if fn.NamespaceName != "" {
		queryParams.Add(queryParamKeyNamespaceName, fn.NamespaceName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiSubGroup, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	fluxCDKustomizationClusterGroupResponse := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse{}
	err := p.Get(requestURL, fluxCDKustomizationClusterGroupResponse)

	return fluxCDKustomizationClusterGroupResponse, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceUpdate updates overwrite a Flux CD kustomization scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceUpdate(request *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest) (*kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Kustomization.FullName.ClusterGroupName, apiSubGroup, apiKind, request.Kustomization.FullName.Name).String()
	fluxCDKustomizationClusterGroupResponse := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse{}
	err := p.Update(requestURL, request, fluxCDKustomizationClusterGroupResponse)

	return fluxCDKustomizationClusterGroupResponse, err
}
