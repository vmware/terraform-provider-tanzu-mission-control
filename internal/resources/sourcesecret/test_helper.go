/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package sourcesecret

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
)

// nolint: gosec
const (
	sourceSecretResource      = ResourceName
	sourceSecretResourceVar   = "test_source_secret"
	sourceSecretDataSourceVar = "test_data_source_source_secret"
	sourceSecretNamePrefix    = "tf-ss-test"
)

type testAcceptanceConfig struct {
	Provider                   *schema.Provider
	SourceSecretResource       string
	SourceSecretResourceVar    string
	SourceSecretResourceName   string
	SourceSecretName           string
	ScopeHelperResources       *ScopeHelperResources
	SourceSecretDataSourceVar  string
	SourceSecretDataSourceName string
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_REPOCRED_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
