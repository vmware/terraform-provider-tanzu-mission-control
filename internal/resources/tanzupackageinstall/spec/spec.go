/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

var (
	SpecSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "spec for package install.",
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				PackageRefKey: PackageRefKeySpec,
				RoleBindingScopeKey: {
					Type:        schema.TypeString,
					Description: "Role binding scope for service account which will be used by Package Install.",
					Computed:    true,
				},
				PathToInlineValuesKey: {
					Type:        schema.TypeString,
					Description: "File to read inline values from (in yaml format). User needs to specify the file path for inline values.",
					Optional:    true,
					DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
						yamlDataOld, err := helper.ReadYamlFileAsJSON(old)
						if err != nil {
							return false
						}

						yamlDataNew, err := helper.ReadYamlFileAsJSON(new)
						if err != nil {
							return false
						}

						return yamlDataOld == yamlDataNew
					},
				},
				InlineValuesKey: {
					Type:        schema.TypeMap,
					Description: "Deprecated, Use `path_to_inline_values` instead. Inline values to configure the Package Install.",
					Optional:    true,
					Sensitive:   true,
					Deprecated:  "This field is deprecated. For providing the inline values, use the new field: path_to_inline_values",
				},
			},
		},
	}

	PackageRefKeySpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Reference to the Package which will be installed.",
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				PackageMetadataNameKey: {
					Type:        schema.TypeString,
					Description: "Name of the Package Metadata.",
					Required:    true,
					ForceNew:    true,
				},
				VersionSelectionKey: versionSelectionSpec,
			},
		},
	}

	versionSelectionSpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Version Selection of the Package.",
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				ConstraintsKey: {
					Type:        schema.TypeString,
					Description: "Constraints to select Package. Example: constraints: 'v1.2.3', constraints: '<v1.4.0' etc.",
					Required:    true,
				},
			},
		},
	}
)

func HasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(SpecKey, PackageRefKey, PackageMetadataNameKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, PackageRefKey, VersionSelectionKey, ConstraintsKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, RoleBindingScopeKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, PathToInlineValuesKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, InlineValuesKey)):
		updateRequired = true
	}

	return updateRequired
}

type ValidateInlineValuesType func(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error

func ValidateInlineValues() ValidateInlineValuesType {
	return func(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
		value, ok := diff.GetOk(SpecKey)
		if !ok {
			return fmt.Errorf("spec: %v is not valid: minimum one valid spec block is required", value)
		}

		data, _ := value.([]interface{})

		if len(data) == 0 || data[0] == nil {
			return fmt.Errorf("spec data: %v is not valid: minimum one valid spec block is required", data)
		}

		specData := data[0].(map[string]interface{})
		inlineConfigFound := make([]string, 0)

		if pathToInlineValuesData, ok := specData[PathToInlineValuesKey]; ok {
			if pathToInlineValue, ok := pathToInlineValuesData.([]interface{}); ok && len(pathToInlineValue) != 0 {
				inlineConfigFound = append(inlineConfigFound, PathToInlineValuesKey)
			}
		}

		if clusterGroupData, ok := specData[InlineValuesKey]; ok {
			if clusterGroupValue, ok := clusterGroupData.([]interface{}); ok && len(clusterGroupValue) != 0 {
				inlineConfigFound = append(inlineConfigFound, InlineValuesKey)
			}
		}

		if len(inlineConfigFound) > 1 {
			return fmt.Errorf("found inline configs: %v are not valid: maximum one valid inline config attribute is allowed", strings.Join(inlineConfigFound, `, `))
		}

		return nil
	}
}
