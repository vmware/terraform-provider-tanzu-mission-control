//go:build imagepolicy
// +build imagepolicy

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package imagepolicyresource

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
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	policyworkspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindImage "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	imagePolicyResource    = policykindImage.ResourceName
	imagePolicyResourceVar = "test_image_policy"
	imagePolicyNamePrefix  = "tf-ip-test"
)

type testAcceptanceConfig struct {
	Provider                *schema.Provider
	ImagePolicyResource     string
	ImagePolicyResourceVar  string
	ImagePolicyResourceName string
	ImagePolicyName         string
	ScopeHelperResources    *policy.ScopeHelperResources
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                initTestProvider(t),
		ImagePolicyResource:     imagePolicyResource,
		ImagePolicyResourceVar:  imagePolicyResourceVar,
		ImagePolicyResourceName: fmt.Sprintf("%s.%s", imagePolicyResource, imagePolicyResourceVar),
		ImagePolicyName:         acctest.RandomWithPrefix(imagePolicyNamePrefix),
		ScopeHelperResources:    policy.NewScopeHelperResources(),
	}
}

func TestAcceptanceForImagePolicyResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start image policy resource acceptance tests!")

	// Test case for image policy resource with allowed-name-tag recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestImagePolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindImage.AllowedNameTagRecipe),
				Check:  testConfig.checkImagePolicyResourceAttributes(scope.WorkspaceScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped image policy acceptance test")
					}
				},
				Config: testConfig.getTestImagePolicyResourceBasicConfigValue(scope.OrganizationScope, policykindImage.AllowedNameTagRecipe),
				Check:  testConfig.checkImagePolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Image policy resource acceptance test complete for allowed-name-tag recipe!")

	// Test case for image policy resource with block-latest-tag recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestImagePolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindImage.BlockLatestTagRecipe),
				Check:  testConfig.checkImagePolicyResourceAttributes(scope.WorkspaceScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped image policy acceptance test")
					}
				},
				Config: testConfig.getTestImagePolicyResourceBasicConfigValue(scope.OrganizationScope, policykindImage.BlockLatestTagRecipe),
				Check:  testConfig.checkImagePolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Image policy resource acceptance test complete for block-latest-tag recipe!")

	// Test case for image policy resource with require-digest recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestImagePolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindImage.RequireDigestRecipe),
				Check:  testConfig.checkImagePolicyResourceAttributes(scope.WorkspaceScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped image policy acceptance test")
					}
				},
				Config: testConfig.getTestImagePolicyResourceBasicConfigValue(scope.OrganizationScope, policykindImage.RequireDigestRecipe),
				Check:  testConfig.checkImagePolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Image policy resource acceptance test complete for require-digest recipe!")

	// Test case for image policy resource with custom recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{

			{
				Config: testConfig.getTestImagePolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindImage.CustomRecipe),
				Check:  testConfig.checkImagePolicyResourceAttributes(scope.WorkspaceScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped image policy acceptance test")
					}
				},
				Config: testConfig.getTestImagePolicyResourceBasicConfigValue(scope.OrganizationScope, policykindImage.CustomRecipe),
				Check:  testConfig.checkImagePolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Image policy resource acceptance test complete for custom recipe!")
	t.Log("all image policy resource acceptance tests complete!")
}

func (testConfig *testAcceptanceConfig) getTestImagePolicyResourceBasicConfigValue(scope scope.Scope, recipe policykindImage.Recipe) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestPolicyResourceHelperAndScope(scope, policyoperations.ScopeMap[testConfig.ImagePolicyResource], false)
	inputBlock := testConfig.getTestImagePolicyResourceInput(recipe)

	return fmt.Sprintf(`
	%s

	resource "%s" "%s" {
	 name = "%s"

	 %s

	 spec {
	   %s

	   namespace_selector {
	     match_expressions {
	       key      = "component"
	       operator = "NotIn"
	       values   = [
	         "api-server",
	         "agent-gateway"
	       ]
	     }
	     match_expressions {
	       key      = "not-a-component"
	       operator = "DoesNotExist"
	       values   = []
	     }
	   }
	 }
	}
	`, helperBlock, testConfig.ImagePolicyResource, testConfig.ImagePolicyResourceVar, testConfig.ImagePolicyName, scopeBlock, inputBlock)
}

// getTestImagePolicyResourceInput builds the input block for image policy resource based a recipe.
func (testConfig *testAcceptanceConfig) getTestImagePolicyResourceInput(recipe policykindImage.Recipe) string {
	var inputBlock string

	switch recipe {
	case policykindImage.AllowedNameTagRecipe:
		inputBlock = `
    input {
      allowed_name_tag {
        audit = true
        rules {
          imagename = "bar"
          tag {
            negate = true
            value = "test"
          }
        }
      }
    }
`
	case policykindImage.BlockLatestTagRecipe:
		inputBlock = `
    input {
      block_latest_tag {
        audit = false
      }
    }
`
	case policykindImage.RequireDigestRecipe:
		inputBlock = `
    input {
      require_digest {
        audit = false
      }
    }
`
	case policykindImage.CustomRecipe:
		inputBlock = `
    input {
      custom {
        audit = true
        rules {
          hostname = "foo"
          imagename = "bar"
          port = "80"
          requiredigest = false
          tag {
            negate = false
            value = "test"
          }
        }
      }
    }
`
	case policykindImage.UnknownRecipe:
		log.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(policykindImage.RecipesAllowed[:], `, `))
	}

	return inputBlock
}

// checkImagePolicyResourceAttributes checks for image policy creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkImagePolicyResourceAttributes(scopeType scope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyImagePolicyResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.ImagePolicyResourceName, "name", testConfig.ImagePolicyName),
	}

	switch scopeType {
	case scope.WorkspaceScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.ImagePolicyResourceName, "scope.0.workspace.0.workspace", testConfig.ScopeHelperResources.Workspace.Name))
	case scope.OrganizationScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.ImagePolicyResourceName, "scope.0.organization.0.organization", testConfig.ScopeHelperResources.OrgID))
	case scope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policyoperations.ScopeMap[testConfig.ImagePolicyResource], `, `))
	}

	check = append(check, policy.MetaResourceAttributeCheck(testConfig.ImagePolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyImagePolicyResourceCreation(scopeType scope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.ImagePolicyResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.ImagePolicyResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.ImagePolicyResourceName)
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
		case scope.WorkspaceScope:
			fn := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName{
				WorkspaceName: testConfig.ScopeHelperResources.Workspace.Name,
				Name:          testConfig.ImagePolicyName,
			}

			resp, err := config.TMCConnection.WorkspacePolicyResourceService.ManageV1alpha1WorkspacePolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "workspace scoped image policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "workspace scoped image policy resource is empty, resource: %s", testConfig.ImagePolicyResourceName)
			}
		case scope.OrganizationScope:
			fn := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
				OrgID: testConfig.ScopeHelperResources.OrgID,
				Name:  testConfig.ImagePolicyName,
			}

			resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "organization scoped image policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "organization scoped image policy resource is empty, resource: %s", testConfig.ImagePolicyResourceName)
			}
		case scope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policyoperations.ScopeMap[testConfig.ImagePolicyResource], `, `))
		}

		return nil
	}
}
