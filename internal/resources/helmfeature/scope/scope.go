/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"

	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

// ScopedFullname is a struct for all types of helm release full names.
type ScopedFullname struct {
	Scope                commonscope.Scope
	FullnameCluster      *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName
	FullnameClusterGroup *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmFullName
}

var (
	ScopesAllowed = [...]string{commonscope.ClusterKey, commonscope.ClusterGroupKey}
	ScopeSchema   = commonscope.GetScopeSchema(
		commonscope.WithDescription(fmt.Sprintf("Scope for the Helm feature, having one of the valid scopes: %v.", strings.Join(ScopesAllowed[:], `, `))),
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
				FullnameCluster: ConstructClusterHelmFullname(clusterValue),
			}
		}
	}

	if clusterGroupData, ok := scopeData[commonscope.ClusterGroupKey]; ok && slices.Contains(ScopesAllowed[:], commonscope.ClusterGroupKey) {
		if clusterGroupValue, ok := clusterGroupData.([]interface{}); ok && len(clusterGroupValue) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:                commonscope.ClusterGroupScope,
				FullnameClusterGroup: ConstructClusterGroupHelmFullname(clusterGroupValue),
			}
		}
	}

	return scopedFullnameData
}

func FlattenScope(scopedFullname *ScopedFullname) (data []interface{}) {
	if scopedFullname == nil {
		return data
	}

	flattenScopeData := make(map[string]interface{})

	switch scopedFullname.Scope {
	case commonscope.ClusterScope:
		if slices.Contains(ScopesAllowed[:], commonscope.ClusterKey) {
			flattenScopeData[commonscope.ClusterKey] = FlattenClusterHelmFullname(scopedFullname.FullnameCluster)
		}
	case commonscope.ClusterGroupScope:
		if slices.Contains(ScopesAllowed[:], commonscope.ClusterGroupKey) {
			flattenScopeData[commonscope.ClusterGroupKey] = FlattenClusterGroupHelmFullname(scopedFullname.FullnameClusterGroup)
		}
	case commonscope.UnknownScope:
		fmt.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopesAllowed[:], `, `))
	}

	return []interface{}{flattenScopeData}
}
