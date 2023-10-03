/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/backupschedule"
)

type ReadContextModeKey string
type ReadContextModeValue string

const (
	readContextMode       ReadContextModeKey   = "Mode"
	readContextModeCreate ReadContextModeValue = "Create"
	readContextModeUpdate ReadContextModeValue = "Update"
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
	model := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})
	diags = validateSchema(model, BackupScope(data.Get(ScopeKey).(string)))

	if diags.HasError() {
		return diags
	}

	request := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest{
		Schedule: model,
	}

	_, err := config.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control backup schedule. Name: %s, Cluster: %s",
			request.Schedule.FullName.Name, request.Schedule.FullName.ClusterName))
	}

	return resourceBackupScheduleRead(context.WithValue(ctx, readContextMode, readContextModeCreate), data, m)
}

func resourceBackupScheduleRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	backupScheduleFn := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey, ClusterNameKey, ManagementClusterNameKey, ProvisionerNameKey}).FullName
	getResponse, err := config.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceGet(backupScheduleFn)

	if err != nil {
		if clienterrors.IsNotFoundError(err) && ctx.Value(readContextMode) == nil {
			*data = schema.ResourceData{}

			return diags
		}

		return diag.Errorf("Couldn't read backup schedule. Name: %s, Cluster: %s", backupScheduleFn.Name, backupScheduleFn.ClusterName)
	} else if getResponse.Schedule != nil {
		userExcludedNamespaces := getExcludedNamespaces(data, ExcludedNamespacesKey)
		systemExcludedNamespaces := getResponseSystemExcludedNamespaces(getResponse.Schedule, userExcludedNamespaces)
		getResponse.Schedule.Spec.Template.ExcludedNamespaces = userExcludedNamespaces

		if getSchemaCsiSnapshotTimeout(data) == "" {
			getResponse.Schedule.Spec.Template.CsiSnapshotTimeout = ""
		}

		err = tfModelResourceConverter.FillTFSchema(getResponse.Schedule, data)

		if err != nil {
			return diag.Errorf("Couldn't read backup schedule. Name: %s, Cluster: %s", backupScheduleFn.Name, backupScheduleFn.ClusterName)
		}

		fullNameList := []string{backupScheduleFn.ManagementClusterName, backupScheduleFn.ProvisionerName, backupScheduleFn.ClusterName, backupScheduleFn.Name}

		data.SetId(strings.Join(fullNameList, "/"))
		setSystemExcludedNamespaces(data, systemExcludedNamespaces)
	}

	return diags
}

func resourceBackupScheduleDelete(_ context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	backupScheduleFn := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey, ClusterNameKey, ManagementClusterNameKey, ProvisionerNameKey}).FullName
	err := config.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceDelete(backupScheduleFn)

	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control backup schedule. Name: %s, Cluster: %s", backupScheduleFn.Name, backupScheduleFn.ClusterName))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	_ = schema.RemoveFromState(data, m)

	return diags
}

func resourceBackupScheduleUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})
	diags = validateSchema(model, BackupScope(data.Get(ScopeKey).(string)))

	if diags.HasError() {
		return diags
	}

	systemExcludedNamespaces := getExcludedNamespaces(data, SystemExcludedNamespacesKey)
	model.Spec.Template.ExcludedNamespaces = append(model.Spec.Template.ExcludedNamespaces, systemExcludedNamespaces...)

	request := &backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest{
		Schedule: model,
	}

	_, err := config.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceUpdate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to update Tanzu Mission Control backup schedule. Name: %s, Cluster: %s",
			request.Schedule.FullName.Name, request.Schedule.FullName.ClusterName))
	}

	return resourceBackupScheduleRead(context.WithValue(ctx, readContextMode, readContextModeUpdate), data, m)
}

func resourceBackupScheduleImporter(_ context.Context, data *schema.ResourceData, config any) ([]*schema.ResourceData, error) {
	client := config.(authctx.TanzuContext)

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

	getResponse, err := client.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceGet(backupScheduleFn)

	if err != nil || getResponse.Schedule == nil {
		return nil, errors.Errorf("Couldn't read backup schedule. Name: %s, Cluster: %s", backupScheduleFn.Name, backupScheduleFn.ClusterName)
	} else {
		userExcludedNamespaces := getExcludedNamespaces(data, ExcludedNamespacesKey)
		systemExcludedNamespaces := getResponseSystemExcludedNamespaces(getResponse.Schedule, userExcludedNamespaces)
		getResponse.Schedule.Spec.Template.ExcludedNamespaces = userExcludedNamespaces

		if getSchemaCsiSnapshotTimeout(data) == "" {
			getResponse.Schedule.Spec.Template.CsiSnapshotTimeout = ""
		}

		err = tfModelResourceConverter.FillTFSchema(getResponse.Schedule, data)

		if err != nil {
			return nil, err
		}

		data.SetId(fmt.Sprintf("%s/%s", backupScheduleFn.ClusterName, backupScheduleFn.Name))
		setSystemExcludedNamespaces(data, systemExcludedNamespaces)
	}

	return []*schema.ResourceData{data}, err
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
		if scheduleModel.Spec.Template.LabelSelector == nil {
			d := buildValidationErrorDiag(fmt.Sprintf("(Template) Lable selectors must be configured when scope is %s", scope))
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

	return helper.SetPrimitiveList[string](template[excludedNsKey])
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
