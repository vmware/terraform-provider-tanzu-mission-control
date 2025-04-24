//go:build sourcesecret
// +build sourcesecret

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package sourcesecret

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	sourcesecretscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret/spec"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForSourceSecretDataSource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start source secret data source acceptance tests!")

	// Test case for source secret data source with SSH credential type.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped source secret acceptance test")
					}
				},
				Config: testConfig.getTestSourceSecretDataSourceBasicConfigValue(commonscope.ClusterScope, spec.SSHKey),
				Check:  testConfig.checkSourceSecretDataSourceAttributes(),
			},
			{
				Config: testConfig.getTestSourceSecretDataSourceBasicConfigValue(commonscope.ClusterScope, spec.UsernamePasswordKey, WithUsername("someusername")),
				Check:  testConfig.checkSourceSecretDataSourceAttributes(),
			},
			{
				Config: testConfig.getTestSourceSecretDataSourceBasicConfigValue(commonscope.ClusterGroupScope, spec.UsernamePasswordKey),
				Check:  testConfig.checkSourceSecretDataSourceAttributes(),
			},
			{
				Config: testConfig.getTestSourceSecretDataSourceBasicConfigValue(commonscope.ClusterGroupScope, spec.SSHKey, WithKnownhosts("somehosts")),
				Check:  testConfig.checkSourceSecretDataSourceAttributes(),
			},
		},
	},
	)

	t.Log("source secret data source acceptance test completed")
}

func (testConfig *testAcceptanceConfig) getTestSourceSecretDataSourceBasicConfigValue(scope commonscope.Scope, allowedCredential string, opts ...OperationOption) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestSourceSecretResourceHelperAndScope(scope, sourcesecretscope.CredentialTypesAllowed[:])
	credentialType := testConfig.getTestSourceSecretResourceCredential(allowedCredential, opts...)

	return fmt.Sprintf(`
	%s

	resource "%s" "%s" {
	 name = "%s"

	 %s

	 spec {
	   %s
	 }
	}

	data "%s" "%s" {
		name = tanzu-mission-control_repository_credential.test_source_secret.name

		%s
	}
	`, helperBlock, testConfig.SourceSecretResource, testConfig.SourceSecretResourceVar, testConfig.SourceSecretName, scopeBlock, credentialType, testConfig.SourceSecretResource, testConfig.SourceSecretDataSourceVar, scopeBlock)
}

// checkSourceSecretDataSourceAttributes checks for source secret creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkSourceSecretDataSourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifySourceSecretDataSourceCreation(testConfig.SourceSecretDataSourceName),
		resource.TestCheckResourceAttrPair(testConfig.SourceSecretDataSourceName, "name", testConfig.SourceSecretResourceName, "name"),
		resource.TestCheckResourceAttrSet(testConfig.SourceSecretDataSourceName, "id"),
	}

	check = append(check, MetaDataSourceAttributeCheck(testConfig.SourceSecretDataSourceName, testConfig.SourceSecretResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifySourceSecretDataSourceCreation(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have source secret resource %s", name)
		}

		return nil
	}
}
