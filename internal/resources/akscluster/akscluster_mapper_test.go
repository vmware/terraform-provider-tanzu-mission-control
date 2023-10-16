/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/stretchr/testify/assert"

	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/akscluster"
)

func Test_ConstructAKSCluster(t *testing.T) {
	d := schema.TestResourceDataRaw(t, akscluster.ClusterSchema, aTestClusterDataMap())

	result, _ := akscluster.ConstructCluster(d)
	expected := aTestCluster()

	assert.NotNil(t, result, "no request created")
	assert.Equal(t, expected.FullName, result.FullName, "unexpected full name")
	assert.Equal(t, expected.Spec, result.Spec, "unexpected spec")
}

func Test_FlattenToMap_nilSpec(t *testing.T) {
	got := akscluster.ToAKSClusterMap(nil, nil)
	assert.Equal(t, []any{}, got)
}

func Test_FlattenToMap_fullSpec(t *testing.T) {
	testCluster := aTestCluster()
	testNodepool := []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool()}
	expected := aTestClusterDataMap()

	got := akscluster.ToAKSClusterMap(testCluster, testNodepool)
	assert.Equal(t, expected, got)
}

func Test_FlattenToMap_nilNodepools(t *testing.T) {
	testCluster := aTestCluster()
	expected := aTestClusterDataMap(withoutNodepools)

	got := akscluster.ToAKSClusterMap(testCluster, nil)
	assert.Equal(t, expected, got)
}
