/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tapeula

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
)

// nolint: gosec, unused
const (
	tapEULAResource      = ResourceName
	tapEULAResourceVar   = "test_tap_eula"
	tapEULADataSourceVar = "test_data_tap_eula"
	tapEULATAPVersion    = "1.8.1"
)

// nolint: unused
type testAcceptanceConfig struct {
	Provider                    *schema.Provider
	TAPEULAResource             string
	TAPEULAResourceVar          string
	TAPEULAResourceTAPVersion   string
	TAPEULATAPVersion           string
	TAPEULADataSourceVar        string
	TAPEULADataSourceTAPVersion string
}

// nolint: unused
func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_TAPEULA_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
