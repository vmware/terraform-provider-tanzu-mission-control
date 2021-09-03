/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helper

import (
	"fmt"
)

const (
	ContentLengthKey = "Content-Length"
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
