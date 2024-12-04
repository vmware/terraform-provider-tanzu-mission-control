// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/provider"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	}

	if debugMode {
		opts.Debug = debugMode
		opts.ProviderAddr = "vmware/dev/tanzu-mission-control"
	}

	plugin.Serve(opts)
}
