/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/stretchr/testify/assert"

	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
)

func Test_ConstructAKSCluster(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ClusterSchema, aTestClusterDataMap())

	result, _ := ConstructCluster(d)
	expected := aTestCluster()

	assert.NotNil(t, result, "no request created")
	assert.Equal(t, expected.FullName, result.FullName, "unexpected full name")
	assert.Equal(t, expected.Spec, result.Spec, "unexpected spec")
}

func Test_ConstructAKSCluster_withInvalidNetworkConfig(t *testing.T) {
	tests := []*schema.ResourceData{
		schema.TestResourceDataRaw(t, ClusterSchema, aTestClusterDataMap(withNetworkPlugin("kubenet"))),
		schema.TestResourceDataRaw(t, ClusterSchema, aTestClusterDataMap(withNetworkPlugin("kubenet"), withoutNetworkDNSServiceIP)),
		schema.TestResourceDataRaw(t, ClusterSchema, aTestClusterDataMap(withNetworkPlugin("kubenet"), withoutNetworkServiceCIDR)),
	}

	for _, d := range tests {
		_, err := ConstructCluster(d)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "can not set network_config.dns_service_ip or network_config.service_cidr when network_config.network_plugin is set to kubenet")
	}
}

func Test_FlattenToMap_nilSpec(t *testing.T) {
	got := ToAKSClusterMap(nil, nil)
	assert.Equal(t, []any{}, got)
}

func Test_FlattenToMap_fullSpec(t *testing.T) {
	testCluster := aTestCluster()
	testNodepool := []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool()}
	expected := aTestClusterDataMap()

	got := ToAKSClusterMap(testCluster, testNodepool)
	assert.Equal(t, expected, got)
}

func Test_FlattenToMap_nilNodepools(t *testing.T) {
	testCluster := aTestCluster()
	expected := aTestClusterDataMap(withoutNodepools)

	got := ToAKSClusterMap(testCluster, nil)
	assert.Equal(t, expected, got)
}
