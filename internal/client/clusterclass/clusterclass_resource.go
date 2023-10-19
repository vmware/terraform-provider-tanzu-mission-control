/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclassclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clusterclassmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clusterclass"
)

const (
	apiVersionAndGroup = "/v1alpha1/managementclusters"
	provisioners       = "provisioners"
	clusterClasses     = "clusterclasses"
	nameQueryParamKey  = "name"
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
	ClusterClassResourceServiceGet(fn *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassFullName) (*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassListClusterClassesResponse, error)
}

/*
ClusterClassResourceServiceGet gets or lists cluster classes.
*/
func (c *Client) ClusterClassResourceServiceGet(fn *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassFullName) (*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassListClusterClassesResponse, error) {
	response := &clusterclassmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassListClusterClassesResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ManagementClusterName, provisioners, fn.ProvisionerName, clusterClasses)

	if fn.Name != "" {
		queryParams := url.Values{
			nameQueryParamKey: {fn.Name},
		}

		requestURL = requestURL.AppendQueryParams(queryParams)
	}

	err := c.Get(requestURL.String(), response)

	return response, err
}