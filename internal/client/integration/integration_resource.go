/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integrationclient

import (
	"fmt"
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/integration"
)

func New(transport *transport.Client) ClientService {
	return &client{Client: transport}
}

type client struct {
	*transport.Client
}

type ClientService interface {
	ManageV1alpha1ClusterIntegrationResourceServiceCreate(
		*integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest,
	) (*integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse, error)

	ManageV1alpha1ClusterIntegrationResourceServiceRead(
		*integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName,
	) (*integration.VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse, error)

	ManageV1alpha1ClusterIntegrationResourceServiceDelete(
		*integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName,
	) error
}

func (c *client) ManageV1alpha1ClusterIntegrationResourceServiceCreate(
	request *integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest,
) (*integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse, error) {
	response := &integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse{}

	fn := request.Integration.FullName
	if errors := validateFullName(fn); len(errors) > 0 {
		return nil, fmt.Errorf("incomplete full name: %v (%v)", errors, fn)
	}

	err := c.Create(collectionEndpoint(fn), request, response)

	return response, err
}

func (c *client) ManageV1alpha1ClusterIntegrationResourceServiceRead(
	fn *integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName,
) (*integration.VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse, error) {
	if errors := validateFullName(fn); len(errors) > 0 {
		return nil, fmt.Errorf("incomplete full name: %v (%v)", errors, fn)
	}

	response := &integration.VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse{}
	err := c.Get(resourceEndpoint(fn), response)

	return response, err
}

func (c *client) ManageV1alpha1ClusterIntegrationResourceServiceDelete(
	fn *integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName,
) error {
	if errors := validateFullName(fn); len(errors) > 0 {
		return fmt.Errorf("incomplete full name: %v (%v)", errors, fn)
	}

	return c.Delete(resourceEndpoint(fn))
}

const (
	clusterIntegrationPrefix = "v1alpha1/clusters"
	integrationsResourceName = "integrations"
)

func collectionEndpoint(fn *integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) string {
	return helper.ConstructRequestURL(clusterIntegrationPrefix, fn.ClusterName, integrationsResourceName).String()
}

func resourceEndpoint(fn *integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) string {
	p := url.Values{}
	p.Set("fullName.managementClusterName", fn.ManagementClusterName)
	p.Set("fullName.provisionerName", fn.ProvisionerName)
	p.Set("fullName.name", fn.Name)

	return helper.
		ConstructRequestURL(clusterIntegrationPrefix, fn.ClusterName, integrationsResourceName, fn.Name).
		AppendQueryParams(p).
		String()
}

func validateFullName(fn *integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) []string {
	if fn == nil {
		return []string{"FullName is <nil>"}
	}

	var errors []string

	if fn.Name == "" {
		errors = append(errors, "missing Name")
	}

	if fn.ClusterName == "" {
		errors = append(errors, "missing ClusterName")
	}

	if fn.ManagementClusterName == "" {
		errors = append(errors, "missing ManagementClusterName")
	}

	if fn.ProvisionerName == "" {
		errors = append(errors, "missing ProvisionerName")
	}

	return errors
}
