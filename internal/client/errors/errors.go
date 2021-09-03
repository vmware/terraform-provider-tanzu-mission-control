/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clienterrors

import (
	"fmt"
	"net/http"
	"strings"
)

type ClientErrors struct {
	httpCode int
	err      error
}

func Error(err error) ClientErrors {
	return ClientErrors{
		err: err,
	}
}

func ErrorWithHTTPCode(httpCode int, err error) ClientErrors {
	return ClientErrors{
		httpCode: httpCode,
		err:      err,
	}
}

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	convertedError, ok := err.(ClientErrors)
	if !ok {
		return strings.Contains(err.Error(), fmt.Sprintf("%d", http.StatusNotFound)) &&
			strings.Contains(err.Error(), http.StatusText(http.StatusNotFound))
	}

	if convertedError.httpCode == http.StatusNotFound {
		return true
	}

	return false
}

func (e ClientErrors) Error() string {
	if e.err == nil {
		return ""
	}

	return e.err.Error()
}
