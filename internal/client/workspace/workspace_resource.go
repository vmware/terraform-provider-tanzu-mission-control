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

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/helper"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"
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
	ManageV1alpha1WorkspaceResourceServiceCreate(request *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceRequest) (*workspacemodel.VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceResponse, error)
}

/*
  ManageV1alpha1WorkspaceResourceServiceCreate creates a workspace.
*/
func (a *Client) ManageV1alpha1WorkspaceResourceServiceCreate(request *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceRequest) (*workspacemodel.VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceResponse, error) {
	requestURL := fmt.Sprintf("%s%s", a.config.Host, "/v1alpha1/workspaces")

	body, err := request.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "marshall request body")
	}

	headers := a.config.Headers
	headers.Set(helper.ContentLength, fmt.Sprintf("%d", len(body)))

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
		return nil, errors.Errorf("create tanzu TMC workspace request failed with status : %v, response: %v", resp.Status, string(respBody))
	}

	workspaceResponse := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceResponse{}

	err = workspaceResponse.UnmarshalBinary(respBody)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall")
	}

	return workspaceResponse, nil
}
