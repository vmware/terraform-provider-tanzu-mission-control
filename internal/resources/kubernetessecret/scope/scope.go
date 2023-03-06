/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"

	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
)

var ScopeSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Scope for the secret having one of the valid scopes for secret: currently we have only cluster scope",
	Required:    true,
	ForceNew:    true,
	MaxItems:    1,
	MinItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterKey: ClusterFullname,
		},
	},
}

type (
	Scope int64
	// ScopedFullname is a struct for all types of secret full names.
	ScopedFullname struct {
		Scope           Scope
		FullnameCluster *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName
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
				FullnameCluster: ConstructClusterFullname(v1, name),
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
		flattenScopeData[ClusterKey] = FlattenClusterFullname(scopedFullname.FullnameCluster)
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
