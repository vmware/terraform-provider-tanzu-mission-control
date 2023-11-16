/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/dataprotection/cluster/backupschedule"
)

func ResourceBackupSchedule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBackupScheduleCreate,
		ReadContext:   resourceBackupScheduleRead,
		UpdateContext: resourceBackupScheduleUpdate,
		DeleteContext: resourceBackupScheduleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceBackupScheduleImporter,
		},
		Schema: backupScheduleResourceSchema,
	}
}

func resourceBackupScheduleCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Tanzu Mission Control backup schedule."))
	}

	diags = validateSchema(model, BackupScope(data.Get(ScopeKey).(string)))

	if diags.HasError() {
		return diags
	}

	request := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest{
		Schedule: model,
	}

	_, err = config.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Tanzu Mission Control backup schedule.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s, Schedule Name: %s",
			model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.ClusterName, model.FullName.Name))
	}

	return resourceBackupScheduleRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceBackupScheduleRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse

	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey, ClusterNameKey, ManagementClusterNameKey, ProvisionerNameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read Tanzu Mission Control backup schedule."))
	}

	backupScheduleFn := model.FullName
	resp, err = readResourceWait(ctx, &config, backupScheduleFn)

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

		return diag.FromErr(errors.Wrapf(err, "Couldn't read backup schedule.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s, Schedule Name: %s",
			backupScheduleFn.ManagementClusterName, backupScheduleFn.ProvisionerName, backupScheduleFn.ClusterName, backupScheduleFn.Name))
	} else if resp != nil {
		userExcludedNamespaces := getExcludedNamespaces(data, ExcludedNamespacesKey)
		systemExcludedNamespaces := getResponseSystemExcludedNamespaces(resp.Schedule, userExcludedNamespaces)
		resp.Schedule.Spec.Template.ExcludedNamespaces = userExcludedNamespaces

		if getSchemaCsiSnapshotTimeout(data) == "" {
			resp.Schedule.Spec.Template.CsiSnapshotTimeout = ""
		}

		err = tfModelResourceConverter.FillTFSchema(resp.Schedule, data)

		if err != nil {
			return diag.Errorf("Couldn't read backup schedule.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s, Schedule Name: %s",
				backupScheduleFn.ManagementClusterName, backupScheduleFn.ProvisionerName, backupScheduleFn.ClusterName, backupScheduleFn.Name)
		}

		fullNameList := []string{backupScheduleFn.ManagementClusterName, backupScheduleFn.ProvisionerName, backupScheduleFn.ClusterName, backupScheduleFn.Name}

		data.SetId(strings.Join(fullNameList, "/"))
		setSystemExcludedNamespaces(data, systemExcludedNamespaces)
	}

	return diags
}

func resourceBackupScheduleDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey, ClusterNameKey, ManagementClusterNameKey, ProvisionerNameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete Tanzu Mission Control backup schedule."))
	}

	backupScheduleFn := model.FullName
	err = config.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceDelete(backupScheduleFn)

	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete Tanzu Mission Control backup schedule.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s, Schedule Name: %s",
			backupScheduleFn.ManagementClusterName, backupScheduleFn.ProvisionerName, backupScheduleFn.ClusterName, backupScheduleFn.Name))
	}

	return resourceBackupScheduleRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceBackupScheduleUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update Tanzu Mission Control backup schedule."))
	}

	diags = validateSchema(model, BackupScope(data.Get(ScopeKey).(string)))

	if diags.HasError() {
		return diags
	}

	systemExcludedNamespaces := getExcludedNamespaces(data, SystemExcludedNamespacesKey)
	model.Spec.Template.ExcludedNamespaces = append(model.Spec.Template.ExcludedNamespaces, systemExcludedNamespaces...)

	request := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest{
		Schedule: model,
	}

	_, err = config.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceUpdate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update Tanzu Mission Control backup schedule.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s, Schedule Name: %s",
			model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.ClusterName, model.FullName.Name))
	}

	return resourceBackupScheduleRead(helper.GetContextWithCaller(ctx, helper.UpdateState), data, m)
}

func resourceBackupScheduleImporter(ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)
	backupScheduleID := data.Id()

	if backupScheduleID == "" {
		return nil, errors.New("ID is needed to import an TMC AKS cluster")
	}

	namesArray := strings.Split(backupScheduleID, "/")

	if len(namesArray) != 4 {
		return nil, errors.Errorf("Invalid backup schedule ID.\nBackup schedule id should consists of a full cluster name and the schedule name separated by '/'.\nProvided ID: %s", backupScheduleID)
	}

	backupScheduleFn := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName{
		ManagementClusterName: namesArray[0],
		ProvisionerName:       namesArray[1],
		ClusterName:           namesArray[2],
		Name:                  namesArray[3],
	}
	resp, err := readResourceWait(ctx, &config, backupScheduleFn)

	if err != nil || resp.Schedule == nil {
		return nil, errors.Errorf("Couldn't import backup schedule.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s, Schedule Name: %s",
			backupScheduleFn.ManagementClusterName, backupScheduleFn.ProvisionerName, backupScheduleFn.ClusterName, backupScheduleFn.Name)
	} else {
		userExcludedNamespaces := getExcludedNamespaces(data, ExcludedNamespacesKey)
		systemExcludedNamespaces := getResponseSystemExcludedNamespaces(resp.Schedule, userExcludedNamespaces)
		resp.Schedule.Spec.Template.ExcludedNamespaces = userExcludedNamespaces

		if getSchemaCsiSnapshotTimeout(data) == "" {
			resp.Schedule.Spec.Template.CsiSnapshotTimeout = ""
		}

		err = tfModelResourceConverter.FillTFSchema(resp.Schedule, data)

		if err != nil {
			return nil, err
		}

		setSystemExcludedNamespaces(data, systemExcludedNamespaces)
	}

	return []*schema.ResourceData{data}, err
}

func readResourceWait(ctx context.Context, config *authctx.TanzuContext, resourceFullName *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName) (resp *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse, err error) {
	stopStatuses := map[backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase]bool{
		backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseFAILEDVALIDATION: true,
		backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseENABLED:          true,
		backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePAUSED:           true,
	}

	responseStatus := backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePHASEUNSPECIFIED
	_, isStopStatus := stopStatuses[responseStatus]
	isCtxCallerSet := helper.IsContextCallerSet(ctx)

	for !isStopStatus {
		if isCtxCallerSet || (!isCtxCallerSet && responseStatus != backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePHASEUNSPECIFIED) {
			time.Sleep(5 * time.Second)
		}

		resp, err = config.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceGet(resourceFullName)

		if err != nil || resp == nil || resp.Schedule == nil {
			return nil, err
		}

		responseStatus = *resp.Schedule.Status.Phase
		_, isStopStatus = stopStatuses[responseStatus]
	}

	if responseStatus == backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseFAILEDVALIDATION {
		err = errors.Errorf("Failed validation of backup schedule '%s' in cluster: %s/%s/%s", resourceFullName.Name,
			resourceFullName.ManagementClusterName, resourceFullName.ProvisionerName, resourceFullName.ClusterName)

		return nil, err
	}

	return resp, err
}

func validateSchema(scheduleModel *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule, scope BackupScope) (diags diag.Diagnostics) {
	switch scope {
	case FullClusterBackupScope:
		if len(scheduleModel.Spec.Template.IncludedNamespaces) > 0 {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Included namespaces can't be configured when scope is %s", scope))
			diags = append(diags, d)
		}

		if scheduleModel.Spec.Template.LabelSelector != nil {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Lable selectors can't be configured when scope is %s", scope))
			diags = append(diags, d)
		}

		if len(scheduleModel.Spec.Template.OrLabelSelectors) > 0 {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Or lables selectors can't be configured when scope is %s", scope))
			diags = append(diags, d)
		}
	case NamespacesBackupScope:
		if len(scheduleModel.Spec.Template.IncludedNamespaces) == 0 {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Included namespaces must be configured when scope is %s", scope))
			diags = append(diags, d)
		}

		if len(scheduleModel.Spec.Template.ExcludedNamespaces) > 0 {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Excluded namespaces can't be configured when scope is %s", scope))
			diags = append(diags, d)
		}

		if scheduleModel.Spec.Template.LabelSelector != nil {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Lable selectors can't be configured when scope is %s", scope))
			diags = append(diags, d)
		}

		if len(scheduleModel.Spec.Template.OrLabelSelectors) > 0 {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Or lables selectors can't be configured when scope is %s", scope))
			diags = append(diags, d)
		}

	case LabelSelectorBackupScope:
		if scheduleModel.Spec.Template.LabelSelector == nil && scheduleModel.Spec.Template.OrLabelSelectors == nil {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Or/Lable selectors must be configured when scope is %s", scope))
			diags = append(diags, d)
		}

		if len(scheduleModel.Spec.Template.IncludedNamespaces) > 0 {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Included namespaces can't be configured when scope is %s", scope))
			diags = append(diags, d)
		}
	}

	return diags
}

func getResponseSystemExcludedNamespaces(scheduleModel *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule, userExcludedNamespaces []string) []string {
	var systemExcludedNamespaces []string

	for _, responseNs := range scheduleModel.Spec.Template.ExcludedNamespaces {
		found := false

		for _, userNs := range userExcludedNamespaces {
			if responseNs == userNs {
				found = true
				break
			}
		}

		if !found {
			systemExcludedNamespaces = append(systemExcludedNamespaces, responseNs)
		}
	}

	return systemExcludedNamespaces
}

func buildValidationErrorDiag(msg string) diag.Diagnostic {
	return diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Schema Validation Failed",
		Detail:   msg,
	}
}

func getExcludedNamespaces(data *schema.ResourceData, excludedNsKey string) []string {
	specData := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
	template := specData[TemplateKey].([]interface{})[0].(map[string]interface{})

	return helper.SetPrimitiveList[string](template[excludedNsKey], "")
}

func setSystemExcludedNamespaces(data *schema.ResourceData, systemExcludedNamespaces []string) {
	specData := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
	template := specData[TemplateKey].([]interface{})[0].(map[string]interface{})
	template[SystemExcludedNamespacesKey] = systemExcludedNamespaces

	_ = data.Set(SpecKey, []interface{}{specData})
}

func getSchemaCsiSnapshotTimeout(data *schema.ResourceData) string {
	specData := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
	template := specData[TemplateKey].([]interface{})[0].(map[string]interface{})

	return template[CsiSnapshotTimeoutKey].(string)
}
