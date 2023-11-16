//go:build backupschedule
// +build backupschedule

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupscheduletests

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	backupscheduleres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/backupschedule"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
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
			targetlocationres.ResourceName: targetlocationres.ResourceTargetLocation(),
			dataprotection.ResourceName:    dataprotection.ResourceEnableDataProtection(),
			backupscheduleres.ResourceName: backupscheduleres.ResourceBackupSchedule(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			backupscheduleres.ResourceName: backupscheduleres.DataSourceBackupSchedule(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}

	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
