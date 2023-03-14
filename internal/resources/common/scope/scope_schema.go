/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package commonscope

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type (
	Scope int64

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
