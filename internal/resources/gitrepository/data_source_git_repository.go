/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func DataSourceGitRepository() *schema.Resource {
	return &schema.Resource{
		Schema:        gitRepositorySchema,
		CustomizeDiff: schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey, commonscope.ClusterGroupKey})),
	}
}
