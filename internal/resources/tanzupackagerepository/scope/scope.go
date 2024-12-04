// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"

	repositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

// ScopedFullname is a struct for all types of package repository full names.
type ScopedFullname struct {
	Scope           commonscope.Scope
	FullnameCluster *repositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName
}

var (
	ScopesAllowed = [...]string{commonscope.ClusterKey}
	ScopeSchema   = commonscope.GetScopeSchema(
		commonscope.WithDescription(fmt.Sprintf("Scope for the package repository, having one of the valid scopes: %v.", strings.Join(ScopesAllowed[:], `, `))),
		commonscope.WithScopes(ScopesAllowed[:]))
)

func ConstructScope(d *schema.ResourceData, name string) (scopedFullnameData *ScopedFullname) {
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
				FullnameCluster: ConstructClusterPackageRepositoryFullname(clusterValue, name),
			}
		}
	}

	return scopedFullnameData
}

func FlattenScope(scopedFullname *ScopedFullname) (data []interface{}, name, namespace string) {
	if scopedFullname == nil {
		return data, name, namespace
	}

	flattenScopeData := make(map[string]interface{})

	switch scopedFullname.Scope {
	case commonscope.ClusterScope:
		if slices.Contains(ScopesAllowed[:], commonscope.ClusterKey) {
			name = scopedFullname.FullnameCluster.Name
			namespace = scopedFullname.FullnameCluster.NamespaceName
			flattenScopeData[commonscope.ClusterKey] = FlattenClusterPackageRepositoryFullname(scopedFullname.FullnameCluster)
		}
	case commonscope.UnknownScope:
		fmt.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopesAllowed[:], `, `))
	}

	return []interface{}{flattenScopeData}, name, namespace
}
