/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backupschedule/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestConstructClusterBackupScheduleFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       []interface{}
		expected    *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName
	}{
		{
			description: "check for nil cluster backup schedule full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster backup schedule full name",
			input: []interface{}{
				map[string]interface{}{
					commonscope.NameKey:                  "c",
					commonscope.ManagementClusterNameKey: "m",
					commonscope.ProvisionerNameKey:       "p",
				},
			},
			expected: &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName{
				ClusterName:           "c",
				ManagementClusterName: "m",
				ProvisionerName:       "p",
				Name:                  "test-schedule",
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := ConstructClusterBackupScheduleFullname(test.input, "test-schedule")
			require.Equal(t, test.expected, actual)
		})
	}
}
