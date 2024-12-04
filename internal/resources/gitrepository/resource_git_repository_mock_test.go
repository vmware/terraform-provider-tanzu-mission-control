// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package gitrepository

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
	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

const (
	https                = "https:/"
	clAPIVersionAndGroup = "v1alpha1/clusters"
	apiSubGroup          = "namespaces"
	apiKind              = "fluxcd/gitrepositories"
	cdAPIKind            = "fluxcd/continuousdelivery"
	cgAPIVersionAndGroup = "v1alpha1/clustergroups"
)

func getMockSpec() gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec {
	return gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
		URL:      "https://github.com/tmc-build-integrations/sample-update-configmap",
		Interval: "5m",
		Ref: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{
			Branch: "master",
		},
		GitImplementation: gitrepositoryclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT),
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
		getModel := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository{
			FullName: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName{
				Name:                  testConfig.GitRepositoryName,
				OrgID:                 OrgID,
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				NamespaceName:         testConfig.Namespace,
				ProvisionerName:       "attached",
				ManagementClusterName: "attached",
			},
			Spec: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
				URL:      "https://github.com/tmc-build-integrations/sample-update-configmap",
				Interval: "10m",
				Ref: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{
					Branch: "master",
				},
				GitImplementation: gitrepositoryclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2),
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
			Status: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryStatus{
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					"Ready": {
						Reason: "made successfully",
					},
				},
			},
		}

		getResponse := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse{
			GitRepository: getModel,
		}

		getGitRepoEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind, testConfig.GitRepositoryName)).String()

		httpmock.RegisterResponder("GET", getGitRepoEndpoint,
			bodyInspectingResponder(t, nil, 200, getResponse))

	case commonscope.ClusterGroupScope:
		getCDModel := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
			FullName: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{
				Name:             testConfig.GitRepositoryName,
				OrgID:            OrgID,
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				NamespaceName:    testConfig.Namespace,
			},
			Spec: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec{
				AtomicSpec: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
					URL:      "https://github.com/tmc-build-integrations/sample-update-configmap",
					Interval: "10m",
					Ref: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{
						Branch: "master",
					},
					GitImplementation: gitrepositoryclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2),
				},
			},
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				ParentReferences: referenceArray,
				Description:      "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				UID:             "cdgitrepository1",
				ResourceVersion: "v1",
			},
			Status: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryStatus{
				Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
			},
		}

		getCGResponse := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGetGitRepositoryResponse{
			GitRepository: getCDModel,
		}

		getCGGitRepoEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind, testConfig.GitRepositoryName)).String()

		httpmock.RegisterResponder("GET", getCGGitRepoEndpoint,
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

	// cluster level git repository resource.
	postRequest, postResponse, getResponse, postContinuousDeliveryRequest, postContinuousDeliveryResponse := testConfig.getClRequestResponse(OrgID, referenceArray)

	putRequest := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest{
		GitRepository: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository{
			FullName: postRequest.GitRepository.FullName,
			Meta:     postRequest.GitRepository.Meta,
			Spec: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
				URL:      "https://github.com/tmc-build-integrations/sample-update-configmap",
				Interval: "10m",
				Ref: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{
					Branch: "master",
				},
				GitImplementation: gitrepositoryclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2),
			},
		},
	}

	putResponse := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse{
		GitRepository: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository{
			FullName: postRequest.GitRepository.FullName,
			Meta:     postRequest.GitRepository.Meta,
			Spec: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
				URL:      "https://github.com/tmc-build-integrations/sample-update-configmap",
				Interval: "10m",
				Ref: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{
					Branch: "master",
				},
				GitImplementation: gitrepositoryclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2),
			},
			Status: postResponse.GitRepository.Status,
		},
	}

	postEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind)).String()
	getGitRepoEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind, testConfig.GitRepositoryName)).String()
	deleteEndpoint := getGitRepoEndpoint

	postContinuousDeliveryEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, cdAPIKind)).String()

	httpmock.RegisterResponder("POST", postContinuousDeliveryEndpoint,
		bodyInspectingResponder(t, postContinuousDeliveryRequest, 200, postContinuousDeliveryResponse))

	httpmock.RegisterResponder("POST", postEndpoint,
		bodyInspectingResponder(t, postRequest, 200, postResponse))

	httpmock.RegisterResponder("PUT", getGitRepoEndpoint,
		bodyInspectingResponder(t, putRequest, 200, putResponse))

	httpmock.RegisterResponder("GET", getGitRepoEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))

	httpmock.RegisterResponder("DELETE", deleteEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getGitRepoEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))

	// cluster group level git repository resource.
	postCGRequest, postCGResponse, getCGResponse, postCGContinuousDeliveryRequest, postCGContinuousDeliveryResponse := testConfig.getCGRequestResponse(OrgID, referenceArray)

	putCGRequest := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest{
		GitRepository: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
			FullName: postCGRequest.GitRepository.FullName,
			Meta:     postCGRequest.GitRepository.Meta,
			Spec: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec{
				AtomicSpec: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
					URL:      "https://github.com/tmc-build-integrations/sample-update-configmap",
					Interval: "10m",
					Ref: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{
						Branch: "master",
					},
					GitImplementation: gitrepositoryclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2),
				},
			},
		},
	}

	putCGResponse := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse{
		GitRepository: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
			FullName: postCGRequest.GitRepository.FullName,
			Meta:     postCGRequest.GitRepository.Meta,
			Spec: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec{
				AtomicSpec: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
					URL:      "https://github.com/tmc-build-integrations/sample-update-configmap",
					Interval: "10m",
					Ref: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{
						Branch: "master",
					},
					GitImplementation: gitrepositoryclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2),
				},
			},
			Status: postCGResponse.GitRepository.Status,
		},
	}

	postCGEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind)).String()
	getCGGitRepoEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind, testConfig.GitRepositoryName)).String()
	deleteCGEndpoint := getCGGitRepoEndpoint

	postCGContinuousDeliveryEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, cdAPIKind)).String()

	httpmock.RegisterResponder("POST", postCGContinuousDeliveryEndpoint,
		bodyInspectingResponder(t, postCGContinuousDeliveryRequest, 200, postCGContinuousDeliveryResponse))

	httpmock.RegisterResponder("POST", postCGEndpoint,
		bodyInspectingResponder(t, postCGRequest, 200, postCGResponse))

	httpmock.RegisterResponder("PUT", getCGGitRepoEndpoint,
		bodyInspectingResponder(t, putCGRequest, 200, putCGResponse))

	httpmock.RegisterResponder("GET", getCGGitRepoEndpoint,
		bodyInspectingResponder(t, nil, 200, getCGResponse))

	httpmock.RegisterResponder("DELETE", deleteCGEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getCGGitRepoEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) getCGRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest,
	*gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse,
	*gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGetGitRepositoryResponse,
	*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest,
	*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse,
) {
	secretSpec := getMockSpec()

	cdSpec := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec{
		AtomicSpec: &secretSpec,
	}

	postCGRequestModel := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
		FullName: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{
			Name:             testConfig.GitRepositoryName,
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
			UID:             "cdgitrepository1",
			ResourceVersion: "v1",
		},
	}

	postCGResponseModel := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
		FullName: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{
			Name:             testConfig.GitRepositoryName,
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
			UID:             "cdgitrepository1",
			ResourceVersion: "v1",
		},
		Status: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	postCGRequest := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest{
		GitRepository: postCGRequestModel,
	}

	postCGResponse := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse{
		GitRepository: postCGResponseModel,
	}

	getCDModel := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
		FullName: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{
			Name:             testConfig.GitRepositoryName,
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
			UID:             "cdgitrepository1",
			ResourceVersion: "v1",
		},
		Status: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	getCGResponse := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGetGitRepositoryResponse{
		GitRepository: getCDModel,
	}

	postRequestCDContinuousDeliveryModel := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery{
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
			UID:             "gitrepository1",
			ResourceVersion: "v1",
		},
	}

	postResponseCDContinuousDelivery := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery{
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
			UID:             "gitrepository1",
			ResourceVersion: "v1",
		},
		Status: &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	postCGContinuousDeliveryRequest := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest{
		ContinuousDelivery: postRequestCDContinuousDeliveryModel,
	}

	postCGContinuousDeliveryResponse := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse{
		ContinuousDelivery: postResponseCDContinuousDelivery,
	}

	return postCGRequest, postCGResponse, getCGResponse, postCGContinuousDeliveryRequest, postCGContinuousDeliveryResponse
}

func (testConfig *testAcceptanceConfig) getClRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest,
	*gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse,
	*gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse,
	*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest,
	*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse,
) {
	secretSpec := getMockSpec()
	postRequestModel := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository{
		FullName: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName{
			Name:                  testConfig.GitRepositoryName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &secretSpec,
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

	postResponseModel := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository{
		FullName: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName{
			Name:                  testConfig.GitRepositoryName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &secretSpec,
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
		Status: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
		},
	}

	postRequest := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest{
		GitRepository: postRequestModel,
	}

	postResponse := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse{
		GitRepository: postResponseModel,
	}

	getModel := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository{
		FullName: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName{
			Name:                  testConfig.GitRepositoryName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &secretSpec,
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
		Status: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
		},
	}

	getResponse := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse{
		GitRepository: getModel,
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
			UID:             "gitrepository1",
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
			UID:             "gitrepository1",
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
