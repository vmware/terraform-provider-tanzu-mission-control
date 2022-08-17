/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helper

import (
	"fmt"
	"net/url"
)

type RequestURL string

// ConstructRequestURL constructs a request URL based on paths.
func ConstructRequestURL(paths ...string) (url RequestURL) {
	if len(paths) == 0 || paths[0] == "" {
		return url
	}

	url = RequestURL(paths[0])

	for _, value := range paths[1:] {
		url = RequestURL(fmt.Sprintf("%s/%s", url, value))
	}

	return url
}

// AppendQueryParams appends query parameters to a request URL.
func (ru RequestURL) AppendQueryParams(queryParameters url.Values) (url RequestURL) {
	if len(queryParameters) == 0 {
		return ru
	}

	url = RequestURL(fmt.Sprintf("%s?%s", ru, queryParameters.Encode()))

	return url
}

// String converts a request URL to a string.
func (ru RequestURL) String() (urlString string) {
	return string(ru)
}
