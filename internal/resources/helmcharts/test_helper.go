// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmcharts

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
)

// nolint: gosec, unused
const (
	chartResource      = ResourceName
	chartDataSourceVar = "test_data_source_repo"

	chartMetadataName = "zookeeper"
)

// nolint: unused
type testAcceptanceConfig struct {
	Provider            *schema.Provider
	ChartResource       string
	ChartDataSourceVar  string
	ChartDataSourceName string
	ChartMetadataName   string
}

// nolint: unused
func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_HELMCHART_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
