/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package data

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	tapeulamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tap/eula"
)

func TestFlattenEULAData(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaData
		expected    interface{}
	}{
		{
			description: "check for nil TAP EULA data",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete TAP EULA data",
			input: &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaData{
				Accepted:   true,
				EulaURL:    "https://network.tanzu.vmware.com/legal_documents/vmware_general_terms",
				ReleasedAt: strfmt.DateTime{},
				User:       "test_uer_id",
			},
			expected: []interface{}{
				map[string]interface{}{
					acceptedKey:   true,
					EulaURLKey:    "https://network.tanzu.vmware.com/legal_documents/vmware_general_terms",
					releasedAtKey: strfmt.DateTime{}.String(),
					userKey:       "test_uer_id",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenEULAData(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
