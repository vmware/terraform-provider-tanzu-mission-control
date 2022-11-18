package eksnodepool

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
)

const (
	apiVersionAndGroup                     = "/v1alpha1/eksclusters"
	apiNodepoolsPath                       = "nodepools"
	queryParamKeySearchScopeCredentialName = "searchScope.credentialName"
	queryParamKeySearchScopeRegion         = "searchScope.region"
	queryParamKeySFullNameCredentialName   = "fullName.credentialName" //nolint:gosec
	queryParamKeyFullNameRegion            = "fullName.region"
	queryParamKeySortBy                    = "sortBy"
)

// New creates a new eks cluster node pool resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster node pool resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	EksNodePoolResourceServiceGet(fn *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse, error)

	EksNodePoolResourceServiceCreate(request *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse, error)

	EksNodePoolResourceServiceList(cluster *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse, error)

	EksNodePoolResourceServiceDelete(fn *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName) error

	EksNodePoolResourceServiceUpdate(request *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse, error)
}

// EksNodePoolResourceServiceGet implements ClientService.
func (c *Client) EksNodePoolResourceServiceGet(fn *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse, error) {
	queryParams := url.Values{}

	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeySFullNameCredentialName, fn.CredentialName)
	}

	if fn.Region != "" {
		queryParams.Add(queryParamKeyFullNameRegion, fn.Region)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.EksClusterName, apiNodepoolsPath, fn.Name).AppendQueryParams(queryParams).String()
	response := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse{}

	err := c.Get(requestURL, response)

	return response, err
}

// EksNodePoolResourceServiceCreate implements ClientService.
func (c *Client) EksNodePoolResourceServiceCreate(request *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse, error) {
	response := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Nodepool.FullName.EksClusterName, apiNodepoolsPath).String()

	err := c.Create(requestURL, request, response)

	return response, err
}

// EksNodePoolResourceServiceDelete implements ClientService.
func (c *Client) EksNodePoolResourceServiceDelete(fn *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName) error {
	queryParams := url.Values{}

	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeySFullNameCredentialName, fn.CredentialName)
	}

	if fn.Region != "" {
		queryParams.Add(queryParamKeyFullNameRegion, fn.Region)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.EksClusterName, apiNodepoolsPath, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}

// EksNodePoolResourceServiceList implements ClientService.
func (c *Client) EksNodePoolResourceServiceList(cluster *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse, error) {
	queryParams := url.Values{}

	// for stability of the results
	queryParams.Add(queryParamKeySortBy, "createTime")

	if cluster.CredentialName != "" {
		queryParams.Add(queryParamKeySearchScopeCredentialName, cluster.CredentialName)
	}

	if cluster.Region != "" {
		queryParams.Add(queryParamKeySearchScopeRegion, cluster.Region)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, cluster.Name, apiNodepoolsPath).AppendQueryParams(queryParams).String()
	clusterNodePoolsResponse := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse{}
	err := c.Get(requestURL, clusterNodePoolsResponse)

	return clusterNodePoolsResponse, err
}

// EksNodePoolResourceServiceUpdate implements ClientService.
func (c *Client) EksNodePoolResourceServiceUpdate(request *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest) (*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse, error) {
	response := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Nodepool.FullName.EksClusterName, apiNodepoolsPath, request.Nodepool.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}
