/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyclustergroupclient

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
)

const (
	apiVersionAndGroup = "v1alpha1/clustergroups"
	apiKind            = "policies"
)

// New creates a new cluster group policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster group policy resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1ClustergroupPolicyResourceService Client methods.
type ClientService interface {
	ManageV1alpha1ClustergroupPolicyResourceServiceCreate(request *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse, error)

	ManageV1alpha1ClustergroupPolicyResourceServiceDelete(fn *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) error

	ManageV1alpha1ClustergroupPolicyResourceServiceGet(fn *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse, error)

	ManageV1alpha1ClustergroupPolicyResourceServiceUpdate(request *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse, error)
}

/*
ManageV1alpha1ClustergroupPolicyResourceServiceCreate creates a policy scoped to a cluster group resource.
*/
func (p *Client) ManageV1alpha1ClustergroupPolicyResourceServiceCreate(request *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Policy.FullName.ClusterGroupName, apiKind).String()
	policyClusterGroupResponse := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse{}
	err := p.Create(requestURL, request, policyClusterGroupResponse)

	return policyClusterGroupResponse, err
}

/*
ManageV1alpha1ClustergroupPolicyResourceServiceDelete deletes a policy scoped to a cluster group resource.
*/
func (p *Client) ManageV1alpha1ClustergroupPolicyResourceServiceDelete(fn *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) error {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiKind, fn.Name).String()

	return p.Delete(requestURL)
}

/*
ManageV1alpha1ClustergroupPolicyResourceServiceGet gets a policy scoped to a cluster group resource.
*/
func (p *Client) ManageV1alpha1ClustergroupPolicyResourceServiceGet(fn *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiKind, fn.Name).String()
	policyClusterGroupResponse := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse{}
	err := p.Get(requestURL, policyClusterGroupResponse)

	return policyClusterGroupResponse, err
}

/*
ManageV1alpha1ClustergroupPolicyResourceServiceUpdate updates overwrite a policy scoped to a cluster group resource.
*/
func (p *Client) ManageV1alpha1ClustergroupPolicyResourceServiceUpdate(request *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) (*policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Policy.FullName.ClusterGroupName, apiKind, request.Policy.FullName.Name).String()
	policyClusterGroupResponse := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse{}
	err := p.Update(requestURL, request, policyClusterGroupResponse)

	return policyClusterGroupResponse, err
}
