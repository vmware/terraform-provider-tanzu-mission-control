/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackageinstall

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	packagerepository "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository"
)

// nolint: gosec
const (
	pkgRepoResource      = packagerepository.ResourceName
	pkgRepoResourceVar   = "test_pkg_repository"
	pkgRepoDataSourceVar = "test_data_source_pkg_repository"
	pkgRepoNamePrefix    = "tf-pkg-repository-test"

	pkgInstallResource      = ResourceName
	pkgInstallResourceVar   = "test_pkg_install"
	pkgInstallDataSourceVar = "test_data_source_pkg_install"
	pkgInstallNamePrefix    = "tf-pkg-install-test"
	namespaceNamePrefix     = "test-pkg-install-ns"

	constraints = "3.0.0-rc.1"
)

type testAcceptanceConfig struct {
	Provider               *schema.Provider
	PkgInstallResource     string
	PkgInstallResourceVar  string
	PkgInstallResourceName string
	PkgInstallName         string
	PkgRepoName            string
	ScopeHelperResources   *commonscope.ScopeHelperResources
	Namespace              string
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_PKGINS_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
