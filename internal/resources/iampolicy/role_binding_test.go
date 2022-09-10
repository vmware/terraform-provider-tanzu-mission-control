/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

func TestGetRoleBindingSchemaData(t *testing.T) {
	d := new(schema.Resource)
	d.Schema = iamPolicySchema

	rd := d.TestResourceData()

	err := rd.Set(
		roleBindingsKey,
		[]interface{}{
			map[string]interface{}{
				roleKey: "dummy-role",
				subjectsKey: []interface{}{
					map[string]interface{}{
						subjectNameKey: "sub",
						subjectKindKey: "kind",
					},
					map[string]interface{}{
						subjectNameKey: "sub2",
						subjectKindKey: "kind2",
					},
				},
			},
			map[string]interface{}{
				roleKey: "dummy-role-2",
				subjectsKey: []interface{}{
					map[string]interface{}{
						subjectNameKey: "sub",
						subjectKindKey: "kind",
					},
					map[string]interface{}{
						subjectNameKey: "sub2",
						subjectKindKey: "GROUP",
					},
				},
			},
		})

	require.NoError(t, err)
}

func TestRoleBindingListToUpdate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType
		expected    []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta
	}{
		{
			description: "check for normal scenario of an entry in map",
			input: &map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType{
				"cluster.admin_test_USER": iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD,
			},
			expected: []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta{
				{
					Op:   iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD.Pointer(),
					Role: "cluster.admin",
					Subject: &iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
						Name: "test",
						Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := getRoleBindingListToUpdate(test.input)
			require.EqualValues(t, test.expected, actual)
		})
	}
}

func TestRBOpForUpdate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		inputRBL    []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding
		inputMap    map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType
		action      iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType
		expectedMap map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType
	}{
		{
			description: "check for action as UNSPECIFIED",
			inputRBL: []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
				{
					Role: "cluster.admin",
					Subjects: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
						{
							Name: "test",
							Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
						},
					},
				},
			},
			inputMap: map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType{
				"cluster.admin_test_USER": iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD,
			},
			action: iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeOPTYPEUNSPECIFIED,
			expectedMap: map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType{
				"cluster.admin_test_USER": iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeOPTYPEUNSPECIFIED,
			},
		},
		{
			description: "check for action as DELETE",
			inputRBL: []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
				{
					Role: "cluster.admin",
					Subjects: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
						{
							Name: "test",
							Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
						},
					},
				},
			},
			inputMap: map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType{
				"cluster.admin_test_USER": iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD,
			},
			action: iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeDELETE,
			expectedMap: map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType{
				"cluster.admin_test_USER": iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeDELETE,
			},
		},
		{
			description: "check for action as ADD",
			inputRBL: []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
				{
					Role: "cluster.admin",
					Subjects: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
						{
							Name: "test",
							Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
						},
					},
				},
			},
			inputMap: map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType{
				"cluster.admin_test_USER": iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD,
			},
			action: iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD,
			expectedMap: map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType{
				"cluster.admin_test_USER": iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD,
			},
		},
		{
			description: "check for flip action scenario",
			inputRBL: []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
				{
					Role: "cluster.admin",
					Subjects: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
						{
							Name: "test",
							Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
						},
					},
				},
			},
			inputMap: map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType{
				"cluster.admin_test_USER": iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeDELETE,
			},
			action: iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD,
			expectedMap: map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType{
				"cluster.admin_test_USER": iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeOPTYPEUNSPECIFIED,
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			constructRBOpForUpdate(test.inputRBL, &test.inputMap, test.action)
			require.EqualValues(t, test.expectedMap, test.inputMap)
		})
	}
}

func TestGetIntersectionOfSubs(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		state       []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject
		server      []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject
		expected    []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject
	}{
		{
			description: "check for when state and server are empty lists",
			state:       []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{},
			server:      []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{},
			expected:    []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject(nil),
		},
		{
			description: "check for when state and server have no common subjects",
			state: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
				{
					Name: "test-1",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
				},
				{
					Name: "test-2",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP.Pointer(),
				},
			},
			server: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
				{
					Name: "test-1",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP.Pointer(),
				},
				{
					Name: "test-3",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
				},
			},
			expected: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject(nil),
		},
		{
			description: "check for when state and server have subjects in common",
			state: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
				{
					Name: "test-1",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
				},
				{
					Name: "test-2",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP.Pointer(),
				},
				{
					Name: "test-3",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP.Pointer(),
				},
			},
			server: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
				{
					Name: "test-1",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
				},
				{
					Name: "test-3",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
				},
				{
					Name: "test-4",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP.Pointer(),
				},
			},
			expected: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
				{
					Name: "test-1",
					Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := getIntersectionOfSubs(test.state, test.server)
			require.EqualValues(t, test.expected, actual)
		})
	}
}
