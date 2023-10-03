/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
	targetlocationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/targetlocation"
)

type CredentialsTypeCtxKey string
type ReadContextModeKey string
type ReadContextModeValue string

const (
	readContextMode       ReadContextModeKey   = "Mode"
	readContextModeCreate ReadContextModeValue = "Create"
	readContextModeUpdate ReadContextModeValue = "Update"

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
		Schema: backupTargetLocationResourceSchema,
	}
}

func resourceTargetLocationCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})
	model.FullName.ProviderName = "tmc"
	credentialsType, err := getCredentialsType(config, model.Spec.Credential.Name)

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't create Tanzu Mission Control backup target location."))
	}

	err = validateSchema(model, credentialsType)

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't create Tanzu Mission Control backup target location."))
	}

	request := &targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest{
		BackupLocation: model,
	}

	_, err = config.TMCConnection.TargetLocationService.TargetLocationResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Tanzu Mission Control backup target location. Name: %s, Provider: %s",
			request.BackupLocation.FullName.Name, request.BackupLocation.FullName.ProviderName))
	}

	ctx = context.WithValue(ctx, credentialsTypeCtxKey, credentialsType)

	return resourceTargetLocationRead(context.WithValue(ctx, readContextMode, readContextModeCreate), data, m)
}

func resourceTargetLocationRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	targetLocationFn := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey}).FullName
	targetLocationFn.ProviderName = "tmc"
	getResponse, err := config.TMCConnection.TargetLocationService.TargetLocationResourceServiceGet(targetLocationFn)

	if err != nil {
		if clienterrors.IsNotFoundError(err) && ctx.Value(readContextMode) == nil {
			*data = schema.ResourceData{}

			return diags
		}

		return diag.Errorf("Couldn't read backup target location. Name: %s, Provider: %s", targetLocationFn.Name, targetLocationFn.ProviderName)
	} else if getResponse.BackupLocation != nil {
		var credentialsType credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider

		credentialsTypeCtx := ctx.Value(credentialsTypeCtxKey)

		if credentialsTypeCtx == nil {
			credentialsType, err = getCredentialsType(config, getResponse.BackupLocation.Spec.Credential.Name)

			if err != nil {
				return diag.Errorf("Couldn't read backup target location. Credentials Error: %s", err.Error())
			}
		} else {
			credentialsType = credentialsTypeCtx.(credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider)
		}

		// When Credentials are TMC managed, bucket returns with a generated name, therefore removing it from the response to keep the state consistent
		err = tfModelResourceConverter.FillTFSchema(getResponse.BackupLocation, data)

		if err != nil {
			return diag.Errorf("Couldn't read backup target location. Name: %s, Provider: %s", targetLocationFn.Name, targetLocationFn.ProviderName)
		}

		data.SetId(targetLocationFn.Name)
		modifyData(data, credentialsType)
	}

	return diags
}

func resourceTargetLocationDelete(_ context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	targetLocationFn := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey}).FullName
	targetLocationFn.ProviderName = "tmc"
	err := config.TMCConnection.TargetLocationService.TargetLocationResourceServiceDelete(targetLocationFn)

	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete backup target location. Name: %s, Provider: %s", targetLocationFn.Name, targetLocationFn.ProviderName))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	_ = schema.RemoveFromState(data, m)

	return diags
}

func resourceTargetLocationUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})
	model.FullName.ProviderName = "tmc"
	credentialsType, err := getCredentialsType(config, model.Spec.Credential.Name)

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't update Tanzu Mission Control backup target location."))
	}

	err = validateSchema(model, credentialsType)

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
		return diag.FromErr(errors.Wrapf(err, "Couldn't update Tanzu Mission Control backup target location. Name: %s, Provider: %s",
			request.BackupLocation.FullName.Name, request.BackupLocation.FullName.ProviderName))
	}

	return resourceTargetLocationRead(context.WithValue(ctx, readContextMode, readContextModeUpdate), data, m)
}

func resourceTargetLocationImporter(_ context.Context, data *schema.ResourceData, config any) ([]*schema.ResourceData, error) {
	client := config.(authctx.TanzuContext)

	targetLocationName := data.Id()

	if targetLocationName == "" {
		return nil, errors.New("ID is needed to import a target location")
	}

	targetLocationFn := &targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName{
		ProviderName: "tmc",
		Name:         targetLocationName,
	}
	getResponse, err := client.TMCConnection.TargetLocationService.TargetLocationResourceServiceGet(targetLocationFn)

	if err != nil || getResponse.BackupLocation == nil {
		return nil, errors.Errorf("Couldn't Import backup target location. Name: %s, Provider: %s", targetLocationFn.Name, targetLocationFn.ProviderName)
	} else {
		credentialsType, err := getCredentialsType(client, getResponse.BackupLocation.Spec.Credential.Name)

		if err != nil {
			return nil, errors.Errorf("Couldn't Import backup target location. Credentials Error: %s", err.Error())
		}

		err = tfModelResourceConverter.FillTFSchema(getResponse.BackupLocation, data)

		if err != nil {
			return nil, err
		}

		data.SetId(targetLocationFn.Name)
		modifyData(data, credentialsType)
	}

	return []*schema.ResourceData{data}, err
}

func validateSchema(tlModel *targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation, credentialsType credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider) error {
	if credentialsType == credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderAWSEC2 {
		if tlModel.Spec.Bucket != "" || tlModel.Spec.Config != nil {
			return errors.Errorf("Bucket field and Config block should not be set when credentials are %s", credentialsType)
		}

		tlModel.Spec.Config = &targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderSpecificConfig{}
	} else {
		if tlModel.Spec.Bucket == "" || tlModel.Spec.Config == nil {
			return errors.Errorf("Bucket field and Config block must be set when credentials are %s", credentialsType)
		}

		if tlModel.Spec.Config.S3Config != nil && tlModel.Spec.Config.AzureConfig != nil {
			return errors.New("Config should be set with either AWS or Azure blocks but not both.")
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
