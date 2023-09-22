/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrelease

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	helm "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmfeature"
)

// nolint: gosec
const (
	helmReleaseResource      = ResourceName
	helmReleaseResourceVar   = "test_helm_rel"
	gitRepoResourceVar       = "test_git_repo"
	gitRepoName              = "test-git-repo"
	helmReleaseDataSourceVar = "test_data_source_helm_rel"
	helmReleaseNamePrefix    = "tf-helm-rel-test"
	namespaceNamePrefix      = "test-helm-rel-ns"

	helmfeatureResource    = helm.ResourceName
	helmfeatureResourceVar = "test_helm_feature"

	constraints = "15.0.5"
)

type testAcceptanceConfig struct {
	Provider                *schema.Provider
	HelmReleaseResource     string
	HelmFeatureResource     string
	HelmReleaseResourceVar  string
	HelmFeatureResourceVar  string
	HelmReleaseResourceName string
	HelmReleaseName         string
	ScopeHelperResources    *commonscope.ScopeHelperResources
	Namespace               string
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_HELMRELEASE_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
