/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ConstructNodepools(t *testing.T) {
	expected := aTestNodePool()
	d := schema.TestResourceDataRaw(t, ClusterSchema, aTestClusterDataMap())

	got := ConstructNodepools(d)
	require.Equal(t, 1, len(got))
	assert.Equal(t, expected.Spec, got[0].Spec)
}

func Test_ConstructNodepools_without_nodepool_type(t *testing.T) {
	expected := aTestNodePool()
	d := schema.TestResourceDataRaw(t, ClusterSchema, aTestClusterDataMap(withoutNodepoolType))

	got := ConstructNodepools(d)
	require.Equal(t, 1, len(got))
	assert.Equal(t, expected.Spec, got[0].Spec)
}

func Test_ToNodepoolMap(t *testing.T) {
	np := aTestNodePool()
	expected := aTestNodepoolDataMap()

	got := ToNodepoolMap(np)
	assert.Equal(t, expected, got)
}
