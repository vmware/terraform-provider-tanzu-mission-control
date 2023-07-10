/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
)

const (
	secretResourceVar   = "test_secret"
	secretDataSourceVar = "test_data_secret"
	clusterResource     = "tanzu-mission-control_cluster"
	clusterResourceVar  = "tmc_cluster_test"
)

type testAcceptanceConfig struct {
	Provider           *schema.Provider
	SecretResource     string
	SecretResourceVar  string
	SecretResourceName string
	SecretName         string
	ClusterName        string
	NamespaceName      string
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_SECRET_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}
