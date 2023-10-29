/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetesclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"

	tkcmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
)

const (
	apiVersionAndGroup      = "/v1alpha1/managementclusters"
	provisioners            = "provisioners"
	tanzukubernetesclusters = "tanzukubernetesclusters"
	forceQueryParamKey      = "force"
)

// New creates a new tanzu kubernetes cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for tanzu kubernetes cluster resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	TanzuKubernetesClusterResourceServiceCreate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesClusterResponse, error)

	TanzuKubernetesClusterResourceServiceDelete(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName, force bool) error

	TanzuKubernetesClusterResourceServiceGet(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesClusterResponse, error)

	TanzuKubernetesClusterResourceServiceUpdate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesClusterResponse, error)
}

/*
TanzuKubernetesClusterResourceServiceCreate creates a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceCreate(request *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesClusterResponse, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesClusterResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.TanzuKubernetesCluster.FullName.ManagementClusterName, provisioners,
		request.TanzuKubernetesCluster.FullName.ProvisionerName, tanzukubernetesclusters)
	err := c.Create(requestURL.String(), request, response)

	return response, err
}

/*
TanzuKubernetesClusterResourceServiceDelete deletes a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceDelete(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName, force bool) error {
	queryParams := url.Values{
		forceQueryParamKey: {helper.ConvertToString(force, "")},
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ManagementClusterName, provisioners, fn.ProvisionerName, tanzukubernetesclusters, fn.Name).AppendQueryParams(queryParams)
	err := c.Delete(requestURL.String())

	return err
}

/*
TanzuKubernetesClusterResourceServiceGet gets a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceGet(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesClusterResponse, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesClusterResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ManagementClusterName, provisioners, fn.ProvisionerName, tanzukubernetesclusters, fn.Name)
	err := c.Get(requestURL.String(), response)

	return response, err
}

/*
TanzuKubernetesClusterResourceServiceUpdate updates overwrite a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceUpdate(request *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesClusterResponse, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesClusterResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.TanzuKubernetesCluster.FullName.ManagementClusterName, provisioners, request.TanzuKubernetesCluster.FullName.ProvisionerName, tanzukubernetesclusters, request.TanzuKubernetesCluster.FullName.Name)
	err := c.Update(requestURL.String(), request, response)

	return response, err
}
