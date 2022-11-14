/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
)

const (
	apiVersionAndGroup          = "/v1alpha1/eksclusters"
	queryParamKeyForce          = "force"
	queryParamKeyCredentialName = "fullName.credentialName" //nolint:gosec
	queryParamKeyRegion         = "fullName.region"
)

// New creates a new eks cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for eks cluster resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	EksClusterResourceServiceCreate(request *eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse, error)

	EksClusterResourceServiceDelete(fn *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName, force string) error

	EksClusterResourceServiceGet(fn *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse, error)

	EksClusterResourceServiceUpdate(request *eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse, error)
}

/*
EksClusterResourceServiceCreate creates an eks cluster.
*/
func (c *Client) EksClusterResourceServiceCreate(request *eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse, error) {
	response := &eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse{}
	err := c.Create(apiVersionAndGroup, request, response)

	return response, err
}

/*
EksClusterResourceServiceDelete deletes an eks cluster.
*/
func (c *Client) EksClusterResourceServiceDelete(fn *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName, force string) error {
	queryParams := url.Values{
		queryParamKeyForce: []string{force},
	}
	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeyCredentialName, fn.CredentialName)
	}

	if fn.Region != "" {
		queryParams.Add(queryParamKeyRegion, fn.Region)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}

/*
EksClusterResourceServiceGet gets an eks cluster.
*/
func (c *Client) EksClusterResourceServiceGet(fn *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse, error) {
	queryParams := url.Values{}
	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeyCredentialName, fn.CredentialName)
	}

	if fn.Region != "" {
		queryParams.Add(queryParamKeyRegion, fn.Region)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).AppendQueryParams(queryParams).String()
	clusterResponse := &eksmodel.VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse{}

	err := c.Get(requestURL, clusterResponse)

	return clusterResponse, err
}

/*
EksClusterResourceServiceUpdate updates overwrite an eks cluster.
*/
func (c *Client) EksClusterResourceServiceUpdate(request *eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse, error) {
	response := &eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.EksCluster.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}
