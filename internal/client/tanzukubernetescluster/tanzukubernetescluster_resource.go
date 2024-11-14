// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzukubernetesclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	kubeconfigmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/kubeconfig"

	tkcmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	tkcnodepoolmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/nodepool"
)

const (
	// Cluster API.
	tanzuKubernetesClusterAPIVersionAndGroupPath = "v1alpha1/managementclusters"
	provisionersPath                             = "provisioners"
	tanzuKubernetesClustersPath                  = "tanzukubernetesclusters"
	forceQueryParamKey                           = "force"

	// Node Pools API.
	nodePoolsPath = "nodepools"

	// KubeConfig API.
	clusterAPIVersionAndGroupPath = "v1alpha1/clusters"
	kubeConfigPath                = "kubeconfig"
	managementClusterNameParamKey = "fullName.managementClusterName"
	provisionerNameParamKey       = "fullName.provisionerName"
	cliParamKey                   = "cli"
	cliParamValue                 = "TANZU_CLI"
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

	// Cluster API.
	TanzuKubernetesClusterResourceServiceCreate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error)

	TanzuKubernetesClusterResourceServiceDelete(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName, force bool) error

	TanzuKubernetesClusterResourceServiceGet(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error)

	TanzuKubernetesClusterResourceServiceUpdate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error)

	// Node Pools API.
	TanzuKubernetesClusterNodePoolResourceServiceCreate(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error)

	TanzuKubernetesClusterNodePoolResourceServiceDelete(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) error

	TanzuKubernetesClusterNodePoolResourceServiceGet(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error)

	TanzuKubernetesClusterNodePoolResourceServiceList(clusterFn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error)

	TanzuKubernetesClusterNodePoolResourceServiceUpdate(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error)

	// KubeConfig API.
	KubeConfigResourceServiceGet(clusterFn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*kubeconfigmodels.VmwareTanzuManageV1alpha1ClusterKubeconfigResponse, error)
}

// Cluster APIs.

/*
TanzuKubernetesClusterResourceServiceCreate creates a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceCreate(request *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{}
	requestURL := helper.ConstructRequestURL(tanzuKubernetesClusterAPIVersionAndGroupPath,
		request.TanzuKubernetesCluster.FullName.ManagementClusterName, provisionersPath,
		request.TanzuKubernetesCluster.FullName.ProvisionerName, tanzuKubernetesClustersPath)
	err := c.Create(requestURL.String(), request, response)

	return response, err
}

/*
TanzuKubernetesClusterResourceServiceDelete deletes a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceDelete(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName, force bool) error {
	queryParams := url.Values{
		forceQueryParamKey: {helper.ConvertToString(force, "")},
	}

	requestURL := helper.ConstructRequestURL(tanzuKubernetesClusterAPIVersionAndGroupPath,
		fn.ManagementClusterName, provisionersPath, fn.ProvisionerName, tanzuKubernetesClustersPath, fn.Name).AppendQueryParams(queryParams)
	err := c.Delete(requestURL.String())

	return err
}

/*
TanzuKubernetesClusterResourceServiceGet gets a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceGet(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{}
	requestURL := helper.ConstructRequestURL(tanzuKubernetesClusterAPIVersionAndGroupPath,
		fn.ManagementClusterName, provisionersPath, fn.ProvisionerName, tanzuKubernetesClustersPath, fn.Name)
	err := c.Get(requestURL.String(), response)

	return response, err
}

/*
TanzuKubernetesClusterResourceServiceUpdate updates overwrite a tanzu kubernetes cluster.
*/
func (c *Client) TanzuKubernetesClusterResourceServiceUpdate(request *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
	response := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{}
	requestURL := helper.ConstructRequestURL(tanzuKubernetesClusterAPIVersionAndGroupPath,
		request.TanzuKubernetesCluster.FullName.ManagementClusterName, provisionersPath,
		request.TanzuKubernetesCluster.FullName.ProvisionerName, tanzuKubernetesClustersPath,
		request.TanzuKubernetesCluster.FullName.Name)
	err := c.Update(requestURL.String(), request, response)

	return response, err
}

// Node Pool APIs.

/*
TanzuKubernetesClusterNodePoolResourceServiceCreate creates a tanzu kubernetes cluster node pool.
*/
func (c *Client) TanzuKubernetesClusterNodePoolResourceServiceCreate(request *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
	response := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{}
	requestURL := helper.ConstructRequestURL(tanzuKubernetesClusterAPIVersionAndGroupPath,
		request.Nodepool.FullName.ManagementClusterName, provisionersPath, request.Nodepool.FullName.ProvisionerName,
		tanzuKubernetesClustersPath, request.Nodepool.FullName.TanzuKubernetesClusterName, nodePoolsPath)
	err := c.Create(requestURL.String(), request, response)

	return response, err
}

/*
TanzuKubernetesClusterNodePoolResourceServiceDelete deletes a tanzu kubernetes cluster node pool.
*/
func (c *Client) TanzuKubernetesClusterNodePoolResourceServiceDelete(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) error {
	requestURL := helper.ConstructRequestURL(tanzuKubernetesClusterAPIVersionAndGroupPath, fn.ManagementClusterName,
		provisionersPath, fn.ProvisionerName, tanzuKubernetesClustersPath, fn.TanzuKubernetesClusterName,
		nodePoolsPath, fn.Name)
	err := c.Delete(requestURL.String())

	return err
}

/*
TanzuKubernetesClusterNodePoolResourceServiceGet gets a tanzu kubernetes cluster node pool.
*/
func (c *Client) TanzuKubernetesClusterNodePoolResourceServiceGet(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
	response := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{}
	requestURL := helper.ConstructRequestURL(tanzuKubernetesClusterAPIVersionAndGroupPath, fn.ManagementClusterName,
		provisionersPath, fn.ProvisionerName, tanzuKubernetesClustersPath, fn.TanzuKubernetesClusterName,
		nodePoolsPath, fn.Name)
	err := c.Get(requestURL.String(), response)

	return response, err
}

/*
TanzuKubernetesClusterNodePoolResourceServiceList lists a tanzu kubernetes cluster node pools.
*/
func (c *Client) TanzuKubernetesClusterNodePoolResourceServiceList(clusterFn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
	response := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{}
	requestURL := helper.ConstructRequestURL(tanzuKubernetesClusterAPIVersionAndGroupPath,
		clusterFn.ManagementClusterName, provisionersPath, clusterFn.ProvisionerName, tanzuKubernetesClustersPath,
		clusterFn.Name, nodePoolsPath)

	err := c.Get(requestURL.String(), response)

	return response, err
}

/*
TanzuKubernetesClusterNodePoolResourceServiceUpdate updates overwrite a tanzu kubernetes cluster node pool.
*/
func (c *Client) TanzuKubernetesClusterNodePoolResourceServiceUpdate(request *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
	response := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{}
	requestURL := helper.ConstructRequestURL(tanzuKubernetesClusterAPIVersionAndGroupPath,
		request.Nodepool.FullName.ManagementClusterName, provisionersPath, request.Nodepool.FullName.ProvisionerName,
		tanzuKubernetesClustersPath, request.Nodepool.FullName.TanzuKubernetesClusterName, nodePoolsPath,
		request.Nodepool.FullName.Name)
	err := c.Update(requestURL.String(), request, response)

	return response, err
}

// KubeConfig

/*
KubeConfigResourceServiceGet gets a cluster's kubeconfig.
*/
func (c *Client) KubeConfigResourceServiceGet(clusterFn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*kubeconfigmodels.VmwareTanzuManageV1alpha1ClusterKubeconfigResponse, error) {
	response := &kubeconfigmodels.VmwareTanzuManageV1alpha1ClusterKubeconfigResponse{}
	requestURL := helper.ConstructRequestURL(clusterAPIVersionAndGroupPath, clusterFn.Name, kubeConfigPath)

	queryParams := url.Values{}

	queryParams.Add(managementClusterNameParamKey, clusterFn.ManagementClusterName)
	queryParams.Add(provisionerNameParamKey, clusterFn.ProvisionerName)
	queryParams.Add(cliParamKey, cliParamValue)

	requestURL = requestURL.AppendQueryParams(queryParams)

	err := c.Get(requestURL.String(), response)

	return response, err
}
