// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackageinstall

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/go-test/deep"
	"github.com/jarcoal/httpmock"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	pakageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	tanzupakageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackage"
	packageinstallmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
)

const (
	https                = "https:/"
	clAPIVersionAndGroup = "v1alpha1/clusters"
	apiSubGroup          = "namespaces"
	apiKind              = "tanzupackage/installs"
	apiKindMetadata      = "tanzupackage/metadatas"
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

	getModel := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall{
		FullName: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName{
			Name:                  testConfig.PkgInstallName,
			OrgID:                 OrgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec{
			PackageRef: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef{
				PackageMetadataName: "pkg.test.carvel.dev",
				VersionSelection: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection{
					Constraints: constraints,
				},
			},
			RoleBindingScope: packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER.Pointer(),
			InlineValues: map[string]interface{}{
				"bar": "foo",
			},
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: referenceArray,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "pkginstall1",
			ResourceVersion: "v1",
		},
		Status: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus{
			Managed: false,
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Type:     "Ready",
					Status:   statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE.Pointer(),
					Severity: statusmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO.Pointer(),
					Reason:   "ReconcileSucceeded",
					Message:  "Reconcile Succeeded",
				},
			},
			ResolvedVersion: "1.9.5+vmware.1-tkg.1",
			GeneratedResources: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources{
				ClusterRoleName:    "testclusterrole",
				RoleBindingName:    "testrolebinding",
				ServiceAccountName: "testserviceaccount",
			},
			ReferredBy: []string{"foo", "bar"},
		},
	}

	getResponse := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse{
		Install: getModel,
	}

	getPkgInstallEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind, testConfig.PkgInstallName)).String()

	httpmock.RegisterResponder("GET", getPkgInstallEndpoint,
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

	getTanzuPkgMetadataResponse := func(pkgName string) *pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse {
		return &pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse{
			Package: &pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackage{
				FullName: &pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageFullName{
					ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
					ManagementClusterName: "attached",
					ProvisionerName:       "attached",
					OrgID:                 OrgID,
					Name:                  pkgName,
					NamespaceName:         globalRepoNamespace,
					MetadataName:          pkgMetadataName,
				},
				Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
					ParentReferences: referenceArray,
					Description:      "resource with description",
					Labels: map[string]string{
						"key1": "value1",
						"key2": "value2",
					},
					UID:             "package1",
					ResourceVersion: "v1",
				},
				Spec: &pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec{
					CapacityRequirementsDescription: "someCapacityRequirementsDescription",
					Licenses: []string{
						"some1",
					},
					ReleaseNotes:   "cert-manager 1.1.0 https://github.com/jetstack/cert-manager/1.1.0",
					ReleasedAt:     strfmt.DateTime{},
					RepositoryName: "testRepo",
					ValuesSchema: &pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema{
						Template: &pakageclustermodel.K8sIoApimachineryPkgRuntimeRawExtension{
							Raw: []byte("somevalue"),
						},
					},
				},
			},
		}
	}

	getTanzuPkgMetadataEndpoint1 := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, globalRepoNamespace, apiKindMetadata, pkgMetadataName, "packages", testConfig.PkgName1)).String()
	getTanzuPkgMetadataEndpoint2 := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, globalRepoNamespace, apiKindMetadata, pkgMetadataName, "packages", testConfig.PkgName2)).String()

	httpmock.RegisterResponder("GET", getTanzuPkgMetadataEndpoint1,
		bodyInspectingResponder(t, nil, 200, getTanzuPkgMetadataResponse(testConfig.PkgName1)))
	httpmock.RegisterResponder("GET", getTanzuPkgMetadataEndpoint2,
		bodyInspectingResponder(t, nil, 200, getTanzuPkgMetadataResponse(testConfig.PkgName2)))

	// cluster level package resource.

	getTanzuPackageResponse := &tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse{
		TanzuPackages: []*tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageTanzuPackage{
			{
				FullName: &tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageFullName{
					ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
					ManagementClusterName: "attached",
					ProvisionerName:       "attached",
					OrgID:                 OrgID,
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

	getTanzuPackageEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, "tanzupackage")).String()

	httpmock.RegisterResponder("GET", getTanzuPackageEndpoint,
		bodyInspectingResponder(t, nil, 200, getTanzuPackageResponse))

	// cluster level package install resource.
	postRequest, postResponse, getResponse := testConfig.getClRequestResponse(OrgID, referenceArray)

	putRequest := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest{
		Install: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall{
			FullName: postRequest.Install.FullName,
			Meta:     postRequest.Install.Meta,
			Spec: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec{
				PackageRef: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef{
					PackageMetadataName: "pkg.test.carvel.dev",
					VersionSelection: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection{
						Constraints: constraints,
					},
				},
				RoleBindingScope: packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER.Pointer(),
				InlineValues: map[string]interface{}{
					"bar": "foo",
				},
			},
		},
	}

	putResponse := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest{
		Install: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall{
			FullName: postRequest.Install.FullName,
			Meta:     postRequest.Install.Meta,
			Spec: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec{
				PackageRef: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef{
					PackageMetadataName: "pkg.test.carvel.dev",
					VersionSelection: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection{
						Constraints: constraints,
					},
				},
				RoleBindingScope: packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER.Pointer(),
				InlineValues: map[string]interface{}{
					"bar": "foo",
				},
			},
			Status: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus{
				Managed: false,
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					"Ready": {
						Type:     "Ready",
						Status:   statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE.Pointer(),
						Severity: statusmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO.Pointer(),
						Reason:   "ReconcileSucceeded",
						Message:  "Reconcile Succeeded",
					},
				},
				ResolvedVersion: "1.9.5+vmware.1-tkg.1",
				GeneratedResources: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources{
					ClusterRoleName:    "testclusterrole",
					RoleBindingName:    "testrolebinding",
					ServiceAccountName: "testserviceaccount",
				},
				ReferredBy: []string{"foo", "bar"},
			},
		},
	}

	postEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind)).String()
	getPkgInstallEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind, testConfig.PkgInstallName)).String()
	deleteEndpoint := getPkgInstallEndpoint

	httpmock.RegisterResponder("POST", postEndpoint,
		bodyInspectingResponder(t, postRequest, 200, postResponse))

	httpmock.RegisterResponder("PUT", getPkgInstallEndpoint,
		bodyInspectingResponder(t, putRequest, 200, putResponse))

	httpmock.RegisterResponder("GET", getPkgInstallEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))

	httpmock.RegisterResponder("DELETE", deleteEndpoint, changeStateResponder(
		// Set up the get to return 404 after the package install has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getPkgInstallEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) getClRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest,
	*packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse,
	*packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse,
) {
	pkgInstallSpec := packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec{
		PackageRef: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef{
			PackageMetadataName: "pkg.test.carvel.dev",
			VersionSelection: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection{
				Constraints: "2.0.0",
			},
		},
		RoleBindingScope: packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER.Pointer(),
		InlineValues: map[string]interface{}{
			"bar": "foo",
		},
	}
	postRequestModel := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall{
		FullName: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName{
			Name:                  testConfig.PkgInstallName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &pkgInstallSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "pkginstall1",
			ResourceVersion: "v1",
		},
	}

	postResponseModel := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall{
		FullName: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName{
			Name:                  testConfig.PkgInstallName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &pkgInstallSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "pkginstall1",
			ResourceVersion: "v1",
		},
		Status: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus{
			Managed: false,
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Type:     "Ready",
					Status:   statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE.Pointer(),
					Severity: statusmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO.Pointer(),
					Reason:   "ReconcileSucceeded",
					Message:  "Reconcile Succeeded",
				},
			},
			ResolvedVersion: "1.9.5+vmware.1-tkg.1",
			GeneratedResources: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources{
				ClusterRoleName:    "testclusterrole",
				RoleBindingName:    "testrolebinding",
				ServiceAccountName: "testserviceaccount",
			},
			ReferredBy: []string{"foo", "bar"},
		},
	}

	postRequest := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest{
		Install: postRequestModel,
	}

	postResponse := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse{
		Install: postResponseModel,
	}

	getModel := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall{
		FullName: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName{
			Name:                  testConfig.PkgInstallName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			NamespaceName:         testConfig.Namespace,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &pkgInstallSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: referenceArray,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "pkginstall1",
			ResourceVersion: "v1",
		},
		Status: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus{
			Managed: false,
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Type:     "Ready",
					Status:   statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE.Pointer(),
					Severity: statusmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO.Pointer(),
					Reason:   "ReconcileSucceeded",
					Message:  "Reconcile Succeeded",
				},
			},
			ResolvedVersion: "1.9.5+vmware.1-tkg.1",
			GeneratedResources: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources{
				ClusterRoleName:    "testclusterrole",
				RoleBindingName:    "testrolebinding",
				ServiceAccountName: "testserviceaccount",
			},
			ReferredBy: []string{"foo", "bar"},
		},
	}

	getResponse := &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse{
		Install: getModel,
	}

	return postRequest, postResponse, getResponse
}
