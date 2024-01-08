/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package permissiontemplate

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	permissiontemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/permissiontemplate"
)

func DataSourcePermissionTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePermissionTemplateRead,
		Schema:      permissionTemplateSchema,
	}
}

func validateSchema(data *schema.ResourceData) (err error) {
	capability := data.Get(CapabilityKey).(string)
	provider := data.Get(ProviderKey).(string)

	capabilityMatchingProvider := capabilityProviderMap[capability]

	if provider != capabilityMatchingProvider {
		return errors.Errorf("When %s is set to '%s' %s must be set to '%s'.\nProvider is '%s'.", CapabilityKey, capability, ProviderKey, capabilityMatchingProvider, provider)
	}

	return err
}

func dataSourcePermissionTemplateRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var response *permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse

	err := validateSchema(data)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Schema validation failed."))
	}

	config := m.(authctx.TanzuContext)
	request, err := tfModelRequestConverter.ConvertTFSchemaToAPIModel(data, []string{CredentialsNameKey, CapabilityKey, ProviderKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read permission template."))
	}

	response, err = config.TMCConnection.PermissionTemplateService.PermissionTemplateResourceServiceGet(request)

	if err != nil {
		if !clienterrors.IsNotFoundError(err) {
			return diag.FromErr(errors.Wrapf(err, "Couldn't read permission template."))
		}

		response, err = config.TMCConnection.PermissionTemplateService.PermissionTemplateResourceServiceGenerate(request)

		if err != nil {
			diags = diag.FromErr(errors.Wrapf(err, "Couldn't read permission template."))
		}
	}

	if response.TemplateValues != nil && len(response.TemplateValues) > 0 {
		// This is necessary because sometimes the template parameters definition and the template values returned from the API do not match.
		err = removeUndefinedTemplateValues(response)

		if err != nil {
			diags = diag.FromErr(errors.Wrapf(err, "Couldn't read permission template."))
		}
	}

	err = tfModelResponseConverter.FillTFSchema(response, data)

	if err != nil {
		diags = diag.FromErr(errors.Wrapf(err, "Couldn't read permission template."))
	}

	idFields := []string{request.FullName.Name, request.Capability, string(*request.Provider)}

	data.SetId(strings.Join(idFields, "/"))

	return diags
}

func removeUndefinedTemplateValues(response *permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse) error {
	var templateJSON map[string]interface{}

	templateBytes, err := base64.StdEncoding.DecodeString(response.PermissionTemplate)

	if err != nil {
		return err
	}

	err = json.Unmarshal(templateBytes, &templateJSON)

	if err != nil {
		return err
	}

	definedTemplateValues := make(map[string]string)
	undefinedTemplateValues := make(map[string]string)
	templateParametersDefinition := templateJSON["Parameters"].(map[string]interface{})

	for key, value := range response.TemplateValues {
		_, keyExists := templateParametersDefinition[key]

		if keyExists {
			definedTemplateValues[key] = value
		} else {
			undefinedTemplateValues[key] = value
		}
	}

	response.TemplateValues = definedTemplateValues
	response.UndefinedTemplateValues = undefinedTemplateValues

	return nil
}
