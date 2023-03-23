/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredential

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization/status"
)

func ResourceRepositorycredential() *schema.Resource {
	return &schema.Resource{
		Schema:        repositorycredentialSchema,
		CustomizeDiff: schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey, commonscope.ClusterGroupKey})),
	}
}

var repositorycredentialSchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the Repository Credential.",
		Required:    true,
		ForceNew:    true,
	},
	commonscope.ScopeKey: scope.ScopeSchema,
	common.MetaKey:       common.Meta,
	spec.SpecKey:         spec.SpecSchema,
	statusKey:            status.StatusSchema,
}
