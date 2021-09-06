/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespaceclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/helper"
	namespacemodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/namespace"
)

// New creates a new namespace resource service API client.
func New(transport *transport.Client, config *transport.Config) ClientService {
	return &Client{transport: transport, config: config}
}

/*
Client for namespace resource service API.
*/
type Client struct {
	transport *transport.Client
	config    *transport.Config
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1NamespaceResourceServiceCreate(request *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse, error)

	ManageV1alpha1NamespaceResourceServiceDelete(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) error

	ManageV1alpha1NamespaceResourceServiceGet(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse, error)

	ManageV1alpha1NamespaceResourceServiceUpdate(request *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse, error)
}

/*
ManageV1alpha1NamespaceResourceServiceCreate creates a Namespace.
*/
func (a *Client) ManageV1alpha1NamespaceResourceServiceCreate(request *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse, error) {
	requestURL := fmt.Sprintf("%s%s%s%s", a.config.Host, "/v1alpha1/clusters/", request.Namespace.FullName.ClusterName, "/namespaces")

	return a.invokeAction(http.MethodPost, requestURL, request)
}

/*
ManageV1alpha1NamespaceResourceServiceUpdate updates a Namespace.
*/
func (a *Client) ManageV1alpha1NamespaceResourceServiceUpdate(request *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse, error) {
	requestURL := fmt.Sprintf("%s%s%s%s%s", a.config.Host, "/v1alpha1/clusters/", request.Namespace.FullName.ClusterName, "/namespaces/", request.Namespace.FullName.Name)

	return a.invokeAction(http.MethodPut, requestURL, request)
}

func (a *Client) invokeAction(httpMethodType string, requestURL string, request *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse, error) {
	body, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "marshall request body")
	}

	headers := a.config.Headers
	headers.Set(helper.ContentLengthKey, fmt.Sprintf("%d", len(body)))

	var resp *http.Response

	// nolint:bodyclose // response is being closed outside the switch block
	switch httpMethodType {
	case http.MethodPost:
		resp, err = a.transport.Post(requestURL, bytes.NewReader(body), headers)
		if err != nil {
			return nil, errors.Wrap(err, "create")
		}
	case http.MethodPut:
		resp, err = a.transport.Put(requestURL, bytes.NewReader(body), headers)
		if err != nil {
			return nil, errors.Wrap(err, "update")
		}
	default:
		return nil, errors.New("unsupported http method type invoked")
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "read %v response", httpMethodType)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("%s tanzu TMC namespace request failed with status : %v, response: %v", httpMethodType, resp.Status, string(respBody))
	}

	namespaceResponse := &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceResponse{}

	err = namespaceResponse.UnmarshalBinary(respBody)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall")
	}

	return namespaceResponse, nil
}

/*
ManageV1alpha1NamespaceResourceServiceDelete deletes a Namespace.
*/
func (a *Client) ManageV1alpha1NamespaceResourceServiceDelete(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s%s%s%s%s?%s", a.config.Host, "/v1alpha1/clusters/", fn.ClusterName, "/namespaces/", fn.Name, queryParams.Encode())

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
		return errors.Errorf("delete tanzu TMC namespace request failed with status : %v, response: %v", resp.Status, string(respBody))
	}

	return nil
}

/*
ManageV1alpha1NamespaceResourceServiceGet gets a namespace.
*/
func (a *Client) ManageV1alpha1NamespaceResourceServiceGet(fn *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) (*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s%s%s%s%s?%s", a.config.Host, "/v1alpha1/clusters/", fn.ClusterName, "/namespaces/", fn.Name, queryParams.Encode())

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
		return nil, errors.Errorf("get tanzu TMC namespace request failed with status : %v, response: %v", resp.Status, string(respBody))
	}

	namespaceResponse := &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse{}

	err = namespaceResponse.UnmarshalBinary(respBody)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall")
	}

	return namespaceResponse, nil
}
