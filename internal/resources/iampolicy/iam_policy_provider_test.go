// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package iampolicy

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/namespace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

const (
	clusterGroupNameForCluster = "default"
	clusterRole                = "cluster.view"
	subject1Kind               = "GROUP"
	subject1Name               = "test-1"
	testAttached               = "attached"
	testClusteradmin           = "cluster.admin"
	testClusteradmintestuser   = "cluster.admin;test;USER"
	testClustergroupadmin      = "cluster-group.admin"
	testDummy                  = "dummy"
	testTest                   = "test"
	testTest2                  = "test-2"
	testTest3                  = "test-3"
)

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceName:              ResourceIAMPolicy(),
			cluster.ResourceName:      cluster.ResourceTMCCluster(),
			clustergroup.ResourceName: clustergroup.ResourceClusterGroup(),
			workspace.ResourceName:    workspace.ResourceWorkspace(),
			namespace.ResourceName:    namespace.ResourceNamespace(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
