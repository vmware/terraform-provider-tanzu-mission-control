// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package targetlocation

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
	targetlocationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/targetlocation"
)

type CredentialsTypeCtxKey string

const (
	credentialsTypeCtxKey CredentialsTypeCtxKey = "CredentialsType"
)

func ResourceTargetLocation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTargetLocationCreate,
		ReadContext:   resourceTargetLocationRead,
		UpdateContext: resourceTargetLocationUpdate,
		DeleteContext: resourceTargetLocationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceTargetLocationImporter,
		},
		CustomizeDiff: validateSchema,
		Schema:        backupTargetLocationResourceSchema,
	}
}

func resourceTargetLocationCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't create Tanzu Mission Control backup target location."))
	}

	model.FullName.ProviderName = TMCProviderName

	credentialsType, err := getCredentialsType(config, model.Spec.Credential.Name)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't create Tanzu Mission Control backup target location."))
	}

	err = validateSchemaByCredentials(model, credentialsType)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't create Tanzu Mission Control backup target location."))
	}

	request := &targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest{
		BackupLocation: model,
	}

	_, err = config.TMCConnection.TargetLocationService.TargetLocationResourceServiceCreate(request)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Tanzu Mission Control backup target location.\nName: %s, Provider: %s",
			request.BackupLocation.FullName.Name, request.BackupLocation.FullName.ProviderName))
	}

	ctx = context.WithValue(ctx, credentialsTypeCtxKey, credentialsType)

	return resourceTargetLocationRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceTargetLocationRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't read Tanzu Mission Control backup target location."))
	}

	targetLocationFn := model.FullName
	targetLocationFn.ProviderName = TMCProviderName

	resp, err := readResourceWait(ctx, &config, targetLocationFn)
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

		return diag.FromErr(errors.Wrapf(err, "Couldn't read backup target location.\nName: %s, Provider: %s",
			targetLocationFn.Name, targetLocationFn.ProviderName))
	} else if resp != nil {
		var credentialsType credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider

		credentialsTypeCtx := ctx.Value(credentialsTypeCtxKey)

		if credentialsTypeCtx == nil {
			credentialsType, err = getCredentialsType(config, resp.BackupLocation.Spec.Credential.Name)
			if err != nil {
				return diag.Errorf("Couldn't read backup target location. Credentials Error: %s", err.Error())
			}
		} else {
			credentialsType = credentialsTypeCtx.(credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider)
		}

		// When Credentials are TMC managed, bucket returns with a generated name, therefore removing it from the response to keep the state consistent
		err = tfModelResourceConverter.FillTFSchema(resp.BackupLocation, data)
		if err != nil {
			return diag.Errorf("Couldn't read backup target location.\nName: %s, Provider: %s", targetLocationFn.Name, targetLocationFn.ProviderName)
		}

		data.SetId(targetLocationFn.Name)
		modifyData(data, credentialsType)
	}

	return diags
}

func resourceTargetLocationDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't delete Tanzu Mission Control backup target location."))
	}

	targetLocationFn := model.FullName
	targetLocationFn.ProviderName = TMCProviderName
	err = config.TMCConnection.TargetLocationService.TargetLocationResourceServiceDelete(targetLocationFn)

	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete backup target location.\nName: %s, Provider: %s", targetLocationFn.Name, targetLocationFn.ProviderName))
	}

	return resourceTargetLocationRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceTargetLocationUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't update Tanzu Mission Control backup target location."))
	}

	model.FullName.ProviderName = TMCProviderName

	credentialsType, err := getCredentialsType(config, model.Spec.Credential.Name)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't update Tanzu Mission Control backup target location."))
	}

	err = validateSchemaByCredentials(model, credentialsType)

	if credentialsType == credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderAWSEC2 {
		spec := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
		model.Spec.Bucket = spec[SysBucketKey].(string)
		model.Spec.Region = spec[SysRegionKey].(string)
	}

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't update Tanzu Mission Control backup target location."))
	}

	request := &targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest{
		BackupLocation: model,
	}

	_, err = config.TMCConnection.TargetLocationService.TargetLocationResourceServiceUpdate(request)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update Tanzu Mission Control backup target location.\nName: %s, Provider: %s",
			request.BackupLocation.FullName.Name, request.BackupLocation.FullName.ProviderName))
	}

	return resourceTargetLocationRead(helper.GetContextWithCaller(ctx, helper.UpdateState), data, m)
}

func resourceTargetLocationImporter(ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)
	targetLocationName := data.Id()

	if targetLocationName == "" {
		return nil, errors.New("ID is needed to import a target location")
	}

	targetLocationFn := &targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName{
		ProviderName: TMCProviderName,
		Name:         targetLocationName,
	}
	resp, err := readResourceWait(ctx, &config, targetLocationFn)

	if err != nil || resp.BackupLocation == nil {
		return nil, errors.Errorf("Couldn't Import backup target location.\nName: %s, Provider: %s", targetLocationFn.Name, targetLocationFn.ProviderName)
	} else {
		credentialsType, err := getCredentialsType(config, resp.BackupLocation.Spec.Credential.Name)
		if err != nil {
			return nil, errors.Errorf("Couldn't Import backup target location. Credentials Error: %s", err.Error())
		}

		err = tfModelResourceConverter.FillTFSchema(resp.BackupLocation, data)
		if err != nil {
			return nil, err
		}

		data.SetId(targetLocationFn.Name)
		modifyData(data, credentialsType)
	}

	return []*schema.ResourceData{data}, err
}

func validateSchema(ctx context.Context, data *schema.ResourceDiff, m interface{}) error {
	specData := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
	configData, configExists := specData[ConfigKey]

	if configExists && len(configData.([]interface{})) > 0 {
		configDataMap := configData.([]interface{})[0].(map[string]interface{})
		awsConfigData, _ := configDataMap[AwsConfigKey].([]interface{})
		azureConfigData, _ := configDataMap[AzureConfigKey].([]interface{})

		if len(awsConfigData) > 0 && len(azureConfigData) > 0 {
			return errors.New("Config should be set with either AWS or Azure blocks but not both.")
		}
	}

	return nil
}

func readResourceWait(ctx context.Context, config *authctx.TanzuContext, resourceFullName *targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName) (resp *targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse, err error) {
	stopStatuses := map[targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase]bool{
		targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseERROR: true,
		targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseREADY: true,
	}

	responseStatus := targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhasePHASEUNSPECIFIED
	_, isStopStatus := stopStatuses[responseStatus]
	isCtxCallerSet := helper.IsContextCallerSet(ctx)

	for !isStopStatus {
		if isCtxCallerSet || (!isCtxCallerSet && responseStatus != targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhasePHASEUNSPECIFIED) {
			time.Sleep(5 * time.Second)
		}

		resp, err = config.TMCConnection.TargetLocationService.TargetLocationResourceServiceGet(resourceFullName)

		if err != nil || resp == nil || resp.BackupLocation == nil {
			if clienterrors.IsUnauthorizedError(err) {
				authctx.RefreshUserAuthContext(config, clienterrors.IsUnauthorizedError, err)
				continue
			}

			return nil, err
		}

		responseStatus = *resp.BackupLocation.Status.Phase
		_, isStopStatus = stopStatuses[responseStatus]
	}

	if responseStatus == targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseERROR {
		err = errors.Errorf("Target location '%s' has errored.", resourceFullName.Name)

		return nil, err
	}

	return resp, err
}

func validateSchemaByCredentials(tlModel *targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation, credentialsType credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider) error {
	if credentialsType == credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderAWSEC2 {
		if tlModel.Spec.Bucket != "" || tlModel.Spec.Config != nil {
			return errors.Errorf("Bucket field and Config block should not be set when credentials are %s", credentialsType)
		}

		tlModel.Spec.Config = &targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderSpecificConfig{}
	} else {
		if tlModel.Spec.Bucket == "" || tlModel.Spec.Config == nil {
			return errors.Errorf("Bucket field and Config block must be set when credentials are %s", credentialsType)
		}

		if tlModel.Spec.Config.S3Config != nil && tlModel.Spec.Region == "" {
			return errors.New("Region should be set when configuring self managed AWS Bucket.")
		}
	}

	return nil
}

func modifyData(data *schema.ResourceData, credentialsType credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider) {
	if credentialsType == credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderAWSEC2 {
		spec := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
		spec[SysBucketKey] = spec[BucketKey]
		spec[BucketKey] = ""
		spec[SysRegionKey] = spec[RegionKey]
		spec[RegionKey] = ""
		spec[ConfigKey] = []interface{}{}

		_ = data.Set(SpecKey, []interface{}{spec})
	}
}

func getCredentialsType(config authctx.TanzuContext, credentialsName string) (credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider, error) {
	credentialsFn := &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName{
		Name: credentialsName,
	}

	resp, err := config.TMCConnection.CredentialResourceService.CredentialResourceServiceGet(credentialsFn)
	if err != nil {
		return credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderPROVIDERUNSPECIFIED, err
	}

	return *resp.Credential.Spec.Meta.Provider, nil
}
