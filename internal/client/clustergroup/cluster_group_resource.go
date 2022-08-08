/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupclient

import (
	"fmt"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
)

// New creates a new cluster group resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
  Client for cluster group resource service API
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1ClusterGroupResourceServiceCreate(request *clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupRequest) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupResponse, error)

	ManageV1alpha1ClusterGroupResourceServiceDelete(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) error

	ManageV1alpha1ClusterGroupResourceServiceGet(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse, error)

	ManageV1alpha1ClusterGroupResourceServiceUpdate(request *clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupRequest) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupResponse, error)
}

/*
  ManageV1alpha1ClusterGroupResourceServiceGet gets a cluster group
*/
func (c *Client) ManageV1alpha1ClusterGroupResourceServiceGet(
	fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName,
) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse, error) {
	requestURL := fmt.Sprintf("%s/%s", "v1alpha1/clustergroups", fn.Name)
	clusterGroupResponse := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse{}

	err := c.Get(requestURL, clusterGroupResponse)

	return clusterGroupResponse, err
}

/*
  ManageV1alpha1ClusterGroupResourceServiceDelete deletes a cluster group
*/
func (c *Client) ManageV1alpha1ClusterGroupResourceServiceDelete(
	fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName,
) error {
	requestURL := fmt.Sprintf("%s/%s", "v1alpha1/clustergroups", fn.Name)

	return c.Delete(requestURL)
}

/*
  ManageV1alpha1ClusterGroupResourceServiceCreate creates a cluster group
*/
func (c *Client) ManageV1alpha1ClusterGroupResourceServiceCreate(
	request *clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupRequest,
) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupResponse, error) {
	clusterGroupResponse := &clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupResponse{}
	err := c.Create("v1alpha1/clustergroups", request, clusterGroupResponse)

	return clusterGroupResponse, err
}

/*
  ManageV1alpha1ClusterGroupResourceServiceUpdate updates a cluster group
*/
func (c *Client) ManageV1alpha1ClusterGroupResourceServiceUpdate(
	request *clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupRequest,
) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupResponse, error) {
	requestURL := fmt.Sprintf("%s/%s", "v1alpha1/clustergroups", request.ClusterGroup.FullName.Name)
	clusterGroupResponse := &clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupResponse{}
	err := c.Update(requestURL, request, clusterGroupResponse)

	return clusterGroupResponse, err
}
