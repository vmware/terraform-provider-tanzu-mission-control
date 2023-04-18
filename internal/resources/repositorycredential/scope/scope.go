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

	repositorycredentialclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/repositorycredential/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

// ScopedFullname is a struct for all types of kustomization full names.
type ScopedFullname struct {
	Scope           commonscope.Scope
	FullnameCluster *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName
	// FullnameClusterGroup *repositorycredentialclustermodel.

}

var (
	ScopesAllowed = [...]string{commonscope.ClusterKey, commonscope.ClusterGroupKey}
	ScopeSchema   = commonscope.GetScopeSchema(
		commonscope.WithDescription(fmt.Sprintf("Scope for the kustomization, having one of the valid scopes: %v.", strings.Join(ScopesAllowed[:], `, `))),
		commonscope.WithScopes(ScopesAllowed[:]))
)

func ConstructScope(d *schema.ResourceData, name, orgID string) (scopedFullnameData *ScopedFullname) {
	value, ok := d.GetOk(commonscope.ScopeKey)

	if !ok {
		return scopedFullnameData
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return scopedFullnameData
	}

	scopeData := data[0].(map[string]interface{})

	if v, ok := scopeData[commonscope.ClusterKey]; ok && slices.Contains(ScopesAllowed[:], commonscope.ClusterKey) {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:           commonscope.ClusterScope,
				FullnameCluster: ConstructClusterRepositoryCredentialFullname(v1, name, orgID),
			}
		}
	}

	//cluster group not implmented yet
	// if v, ok := scopeData[commonscope.ClusterGroupKey]; ok && slices.Contains(ScopesAllowed[:], commonscope.ClusterGroupKey) {
	// 	if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
	// 		scopedFullnameData = &ScopedFullname{
	// 			Scope:                commonscope.ClusterGroupScope,
	// 			FullnameClusterGroup: ConstructClusterGroupKustomizationFullname(v1, name, namespace, orgID),
	// 		}
	// 	}
	// }

	return scopedFullnameData
}

func FlattenScope(scopedFullname *ScopedFullname) (data []interface{}, name, orgID string) {
	if scopedFullname == nil {
		return data, name, orgID
	}

	flattenScopeData := make(map[string]interface{})

	switch scopedFullname.Scope {
	case commonscope.ClusterScope:
		if slices.Contains(ScopesAllowed[:], commonscope.ClusterKey) {
			name = scopedFullname.FullnameCluster.Name
			orgID = scopedFullname.FullnameCluster.OrgId
			flattenScopeData[commonscope.ClusterKey] = FlattenClusterKustomizationFullname(scopedFullname.FullnameCluster)
		}
	// cluster group scope not implmented yet
	// case commonscope.ClusterGroupScope:
	// 	if slices.Contains(ScopesAllowed[:], commonscope.ClusterGroupKey) {
	// 		name = scopedFullname.FullnameClusterGroup.Name
	// 		namespace = scopedFullname.FullnameClusterGroup.NamespaceName
	// 		orgID = scopedFullname.FullnameClusterGroup.OrgID
	// 		flattenScopeData[commonscope.ClusterGroupKey] = FlattenClusterGroupKustomizationFullname(scopedFullname.FullnameClusterGroup)
	// 	}
	case commonscope.UnknownScope:
		fmt.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopesAllowed[:], `, `))
	}

	return []interface{}{flattenScopeData}, name, orgID
}
