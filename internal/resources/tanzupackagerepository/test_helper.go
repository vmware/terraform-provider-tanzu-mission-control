/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

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
	PkgRepoResource      = ResourceName
	PkgRepoResourceVar   = "test_pkg_repository"
	pkgRepoDataSourceVar = "test_data_source_pkg_repository"
	pkgRepoNamePrefix    = "tf-pkg-repository-test"
	globalRepoNamespace  = "tanzu-package-repo-global"

	imageURL        = "extensions.aws-usw2.tmc-dev.cloud.vmware.com/packages/standard/repo:v2.2.0_update.2"
	updatedImageURL = "extensions.aws-usw2.tmc-dev.cloud.vmware.com/packages/standard/repo:v2.2.0_update.1"
)

type testAcceptanceConfig struct {
	Provider                    *schema.Provider
	PkgRepoResource             string
	PkgRepoResourceVar          string
	PkgRepoResourceName         string
	PkgRepoName                 string
	PkgRepoDataSourceVar        string
	PkgRepositoryDataSourceName string
	ScopeHelperResources        *commonscope.ScopeHelperResources
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_PKGREPO_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
