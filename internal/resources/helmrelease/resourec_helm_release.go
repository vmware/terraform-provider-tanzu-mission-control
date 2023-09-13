/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrelease

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease/status"
)

const (
	ResourceName = "tanzu-mission-control_helm_release"

	nameKey          = "name"
	namespaceNameKey = "namespace_name"
	statusKey        = "status"
)

func ResourceHelmRelease() *schema.Resource {
	return &schema.Resource{
		Schema:        getHelmReleaseSchema(false),
		CustomizeDiff: schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey, commonscope.ClusterGroupKey})),
	}
}

func getHelmReleaseSchema(isDataSource bool) map[string]*schema.Schema {
	var helmreleaseSchema = map[string]*schema.Schema{
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
		commonscope.ScopeKey: scope.ScopeSchema,
		common.MetaKey:       common.Meta,
		statusKey:            status.StatusSchema,
	}

	innerMap := map[string]*schema.Schema{
		spec.SpecKey: spec.SpecSchema,
	}

	for key, value := range innerMap {
		if isDataSource {
			helmreleaseSchema[key] = helper.UpdateDataSourceSchema(value)
		} else {
			helmreleaseSchema[key] = value
		}
	}

	return helmreleaseSchema
}
