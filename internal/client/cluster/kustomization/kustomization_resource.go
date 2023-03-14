/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kustomizationclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiSubGroup                        = "namespaces"
	apiKind                            = "fluxcd/kustomizations"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
	queryParamKeyOrgID                 = "fullName.orgID"
)

// New creates a new cluster Flux CD kustomization resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster Flux CD kustomization resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceCreate(request *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest) (*kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse, error)

	VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceDelete(fn *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName) error

	VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceGet(fn *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName) (*kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationGetKustomizationResponse, error)

	VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceUpdate(request *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest) (*kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceCreate creates a Flux CD kustomization scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceCreate(request *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest) (*kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Kustomization.FullName.ClusterName, apiSubGroup, request.Kustomization.FullName.NamespaceName, apiKind).String()
	fluxCDKustomizationClusterResponse := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse{}
	err := p.Create(requestURL, request, fluxCDKustomizationClusterResponse)

	return fluxCDKustomizationClusterResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceDelete deletes a Flux CD kustomization scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceDelete(fn *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName) error {
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

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiSubGroup, fn.NamespaceName, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceGet gets a Flux CD kustomization scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceGet(fn *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName) (*kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationGetKustomizationResponse, error) {
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

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiSubGroup, fn.NamespaceName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	fluxCDKustomizationClusterResponse := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationGetKustomizationResponse{}
	err := p.Get(requestURL, fluxCDKustomizationClusterResponse)

	return fluxCDKustomizationClusterResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceUpdate updates overwrite a Flux CD kustomization scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceUpdate(request *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest) (*kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Kustomization.FullName.ClusterName, apiSubGroup, request.Kustomization.FullName.NamespaceName, apiKind, request.Kustomization.FullName.Name).String()
	fluxCDKustomizationClusterResponse := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse{}
	err := p.Update(requestURL, request, fluxCDKustomizationClusterResponse)

	return fluxCDKustomizationClusterResponse, err
}
