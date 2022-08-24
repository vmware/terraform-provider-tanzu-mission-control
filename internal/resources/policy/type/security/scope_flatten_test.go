/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package security

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	scoperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

func TestFlattenScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description  string
		input        *scopedFullname
		expectedData []interface{}
		expectedName string
	}{
		{
			description:  "check for nil scope",
			input:        nil,
			expectedData: nil,
			expectedName: "",
		},
		{
			description: "normal scenario with complete cluster scope",
			input: &scopedFullname{
				scope: clusterScope,
				fullnameCluster: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName{
					Name:                  "n",
					ClusterName:           "c",
					ManagementClusterName: "m",
					ProvisionerName:       "p",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					clusterKey: []interface{}{
						map[string]interface{}{
							scoperesource.ManagementClusterNameKey: "m",
							scoperesource.ClusterNameKey:           "c",
							scoperesource.ProvisionerNameKey:       "p",
						},
					},
				},
			},
			expectedName: "n",
		},
		{
			description: "normal scenario with complete cluster group scope",
			input: &scopedFullname{
				scope: clusterGroupScope,
				fullnameClusterGroup: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
					Name:             "n",
					ClusterGroupName: "c",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					clusterGroupKey: []interface{}{
						map[string]interface{}{
							scoperesource.ClusterGroupNameKey: "c",
						},
					},
				},
			},
			expectedName: "n",
		},
		{
			description: "normal scenario with complete organization scope",
			input: &scopedFullname{
				scope: organizationScope,
				fullnameOrganization: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
					Name:  "n",
					OrgID: "o",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					organizationKey: []interface{}{
						map[string]interface{}{
							scoperesource.OrganizationIDKey: "o",
						},
					},
				},
			},
			expectedName: "n",
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actualData, actualName := flattenScope(test.input)
			require.Equal(t, test.expectedData, actualData)
			require.Equal(t, test.expectedName, actualName)
		})
	}
}
