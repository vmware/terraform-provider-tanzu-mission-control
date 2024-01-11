/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	cgbackupschedulemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backupschedule/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestConstructClusterGroupBackupScheduleFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       []interface{}
		expected    *cgbackupschedulemodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleFullName
	}{
		{
			description: "check for nil cluster group backup schedule full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group backup schedule full name",
			expected: &cgbackupschedulemodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleFullName{
				ClusterGroupName: "c",
				Name:             "sch",
			},
			input: []interface{}{
				map[string]interface{}{
					commonscope.ClusterGroupNameKey: "c",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := ConstructClusterGroupBackupScheduleFullname(test.input, "sch")
			require.Equal(t, test.expected, actual)
		})
	}
}
