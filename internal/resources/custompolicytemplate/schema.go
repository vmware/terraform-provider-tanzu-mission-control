/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplate

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	ResourceName = "tanzu-mission-control_custom_policy_template"

	// Root Keys.
	NameKey = "name"
	SpecKey = "spec"

	// Spec Directive Keys.
	IsDeprecatedKey     = "is_deprecated"
	DataInventoryKey    = "data_inventory"
	TemplateManifestKey = "template_manifest"
	ObjectTypeKey       = "object_type"
	TemplateTypeKey     = "template_type"

	// Data Inventory Directive Keys.
	GroupKey   = "group"
	VersionKey = "version"
	KindKey    = "kind"
)

var customPolicyTemplateResourceSchema = map[string]*schema.Schema{
	NameKey:        nameSchema,
	SpecKey:        specSchema,
	common.MetaKey: common.Meta,
}

var nameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "The name of the custom policy template",
	Required:    true,
	ForceNew:    true,
}

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec block of the custom policy template",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			IsDeprecatedKey: {
				Type:        schema.TypeBool,
				Description: "Flag representing whether the custom policy template is deprecated.",
				Default:     false,
				Optional:    true,
			},
			TemplateManifestKey: {
				Type:        schema.TypeString,
				Description: "YAML formatted Kubernetes resource.\nThe Kubernetes object has to be of the type defined in ObjectType ('ConstraintTemplate').\nThe object name must match the name of the wrapping policy template.\nThis will be applied on the cluster after a policy is created using this version of the template.\nThis contains the latest version of the object. For previous versions, check Versions API.",
				Required:    true,
			},
			ObjectTypeKey: {
				Type:        schema.TypeString,
				Description: "The type of Kubernetes resource encoded in Object.\nCurrently, we only support OPAGatekeeper based 'ConstraintTemplate' object.",
				Optional:    true,
				Default:     "ConstraintTemplate",
			},
			TemplateTypeKey: {
				Type:        schema.TypeString,
				Description: "The type of policy template.\nCurrently, we only support 'OPAGatekeeper' based policy templates.",
				Optional:    true,
				Default:     "OPAGatekeeper",
			},
			DataInventoryKey: DataInventorySchema,
		},
	},
}

var DataInventorySchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "List of Kubernetes api-resource kinds that need to be synced/replicated in Gatekeeper in order to enforce policy rules on those resources.\nNote: This is used for OPAGatekeeper based templates, and should be used if the policy enforcement logic in Rego code uses cached data using \"data.inventory\" fields.",
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			GroupKey: {
				Type:        schema.TypeString,
				Description: "API resource group",
				Required:    true,
			},
			KindKey: {
				Type:        schema.TypeString,
				Description: "API resource kind",
				Required:    true,
			},
			VersionKey: {
				Type:        schema.TypeString,
				Description: "API resource version",
				Required:    true,
			},
		},
	},
}
