// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package mutationpolicyresource

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
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	policyrecipemutationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation"
	policyrecipemutationcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation/common"
)

const (
	https         = "https:/"
	apiVersion    = "v1alpha1"
	clustergroups = "clustergroups"
	org           = "organization"
	apiKind       = "policies"
)

func (testConfig *testAcceptanceConfig) setupHTTPMocks(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)

	endpoint := os.Getenv("TMC_ENDPOINT")

	for _, recipe := range []string{annotation, label, podSecurity} {
		testConfig.setUpOrgPolicyEndPointMocks(t, recipe, endpoint)
		testConfig.setUpClusterGroupEndPointMocks(t, endpoint)
		testConfig.setUpClusterGroupPolicyEndpointMocks(t, recipe, endpoint)
	}
}

func (testConfig *testAcceptanceConfig) setUpClusterGroupPolicyEndpointMocks(t *testing.T, recipe string, endpoint string) {
	clusterGroupPolicyModel := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy{
		FullName: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			Name:             testConfig.MutationPolicyName + recipe,
			OrgID:            testConfig.ScopeHelperResources.OrgID,
		},
		Spec: getMockPolicyCreateSpec(recipe),
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			Description: "description",
			Labels: map[string]string{
				"key1": "value11",
				"key2": "value22",
			},
			UID:             fmt.Sprintf("%s-org-nw-policy", recipe),
			ResourceVersion: "v1",
		},
	}

	postClusterGroupPolicyRequest := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest{Policy: clusterGroupPolicyModel}
	postClusterGroupPolicyResponse := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse{Policy: clusterGroupPolicyModel}

	postClusterGroupPolicyEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, clustergroups, testConfig.ScopeHelperResources.ClusterGroup.Name, apiKind).String()
	getClusterGroupPolicyEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, clustergroups, testConfig.ScopeHelperResources.ClusterGroup.Name, apiKind, testConfig.MutationPolicyName+recipe).String()
	deleteClusterGroupPolicyEndpoint := getClusterGroupPolicyEndpoint

	httpmock.RegisterResponder("POST", postClusterGroupPolicyEndpoint,
		bodyInspectingResponder(t, postClusterGroupPolicyRequest, http.StatusOK, postClusterGroupPolicyResponse))

	httpmock.RegisterResponder("GET", getClusterGroupPolicyEndpoint,
		bodyInspectingResponder(t, nil, http.StatusOK, postClusterGroupPolicyResponse))

	httpmock.RegisterResponder("DELETE", deleteClusterGroupPolicyEndpoint, changeStateResponder(
		func() {
			httpmock.RegisterResponder("GET", getClusterGroupPolicyEndpoint,
				httpmock.NewStringResponder(http.StatusBadRequest, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) setUpClusterGroupEndPointMocks(t *testing.T, endpoint string) {
	clusterGroupResponse := &clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupResponse{
		ClusterGroup: &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupClusterGroup{
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				UID:         "12345",
				Description: "resource with description",
				Labels: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
	}

	postClusterGroupEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, clustergroups).String()
	getClusterGroupEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, clustergroups, testConfig.ScopeHelperResources.ClusterGroup.Name).String()
	deleteClusterGroupEndpoint := getClusterGroupEndpoint

	httpmock.RegisterResponder("POST", postClusterGroupEndpoint,
		bodyInspectingResponder(t, nil, http.StatusOK, clusterGroupResponse))

	httpmock.RegisterResponder("GET", getClusterGroupEndpoint,
		bodyInspectingResponder(t, nil, http.StatusOK, clusterGroupResponse))

	httpmock.RegisterResponder("DELETE", deleteClusterGroupEndpoint, changeStateResponder(
		func() {
			httpmock.RegisterResponder("GET", getClusterGroupEndpoint,
				httpmock.NewStringResponder(http.StatusBadRequest, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) setUpOrgPolicyEndPointMocks(t *testing.T, recipe string, endpoint string) {
	generatedOrgPolicyModel := testConfig.generateOrgPolicy(recipe)
	postOrgPolicyRequest := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{Policy: generatedOrgPolicyModel}
	postOrgPolicyResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse{Policy: generatedOrgPolicyModel}
	getOrgPolicyResponse := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse{Policy: generatedOrgPolicyModel}

	queryParams := url.Values{}
	queryParams.Add("fullName.orgId", testConfig.ScopeHelperResources.OrgID)

	postOrgPolicyEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, org, apiKind).String()
	getOrgPolicyEndpoint := helper.ConstructRequestURL(https, endpoint, apiVersion, org, apiKind, testConfig.MutationPolicyName+recipe).AppendQueryParams(queryParams).String()
	deleteOrgPolicyEndpoint := getOrgPolicyEndpoint

	httpmock.RegisterResponder("POST", postOrgPolicyEndpoint,
		bodyInspectingResponder(t, postOrgPolicyRequest, http.StatusOK, postOrgPolicyResponse))

	httpmock.RegisterResponder("GET", getOrgPolicyEndpoint,
		bodyInspectingResponder(t, nil, http.StatusOK, getOrgPolicyResponse))

	httpmock.RegisterResponder("DELETE", deleteOrgPolicyEndpoint, changeStateResponder(
		func() {
			httpmock.RegisterResponder("GET", getOrgPolicyEndpoint,
				httpmock.NewStringResponder(http.StatusBadRequest, "Not found"))
		},
		http.StatusOK,
		nil))
}

func (testConfig *testAcceptanceConfig) generateOrgPolicy(recipe string) *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy {
	return &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
		FullName: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
			Name:  testConfig.MutationPolicyName + recipe,
			OrgID: testConfig.ScopeHelperResources.OrgID,
		},
		Spec: getMockPolicyCreateSpec(recipe),
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			Description: "description",
			Labels: map[string]string{
				"key1": "value11",
				"key2": "value22",
			},
			UID:             fmt.Sprintf("%s-org-nw-policy", recipe),
			ResourceVersion: "v1",
		},
	}
}

func getMockPolicyCreateSpec(recipe string) *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec {
	spec := &policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec{
		Recipe:        recipe,
		RecipeVersion: "v1",
		Type:          "mutation-policy",
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
	spec.Input = struct{}{}

	switch spec.Recipe {
	case annotation:
		spec.Input = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation{
			Annotation: &policyrecipemutationcommonmodel.KeyValue{
				Key:   "test",
				Value: "optional"},
			Scope: policyrecipemutationcommonmodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope(policyrecipemutationcommonmodel.Cluster),
			TargetKubernetesResources: []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
				{
					APIGroups: []string{"apps"},
					Kinds:     []string{"Event"},
				},
			},
		}
	case label:
		spec.Input = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label{
			Label: &policyrecipemutationcommonmodel.KeyValue{
				Key:   "test",
				Value: "optional"},
			Scope: policyrecipemutationcommonmodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope(policyrecipemutationcommonmodel.Cluster),
			TargetKubernetesResources: []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
				{
					APIGroups: []string{"apps"},
					Kinds:     []string{"Event"},
				},
			},
		}
	case podSecurity:
		spec.Input = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity{
			AllowPrivilegeEscalation: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation{
				Condition: helper.StringPointer("Always"),
				Value:     helper.BoolPointer(true),
			},
			CapabilitiesAdd: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd{
				Operation: helper.StringPointer("merge"),
				Values:    []string{"AUDIT_CONTROL", "AUDIT_WRITE"},
			},
			CapabilitiesDrop: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop{
				Operation: helper.StringPointer("merge"),
				Values:    []string{"AUDIT_WRITE"},
			},
			FsGroup: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup{
				Condition: helper.StringPointer("Always"),
				Value:     helper.Float64Pointer(4),
			},
			Privileged: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged{
				Condition: helper.StringPointer("Always"),
				Value:     helper.BoolPointer(true),
			},
			ReadOnlyRootFilesystem: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem{
				Condition: helper.StringPointer("Always"),
				Value:     helper.BoolPointer(true),
			},
			RunAsGroup: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup{
				Condition: helper.StringPointer("Always"),
				Value:     helper.Float64Pointer(5),
			},
			RunAsNonRoot: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot{
				Condition: helper.StringPointer("Always"),
				Value:     helper.BoolPointer(true),
			},
			RunAsUser: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser{
				Condition: helper.StringPointer("Always"),
				Value:     helper.Float64Pointer(7),
			},
			SeLinuxOptions: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions{
				Condition: helper.StringPointer("IfFieldDoesNotExist"),
				Value: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue{
					Level: "level_test",
					Role:  "role_test",
					Type:  "type_test",
					User:  "user_test",
				},
			},
			SupplementalGroups: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups{
				Condition: helper.StringPointer("Always"),
				Values:    []*float64{helper.Float64Pointer(0), helper.Float64Pointer(1), helper.Float64Pointer(2), helper.Float64Pointer(3)},
			},
		}
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

func changeStateResponder(registerFunc func(), successResponse int, successResponseBody interface{}) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		registerFunc()
		return httpmock.NewJsonResponse(successResponse, successResponseBody)
	}
}
