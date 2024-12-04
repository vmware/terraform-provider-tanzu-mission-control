// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipemutationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation"
)

func TestFlattenPodSecurity(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity
		expected    []interface{}
	}{
		{
			description: "check for nil mutation label",
		},
		{
			description: "flatten pod security mutation policy struct",
			input: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity{
				AllowPrivilegeEscalation: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation{
					Condition: helper.StringPointer("IfFieldExists"),
					Value:     helper.BoolPointer(true),
				},
				CapabilitiesAdd: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd{
					Operation: helper.StringPointer("prune"),
					Values:    []string{"AUDIT_CONTROL", "AUDIT_READ"},
				},
				CapabilitiesDrop: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop{
					Operation: helper.StringPointer("prune"),
					Values:    []string{"SETFCAP"},
				},
				FsGroup: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup{
					Condition: helper.StringPointer("Always"),
					Value:     helper.Float64Pointer(4),
				},
				Privileged: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged{
					Condition: helper.StringPointer("IfFieldExists"),
					Value:     helper.BoolPointer(false),
				},
				ReadOnlyRootFilesystem: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem{
					Condition: helper.StringPointer("IfFieldDoesNotExist"),
					Value:     helper.BoolPointer(false),
				},
				RunAsGroup: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup{
					Condition: helper.StringPointer("IfFieldExists"),
					Value:     helper.Float64Pointer(567),
				},
				RunAsNonRoot: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot{
					Condition: helper.StringPointer("Always"),
					Value:     helper.BoolPointer(true),
				},
				RunAsUser: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser{
					Condition: helper.StringPointer("Always"),
					Value:     helper.Float64Pointer(2444),
				},
				SeLinuxOptions: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions{
					Condition: helper.StringPointer("IfFieldDoesNotExist"),
					Value: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue{
						Level: "level_up",
						Role:  "role_role",
						Type:  "type_type",
						User:  "user_user",
					},
				},
				SupplementalGroups: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups{
					Condition: helper.StringPointer("IfFieldDoesNotExist"),
					Values:    []*float64{helper.Float64Pointer(4546), helper.Float64Pointer(4)},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					allowPrivilegeEscalationKey: []interface{}{
						map[string]interface{}{
							conditionKey: "IfFieldExists",
							valueKey:     true,
						},
					},
					capabilitiesAddKey: []interface{}{
						map[string]interface{}{
							operationKey: "prune",
							valuesKey:    []interface{}{"AUDIT_CONTROL", "AUDIT_READ"},
						},
					},
					capabilitiesDropKey: []interface{}{
						map[string]interface{}{
							operationKey: "prune",
							valuesKey:    []interface{}{"SETFCAP"},
						},
					},
					fsGroupKey: []interface{}{
						map[string]interface{}{
							conditionKey: "Always",
							valueKey:     float64(4),
						},
					},
					privilegedKey: []interface{}{
						map[string]interface{}{
							conditionKey: "IfFieldExists",
							valueKey:     false,
						},
					},
					readOnlyRootFilesystemKey: []interface{}{
						map[string]interface{}{
							conditionKey: "IfFieldDoesNotExist",
							valueKey:     false,
						},
					},
					runAsGroupKey: []interface{}{
						map[string]interface{}{
							conditionKey: "IfFieldExists",
							valueKey:     float64(567),
						},
					},
					runAsNonRootKey: []interface{}{
						map[string]interface{}{
							conditionKey: "Always",
							valueKey:     true,
						},
					},
					runAsUserKey: []interface{}{
						map[string]interface{}{
							conditionKey: "Always",
							valueKey:     float64(2444),
						},
					},
					seLinuxOptionsKey: []interface{}{
						map[string]interface{}{
							conditionKey: "IfFieldDoesNotExist",
							levelKey:     "level_up",
							roleKey:      "role_role",
							typeKey:      "type_type",
							userKey:      "user_user",
						},
					},
					supplementalGroupsKey: []interface{}{
						map[string]interface{}{
							conditionKey: "IfFieldDoesNotExist",
							valuesKey:    []interface{}{float64(4546), float64(4)},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenPodSecurity(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
