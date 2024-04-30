/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tapeula

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/go-test/deep"
	"github.com/jarcoal/httpmock"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tapeulamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tap/eula"
)

// nolint: unused
const (
	https                   = "https:/"
	apiVersionAndGroup      = "v1alpha1/tanzupackage"
	apiKind                 = "tap/eulas"
	queryParamKeyOrgID      = "orgId"
	queryParamKeyTAPVersion = "tapVersion"
)

// nolint: unused
func getMockData() *tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaData {
	return &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaData{
		Accepted:   true,
		EulaURL:    "https://network.tanzu.vmware.com/legal_documents/vmware_general_terms",
		ReleasedAt: strfmt.DateTime{},
		User:       "test_uer_id",
	}
}

// nolint: unparam, unused
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

// nolint: unused
func (testConfig *testAcceptanceConfig) setupHTTPMocks(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)

	endpoint := os.Getenv("TMC_ENDPOINT")

	OrgID := os.Getenv("ORG_ID")

	postRequest := &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaRequest{
		Eula: &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaEula{
			OrgID:      OrgID,
			TapVersion: tapEULATAPVersion,
		},
	}

	postResponse := &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaResponse{
		Eula: &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaEula{
			OrgID:      OrgID,
			TapVersion: tapEULATAPVersion,
			Data:       getMockData(),
		},
	}

	postEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersionAndGroup, apiKind).String()

	httpmock.RegisterResponder("POST", postEndpoint,
		bodyInspectingResponder(t, postRequest, 200, postResponse))

	getModel := &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaEula{
		OrgID:      OrgID,
		TapVersion: tapEULATAPVersion,
		Data:       getMockData(),
	}

	getResponse := &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaValidateEulaResponse{
		Eula: getModel,
	}

	queryParams := url.Values{}

	if getModel.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, getModel.OrgID)
	}

	if getModel.TapVersion != "" {
		queryParams.Add(queryParamKeyTAPVersion, getModel.TapVersion)
	}

	getTAPEULAEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersionAndGroup, apiKind).AppendQueryParams(queryParams).String()

	httpmock.RegisterResponder("GET", getTAPEULAEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))
}
