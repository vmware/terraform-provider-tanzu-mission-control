/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	iampolicytemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iampolicytemplate"
)

func ResourceIAMPolicyTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIAMPolicyTemplateCreate,
		UpdateContext: resourceIAMPolicyTemplateUpdate,
		ReadContext:   resourceIAMPolicyTemplateRead,
		DeleteContext: resourceIAMPolicyTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceIAMPolicyTemplateImporter,
		},
		Schema: iamPolicyTemplateResourceSchema,
	}
}

func resourceIAMPolicyTemplateCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create IAM policy template."))
	}

	request := &iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{
		Template: model,
	}

	_, err = config.TMCConnection.IAMPolicyTemplateResourceService.IAMPolicyTemplateResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create IAM policy template.\nName: %s", model.FullName.Name))
	}

	return resourceIAMPolicyTemplateRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceIAMPolicyTemplateUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})
	model.Spec.PolicyUpdateStrategy = &iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategy{
		Type: iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyTypeINPLACEUPDATE.Pointer(),
	}

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update IAM policy template."))
	}

	request := &iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{
		Template: model,
	}

	_, err = config.TMCConnection.IAMPolicyTemplateResourceService.IAMPolicyTemplateResourceServiceUpdate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update IAM policy template.\nName: %s", model.FullName.Name))
	}

	return resourceIAMPolicyTemplateRead(helper.GetContextWithCaller(ctx, helper.UpdateState), data, m)
}

func resourceIAMPolicyTemplateRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData

	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read IAM policy template."))
	}

	customIAMRoleFn := model.FullName
	resp, err = config.TMCConnection.IAMPolicyTemplateResourceService.IAMPolicyTemplateResourceServiceGet(customIAMRoleFn)

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

		return diag.FromErr(errors.Wrapf(err, "Couldn't read IAM policy template.\nName: %s", customIAMRoleFn.Name))
	} else if resp != nil {
		err = tfModelConverter.FillTFSchema(resp.Template, data)

		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Couldn't read IAM policy template.\nName: %s", customIAMRoleFn.Name))
		}

		data.SetId(customIAMRoleFn.Name)
	}

	return diags
}

func resourceIAMPolicyTemplateDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete IAM policy template."))
	}

	customIAMRoleFn := model.FullName
	err = config.TMCConnection.IAMPolicyTemplateResourceService.IAMPolicyTemplateResourceServiceDelete(customIAMRoleFn)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete IAM policy template.\nName: %s", customIAMRoleFn.Name))
	}

	return resourceIAMPolicyTemplateRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceIAMPolicyTemplateImporter(_ context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)
	iamPolicyTemplateName := data.Id()

	if iamPolicyTemplateName == "" {
		return nil, errors.New("Cluster ID must be set to the custom IAM role name.")
	}

	customIAMRoleFn := &iampolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName{
		Name: iamPolicyTemplateName,
	}

	resp, err := config.TMCConnection.IAMPolicyTemplateResourceService.IAMPolicyTemplateResourceServiceGet(customIAMRoleFn)

	if err != nil || resp.Template == nil {
		return nil, errors.Wrapf(err, "Couldn't read IAM policy template.\nName: %s", customIAMRoleFn.Name)
	}

	err = tfModelConverter.FillTFSchema(resp.Template, data)

	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't read IAM policy template.\nName: %s", customIAMRoleFn.Name)
	}

	return []*schema.ResourceData{data}, err
}
