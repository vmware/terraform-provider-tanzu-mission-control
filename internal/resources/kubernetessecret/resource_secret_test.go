//go:build clustersecret
// +build clustersecret

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

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
	secretclmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	secretcgmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	secretscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:             initTestProvider(t),
		SecretResource:       ResourceName,
		SecretResourceVar:    secretResourceVar,
		SecretResourceName:   fmt.Sprintf("%s.%s", ResourceName, secretResourceVar),
		SecretName:           acctest.RandomWithPrefix("tf-sc-test"),
		ScopeHelperResources: commonscope.NewScopeHelperResources(),
		NamespaceName:        "default",
		DataSourceName:       fmt.Sprintf("data.%s.%s", ResourceName, secretDataSourceVar),
		DataSourceVar:        secretDataSourceVar,
	}
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_SECRET_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForSecretResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute kubernetes secret tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_SECRET_ENV_TEST")
	if !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.cloud.vmware.com")

		log.Println("Setting up the mock endpoints...")
	} else {
		// Environment variables with non default values required for a successful call to MKP
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
			"ORG_ID",
		}

		// Check if the required environment variables are set
		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	t.Log("start cluster secret resource acceptance tests!", testConfig.ScopeHelperResources.Cluster.Name)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if !found {
						testConfig.setupHTTPMocks(t, DockerSecretType)
					}
				},
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterGroupScope, WithSecretType(DockerSecretType)),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if !found {
						testConfig.setupHTTPMocks(t, DockerSecretType)
					}
				},
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterScope, WithSecretType(DockerSecretType)),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterScope),
			},
			{
				PreConfig: func() {
					if !found {
						testConfig.setupHTTPMocks(t, OpaqueSecretType)
					}
				},
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterGroupScope, WithSecretType(OpaqueSecretType)),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if !found {
						testConfig.setupHTTPMocks(t, OpaqueSecretType)
					}
				},
				Config: testConfig.getTestResourceBasicConfigValue(commonscope.ClusterScope, WithSecretType(OpaqueSecretType)),
				Check:  testConfig.checkResourceAttributes(commonscope.ClusterScope),
			},
		},
	},
	)
	t.Log("secret resource acceptance test complete!")
}

func (testConfig *testAcceptanceConfig) getTestResourceBasicConfigValue(scope commonscope.Scope, opts ...OperationOption) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, secretscope.ScopesAllowed[:])
	secretSpec := testConfig.getTestSecretResourceSpec(opts...)

	if _, found := os.LookupEnv("ENABLE_SECRET_ENV_TEST"); !found {
		clStr := fmt.Sprintf(`
	resource "%s" "%s" {
	 name = "%s"

	 namespace_name = "default"
	
	 scope {
		cluster {
			name = "%s"
			management_cluster_name = "attached"
			provisioner_name = "attached"
		}
	 }

	%s
	}
	`, testConfig.SecretResource, testConfig.SecretResourceVar, testConfig.SecretName, testConfig.ScopeHelperResources.Cluster.Name, secretSpec)

		cgStr := fmt.Sprintf(`
	resource "%s" "%s" {
	 name = "%s"

	 namespace_name = "default"
	
	 scope {
		cluster_group {
			name = "%s"
		}
	 }
	
	 %s
	}
	`, testConfig.SecretResource, testConfig.SecretResourceVar, testConfig.SecretName, testConfig.ScopeHelperResources.ClusterGroup.Name, secretSpec)

		switch scope {
		case commonscope.ClusterScope:
			return clStr
		case commonscope.ClusterGroupScope:
			return cgStr
		default:
			return ""
		}
	}

	return fmt.Sprintf(`
	%s

	resource "%s" "%s" {
	 name = "%s"

	 namespace_name = "default"

	 %s

	 %s
	}
	`, helperBlock, testConfig.SecretResource, testConfig.SecretResourceVar, testConfig.SecretName, scopeBlock, secretSpec)
}

func (testConfig *testAcceptanceConfig) checkResourceAttributes(scopeType commonscope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.SecretResourceName, "name", testConfig.SecretName),
	}

	switch scopeType {
	case commonscope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SecretResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case commonscope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SecretResourceName, "scope.0.cluster_group.0.name", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case commonscope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(secretscope.ScopesAllowed[:], `, `))
	}

	check = append(check, commonscope.MetaResourceAttributeCheck(testConfig.SecretResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyResourceCreation(scopeType commonscope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.SecretResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.SecretResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.SecretResourceName)
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
			fn := &secretclmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ManagementClusterName: commonscope.AttachedValue,
				Name:                  testConfig.SecretName,
				NamespaceName:         testConfig.NamespaceName,
				ProvisionerName:       commonscope.AttachedValue,
			}

			resp, err := config.TMCConnection.SecretResourceService.SecretResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped secret resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped secret resource is empty, resource: %s", testConfig.SecretResourceName)
			}
		case commonscope.ClusterGroupScope:
			fn := &secretcgmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName{
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				Name:             testConfig.SecretName,
				NamespaceName:    testConfig.NamespaceName,
			}

			resp, err := config.TMCConnection.ClusterGroupSecretResourceService.SecretResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped secret resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped secret resource is empty, resource: %s", testConfig.SecretResourceName)
			}
		case commonscope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(secretscope.ScopesAllowed[:], `, `))
		}

		return nil
	}
}

type (
	OperationConfig struct {
		username         string
		password         string
		imageRegistryURL string
		secretType       string
	}

	OperationOption func(*OperationConfig)
)

func WithUsername(val string) OperationOption {
	return func(config *OperationConfig) {
		config.username = val
	}
}

func WithPassword(val string) OperationOption {
	return func(config *OperationConfig) {
		config.password = val
	}
}

func WithURL(val string) OperationOption {
	return func(config *OperationConfig) {
		config.imageRegistryURL = val
	}
}

func WithSecretType(val string) OperationOption {
	return func(config *OperationConfig) {
		config.secretType = val
	}
}

// getTestSecretResourceSpec builds the input block for cluster secret resource based a recipe.
func (testConfig *testAcceptanceConfig) getTestSecretResourceSpec(opts ...OperationOption) string {
	cfg := &OperationConfig{
		username:         "someusername",
		password:         "somepassword",
		imageRegistryURL: "someregistryurl",
		secretType:       DockerSecretType,
	}

	for _, o := range opts {
		o(cfg)
	}

	var secretSpec string

	switch cfg.secretType {
	case DockerSecretType:
		secretSpec = fmt.Sprintf(`  spec {
			docker_config_json {
				username = "%s"
				password = "%s"
				image_registry_url = "%s"
			}
		}`, cfg.username, cfg.password, cfg.imageRegistryURL)
	case OpaqueSecretType:
		secretSpec = fmt.Sprintf(`  spec {
			opaque = {
				username = "%s"
				password = "%s"
			}
		}`, cfg.username, cfg.password)
	}

	return secretSpec
}
