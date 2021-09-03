/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	clienterrors "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/errors"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/helper"
	clustergroupmodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/clustergroup"
)

// New creates a new cluster group resource service API client.
func New(transport *transport.Client, config *transport.Config) ClientService {
	return &Client{transport: transport, config: config}
}

/*
  Client for cluster group resource service API
*/
type Client struct {
	transport *transport.Client
	config    *transport.Config
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1ClusterGroupResourceServiceCreate(request *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupRequest) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupResponse, error)

	ManageV1alpha1ClusterGroupResourceServiceDelete(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) error

	ManageV1alpha1ClusterGroupResourceServiceGet(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse, error)
}

/*
  ManageV1alpha1ClusterGroupResourceServiceGet gets a cluster group
*/
func (a *Client) ManageV1alpha1ClusterGroupResourceServiceGet(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse, error) {
	requestURL := fmt.Sprintf("%s%s%s", a.config.Host, "/v1alpha1/clustergroups/", fn.Name)

	resp, err := a.transport.Get(requestURL, a.config.Headers)
	if err != nil {
		return nil, errors.Wrap(err, "read")
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("get tanzu TMC cluster group request failed with status : %v, response : %v", resp.Status, string(respBody))
	}

	clusterGroupResponse := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse{}

	err = clusterGroupResponse.UnmarshalBinary(respBody)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall")
	}

	return clusterGroupResponse, nil
}

/*
  ManageV1alpha1ClusterGroupResourceServiceDelete deletes a cluster group
*/
func (a *Client) ManageV1alpha1ClusterGroupResourceServiceDelete(fn *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) error {
	requestURL := fmt.Sprintf("%s%s%s", a.config.Host, "/v1alpha1/clustergroups/", fn.Name)

	resp, err := a.transport.Delete(requestURL, a.config.Headers)
	if err != nil {
		return errors.Wrap(err, "delete")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read delete response")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return clienterrors.ErrorWithHTTPCode(resp.StatusCode, errors.Errorf("delete tanzu TMC cluster group request failed with status : %v, response : %v", resp.Status, string(respBody)))
	}

	return nil
}

/*
  ManageV1alpha1ClusterGroupResourceServiceCreate creates a cluster group
*/
func (a *Client) ManageV1alpha1ClusterGroupResourceServiceCreate(request *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupRequest) (*clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupResponse, error) {
	requestURL := fmt.Sprintf("%s%s", a.config.Host, "/v1alpha1/clustergroups")

	body, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "marshall request body")
	}

	headers := a.config.Headers
	headers.Set(helper.ContentLengthKey, fmt.Sprintf("%d", len(body)))

	resp, err := a.transport.Post(requestURL, bytes.NewReader(body), headers)
	if err != nil {
		return nil, errors.Wrap(err, "create")
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read create response")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("create tanzu TMC cluster group request failed with status : %v, response : %v", resp.Status, string(respBody))
	}

	clusterGroupResponse := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupResponse{}

	err = clusterGroupResponse.UnmarshalBinary(respBody)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall")
	}

	return clusterGroupResponse, nil
}
