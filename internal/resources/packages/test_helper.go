// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackages

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

// nolint: gosec
const (
	PkgsResource        = ResourceName
	pkgsDataSourceVar   = "test_data_source_pkgs"
	pkgsMetadataName    = "cert-manager.tanzu.vmware.com"
	globalRepoNamespace = "tanzu-package-repo-global"

	imageURL = "projects.registry.vmware.com/tmc/build-integrations/package/repository/e2e-test-unauth-repo@sha256:87a5f7e0c44523fbc35a9432c657bebce246138bbd0f16d57f5615933ceef632"
)

type testAcceptanceConfig struct {
	Provider             *schema.Provider
	PkgsResource         string
	PkgsDataSourceVar    string
	PkgsDataSourceName   string
	ScopeHelperResources *commonscope.ScopeHelperResources
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_PKGS_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
