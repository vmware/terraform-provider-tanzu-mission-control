/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	custompolicytemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/custompolicytemplate"
)

func ResourceCustomPolicyTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCustomPolicyTemplateCreate,
		UpdateContext: resourceCustomPolicyTemplateUpdate,
		ReadContext:   resourceCustomPolicyTemplateRead,
		DeleteContext: resourceCustomPolicyTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCustomPolicyTemplateImporter,
		},
		Schema: customPolicyTemplateResourceSchema,
	}
}

func resourceCustomPolicyTemplateCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create IAM policy template."))
	}

	request := &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{
		Template: model,
	}

	_, err = config.TMCConnection.CustomPolicyTemplateResourceService.CustomPolicyTemplateResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create IAM policy template.\nName: %s", model.FullName.Name))
	}

	return resourceCustomPolicyTemplateRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceCustomPolicyTemplateUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})
	model.Spec.PolicyUpdateStrategy = &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategy{
		Type: custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyTypeINPLACEUPDATE.Pointer(),
	}

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update custom policy template."))
	}

	request := &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{
		Template: model,
	}

	_, err = config.TMCConnection.CustomPolicyTemplateResourceService.CustomPolicyTemplateResourceServiceUpdate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update custom policy template.\nName: %s", model.FullName.Name))
	}

	return resourceCustomPolicyTemplateRead(helper.GetContextWithCaller(ctx, helper.UpdateState), data, m)
}

func resourceCustomPolicyTemplateRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData

	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read custom policy template."))
	}

	customIAMRoleFn := model.FullName
	resp, err = config.TMCConnection.CustomPolicyTemplateResourceService.CustomPolicyTemplateResourceServiceGet(customIAMRoleFn)

	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			if !helper.IsContextCallerSet(ctx) {
				*data = schema.ResourceData{}

				return diags
			} else if helper.IsDeleteState(ctx) {
				// d.SetId("") is automatically called assuming delete returns no errors, but
				// it is added here for explicitness.
				_ = schema.RemoveFromState(data, m)

				return diags
			}
		}

		return diag.FromErr(errors.Wrapf(err, "Couldn't read custom policy template.\nName: %s", customIAMRoleFn.Name))
	} else if resp != nil {
		err = tfModelConverter.FillTFSchema(resp.Template, data)

		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Couldn't read custom policy template.\nName: %s", customIAMRoleFn.Name))
		}

		data.SetId(customIAMRoleFn.Name)
	}

	return diags
}

func resourceCustomPolicyTemplateDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete custom policy template."))
	}

	customIAMRoleFn := model.FullName
	err = config.TMCConnection.CustomPolicyTemplateResourceService.CustomPolicyTemplateResourceServiceDelete(customIAMRoleFn)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete custom policy template.\nName: %s", customIAMRoleFn.Name))
	}

	return resourceCustomPolicyTemplateRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceCustomPolicyTemplateImporter(_ context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)
	customPolicyTemplateName := data.Id()

	if customPolicyTemplateName == "" {
		return nil, errors.New("Cluster ID must be set to the custom IAM role name.")
	}

	customIAMRoleFn := &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName{
		Name: customPolicyTemplateName,
	}

	resp, err := config.TMCConnection.CustomPolicyTemplateResourceService.CustomPolicyTemplateResourceServiceGet(customIAMRoleFn)

	if err != nil || resp.Template == nil {
		return nil, errors.Wrapf(err, "Couldn't read custom policy template.\nName: %s", customIAMRoleFn.Name)
	}

	err = tfModelConverter.FillTFSchema(resp.Template, data)

	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't read custom policy template.\nName: %s", customIAMRoleFn.Name)
	}

	return []*schema.ResourceData{data}, err
}
