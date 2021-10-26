/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/provider"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "vmware/tanzu/tmc", opts)
		if err != nil {
			log.Fatal(err.Error())
		}

		return
	}

	plugin.Serve(opts)
}
