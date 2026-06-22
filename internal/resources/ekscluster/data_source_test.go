// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package ekscluster

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
)

const (
	test00001                   = "0.0.0.0/1"
	test1000010                 = "10.0.0.0/10"
	test10001                   = "1.0.0.0/1"
	test100101024               = "100.10.1.0/24"
	test112                     = "1.12"
	test126420230703            = "1.26.4-20230703"
	test196191024               = "196.19.1.0/24"
	testAmi2qu8409oisdfj0qw     = "ami-2qu8409oisdfj0qw"
	testArnAwsIam000000000000   = "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com"
	testBootstrapCmd            = "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami"
	testCustom                  = "CUSTOM"
	testDescription             = "description"
	testKey1                    = "key1"
	testKey2                    = "key2"
	testM3large                 = "m3.large"
	testNewtesttag              = "newtesttag"
	testOnDemand                = "ON_DEMAND"
	testReady                   = "ready"
	testResourceWithDescription = "resource with description"
	testRoleArn                 = "role-arn"
	testSg0a6768722e9716768     = "sg-0a6768722e9716768"
	testSg1                     = "sg-1"
	testSg2                     = "sg-2"
	testSg3                     = "sg-3"
	testSg4                     = "sg-4"
	testSgXxxxxxx               = "sg-xxxxxxx"
	testSubnet022e4a6bc8c8ee7a6 = "subnet-022e4a6bc8c8ee7a6"
	testSubnet06497e6063c209f4d = "subnet-06497e6063c209f4d"
	testSubnet0a184f9301ae39a86 = "subnet-0a184f9301ae39a86"
	testSubnet0b495d7c212fc92a1 = "subnet-0b495d7c212fc92a1"
	testSubnet0c86ec9ecde7b9bf7 = "subnet-0c86ec9ecde7b9bf7"
	testSubnet1                 = "subnet-1"
	testSubnet2                 = "subnet-2"
	testSubnet3                 = "subnet-3"
	testSubnetId1               = "subnet-id-1"
	testSubnetId2               = "subnet-id-2"
	testT3medium                = "t3.medium"
	testTag1                    = "tag1"
	testTag2                    = "tag2"
	testTempl                   = "templ"
	testTestKey                 = "test-key"
	testTesttag                 = "testtag"
	testTestCg                  = "test-cg"
	testTestNp                  = "test np"
	testTestNp2                 = "test-np"
	testTestProoxy              = "test-prooxy"
	testTg1                     = "tg1"
	testTkey                    = "tkey"
	testTkey2                   = "tkey2"
	testTvalue                  = "tvalue"
	testTvalue2                 = "tvalue2"
	testTv1                     = "tv1"
	testVal1                    = "val1"
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

func TestIsManagemetClusterHealthy(t *testing.T) {
	tests := []struct {
		name     string
		cluster  *clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse
		response bool
		err      error
	}{
		{
			name: "Not healthy",
			cluster: &clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse{
				Cluster: &clustermodel.VmwareTanzuManageV1alpha1ClusterCluster{
					Status: &clustermodel.VmwareTanzuManageV1alpha1ClusterStatus{
						Health: clustermodel.NewVmwareTanzuManageV1alpha1CommonClusterHealth(clustermodel.VmwareTanzuManageV1alpha1CommonClusterHealthUNHEALTHY),
					},
				},
			},
			response: false,
			err:      nil,
		},
		{
			name: "Healthy",
			cluster: &clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse{
				Cluster: &clustermodel.VmwareTanzuManageV1alpha1ClusterCluster{
					Status: &clustermodel.VmwareTanzuManageV1alpha1ClusterStatus{
						Health: clustermodel.NewVmwareTanzuManageV1alpha1CommonClusterHealth(clustermodel.VmwareTanzuManageV1alpha1CommonClusterHealthHEALTHY),
					},
				},
			},
			response: true,
			err:      nil,
		},
		{
			name: "Error",
			cluster: &clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse{
				Cluster: &clustermodel.VmwareTanzuManageV1alpha1ClusterCluster{
					Status: nil,
				},
			},
			response: false,
			err:      errors.New("cluster data is invalid or nil"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := isClusterHealthy(test.cluster)
			if err != nil {
				if err.Error() != test.err.Error() {
					t.Errorf("expected error to match")
				}
			} else if test.response != result {
				t.Errorf("expected function output to match")
			}
		})
	}
}
