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

	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

// ScopedFullname is a struct for all types of git repository full names.
type ScopedFullname struct {
	Scope                commonscope.Scope
	FullnameCluster      *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName
	FullnameClusterGroup *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName
}

var (
	ScopesAllowed = [...]string{commonscope.ClusterKey, commonscope.ClusterGroupKey}
	ScopeSchema   = commonscope.GetScopeSchema(
		commonscope.WithDescription(fmt.Sprintf("Scope for the git repository, having one of the valid scopes: %v.", strings.Join(ScopesAllowed[:], `, `))),
		commonscope.WithScopes(ScopesAllowed[:]))
)

func ConstructScope(d *schema.ResourceData, name, namespace, orgID string) (scopedFullnameData *ScopedFullname) {
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
				FullnameCluster: ConstructClusterGitRepositoryFullname(v1, name, namespace, orgID),
			}
		}
	}

	if v, ok := scopeData[commonscope.ClusterGroupKey]; ok && slices.Contains(ScopesAllowed[:], commonscope.ClusterGroupKey) {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:                commonscope.ClusterGroupScope,
				FullnameClusterGroup: ConstructClusterGroupGitRepositoryFullname(v1, name, namespace, orgID),
			}
		}
	}

	return scopedFullnameData
}

func FlattenScope(scopedFullname *ScopedFullname) (data []interface{}, name, namespace, orgID string) {
	if scopedFullname == nil {
		return data, name, namespace, orgID
	}

	flattenScopeData := make(map[string]interface{})

	switch scopedFullname.Scope {
	case commonscope.ClusterScope:
		if slices.Contains(ScopesAllowed[:], commonscope.ClusterKey) {
			name = scopedFullname.FullnameCluster.Name
			namespace = scopedFullname.FullnameCluster.NamespaceName
			orgID = scopedFullname.FullnameCluster.OrgID
			flattenScopeData[commonscope.ClusterKey] = FlattenClusterGitRepositoryFullname(scopedFullname.FullnameCluster)
		}
	case commonscope.ClusterGroupScope:
		if slices.Contains(ScopesAllowed[:], commonscope.ClusterGroupKey) {
			name = scopedFullname.FullnameClusterGroup.Name
			namespace = scopedFullname.FullnameClusterGroup.NamespaceName
			orgID = scopedFullname.FullnameClusterGroup.OrgID
			flattenScopeData[commonscope.ClusterGroupKey] = FlattenClusterGroupGitRepositoryFullname(scopedFullname.FullnameClusterGroup)
		}
	case commonscope.UnknownScope:
		fmt.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopesAllowed[:], `, `))
	}

	return []interface{}{flattenScopeData}, name, namespace, orgID
}
