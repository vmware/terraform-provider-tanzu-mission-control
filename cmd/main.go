/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
