/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"testing"

	"github.com/stretchr/testify/require"

	organizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/organization"
)

func TestFlattenOrganizationFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName
		expected    []interface{}
	}{
		{
			description: "check for nil organization full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete organization full name",
			input: &organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName{
				OrgID: "default",
			},
			expected: []interface{}{
				map[string]interface{}{
					organizationIDKey: "default",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenOrganizationFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
