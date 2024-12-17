// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"

	dataprotectionmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/dataprotection"
	dataprotectioncgmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup/dataprotection"
)

type (
	Scope int64
	// ScopedFullname is a struct for all types of policy full names.
	ScopedFullname struct {
		Scope                Scope
		FullnameCluster      *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName
		FullnameClusterGroup *dataprotectioncgmodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionFullName
	}
)

func ConstructScope(d *schema.ResourceData) (scopedFullnameData *ScopedFullname) {
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
				FullnameCluster: ConstructClusterScopeFullname(v1),
			}
		}
	}

	if v, ok := scopeData[ClusterGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			scopedFullnameData = &ScopedFullname{
				Scope:                ClusterGroupScope,
				FullnameClusterGroup: ConstructClusterGroupScopeFullname(v1),
			}
		}
	}

	return scopedFullnameData
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
