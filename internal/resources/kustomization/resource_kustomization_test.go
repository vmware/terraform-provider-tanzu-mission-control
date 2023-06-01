/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	gitrepositoryhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository"
	kustomizationscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

// nolint: gosec
const (
	kustomizationResource      = ResourceName
	kustomizationResourceVar   = "test_kustomizationy"
	kustomizationDataSourceVar = "test_data_source_kustomizationy"
	kustomizationNamePrefix    = "tf-kustomization-test"
	namespaceName              = "tanzu-continuousdelivery-resources"

	gitRepositoryResource    = gitrepositoryhelper.ResourceName
	gitRepositoryResourceVar = "test_git_repository"
	gitRepositoryNamePrefix  = "tf-gr-test"
)

type testAcceptanceConfig struct {
	Provider                  *schema.Provider
	KustomizationResource     string
	KustomizationResourceVar  string
	KustomizationResourceName string
	KustomizationName         string
	ScopeHelperResources      *commonscope.ScopeHelperResources
	GitRepositoryResource     string
	GitRepositoryResourceVar  string
	GitRepositoryName         string
}

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
	}
}

func TestAcceptanceForKustomizationResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start kustomization resource acceptance tests!")

	// Test case for kustomization resource.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped kustomization acceptance test")
					}
				},
				Config: testConfig.getTestKustomizationResourceBasicConfigValue(commonscope.ClusterScope),
				Check:  testConfig.checkKustomizationResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestKustomizationResourceBasicConfigValue(commonscope.ClusterScope, WithPruneEnabled(true), WithInterval("10m")),
				Check:  testConfig.checkKustomizationResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestKustomizationResourceBasicConfigValue(commonscope.ClusterGroupScope),
				Check:  testConfig.checkKustomizationResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
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
	`, helperBlock, testConfig.GitRepositoryResource, testConfig.GitRepositoryResourceVar, testConfig.GitRepositoryName, namespaceName, scopeBlock, testConfig.KustomizationResource, testConfig.KustomizationResourceVar, testConfig.KustomizationName, namespaceName, scopeBlock, kustomizationSpec)
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

		err := config.Setup()
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
