/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclass

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ResourceName = "tanzu-mission-control_cluster_class"

	// Root Keys.
	NameKey                  = "name"
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"

	// Computed Keys.
	WorkerClassesKey     = "worker_classes"
	VariablesSchemaKey   = "variables_schema"
	VariablesTemplateKey = "variables_template"
)

var clusterClassSchema = map[string]*schema.Schema{
	ManagementClusterNameKey: managementClusterNameSchema,
	ProvisionerNameKey:       provisionerNameSchema,
	NameKey:                  nameSchema,
	WorkerClassesKey:         WorkerClassesSchema,
	VariablesSchemaKey:       VariablesSchema,
	VariablesTemplateKey:     VariablesTemplateSchema,
}

var managementClusterNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Management cluster name",
	Required:    true,
	ForceNew:    true,
}

var provisionerNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster provisioner name",
	Required:    true,
	ForceNew:    true,
}

var nameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster class name",
	Required:    true,
	ForceNew:    true,
}

var WorkerClassesSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Available worker classes in cluster class",
	Computed:    true,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
}

var VariablesSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "JSON encoded OpenAPIV3 schema for the cluster class variables",
	Computed:    true,
}

var VariablesTemplateSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "JSON encoded example template for the cluster class variables",
	Computed:    true,
}
