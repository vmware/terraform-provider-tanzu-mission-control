/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupscheduleclient

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backupschedule/cluster"
)

const (
	apiVersionAndGroup         = "v1alpha1/clusters"
	dataProtectionSchedulePath = "dataprotection/schedules"
)

// New creates a new backup schedule resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for backup schedule resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	BackupScheduleResourceServiceCreate(request *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest) (*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse, error)

	BackupScheduleResourceServiceUpdate(request *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest) (*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse, error)

	BackupScheduleResourceServiceDelete(fn *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName) error

	BackupScheduleResourceServiceGet(fn *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName) (*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse, error)

	BackupScheduleResourceServiceList(request *backupschedulemodels.ListBackupSchedulesRequest) (*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse, error)
}

/*
BackupScheduleResourceServiceCreate creates a backup schedule.
*/
func (c *Client) BackupScheduleResourceServiceCreate(request *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest) (*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse, error) {
	response := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Schedule.FullName.ClusterName, dataProtectionSchedulePath).String()
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
BackupScheduleResourceServiceUpdate updates a backup schedule.
*/
func (c *Client) BackupScheduleResourceServiceUpdate(request *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest) (*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse, error) {
	response := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Schedule.FullName.ClusterName, dataProtectionSchedulePath, request.Schedule.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}

/*
BackupScheduleResourceServiceDelete deletes a backup schedule.
*/
func (c *Client) BackupScheduleResourceServiceDelete(fullName *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName) error {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fullName.ClusterName, dataProtectionSchedulePath, fullName.Name)
	queryParams := url.Values{}

	queryParams.Add("fullName.managementClusterName", fullName.ManagementClusterName)
	queryParams.Add("fullName.provisionerName", fullName.ProvisionerName)

	requestURL = requestURL.AppendQueryParams(queryParams)

	return c.Delete(requestURL.String())
}

/*
BackupScheduleResourceServiceGet gets a backup schedule.
*/
func (c *Client) BackupScheduleResourceServiceGet(fullName *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName) (*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fullName.ClusterName, dataProtectionSchedulePath, fullName.Name)
	queryParams := url.Values{}

	queryParams.Add("fullName.managementClusterName", fullName.ManagementClusterName)
	queryParams.Add("fullName.provisionerName", fullName.ProvisionerName)
	requestURL = requestURL.AppendQueryParams(queryParams)

	resp := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse{}
	err := c.Get(requestURL.String(), resp)

	return resp, err
}

/*
BackupScheduleResourceServiceList lists backup schedules.
*/
func (c *Client) BackupScheduleResourceServiceList(request *backupschedulemodels.ListBackupSchedulesRequest) (*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse, error) {
	resp := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse{}

	if request.SearchScope == nil || request.SearchScope.ClusterName == "" {
		return nil, errors.New("scope must be set with either provider name or cluster name")
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SearchScope.ClusterName, dataProtectionSchedulePath)
	queryParams := url.Values{}

	if request.SearchScope.ManagementClusterName != "" {
		queryParams.Add("searchScope.managementClusterName", request.SearchScope.ManagementClusterName)
	}

	if request.SearchScope.ProvisionerName != "" {
		queryParams.Add("searchScope.provisionerName", request.SearchScope.ProvisionerName)
	}

	if request.SearchScope.Name != "" {
		queryParams.Add("searchScope.name", request.SearchScope.Name)
	}

	if len(queryParams) > 0 {
		requestURL = requestURL.AppendQueryParams(queryParams)
	}

	err := c.Get(requestURL.String(), resp)

	return resp, err
}
