/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/status"
)

func ResourcePackageRepository() *schema.Resource {
	return &schema.Resource{
		Schema: packageRepositorySchema,
	}
}

var packageRepositorySchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the package repository resource.",
		Required:    true,
		ForceNew:    true,
	},
	NamespaceKey: {
		Type:        schema.TypeString,
		Description: "Name of Namespace where package repository will be created.",
		Computed:    true,
	},
	disabledKey: {
		Type:        schema.TypeBool,
		Description: "If true, Package Repository is disabled for cluster.",
		Optional:    true,
		Default:     false,
	},
	commonscope.ScopeKey: scope.ScopeSchema,
	spec.SpecKey:         spec.SpecSchema,
	status.StatusKey:     status.StatusSchema,
}
