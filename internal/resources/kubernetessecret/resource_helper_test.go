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

const (
	enableNamespaceExclusionsSpecKey = "enableNamespaceExclusions"
	namespaceExclusionsSpecKey       = "namespaceExclusions"
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
	specKey:        specSchema,
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
		Type:        specSchema.Type,
		Description: specSchema.Description,
		Computed:    true,
		Elem:        specSchema.Elem,
	},
	ExportKey: {
		Type:        schema.TypeBool,
		Description: "Export the secret to all namespaces.",
		Computed:    true,
	},
}

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the kubernetes secret",
	Required:    true,
	MaxItems:    1,
	MinItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			DockerConfigjsonKey: {
				Type:        schema.TypeList,
				Required:    true,
				Description: "SecretType definition - SECRET_TYPE_DOCKERCONFIGJSON, Kubernetes secrets type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						UsernameKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Username of the registry.",
							Required:    true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
						PasswordKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Password of the registry.",
							Required:    true,
							Sensitive:   true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
						ImageRegistryURLKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Server URL of the registry.",
							Required:    true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
					},
				},
			},
		},
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
		v2 := d2[k]

		switch v1.Type {
		case schema.TypeList:
			elem1, ok := v1.Elem.(*schema.Resource)

			elem2, ok1 := v2.Elem.(*schema.Resource)

			if ok && ok1 {
				ans = equalSchema(elem1.Schema, elem2.Schema)
			} else if (ok && !ok1) || (!ok && ok1) {
				ans = false
			} else {
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
