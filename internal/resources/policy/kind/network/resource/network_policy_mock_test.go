// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package networkpolicyresource

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/go-test/deep"
	"github.com/jarcoal/httpmock"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
	policyworkspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
)

const (
	https      = "https:/"
	apiVersion = "v1alpha1"
	workspaces = "workspaces"
	org        = "organization"
	apiKind    = "policies"
)

// Function to set up HTTP mocks for the network policy requests anticipated by this test, when not being run against a real TMC stack.
func (testConfig *testAcceptanceConfig) setupHTTPMocks(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)

	endpoint := os.Getenv("TMC_ENDPOINT")

	reference := objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference{
		Rid: "test_rid",
		UID: "test_uid",
	}
	referenceArray := make([]*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference, 0)
	referenceArray = append(referenceArray, &reference)

	for _, recipe := range []string{"allow-all", "allow-all-to-pods", "allow-all-egress", "deny-all", "deny-all-to-pods", "deny-all-egress", "custom-egress", "custom-ingress"} {
		postWorkspacePolicyModel := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicy{
			FullName: &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName{
				Name:          testConfig.NetworkPolicyName + recipe,
				OrgID:         testConfig.ScopeHelperResources.OrgID,
				WorkspaceName: "workspace1",
			},
			Spec: getMockPolicyCreateSpec(recipe),
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				ParentReferences: nil,
				Description:      "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				UID:             fmt.Sprintf("%s-ws-nw-policy", recipe),
				ResourceVersion: "v1",
			},
		}
		getWorkspacePolicyModel := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicy{
			FullName: &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName{
				Name:          testConfig.NetworkPolicyName + recipe,
				OrgID:         testConfig.ScopeHelperResources.OrgID,
				WorkspaceName: "workspace1",
			},
			Spec: getMockPolicyCreateSpec(recipe),
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				ParentReferences: referenceArray,
				Description:      "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				UID:             fmt.Sprintf("%s-ws-nw-policy", recipe),
				ResourceVersion: "v1",
			},
		}

		postWorkspacePolicyRequest := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest{Policy: postWorkspacePolicyModel}
		postWorkspacePolicyResponse := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse{Policy: postWorkspacePolicyModel}
		getWorkspacePolicyResponse := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse{Policy: getWorkspacePolicyModel}

		postWorkspacePolicyEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, workspaces, "workspace1", apiKind).String()
		getWorkspacePolicyEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, workspaces, "workspace1", apiKind, testConfig.NetworkPolicyName+recipe).String()
		deleteWorkspacePolicyEndpoint := getWorkspacePolicyEndpoint

		postOrgPolicyModel := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
			FullName: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
				Name:  testConfig.NetworkPolicyName + recipe,
				OrgID: testConfig.ScopeHelperResources.OrgID,
			},
			Spec: getMockPolicyCreateSpec(recipe),
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				ParentReferences: nil,
				Description:      "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				UID:             fmt.Sprintf("%s-org-nw-policy", recipe),
				ResourceVersion: "v1",
			},
		}
		getOrgPolicyModel := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
			FullName: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
				Name:  testConfig.NetworkPolicyName + recipe,
				OrgID: testConfig.ScopeHelperResources.OrgID,
			},
			Spec: getMockPolicyCreateSpec(recipe),
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				ParentReferences: referenceArray,
				Description:      "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				UID:             fmt.Sprintf("%s-org-nw-policy", recipe),
				ResourceVersion: "v1",
			},
		}

		postOrgPolicyRequest := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{Policy: postOrgPolicyModel}
		postOrgPolicyResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse{Policy: postOrgPolicyModel}
		getOrgPolicyResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse{Policy: getOrgPolicyModel}

		queryParams := url.Values{}
		queryParams.Add("fullName.orgId", testConfig.ScopeHelperResources.OrgID)

		postOrgPolicyEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, org, apiKind).String()
		getOrgPolicyEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, org, apiKind, testConfig.NetworkPolicyName+recipe).AppendQueryParams(queryParams).String()

		deleteOrgPolicyEndpoint := getOrgPolicyEndpoint

		httpmock.RegisterResponder("POST", postWorkspacePolicyEndpoint,
			bodyInspectingResponder(t, postWorkspacePolicyRequest, http.StatusOK, postWorkspacePolicyResponse))

		httpmock.RegisterResponder("GET", getWorkspacePolicyEndpoint,
			bodyInspectingResponder(t, nil, http.StatusOK, getWorkspacePolicyResponse))

		httpmock.RegisterResponder("DELETE", deleteWorkspacePolicyEndpoint, changeStateResponder(
			// Set up the get to return 404 after the policy has been 'deleted'
			func() {
				httpmock.RegisterResponder("GET", getWorkspacePolicyEndpoint,
					httpmock.NewStringResponder(http.StatusBadRequest, "Not found"))
			},
			http.StatusOK,
			nil))

		httpmock.RegisterResponder("POST", postOrgPolicyEndpoint,
			bodyInspectingResponder(t, postOrgPolicyRequest, http.StatusOK, postOrgPolicyResponse))

		httpmock.RegisterResponder("GET", getOrgPolicyEndpoint,
			bodyInspectingResponder(t, nil, http.StatusOK, getOrgPolicyResponse))

		httpmock.RegisterResponder("DELETE", deleteOrgPolicyEndpoint, changeStateResponder(
			// Set up the get to return 404 after the policy has been 'deleted'
			func() {
				httpmock.RegisterResponder("GET", getOrgPolicyEndpoint,
					httpmock.NewStringResponder(http.StatusBadRequest, "Not found"))
			},
			http.StatusOK,
			nil))
	}
}

func getMockPolicyCreateSpec(recipe string) *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec {
	spec := &policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec{
		Recipe:        recipe,
		RecipeVersion: "v1",
		Type:          "network-policy",
		NamespaceSelector: &policymodel.VmwareTanzuManageV1alpha1CommonPolicyLabelSelector{
			MatchExpressions: []*policymodel.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement{
				{
					Key:      "component",
					Operator: "NotIn",
					Values:   []string{"api-server", "agent-gateway"},
				},
				{
					Key:      "not-a-component",
					Operator: "DoesNotExist",
					Values:   []string{},
				},
			},
		},
	}

	switch spec.Recipe {
	case "allow-all":
		spec.Input = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll{
			FromOwnNamespace: helper.BoolPointer(false),
		}
	case "allow-all-to-pods":
		spec.Input = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAllToPods{
			FromOwnNamespace: helper.BoolPointer(false),
			ToPodLabels: []*policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
				{
					Key:   "key1",
					Value: "value1",
				},
				{
					Key:   "key2",
					Value: "value2",
				},
			},
		}
	case "deny-all-to-pods":
		spec.Input = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1DenyAllToPods{
			ToPodLabels: []*policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
				{
					Key:   "key1",
					Value: "value1",
				},
				{
					Key:   "key2",
					Value: "value2",
				},
			},
		}
	// TODO: Add Case for Custom Egress and Custom Ingress mock run
	// case "custom-egress":
	//	spec.Input = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomEgress{
	//		Rules: []policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRules{
	//			{
	//				Ports: func() *[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPorts {
	//					return &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPorts{
	//						{
	//							Port: func() *string {
	//								port := "8443"
	//								return &port
	//							}(),
	//							Protocol: policyrecipenetworkcommonmodel.NewV1alpha1CommonPolicySpecNetworkV1CustomRulesPortsProtocol(policyrecipenetworkcommonmodel.TCP),
	//						},
	//					}
	//				}(),
	//				RuleSpec: func() []interface{} {
	//					ruleSpec := &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec1{
	//						{
	//							NamespaceSelector: &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
	//								{
	//									Key:   "key-1",
	//									Value: "value-1",
	//								},
	//								{
	//									Key:   "key-2",
	//									Value: "value-2",
	//								},
	//							},
	//							PodSelector: &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
	//								{
	//									Key:   "key-1",
	//									Value: "value-1",
	//								},
	//								{
	//									Key:   "key-2",
	//									Value: "value-2",
	//								},
	//							},
	//						},
	//					}
	//
	//					return []interface{}{ruleSpec}
	//				}(),
	//			},
	//		},
	//		ToPodLabels: func() *[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels {
	//			return &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
	//				{
	//					Key:   "key1",
	//					Value: "value1",
	//				},
	//				{
	//					Key:   "key2",
	//					Value: "value2",
	//				},
	//			}
	//		}(),
	//	}
	// case "custom-ingress":
	//	spec.Input = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomIngress{
	//		Rules: []policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRules{
	//			{
	//				Ports: func() *[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPorts {
	//					return &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPorts{
	//						{
	//							Port: func() *string {
	//								port := "8443"
	//								return &port
	//							}(),
	//							Protocol: policyrecipenetworkcommonmodel.NewV1alpha1CommonPolicySpecNetworkV1CustomRulesPortsProtocol(policyrecipenetworkcommonmodel.TCP),
	//						},
	//					}
	//				}(),
	//				RuleSpec: func() []interface{} {
	//					ruleSpec := &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0{
	//						{
	//							IpBlock: &policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0IpBlock{
	//								Cidr:   "192.168.1.1/24",
	//								Except: &[]string{"2001:db9::/64"},
	//							},
	//						},
	//					}
	//
	//					return []interface{}{ruleSpec}
	//				}(),
	//			},
	//		},
	//		ToPodLabels: func() *[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels {
	//			return &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
	//				{
	//					Key:   "key1",
	//					Value: "value1",
	//				},
	//				{
	//					Key:   "key2",
	//					Value: "value2",
	//				},
	//			}
	//		}(),
	//	}
	case "allow-all-egress", "deny-all", "deny-all-egress":
		spec.Input = struct{}{}
	}

	return spec
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
