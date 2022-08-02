/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helper

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstructRequestURL(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		paths    []string
		expected RequestURL
	}{
		{
			name:     "case for no paths",
			paths:    []string{},
			expected: "",
		},
		{
			name:     "case for a single path",
			paths:    []string{"p1"},
			expected: "p1",
		},
		{
			name:     "case for multiple successive paths",
			paths:    []string{"p1", "p2", "p3"},
			expected: "p1/p2/p3",
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := ConstructRequestURL(test.paths...)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestAppendQueryParams(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name            string
		requestURL      RequestURL
		queryParameters url.Values
		expected        RequestURL
	}{
		{
			name:            "case for no query parameter",
			requestURL:      "p1",
			queryParameters: nil,
			expected:        "p1",
		},
		{
			name:       "case for one query parameter",
			requestURL: "p1",
			queryParameters: url.Values{
				"k1": []string{"v1"},
			},
			expected: "p1?k1=v1",
		},
		{
			name:       "case for multiple query parameters",
			requestURL: "p1/p2/p3",
			queryParameters: url.Values{
				"k1": []string{"v1"},
				"k2": []string{"v2"},
			},
			expected: "p1/p2/p3?k1=v1&k2=v2",
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := test.requestURL.AppendQueryParams(test.queryParameters)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name       string
		requestURL RequestURL
		expected   string
	}{
		{
			name:       "case for empty path",
			requestURL: "",
			expected:   "",
		},
		{
			name:       "case for multiple paths and query parameters",
			requestURL: "p1/p2/p3?k1=v1&k2=v2",
			expected:   "p1/p2/p3?k1=v1&k2=v2",
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := test.requestURL.String()
			require.Equal(t, test.expected, actual)
		})
	}
}
