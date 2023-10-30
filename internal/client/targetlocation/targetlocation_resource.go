/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	targetlocationsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/targetlocation"
)

const (
	clustersAPIVersionAndGroup       = "v1alpha1/clusters"
	clustersListAPI                  = "dataprotection/backuplocations"
	dataProtectionAPIVersionAndGroup = "v1alpha1/dataprotection/providers"
	backupLocationsAPI               = "backuplocations"
)

// New creates a new cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for credentials resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	TargetLocationResourceServiceList(request *targetlocationsmodel.ListBackupLocationsRequest) (*targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse, error)

	TargetLocationResourceServiceCreate(request *targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest) (*targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse, error)

	TargetLocationResourceServiceUpdate(request *targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest) (*targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse, error)

	TargetLocationResourceServiceDelete(fn *targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName) error

	TargetLocationResourceServiceGet(fn *targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName) (*targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse, error)
}

/*
TargetLocationResourceServiceList lists target locations.
*/
func (c *Client) TargetLocationResourceServiceList(request *targetlocationsmodel.ListBackupLocationsRequest) (*targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse, error) {
	resp := &targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse{}

	if request.SearchScope == nil || (request.SearchScope.ClusterName == "" && request.SearchScope.ProviderName == "") {
		return nil, errors.New("scope must be set with either provider name or cluster name")
	}

	requestURL, err := buildTargetLocationsRequestURL(request)

	if err != nil {
		return nil, err
	}

	err = c.Get(requestURL, resp)

	return resp, err
}

/*
TargetLocationResourceServiceList gets a target location.
*/

func (c *Client) TargetLocationResourceServiceGet(fullName *targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName) (*targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse, error) {
	requestURL := helper.ConstructRequestURL(dataProtectionAPIVersionAndGroup, fullName.ProviderName, backupLocationsAPI, fullName.Name).String()
	resp := &targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse{}
	err := c.Get(requestURL, resp)

	return resp, err
}

/*
TargetLocationResourceServiceCreate creates a target location.
*/
func (c *Client) TargetLocationResourceServiceCreate(request *targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest) (*targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse, error) {
	response := &targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse{}
	requestURL := helper.ConstructRequestURL(dataProtectionAPIVersionAndGroup, request.BackupLocation.FullName.ProviderName, backupLocationsAPI).String()
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
TargetLocationResourceServiceUpdate updates a target location.
*/
func (c *Client) TargetLocationResourceServiceUpdate(request *targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest) (*targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse, error) {
	response := &targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse{}
	requestURL := helper.ConstructRequestURL(dataProtectionAPIVersionAndGroup, request.BackupLocation.FullName.ProviderName, backupLocationsAPI, request.BackupLocation.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}

/*
TargetLocationResourceServiceDelete deletes a target location.
*/
func (c *Client) TargetLocationResourceServiceDelete(fullName *targetlocationsmodel.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName) error {
	requestURL := helper.ConstructRequestURL(dataProtectionAPIVersionAndGroup, fullName.ProviderName, backupLocationsAPI, fullName.Name).String()

	return c.Delete(requestURL)
}

func buildTargetLocationsRequestURL(request *targetlocationsmodel.ListBackupLocationsRequest) (string, error) {
	var (
		requestURL helper.RequestURL
		requestMap map[string]interface{}
	)

	queryParams := url.Values{}
	requestJSONByte, _ := request.MarshalBinary()
	_ = json.Unmarshal(requestJSONByte, &requestMap)

	if request.SearchScope.ClusterName != "" {
		requestURL = helper.ConstructRequestURL(clustersAPIVersionAndGroup, request.SearchScope.ClusterName, clustersListAPI)
	} else {
		requestURL = helper.ConstructRequestURL(dataProtectionAPIVersionAndGroup, request.SearchScope.ProviderName, backupLocationsAPI)
	}

	err := buildQueryParams(&queryParams, "", requestMap)

	if err != nil {
		return "", errors.New("couldn't create request url")
	}

	requestURL = requestURL.AppendQueryParams(queryParams)

	return requestURL.String(), nil
}

func buildQueryParams(queryParams *url.Values, parentParam string, request map[string]interface{}) error {
	var err error

	for key, value := range request {
		if key != "clusterName" && key != "providerName" {
			parent := key

			if parentParam != "" {
				parent = fmt.Sprintf("%s.%s", parentParam, key)
			}

			switch value := value.(type) {
			case map[string]interface{}:
				err = buildQueryParams(queryParams, parent, value)

				if err != nil {
					return err
				}
			default:
				if value != nil {
					queryParams.Add(parent, helper.ConvertToString(value, ","))
				}
			}
		}
	}

	return err
}
