// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package converter

import (
	"fmt"
	"strings"
)

func BuildModelPath(modelPathSeparator string, paths ...string) string {
	return strings.Join(paths, modelPathSeparator)
}

func BuildDefaultModelPath(paths ...string) string {
	return BuildModelPath(DefaultModelPathSeparator, paths...)
}

func BuildArrayField(field string) string {
	return fmt.Sprintf("%s%s", field, ArrayFieldMarker)
}
