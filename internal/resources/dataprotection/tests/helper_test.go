//go:build dataprotection
// +build dataprotection

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectiontests

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

var (
	testScopeHelper = commonscope.NewScopeHelperResources()
)

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			dataprotection.ResourceName: dataprotection.ResourceEnableDataProtection(),
			cluster.ResourceName:        cluster.ResourceTMCCluster(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}

	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
