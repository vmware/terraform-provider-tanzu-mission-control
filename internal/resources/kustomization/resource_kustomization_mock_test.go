// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kustomization

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
	continuousdeliveryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/continuousdelivery/cluster"
	continuousdeliveryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/continuousdelivery/clustergroup"
	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

const (
	https                = "https:/"
	clAPIVersionAndGroup = "v1alpha1/clusters"
	apiSubGroup          = "namespaces"
	apiKind              = "fluxcd/kustomizations"
	cdAPIKind            = "fluxcd/continuousdelivery"
	cgAPIVersionAndGroup = "v1alpha1/clustergroups"
)

func getMockSpec() kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec {
	return kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{
		Path:     "manifests/",
		Interval: "5m",
		Prune:    false,
		Source: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{
			Name:      "someGitRepository",
			Namespace: "tanzu-continuousdelivery-resources",
		},
	}
}

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

func (testConfig *testAcceptanceConfig) setupHTTPMocksUpdate(t *testing.T, scope commonscope.Scope) {
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

	switch scope {
	case commonscope.ClusterScope:
		getModel := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization{
			FullName: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName{
				Name:                  testConfig.KustomizationName,
				OrgID:                 OrgID,
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				NamespaceName:         testConfig.Namespace,
				ProvisionerName:       "attached",
				ManagementClusterName: "attached",
			},
			Spec: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{
				Path:     "manifests/",
				Interval: "10m",
				Prune:    true,
				Source: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{
					Name:      "someGitRepository",
					Namespace: "tanzu-continuousdelivery-resources",
				},
			},
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				ParentReferences: referenceArray,
				Description:      "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				UID:             "kustomization1",
				ResourceVersion: "v1",
			},
			Status: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationStatus{
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					"Ready": {
						Reason: "made successfully",
					},
				},
			},
		}

		getResponse := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationGetKustomizationResponse{
			Kustomization: getModel,
		}
		getKustomizationEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind, testConfig.KustomizationName)).String()

		httpmock.RegisterResponder("GET", getKustomizationEndpoint,
			bodyInspectingResponder(t, nil, 200, getResponse))
	case commonscope.ClusterGroupScope:
		getCGModel := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization{
			FullName: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName{
				Name:             testConfig.KustomizationName,
				OrgID:            OrgID,
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				NamespaceName:    testConfig.Namespace,
			},
			Spec: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec{
				AtomicSpec: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{
					Path:     "manifests/",
					Interval: "10m",
					Prune:    true,
					Source: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{
						Name:      "someGitRepository",
						Namespace: "tanzu-continuousdelivery-resources",
					},
				},
			},
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				ParentReferences: referenceArray,
				Description:      "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				UID:             "cdkustomization1",
				ResourceVersion: "v1",
			},
			Status: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationStatus{
				Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
			},
		}

		getCGResponse := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse{
			Kustomization: getCGModel,
		}

		getCGKustomizationEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind, testConfig.KustomizationName)).String()

		httpmock.RegisterResponder("GET", getCGKustomizationEndpoint,
			bodyInspectingResponder(t, nil, 200, getCGResponse))
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

	// cluster level Kustomization resorce.
	postRequest, postResponse, getResponse, postContinuousDeliveryRequest, postContinuousDeliveryResponse := testConfig.getClRequestResponse(OrgID, referenceArray)

	putRequest := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest{
		Kustomization: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization{
			FullName: postRequest.Kustomization.FullName,
			Meta:     postRequest.Kustomization.Meta,
			Spec: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{
				Path:     "manifests/",
				Interval: "10m",
				Prune:    true,
				Source: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{
					Name:      "someGitRepository",
					Namespace: "tanzu-continuousdelivery-resources",
				},
			},
		},
	}

	putResponse := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse{
		Kustomization: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization{
			FullName: postRequest.Kustomization.FullName,
			Meta:     postRequest.Kustomization.Meta,
			Spec: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{
				Path:     "manifests/",
				Interval: "10m",
				Prune:    true,
				Source: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{
					Name:      "someGitRepository",
					Namespace: "tanzu-continuousdelivery-resources",
				},
			},
			Status: postResponse.Kustomization.Status,
		},
	}

	postEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind)).String()
	getKustomizationEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind, testConfig.KustomizationName)).String()
	deleteEndpoint := getKustomizationEndpoint

	postContinuousDeliveryEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, cdAPIKind)).String()

	httpmock.RegisterResponder("POST", postContinuousDeliveryEndpoint,
		bodyInspectingResponder(t, postContinuousDeliveryRequest, 200, postContinuousDeliveryResponse))

	httpmock.RegisterResponder("POST", postEndpoint,
		bodyInspectingResponder(t, postRequest, 200, postResponse))

	httpmock.RegisterResponder("PUT", getKustomizationEndpoint,
		bodyInspectingResponder(t, putRequest, 200, putResponse))

	httpmock.RegisterResponder("GET", getKustomizationEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))

	httpmock.RegisterResponder("DELETE", deleteEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getKustomizationEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))

	// cluster group level kustomization resource.
	postCGRequest, postCGResponse, getCGResponse, postCGContinuousDeliveryRequest, postCGContinuousDeliveryResponse := testConfig.getCGRequestResponse(OrgID, referenceArray)

	putCGRequest := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest{
		Kustomization: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization{
			FullName: postCGRequest.Kustomization.FullName,
			Meta:     postResponse.Kustomization.Meta,
			Spec: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec{
				AtomicSpec: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{
					Path:     "manifests/",
					Interval: "10m",
					Prune:    true,
					Source: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{
						Name:      "someGitRepository",
						Namespace: "tanzu-continuousdelivery-resources",
					},
				},
			},
		},
	}

	putCGResponse := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse{
		Kustomization: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization{
			FullName: postCGRequest.Kustomization.FullName,
			Meta:     postResponse.Kustomization.Meta,
			Spec: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec{
				AtomicSpec: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{
					Path:     "manifests/",
					Interval: "10m",
					Prune:    true,
					Source: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{
						Name:      "someGitRepository",
						Namespace: "tanzu-continuousdelivery-resources",
					},
				},
			},
			Status: postCGResponse.Kustomization.Status,
		},
	}

	postCGEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind)).String()
	getCGKustomizationEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind, testConfig.KustomizationName)).String()
	deleteCGEndpoint := getCGKustomizationEndpoint

	postCDContinuousDeliveryEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, cdAPIKind)).String()

	httpmock.RegisterResponder("POST", postCDContinuousDeliveryEndpoint,
		bodyInspectingResponder(t, postCGContinuousDeliveryRequest, 200, postCGContinuousDeliveryResponse))

	httpmock.RegisterResponder("POST", postCGEndpoint,
		bodyInspectingResponder(t, postCGRequest, 200, postCGResponse))

	httpmock.RegisterResponder("GET", getCGKustomizationEndpoint,
		bodyInspectingResponder(t, nil, 200, getCGResponse))

	httpmock.RegisterResponder("PUT", getCGKustomizationEndpoint,
		bodyInspectingResponder(t, putCGRequest, 200, putCGResponse))

	httpmock.RegisterResponder("DELETE", deleteCGEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getCGKustomizationEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) getCGRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest,
	*kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse,
	*kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse,
	*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest,
	*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse,
) {
	kustomizationSpec := getMockSpec()

	cdSpec := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec{
		AtomicSpec: &kustomizationSpec,
	}

	postCGRequestModel := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization{
		FullName: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName{
			Name:             testConfig.KustomizationName,
			OrgID:            orgID,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			NamespaceName:    testConfig.Namespace,
		},
		Spec: cdSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "cdkustomization1",
			ResourceVersion: "v1",
		},
	}

	postCGResponseModel := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization{
		FullName: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName{
			Name:             testConfig.KustomizationName,
			OrgID:            orgID,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			NamespaceName:    testConfig.Namespace,
		},
		Spec: cdSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "cdkustomization1",
			ResourceVersion: "v1",
		},
		Status: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	postCGRequest := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest{
		Kustomization: postCGRequestModel,
	}

	postCGResponse := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse{
		Kustomization: postCGResponseModel,
	}

	getCGModel := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization{
		FullName: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName{
			Name:             testConfig.KustomizationName,
			OrgID:            orgID,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			NamespaceName:    testConfig.Namespace,
		},
		Spec: cdSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: referenceArray,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "cdkustomization1",
			ResourceVersion: "v1",
		},
		Status: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	getCGResponse := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse{
		Kustomization: getCGModel,
	}

	postRequestCGContinuousDeliveryModel := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery{
		FullName: &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName{
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
			UID:             "cdcontinuousdelivery1",
			ResourceVersion: "v1",
		},
	}

	postResponseCGContinuousDelivery := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery{
		FullName: &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName{
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
			UID:             "cdcontinuousdelivery1",
			ResourceVersion: "v1",
		},
		Status: &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	postCGContinuousDeliveryRequest := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest{
		ContinuousDelivery: postRequestCGContinuousDeliveryModel,
	}

	postCGContinuousDeliveryResponse := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse{
		ContinuousDelivery: postResponseCGContinuousDelivery,
	}

	return postCGRequest, postCGResponse, getCGResponse, postCGContinuousDeliveryRequest, postCGContinuousDeliveryResponse
}

func (testConfig *testAcceptanceConfig) getClRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest,
	*kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse,
	*kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationGetKustomizationResponse,
	*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest,
	*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse,
) {
	kustomizationSpec := getMockSpec()
	postRequestModel := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization{
		FullName: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName{
			Name:                  testConfig.KustomizationName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &kustomizationSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "kustomization1",
			ResourceVersion: "v1",
		},
	}

	postResponseModel := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization{
		FullName: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName{
			Name:                  testConfig.KustomizationName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &kustomizationSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "kustomization1",
			ResourceVersion: "v1",
		},
		Status: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
		},
	}

	postRequest := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest{
		Kustomization: postRequestModel,
	}

	postResponse := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse{
		Kustomization: postResponseModel,
	}

	getModel := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization{
		FullName: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName{
			Name:                  testConfig.KustomizationName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &kustomizationSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: referenceArray,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "kustomization1",
			ResourceVersion: "v1",
		},
		Status: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
		},
	}

	getResponse := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationGetKustomizationResponse{
		Kustomization: getModel,
	}

	postRequestContinuousDeliveryModel := &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery{
		FullName: &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryFullName{
			ClusterName: testConfig.ScopeHelperResources.Cluster.Name,
			OrgID:       orgID,
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "continuousdelivery1",
			ResourceVersion: "v1",
		},
	}

	postResponseContinuousDelivery := &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery{
		FullName: &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryFullName{
			ClusterName: testConfig.ScopeHelperResources.Cluster.Name,
			OrgID:       orgID,
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "continuousdelivery1",
			ResourceVersion: "v1",
		},
		Status: &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
		},
	}

	postContinuousDeliveryRequest := &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest{
		ContinuousDelivery: postRequestContinuousDeliveryModel,
	}

	postContinuousDeliveryResponse := &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse{
		ContinuousDelivery: postResponseContinuousDelivery,
	}

	return postRequest, postResponse, getResponse, postContinuousDeliveryRequest, postContinuousDeliveryResponse
}
