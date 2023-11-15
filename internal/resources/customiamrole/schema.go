/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrole

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	customiamrolemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/customiamrole"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	ResourceName = "tanzu-mission-control_custom_iam_role"

	// Root Keys.
	NameKey = "name"
	SpecKey = "spec"

	// Spec Directive Keys.
	AllowedScopesKey         = "allowed_scopes"
	AggregationRuleKey       = "aggregation_rule"
	IsDeprecatedKey          = "is_deprecated"
	TanzuPermissionsKey      = "tanzu_permissions"
	KubernetesPermissionsKey = "kubernetes_permissions"

	// Kubernetes Permissions Directive Keys.
	RuleKey = "rule"

	// Rule Directive Keys.
	ResourcesKey     = "resources"
	APIGroupsKey     = "api_groups"
	URLPathsKey      = "url_paths"
	ResourceNamesKey = "resource_names"
	VerbsKey         = "verbs"

	// Aggregation Rule Directive Keys.
	ClusterRoleSelectorKey = "cluster_role_selector"

	// Cluster Role Selector Directive Keys.
	MatchLabelsKey     = "match_labels"
	MatchExpressionKey = "match_expression"

	// Match Expression Directive Keys.
	MeKey         = "key"
	MeOperatorKey = "operator"
	MeValuesKey   = "values"
)

var AllowedScopesValidValues = []string{
	string(customiamrolemodels.VmwareTanzuManageV1alpha1IamPermissionResourceORGANIZATION),
	string(customiamrolemodels.VmwareTanzuManageV1alpha1IamPermissionResourceMANAGEMENTCLUSTER),
	string(customiamrolemodels.VmwareTanzuManageV1alpha1IamPermissionResourcePROVISIONER),
	string(customiamrolemodels.VmwareTanzuManageV1alpha1IamPermissionResourceCLUSTERGROUP),
	string(customiamrolemodels.VmwareTanzuManageV1alpha1IamPermissionResourceCLUSTER),
	string(customiamrolemodels.VmwareTanzuManageV1alpha1IamPermissionResourceWORKSPACE),
	string(customiamrolemodels.VmwareTanzuManageV1alpha1IamPermissionResourceNAMESPACE),
}

var customIAMRoleResourceSchema = map[string]*schema.Schema{
	NameKey:        nameSchema,
	SpecKey:        specSchema,
	common.MetaKey: common.Meta,
}

var nameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "The name of the iam role",
	Required:    true,
	ForceNew:    true,
}

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec block of iam role",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AllowedScopesKey:         AllowedScopesSchema,
			AggregationRuleKey:       AggregationRuleSchema,
			KubernetesPermissionsKey: KubernetesPermissionsSchema,
			TanzuPermissionsKey: {
				Type:        schema.TypeList,
				Description: "Tanzu-specific permissions for the role.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			IsDeprecatedKey: {
				Type:        schema.TypeBool,
				Description: "Flag representing whether role is deprecated.",
				Default:     false,
				Optional:    true,
			},
		},
	},
}

var AllowedScopesSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: fmt.Sprintf("The allowed scopes for the iam role.\nValid values are (%s)", strings.Join(AllowedScopesValidValues, ", ")),
	Required:    true,
	Elem: &schema.Schema{
		Type:             schema.TypeString,
		ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(AllowedScopesValidValues, false)),
	},
}

var AggregationRuleSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Aggregation rules for the iam role.",
	MaxItems:    1,
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterRoleSelectorKey: {
				Type:        schema.TypeList,
				Description: "Cluster role selector for the iam role.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						MatchLabelsKey: {
							Type:        schema.TypeMap,
							Description: "Map of {key,value} pairs.\nA single {key,value} in the match_labels map is equivalent to an element of match_expression, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". \nThe requirements are ANDed.",
							Optional:    true,
						},
						MatchExpressionKey: {
							Type:        schema.TypeList,
							Description: "List of label selector requirements.\nThe requirements are ANDed.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									MeKey: {
										Type:        schema.TypeString,
										Description: "Key is the label key that the selector applies to.",
										Required:    true,
									},
									MeOperatorKey: {
										Type:        schema.TypeString,
										Description: "Operator represents a key's relationship to a set of values.\nValid operators are \"In\", \"NotIn\", \"Exists\" and \"DoesNotExist\".",
										Required:    true,
									},
									MeValuesKey: {
										Type:        schema.TypeList,
										Description: "Values is an array of string values.\nIf the operator is \"In\" or \"NotIn\", the values array must be non-empty.\nIf the operator is \"Exists\" or \"DoesNotExist\", the values array must be empty.\nThis array is replaced during a strategic merge patch.",
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

var KubernetesPermissionsSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Kubernetes permissions for the iam role.",
	MaxItems:    1,
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			RuleKey: RuleSchema,
		},
	},
}

var RuleSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Kubernetes rules.",
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			APIGroupsKey: {
				Type:        schema.TypeList,
				Description: "API groups.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			VerbsKey: {
				Type:        schema.TypeList,
				Description: "Verbs.",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			ResourcesKey: {
				Type:        schema.TypeList,
				Description: "Resources for the role.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			ResourceNamesKey: {
				Type:        schema.TypeList,
				Description: "Restricts the rule to resources by name.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			URLPathsKey: {
				Type:        schema.TypeList,
				Description: "Non-resource urls for the role.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}
