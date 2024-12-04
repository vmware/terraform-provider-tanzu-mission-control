// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"

	packageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

// ScopedFullname is a struct for all types of package full names.
type ScopedFullname struct {
	Scope           commonscope.Scope
	FullnameCluster *packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageFullName
}

var (
	ScopeAllowed = [...]string{commonscope.ClusterKey}
	ScopeSchema  = commonscope.GetScopeSchema(
		commonscope.WithDescription(fmt.Sprintf("Scope for the data source, having one of the valid scopes: %v.", strings.Join(ScopeAllowed[:], `, `))),
		commonscope.WithScopes(ScopeAllowed[:]))
)

func ConstructScope(d *schema.ResourceData) (scopedFullnameData *ScopedFullname, scopesFound []string) {
	value, ok := d.GetOk(commonscope.ScopeKey)

	if !ok {
		return scopedFullnameData, scopesFound
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return scopedFullnameData, scopesFound
	}

	scopeData := data[0].(map[string]interface{})

	if clusterData, ok := scopeData[commonscope.ClusterKey]; ok && slices.Contains(ScopeAllowed[:], commonscope.ClusterKey) {
		if clusterValue, ok := clusterData.([]interface{}); ok && len(clusterValue) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:           commonscope.ClusterScope,
				FullnameCluster: ConstructClusterPackageFullname(clusterValue),
			}

			scopesFound = append(scopesFound, commonscope.ClusterKey)
		}
	}

	return scopedFullnameData, scopesFound
}

func FlattenScope(scopedFullname *ScopedFullname) (data []interface{}) {
	if scopedFullname == nil {
		return data
	}

	flattenScopeData := make(map[string]interface{})

	switch scopedFullname.Scope {
	case commonscope.ClusterScope:
		if slices.Contains(ScopeAllowed[:], commonscope.ClusterKey) {
			flattenScopeData[commonscope.ClusterKey] = FlattenClusterPackageFullname(scopedFullname.FullnameCluster)
		}
	case commonscope.UnknownScope:
		fmt.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopeAllowed[:], `, `))
	}

	return []interface{}{flattenScopeData}
}
