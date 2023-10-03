/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclient

import (
	"net/url"
	"strconv"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	dataprotectionmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/dataprotection"
)

const (
	apiVersionAndGroup = "v1alpha1/clusters"
	dataProtectionPath = "dataprotection"
)

// New creates a new cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for credentials resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	DataProtectionResourceServiceCreate(request *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest) (*dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse, error)

	DataProtectionResourceServiceDelete(fn *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName, destroyBackups bool) error

	DataProtectionResourceServiceList(clusterName string) (*dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse, error)

	DataProtectionResourceServiceUpdate(request *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest) (*dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse, error)
}

/*
DataProtectionResourceServiceCreate enables data protection on a cluster.
*/
func (c *Client) DataProtectionResourceServiceCreate(request *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest,
) (*dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse, error) {
	response := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.DataProtection.FullName.ClusterName, dataProtectionPath).String()
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
DataProtectionResourceServiceDelete disables data protection on a cluster.
*/
func (c *Client) DataProtectionResourceServiceDelete(fn *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName, deleteBackups bool) error {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, dataProtectionPath)
	queryParams := url.Values{}

	queryParams.Add("fullName.provisionerName", fn.ProvisionerName)
	queryParams.Add("fullName.managementClusterName", fn.ManagementClusterName)
	queryParams.Add("delete_backups", strconv.FormatBool(deleteBackups))

	requestURL = requestURL.AppendQueryParams(queryParams)

	return c.Delete(requestURL.String())
}

/*
DataProtectionResourceServiceList gets data protection details.
*/
func (c *Client) DataProtectionResourceServiceList(clusterName string) (*dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, clusterName, dataProtectionPath).String()
	resp := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse{}
	err := c.Get(requestURL, resp)

	return resp, err
}

/*
DataProtectionResourceServiceUpdate updates a data protection configuration on a cluster.
*/
func (c *Client) DataProtectionResourceServiceUpdate(request *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest,
) (*dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse, error) {
	response := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.DataProtection.FullName.ClusterName, dataProtectionPath).String()
	err := c.Update(requestURL, request, response)

	return response, err
}
