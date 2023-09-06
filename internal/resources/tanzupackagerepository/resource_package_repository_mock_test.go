/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

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
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	tanzupakageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackage"
	pkgrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
)

const (
	https                = "https:/"
	clAPIVersionAndGroup = "v1alpha1/clusters"
	apiSubGroup          = "namespaces"
	apiKind              = "tanzupackage/repositories"
	availabilityAPIKind  = "tanzupackage/repositories:setavailability"
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

func (testConfig *testAcceptanceConfig) setupHTTPMocksUpdate(t *testing.T) {
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

	getModel := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
		FullName: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName{
			Name:                  testConfig.PkgRepoName,
			OrgID:                 OrgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         globalRepoNamespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec{
			ImgpkgBundle: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec{
				Image: "extensions.aws-usw2.tmc-dev.cloud.vmware.com/packages/standard/repo:v2.2.0_update.1",
			},
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: referenceArray,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "pkgrepository1",
			ResourceVersion: "v1",
		},
		Status: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
			Disabled:   true,
			Managed:    false,
			Subscribed: false,
		},
	}

	getResponse := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryGetResponse{
		Repository: getModel,
	}

	getPkgRepoEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, globalRepoNamespace, apiKind, testConfig.PkgRepoName)).String()

	httpmock.RegisterResponder("GET", getPkgRepoEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))
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

	// cluster level package repository resource.
	postRequest, postResponse, getResponse, postAvailabilityRequestRequest, postAvailabilityResponse := testConfig.getClRequestResponse(OrgID, referenceArray)

	putRequest := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest{
		Repository: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
			FullName: postRequest.Repository.FullName,
			Meta:     postRequest.Repository.Meta,
			Spec: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec{
				ImgpkgBundle: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec{
					Image: "extensions.aws-usw2.tmc-dev.cloud.vmware.com/packages/standard/repo:v2.2.0_update.1",
				},
			},
		},
	}

	putResponse := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest{
		Repository: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
			FullName: postResponse.Repository.FullName,
			Meta:     postResponse.Repository.Meta,
			Spec: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec{
				ImgpkgBundle: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec{
					Image: "extensions.aws-usw2.tmc-dev.cloud.vmware.com/packages/standard/repo:v2.2.0_update.1",
				},
			},
			Status: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus{
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					"Ready": {
						Reason: "made successfully",
					},
				},
				Disabled:   true,
				Managed:    false,
				Subscribed: false,
			},
		},
	}

	getTanzuPackageResponse := &tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse{
		TanzuPackages: []*tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageTanzuPackage{
			{
				FullName: &tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageFullName{
					ClusterName:           postRequest.Repository.FullName.ClusterName,
					ManagementClusterName: postRequest.Repository.FullName.ManagementClusterName,
					ProvisionerName:       postRequest.Repository.FullName.ProvisionerName,
					OrgID:                 postRequest.Repository.FullName.OrgID,
				},
				Status: &tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageStatus{
					Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
						"Ready": {
							Reason: "made successfully",
						},
					},
					PackageRepositoryGlobalNamespace: globalRepoNamespace,
				},
			},
		},
	}

	postEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, globalRepoNamespace, apiKind)).String()
	getTanzuPackageEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, "tanzupackage")).String()
	getPkgRepoEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, globalRepoNamespace, apiKind, testConfig.PkgRepoName)).String()
	deleteEndpoint := getPkgRepoEndpoint

	postAvailabilityRequestEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, availabilityAPIKind)).String()

	httpmock.RegisterResponder("POST", postAvailabilityRequestEndpoint,
		bodyInspectingResponder(t, postAvailabilityRequestRequest, 200, postAvailabilityResponse))

	httpmock.RegisterResponder("POST", postEndpoint,
		bodyInspectingResponder(t, postRequest, 200, postResponse))

	httpmock.RegisterResponder("PUT", getPkgRepoEndpoint,
		bodyInspectingResponder(t, putRequest, 200, putResponse))

	httpmock.RegisterResponder("GET", getPkgRepoEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))

	httpmock.RegisterResponder("GET", getTanzuPackageEndpoint,
		bodyInspectingResponder(t, nil, 200, getTanzuPackageResponse))

	httpmock.RegisterResponder("DELETE", deleteEndpoint, changeStateResponder(
		// Set up the get to return 404 after the package repository has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getPkgRepoEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) getClRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest,
	*pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryResponse,
	*pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryGetResponse,
	*pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest,
	*pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse,
) {
	pkgRepoSpec := pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec{
		ImgpkgBundle: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec{
			Image: "extensions.aws-usw2.tmc-dev.cloud.vmware.com/packages/standard/repo:v2.2.0_update.2",
		},
	}
	postRequestModel := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
		FullName: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName{
			Name:                  testConfig.PkgRepoName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         globalRepoNamespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &pkgRepoSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "pkgrepository1",
			ResourceVersion: "v1",
		},
	}

	postResponseModel := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
		FullName: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName{
			Name:                  testConfig.PkgRepoName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         globalRepoNamespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &pkgRepoSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "pkgrepository1",
			ResourceVersion: "v1",
		},
		Status: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
			Disabled:   false,
			Managed:    false,
			Subscribed: false,
		},
	}

	postRequest := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest{
		Repository: postRequestModel,
	}

	postResponse := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryResponse{
		Repository: postResponseModel,
	}

	getModel := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
		FullName: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName{
			Name:                  testConfig.PkgRepoName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         globalRepoNamespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &pkgRepoSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: referenceArray,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "pkgrepository1",
			ResourceVersion: "v1",
		},
		Status: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
			Disabled:   false,
			Managed:    false,
			Subscribed: false,
		},
	}

	getResponse := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryGetResponse{
		Repository: getModel,
	}

	availabilityRequest := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest{
		Disabled: true,
		FullName: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName{
			Name:                  testConfig.PkgRepoName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         globalRepoNamespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
	}

	availabilityResponse := &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse{
		Repository: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository{
			FullName: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName{
				Name:                  testConfig.PkgRepoName,
				OrgID:                 orgID,
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				NamespaceName:         globalRepoNamespace,
				ProvisionerName:       "attached",
				ManagementClusterName: "attached",
			},
			Spec: &pkgRepoSpec,
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				ParentReferences: referenceArray,
				Description:      "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				UID:             "pkgrepository1",
				ResourceVersion: "v1",
			},
			Status: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus{
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					"Ready": {
						Reason: "made successfully",
					},
				},
				Disabled:   true,
				Managed:    false,
				Subscribed: false,
			},
		},
	}

	return postRequest, postResponse, getResponse, availabilityRequest, availabilityResponse
}
