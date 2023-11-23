/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integrationschema

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ResourceName = "tanzu-mission-control_tanzu_observability_integration"

	// Common Keys.
	NameKey = "name"

	// Root Keys.
	ScopeKey = "scope"
	SpecKey  = "spec"

	// Scope Directive Keys.
	ClusterScopeKey      = "cluster"
	ClusterGroupScopeKey = "cluster_group"

	// Cluster Scope Directive Keys.
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"

	// Spec Directive Keys.
	ConfigurationsKey = "configurations"
	CredentialsKey    = "credentials_name"
	SecretsKey        = "secrets"
)

const (
	TanzuObservabilitySaaSValue = "tanzu-observability-saas"
)

var specAtLeastOneOf = []string{
	strings.Join([]string{SpecKey, "0", SecretsKey}, "."),
	strings.Join([]string{SpecKey, "0", CredentialsKey}, "."),
}

var clusterScopeConflictsWith = []string{strings.Join([]string{ScopeKey, "0", ClusterGroupScopeKey}, ".")}

var TOIntegrationSchema = map[string]*schema.Schema{
	ScopeKey: scopeSchema,
	SpecKey:  specSchema,
}

var scopeSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The scope to apply the integration on.",
	MaxItems:    1,
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterScopeKey: {
				Type:          schema.TypeList,
				Description:   "Cluster scope.",
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: clusterScopeConflictsWith,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ManagementClusterNameKey: {
							Type:        schema.TypeString,
							Description: "Management cluster name",
							Required:    true,
							ForceNew:    true,
						},
						ProvisionerNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster provisioner name",
							Required:    true,
							ForceNew:    true,
						},
						NameKey: {
							Type:        schema.TypeString,
							Description: "Cluster name",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			ClusterGroupScopeKey: {
				Type:        schema.TypeList,
				Description: "Cluster group scope.",
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						NameKey: {
							Type:        schema.TypeString,
							Description: "Cluster group name",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	},
}

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Integration specs.",
	MaxItems:    1,
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			CredentialsKey: {
				Type:         schema.TypeString,
				Description:  "Credential name is the name of the Organization's Account Credential to be used instead of secrets to add an integration.",
				Optional:     true,
				ExactlyOneOf: specAtLeastOneOf,
			},
			ConfigurationsKey: {
				Type:        schema.TypeString,
				Description: "The expected JSON encoded input schema for the integration. Can be found in v1alpha1/integration API.",
				Optional:    true,
			},
			SecretsKey: {
				Type:        schema.TypeMap,
				Description: "Secrets are for sensitive configurations. The values are write-only and will be masked when read.",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	},
}
