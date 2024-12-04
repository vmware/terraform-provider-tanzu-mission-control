// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package sourcesecret

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
	continuousdeliveryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/continuousdelivery/cluster"
	continuousdeliveryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/continuousdelivery/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	sourcesecretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
	sourcesecretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret/spec"
)

const (
	https                = "https:/"
	clAPIVersionAndGroup = "v1alpha1/clusters"
	apiKind              = "fluxcd/sourcesecrets"
	cdAPIKind            = "fluxcd/continuousdelivery"
	cgAPIVersionAndGroup = "v1alpha1/clustergroups"
)

func getMockSpec(allowedCredential string) sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec {
	switch allowedCredential {
	case spec.UsernamePasswordKey:
		return sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
			SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeUSERNAMEPASSWORD),
			Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
				Data: map[string]strfmt.Base64{
					"username": []byte("testusername"),
					"password": []byte(""),
				},
			},
		}
	case spec.SSHKey:
		return sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
			SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeSSH),
			Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
				Data: map[string]strfmt.Base64{
					"identity":    []byte(""),
					"known_hosts": []byte("testhostes"),
				},
			},
		}
	default:
		return sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{}
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
		getModel := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret{
			FullName: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName{
				Name:                  testConfig.SourceSecretName,
				OrgID:                 OrgID,
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ProvisionerName:       "attached",
				ManagementClusterName: "attached",
			},
			Spec: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
				SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeUSERNAMEPASSWORD),
				Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
					Data: map[string]strfmt.Base64{
						"username": []byte("someusername"),
						"password": []byte(""),
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
				UID:             "kustomization1",
				ResourceVersion: "v1",
			},
		}

		getResponse := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdGetSourceSecretResponse{
			SourceSecret: getModel,
		}
		getSourcesecretEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiKind, testConfig.SourceSecretName)).String()

		httpmock.RegisterResponder("GET", getSourcesecretEndpoint,
			bodyInspectingResponder(t, nil, 200, getResponse))
	case commonscope.ClusterGroupScope:
		getCGModel := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret{
			FullName: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName{
				Name:             testConfig.SourceSecretName,
				OrgID:            OrgID,
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			},
			Spec: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec{
				AtomicSpec: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
					SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeSSH),
					Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
						Data: map[string]strfmt.Base64{
							"identity":    []byte(""),
							"known_hosts": []byte("somehosts"),
						},
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
		}

		getCGResponse := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdGetSourceSecretResponse{
			SourceSecret: getCGModel,
		}

		getCGSourcesecretEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, apiKind, testConfig.SourceSecretName)).String()

		httpmock.RegisterResponder("GET", getCGSourcesecretEndpoint,
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

	// cluster level source secret resorce.
	postRequest, postResponse, getResponse, postContinuousDeliveryRequest, postContinuousDeliveryResponse := testConfig.getClRequestResponse(OrgID, referenceArray)

	putRequest := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest{
		SourceSecret: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret{
			FullName: postRequest.SourceSecret.FullName,
			Meta:     postRequest.SourceSecret.Meta,
			Spec: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
				SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeUSERNAMEPASSWORD),
				Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
					Data: map[string]strfmt.Base64{
						"username": []byte("someusername"),
						"password": []byte(""),
					},
				},
			},
		},
	}

	putResponse := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse{
		SourceSecret: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret{
			FullName: postRequest.SourceSecret.FullName,
			Meta:     postRequest.SourceSecret.Meta,
			Spec: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
				SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeUSERNAMEPASSWORD),
				Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
					Data: map[string]strfmt.Base64{
						"username": []byte("someusername"),
						"password": []byte(""),
					},
				},
			},
		},
	}

	postEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiKind)).String()
	getSourcesecretEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiKind, testConfig.SourceSecretName)).String()
	deleteEndpoint := getSourcesecretEndpoint

	postContinuousDeliveryEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, cdAPIKind)).String()

	httpmock.RegisterResponder("POST", postContinuousDeliveryEndpoint,
		bodyInspectingResponder(t, postContinuousDeliveryRequest, 200, postContinuousDeliveryResponse))

	httpmock.RegisterResponder("POST", postEndpoint,
		bodyInspectingResponder(t, postRequest, 200, postResponse))

	httpmock.RegisterResponder("PUT", getSourcesecretEndpoint,
		bodyInspectingResponder(t, putRequest, 200, putResponse))

	httpmock.RegisterResponder("GET", getSourcesecretEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))

	httpmock.RegisterResponder("DELETE", deleteEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getSourcesecretEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))

	// cluster group level source secret resorce.
	postCGRequest, postCGResponse, getCGResponse, postCGContinuousDeliveryRequest, postCGContinuousDeliveryResponse := testConfig.getCGRequestResponse(OrgID, referenceArray)

	putCGRequest := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest{
		SourceSecret: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret{
			FullName: postCGRequest.SourceSecret.FullName,
			Meta:     postCGRequest.SourceSecret.Meta,
			Spec: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec{
				AtomicSpec: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
					SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeSSH),
					Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
						Data: map[string]strfmt.Base64{
							"identity":    []byte(""),
							"known_hosts": []byte("somehosts"),
						},
					},
				},
			},
		},
	}

	putCGResponse := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse{
		SourceSecret: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret{
			FullName: postCGRequest.SourceSecret.FullName,
			Meta:     postCGRequest.SourceSecret.Meta,
			Spec: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec{
				AtomicSpec: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
					SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeSSH),
					Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
						Data: map[string]strfmt.Base64{
							"identity":    []byte(""),
							"known_hosts": []byte("somehosts"),
						},
					},
				},
			},
		},
	}

	postCDEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, apiKind)).String()
	getCGSourcesecretEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, apiKind, testConfig.SourceSecretName)).String()
	deleteCGEndpoint := getCGSourcesecretEndpoint

	postCDContinuousDeliveryEndpoint := (helper.ConstructRequestURL(https, endpoint, cgAPIVersionAndGroup, testConfig.ScopeHelperResources.ClusterGroup.Name, cdAPIKind)).String()

	httpmock.RegisterResponder("POST", postCDContinuousDeliveryEndpoint,
		bodyInspectingResponder(t, postCGContinuousDeliveryRequest, 200, postCGContinuousDeliveryResponse))

	httpmock.RegisterResponder("POST", postCDEndpoint,
		bodyInspectingResponder(t, postCGRequest, 200, postCGResponse))

	httpmock.RegisterResponder("GET", getCGSourcesecretEndpoint,
		bodyInspectingResponder(t, nil, 200, getCGResponse))

	httpmock.RegisterResponder("PUT", getCGSourcesecretEndpoint,
		bodyInspectingResponder(t, putCGRequest, 200, putCGResponse))

	httpmock.RegisterResponder("DELETE", deleteCGEndpoint, changeStateResponder(
		// Set up the get to return 404 after the Secret has been 'deleted'.
		func() {
			httpmock.RegisterResponder("GET", getCGSourcesecretEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) getCGRequestResponse(orgID string, referenceArray []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference) (
	*sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest,
	*sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse,
	*sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdGetSourceSecretResponse,
	*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest,
	*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse,
) {
	sourcesecretSpec := getMockSpec(spec.UsernamePasswordKey)

	cgSourcesecretSpec := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec{
		AtomicSpec: &sourcesecretSpec,
	}

	postCGRequestModel := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret{
		FullName: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName{
			Name:             testConfig.SourceSecretName,
			OrgID:            orgID,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
		},
		Spec: cgSourcesecretSpec,
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

	postCGResponseModel := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret{
		FullName: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName{
			Name:             testConfig.SourceSecretName,
			OrgID:            orgID,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
		},
		Spec: cgSourcesecretSpec,
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

	postCGRequest := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest{
		SourceSecret: postCGRequestModel,
	}

	postCGResponse := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse{
		SourceSecret: postCGResponseModel,
	}

	getCGModel := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret{
		FullName: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName{
			Name:             testConfig.SourceSecretName,
			OrgID:            orgID,
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
		},
		Spec: cgSourcesecretSpec,
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
	}

	getCGResponse := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdGetSourceSecretResponse{
		SourceSecret: getCGModel,
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
	*sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest,
	*sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse,
	*sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdGetSourceSecretResponse,
	*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest,
	*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse,
) {
	sourcesecretSpec := getMockSpec(spec.SSHKey)
	postRequestModel := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret{
		FullName: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName{
			Name:                  testConfig.SourceSecretName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &sourcesecretSpec,
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

	postResponseModel := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret{
		FullName: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName{
			Name:                  testConfig.SourceSecretName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &sourcesecretSpec,
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

	postRequest := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest{
		SourceSecret: postRequestModel,
	}

	postResponse := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse{
		SourceSecret: postResponseModel,
	}

	getModel := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret{
		FullName: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName{
			Name:                  testConfig.SourceSecretName,
			OrgID:                 orgID,
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			ProvisionerName:       "attached",
			ManagementClusterName: "attached",
		},
		Spec: &sourcesecretSpec,
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
	}

	getResponse := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdGetSourceSecretResponse{
		SourceSecret: getModel,
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
