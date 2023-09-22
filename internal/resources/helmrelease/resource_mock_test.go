/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrelease

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
	gitrepoclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	gitrepoclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
	releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
	releaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

const (
	https                = "https:/"
	clAPIVersionAndGroup = "v1alpha1/clusters"
	apiSubGroup          = "namespaces"
	apiKind              = "fluxcd/helm/releases"
	helmAPIKind          = "fluxcd/helm"
	cgAPIVersionAndGroup = "v1alpha1/clustergroups"
)

func getMockSpec() releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec {
	return releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec{
		ChartRef: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef{
			Chart:               "airflow",
			Version:             "15.0.3",
			RepositoryName:      "bitnami",
			RepositoryNamespace: "tanzu-helm-resources",
			RepositoryType:      releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeHELM),
		},
		Interval: "5m",
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
		getModel := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease{
			FullName: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName{
				Name:                  testConfig.HelmReleaseName,
				OrgID:                 OrgID,
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				NamespaceName:         testConfig.Namespace,
				ProvisionerName:       "attached",
				ManagementClusterName: "attached",
			},
			Spec: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec{
				ChartRef: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef{
					Chart:               "airflow",
					Version:             "15.0.5",
					RepositoryName:      "bitnami",
					RepositoryNamespace: "tanzu-helm-resources",
					RepositoryType:      releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeHELM),
				},
				Interval: "5m",
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
			Status: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus{
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					"Ready": {
						Reason: "made successfully",
					},
				},
				GeneratedResources: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources{
					ClusterRoleName:    "testclusterrole",
					RoleBindingName:    "testrolebinding",
					ServiceAccountName: "testserviceaccount",
				},
			},
		}

		getResponse := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse{
			Release: getModel,
		}

		getEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind, testConfig.HelmReleaseName)).String()

		httpmock.RegisterResponder("GET", getEndpoint,
			bodyInspectingResponder(t, nil, 200, getResponse))

	case commonscope.ClusterGroupScope:
		getCDModel := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease{
			FullName: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName{
				Name:             testConfig.HelmReleaseName,
				OrgID:            OrgID,
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				NamespaceName:    testConfig.Namespace,
			},
			Spec: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec{
				AtomicSpec: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec{
					ChartRef: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef{
						Chart:               "constraints",
						RepositoryName:      "test-git-repo",
						RepositoryNamespace: "tanzu-continuousdelivery-resources",
						RepositoryType:      releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeGIT),
					},
					Interval: "5m",
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
			Status: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus{
				Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
			},
		}

		getCGResponse := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseGetResponse{
			Release: getCDModel,
		}

		getCGHelmReleaseEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind, testConfig.HelmReleaseName)).String()

		httpmock.RegisterResponder("GET", getCGHelmReleaseEndpoint,
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

	// cluster level Helm resource.
	postRequest, postResponse, getResponse, postHelmFeatureRequest, postHelmFeatureResponse := testConfig.getClRequestResponse(OrgID, referenceArray)

	putRequest := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest{
		Release: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease{
			FullName: postRequest.Release.FullName,
			Meta:     postRequest.Release.Meta,
			Spec: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec{
				ChartRef: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef{
					Chart:               "airflow",
					Version:             constraints,
					RepositoryName:      "bitnami",
					RepositoryNamespace: "tanzu-helm-resources",
					RepositoryType:      releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeHELM),
				},
				Interval: "5m",
			},
		},
	}

	putResponse := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse{
		Release: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease{
			FullName: postRequest.Release.FullName,
			Meta:     postRequest.Release.Meta,
			Spec: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec{
				ChartRef: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef{
					Chart:               "airflow",
					Version:             constraints,
					RepositoryName:      "bitnami",
					RepositoryNamespace: "tanzu-helm-resources",
					RepositoryType:      releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeHELM),
				},
				Interval: "5m",
			},
			Status: postResponse.Release.Status,
		},
	}

	listRespponse := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse{
		Helms: []*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm{
			postHelmFeatureResponse.Helm,
		},
	}

	postEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind)).String()
	getHelmReleaseEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind, testConfig.HelmReleaseName)).String()
	deleteEndpoint := getHelmReleaseEndpoint

	postHelmFeatureEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, helmAPIKind)).String()

	httpmock.RegisterResponder("POST", postHelmFeatureEndpoint,
		bodyInspectingResponder(t, postHelmFeatureRequest, 200, postHelmFeatureResponse))

	httpmock.RegisterResponder("GET", postHelmFeatureEndpoint,
		bodyInspectingResponder(t, nil, 200, listRespponse))

	httpmock.RegisterResponder("POST", postEndpoint,
		bodyInspectingResponder(t, postRequest, 200, postResponse))

	httpmock.RegisterResponder("PUT", getHelmReleaseEndpoint,
		bodyInspectingResponder(t, putRequest, 200, putResponse))

	httpmock.RegisterResponder("GET", getHelmReleaseEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))

	httpmock.RegisterResponder("DELETE", deleteEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getHelmReleaseEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))

	// cluster group level Helm resource.
	postCGRequest, postCGResponse, getCGResponse, postCGHelmFeatureRequest, postCGHelmFeatureResponse, postCGgitRepoRequest, postCGgitRepoResponse := testConfig.getCGRequestResponse(OrgID, referenceArray)

	putCGRequest := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest{
		Release: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease{
			FullName: postCGRequest.Release.FullName,
			Meta:     postCGRequest.Release.Meta,
			Spec: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec{
				AtomicSpec: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec{
					ChartRef: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef{
						Chart:               "constraints",
						RepositoryName:      "test-git-repo",
						RepositoryNamespace: "tanzu-continuousdelivery-resources",
						RepositoryType:      releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeGIT),
					},
					Interval: "5m",
				},
			},
		},
	}

	putCGResponse := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest{
		Release: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease{
			FullName: postCGRequest.Release.FullName,
			Meta:     postCGRequest.Release.Meta,
			Spec: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec{
				AtomicSpec: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec{
					ChartRef: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef{
						Chart:               "constraints",
						RepositoryName:      "test-git-repo",
						RepositoryNamespace: "tanzu-continuousdelivery-resources",
						RepositoryType:      releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeGIT),
					},
					Interval: "5m",
				},
			},
			Status: postCGResponse.Release.Status,
		},
	}

	listCGResponse := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse{
		Helms: []*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm{
			postCGHelmFeatureResponse.Helm,
		},
	}

	postCGEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind)).String()
	getCGHelmReleaseEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind, testConfig.HelmReleaseName)).String()
	deleteCGEndpoint := getCGHelmReleaseEndpoint

	postCGgitRepoEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", "fluxcd/gitrepositories")).String()

	postCGHelmFeatureEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, helmAPIKind)).String()

	httpmock.RegisterResponder("POST", postCGHelmFeatureEndpoint,
		bodyInspectingResponder(t, postCGHelmFeatureRequest, 200, postCGHelmFeatureResponse))

	httpmock.RegisterResponder("GET", postCGHelmFeatureEndpoint,
		bodyInspectingResponder(t, nil, 200, listCGResponse))

	httpmock.RegisterResponder("POST", postCGgitRepoEndpoint,
		bodyInspectingResponder(t, postCGgitRepoRequest, 200, postCGgitRepoResponse))

	httpmock.RegisterResponder("POST", postCGEndpoint,
		bodyInspectingResponder(t, postCGRequest, 200, postCGResponse))

	httpmock.RegisterResponder("PUT", getCGHelmReleaseEndpoint,
		bodyInspectingResponder(t, putCGRequest, 200, putCGResponse))

	httpmock.RegisterResponder("GET", getCGHelmReleaseEndpoint,
		bodyInspectingResponder(t, nil, 200, getCGResponse))

	httpmock.RegisterResponder("DELETE", deleteCGEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getCGHelmReleaseEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) getCGRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest,
	*releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse,
	*releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseGetResponse,
	*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest,
	*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse,
	*gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest,
	*gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse,
) {
	cdSpec := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec{
		AtomicSpec: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec{
			ChartRef: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef{
				Chart:               "manifests",
				RepositoryName:      "test-git-repo",
				RepositoryNamespace: "tanzu-continuousdelivery-resources",
				RepositoryType:      releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeGIT),
			},
			Interval: "5m",
		},
	}

	postCGRequestModel := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease{
		FullName: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName{
			Name:             testConfig.HelmReleaseName,
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

	postCGResponseModel := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease{
		FullName: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName{
			Name:             testConfig.HelmReleaseName,
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
		Status: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	postCGRequest := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest{
		Release: postCGRequestModel,
	}

	postCGResponse := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse{
		Release: postCGResponseModel,
	}

	getCDModel := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease{
		FullName: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName{
			Name:             testConfig.HelmReleaseName,
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
		Status: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	getCGResponse := &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseGetResponse{
		Release: getCDModel,
	}

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
			ParentReferences: nil,
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

	postCGgitRepoRequestModel := &gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
		FullName: &gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{
			Name:             gitRepoName,
			OrgID:            orgID,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			NamespaceName:    testConfig.Namespace,
		},
		Spec: &gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec{
			AtomicSpec: &gitrepoclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
				URL: "https://github.com/tmc-build-integrations/sample-update-configmap",
			},
		},
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

	postCGgitRepoResponseModel := &gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
		FullName: &gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{
			Name:             gitRepoName,
			OrgID:            orgID,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			NamespaceName:    testConfig.Namespace,
		},
		Spec: &gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec{
			AtomicSpec: &gitrepoclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
				URL: "https://github.com/tmc-build-integrations/sample-update-configmap",
			},
		},
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
		Status: &gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	postCGgitRepoRequest := &gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest{
		GitRepository: postCGgitRepoRequestModel,
	}

	postCGgitRepoResponse := &gitrepoclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse{
		GitRepository: postCGgitRepoResponseModel,
	}

	return postCGRequest, postCGResponse, getCGResponse, postCGHelmFeatureRequest, postCGHelmFeatureResponse, postCGgitRepoRequest, postCGgitRepoResponse
}

func (testConfig *testAcceptanceConfig) getClRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest,
	*releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse,
	*releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse,
	*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest,
	*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse,
) {
	secretSpec := getMockSpec()
	postRequestModel := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease{
		FullName: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName{
			Name:                  testConfig.HelmReleaseName,
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

	postResponseModel := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease{
		FullName: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName{
			Name:                  testConfig.HelmReleaseName,
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
		Status: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
			GeneratedResources: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources{
				ClusterRoleName:    "testclusterrole",
				RoleBindingName:    "testrolebinding",
				ServiceAccountName: "testserviceaccount",
			},
		},
	}

	postRequest := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest{
		Release: postRequestModel,
	}

	postResponse := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse{
		Release: postResponseModel,
	}

	getModel := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease{
		FullName: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName{
			Name:                  testConfig.HelmReleaseName,
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
		Status: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
			GeneratedResources: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources{
				ClusterRoleName:    "testclusterrole",
				RoleBindingName:    "testrolebinding",
				ServiceAccountName: "testserviceaccount",
			},
		},
	}

	getResponse := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse{
		Release: getModel,
	}

	postRequestHelmFeatureModel := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm{
		FullName: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName{
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

	postResponseHelmFeature := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm{
		FullName: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName{
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

	return postRequest, postResponse, getResponse, postHelmFeatureRequest, postHelmFeatureResponse
}
