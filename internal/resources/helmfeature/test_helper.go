// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmfeature

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

const (
	helmfeatureResource    = ResourceName
	helmfeatureResourceVar = "test_helm_feature"

	helmReleaseResource    = "tanzu-mission-control_helm_release"
	helmReleaseResourceVar = "test_helm_rel"
	helmReleaseNamePrefix  = "tf-helm-rel-test"
	namespace              = "tanzu-helm-resources"
)

type testAcceptanceConfig struct {
	Provider                *schema.Provider
	HelmFeatureResource     string
	HelmFeatureResourceVar  string
	HelmFeatureResourceName string
	HelmReleaseName         string
	ScopeHelperResources    *commonscope.ScopeHelperResources
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_HELM_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
