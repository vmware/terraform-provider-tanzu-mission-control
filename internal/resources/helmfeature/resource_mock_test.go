/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmfeature

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
	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

const (
	https                = "https:/"
	clAPIVersionAndGroup = "v1alpha1/clusters"
	helmAPIKind          = "fluxcd/helm"
	cgAPIVersionAndGroup = "v1alpha1/clustergroups"
)

// nolint: unparam
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

func (testConfig *testAcceptanceConfig) setupHTTPMocks(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)

	endpoint := os.Getenv("TMC_ENDPOINT")

	OrgID := os.Getenv("ORG_ID")

	reference := objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference{
		Rid: "test_rid",
		UID: "test_uid",
	}
	referenceArray := make([]*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference, 0)
	referenceArray = append(referenceArray, &reference)

	// cluster level Helm resource.
	postHelmFeatureRequest, postHelmFeatureResponse, listResponse := testConfig.getClRequestResponse(OrgID, referenceArray)

	listDeleteResp := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse{
		Helms: []*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm{
			{
				FullName: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName{
					ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
					OrgID:                 OrgID,
					ManagementClusterName: "attached",
					ProvisionerName:       "attached",
				},
				Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
					ParentReferences: referenceArray,
					Description:      "resource with description",
					Labels: map[string]string{
						"key1": "value1",
						"key2": "value2",
					},
					UID:             "gitrepository1",
					ResourceVersion: "v1",
				},
				Status: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus{
					Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
						"Disabled": {
							Reason: "feature is disable on the cluster",
						},
					},
				},
			},
		},
	}

	HelmFeatureEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, helmAPIKind)).String()

	httpmock.RegisterResponder("POST", HelmFeatureEndpoint,
		bodyInspectingResponder(t, postHelmFeatureRequest, 200, postHelmFeatureResponse))

	httpmock.RegisterResponder("GET", HelmFeatureEndpoint,
		bodyInspectingResponder(t, nil, 200, listResponse))

	// httpmock.RegisterResponder("DELETE", HelmFeatureEndpoint,
	// 	bodyInspectingResponder(t, nil, 200, nil))

	httpmock.RegisterResponder("DELETE", HelmFeatureEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", HelmFeatureEndpoint,
				bodyInspectingResponder(t, nil, 200, listDeleteResp))
		},
		http.StatusOK,
		nil))

	// cluster group level Helm resource.
	postCGHelmFeatureRequest, postCGHelmFeatureResponse, listCGResponse := testConfig.getCGRequestResponse(OrgID, referenceArray)

	CGHelmFeatureEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, helmAPIKind)).String()

	httpmock.RegisterResponder("POST", CGHelmFeatureEndpoint,
		bodyInspectingResponder(t, postCGHelmFeatureRequest, 200, postCGHelmFeatureResponse))

	httpmock.RegisterResponder("GET", CGHelmFeatureEndpoint,
		bodyInspectingResponder(t, nil, 200, listCGResponse))

	httpmock.RegisterResponder("DELETE", CGHelmFeatureEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", CGHelmFeatureEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) getCGRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest,
	*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse,
	*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse,
) {
	postRequestCDHelmFeatureModel := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm{
		FullName: &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmFullName{
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			OrgID:            orgID,
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "gitrepository1",
			ResourceVersion: "v1",
		},
	}

	postResponseCGHelmFeature := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm{
		FullName: &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmFullName{
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			OrgID:            orgID,
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: referenceArray,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "gitrepository1",
			ResourceVersion: "v1",
		},
		Status: &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	postCGHelmFeatureRequest := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest{
		Helm: postRequestCDHelmFeatureModel,
	}

	postCGHelmFeatureResponse := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse{
		Helm: postResponseCGHelmFeature,
	}

	listResponse := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse{
		Helms: []*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm{
			postResponseCGHelmFeature,
		},
	}

	return postCGHelmFeatureRequest, postCGHelmFeatureResponse, listResponse
}

func (testConfig *testAcceptanceConfig) getClRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest,
	*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse,
	*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse,
) {
	postRequestHelmFeatureModel := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm{
		FullName: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName{
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			OrgID:                 orgID,
			ManagementClusterName: "attached",
			ProvisionerName:       "attached",
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "gitrepository1",
			ResourceVersion: "v1",
		},
	}

	postResponseHelmFeature := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm{
		FullName: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName{
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			OrgID:                 orgID,
			ManagementClusterName: "attached",
			ProvisionerName:       "attached",
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "gitrepository1",
			ResourceVersion: "v1",
		},
		Status: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
		},
	}

	postHelmFeatureRequest := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest{
		Helm: postRequestHelmFeatureModel,
	}

	postHelmFeatureResponse := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse{
		Helm: postResponseHelmFeature,
	}

	listResponse := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse{
		Helms: []*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm{
			{
				FullName: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName{
					ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
					OrgID:                 orgID,
					ManagementClusterName: "attached",
					ProvisionerName:       "attached",
				},
				Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
					ParentReferences: referenceArray,
					Description:      "resource with description",
					Labels: map[string]string{
						"key1": "value1",
						"key2": "value2",
					},
					UID:             "gitrepository1",
					ResourceVersion: "v1",
				},
				Status: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus{
					Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
						"Ready": {
							Reason: "made successfully",
						},
					},
				},
			},
		},
	}

	return postHelmFeatureRequest, postHelmFeatureResponse, listResponse
}
