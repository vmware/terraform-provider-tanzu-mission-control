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

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/helper"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"
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
