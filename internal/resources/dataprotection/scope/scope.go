/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionscope

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dataprotectionclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/dataprotection/cluster/dataprotection"
	dataprotectionclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/dataprotection/clustergroup/dataprotection"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

// ScopedFullname is a struct for all types of helm release full names.
type ScopedFullname struct {
	Scope                commonscope.Scope
	FullnameCluster      *dataprotectionclustermodel.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName
	FullnameClusterGroup *dataprotectionclustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName
}

var (
	ScopesAllowed = [...]string{commonscope.ClusterKey, commonscope.ClusterGroupKey}
	ScopeSchema   = commonscope.GetScopeSchema(
		commonscope.WithDescription(fmt.Sprintf("Scope for the Data Protection, having one of the valid scopes: %v.", strings.Join(ScopesAllowed[:], `, `))),
		commonscope.WithScopes(ScopesAllowed[:]))
)

func ConstructScope(d *schema.ResourceData) (scopedFullnameData *ScopedFullname) {
	value, ok := d.GetOk(commonscope.ScopeKey)

	if !ok {
		return scopedFullnameData
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return scopedFullnameData
	}

	scopeData := data[0].(map[string]interface{})

	if clusterData, ok := scopeData[commonscope.ClusterKey]; ok && slices.Contains(ScopesAllowed[:], commonscope.ClusterKey) {
		if clusterValue, ok := clusterData.([]interface{}); ok && len(clusterValue) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:           commonscope.ClusterScope,
				FullnameCluster: ConstructDataProtectionClusterFullname(clusterValue),
			}
		}
	}

	if clusterGroupData, ok := scopeData[commonscope.ClusterGroupKey]; ok && slices.Contains(ScopesAllowed[:], commonscope.ClusterGroupKey) {
		if clusterGroupValue, ok := clusterGroupData.([]interface{}); ok && len(clusterGroupValue) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:                commonscope.ClusterGroupScope,
				FullnameClusterGroup: ConstructDataProtectionClusterGroupFullname(clusterGroupValue),
			}
		}
	}

	return scopedFullnameData
}
