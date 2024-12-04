//go:build backupschedule
// +build backupschedule

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package backupscheduletests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backupschedule"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	dataprotectionres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection"
	targetlocationres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/targetlocation"
)

const (
	tmcManagedCredentialsEnv = "TMC_MANAGED_CREDENTIALS_NAME" // #nosec G101
)

var (
	testScopeHelper = commonscope.NewScopeHelperResources()
)

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			cluster.ResourceName:           cluster.ResourceTMCCluster(),
			clustergroup.ResourceName:      clustergroup.ResourceClusterGroup(),
			targetlocationres.ResourceName: targetlocationres.ResourceTargetLocation(),
			dataprotectionres.ResourceName: dataprotectionres.ResourceEnableDataProtection(),
			backupschedule.ResourceName:    backupschedule.ResourceBackupSchedule(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			backupschedule.ResourceName: backupschedule.DataSourceBackupSchedule(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}

	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
