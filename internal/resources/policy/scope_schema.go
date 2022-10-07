/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policy

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	scoperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

var (
	ScopeSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Scope for the security and custom policy, having one of the valid scopes: cluster, cluster_group or organization.",
		Required:    true,
		ForceNew:    true,
		MaxItems:    1,
		MinItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				clusterKey:      scoperesource.ClusterPolicyFullname,
				clusterGroupKey: scoperesource.ClusterGroupPolicyFullname,
				organizationKey: scoperesource.OrganizationPolicyFullname,
			},
		},
	}
	ScopesAllowed = [...]string{clusterKey, clusterGroupKey, organizationKey}
)

type (
	Scope int64
	// ScopedFullname is a struct for all types of policy full names.
	ScopedFullname struct {
		Scope                Scope
		FullnameCluster      *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName
		FullnameClusterGroup *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName
		FullnameOrganization *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName
	}
)

func ConstructScope(d *schema.ResourceData, name string) (scopedFullnameData *ScopedFullname) {
	value, ok := d.GetOk(ScopeKey)
	if !ok {
		return scopedFullnameData
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return scopedFullnameData
	}

	scopeData := data[0].(map[string]interface{})

	if v, ok := scopeData[clusterKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:           ClusterScope,
				FullnameCluster: scoperesource.ConstructClusterPolicyFullname(v1, name),
			}
		}
	}

	if v, ok := scopeData[clusterGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:                ClusterGroupScope,
				FullnameClusterGroup: scoperesource.ConstructClusterGroupPolicyFullname(v1, name),
			}
		}
	}

	if v, ok := scopeData[organizationKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:                OrganizationScope,
				FullnameOrganization: scoperesource.ConstructOrganizationPolicyFullname(v1, name),
			}
		}
	}

	return scopedFullnameData
}

func FlattenScope(scopedFullname *ScopedFullname) (data []interface{}, name string) {
	if scopedFullname == nil {
		return data, name
	}

	flattenScopeData := make(map[string]interface{})

	switch scopedFullname.Scope {
	case ClusterScope:
		name = scopedFullname.FullnameCluster.Name
		flattenScopeData[clusterKey] = scoperesource.FlattenClusterPolicyFullname(scopedFullname.FullnameCluster)
	case ClusterGroupScope:
		name = scopedFullname.FullnameClusterGroup.Name
		flattenScopeData[clusterGroupKey] = scoperesource.FlattenClusterGroupPolicyFullname(scopedFullname.FullnameClusterGroup)
	case OrganizationScope:
		name = scopedFullname.FullnameOrganization.Name
		flattenScopeData[organizationKey] = scoperesource.FlattenOrganizationPolicyFullname(scopedFullname.FullnameOrganization)
	case UnknownScope:
		fmt.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopesAllowed[:], `, `))
	}

	return []interface{}{flattenScopeData}, name
}

func ValidateScope(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
	value, ok := diff.GetOk(ScopeKey)
	if !ok {
		return fmt.Errorf("scope: %v is not valid: minimum one valid scope block is required", value)
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("scope data: %v is not valid: minimum one valid scope block is required among: %v", data, strings.Join(ScopesAllowed[:], `, `))
	}

	scopeData := data[0].(map[string]interface{})
	scopesFound := make([]string, 0)

	if v, ok := scopeData[clusterKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopesFound = append(scopesFound, clusterKey)
		}
	}

	if v, ok := scopeData[clusterGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopesFound = append(scopesFound, clusterGroupKey)
		}
	}

	if v, ok := scopeData[organizationKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopesFound = append(scopesFound, organizationKey)
		}
	}

	if len(scopesFound) == 0 {
		return fmt.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v", strings.Join(ScopesAllowed[:], `, `))
	} else if len(scopesFound) > 1 {
		return fmt.Errorf("found scopes: %v are not valid: maximum one valid scope type block is allowed", strings.Join(scopesFound, `, `))
	}

	return nil
}
