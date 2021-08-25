/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"
	clustermodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/cluster"
)

// New creates a new cluster resource service API client.
func New(transport *transport.Client, config *transport.Config) ClientService {
	return &Client{transport: transport, config: config}
}

/*
Client for cluster resource service API.
*/
type Client struct {
	transport *transport.Client
	config    *transport.Config
}

// ClientOption is the option for Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1ClusterResourceServiceCreate(request *clustermodel.VmwareTanzuManageV1alpha1ClusterCreateClusterRequest) (*clustermodel.VmwareTanzuManageV1alpha1ClusterCreateClusterResponse, error)

	ManageV1alpha1ClusterResourceServiceDelete(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName, force string) error

	ManageV1alpha1ClusterResourceServiceGet(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) (*clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse, error)
}

/*
  ManageV1alpha1ClusterResourceServiceCreate creates a cluster.
*/
func (a *Client) ManageV1alpha1ClusterResourceServiceCreate(request *clustermodel.VmwareTanzuManageV1alpha1ClusterCreateClusterRequest) (*clustermodel.VmwareTanzuManageV1alpha1ClusterCreateClusterResponse, error) {
	requestURL := fmt.Sprintf("%s%s", a.config.Host, "/v1alpha1/clusters")

	body, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "create tanzu TMC cluster marshall request")
	}

	headers := a.config.Headers
	headers.Set("Content-Length", fmt.Sprintf("%d", len(body)))

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
		return nil, errors.Errorf("create tanzu TMC cluster request failed with status : %v, response: %v", resp.Status, string(respBody))
	}

	clusterResponse := &clustermodel.VmwareTanzuManageV1alpha1ClusterCreateClusterResponse{}

	err = clusterResponse.UnmarshalBinary(respBody)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall")
	}

	return clusterResponse, nil
}

/*
  ManageV1alpha1ClusterResourceServiceDelete deletes a cluster.
*/
func (a *Client) ManageV1alpha1ClusterResourceServiceDelete(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName, force string) error {
	queryParams := url.Values{
		"force": []string{force},
	}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s%s%s?%s", a.config.Host, "/v1alpha1/clusters/", fn.Name, queryParams.Encode())

	resp, err := a.transport.Delete(requestURL, a.config.Headers)
	if err != nil {
		return errors.Wrap(err, "delete")
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("delete tanzu TMC cluster request failed with status : %v, response: %v", resp.Status, string(respBody))
	}

	return nil
}

/*
  ManageV1alpha1ClusterResourceServiceGet gets a cluster.
*/
func (a *Client) ManageV1alpha1ClusterResourceServiceGet(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) (*clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s%s%s?%s", a.config.Host, "/v1alpha1/clusters/", fn.Name, queryParams.Encode())

	resp, err := a.transport.Get(requestURL, a.config.Headers)
	if err != nil {
		return nil, errors.Wrap(err, "get request")
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("get tanzu TMC cluster request failed with status : %v, response: %v", resp.Status, string(respBody))
	}

	clusterResponse := &clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse{}

	err = clusterResponse.UnmarshalBinary(respBody)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall")
	}

	return clusterResponse, nil
}
