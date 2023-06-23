/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackage

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package/spec"
)

func DataSourceTanzuPackage() *schema.Resource {
	return &schema.Resource{
		Schema: packageSchema,
	}
}

var packageSchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the package. It represents version of the Package metadata",
		Required:    true,
	},
	namespaceKey: {
		Type:        schema.TypeString,
		Description: "Namespae of package.",
		Required:    true,
	},
	metadataNameKey: {
		Type:        schema.TypeString,
		Description: "Metadata name of package.",
		Required:    true,
	},
	commonscope.ScopeKey: scope.ScopeSchema,
	common.MetaKey:       common.Meta,
	spec.SpecKey:         spec.SpecSchema,
}
