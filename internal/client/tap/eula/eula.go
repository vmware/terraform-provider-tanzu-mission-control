/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tapeula

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tapeula "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tap/eula"
)

const (
	apiVersionAndGroup      = "v1alpha1/tanzupackage"
	apiKind                 = "tap/eulas"
	queryParamKeyOrgID      = "orgId"
	queryParamKeyTAPVersion = "tapVersion"
)

// New creates a new EULA resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for EULA resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	EulaResourceServiceAccept(request *tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaRequest) (*tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaResponse, error)

	EulaResourceServiceValidate(e *tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaEula) (*tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaValidateEulaResponse, error)
}

/*
EulaResourceServiceAccept accepts a EULA.
*/
func (c *Client) EulaResourceServiceAccept(request *tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaRequest) (*tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, apiKind).String()
	eulaResponse := &tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaResponse{}
	err := c.Create(requestURL, request, eulaResponse)

	return eulaResponse, err
}

/*
EulaResourceServiceValidate validates a EULA.
*/
func (c *Client) EulaResourceServiceValidate(e *tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaEula) (*tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaValidateEulaResponse, error) {
	queryParams := url.Values{}

	if e.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, e.OrgID)
	}

	if e.TapVersion != "" {
		queryParams.Add(queryParamKeyTAPVersion, e.TapVersion)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, apiKind).AppendQueryParams(queryParams).String()
	eulaResponse := &tapeula.VmwareTanzuManageV1alpha1TanzupackageTapEulaValidateEulaResponse{}
	err := c.Get(requestURL, eulaResponse)

	return eulaResponse, err
}
