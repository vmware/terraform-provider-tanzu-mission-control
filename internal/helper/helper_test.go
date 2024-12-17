// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

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

func TestSetPrimitiveValueFloat64toFloat32(t *testing.T) {
	var (
		f32           float32
		f64           float64
		i64           int64
		s             string
		b             bool
		interfaceType interface{}
	)

	// pass: convert string to string
	SetPrimitiveValue("TMC", &s, "key")
	require.Equal(t, s, "TMC")

	// pass: convert bool to bool
	SetPrimitiveValue(true, &b, "key")
	require.Equal(t, b, true)

	// pass: convert int to int64
	SetPrimitiveValue(10, &i64, "key")
	require.Equal(t, i64, int64(10))

	// pass: convert float64 to float32
	SetPrimitiveValue(4.66, &f32, "key")
	require.Equal(t, f32, float32(4.66))

	// fail: convert string to float64
	SetPrimitiveValue("4.55", &f64, "key")
	require.Equal(t, f64, float64(0))

	// fail: convert non-primitive type to float32
	SetPrimitiveValue(interfaceType, &f32, "key")
	require.Equal(t, f32, float32(0))
}

func TestRetryUntilTimeout(t *testing.T) {
	testFun := func() (bool, error) {
		return true, nil
	}

	retries, err := RetryUntilTimeout(testFun, 1*time.Second, 2*time.Second)

	require.NoError(t, err)
	require.Equal(t, 2, retries)
}
