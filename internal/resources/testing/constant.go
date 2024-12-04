// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// testing is a helper package created for testing purpose.
// go linker would not include this package in the binary, as it is not imported anywhere else other than for testing

package testing

const (
	providerName = "tanzu-mission-control"
	value1       = "value1"
	value2       = "value2"
	description  = "resource with description"
)
