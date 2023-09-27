/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetesclusterclient

import (
	// "fmt"
	// "net/http"
	"net/url"

	// "github.com/pkg/errors"

	// clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
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

func getBaseReqURL(mgmtClsName string, provisionerName string) helper.RequestURL {
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
	TanzuKubernetesClusterResourceServiceCreate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse, error)

	TanzuKubernetesClusterResourceServiceDelete(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName, force string) error

	TanzuKubernetesClusterResourceServiceGet(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse, error)

	// TanzuKubernetesClusterResourceServiceGetByID(id string) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse, error)

	TanzuKubernetesClusterResourceServiceUpdate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse, error)
}

/*
TanzuKubernetesClusterResourceServiceCreate creates a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceCreate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse{}

	requestURL := getBaseReqURL(req.TanzuKubernetesCluster.FullName.ManagementClusterName, req.TanzuKubernetesCluster.FullName.ProvisionerName).String()

	err := c.Create(requestURL, req, response)

	return response, err
}

/*
TanzuKubernetesClusterResourceServiceDelete deletes a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceDelete(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName, force string) error {
	queryParams := url.Values{
		queryParamKeyForce: []string{force},
	}

	requestURL := getBaseReqURL(fn.ManagementClusterName, fn.ProvisionerName).String()
	requestURL = helper.ConstructRequestURL(requestURL, fn.Name).AppendQueryParams(queryParams).String()

	err := c.Delete(requestURL)

	return err
}

/*
TanzuKubernetesClusterResourceServiceGet gets a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceGet(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse{}

	requestURL := getBaseReqURL(fn.ManagementClusterName, fn.ProvisionerName).String()
	requestURL = helper.ConstructRequestURL(requestURL, fn.Name).String()

	err := c.Get(requestURL, response)

	return response, err
}

/*
EksClusterResourceServiceGetByID gets an eks cluster by its ID.
*/
// func (c *Client) TanzuKubernetesClusterResourceServiceGetByID(id string) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse, error) {
// 	queryParams := url.Values{
// 		"query": []string{fmt.Sprintf("uid=\"%s\"", id)},
// 	}

// 	requestURL := getBaseReqURL(fn.ManagementClusterName, fn.ProvisionerName).String()
// 	requestURL = helper.ConstructRequestURL(requestURL, fn.Name).String()

// 	requestURL := helper.ConstructRequestURL(apiVersionAndGroup).AppendQueryParams(queryParams).String()
// 	clusterListResponse := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterListTanzuKubernetesClustersResponse{}

// 	err := c.Get(requestURL, clusterListResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(clusterListResponse.TanzuKubernetesClusters) == 0 {
// 		return nil, clienterrors.ErrorWithHTTPCode(http.StatusNotFound, errors.New("cluster list by ID was empty"))
// 	}

// 	if len(clusterListResponse.TanzuKubernetesClusters) > 1 {
// 		return nil, clienterrors.ErrorWithHTTPCode(http.StatusExpectationFailed, errors.New("cluster list by ID returned more than one cluster"))
// 	}

// 	clusterResponse := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterGetTanzuKubernetesClusterResponse{}
// 	clusterResponse.TanzuKubernetesCluster = clusterListResponse.TanzuKubernetesClusters[0]

// 	return clusterResponse, nil
// }

/*
TanzuKubernetesClusterResourceServiceUpdate updates overwrite a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceUpdate(
	req *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse{}

	requestURL := getBaseReqURL(req.TanzuKubernetesCluster.FullName.ManagementClusterName, req.TanzuKubernetesCluster.FullName.ProvisionerName).String()
	requestURL = helper.ConstructRequestURL(requestURL, req.TanzuKubernetesCluster.FullName.Name).String()

	err := c.Update(requestURL, req, response)

	return response, err
}
