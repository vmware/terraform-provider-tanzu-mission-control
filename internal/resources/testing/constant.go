/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// testing is a helper package created for testing purpose.
// go linker would not include this package in the binary, as it is not imported anywhere else other than for testing

package testing

const (
	providerName = "tanzu-mission-control"
	value1       = "value1"
	value2       = "value2"
	description  = "resource with description"

	// if set, used mocked API server
	EKSMockEnv = "ENABLE_EKS_ENV_TEST"
)
