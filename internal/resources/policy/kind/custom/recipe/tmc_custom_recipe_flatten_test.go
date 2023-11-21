/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

const (
	customRecipeTemplateName = "some-custom-template"
)

func TestFlattenTMCCustom(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustom
		expected    []interface{}
	}{
		{
			description: "check for nil custom policy recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete custom policy recipe",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustom{
				Audit: true,
				Parameters: map[string]interface{}{
					"ranges": []map[string]interface{}{
						{
							"min_replicas": 3,
							"max_replicas": 7,
						},
					},
				},
				TargetKubernetesResources: []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
					{
						APIGroups: []string{"apps"},
						Kinds:     []string{"Deployment", "StatefulSet"},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					AuditKey:        true,
					TemplateNameKey: customRecipeTemplateName,
					ParametersKey:   "{\"ranges\":[{\"max_replicas\":7,\"min_replicas\":3}]}",
					TargetKubernetesResourcesKey: []interface{}{
						map[string]interface{}{
							APIGroupsKey: []string{"apps"},
							KindsKey:     []string{"Deployment", "StatefulSet"},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenTMCCustom(customRecipeTemplateName, test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestValidateRecipeParameters(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description      string
		recipeSchema     string
		recipeParameters string
		expectedErrors   bool
	}{
		{
			description:      "scenario for valid recipe parameters",
			recipeSchema:     "{\"description\":\"The input schema for replica-count-range-enforcement recipe\",\"type\":\"object\",\"title\":\"replica-count-range-enforcement recipe schema\",\"required\":[\"targetKubernetesResources\"],\"properties\":{\"audit\":{\"description\":\"Creates this policy for dry-run. Violations will be logged but not denied. Defaults to false (deny). (This is deprecated, please use enforcementAction instead)\",\"type\":\"boolean\",\"title\":\"Audit (dry-run)\"},\"enforcementAction\":{\"description\":\"Select the action to take when the policy is violated.\",\"type\":\"string\",\"title\":\"Enforcement Action\",\"pattern\":\"dryrun|warn|deny\"},\"parameters\":{\"type\":\"object\",\"properties\":{\"ranges\":{\"description\":\"Allowed ranges for numbers of replicas.  Values are inclusive.\",\"type\":\"array\",\"items\":{\"description\":\"A range of allowed replicas.  Values are inclusive.\",\"type\":\"object\",\"properties\":{\"max_replicas\":{\"description\":\"The maximum number of replicas allowed, inclusive.\",\"type\":\"integer\"},\"min_replicas\":{\"description\":\"The minimum number of replicas allowed, inclusive.\",\"type\":\"integer\"}}}}}},\"targetKubernetesResources\":{\"description\":\"TargetKubernetesResources is a list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. You can use 'kubectl api-resources' to view the list of available api resources on your cluster.\",\"type\":\"array\",\"minItems\":1,\"items\":{\"required\":[\"apiGroups\",\"kinds\"],\"properties\":{\"apiGroups\":{\"description\":\"apiGroup is group containing the resource type, for example 'rbac.authorization.k8s.io', 'networking.k8s.io', 'extensions', '' (some resources like Namespace, Pod have empty apiGroup).\",\"type\":\"array\",\"minItems\":1,\"items\":{\"type\":\"string\"}},\"kinds\":{\"description\":\"kind is the name of the object schema (resource type), for example 'Namespace', 'Pod', 'Ingress'\",\"type\":\"array\",\"minItems\":1,\"items\":{\"type\":\"string\"}}}}}}}",
			recipeParameters: "{\"ranges\":[{\"max_replicas\":7,\"min_replicas\":3}]}",
			expectedErrors:   false,
		},
		{
			description:      "scenario for invalid recipe parameters",
			recipeSchema:     "{\"description\":\"The input schema for replica-count-range-enforcement recipe\",\"type\":\"object\",\"title\":\"replica-count-range-enforcement recipe schema\",\"required\":[\"targetKubernetesResources\"],\"properties\":{\"audit\":{\"description\":\"Creates this policy for dry-run. Violations will be logged but not denied. Defaults to false (deny). (This is deprecated, please use enforcementAction instead)\",\"type\":\"boolean\",\"title\":\"Audit (dry-run)\"},\"enforcementAction\":{\"description\":\"Select the action to take when the policy is violated.\",\"type\":\"string\",\"title\":\"Enforcement Action\",\"pattern\":\"dryrun|warn|deny\"},\"parameters\":{\"type\":\"object\",\"properties\":{\"ranges\":{\"description\":\"Allowed ranges for numbers of replicas.  Values are inclusive.\",\"type\":\"array\",\"items\":{\"description\":\"A range of allowed replicas.  Values are inclusive.\",\"type\":\"object\",\"properties\":{\"max_replicas\":{\"description\":\"The maximum number of replicas allowed, inclusive.\",\"type\":\"integer\"},\"min_replicas\":{\"description\":\"The minimum number of replicas allowed, inclusive.\",\"type\":\"integer\"}}}}}},\"targetKubernetesResources\":{\"description\":\"TargetKubernetesResources is a list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. You can use 'kubectl api-resources' to view the list of available api resources on your cluster.\",\"type\":\"array\",\"minItems\":1,\"items\":{\"required\":[\"apiGroups\",\"kinds\"],\"properties\":{\"apiGroups\":{\"description\":\"apiGroup is group containing the resource type, for example 'rbac.authorization.k8s.io', 'networking.k8s.io', 'extensions', '' (some resources like Namespace, Pod have empty apiGroup).\",\"type\":\"array\",\"minItems\":1,\"items\":{\"type\":\"string\"}},\"kinds\":{\"description\":\"kind is the name of the object schema (resource type), for example 'Namespace', 'Pod', 'Ingress'\",\"type\":\"array\",\"minItems\":1,\"items\":{\"type\":\"string\"}}}}}}}",
			recipeParameters: "{\"replica_ranges\":[{\"maximum\":7,\"minimum\":3}]}",
			expectedErrors:   true,
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := ValidateRecipeParameters(test.recipeSchema, test.recipeParameters)
			require.Equal(t, test.expectedErrors, len(actual) > 0)
		})
	}
}
