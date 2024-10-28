//go:build kustomization
// +build kustomization

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kustomization

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
	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	kustomizationscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                  initTestProvider(t),
		KustomizationResource:     kustomizationResource,
		KustomizationResourceVar:  kustomizationResourceVar,
		KustomizationResourceName: fmt.Sprintf("%s.%s", kustomizationResource, kustomizationResourceVar),
		KustomizationName:         acctest.RandomWithPrefix(kustomizationNamePrefix),
		ScopeHelperResources:      commonscope.NewScopeHelperResources(),
		GitRepositoryResource:     gitRepositoryResource,
		GitRepositoryResourceVar:  gitRepositoryResourceVar,
		GitRepositoryName:         acctest.RandomWithPrefix(gitRepositoryNamePrefix),
		Namespace:                 "tanzu-continuousdelivery-resources",
	}
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_KUSTOMIZATION_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForKustomizationResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute kustomization tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_KUSTOMIZATION_ENV_TEST")
	if !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.tanzu.broadcom.com")

		log.Println("Setting up the mock endpoints...")

		testConfig.setupHTTPMocks(t)
	} else {
		// Environment variables with non default values required for a successful call to Cluster Config Service.
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

	t.Log("start kustomization resource acceptance tests!")

	// Test case for kustomization resource.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" && found {
						t.Skip("KUBECONFIG env var is not set for cluster scoped kustomization acceptance test")
					}
				},
				Config: testConfig.getTestKustomizationResourceBasicConfigValue(commonscope.ClusterScope),
				Check:  testConfig.checkKustomizationResourceAttributes(commonscope.ClusterScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder for cluster scope...")
						testConfig.setupHTTPMocksUpdate(t, commonscope.ClusterScope)
					}
				},
				Config: testConfig.getTestKustomizationResourceBasicConfigValue(commonscope.ClusterScope, WithPruneEnabled(true), WithInterval("10m")),
				Check:  testConfig.checkKustomizationResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestKustomizationResourceBasicConfigValue(commonscope.ClusterGroupScope),
				Check:  testConfig.checkKustomizationResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder for cluster group scope...")
						testConfig.setupHTTPMocksUpdate(t, commonscope.ClusterGroupScope)
					}
				},
				Config: testConfig.getTestKustomizationResourceBasicConfigValue(commonscope.ClusterGroupScope, WithPruneEnabled(true), WithInterval("10m")),
				Check:  testConfig.checkKustomizationResourceAttributes(commonscope.ClusterGroupScope),
			},
		},
	},
	)

	t.Log("kustomization resource acceptance test complete")
}

func (testConfig *testAcceptanceConfig) getTestKustomizationResourceBasicConfigValue(scope commonscope.Scope, opts ...OperationOption) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, kustomizationscope.ScopesAllowed[:])
	kustomizationSpec := getKustomizationspec(opts...)

	if _, found := os.LookupEnv("ENABLE_KUSTOMIZATION_ENV_TEST"); !found {
		clStr := fmt.Sprintf(`
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

		%s
	}
	`, testConfig.KustomizationResource, testConfig.KustomizationResourceVar, testConfig.KustomizationName, testConfig.Namespace, testConfig.ScopeHelperResources.Cluster.Name, kustomizationSpec)

		cgStr := fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"

		namespace_name = "%s"

		 scope {
		cluster_group {
			name = "%s"
		}
	 }

		%s
	}
	`, testConfig.KustomizationResource, testConfig.KustomizationResourceVar, testConfig.KustomizationName, testConfig.Namespace, testConfig.ScopeHelperResources.ClusterGroup.Name, kustomizationSpec)

		switch scope {
		case commonscope.ClusterScope:
			return clStr
		case commonscope.ClusterGroupScope:
			return cgStr
		}
	}

	return fmt.Sprintf(`
	%s
	
	resource "%s" "%s" {
	 name = "%s"

	 namespace_name = "%s"
	
	 %s
	
	 spec {
		url = "https://github.com/tmc-build-integrations/sample-update-configmap"
		ref {
			branch = "master"
		}
	 }
	}

	resource "%s" "%s" {
		name = "%s"

		namespace_name = "%s"

		%s

		%s
	}
	`, helperBlock, testConfig.GitRepositoryResource, testConfig.GitRepositoryResourceVar, testConfig.GitRepositoryName, testConfig.Namespace, scopeBlock, testConfig.KustomizationResource, testConfig.KustomizationResourceVar, testConfig.KustomizationName, testConfig.Namespace, scopeBlock, kustomizationSpec)
}

type (
	OperationConfig struct {
		pruneEnabled bool
		interval     string
	}

	OperationOption func(*OperationConfig)
)

func WithPruneEnabled(val bool) OperationOption {
	return func(config *OperationConfig) {
		config.pruneEnabled = val
	}
}

func WithInterval(val string) OperationOption {
	return func(config *OperationConfig) {
		config.interval = val
	}
}

func getKustomizationspec(opts ...OperationOption) string {
	cfg := &OperationConfig{
		pruneEnabled: false,
		interval:     "5m",
	}

	for _, o := range opts {
		o(cfg)
	}

	if _, found := os.LookupEnv("ENABLE_KUSTOMIZATION_ENV_TEST"); !found {
		inputBlock := `
    spec {
			path = "manifests/"
			prune = %t
			interval = "%s"
			source {
				name = "someGitRepository"
				namespace = "tanzu-continuousdelivery-resources"
			}
		}
`
		inputBlock = fmt.Sprintf(inputBlock, cfg.pruneEnabled, cfg.interval)

		return inputBlock
	}

	inputBlock := `
    spec {
			path = "manifests/"
			prune = %t
			interval = "%s"
			source {
				name = tanzu-mission-control_git_repository.test_git_repository.name
				namespace = tanzu-mission-control_git_repository.test_git_repository.namespace_name
			}
		}
`
	inputBlock = fmt.Sprintf(inputBlock, cfg.pruneEnabled, cfg.interval)

	return inputBlock
}

// checkKustomizationResourceAttributes checks for kustomization creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkKustomizationResourceAttributes(scopeType commonscope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyKustomizationResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.KustomizationResourceName, "name", testConfig.KustomizationName),
	}

	switch scopeType {
	case commonscope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.KustomizationResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case commonscope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.KustomizationResourceName, "scope.0.cluster_group.0.name", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case commonscope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(kustomizationscope.ScopesAllowed[:], `, `))
	}

	check = append(check, commonscope.MetaResourceAttributeCheck(testConfig.KustomizationResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyKustomizationResourceCreation(scopeType commonscope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.KustomizationResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.KustomizationResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.KustomizationResourceName)
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
			fn := &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName{
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ManagementClusterName: commonscope.AttachedValue,
				Name:                  testConfig.KustomizationName,
				NamespaceName:         "tanzu-continuousdelivery-resources",
				ProvisionerName:       commonscope.AttachedValue,
			}

			resp, err := config.TMCConnection.ClusterKustomizationResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdKustomizationResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped kustomization resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped kustomization resource is empty, resource: %s", testConfig.KustomizationResourceName)
			}
		case commonscope.ClusterGroupScope:
			fn := &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName{
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				Name:             testConfig.KustomizationName,
				NamespaceName:    "tanzu-continuousdelivery-resources",
			}

			resp, err := config.TMCConnection.ClusterGroupKustomizationResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdKustomizationResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped kustomization resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped kustomization resource is empty, resource: %s", testConfig.KustomizationResourceName)
			}
		case commonscope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(kustomizationscope.ScopesAllowed[:], `, `))
		}

		return nil
	}
}
