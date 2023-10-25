/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmcharts

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
)

// nolint: gosec
const (
	chartResource      = ResourceName
	chartDataSourceVar = "test_data_source_repo"

	chartMetadataName = "zookeeper"
)

type testAcceptanceConfig struct {
	Provider            *schema.Provider
	ChartResource       string
	ChartDataSourceVar  string
	ChartDataSourceName string
	ChartMetadataName   string
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_HELMCHART_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
