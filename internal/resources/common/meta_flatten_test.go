// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"testing"

	"github.com/stretchr/testify/require"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

const (
	testDescriptionOfResource = "description of resource"
	testTest                  = "test"
)

func TestFlattenMeta(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
		expected []interface{}
	}{
		{
			name:     "check for nil meta data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with all fields of meta data",
			input: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				Annotations: map[string]string{testTest: testTest},
				Labels:      map[string]string{testTest: testTest},
				Description: testDescriptionOfResource,
				UID:         "abc",
			},
			expected: []interface{}{
				map[string]interface{}{
					AnnotationsKey:     map[string]string{testTest: testTest},
					LabelsKey:          map[string]string{testTest: testTest},
					DescriptionKey:     testDescriptionOfResource,
					resourceVersionKey: "",
					uidKey:             "abc",
				},
			},
		},
		{
			name: "normal scenario with annotation and description of meta data",
			input: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				Annotations: map[string]string{testTest: testTest},
				Labels:      map[string]string{},
				Description: testDescriptionOfResource,
				UID:         "",
			},
			expected: []interface{}{
				map[string]interface{}{
					AnnotationsKey:     map[string]string{testTest: testTest},
					LabelsKey:          map[string]string{},
					DescriptionKey:     testDescriptionOfResource,
					resourceVersionKey: "",
					uidKey:             "",
				},
			},
		},
		{
			name: "normal scenario with labels and UID of meta data",
			input: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				Annotations: map[string]string{},
				Labels:      map[string]string{testTest: testTest},
				Description: "",
				UID:         "123",
			},
			expected: []interface{}{
				map[string]interface{}{
					AnnotationsKey:     map[string]string{},
					LabelsKey:          map[string]string{testTest: testTest},
					DescriptionKey:     "",
					resourceVersionKey: "",
					uidKey:             "123",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := FlattenMeta(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
