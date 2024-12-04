// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
)

func TestFlattenOrganizationPolicyFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName
		expected    []interface{}
	}{
		{
			description: "check for nil organization policy full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete organization policy full name",
			input: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
				OrgID: "o",
			},
			expected: []interface{}{
				map[string]interface{}{
					OrganizationIDKey: "o",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenOrganizationPolicyFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
