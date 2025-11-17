// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package dataprotection

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	dataprotectionmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/dataprotection"
	dataprotectioncgmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup/dataprotection"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection/scope"
)

func ResourceEnableDataProtection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEnableDataProtectionCreate,
		ReadContext:   resourceEnableDataProtectionRead,
		UpdateContext: resourceEnableDataProtectionUpdate,
		DeleteContext: resourceEnableDataProtectionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceEnableDataProtectionImporter,
		},
		Schema:        enableDataProtectionSchema,
		CustomizeDiff: schema.CustomizeDiffFunc(scope.ValidateScope([]string{scope.ClusterKey, scope.ClusterGroupKey})),
	}
}

func resourceEnableDataProtectionCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	scopedFullnameData := scope.ConstructScope(data)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control data protection entry; Scope full name is empty")
	}

	err := enableDataProtection(&config, scopedFullnameData, data, false)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't create Tanzu Mission Control data protection configuration."))
	}

	return resourceEnableDataProtectionRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceEnableDataProtectionRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	scopedFullnameData := scope.ConstructScope(data)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control data protection entry; Scope full name is empty")
	}

	config := m.(authctx.TanzuContext)

	err := populateDataFromServer(ctx, config, scopedFullnameData, data)
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

		return diag.FromErr(err)
	}

	// remove the existing cluster level resource from state if it is now
	// managed at the cluster group level.
	if scopedFullnameData.Scope == scope.ClusterScope {
		value, ok := data.GetOk(common.MetaKey)
		if ok && len(value.([]interface{})) > 0 {
			metaData := value.([]interface{})[0].(map[string]interface{})
			annotations := metaData[common.AnnotationsKey].(map[string]interface{})

			if _, ok := annotations[commonscope.BatchUIDAnnotationKey]; ok {
				_ = schema.RemoveFromState(data, m)

				return diags
			}
		}
	}

	return diags
}

func resourceEnableDataProtectionDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	scopedFullnameData := scope.ConstructScope(data)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control data protection entry; Scope full name is empty")
	}

	deletionPolicy := data.Get(DeletionPolicyKey).([]interface{})
	deleteBackups := false
	force := false

	if len(deletionPolicy) > 0 {
		deleteBackups = deletionPolicy[0].(map[string]interface{})[DeleteBackupsKey].(bool)
		force = deletionPolicy[0].(map[string]interface{})[ForceDeleteKey].(bool)
	}

	switch scopedFullnameData.Scope {
	case scope.ClusterScope:
		err := config.TMCConnection.DataProtectionService.DataProtectionResourceServiceDelete(scopedFullnameData.FullnameCluster, deleteBackups)
		if err != nil && !clienterrors.IsNotFoundError(err) {
			return diag.FromErr(errors.Wrap(err, "Unable to delete Tanzu Mission Control data protection"))
		}
	case scope.ClusterGroupScope:
		err := config.TMCConnection.ClusterGroupDataProtectionService.DataProtectionResourceServiceDelete(scopedFullnameData.FullnameClusterGroup, deleteBackups, force)
		if err != nil && !clienterrors.IsNotFoundError(err) {
			return diag.FromErr(errors.Wrap(err, "Unable to delete Tanzu Mission Control data protection"))
		}
	case scope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required. Please check the schema.")
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return resourceEnableDataProtectionRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceEnableDataProtectionUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	scopedFullnameData := scope.ConstructScope(data)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control data protection entry; Scope full name is empty")
	}

	err := enableDataProtection(&config, scopedFullnameData, data, true)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't create Tanzu Mission Control data protection configuration."))
	}

	return resourceEnableDataProtectionRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceEnableDataProtectionImporter(ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)

	fullName := data.Id()
	fullNameParts := strings.Split(fullName, "/")

	scopedFullname := &scope.ScopedFullname{}

	switch len(fullNameParts) {
	case 1:
		scopedFullname.Scope = scope.ClusterGroupScope
		scopedFullname.FullnameClusterGroup = &dataprotectioncgmodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionFullName{
			ClusterGroupName: fullNameParts[0],
		}

	case 3:
		scopedFullname.Scope = scope.ClusterScope
		scopedFullname.FullnameCluster = &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName{
			ManagementClusterName: fullNameParts[0],
			ProvisionerName:       fullNameParts[1],
			ClusterName:           fullNameParts[2],
		}
	default:
		return nil, errors.New("unexpected resource id")
	}

	err := populateDataFromServer(ctx, config, scopedFullname, data)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't import data protection configuration.")
	}

	if err == nil {
		if _, ok := data.GetOk(DeletionPolicyKey); !ok {
			deletionPolicyMap := map[string]interface{}{
				DeleteBackupsKey: false,
			}

			_ = data.Set(DeletionPolicyKey, []interface{}{deletionPolicyMap})
		}
	}

	return []*schema.ResourceData{data}, err
}

func populateDataFromServer(ctx context.Context, config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, data *schema.ResourceData) error {
	var fullNameList []string

	switch scopedFullnameData.Scope {
	case scope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			stopStatuses := map[dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhase]bool{
				dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseERROR: true,
				dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseREADY: true,
			}

			var dp *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection

			responseStatus := dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePHASEUNSPECIFIED
			isStopStatus := false
			isCtxCallerSet := helper.IsContextCallerSet(ctx)

			for !isStopStatus {
				if isCtxCallerSet || (!isCtxCallerSet && responseStatus != dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePHASEUNSPECIFIED) {
					time.Sleep(5 * time.Second)
				}

				resp, err := config.TMCConnection.DataProtectionService.DataProtectionResourceServiceList(scopedFullnameData.FullnameCluster)
				if err != nil || resp == nil {
					if clienterrors.IsUnauthorizedError(err) {
						authctx.RefreshUserAuthContext(&config, clienterrors.IsUnauthorizedError, err)
						continue
					}

					return errors.Wrap(err, "list data protections")
				}

				if len(resp.DataProtections) == 0 {
					return clienterrors.ErrorWithHTTPCode(http.StatusNotFound, errors.New("data protection not found"))
				}

				dp = resp.DataProtections[0]
				responseStatus = *dp.Status.Phase
				_, isStopStatus = stopStatuses[responseStatus]
			}

			converter := getTFModelConverterCluster()

			err := converter.FillTFSchema(dp, data)
			if err != nil {
				return errors.Wrap(err, "convert api response to schema")
			}

			fullNameList = []string{dp.FullName.ManagementClusterName, dp.FullName.ProvisionerName, dp.FullName.ClusterName}
		}
	case scope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupDataProtectionService.DataProtectionResourceServiceList(scopedFullnameData.FullnameClusterGroup)
			if err != nil || resp == nil {
				return errors.Wrap(err, "list data protections")
			}

			if len(resp.DataProtections) == 0 {
				return clienterrors.ErrorWithHTTPCode(http.StatusNotFound, errors.New("data protection not found"))
			}

			// for cluster group data protection, we don't care about the status
			dp := resp.DataProtections[0]

			converter := getTFModelConverterClusterGroup()

			err = converter.FillTFSchema(dp, data)
			if err != nil {
				return errors.Wrap(err, "convert api response to schema")
			}

			fullNameList = []string{dp.FullName.ClusterGroupName}
		}
	case scope.UnknownScope:
		return errors.New("no valid scope type block found: minimum one valid scope type block is required. Please check the schema.")
	}

	data.SetId(strings.Join(fullNameList, "/"))

	return nil
}

func enableDataProtection(config *authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, data *schema.ResourceData, isUpdate bool) error {
	if config == nil || scopedFullnameData == nil {
		return errors.New("missing variables: error while enabling Tanzu Mission Control cluster data protection feature")
	}

	switch scopedFullnameData.Scope {
	case scope.ClusterScope:
		converter := getTFModelConverterCluster()

		model, err := converter.ConvertTFSchemaToAPIModel(data, []string{})
		if err != nil {
			return errors.Wrap(err, "convert schema to api model")
		}

		request := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest{
			DataProtection: model,
		}

		if request.DataProtection.Spec == nil {
			request.DataProtection.Spec = &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionSpec{}
		}

		fn := config.TMCConnection.DataProtectionService.DataProtectionResourceServiceCreate
		if isUpdate {
			fn = config.TMCConnection.DataProtectionService.DataProtectionResourceServiceUpdate
		}

		_, err = fn(request)
		if err != nil {
			return err
		}
	case scope.ClusterGroupScope:
		converter := getTFModelConverterClusterGroup()

		model, err := converter.ConvertTFSchemaToAPIModel(data, []string{})
		if err != nil {
			return errors.Wrap(err, "convert schema to api model")
		}

		request := &dataprotectioncgmodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionRequest{
			DataProtection: model,
		}

		if request.DataProtection.Spec == nil {
			request.DataProtection.Spec = &dataprotectioncgmodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionSpec{}
		}

		fn := config.TMCConnection.ClusterGroupDataProtectionService.DataProtectionResourceServiceCreate
		if isUpdate {
			fn = config.TMCConnection.ClusterGroupDataProtectionService.DataProtectionResourceServiceUpdate
		}

		_, err = fn(request)
		if err != nil {
			return err
		}
	case scope.UnknownScope:
		return errors.New("no valid scope type block found: minimum one valid scope type block is required. Please check the schema")
	}

	return nil
}
