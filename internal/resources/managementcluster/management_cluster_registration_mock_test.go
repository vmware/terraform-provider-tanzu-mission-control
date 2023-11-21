/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementcluster

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/jarcoal/httpmock"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	registrationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/managementcluster"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

func setupHTTPMocks(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)
}

const (
	https      = "https:/"
	apiVersion = "v1alpha1"
	groups     = "managementclusters"
)

func setUpOrgPolicyEndPointMocks(t *testing.T, endpoint string, resourceName string, kubernetesProviderTypeValue *clustermodel.VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType) {
	managementClusterRequest := &registrationmodel.VmwareTanzuManageV1alpha1ManagementclusterCreateManagementClusterRequest{}

	managementClusterResponse := &registrationmodel.VmwareTanzuManageV1alpha1ManagementclusterCreateManagementClusterResponse{
		ManagementCluster: &registrationmodel.VmwareTanzuManageV1alpha1ManagementclusterManagementCluster{
			FullName: &registrationmodel.VmwareTanzuManageV1alpha1ManagementclusterFullName{
				Name: resourceName,
			},
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				UID: "1234",
			},
			Spec: &registrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec{
				DefaultClusterGroup:    "default",
				KubernetesProviderType: kubernetesProviderTypeValue,
			},
			Status: &registrationmodel.VmwareTanzuManageV1alpha1ManagementclusterStatus{
				KubernetesProvider: &clustermodel.VmwareTanzuManageV1alpha1CommonClusterKubernetesProvider{
					Type: kubernetesProviderTypeValue,
				},
				Phase: registrationmodel.NewVmwareTanzuManageV1alpha1ManagementclusterPhase(registrationmodel.VmwareTanzuManageV1alpha1ManagementclusterPhaseDELETING),
			},
		}}

	createManagementClusterRegistrationEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, groups).String()
	readManagementClusterRegistrationEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, groups, resourceName).String()
	deleteManagementClusterRegistrationResourceEndpoint := readManagementClusterRegistrationEndpoint

	httpmock.RegisterResponder("POST", createManagementClusterRegistrationEndpoint,
		bodyInspectingResponder(t, managementClusterRequest, http.StatusOK, managementClusterResponse))

	httpmock.RegisterResponder("GET", readManagementClusterRegistrationEndpoint,
		bodyInspectingResponder(t, nil, http.StatusOK, managementClusterResponse))

	httpmock.RegisterResponder("DELETE", deleteManagementClusterRegistrationResourceEndpoint, changeStateResponder(
		func() {
			httpmock.RegisterResponder("GET", readManagementClusterRegistrationEndpoint,
				httpmock.NewStringResponder(http.StatusBadRequest, "Not found"))
		},
		http.StatusOK,
		nil))
}

func bodyInspectingResponder(t *testing.T, expectedContent interface{}, successResponse int, successResponseBody interface{}) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		successFunc := func() (*http.Response, error) {
			return httpmock.NewJsonResponse(successResponse, successResponseBody)
		}

		if expectedContent == nil {
			return successFunc()
		}

		// Compare to expected content.
		expectedBytes, err := json.Marshal(expectedContent)
		if err != nil {
			t.Fail()
			return nil, err
		}

		if r.Body == nil {
			t.Fail()
			return nil, fmt.Errorf("expected body on request")
		}

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fail()
			return nil, err
		}

		var bodyInterface map[string]interface{}
		if err = json.Unmarshal(bodyBytes, &bodyInterface); err == nil {
			var expectedInterface map[string]interface{}

			err = json.Unmarshal(expectedBytes, &expectedInterface)
			if err != nil {
				return nil, err
			}

			diff := deep.Equal(bodyInterface, expectedInterface)
			if diff == nil {
				return successFunc()
			}
		} else {
			return nil, err
		}

		return successFunc()
	}
}

// Register a new responder when the given call is made.
func changeStateResponder(registerFunc func(), successResponse int, successResponseBody interface{}) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		registerFunc()
		return httpmock.NewJsonResponse(successResponse, successResponseBody)
	}
}
