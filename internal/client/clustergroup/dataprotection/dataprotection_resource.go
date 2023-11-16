/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupdataprotectionclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	dataprotectionclustergroupmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/dataprotection/clustergroup/dataprotection"
)

const (
	apiVersionAndGroup = "v1alpha1/clustergroups"
	dataProtectionPath = "dataprotection"
)

// New creates a new data protection resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for data protection resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ClusterGroupDataProtectionResourceServiceCreate(request *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionRequest) (*dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse, error)

	ClusterGroupDataProtectionResourceServiceDelete(fn *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName, destroyBackups bool) error

	ClusterGroupDataProtectionResourceServiceList(fn *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName) (*dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionListDataProtectionsResponse, error)

	ClusterGroupDataProtectionResourceServiceUpdate(request *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionRequest) (*dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse, error)
}

/*
ClusterGroupDataProtectionResourceServiceCreate enables data protection on a cluster.
*/
func (c *Client) ClusterGroupDataProtectionResourceServiceCreate(request *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionRequest,
) (*dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse, error) {
	response := &dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.DataProtection.FullName.ClusterGroupName, dataProtectionPath).String()
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
ClusterGroupDataProtectionResourceServiceDelete disables data protection on a cluster group.
*/
func (c *Client) ClusterGroupDataProtectionResourceServiceDelete(fn *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName, deleteBackups bool) error {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, dataProtectionPath)
	queryParams := url.Values{}

	//queryParams.Add("fullName.clusterGroupName", fn.ClusterGroupName)
	//queryParams.Add("delete_backups", strconv.FormatBool(deleteBackups))

	requestURL = requestURL.AppendQueryParams(queryParams)

	return c.Delete(requestURL.String())
}

/*
ClusterGroupDataProtectionResourceServiceList gets data protection details.
*/
func (c *Client) ClusterGroupDataProtectionResourceServiceList(fn *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName) (*dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionListDataProtectionsResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, dataProtectionPath)
	queryParams := url.Values{}

	/*if fn.ManagementClusterName != "" {
		queryParams.Add("searchScope.managementClusterName", fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add("searchScope.provisionerName", fn.ProvisionerName)
	}*/

	if len(queryParams) > 0 {
		requestURL = requestURL.AppendQueryParams(queryParams)
	}

	resp := &dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionListDataProtectionsResponse{}
	err := c.Get(requestURL.String(), resp)

	return resp, err
}

/*
ClusterGroupDataProtectionResourceServiceUpdate updates a data protection configuration on a cluster.
*/
func (c *Client) ClusterGroupDataProtectionResourceServiceUpdate(request *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionRequest,
) (*dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse, error) {
	response := &dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.DataProtection.FullName.ClusterGroupName, dataProtectionPath).String()
	err := c.Update(requestURL, request, response)

	return response, err
}
