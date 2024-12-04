// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package iampolicy

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	namespacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/namespace"
	organizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/organization"
	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/namespace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

var (
	scopeSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Scope of the resource on which the rolebinding has to be added, having one of the valid scopes: organization, cluster_group, cluster, workspace or namespace.",
		Required:    true,
		ForceNew:    true,
		MinItems:    1,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				organizationKey: organizationFullname,
				clusterGroupKey: clustergroup.ClusterGroupFullname,
				clusterKey:      cluster.ClusterFullname,
				workspaceKey:    workspace.WorkspaceFullname,
				namespaceKey:    namespace.NamespaceFullname,
			},
		},
	}
	scopesAllowed = [...]string{organizationKey, clusterGroupKey, clusterKey, workspaceKey, namespaceKey}
)

type (
	scope int64
	// ScopedFullname is a struct for all types of policy full names.
	scopedFullname struct {
		scope                scope
		fullnameOrganization *organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName
		fullnameClusterGroup *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName
		fullnameCluster      *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName
		fullnameWorkspace    *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName
		fullnameNamespace    *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName
	}
)

func constructScope(d *schema.ResourceData) (scopedFullnameData *scopedFullname) {
	value, ok := d.GetOk(scopeKey)
	if !ok {
		return scopedFullnameData
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return scopedFullnameData
	}

	scopeData := data[0].(map[string]interface{})

	if v, ok := scopeData[organizationKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &scopedFullname{
				scope:                organizationScope,
				fullnameOrganization: constructOrganizationFullname(v1),
			}
		}
	}

	if v, ok := scopeData[clusterGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &scopedFullname{
				scope:                clusterGroupScope,
				fullnameClusterGroup: clustergroup.ConstructClusterGroupFullname(v1),
			}
		}
	}

	if v, ok := scopeData[clusterKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &scopedFullname{
				scope:           clusterScope,
				fullnameCluster: cluster.ConstructClusterFullname(v1),
			}
		}
	}

	if v, ok := scopeData[workspaceKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &scopedFullname{
				scope:             workspaceScope,
				fullnameWorkspace: workspace.ConstructWorkspaceFullname(v1),
			}
		}
	}

	if v, ok := scopeData[namespaceKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &scopedFullname{
				scope:             namespaceScope,
				fullnameNamespace: namespace.ConstructNamespaceFullname(v1),
			}
		}
	}

	return scopedFullnameData
}

func flattenScope(scopedFullname *scopedFullname) (data []interface{}) {
	if scopedFullname == nil {
		return data
	}

	flattenScopeData := make(map[string]interface{})

	switch scopedFullname.scope {
	case organizationScope:
		flattenScopeData[organizationKey] = flattenOrganizationFullname(scopedFullname.fullnameOrganization)
	case clusterGroupScope:
		flattenScopeData[clusterGroupKey] = clustergroup.FlattenClusterGroupFullname(scopedFullname.fullnameClusterGroup)
	case clusterScope:
		flattenScopeData[clusterKey] = cluster.FlattenClusterFullname(scopedFullname.fullnameCluster)
	case workspaceScope:
		flattenScopeData[workspaceKey] = workspace.FlattenWorkspaceFullname(scopedFullname.fullnameWorkspace)
	case namespaceScope:
		flattenScopeData[namespaceKey] = namespace.FlattenNamespaceFullname(scopedFullname.fullnameNamespace)
	case unknownScope:
		fmt.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.\n", strings.Join(scopesAllowed[:], `, `))
	}

	return []interface{}{flattenScopeData}
}

func validateScope(_ context.Context, diff *schema.ResourceDiff, i interface{}) error {
	value, ok := diff.GetOk(scopeKey)
	if !ok {
		return fmt.Errorf("scope: %v is not valid: minimum one valid scope block is required", value)
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("scope data: %v is not valid: minimum one valid scope block is required among: %v", data, strings.Join(scopesAllowed[:], `, `))
	}

	scopeData := data[0].(map[string]interface{})
	scopesFound := make([]string, 0)

	if v, ok := scopeData[organizationKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopesFound = append(scopesFound, organizationKey)
		}
	}

	if v, ok := scopeData[clusterGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopesFound = append(scopesFound, clusterGroupKey)
		}
	}

	if v, ok := scopeData[clusterKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopesFound = append(scopesFound, clusterKey)
		}
	}

	if v, ok := scopeData[workspaceKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopesFound = append(scopesFound, workspaceKey)
		}
	}

	if v, ok := scopeData[namespaceKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopesFound = append(scopesFound, namespaceKey)
		}
	}

	if len(scopesFound) == 0 {
		return fmt.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v", strings.Join(scopesAllowed[:], `, `))
	} else if len(scopesFound) > 1 {
		return fmt.Errorf("found scopes: %v are not valid: maximum one valid scope type block is allowed", strings.Join(scopesFound, `, `))
	}

	return nil
}
