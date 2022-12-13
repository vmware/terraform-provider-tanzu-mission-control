/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipeimagemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image"
	policyrecipeimagecommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image/common"
)

func TestFlattenCustom(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1Custom
		expected    []interface{}
	}{
		{
			description: "check for nil image policy custom recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "scenario with nil value of Audit value in custom recipe",
			input: &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1Custom{
				Audit: nil,
				Rules: []*policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CustomRules{
					{
						Hostname:      "bar",
						ImageName:     "foo",
						Port:          "80",
						RequireDigest: helper.BoolPointer(true),
						Tag: &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{
							Negate: helper.BoolPointer(false),
							Value:  "test",
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					RulesKey: []interface{}{
						map[string]interface{}{
							HostNameKey:  "bar",
							ImageNameKey: "foo",
							PortKey:      "80",
							RequireKey:   true,
							TagKey: []interface{}{
								map[string]interface{}{
									NegateKey: false,
									ValueKey:  "test",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "normal scenario with complete image policy custom recipe",
			input: &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1Custom{
				Audit: helper.BoolPointer(true),
				Rules: []*policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CustomRules{
					{
						Hostname:      "bar",
						ImageName:     "foo",
						Port:          "80",
						RequireDigest: helper.BoolPointer(true),
						Tag: &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{
							Negate: helper.BoolPointer(false),
							Value:  "test",
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					AuditKey: true,
					RulesKey: []interface{}{
						map[string]interface{}{
							HostNameKey:  "bar",
							ImageNameKey: "foo",
							PortKey:      "80",
							RequireKey:   true,
							TagKey: []interface{}{
								map[string]interface{}{
									NegateKey: false,
									ValueKey:  "test",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenCustom(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenCustomRules(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CustomRules
		expected    interface{}
	}{
		{
			description: "check for nil rules value of custom recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with all fields of rules spec of custom recipe",
			input: &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CustomRules{
				Hostname:      "bar",
				ImageName:     "foo",
				Port:          "80",
				RequireDigest: helper.BoolPointer(true),
				Tag: &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{
					Negate: helper.BoolPointer(false),
					Value:  "test",
				},
			},
			expected: map[string]interface{}{
				HostNameKey:  "bar",
				ImageNameKey: "foo",
				PortKey:      "80",
				RequireKey:   true,
				TagKey: []interface{}{
					map[string]interface{}{
						NegateKey: false,
						ValueKey:  "test",
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenCustomRules(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
