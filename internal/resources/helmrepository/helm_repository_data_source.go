/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrepository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	helmscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrepository/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrepository/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrepository/status"
)

const (
	ResourceName = "tanzu-mission-control_helm_repository"

	nameKey          = "name"
	namespaceNameKey = "namespace_name"
	statusKey        = "status"
)

func DataSourceHelmRepository() *schema.Resource {
	return &schema.Resource{
		Schema: helmSchema,
	}
}

var helmSchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the helm repository.",
		Optional:    true,
	},
	namespaceNameKey: {
		Type:        schema.TypeString,
		Description: "Name of Namespace.",
		Optional:    true,
		Default:     "*",
	},
	commonscope.ScopeKey: helmscope.ScopeSchema,
	common.MetaKey:       common.Meta,
	spec.SpecKey:         spec.SpecSchema,
	statusKey:            status.StatusSchema,
}
