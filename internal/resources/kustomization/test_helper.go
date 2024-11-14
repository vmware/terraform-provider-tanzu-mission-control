// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kustomization

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	gitrepositoryhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository"
)

// nolint: gosec
const (
	kustomizationResource      = ResourceName
	kustomizationResourceVar   = "test_kustomization"
	kustomizationDataSourceVar = "test_data_source_kustomization"
	kustomizationNamePrefix    = "tf-kustomization-test"

	gitRepositoryResource    = gitrepositoryhelper.ResourceName
	gitRepositoryResourceVar = "test_git_repository"
	gitRepositoryNamePrefix  = "tf-gr-test"
)

type testAcceptanceConfig struct {
	Provider                  *schema.Provider
	KustomizationResource     string
	KustomizationResourceVar  string
	KustomizationResourceName string
	KustomizationName         string
	ScopeHelperResources      *commonscope.ScopeHelperResources
	GitRepositoryResource     string
	GitRepositoryResourceVar  string
	GitRepositoryName         string
	Namespace                 string
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_KUSTOMIZATION_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
