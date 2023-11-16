/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotection

import (
	"context"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	dataprotectionscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection/scope"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	dataprotectionmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/dataprotection/cluster/dataprotection"
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
		Schema: enableDataProtectionSchema,
	}
}

func resourceEnableDataProtectionCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	/*model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Tanzu Mission Control data protection configurations."))
	}*/

	scopedFullNameData := dataprotectionscope.ConstructScope(data)

	if scopedFullNameData == nil {
		return diag.Errorf("Unable to enable Tanzu Mission Control Data Protection; Scope full name is empty")
	}

	var (
		UID string
		meta = common.ConstructMeta(data)
	)

	switch scopedFullNameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullNameData.FullnameCluster != nil {
			specVal, err := 
		}
	}

	request := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest{
		DataProtection: model,
	}

	if request.DataProtection.Spec == nil {
		request.DataProtection.Spec = &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionSpec{}
	}

	_, err = config.TMCConnection.DataProtectionService.DataProtectionResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Tanzu Mission Control data protection configurations.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.ClusterName))
	}

	return resourceEnableDataProtectionRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceEnableDataProtectionRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse

	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{ClusterNameKey, ProvisionerNameKey, ManagementClusterNameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read Tanzu Mission Control data protection configurations."))
	}

	dataProtectionFn := model.FullName
	resp, err = readResourceWait(ctx, &config, dataProtectionFn)

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

		return diag.FromErr(errors.Wrapf(err, "Couldn't read data protection configuration for cluster.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			dataProtectionFn.ManagementClusterName, dataProtectionFn.ProvisionerName, dataProtectionFn.ClusterName))
	} else if resp != nil {
		var (
			oldSpecMap map[string]interface{}

			dataProtection = resp.DataProtections[0]
			oldSpec        = data.Get(SpecKey).([]interface{})
		)

		if len(oldSpec) > 0 {
			oldSpecMap = oldSpec[0].(map[string]interface{})

			// Disable restic doesn't return from API.
			if schemaDisableRestic, ok := oldSpecMap[DisableResticKey]; ok {
				dataProtection.Spec.DisableRestic = schemaDisableRestic.(bool)
			}
		} else {
			dataProtection.Spec = nil
		}

		err = tfModelConverter.FillTFSchema(dataProtection, data)

		if err != nil {
			diags = diag.FromErr(err)
		}

		fullNameList := []string{dataProtection.FullName.ManagementClusterName, dataProtection.FullName.ProvisionerName, dataProtection.FullName.ClusterName}

		data.SetId(strings.Join(fullNameList, "/"))
	}

	return diags
}

func resourceEnableDataProtectionDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{ClusterNameKey, ProvisionerNameKey, ManagementClusterNameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete Tanzu Mission Control data protection configurations."))
	}

	dataProtectionFn := model.FullName
	deletionPolicy := data.Get(DeletionPolicyKey).([]interface{})
	deleteBackups := false

	if len(deletionPolicy) > 0 {
		deleteBackups = deletionPolicy[0].(map[string]interface{})[DeleteBackupsKey].(bool)
	}

	err = config.TMCConnection.DataProtectionService.DataProtectionResourceServiceDelete(dataProtectionFn, deleteBackups)

	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete Tanzu Mission Control data protection configurations.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			dataProtectionFn.ManagementClusterName, dataProtectionFn.ProvisionerName, dataProtectionFn.ClusterName))
	}

	return resourceEnableDataProtectionRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceEnableDataProtectionUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update Tanzu Mission Control data protection configurations."))
	}

	request := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest{
		DataProtection: model,
	}

	if request.DataProtection.Spec == nil {
		request.DataProtection.Spec = &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionSpec{}
	}

	_, err = config.TMCConnection.DataProtectionService.DataProtectionResourceServiceUpdate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Tanzu Mission Control data protection configurations.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.ClusterName))
	}

	return resourceEnableDataProtectionRead(helper.GetContextWithCaller(ctx, helper.UpdateState), data, m)
}

func resourceEnableDataProtectionImporter(ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)
	clusterFullName := data.Id()
	clusterFullNameParts := strings.Split(clusterFullName, "/")

	if len(clusterFullNameParts) != 3 {
		return nil, errors.New("Cluster ID must be comprised of management_cluster_name, provisioner_name and cluster_name - separated by /")
	}

	clusterFn := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName{
		ManagementClusterName: clusterFullNameParts[0],
		ProvisionerName:       clusterFullNameParts[1],
		ClusterName:           clusterFullNameParts[2],
	}

	resp, err := readResourceWait(ctx, &config, clusterFn)

	if err != nil || len(resp.DataProtections) == 0 {
		return nil, errors.Wrapf(err, "Couldn't import data protection configuration.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.ClusterName)
	}

	dataProtection := resp.DataProtections[0]
	err = tfModelConverter.FillTFSchema(dataProtection, data)

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

func readResourceWait(ctx context.Context, config *authctx.TanzuContext, resourceFullName *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName) (resp *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse, err error) {
	stopStatuses := map[dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhase]bool{
		dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseERROR: true,
		dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseREADY: true,
	}

	responseStatus := dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePHASEUNSPECIFIED
	_, isStopStatus := stopStatuses[responseStatus]
	isCtxCallerSet := helper.IsContextCallerSet(ctx)

	for !isStopStatus {
		if isCtxCallerSet || (!isCtxCallerSet && responseStatus != dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePHASEUNSPECIFIED) {
			time.Sleep(5 * time.Second)
		}

		resp, err = config.TMCConnection.DataProtectionService.DataProtectionResourceServiceList(resourceFullName)

		if err != nil || resp == nil || resp.DataProtections == nil {
			return nil, err
		}

		dataProtection := resp.DataProtections[0]
		responseStatus = *dataProtection.Status.Phase
		_, isStopStatus = stopStatuses[responseStatus]
	}

	if responseStatus == dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseERROR {
		dataProtectionFn := resp.DataProtections[0].FullName
		err = errors.Errorf("data protection configurations errored.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			dataProtectionFn.ManagementClusterName, dataProtectionFn.ProvisionerName, dataProtectionFn.ClusterName)

		return nil, err
	}

	return resp, err
}
