/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspaceclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/helper"
	workspacemodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/workspace"
)

// New creates a new workspace resource service API client.
func New(transport *transport.Client, config *transport.Config) ClientService {
	return &Client{transport: transport, config: config}
}

/*
Client for workspace resource service API.
*/
type Client struct {
	transport *transport.Client
	config    *transport.Config
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1WorkspaceResourceServiceCreate(request *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest) (*workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse, error)

	ManageV1alpha1WorkspaceResourceServiceDelete(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) error

	ManageV1alpha1WorkspaceResourceServiceGet(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) (*workspacemodel.VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse, error)

	ManageV1alpha1WorkspaceResourceServiceUpdate(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest) (*workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse, error)
}

/*
  ManageV1alpha1WorkspaceResourceServiceUpdate updates a workspace.
*/
func (a *Client) ManageV1alpha1WorkspaceResourceServiceUpdate(request *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest) (*workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse, error) {
	requestURL := fmt.Sprintf("%s%s%s", a.config.Host, "/v1alpha1/workspaces/", request.Workspace.FullName.Name)

	return a.invokeAction(http.MethodPut, requestURL, request)
}

/*
  ManageV1alpha1WorkspaceResourceServiceCreate creates a workspace.
*/
func (a *Client) ManageV1alpha1WorkspaceResourceServiceCreate(request *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest) (*workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse, error) {
	requestURL := fmt.Sprintf("%s%s", a.config.Host, "/v1alpha1/workspaces")

	return a.invokeAction(http.MethodPost, requestURL, request)
}

func (a *Client) invokeAction(httpMethodType string, requestURL string, request *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest) (*workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse, error) {
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
		return nil, errors.Errorf("%s tanzu TMC workspace request failed with status : %v, response: %v", httpMethodType, resp.Status, string(respBody))
	}

	workspaceResponse := &workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse{}

	err = workspaceResponse.UnmarshalBinary(respBody)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall")
	}

	return workspaceResponse, nil
}

/*
  ManageV1alpha1WorkspaceResourceServiceGet gets a workspace.
*/
func (a *Client) ManageV1alpha1WorkspaceResourceServiceGet(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) (*workspacemodel.VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse, error) {
	requestURL := fmt.Sprintf("%s%s%s", a.config.Host, "/v1alpha1/workspaces/", fn.Name)

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
		return nil, errors.Errorf("get tanzu TMC workspace request failed with status : %v, response : %v", resp.Status, string(respBody))
	}

	workspaceResponse := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse{}

	err = workspaceResponse.UnmarshalBinary(respBody)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall")
	}

	return workspaceResponse, nil
}

/*
  ManageV1alpha1WorkspaceResourceServiceDelete deletes a workspace.
*/
func (a *Client) ManageV1alpha1WorkspaceResourceServiceDelete(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) error {
	requestURL := fmt.Sprintf("%s%s%s", a.config.Host, "/v1alpha1/workspaces/", fn.Name)

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
		return errors.Errorf("delete tanzu TMC workspace request failed with status : %v, response: %v", resp.Status, string(respBody))
	}

	return nil
}
