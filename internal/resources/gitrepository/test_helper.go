/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepository

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

const (
	gitRepositoryResource      = ResourceName
	gitRepositoryResourceVar   = "test_git_repository"
	gitRepositoryDataSourceVar = "test_data_source_git_repository"
	gitRepositoryNamePrefix    = "tf-gr-test"
)

type testAcceptanceConfig struct {
	Provider                    *schema.Provider
	GitRepositoryResource       string
	GitRepositoryResourceVar    string
	GitRepositoryResourceName   string
	GitRepositoryName           string
	ScopeHelperResources        *commonscope.ScopeHelperResources
	GitRepositoryDataSourceVar  string
	GitRepositoryDataSourceName string
	Namespace                   string
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_GITREPO_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
