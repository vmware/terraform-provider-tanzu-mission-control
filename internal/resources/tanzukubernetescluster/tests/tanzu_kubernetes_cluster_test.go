//go:build tanzukubernetescluster
// +build tanzukubernetescluster

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzuekubernetesclustertests

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
	tanzukubernetesclustermodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
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

func TestAcceptanceUTKGResource(t *testing.T) {
	err := context.Setup()

	if err != nil {
		t.Error(errors.Wrap(err, "unable to set the context"))
		t.FailNow()
	}

	// See cluster_env_vars.go for required environment variables.
	environmentVars, errs := ReadClusterEnvironmentVariables()

	if len(errs) > 0 {
		errMsg := ""

		for _, e := range errs {
			errMsg = fmt.Sprintf("%s\n%s", errMsg, e.Error())
		}

		t.Error(errors.Errorf("Required environment variables are missing: %s", errMsg))
		t.FailNow()
	}

	var (
		provider                = initTestProvider(t)
		tfResourceConfigBuilder = InitResourceTFConfigBuilder()
		tkgmEnvironmentVars     = environmentVars[TKGMClusterType]
		tkgsEnvironmentVars     = environmentVars[TKGSClusterType]
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: tfResourceConfigBuilder.GetTKGMClusterConfig(tkgmEnvironmentVars, 1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(TKGMClusterResourceFullName, "name", TKGMClusterName),
					verifyTanzuKubernetesClusterResourceCreation(provider, TKGMClusterResourceFullName, tkgmEnvironmentVars[TKGMManagementClusterNameEnv],
						tkgmEnvironmentVars[TKGMProvisionerNameEnv], TKGMClusterName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetTKGMClusterConfig(tkgmEnvironmentVars, 2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(TKGMClusterResourceFullName, "name", TKGMClusterName),
					verifyTanzuKubernetesClusterResourceCreation(provider, TKGMClusterResourceFullName, tkgmEnvironmentVars[TKGMManagementClusterNameEnv],
						tkgmEnvironmentVars[TKGMProvisionerNameEnv], TKGMClusterName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetTKGSClusterConfig(tkgsEnvironmentVars, 1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(TKGSClusterResourceFullName, "name", TKGSClusterName),
					verifyTanzuKubernetesClusterResourceCreation(provider, TKGSClusterResourceFullName, tkgsEnvironmentVars[TKGSManagementClusterNameEnv],
						tkgsEnvironmentVars[TKGSProvisionerNameEnv], TKGSClusterName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetTKGSClusterConfig(tkgsEnvironmentVars, 2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(TKGSClusterResourceFullName, "name", TKGSClusterName),
					verifyTanzuKubernetesClusterResourceCreation(provider, TKGSClusterResourceFullName, tkgsEnvironmentVars[TKGSManagementClusterNameEnv],
						tkgsEnvironmentVars[TKGSProvisionerNameEnv], TKGSClusterName),
				),
			},
		},
	},
	)

	t.Log("Tanzu kubernetes cluster resource acceptance test complete!")
}

func verifyTanzuKubernetesClusterResourceCreation(
	provider *schema.Provider,
	resourceName string,
	managementClusterName string,
	provisionerName string,
	clusterName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("could not find resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource %s", resourceName)
		}

		fn := &tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
			ManagementClusterName: managementClusterName,
			ProvisionerName:       provisionerName,
			Name:                  clusterName,
		}

		resp, err := context.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterResourceServiceGet(fn)

		if err != nil {
			return fmt.Errorf("TKG resource not found, resource: %s | err: %s", resourceName, err)
		}

		if resp == nil {
			return fmt.Errorf("TKG resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
