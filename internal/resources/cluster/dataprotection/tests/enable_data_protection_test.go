//go:build dataprotection
// +build dataprotection

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectiontests

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
	dataprotectionmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/dataprotection"
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

func TestAcceptanceEnableDataProtectionResource(t *testing.T) {
	err := context.Setup()

	if err != nil {
		t.Error(errors.Wrap(err, "unable to set the context"))
		t.FailNow()
	}

	var (
		tfResourceConfigBuilder = InitResourceTFConfigBuilder(testScopeHelper, RsFullBuild)
		provider                = initTestProvider(t)
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: tfResourceConfigBuilder.GetEnableDataProtectionConfig(),
				Check: resource.ComposeTestCheckFunc(
					verifyEnableDataProtectionResourceCreation(provider, EnableDataProtectionResourceFullName),
				),
			},
		},
	},
	)

	t.Log("data protection resource acceptance test complete!")
}

func verifyEnableDataProtectionResourceCreation(
	provider *schema.Provider,
	resourceName string,
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

		fn := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName{
			ClusterName:           testScopeHelper.Cluster.Name,
			ManagementClusterName: testScopeHelper.Cluster.ManagementClusterName,
			ProvisionerName:       testScopeHelper.Cluster.ProvisionerName,
		}

		resp, err := context.TMCConnection.DataProtectionService.DataProtectionResourceServiceList(fn)

		if err != nil {
			return fmt.Errorf("data protection enablement not found, resource: %s | err: %s", resourceName, err)
		}

		if resp == nil || len(resp.DataProtections) == 0 {
			return fmt.Errorf("data protection enablement resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
