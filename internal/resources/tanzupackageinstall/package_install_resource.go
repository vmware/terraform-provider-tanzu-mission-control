/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackageinstall

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/status"
)

func ResourcePackageInstall() *schema.Resource {
	return &schema.Resource{
		Schema: getResourceSchema(),
	}
}

func getResourceSchema() map[string]*schema.Schema {
	return getSecretSchema(false)
}

func getSecretSchema(isDataSource bool) map[string]*schema.Schema {
	var packageInstallSchema = map[string]*schema.Schema{
		nameKey: {
			Type:        schema.TypeString,
			Description: "Name of the package install resource.",
			Required:    true,
			ForceNew:    true,
		},
		NamespaceKey: {
			Type:        schema.TypeString,
			Description: "Name of Namespace where package install will be created.",
			Required:    true,
			ForceNew:    true,
		},
		commonscope.ScopeKey: scope.ScopeSchema,
		common.MetaKey:       common.Meta,
		status.StatusKey:     status.StatusSchema,
	}

	innerMap := map[string]*schema.Schema{
		spec.SpecKey: spec.SpecSchema,
	}

	for key, value := range innerMap {
		if isDataSource {
			packageInstallSchema[key] = helper.UpdateDataSourceSchema(value)
		} else {
			packageInstallSchema[key] = value
		}
	}

	return packageInstallSchema
}
