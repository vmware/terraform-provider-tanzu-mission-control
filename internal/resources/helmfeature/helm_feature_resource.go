/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmfeature

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmfeature/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmfeature/status"
)

const (
	ResourceName = "tanzu-mission-control_helm_feature"

	statusKey = "status"
)

func ResourceHelm() *schema.Resource {
	return &schema.Resource{
		Schema:        helmSchema,
		CustomizeDiff: schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey, commonscope.ClusterGroupKey})),
	}
}

var helmSchema = map[string]*schema.Schema{
	commonscope.ScopeKey: scope.ScopeSchema,
	common.MetaKey:       common.Meta,
	statusKey:            status.StatusSchema,
}
