//go:build helmfeature
// +build helmfeature

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmfeature

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	helmscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmfeature/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                initTestProvider(t),
		ScopeHelperResources:    commonscope.NewScopeHelperResources(commonscope.WithRandomClusterGroupNameForCluster()), // Don't use the default cluster group
		HelmFeatureResource:     helmfeatureResource,
		HelmFeatureResourceVar:  helmfeatureResourceVar,
		HelmReleaseName:         acctest.RandomWithPrefix(helmReleaseNamePrefix),
		HelmFeatureResourceName: fmt.Sprintf("%s.%s", helmfeatureResource, helmfeatureResourceVar),
	}
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_HELM_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForHelmFeatureResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute helm feature tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_HELM_ENV_TEST")
	if !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.cloud.vmware.com")
		os.Setenv("ORG_ID", "bc27608b-4809-4cac-9e04-778803963da2")

		log.Println("Setting up the mock endpoints...")

		testConfig.setupHTTPMocks(t)
	} else {
		// Environment variables with non default values required for a successful call to helm deployment service.
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
			"ORG_ID",
		}

		// Check if the required environment variables are set.
		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	t.Log("start helm feature resource acceptance tests!")

	// Test case for helm feature resource.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		CheckDestroy:      nil,
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "registry.terraform.io/hashicorp/time",
			},
			"null": {
				Source: "registry.terraform.io/hashicorp/null",
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" && found {
						t.Skip("KUBECONFIG env var is not set for cluster scoped helm feature acceptance test")
					}
				},
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterScope),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterGroupScope),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterGroupScope),
			},
		},
	},
	)

	t.Log("helm feature resource acceptance test complete")
}

func (testConfig *testAcceptanceConfig) getTestResourceBasicConfigValue(scope commonscope.Scope) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, helmscope.ScopesAllowed[:])

	if _, found := os.LookupEnv("ENABLE_HELM_ENV_TEST"); !found {
		switch scope {
		case commonscope.ClusterScope:
			return fmt.Sprintf(`
			resource "%s" "%s" {
			
			 scope {
				cluster {
					name = "%s"
					management_cluster_name = "attached"
					provisioner_name = "attached"
				}
			 }

			}
			`, testConfig.HelmFeatureResource, testConfig.HelmFeatureResourceVar, testConfig.ScopeHelperResources.Cluster.Name)
		case commonscope.ClusterGroupScope:
			return fmt.Sprintf(`
			resource "%s" "%s" {
			
			 scope {
				cluster_group {
					name = "%s"
				}
			 }
			}
			`, testConfig.HelmFeatureResource, testConfig.HelmFeatureResourceVar, testConfig.ScopeHelperResources.ClusterGroup.Name)
		default:
			return ""
		}
	}

	switch scope {
	case commonscope.ClusterScope:
		return fmt.Sprintf(`
		%s

		resource "time_sleep" "wait_for_3m" {
			create_duration = "50s"

			depends_on = [tanzu-mission-control_cluster.test_cluster]
		}

		resource "%s" "%s" {

			%s
		}

		resource "time_sleep" "wait_for_2m" {
			destroy_duration = "300s"
		  
			depends_on = [tanzu-mission-control_helm_feature.test_helm_feature]
		}

		resource "null_resource" "next" {
			depends_on = [time_sleep.wait_for_2m]
		}
		`, helperBlock, testConfig.HelmFeatureResource, testConfig.HelmFeatureResourceVar, scopeBlock)
	case commonscope.ClusterGroupScope:
		return fmt.Sprintf(`
		%s

		resource "time_sleep" "wait_for_3m" {
			create_duration = "180s"

			depends_on = [tanzu-mission-control_cluster_group.test_cluster_group]
		}

		resource "%s" "%s" {

			%s

			depends_on = [time_sleep.wait_for_3m]
		}
		`, helperBlock, testConfig.HelmFeatureResource, testConfig.HelmFeatureResourceVar, scopeBlock)
	default:
		return ""
	}
}

// checkResourceAttributes checks for git repository creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkResourceAttributes(scopeType commonscope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyResourceCreation(scopeType),
	}

	switch scopeType {
	case commonscope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.HelmFeatureResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case commonscope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.HelmFeatureResourceName, "scope.0.cluster_group.0.name", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case commonscope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(helmscope.ScopesAllowed[:], `, `))
	}

	check = append(check, commonscope.MetaResourceAttributeCheck(testConfig.HelmFeatureResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyResourceCreation(scopeType commonscope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.HelmFeatureResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.HelmFeatureResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.HelmFeatureResourceName)
		}

		config := authctx.TanzuContext{
			ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
			Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
			VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
			TLSConfig:        &proxy.TLSConfig{},
		}

		err := getSetupConfig(&config)
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		switch scopeType {
		case commonscope.ClusterScope:
			fn := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmRequestParameters{
				SearchScope: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmSearchScope{
					ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
					ManagementClusterName: commonscope.AttachedValue,
					ProvisionerName:       commonscope.AttachedValue,
				},
			}

			resp, err := config.TMCConnection.ClusterHelmResourceService.VmwareTanzuManageV1alpha1ClusterHelmResourceServiceList(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped git repository resource not found")
			}

			if len(resp.Helms) == 0 {
				return errors.Errorf("Tanzu mission control helm feature is disable on cluster, name: %s", testConfig.ScopeHelperResources.Cluster.Name)
			}

			if _, ok := resp.Helms[0].Status.Conditions[disabledKey]; ok {
				return errors.Errorf("Tanzu mission control helm feature is disable on cluster, name: %s", testConfig.ScopeHelperResources.Cluster.Name)
			}
		case commonscope.ClusterGroupScope:
			fn := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupHelmListHelmRequestParameters{
				SearchScope: &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmSearchScope{
					ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				},
			}

			resp, err := config.TMCConnection.ClusterGroupHelmResourceService.VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceList(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped git repository resource not found")
			}

			if len(resp.Helms) == 0 {
				return errors.Errorf("Tanzu mission control helm feature is disable on cluster group, name: %s", testConfig.ScopeHelperResources.ClusterGroup.Name)
			}

			if resp.Helms[0].Status.Phase == nil {
				return errors.Errorf("Tanzu mission control helm feature is disable on cluster group, name: %s", testConfig.ScopeHelperResources.ClusterGroup.Name)
			}
		case commonscope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(helmscope.ScopesAllowed[:], `, `))
		}

		return nil
	}
}
