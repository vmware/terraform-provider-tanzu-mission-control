/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
)

func TestNodepoolPosMap(t *testing.T) {
	tests := []struct {
		name string
		nps  []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition
		res  map[string]int
	}{
		{
			name: "empty list",
			nps:  []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{},
			res:  map[string]int{},
		},
		{
			name: "with some nodepool defs",
			nps: []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
				{
					Info: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{
						Name: "np-1",
					},
					Spec: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec{
						AmiType: "some-type",
						RoleArn: "some-arn",
					},
				},
				{
					Info: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{
						Name: "a-np-2",
					},
					Spec: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec{
						AmiType: "some-type-2",
						RoleArn: "some-arn-2",
					},
				},
			},
			res: map[string]int{
				"np-1":   0,
				"a-np-2": 1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.res, nodepoolPosMap(test.nps), "expected function output to match")
		})
	}
}
