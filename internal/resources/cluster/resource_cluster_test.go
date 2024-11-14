//go:build cluster
// +build cluster

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

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
				SkipFunc: func() (bool, error) {
					if os.Getenv("ATTACH_KUBECONFIG") == "" {
						t.Log("ATTACH_KUBECONFIG env var is not set")
						return true, nil
					}
					return false, nil
				},
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["attachWithKubeConfigImageRegistry"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["attachWithKubeConfigImageRegistry"]...),
				),
				SkipFunc: func() (bool, error) {
					if os.Getenv("ATTACH_WITH_IMAGE_REGISTRY_KUBECONFIG") == "" || os.Getenv("IMAGE_REGISTRY") == "" {
						t.Log("ATTACH_WITH_IMAGE_REGISTRY_KUBECONFIG or IMAGE_REGISTRY env var is not set")
						return true, nil
					}
					return false, nil
				},
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["attachWithKubeConfigProxy"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["attachWithKubeConfigProxy"]...),
				),
				SkipFunc: func() (bool, error) {
					if os.Getenv("ATTACH_WITH_PROXY_KUBECONFIG") == "" || os.Getenv("PROXY") == "" {
						t.Log("ATTACH_WITH_PROXY_KUBECONFIG or PROXY env var is not set")
						return true, nil
					}
					return false, nil
				},
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["tkgAWS"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["tkgAWS"]...),
				),
				SkipFunc: func() (bool, error) {
					if os.Getenv("TKGM_AWS_MANAGEMENT_CLUSTER") == "" || os.Getenv("TKGM_AWS_PROVISIONER_NAME") == "" {
						t.Log("TKGM_AWS_MANAGEMENT_CLUSTER or TKGM_AWS_PROVISIONER_NAME env var is not set for TKGm AWS acceptance test")
						return true, nil
					}
					return false, nil
				},
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["tkgs"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["tkgs"]...),
				),
				SkipFunc: func() (bool, error) {
					if os.Getenv("TKGS_MANAGEMENT_CLUSTER") == "" || os.Getenv("TKGS_PROVISIONER_NAME") == "" ||
						os.Getenv("VERSION") == "" || os.Getenv("STORAGE_CLASS") == "" {
						t.Log("TKGS_MANAGEMENT_CLUSTER, TKGS_PROVISIONER_NAME, VERSION or STORAGE CLASS env var is not set for TKGs acceptance test")
						return true, nil
					}
					return false, nil
				},
			},
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["tkgVsphere"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["tkgVsphere"]...),
				),
				ExpectNonEmptyPlan: true,
				SkipFunc: func() (bool, error) {
					if os.Getenv("TKGM_VSPHERE_MANAGEMENT_CLUSTER") == "" || os.Getenv("TKGM_VSPHERE_PROVISIONER_NAME") == "" {
						t.Log("TKGM_VSPHERE_MANAGEMENT_CLUSTER or TKGM_VSPHERE_PROVISIONER_NAME env var is not set for TKGm Vsphere acceptance test")
						return true, nil
					}
					return false, nil
				},
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
