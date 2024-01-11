/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provisioner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/go-test/deep"
	"github.com/jarcoal/httpmock"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	provisionermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/provisioner"
)

func setupHTTPMocks(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)
}

const (
	https        = "https:/"
	apiVersion   = "v1alpha1"
	mgmtClusters = "managementclusters"
	provisioners = "provisioners"
)

func setUpProvisionerEndPointMocks(t *testing.T, endpoint, prvName string) {
	provisionerRequest := &provisionermodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerCreateProvisionerRequest{}

	provisionerResponse := &provisionermodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerCreateProvisionerResponse{
		Provisioner: &provisionermodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerProvisioner{
			FullName: &provisionermodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerFullName{
				ManagementClusterName: eksManagementCluster,
				Name:                  prvName,
			},
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				Description: "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				UID:             "dummy-uid-1234",
				ResourceVersion: "res-version",
			},
		},
	}

	createProvisionerEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, mgmtClusters, eksManagementCluster, provisioners).String()
	readProvisionerEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, mgmtClusters, eksManagementCluster, provisioners, prvName).String()
	deleteProvisionerEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, mgmtClusters, eksManagementCluster, provisioners, prvName).String()

	httpmock.RegisterResponder("POST", createProvisionerEndpoint,
		bodyInspectingResponder(t, provisionerRequest, http.StatusOK, provisionerResponse))

	httpmock.RegisterResponder("GET", readProvisionerEndpoint,
		bodyInspectingResponder(t, nil, http.StatusOK, provisionerResponse))

	httpmock.RegisterResponder("PUT", readProvisionerEndpoint,
		bodyInspectingResponder(t, nil, http.StatusOK, provisionerResponse))

	httpmock.RegisterResponder("GET", readProvisionerEndpoint,
		bodyInspectingResponder(t, nil, http.StatusOK, provisionerResponse))

	httpmock.RegisterResponder("DELETE", deleteProvisionerEndpoint, changeStateResponder(
		func() {
			httpmock.RegisterResponder("GET", readProvisionerEndpoint,
				httpmock.NewStringResponder(http.StatusBadRequest, "Not found"))
		},
		http.StatusOK,
		nil))
}

func setUpHTTPMockUpdate(t *testing.T, prvName string) {
	endpoint := os.Getenv("TMC_ENDPOINT")

	provisionerResponse := &provisionermodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerCreateProvisionerResponse{
		Provisioner: &provisionermodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerProvisioner{
			FullName: &provisionermodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerFullName{
				ManagementClusterName: eksManagementCluster,
				Name:                  prvName,
			},
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				Description: "resource with updated description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				UID:             "dummy-uid-1234",
				ResourceVersion: "res-version",
			},
		},
	}

	readProvisionerEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, mgmtClusters, eksManagementCluster, provisioners, prvName).String()

	httpmock.RegisterResponder("GET", readProvisionerEndpoint,
		bodyInspectingResponder(t, nil, 200, provisionerResponse))
}

//nolint:unparam
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
