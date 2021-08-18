/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testProvider *schema.Provider

func init() {
	testProvider = Provider()
}

func TestProvider(t *testing.T) {
	fmt.Printf("provider internal validate")

	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}
