/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package common

import (
	"fmt"
	"time"
)

func GetFirstElementOf(parent string, children ...string) (key string) {
	if len(children) == 0 {
		return parent
	}

	key = parent
	for _, value := range children {
		key = fmt.Sprintf("%s.0.%s", key, value)
	}

	return key
}

// Retryable is a simple function which can be retried, returns (retry[yes/no], error).
type Retryable func() (bool, error)

// Retry is a wrapper to retry functions.
func Retry(f Retryable, interval time.Duration, attempts int) (int, error) {
	var (
		err   error
		retry bool
	)

	retries := 0
	for retries < attempts {
		retry, err = f()
		if !retry {
			break
		}

		time.Sleep(interval)
		retries++
	}

	return retries, err
}
