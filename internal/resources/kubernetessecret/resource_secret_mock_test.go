// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kubernetessecret

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
	secretmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	secretcgmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

const (
	https                = "https:/"
	apiVersionAndGroup   = "v1alpha1/clusters"
	apiSubGroup          = "namespaces"
	apiKind              = "secrets"
	exportAPIKind        = "secretexports"
	cgAPIVersionAndGroup = "v1alpha1/clustergroups"
)

func getMockSpec(secretType string) secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec {
	switch secretType {
	case DockerSecretType:
		return secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{
			SecretType: secretmodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEDOCKERCONFIGJSON),
			Data: map[string]strfmt.Base64{
				".dockerconfigjson": []byte(`{"auths":{"someregistryurl":{"auth":"","password":"","username":"someusername"}}}`),
			},
		}
	case OpaqueSecretType:
		return secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{
			SecretType: secretmodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEOPAQUE),
			Data: map[string]strfmt.Base64{
				"username": []byte(`someusername`),
				"password": []byte(`somepassword`),
			},
		}
	default:
		return secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{}
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

// nolint: unparam
// Register a new responder when the given call is made.
func changeStateResponder(registerFunc func(), successResponse int, successResponseBody interface{}) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		registerFunc()
		return httpmock.NewJsonResponse(successResponse, successResponseBody)
	}
}

// Function to set up HTTP mocks for the kubernetes secret requests anticipated by this test, when not being run against a real TMC stack.
func (testConfig *testAcceptanceConfig) setupHTTPMocks(t *testing.T, secretType string) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)

	endpoint := os.Getenv("TMC_ENDPOINT")

	OrgID := os.Getenv("ORG_ID")

	// cluster level cluster secret resource.
	secretSpec := getMockSpec(secretType)
	postRequestModel := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecret{
		FullName: &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{
			Name:                  testConfig.SecretName,
			OrgID:                 OrgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
			NamespaceName:         testConfig.NamespaceName,
		},
		Spec: &secretSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "secret1",
			ResourceVersion: "v1",
		},
	}

	postResponseModel := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecret{
		FullName: &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{
			Name:                  testConfig.SecretName,
			OrgID:                 OrgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
			NamespaceName:         testConfig.NamespaceName,
		},
		Spec: &secretSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "secret1",
			ResourceVersion: "v1",
		},
	}

	postRequest := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest{
		Secret: postRequestModel,
	}

	postResponse := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse{
		Secret: postResponseModel,
	}

	// GET Network Secret mock setup
	getModel := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecret{
		FullName: &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{
			Name:                  testConfig.SecretName,
			OrgID:                 OrgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
			NamespaceName:         testConfig.NamespaceName,
		},
		Spec: &secretSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "secret1",
			ResourceVersion: "v1",
		},
		Status: &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretStatus{
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"Ready": {
					Reason: "made successfully",
				},
			},
		},
	}

	getResponse := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse{
		Secret: getModel,
	}

	postEndpoint := (helper.ConstructRequestURL(https, endpoint, apiVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.NamespaceName, apiKind)).String()
	getSecretEndpoint := (helper.ConstructRequestURL(https, endpoint, apiVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.NamespaceName, apiKind, testConfig.SecretName)).String()
	deleteEndpoint := getSecretEndpoint

	getSecretExportEndpoint := (helper.ConstructRequestURL(https, endpoint, apiVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.NamespaceName, exportAPIKind, testConfig.SecretName)).String()
	deleteExportEndpoint := getSecretExportEndpoint

	httpmock.RegisterResponder("POST", postEndpoint,
		bodyInspectingResponder(t, postRequest, 200, postResponse))

	httpmock.RegisterResponder("GET", getSecretEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))

	httpmock.RegisterResponder("DELETE", deleteEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'
		func() {
			httpmock.RegisterResponder("GET", getSecretEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))

	httpmock.RegisterResponder("GET", getSecretExportEndpoint,
		httpmock.NewStringResponder(404, "Not found"))

	httpmock.RegisterResponder("DELETE", deleteExportEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'
		func() {
			httpmock.RegisterResponder("GET", getSecretExportEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))

	// cluster group level cluster secret resource.

	postCGRequestModel := &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret{
		FullName: &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName{
			Name:             testConfig.SecretName,
			OrgID:            OrgID,
			NamespaceName:    testConfig.NamespaceName,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
		},
		Spec: &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec{
			AtomicSpec: &secretSpec,
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "secret1",
			ResourceVersion: "v1",
		},
	}

	postCGResponseModel := &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret{
		FullName: &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName{
			Name:             testConfig.SecretName,
			OrgID:            OrgID,
			NamespaceName:    testConfig.NamespaceName,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
		},
		Spec: &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec{
			AtomicSpec: &secretSpec,
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "secret1",
			ResourceVersion: "v1",
		},
	}

	postCGRequest := &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest{
		Secret: postCGRequestModel,
	}

	postCGResponse := &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse{
		Secret: postCGResponseModel,
	}

	// GET Network Secret mock setup
	getCGModel := &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret{
		FullName: &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName{
			Name:             testConfig.SecretName,
			OrgID:            OrgID,
			NamespaceName:    testConfig.NamespaceName,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
		},
		Spec: &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec{
			AtomicSpec: &secretSpec,
		},
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			UID:             "secret1",
			ResourceVersion: "v1",
		},
		Status: &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretStatus{
			Phase: statusmodel.NewVmwareTanzuManageV1alpha1CommonBatchPhase(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
		},
	}

	getCGResponse := &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretGetSecretResponse{
		Secret: getCGModel,
	}

	postCGEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind)).String()
	getSecretCGEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", apiKind, testConfig.SecretName)).String()
	deleteCGEndpoint := getSecretCGEndpoint

	getSecretExportCGEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, "namespace", exportAPIKind, testConfig.SecretName)).String()
	deleteExportCGEndpoint := getSecretExportCGEndpoint

	httpmock.RegisterResponder("POST", postCGEndpoint,
		bodyInspectingResponder(t, postCGRequest, 200, postCGResponse))

	httpmock.RegisterResponder("GET", getSecretCGEndpoint,
		bodyInspectingResponder(t, nil, 200, getCGResponse))

	httpmock.RegisterResponder("DELETE", deleteCGEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'
		func() {
			httpmock.RegisterResponder("GET", getSecretCGEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))

	httpmock.RegisterResponder("GET", getSecretExportCGEndpoint,
		httpmock.NewStringResponder(404, "Not found"))

	httpmock.RegisterResponder("DELETE", deleteExportCGEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'
		func() {
			httpmock.RegisterResponder("GET", getSecretExportCGEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))
}
