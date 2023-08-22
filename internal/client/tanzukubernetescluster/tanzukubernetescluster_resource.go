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
	queryParamKeyForce      = "force"
)

func getBaseReqURL(mgmtClsName, provisionerName string) helper.RequestURL {
	return helper.ConstructRequestURL(apiVersionAndGroup, mgmtClsName, provisioners, provisionerName, tanzukubernetesclusters)
}

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
	TanzuKubernetesClusterResourceServiceCreate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterResponse, error)

	TanzuKubernetesClusterResourceServiceDelete(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName, force string) error

	TanzuKubernetesClusterResourceServiceGet(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse, error)

	TanzuKubernetesClusterResourceServiceUpdate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterResponse, error)
}

/*
TanzuKubernetesClusterResourceServiceCreate creates a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceCreate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterResponse, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterResponse{}

	var reqURL helper.RequestURL

	if req.TanzuKubernetesCluster.FullName != nil {
		if req.TanzuKubernetesCluster.FullName.ManagementClusterName != "" && req.TanzuKubernetesCluster.FullName.ProvisionerName != "" {
			reqURL = getBaseReqURL(req.TanzuKubernetesCluster.FullName.ManagementClusterName, req.TanzuKubernetesCluster.FullName.ProvisionerName)
		}
	}

	err := c.Create(reqURL.String(), req, response)

	return response, err
}

/*
TanzuKubernetesClusterResourceServiceDelete deletes a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceDelete(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName, force string) error {
	var reqURL helper.RequestURL

	queryParams := url.Values{
		queryParamKeyForce: []string{force},
	}

	if fn.ManagementClusterName != "" && fn.ProvisionerName != "" {
		reqURL = getBaseReqURL(fn.ManagementClusterName, fn.ProvisionerName)
	}

	if fn.Name != "" {
		reqURL = helper.ConstructRequestURL(reqURL.String(), fn.Name).AppendQueryParams(queryParams)
	}

	err := c.Delete(reqURL.String())

	return err
}

/*
TanzuKubernetesClusterResourceServiceGet gets a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceGet(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse{}

	var reqURL helper.RequestURL

	if fn.ManagementClusterName != "" && fn.ProvisionerName != "" {
		reqURL = getBaseReqURL(fn.ManagementClusterName, fn.ProvisionerName)
	}

	if fn.Name != "" {
		reqURL = helper.ConstructRequestURL(reqURL.String(), fn.Name)
	}

	err := c.Get(reqURL.String(), response)

	return response, err
}

/*
TanzuKubernetesClusterResourceServiceUpdate updates overwrite a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceUpdate(
	req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterResponse, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateTanzuKubernetesClusterResponse{}

	var reqURL helper.RequestURL

	if req.TanzuKubernetesCluster.FullName != nil {
		if req.TanzuKubernetesCluster.FullName.ManagementClusterName != "" && req.TanzuKubernetesCluster.FullName.ProvisionerName != "" {
			getBaseReqURL(req.TanzuKubernetesCluster.FullName.ManagementClusterName, req.TanzuKubernetesCluster.FullName.ProvisionerName)
		}

		if req.TanzuKubernetesCluster.FullName.Name != "" {
			reqURL = helper.ConstructRequestURL(reqURL.String(), req.TanzuKubernetesCluster.FullName.Name)
		}
	}

	err := c.Update(reqURL.String(), req, response)

	return response, err
}
