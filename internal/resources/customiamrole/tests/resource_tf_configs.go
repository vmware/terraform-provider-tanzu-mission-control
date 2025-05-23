// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package customiamroletests

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"

	customiamroleres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/customiamrole"
)

const (
	CustomIAMRoleResourceName = "test_custom_iam_role"
)

var (
	CustomIAMRoleResourceFullName = fmt.Sprintf("%s.%s", customiamroleres.ResourceName, CustomIAMRoleResourceName)
	CustomIAMRoleName             = acctest.RandomWithPrefix("acc-test-custom-iam-role")
)

type ResourceTFConfigBuilder struct {
	NodePoolDefinition string
}

func InitResourceTFConfigBuilder() *ResourceTFConfigBuilder {
	tfConfigBuilder := &ResourceTFConfigBuilder{}

	return tfConfigBuilder
}

func (builder *ResourceTFConfigBuilder) GetCustomFullIAMRoleConfig() string {
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"

		  spec {
			is_deprecated = false

			aggregation_rule {
			  cluster_role_selector {
				match_labels = {
				  key = "value"
				}
			  }

			  cluster_role_selector {
				match_expression {
				  key      = "aa"
				  operator = "Exists"
				  values   = ["aa", "bb", "cc"]
				}
			  }
			}

			allowed_scopes = ["ORGANIZATION", "CLUSTER_GROUP", "CLUSTER"]
			tanzu_permissions = ["account.credential.iam.get"]

			kubernetes_permissions {
			 rule {
			   resources      = ["deployments"]
			   verbs          = ["get", "list"]
			   api_groups     = ["*"]
			 }

			 rule {
			   verbs      = ["get", "list"]
			   api_groups = ["*"]
			   url_paths  = ["/healthz"]
			 }
			}
		  }
		}
		`,
		customiamroleres.ResourceName,
		CustomIAMRoleResourceName,
		CustomIAMRoleName,
	)
}

func (builder *ResourceTFConfigBuilder) GetCustomSlimIAMRoleConfig() string {
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"

		  spec {
			allowed_scopes = ["ORGANIZATION", "CLUSTER_GROUP", "CLUSTER"]

			kubernetes_permissions {
			  rule {
			    resources      = ["deployments"]
			    verbs          = ["get", "list"]
			    api_groups     = ["*"]
			  }
			}
		  }
		}
		`,
		customiamroleres.ResourceName,
		CustomIAMRoleResourceName,
		CustomIAMRoleName,
	)
}
