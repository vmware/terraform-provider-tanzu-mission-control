//go:build backupschedule
// +build backupschedule

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupscheduletests

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/dataprotection/cluster/backupschedule"
	backupscheduleres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/backupschedule"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

var (
	context = authctx.TanzuContext{
		ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
		Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
		VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		TLSConfig:        &proxy.TLSConfig{},
	}
)

func TestAcceptanceBackupScheduleResource(t *testing.T) {
	err := context.Setup()

	if err != nil {
		t.Error(errors.Wrap(err, "unable to set the context"))
		t.FailNow()
	}

	tmcManagedCredentialsName, tmcManagedCredentialsExist := os.LookupEnv(tmcManagedCredentialsEnv)

	if !tmcManagedCredentialsExist {
		t.Error("TMC Managed credentials name is missing!")
		t.FailNow()
	}

	var (
		tfResourceConfigBuilder   = InitResourceTFConfigBuilder(testScopeHelper, RsFullBuild, tmcManagedCredentialsName)
		tfDataSourceConfigBuilder = InitDataSourceTFConfigBuilder(testScopeHelper, tfResourceConfigBuilder, DsFullBuild)
		provider                  = initTestProvider(t)
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: tfResourceConfigBuilder.GetFullClusterBackupScheduleConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(FullClusterBackupScheduleResourceFullName, "name", FullClusterBackupScheduleName),
					verifyBackupScheduleResourceCreation(provider, FullClusterBackupScheduleResourceFullName, FullClusterBackupScheduleName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetNamespacesBackupScheduleConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(NamespacesBackupScheduleResourceFullName, "name", NamespacesBackupScheduleName),
					verifyBackupScheduleResourceCreation(provider, NamespacesBackupScheduleResourceFullName, NamespacesBackupScheduleName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetLabelsBackupScheduleConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(LabelsBackupScheduleResourceFullName, "name", LabelsBackupScheduleName),
					verifyBackupScheduleResourceCreation(provider, LabelsBackupScheduleResourceFullName, LabelsBackupScheduleName),
				),
			},
			{
				Config: tfDataSourceConfigBuilder.GetDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					verifyBackupScheduleDataSource(provider, DataSourceFullName, LabelsBackupScheduleName),
				),
			},
		},
	},
	)

	t.Log("backup schedule resource acceptance test complete!")
}

func verifyBackupScheduleResourceCreation(
	provider *schema.Provider,
	resourceName string,
	backupScheduleName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("could not found resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource %s", resourceName)
		}

		fn := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName{
			Name:                  backupScheduleName,
			ManagementClusterName: testScopeHelper.Cluster.ManagementClusterName,
			ClusterName:           testScopeHelper.Cluster.Name,
			ProvisionerName:       testScopeHelper.Cluster.ProvisionerName,
		}

		resp, err := context.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceGet(fn)

		if err != nil {
			return fmt.Errorf("target location resource not found, resource: %s | err: %s", resourceName, err)
		}

		if resp == nil {
			return fmt.Errorf("target location resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}

func verifyBackupScheduleDataSource(
	provider *schema.Provider,
	dataSourceName string,
	backupScheduleName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[dataSourceName]

		if !ok {
			return fmt.Errorf("could not found data source %s", dataSourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, data source %s", dataSourceName)
		}

		firstTargetLocation := fmt.Sprintf("%s.0.%s", backupscheduleres.SchedulesKey, backupscheduleres.NameKey)

		if rs.Primary.Attributes[firstTargetLocation] != backupScheduleName {
			return fmt.Errorf("target location wasn't found at index 0 (%s)", backupScheduleName)
		}

		return nil
	}
}
