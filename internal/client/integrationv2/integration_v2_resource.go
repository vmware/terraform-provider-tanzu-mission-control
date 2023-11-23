/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integrationv2client

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"

	clusterintegrationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/integration/cluster"
	clustergroupintegrationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/integration/clustergroup"
)

const (
	clusterIntegrationAPIRootPath      = "v1alpha1/clusters"
	clusterGroupIntegrationAPIRootPath = "v1alpha1/clustergroups"
	integrationsPath                   = "integrations"

	managementClusterNameQueryParam = "fullName.managementClusterName"
	provisionerNameQueryParam       = "fullName.provisionerName"
)

func New(transport *transport.Client) ClientService {
	return &client{Client: transport}
}

type client struct {
	*transport.Client
}

type ClientService interface {
	// Cluster API.

	ClusterIntegrationResourceServiceCreate(
		data *clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData,
	) (*clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData, error)

	ClusterIntegrationResourceServiceUpdate(
		request *clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData,
	) (*clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData, error)

	ClusterIntegrationResourceServiceRead(
		*clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationFullName,
	) (*clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData, error)

	ClusterIntegrationResourceServiceDelete(
		*clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationFullName,
	) error

	// Cluster Group API.

	ClusterGroupIntegrationResourceServiceCreate(
		request *clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData,
	) (*clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData, error)

	ClusterGroupIntegrationResourceServiceRead(
		fn *clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationFullName,
	) (*clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData, error)

	ClusterGroupIntegrationResourceServiceDelete(
		fn *clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationFullName,
	) error
}

// Cluster API.

func (c *client) ClusterIntegrationResourceServiceCreate(
	request *clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData,
) (*clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData, error) {
	response := &clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData{}
	requestURL := helper.ConstructRequestURL(clusterIntegrationAPIRootPath, request.Integration.FullName.ClusterName, integrationsPath)
	err := c.Create(requestURL.String(), request, response)

	return response, err
}

func (c *client) ClusterIntegrationResourceServiceUpdate(
	request *clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData,
) (*clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData, error) {
	response := &clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData{}
	requestURL := helper.ConstructRequestURL(clusterIntegrationAPIRootPath, request.Integration.FullName.ClusterName, integrationsPath, request.Integration.FullName.Name)
	err := c.Update(requestURL.String(), request, response)

	return response, err
}

func (c *client) ClusterIntegrationResourceServiceRead(
	fn *clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationFullName,
) (*clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData, error) {
	response := &clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData{}
	requestURL := helper.ConstructRequestURL(clusterIntegrationAPIRootPath, fn.ClusterName, integrationsPath, fn.Name)
	queryParams := url.Values{}

	queryParams.Add(managementClusterNameQueryParam, fn.ManagementClusterName)
	queryParams.Add(provisionerNameQueryParam, fn.ProvisionerName)

	requestURL = requestURL.AppendQueryParams(queryParams)
	err := c.Get(requestURL.String(), response)

	return response, err
}

func (c *client) ClusterIntegrationResourceServiceDelete(
	fn *clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationFullName,
) error {
	requestURL := helper.ConstructRequestURL(clusterIntegrationAPIRootPath, fn.ClusterName, integrationsPath, fn.Name)
	queryParams := url.Values{}

	queryParams.Add(managementClusterNameQueryParam, fn.ManagementClusterName)
	queryParams.Add(provisionerNameQueryParam, fn.ProvisionerName)

	requestURL = requestURL.AppendQueryParams(queryParams)

	return c.Delete(requestURL.String())
}

// Cluster Group API.

func (c *client) ClusterGroupIntegrationResourceServiceCreate(
	request *clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData,
) (*clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData, error) {
	response := &clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData{}
	requestURL := helper.ConstructRequestURL(clusterGroupIntegrationAPIRootPath, request.Integration.FullName.ClusterGroupName, integrationsPath)
	err := c.Create(requestURL.String(), request, response)

	return response, err
}

func (c *client) ClusterGroupIntegrationResourceServiceRead(
	fn *clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationFullName,
) (*clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData, error) {
	response := &clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData{}
	requestURL := helper.ConstructRequestURL(clusterGroupIntegrationAPIRootPath, fn.ClusterGroupName, integrationsPath, fn.Name)
	err := c.Get(requestURL.String(), response)

	return response, err
}

func (c *client) ClusterGroupIntegrationResourceServiceDelete(
	fn *clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationFullName,
) error {
	requestURL := helper.ConstructRequestURL(clusterGroupIntegrationAPIRootPath, fn.ClusterGroupName, integrationsPath, fn.Name)

	return c.Delete(requestURL.String())
}
