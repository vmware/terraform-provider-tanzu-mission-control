/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/scope"
)

var resourceSchema = map[string]*schema.Schema{
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of the secret resource.",
		Required:    true,
		ForceNew:    true,
		ValidateFunc: validation.All(
			validation.StringLenBetween(1, 126),
			validation.StringIsNotEmpty,
			validation.StringIsNotWhiteSpace,
		),
	},
	NamespaceNameKey: {
		Type:        schema.TypeString,
		Description: "Name of Namespace where secret will be created.",
		Required:    true,
		ForceNew:    true,
		ValidateFunc: validation.All(
			validation.StringIsNotEmpty,
			validation.StringIsNotWhiteSpace,
		),
	},
	OrgIDKey: {
		Type:        schema.TypeString,
		Description: "ID of Organization.",
		Optional:    true,
	},
	scope.ScopeKey: scope.ScopeSchema,
	statusKey: {
		Type:        schema.TypeMap,
		Description: "Status for the Secret Export.",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	common.MetaKey: common.Meta,
	specKey:        secretSpec,
	ExportKey: {
		Type:        schema.TypeBool,
		Description: "Export the secret to all namespaces.",
		Optional:    true,
		Default:     false,
	},
}

var dataSourceSchema = map[string]*schema.Schema{
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of the secret resource.",
		Required:    true,
		ForceNew:    true,
		ValidateFunc: validation.All(
			validation.StringLenBetween(1, 126),
			validation.StringIsNotEmpty,
			validation.StringIsNotWhiteSpace,
		),
	},
	NamespaceNameKey: {
		Type:        schema.TypeString,
		Description: "Name of Namespace where secret will be created.",
		Required:    true,
		ForceNew:    true,
		ValidateFunc: validation.All(
			validation.StringIsNotEmpty,
			validation.StringIsNotWhiteSpace,
		),
	},
	OrgIDKey: {
		Type:        schema.TypeString,
		Description: "ID of Organization.",
		Optional:    true,
	},
	scope.ScopeKey: scope.ScopeSchema,
	statusKey: {
		Type:        schema.TypeMap,
		Description: "Status for the Secret Export.",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	common.MetaKey: common.Meta,
	specKey: {
		Type:        secretSpec.Type,
		Description: secretSpec.Description,
		Computed:    true,
		Elem:        secretSpec.Elem,
	},
	ExportKey: {
		Type:        schema.TypeBool,
		Description: "Export the secret to all namespaces.",
		Computed:    true,
	},
}

func TestGetSecretSchema(t *testing.T) {
	testCasea := []struct {
		name            string
		expectedSchema  map[string]*schema.Schema
		generatedSchema map[string]*schema.Schema
		expectedResult  bool
	}{
		{
			name:            "resource schema test",
			expectedSchema:  resourceSchema,
			generatedSchema: getResourceSchema(),
			expectedResult:  true,
		},
		{
			name:            "data source schema test",
			expectedSchema:  dataSourceSchema,
			generatedSchema: getDataSourceSchema(),
			expectedResult:  true,
		},
		{
			name: "generate wrong resource schema",
			expectedSchema: map[string]*schema.Schema{
				NameKey:          resourceSchema[NameKey],
				NamespaceNameKey: resourceSchema[NamespaceNameKey],
				OrgIDKey:         resourceSchema[OrgIDKey],
				scope.ScopeKey:   resourceSchema[scope.ScopeKey],
				common.MetaKey:   common.Meta,
				statusKey:        resourceSchema[statusKey],
				specKey: {
					Description: secretSpec.Description,
					Type:        secretSpec.Type,
					Elem:        secretSpec.Elem,
					Optional:    true,
				},
				ExportKey: resourceSchema[ExportKey],
			},
			generatedSchema: getResourceSchema(),
			expectedResult:  false,
		},
	}
	for _, test := range testCasea {
		t.Run(test.name,
			func(t *testing.T) {
				result := equalSchema(test.expectedSchema, test.generatedSchema)

				if result != test.expectedResult {
					t.Errorf("expected schema is not equal with generated schema, expected schema ")
				}
			},
		)
	}
}

func equalSchema(d1, d2 map[string]*schema.Schema) bool {
	if len(d1) != len(d2) {
		return false
	}

	ans := false

	for k, v1 := range d1 {
		v2, ok := d2[k]
		if !ok {
			return false
		}

		switch v1.Type {
		case schema.TypeList:
			elem1, ok := v1.Elem.(*schema.Resource)

			elem2, ok1 := v2.Elem.(*schema.Resource)

			switch {
			case ok && ok1:
				ans = equalSchema(elem1.Schema, elem2.Schema)
			case (ok && !ok1) || (!ok && ok1):
				ans = false
			default:
				ans = (v1.GoString() == v2.GoString())
			}

		default:
			ans = (v1.GoString() == v2.GoString())
		}

		v1.Elem = nil
		v2.Elem = nil

		if !ans || !(v1.GoString() == v2.GoString()) {
			return false
		}
	}

	return ans
}
