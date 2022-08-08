/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
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

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	clusterResource      = "tanzu-mission-control_cluster"
	clusterResourceVar   = "test_attach_cluster"
	clusterDataSourceVar = "test_data_attach_cluster"
)

var (
	resourceName   = fmt.Sprintf("%s.%s", clusterResource, clusterResourceVar)
	dataSourceName = fmt.Sprintf("data.%s.%s", clusterResource, clusterDataSourceVar)
)

func TestAcceptanceForAttachClusterResource(t *testing.T) {
	var provider = initTestProvider(t)

	clusterConfig := map[string][]testAcceptanceOption{
		"attach":               {withClusterName("tf-attach-test")},
		"attachWithKubeConfig": {withKubeConfig(), withClusterName("tf-attach-kf-test")},
		"tkgAWS":               {withClusterName("tf-tkgm-aws-test"), withTKGmAWSCluster()},
		"tkgs":                 {withClusterName("tf-tkgs-test"), withTKGsCluster()},
		"tkgVsphere":           {withClusterName("tf-tkgm-vsphere-test"), withTKGmVsphereCluster()},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["attach"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["attach"]...),
				),
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["attachWithKubeConfig"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["attachWithKubeConfig"]...),
				),
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["tkgAWS"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["tkgAWS"]...),
				),
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["tkgs"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["tkgs"]...),
				),
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["tkgVsphere"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["tkgVsphere"]...),
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

	switch templateConfig.accTestType {
	case attachClusterTypeWithKubeConfig:
		if templateConfig.KubeConfigPath == "" {
			t.Skipf("KUBECONFIG env var is not set: %s", templateConfig.KubeConfigPath)
		}

	case tkgAWSCluster:
		if templateConfig.ManagementClusterName == "" || templateConfig.ProvisionerName == "" {
			t.Skip("MANAGEMENT CLUSTER or PROVISIONER env var is not set for TKGm AWS acceptance test")
		}

	case tkgVsphereCluster:
		if templateConfig.ManagementClusterName == "" || templateConfig.ProvisionerName == "" || templateConfig.ControlPlaneEndPoint == "" {
			t.Skip("MANAGEMENT CLUSTER, PROVISIONER or CONTROL PLANE ENDPOINT env var is not set for TKGm Vsphere acceptance test")
		}

	case tkgsCluster:
		if templateConfig.ManagementClusterName == "" || templateConfig.ProvisionerName == "" || templateConfig.Version == "" || templateConfig.StorageClass == "" {
			t.Skip("MANAGEMENT CLUSTER, PROVISIONER, VERSION or STORAGE CLASS env var is not set for TKGs acceptance test")
		}
	}

	definition, err := parse(templateConfig, templateConfig.templateData)
	if err != nil {
		t.Skipf("unable to parse cluster script: %s", err)
	}

	return definition
}

func checkResourceAttributes(provider *schema.Provider, opts ...testAcceptanceOption) resource.TestCheckFunc {
	testConfig := testGetDefaultAcceptanceConfig()
	for _, option := range opts {
		option(testConfig)
	}

	var check = []resource.TestCheckFunc{
		verifyClusterResourceCreation(provider, resourceName, testConfig),
		resource.TestCheckResourceAttr(resourceName, "name", testConfig.Name),
		resource.TestCheckResourceAttr(resourceName, helper.GetFirstElementOf("spec", "cluster_group"), "default"),
	}

	if testConfig.accTestType == attachClusterType || testConfig.accTestType == attachClusterTypeWithKubeConfig {
		check = append(check, testhelper.MetaResourceAttributeCheck(resourceName)...)
	}

	return resource.ComposeTestCheckFunc(check...)
}

func verifyClusterResourceCreation(
	provider *schema.Provider,
	resourceName string,
	testConfig *testAcceptanceConfig,
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
			ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
			Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
			VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		}

		err := config.Setup()
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		fn := &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
			Name:                  testConfig.Name,
			ManagementClusterName: testConfig.ManagementClusterName,
			ProvisionerName:       testConfig.ProvisionerName,
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
