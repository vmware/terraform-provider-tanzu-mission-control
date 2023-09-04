//go:build cluster
// +build cluster

/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForAttachClusterResource(t *testing.T) {
	var provider = initTestProvider(t)

	clusterConfig := map[string][]testhelper.TestAcceptanceOption{
		"attach":                            {testhelper.WithClusterName("tf-attach-test")},
		"attachWithKubeConfig":              {testhelper.WithKubeConfig(), testhelper.WithClusterName("tf-attach-kf-test")},
		"attachWithKubeConfigImageRegistry": {testhelper.WithClusterName("tf-attach-img-reg"), testhelper.WithKubeConfigImageRegistry()},
		"attachWithKubeConfigProxy":         {testhelper.WithClusterName("tf-attach-proxy"), testhelper.WithKubeConfigProxy()},
		"tkgAWS":                            {testhelper.WithClusterName("tf-tkgm-aws-test"), testhelper.WithTKGmAWSCluster()},
		"tkgs":                              {testhelper.WithClusterName("tf-tkgs-test"), testhelper.WithTKGsCluster()},
		"tkgVsphere":                        {testhelper.WithClusterName("tf-tkgm-vsphere-test"), testhelper.WithTKGmVsphereCluster()},
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
				Config: testGetResourceClusterDefinition(t, clusterConfig["attachWithKubeConfigImageRegistry"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["attachWithKubeConfigImageRegistry"]...),
				),
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["attachWithKubeConfigProxy"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["attachWithKubeConfigProxy"]...),
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

func testGetResourceClusterDefinition(t *testing.T, opts ...testhelper.TestAcceptanceOption) string {
	templateConfig := testhelper.TestGetDefaultAcceptanceConfig()
	for _, option := range opts {
		option(templateConfig)
	}

	switch templateConfig.AccTestType {
	case testhelper.AttachClusterTypeWithKubeConfig:
		if templateConfig.KubeConfigPath == "" {
			t.Skipf("KUBECONFIG env var is not set: %s", templateConfig.KubeConfigPath)
		}

	case testhelper.AttachClusterTypeWithKubeconfigImageRegistry:
		if templateConfig.KubeConfigPath == "" || templateConfig.ImageRegistry == "" {
			t.Skipf("KUBECONFIG or IMAGE_REGISTRY env var is not set")
		}

	case testhelper.AttachClusterTypeWithKubeconfigProxy:
		if templateConfig.KubeConfigPath == "" || templateConfig.Proxy == "" {
			t.Skipf("KUBECONFIG or PROXY env var is not set")
		}

	case testhelper.TkgAWSCluster:
		if templateConfig.ManagementClusterName == "" || templateConfig.ProvisionerName == "" {
			t.Skip("MANAGEMENT CLUSTER or PROVISIONER env var is not set for TKGm AWS acceptance test")
		}

	case testhelper.TkgVsphereCluster:
		if templateConfig.ManagementClusterName == "" || templateConfig.ProvisionerName == "" || templateConfig.ControlPlaneEndPoint == "" {
			t.Skip("MANAGEMENT CLUSTER, PROVISIONER or CONTROL PLANE ENDPOINT env var is not set for TKGm Vsphere acceptance test")
		}

	case testhelper.TkgsCluster:
		if templateConfig.ManagementClusterName == "" || templateConfig.ProvisionerName == "" || templateConfig.Version == "" || templateConfig.StorageClass == "" {
			t.Skip("MANAGEMENT CLUSTER, PROVISIONER, VERSION or STORAGE CLASS env var is not set for TKGs acceptance test")
		}
	}

	definition, err := testhelper.Parse(templateConfig, templateConfig.TemplateData)
	if err != nil {
		t.Skipf("unable to parse cluster script: %s", err)
	}

	return definition
}

func checkResourceAttributes(provider *schema.Provider, opts ...testhelper.TestAcceptanceOption) resource.TestCheckFunc {
	testConfig := testhelper.TestGetDefaultAcceptanceConfig()
	for _, option := range opts {
		option(testConfig)
	}

	var check = []resource.TestCheckFunc{
		verifyClusterResourceCreation(provider, testhelper.ClusterResourceName, testConfig),
		resource.TestCheckResourceAttr(testhelper.ClusterResourceName, "name", testConfig.Name),
		resource.TestCheckResourceAttr(testhelper.ClusterResourceName, helper.GetFirstElementOf("spec", "cluster_group"), "default"),
	}

	if testConfig.AccTestType == testhelper.AttachClusterType || testConfig.AccTestType == testhelper.AttachClusterTypeWithKubeConfig ||
		testConfig.AccTestType == testhelper.AttachClusterTypeWithKubeconfigProxy || testConfig.AccTestType == testhelper.AttachClusterTypeWithKubeconfigImageRegistry {
		check = append(check, testhelper.MetaResourceAttributeCheck(testhelper.ClusterResourceName)...)
	}

	return resource.ComposeTestCheckFunc(check...)
}

func verifyClusterResourceCreation(
	provider *schema.Provider,
	resourceName string,
	testConfig *testhelper.TestAcceptanceConfig,
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
			TLSConfig:        &proxy.TLSConfig{},
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
