// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"

	sourcesecretlustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
	sourcesecretlustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

// ScopedFullname is a struct for all types of source secret full names.
type ScopedFullname struct {
	Scope                commonscope.Scope
	FullnameCluster      *sourcesecretlustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName
	FullnameClusterGroup *sourcesecretlustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName
}

var (
	CredentialTypesAllowed = [...]string{commonscope.ClusterKey, commonscope.ClusterGroupKey}
	ScopeSchema            = commonscope.GetScopeSchema(
		commonscope.WithDescription(fmt.Sprintf("Scope for the source secret, having one of the valid scopes: %v.", strings.Join(CredentialTypesAllowed[:], `, `))),
		commonscope.WithScopes(CredentialTypesAllowed[:]))
)

func ConstructScope(d *schema.ResourceData, name string) (scopedFullnameData *ScopedFullname, scopesFound []string) {
	value, ok := d.GetOk(commonscope.ScopeKey)

	if !ok {
		return scopedFullnameData, scopesFound
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return scopedFullnameData, scopesFound
	}

	scopeData := data[0].(map[string]interface{})

	if clusterData, ok := scopeData[commonscope.ClusterKey]; ok && slices.Contains(CredentialTypesAllowed[:], commonscope.ClusterKey) {
		if clusterValue, ok := clusterData.([]interface{}); ok && len(clusterValue) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:           commonscope.ClusterScope,
				FullnameCluster: ConstructClusterSourcesecretFullname(clusterValue, name),
			}

			scopesFound = append(scopesFound, commonscope.ClusterKey)
		}
	}

	if clusterGroupData, ok := scopeData[commonscope.ClusterGroupKey]; ok && slices.Contains(CredentialTypesAllowed[:], commonscope.ClusterGroupKey) {
		if clusterGroupValue, ok := clusterGroupData.([]interface{}); ok && len(clusterGroupValue) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:                commonscope.ClusterGroupScope,
				FullnameClusterGroup: ConstructClusterGroupSourcesecretFullname(clusterGroupValue, name),
			}

			scopesFound = append(scopesFound, commonscope.ClusterGroupKey)
		}
	}

	return scopedFullnameData, scopesFound
}

func FlattenScope(scopedFullname *ScopedFullname) (data []interface{}, name string) {
	if scopedFullname == nil {
		return data, name
	}

	flattenScopeData := make(map[string]interface{})

	switch scopedFullname.Scope {
	case commonscope.ClusterScope:
		if slices.Contains(CredentialTypesAllowed[:], commonscope.ClusterKey) {
			name = scopedFullname.FullnameCluster.Name
			flattenScopeData[commonscope.ClusterKey] = FlattenClusterSourcesecretFullname(scopedFullname.FullnameCluster)
		}
	case commonscope.ClusterGroupScope:
		if slices.Contains(CredentialTypesAllowed[:], commonscope.ClusterGroupKey) {
			name = scopedFullname.FullnameClusterGroup.Name
			flattenScopeData[commonscope.ClusterGroupKey] = FlattenClusterGroupSourcesecretFullname(scopedFullname.FullnameClusterGroup)
		}
	case commonscope.UnknownScope:
		fmt.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(CredentialTypesAllowed[:], `, `))
	}

	return []interface{}{flattenScopeData}, name
}
