/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackageinstall

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	packageinstall "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiKind                            = "tanzupackage/installs"
	namespaces                         = "namespaces"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
)

// New creates a new repository resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for repository resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	InstallResourceServiceCreate(request *packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest) (*packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse, error)

	InstallResourceServiceDelete(fn *packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName) error

	InstallResourceServiceGet(fn *packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName) (*packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse, error)

	InstallResourceServiceUpdate(request *packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest) (*packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse, error)
}

/*
InstallResourceServiceCreate creates a repository.
*/
func (c *Client) InstallResourceServiceCreate(request *packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest) (*packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Install.FullName.ClusterName, namespaces, request.Install.FullName.NamespaceName, apiKind).String()
	pkgInstallResponse := &packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse{}
	err := c.Create(requestURL, request, pkgInstallResponse)

	return pkgInstallResponse, err
}

/*
InstallResourceServiceDelete deletes a repository.
*/
func (c *Client) InstallResourceServiceDelete(fn *packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, namespaces, fn.NamespaceName, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}

/*
InstallResourceServiceGet gets a repository.
*/
func (c *Client) InstallResourceServiceGet(fn *packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName) (*packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, namespaces, fn.NamespaceName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	pkgInstallResponse := &packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse{}
	err := c.Get(requestURL, pkgInstallResponse)

	return pkgInstallResponse, err
}

/*
InstallResourceServiceUpdate updates overwrite a repository.
*/
func (c *Client) InstallResourceServiceUpdate(request *packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest) (*packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Install.FullName.ClusterName, namespaces, request.Install.FullName.NamespaceName, apiKind, request.Install.FullName.Name).String()
	pkgInstallResponse := &packageinstall.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse{}
	err := c.Update(requestURL, request, pkgInstallResponse)

	return pkgInstallResponse, err
}
