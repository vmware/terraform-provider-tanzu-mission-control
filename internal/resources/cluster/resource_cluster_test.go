/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster"
	testhelper "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	clusterResource      = "tmc_cluster"
	clusterResourceVar   = "test_attach_cluster"
	clusterDataSourceVar = "test_data_attach_cluster"
)

var (
	resourceName   = fmt.Sprintf("%s.%s", clusterResource, clusterResourceVar)
	dataSourceName = fmt.Sprintf("data.%s.%s", clusterResource, clusterDataSourceVar)
)

func TestAcceptanceForAttachClusterResource(t *testing.T) {
	var provider = initTestProvider(t)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testGetResourceClusterDefinition(t, withClusterName("tf-attach-test")),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, "tf-attach-test"),
				),
			},
			{
				Config: testGetResourceClusterDefinition(t, withClusterName("tf-attach-kubeconfig-test"), withKubeConfig()),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, "tf-attach-kubeconfig-test"),
				),
			},
		},
	})
	t.Log("cluster resource acceptance test complete!")
}

func testGetResourceClusterDefinition(t *testing.T, opts ...testAcceptanceOption) string {
	templateConfig := testGetDefaultAcceptanceConfig()
	for _, option := range opts {
		option(templateConfig)
	}

	if templateConfig.accTestType == attachClusterTypeWithKubeConfig {
		if templateConfig.KubeConfigPath == "" {
			t.Skipf("KUBECONFIG env var is not set: %s", templateConfig.KubeConfigPath)
		}
	}

	definition, err := parse(templateConfig, templateConfig.templateData)
	if err != nil {
		t.Skipf("unable to parse cluster script: %s", definition)
	}

	return definition
}

func checkResourceAttributes(provider *schema.Provider, clusterName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyClusterResourceCreation(provider, resourceName, clusterName),
		resource.TestCheckResourceAttr(resourceName, "name", clusterName),
		resource.TestCheckResourceAttr(resourceName, helper.GetFirstElementOf("spec", "cluster_group"), "default"),
	}

	check = append(check, testhelper.MetaResourceAttributeCheck(resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifyClusterResourceCreation(
	provider *schema.Provider,
	resourceName string,
	clusterName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return errors.New("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return errors.Errorf("not found resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return errors.Errorf("ID not set, resource %s", resourceName)
		}

		config := authctx.TanzuContext{
			ServerEndpoint: os.Getenv(authctx.ServerEndpointEnvVar),
			Token:          os.Getenv(authctx.CSPTokenEnvVar),
			CSPEndPoint:    os.Getenv(authctx.CSPEndpointEnvVar),
		}

		err := config.Setup()
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		fn := &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
			Name:                  clusterName,
			ManagementClusterName: "attached",
			ProvisionerName:       "attached",
		}

		resp, err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(fn)
		if err != nil {
			return errors.Errorf("cluster resource not found: %s", err)
		}

		if resp == nil {
			return errors.Errorf("cluster resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
