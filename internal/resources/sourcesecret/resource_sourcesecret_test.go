//go:build sourcesecret
// +build sourcesecret

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package sourcesecret

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	sourcesecretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
	sourcesecretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	sourcesecretscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret/spec"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                   initTestProvider(t),
		SourceSecretResource:       sourceSecretResource,
		SourceSecretResourceVar:    sourceSecretResourceVar,
		SourceSecretResourceName:   fmt.Sprintf("%s.%s", sourceSecretResource, sourceSecretResourceVar),
		SourceSecretName:           acctest.RandomWithPrefix(sourceSecretNamePrefix),
		ScopeHelperResources:       NewScopeHelperResources(),
		SourceSecretDataSourceVar:  sourceSecretDataSourceVar,
		SourceSecretDataSourceName: fmt.Sprintf("data.%s.%s", ResourceName, sourceSecretDataSourceVar),
	}
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_REPOCRED_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForSourceSecretResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute source secret tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_REPOCRED_ENV_TEST")
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

	t.Log("start source secret resource acceptance tests!")

	// Test case for source secret resource.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" && found {
						t.Skip("KUBECONFIG env var is not set for cluster scoped source secret acceptance test")
					}
				},
				Config: testConfig.getTestSourceSecretResourceBasicConfigValue(commonscope.ClusterScope, spec.SSHKey),
				Check:  testConfig.checkSourceSecretResourceAttributes(commonscope.ClusterScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder for cluster scope...")
						testConfig.setupHTTPMocksUpdate(t, commonscope.ClusterScope)
					}
				},
				Config: testConfig.getTestSourceSecretResourceBasicConfigValue(commonscope.ClusterScope, spec.UsernamePasswordKey, WithUsername("someusername")),
				Check:  testConfig.checkSourceSecretResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestSourceSecretResourceBasicConfigValue(commonscope.ClusterGroupScope, spec.UsernamePasswordKey),
				Check:  testConfig.checkSourceSecretResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder for cluster group scope...")
						testConfig.setupHTTPMocksUpdate(t, commonscope.ClusterGroupScope)
					}
				},
				Config: testConfig.getTestSourceSecretResourceBasicConfigValue(commonscope.ClusterGroupScope, spec.SSHKey, WithKnownhosts("somehosts")),
				Check:  testConfig.checkSourceSecretResourceAttributes(commonscope.ClusterGroupScope),
			},
		},
	},
	)

	t.Log("source secret resource acceptance test completed")
}

func (testConfig *testAcceptanceConfig) getTestSourceSecretResourceBasicConfigValue(scope commonscope.Scope, allowedCredential string, opts ...OperationOption) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestSourceSecretResourceHelperAndScope(scope, sourcesecretscope.CredentialTypesAllowed[:])
	credentialType := testConfig.getTestSourceSecretResourceCredential(allowedCredential, opts...)

	if _, found := os.LookupEnv("ENABLE_REPOCRED_ENV_TEST"); !found {
		clStr := fmt.Sprintf(`
	resource "%s" "%s" {
	 name = "%s"

	 scope {
		cluster {
			name = "%s"
			management_cluster_name = "attached"
			provisioner_name = "attached"
		}
	 }

	 spec {
	   %s
	 }
	}
	`, testConfig.SourceSecretResource, testConfig.SourceSecretResourceVar, testConfig.SourceSecretName, testConfig.ScopeHelperResources.Cluster.Name, credentialType)

		cgStr := fmt.Sprintf(`
	resource "%s" "%s" {
	 name = "%s"

	 scope {
		cluster_group {
			name = "%s"
		}
	 }

	 spec {
	   %s
	 }
	}
	`, testConfig.SourceSecretResource, testConfig.SourceSecretResourceVar, testConfig.SourceSecretName, testConfig.ScopeHelperResources.ClusterGroup.Name, credentialType)

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

	 %s

	 spec {
	   %s
	 }
	}
	`, helperBlock, testConfig.SourceSecretResource, testConfig.SourceSecretResourceVar, testConfig.SourceSecretName, scopeBlock, credentialType)
}

type (
	OperationConfig struct {
		username   string
		knownhosts string
	}

	OperationOption func(*OperationConfig)
)

func WithUsername(val string) OperationOption {
	return func(config *OperationConfig) {
		config.username = val
	}
}

func WithKnownhosts(val string) OperationOption {
	return func(config *OperationConfig) {
		config.knownhosts = val
	}
}

// getTestSourceSecretResourceCredential builds the input block for source secret resource.
// nolint: gosec
func (testConfig *testAcceptanceConfig) getTestSourceSecretResourceCredential(allowedCredential string, opts ...OperationOption) string {
	cfg := &OperationConfig{
		username:   "testusername",
		knownhosts: "testhostes",
	}

	for _, o := range opts {
		o(cfg)
	}

	var credentialType string

	switch allowedCredential {
	case spec.UsernamePasswordKey:
		credentialType = fmt.Sprintf(`
    data {
      username_password {
        username = "%s"
        password = "testpassword"
      }
    }
`, cfg.username)
	case spec.SSHKey:
		credentialType = fmt.Sprintf(`
    data {
      ssh_key {
        identity    = "testidentity"
        known_hosts = "%s"
      }
    }
`, cfg.knownhosts)
	}

	return credentialType
}

// checkSourceSecretResourceAttributes checks for source secret creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkSourceSecretResourceAttributes(scopeType commonscope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifySourceSecretResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.SourceSecretResourceName, "name", testConfig.SourceSecretName),
	}

	switch scopeType {
	case commonscope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SourceSecretResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case commonscope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SourceSecretResourceName, "scope.0.cluster_group.0.name", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case commonscope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(sourcesecretscope.CredentialTypesAllowed[:], `, `))
	}

	check = append(check, MetaResourceAttributeCheck(testConfig.SourceSecretResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifySourceSecretResourceCreation(scopeType commonscope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.SourceSecretResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.SourceSecretResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.SourceSecretResourceName)
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
			fn := &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName{
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ManagementClusterName: commonscope.AttachedValue,
				Name:                  testConfig.SourceSecretName,
				ProvisionerName:       commonscope.AttachedValue,
			}

			resp, err := config.TMCConnection.ClusterSourcesecretResourceService.ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped source secret resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped source secret resource is empty, resource: %s", testConfig.SourceSecretResourceName)
			}
		case commonscope.ClusterGroupScope:
			fn := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName{
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				Name:             testConfig.SourceSecretName,
			}

			resp, err := config.TMCConnection.ClusterGroupSourcesecretResourceService.ManageV1alpha1ClustergroupFluxcdSourcesecretResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped source secret resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped source secret resource is empty, resource: %s", testConfig.SourceSecretResourceName)
			}
		case commonscope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(sourcesecretscope.CredentialTypesAllowed[:], `, `))
		}

		return nil
	}
}
