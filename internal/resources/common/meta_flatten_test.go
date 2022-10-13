/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package common

import (
	"testing"

	"github.com/stretchr/testify/require"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
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
				Annotations: map[string]string{"test": "test"},
				Labels:      map[string]string{"test": "test"},
				Description: "description of resource",
				UID:         "abc",
			},
			expected: []interface{}{
				map[string]interface{}{
					annotationsKey:     map[string]string{"test": "test"},
					LabelsKey:          map[string]string{"test": "test"},
					DescriptionKey:     "description of resource",
					resourceVersionKey: "",
					uidKey:             "abc",
				},
			},
		},
		{
			name: "normal scenario with annotation and description of meta data",
			input: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				Annotations: map[string]string{"test": "test"},
				Labels:      map[string]string{},
				Description: "description of resource",
				UID:         "",
			},
			expected: []interface{}{
				map[string]interface{}{
					annotationsKey:     map[string]string{"test": "test"},
					LabelsKey:          map[string]string{},
					DescriptionKey:     "description of resource",
					resourceVersionKey: "",
					uidKey:             "",
				},
			},
		},
		{
			name: "normal scenario with labels and UID of meta data",
			input: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				Annotations: map[string]string{},
				Labels:      map[string]string{"test": "test"},
				Description: "",
				UID:         "123",
			},
			expected: []interface{}{
				map[string]interface{}{
					annotationsKey:     map[string]string{},
					LabelsKey:          map[string]string{"test": "test"},
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
