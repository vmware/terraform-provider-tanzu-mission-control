/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrepository

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

// nolint: gosec, unused
const (
	repoResource      = ResourceName
	repoDataSourceVar = "test_data_source_repo"
)

// nolint: unused
type testAcceptanceConfig struct {
	Provider             *schema.Provider
	RepoResource         string
	RepoDataSourceVar    string
	RepoDataSourceName   string
	Namespace            string
	CgName               string
	ScopeHelperResources *commonscope.ScopeHelperResources
}

// nolint: unused
func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_HELMREPO_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
