/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"

	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	policyworkspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
)

type (
	SchemaOption func(*SchemaConfig)

	SchemaConfig struct {
		Description string
		Scopes      []string
	}
)

func WithDescription(d string) SchemaOption {
	return func(config *SchemaConfig) {
		config.Description = d
	}
}

func WithScopes(s []string) SchemaOption {
	return func(config *SchemaConfig) {
		config.Scopes = s
	}
}

func GetScopeSchema(opts ...SchemaOption) *schema.Schema {
	cfg := &SchemaConfig{}

	for _, o := range opts {
		o(cfg)
	}

	schemaForAllowedScopes := func() map[string]*schema.Schema {
		scopeSchemaMap := make(map[string]*schema.Schema)

		for _, scope := range cfg.Scopes {
			scopeSchemaMap[scope] = getSchemaForScope()(scope)
		}

		return scopeSchemaMap
	}()

	return &schema.Schema{
		Type:        schema.TypeList,
		Description: cfg.Description,
		Required:    true,
		ForceNew:    true,
		MaxItems:    1,
		MinItems:    1,
		Elem: &schema.Resource{
			Schema: schemaForAllowedScopes,
		},
	}
}

type (
	Scope int64
	// ScopedFullname is a struct for all types of policy full names.
	ScopedFullname struct {
		Scope                Scope
		FullnameCluster      *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName
		FullnameClusterGroup *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName
		FullnameWorkspace    *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName
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

	if v, ok := scopeData[ClusterKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:           ClusterScope,
				FullnameCluster: ConstructClusterPolicyFullname(v1, name),
			}
		}
	}

	if v, ok := scopeData[ClusterGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:                ClusterGroupScope,
				FullnameClusterGroup: ConstructClusterGroupPolicyFullname(v1, name),
			}
		}
	}

	if v, ok := scopeData[WorkspaceKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:             WorkspaceScope,
				FullnameWorkspace: ConstructWorkspacePolicyFullname(v1, name),
			}
		}
	}

	if v, ok := scopeData[OrganizationKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:                OrganizationScope,
				FullnameOrganization: ConstructOrganizationPolicyFullname(v1, name),
			}
		}
	}

	return scopedFullnameData
}

func FlattenScope(scopedFullname *ScopedFullname, scopesAllowed []string) (data []interface{}, name string) {
	if scopedFullname == nil {
		return data, name
	}

	flattenScopeData := make(map[string]interface{})

	switch scopedFullname.Scope {
	case ClusterScope:
		name = scopedFullname.FullnameCluster.Name
		flattenScopeData[ClusterKey] = FlattenClusterPolicyFullname(scopedFullname.FullnameCluster)
	case ClusterGroupScope:
		name = scopedFullname.FullnameClusterGroup.Name
		flattenScopeData[ClusterGroupKey] = FlattenClusterGroupPolicyFullname(scopedFullname.FullnameClusterGroup)
	case WorkspaceScope:
		name = scopedFullname.FullnameWorkspace.Name
		flattenScopeData[WorkspaceKey] = FlattenWorkspacePolicyFullname(scopedFullname.FullnameWorkspace)
	case OrganizationScope:
		name = scopedFullname.FullnameOrganization.Name
		flattenScopeData[OrganizationKey] = FlattenOrganizationPolicyFullname(scopedFullname.FullnameOrganization)
	case UnknownScope:
		fmt.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed, `, `))
	}

	return []interface{}{flattenScopeData}, name
}

type ValidateScopeType func(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error

func ValidateScope(scopesAllowed []string) ValidateScopeType {
	return func(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
		value, ok := diff.GetOk(ScopeKey)
		if !ok {
			return fmt.Errorf("scope: %v is not valid: minimum one valid scope block is required", value)
		}

		data, _ := value.([]interface{})

		if len(data) == 0 || data[0] == nil {
			return fmt.Errorf("scope data: %v is not valid: minimum one valid scope block is required among: %v", data, strings.Join(scopesAllowed, `, `))
		}

		scopeData := data[0].(map[string]interface{})
		scopesFound := make([]string, 0)

		if v, ok := scopeData[ClusterKey]; ok {
			if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
				scopesFound = append(scopesFound, ClusterKey)
			}
		}

		if v, ok := scopeData[ClusterGroupKey]; ok {
			if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
				scopesFound = append(scopesFound, ClusterGroupKey)
			}
		}

		if v, ok := scopeData[WorkspaceKey]; ok {
			if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
				scopesFound = append(scopesFound, WorkspaceKey)
			}
		}

		if v, ok := scopeData[OrganizationKey]; ok {
			if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
				scopesFound = append(scopesFound, OrganizationKey)
			}
		}

		if len(scopesFound) == 0 {
			return fmt.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v", strings.Join(scopesAllowed, `, `))
		} else if len(scopesFound) > 1 {
			return fmt.Errorf("found scopes: %v are not valid: maximum one valid scope type block is allowed", strings.Join(scopesFound, `, `))
		}

		if !slices.Contains(scopesAllowed, scopesFound[0]) {
			return fmt.Errorf("found scope: %v is not valid: minimum one valid scope type block is required among: %v", scopesFound[0], strings.Join(scopesAllowed, `, `))
		}

		return nil
	}
}
