/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iamclustergroupclient

import (
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/helper"
	clustergroupmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	iammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	clustergroupiammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy/clustergroup"
)

const apiVersionAndGroup = "v1alpha1/clustergroups:iam"

// New creates a new cluster group IAM policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
  Client for cluster group IAM policy resource service API.
*/

type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1ClusterGroupIAMPolicy Client methods.
type ClientService interface {
	ManageV1alpha1ClusterGroupIAMPolicyGet(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) (*clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupIAMPolicyResponse, error)

	ManageV1alpha1ClusterGroupIAMPolicyPatch(request *clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest) (*clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyResponse, error)

	ManageV1alpha1ClusterGroupIAMPolicyUpdate(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupUpdateClusterGroupIAMPolicyResponse, error)
}

/*
  ManageV1alpha1ClusterGroupIAMPolicyGet gets all iam policies scoped to a cluster group.
*/

func (c *Client) ManageV1alpha1ClusterGroupIAMPolicyGet(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) (*clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupIAMPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).String()
	response := &clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupIAMPolicyResponse{}
	err := c.Get(requestURL, response)

	return response, err
}

/*
  ManageV1alpha1ClusterGroupIAMPolicyPatch patches all iam policies scoped to a cluster group.
*/

func (c *Client) ManageV1alpha1ClusterGroupIAMPolicyPatch(request *clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest) (*clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyResponse, error) {
	response := &clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.FullName.Name).String()
	err := c.Patch(requestURL, request, response)

	return response, err
}

/*
  ManageV1alpha1ClusterGroupIAMPolicyUpdate updates overwrite all iam policies scoped to a cluster group, deletes if body is empty.
*/

func (c *Client) ManageV1alpha1ClusterGroupIAMPolicyUpdate(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupUpdateClusterGroupIAMPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).String()
	response := &clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupUpdateClusterGroupIAMPolicyResponse{}
	err := c.Update(requestURL, request, response)

	return response, err
}
