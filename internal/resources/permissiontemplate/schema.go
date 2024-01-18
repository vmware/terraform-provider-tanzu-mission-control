/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package permissiontemplate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
)

const (
	ResourceName = "tanzu-mission-control_permission_template"

	// Root Keys.
	CredentialsNameKey = "credentials_name"
	CapabilityKey      = "tanzu_capability"
	ProviderKey        = "tanzu_provider"

	// Computed Fields.
	TemplateKey                = "template"
	TemplateURLKey             = "template_url"
	TemplateValuesKey          = "template_values"
	UndefinedTemplateValuesKey = "undefined_template_values"
)

var capabilityProviderMap = map[string]string{
	"DATA_PROTECTION":      string(credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderAWSEC2),
	"MANAGED_K8S_PROVIDER": string(credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderAWSEKS),
}

var permissionTemplateSchema = map[string]*schema.Schema{
	CredentialsNameKey:         credentialsNameSchema,
	CapabilityKey:              capabilitySchema,
	ProviderKey:                providerSchema,
	TemplateKey:                templateSchema,
	TemplateURLKey:             templateURLSchema,
	TemplateValuesKey:          templateValuesSchema,
	UndefinedTemplateValuesKey: undefinedTemplateValuesSchema,
}

var credentialsNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "The name of the credentials to get permission template for.",
	Required:    true,
	ForceNew:    true,
}

var capabilitySchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: buildCapabilityProviderDescription(CapabilityKey),
	Required:    true,
	ForceNew:    true,
}

var providerSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: buildCapabilityProviderDescription(ProviderKey),
	Required:    true,
	ForceNew:    true,
}

var templateSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Base64 encoded permission template.",
	Computed:    true,
}

var templateURLSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "URL for permission template.",
	Computed:    true,
}

var templateValuesSchema = &schema.Schema{
	Type:        schema.TypeMap,
	Description: "Values to be sent as parameters for the template.",
	Computed:    true,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
}

var undefinedTemplateValuesSchema = &schema.Schema{
	Type:        schema.TypeMap,
	Description: "Values which are not defined in the template parameters definition.",
	Computed:    true,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
}

func buildCapabilityProviderDescription(schemaKey string) (description string) {
	if schemaKey == CapabilityKey {
		description = "The Tanzu capability of the credentials."
		validValues := make([]string, 0)

		for k, v := range capabilityProviderMap {
			valueDescription := fmt.Sprintf("When %s is set to '%s' %s must be set to '%s'.", CapabilityKey, k, ProviderKey, v)
			description = fmt.Sprintf("%s\n%s", description, valueDescription)

			validValues = append(validValues, k)
		}

		description = fmt.Sprintf("%s\nValid values are: %v", description, validValues)
	} else if schemaKey == ProviderKey {
		description = "The Tanzu provider of the credentials."
		validValues := make([]string, 0)

		for k, v := range capabilityProviderMap {
			valueDescription := fmt.Sprintf("When %s is set to '%s' %s must be set to '%s'.", ProviderKey, v, CapabilityKey, k)
			description = fmt.Sprintf("%s\n%s", description, valueDescription)
			validValues = append(validValues, v)
		}

		description = fmt.Sprintf("%s\nValid values are: %v", description, validValues)
	}

	return description
}
