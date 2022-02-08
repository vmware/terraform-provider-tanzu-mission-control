/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetFirstElementOf(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		parent   string
		children []string
		expected string
	}{
		{
			name:     "case for no children",
			parent:   "p1",
			children: []string{},
			expected: "p1",
		},
		{
			name:     "case when parent has one child",
			parent:   "p2",
			children: []string{"c1"},
			expected: "p2.0.c1",
		},
		{
			name:     "case when parent has more than one child",
			parent:   "p3",
			children: []string{"c2", "c3", "c4"},
			expected: "p3.0.c2.0.c3.0.c4",
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := GetFirstElementOf(test.parent, test.children...)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestRetryUntilTimeout(t *testing.T) {
	testFun := func() (bool, error) {
		return true, nil
	}

	retries, err := RetryUntilTimeout(testFun, 1*time.Second, 6*time.Second)

	require.NoError(t, err)
	require.Equal(t, 4, retries)
}
