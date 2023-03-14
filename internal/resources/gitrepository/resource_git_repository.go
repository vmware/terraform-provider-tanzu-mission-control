/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/status"
)

func ResourceGitRepository() *schema.Resource {
	return &schema.Resource{
		Schema:        gitRepositorySchema,
		CustomizeDiff: schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey, commonscope.ClusterGroupKey})),
	}
}

var gitRepositorySchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the Repository.",
		Required:    true,
		ForceNew:    true,
	},
	namespaceNameKey: {
		Type:        schema.TypeString,
		Description: "Name of Namespace.",
		Required:    true,
		ForceNew:    true,
	},
	orgIDKey: {
		Type:        schema.TypeString,
		Description: "ID of Organization.",
		Optional:    true,
		ForceNew:    true,
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			return old != "" && new == ""
		},
	},
	commonscope.ScopeKey: scope.ScopeSchema,
	common.MetaKey:       common.Meta,
	spec.SpecKey:         spec.SpecSchema,
	statusKey:            status.StatusSchema,
}
