//go:build helmrelease
// +build helmrelease

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmrelease

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
	helmreleaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
	helmreleaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	gitrepo "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository"
	releasescope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                initTestProvider(t),
		HelmReleaseResource:     helmReleaseResource,
		HelmReleaseResourceVar:  helmReleaseResourceVar,
		HelmReleaseResourceName: fmt.Sprintf("%s.%s", helmReleaseResource, helmReleaseResourceVar),
		HelmReleaseName:         acctest.RandomWithPrefix(helmReleaseNamePrefix),
		ScopeHelperResources:    commonscope.NewScopeHelperResources(commonscope.WithRandomClusterGroupNameForCluster()), // Don't use the default cluster group
		Namespace:               "tanzu-helm-resources",
		HelmFeatureResource:     helmfeatureResource,
		HelmFeatureResourceVar:  helmfeatureResourceVar,
	}
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_HELMRELEASE_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForHelmReleaseResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute helm release tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_HELMRELEASE_ENV_TEST")
	if !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.tanzu.broadcom.com")
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

	t.Log("start helm release resource acceptance tests!")

	// Test case for helm release resource.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "registry.terraform.io/hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" && found {
						t.Skip("KUBECONFIG env var is not set for cluster scoped helm release acceptance test")
					}
				},
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterScope, "15.0.3"),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder for cluster scope...")
						testConfig.setupHTTPMocksUpdate(t, commonscope.ClusterScope)
					}
				},
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterScope, constraints),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterGroupScope, "manifests"),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder for cluster group scope...")
						testConfig.setupHTTPMocksUpdate(t, commonscope.ClusterGroupScope)
					}
				},
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterGroupScope, "constraints"),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterGroupScope),
			},
		},
	},
	)

	t.Log("helm release resource acceptance test complete")
}

func (testConfig *testAcceptanceConfig) getTestResourceBasicConfigValue(scope commonscope.Scope, constraints string) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, releasescope.ScopesAllowed[:])

	if _, found := os.LookupEnv("ENABLE_HELMRELEASE_ENV_TEST"); !found {
		switch scope {
		case commonscope.ClusterScope:
			return fmt.Sprintf(`
			resource "%s" "%s" {
			 name = "%s"

			 namespace_name = "%s"

			 scope {
				cluster {
					name = "%s"
					management_cluster_name = "attached"
					provisioner_name = "attached"
				}
			 }

			 spec {
				chart_ref {
					helm_repository{
						repository_name = "bitnami"
						repository_namespace = "tanzu-helm-resources"
						chart_name = "airflow"
						version = "%s"
					}
				}
				interval = "5m"
			}
			}
			`, testConfig.HelmReleaseResource, testConfig.HelmReleaseResourceVar, testConfig.HelmReleaseName, testConfig.Namespace, testConfig.ScopeHelperResources.Cluster.Name, constraints)
		case commonscope.ClusterGroupScope:
			return fmt.Sprintf(`
			resource "%s" "%s" {
			 name = "%s"

			 namespace_name = "%s"

			 scope {
				cluster_group {
					name = "%s"
				}
			 }

			 spec {
				chart_ref {
					git_repository{
						repository_name = "test-git-repo"
						repository_namespace = "tanzu-continuousdelivery-resources"
						chart_path = "%s"
					}
				}
				interval = "5m"
			}
			}
			`, testConfig.HelmReleaseResource, testConfig.HelmReleaseResourceVar, testConfig.HelmReleaseName, testConfig.Namespace, testConfig.ScopeHelperResources.ClusterGroup.Name, constraints)
		default:
			return ""
		}
	}

	switch scope {
	case commonscope.ClusterScope:
		return fmt.Sprintf(`
	%s

	resource "time_sleep" "wait_for_3m" {
		duration = "50s"

		depends_on = [tanzu-mission-control_cluster.test_cluster]
	}

	resource "%s" "%s" {

		%s

		depends_on = [time_sleep.wait_for_3m]
	}

	resource "time_sleep" "wait_for_2m" {
		duration = "300s"

		depends_on = [tanzu-mission-control_helm_feature.test_helm_feature]
	}

	resource "%s" "%s" {
		name = "%s"

		namespace_name = "%s"

		%s

		spec {
			chart_ref {
				helm_repository{
					repository_name = "bitnami"
					repository_namespace = "tanzu-helm-resources"
					chart_name = "airflow"
					version = "%s"
				}
			}
		}

		depends_on = [time_sleep.wait_for_2m]
	}
	`, helperBlock, testConfig.HelmFeatureResource, testConfig.HelmFeatureResourceVar, scopeBlock,
			testConfig.HelmReleaseResource, testConfig.HelmReleaseResourceVar, testConfig.HelmReleaseName,
			testConfig.Namespace, scopeBlock, constraints)

	case commonscope.ClusterGroupScope:
		return fmt.Sprintf(`
	%s

	resource "time_sleep" "wait_for_1m" {
		duration = "60s"

		depends_on = [tanzu-mission-control_cluster_group.test_cluster_group]
	}

	resource "%s" "%s" {

		%s

		depends_on = [time_sleep.wait_for_1m]
	}

	resource "%s" "%s" {
		name = "%s"

		namespace_name = "tanzu-continuousdelivery-resources"

		%s

		spec {
			url = "https://github.com/tmc-build-integrations/sample-update-configmap"
			ref {
				branch = "main"
			}
		}
	}

	resource "time_sleep" "wait_for_2m" {
		duration = "180s"

		depends_on = [tanzu-mission-control_helm_feature.test_helm_feature, tanzu-mission-control_git_repository.test_git_repo]
	}

	resource "%s" "%s" {
		name = "%s"

		namespace_name = "%s"

		%s

		spec {
			chart_ref {
				git_repository{
					repository_name = tanzu-mission-control_git_repository.test_git_repo.name
					repository_namespace = tanzu-mission-control_git_repository.test_git_repo.namespace_name
					chart_path = "%s"
				}
			}
		}

		depends_on = [time_sleep.wait_for_2m]
	}
	`, helperBlock, testConfig.HelmFeatureResource, testConfig.HelmFeatureResourceVar, scopeBlock, gitrepo.ResourceName, gitRepoResourceVar, gitRepoName, scopeBlock,
			testConfig.HelmReleaseResource, testConfig.HelmReleaseResourceVar, testConfig.HelmReleaseName,
			testConfig.Namespace, scopeBlock, constraints)
	default:
		return ""
	}
}

// checkResourceAttributes checks for git repository creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkResourceAttributes(scopeType commonscope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.HelmReleaseResourceName, "name", testConfig.HelmReleaseName),
	}

	switch scopeType {
	case commonscope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.HelmReleaseResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case commonscope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.HelmReleaseResourceName, "scope.0.cluster_group.0.name", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case commonscope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(releasescope.ScopesAllowed[:], `, `))
	}

	check = append(check, commonscope.MetaResourceAttributeCheck(testConfig.HelmReleaseResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyResourceCreation(scopeType commonscope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.HelmReleaseResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.HelmReleaseResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.HelmReleaseResourceName)
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
			fn := &helmreleaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName{
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ManagementClusterName: commonscope.AttachedValue,
				Name:                  testConfig.HelmReleaseName,
				NamespaceName:         testConfig.Namespace,
				ProvisionerName:       commonscope.AttachedValue,
			}

			resp, err := config.TMCConnection.ClusterHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped git repository resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped git repository resource is empty, resource: %s", testConfig.HelmReleaseResourceName)
			}
		case commonscope.ClusterGroupScope:
			fn := &helmreleaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName{
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				Name:             testConfig.HelmReleaseName,
				NamespaceName:    testConfig.Namespace,
			}

			resp, err := config.TMCConnection.ClusterGroupHelmReleaseResourceService.VmwareTanzuManageV1alpha1ClustergroupReleaseResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped git repository resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped git repository resource is empty, resource: %s", testConfig.HelmReleaseResourceName)
			}
		case commonscope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(releasescope.ScopesAllowed[:], `, `))
		}

		return nil
	}
}
