// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package backupscheduleclustergroupclient

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	backupscheduleclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backupschedule/clustergroup"
)

const (
	apiVersionAndGroup         = "v1alpha1/clustergroups"
	dataProtectionSchedulePath = "dataprotection/schedules"
	queryParamKeyName          = "fullName.name"
	queryParamKeyOrgID         = "fullName.orgID"
)

// New creates a new schedule resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for schedule resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceCreate(request *backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleRequest) (*backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceDelete(fn *backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleFullName) error

	VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceGet(fn *backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleFullName) (*backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleGetScheduleResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceList(request *backupscheduleclustergroupmodel.ListClusterGroupBackupSchedulesRequest) (*backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleListSchedulesResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceUpdate(request *backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleRequest) (*backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceCreate creates a schedule.
*/
func (a *Client) VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceCreate(
	request *backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleRequest) (
	*backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleResponse, error) {
	response := &backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Schedule.FullName.ClusterGroupName, dataProtectionSchedulePath).String()
	err := a.Create(requestURL, request, response)

	return response, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceUpdate updates overwrite a schedule.
*/
func (a *Client) VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceUpdate(
	request *backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleRequest,
) (*backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleResponse, error) {
	response := &backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleScheduleResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Schedule.FullName.ClusterGroupName, dataProtectionSchedulePath,
		request.Schedule.FullName.Name).String()
	err := a.Update(requestURL, request, response)

	return response, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceDelete deletes a schedule.
*/
func (a *Client) VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceDelete(fullName *backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleFullName) error {
	queryParams := url.Values{}

	if fullName.Name != "" {
		queryParams.Add(queryParamKeyName, fullName.Name)
	}

	if fullName.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fullName.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fullName.ClusterGroupName, dataProtectionSchedulePath, fullName.Name).AppendQueryParams(queryParams).String()

	return a.Delete(requestURL)
}

/*
VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceGet gets a schedule.
*/
func (a *Client) VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceGet(
	fullName *backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleFullName) (
	*backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleGetScheduleResponse, error) {
	queryParams := url.Values{}

	if fullName.Name != "" {
		queryParams.Add(queryParamKeyName, fullName.Name)
	}

	if fullName.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fullName.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fullName.ClusterGroupName, dataProtectionSchedulePath, fullName.Name).AppendQueryParams(queryParams).String()
	backupScheduleClusterGroupResponse := &backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleGetScheduleResponse{}
	err := a.Get(requestURL, backupScheduleClusterGroupResponse)

	return backupScheduleClusterGroupResponse, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceList lists schedules.
*/
func (a *Client) VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceList(
	request *backupscheduleclustergroupmodel.ListClusterGroupBackupSchedulesRequest) (
	*backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleListSchedulesResponse, error) {
	resp := &backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleListSchedulesResponse{}

	if request.SearchScope == nil {
		return nil, errors.New("nil search scope")
	}

	if request.SearchScope.ClusterGroupName == "" {
		return nil, errors.New("scope must be set with cluster group name")
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SearchScope.ClusterGroupName, dataProtectionSchedulePath)
	queryParams := url.Values{}

	queryParams.Add("searchScope.clusterGroupName", request.SearchScope.ClusterGroupName)
	requestURL = requestURL.AppendQueryParams(queryParams)

	err := a.Get(requestURL.String(), resp)

	return resp, err
}
