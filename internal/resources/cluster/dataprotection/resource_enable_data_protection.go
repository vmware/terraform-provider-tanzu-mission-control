/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotection

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	dataprotectionmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/dataprotection"
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
	model := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})
	request := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest{
		DataProtection: model,
	}

	if request.DataProtection.Spec == nil {
		request.DataProtection.Spec = &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionSpec{}
	}

	_, err := config.TMCConnection.DataProtectionService.DataProtectionResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control data protection configurations, cluster: %s/%s/%s",
			model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.ClusterName))
	}

	return resourceEnableDataProtectionRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceEnableDataProtectionRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var (
		resp *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse
		err  error
	)

	resourceFullName := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{ClusterNameKey, ProvisionerNameKey, ManagementClusterNameKey}).FullName
	resp, err = readResourceWait(config, resourceFullName)

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

		return diag.Errorf("Couldn't find data protection configuration for cluster: %s/%s/%s",
			resourceFullName.ManagementClusterName, resourceFullName.ProvisionerName, resourceFullName.ClusterName)
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
	resourceFullName := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{ClusterNameKey, ProvisionerNameKey, ManagementClusterNameKey}).FullName
	deletionPolicy := data.Get(DeletionPolicyKey).([]interface{})
	deleteBackups := false

	if len(deletionPolicy) > 0 {
		deleteBackups = deletionPolicy[0].(map[string]interface{})[DeleteBackupsKey].(bool)
	}

	err := config.TMCConnection.DataProtectionService.DataProtectionResourceServiceDelete(resourceFullName, deleteBackups)

	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control data protection configurations, cluster: %s/%s/%s",
			resourceFullName.ManagementClusterName, resourceFullName.ProvisionerName, resourceFullName.ClusterName))
	}

	return resourceEnableDataProtectionRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceEnableDataProtectionUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})
	request := &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest{
		DataProtection: model,
	}

	if request.DataProtection.Spec == nil {
		request.DataProtection.Spec = &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionSpec{}
	}

	_, err := config.TMCConnection.DataProtectionService.DataProtectionResourceServiceUpdate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control data protection configurations, cluster: %s/%s/%s",
			model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.ClusterName))
	}

	return resourceEnableDataProtectionRead(helper.GetContextWithCaller(ctx, helper.UpdateState), data, m)
}

func resourceEnableDataProtectionImporter(_ context.Context, data *schema.ResourceData, config any) ([]*schema.ResourceData, error) {
	tc, ok := config.(authctx.TanzuContext)

	if !ok {
		return nil, errors.New("error while retrieving Tanzu auth config")
	}

	clusterFullName := data.Id()
	clusterFullNameParts := strings.Split(clusterFullName, "/")

	if len(clusterFullNameParts) != 3 {
		return nil, errors.New("Cluster ID must be comprised of management_cluster_name, provisioner_name and cluster_name - separated by /")
	}

	clusterFn := dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName{
		ManagementClusterName: clusterFullNameParts[0],
		ProvisionerName:       clusterFullNameParts[1],
		ClusterName:           clusterFullNameParts[2],
	}

	resp, err := tc.TMCConnection.DataProtectionService.DataProtectionResourceServiceList(&clusterFn)

	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get data protection configuration for cluster: %s/%s/%s", clusterFn.ManagementClusterName,
			clusterFn.ProvisionerName, clusterFn.ClusterName)
	}

	if len(resp.DataProtections) == 0 {
		err = errors.Errorf("Couldn't find data protection configuration for cluster: %s/%s/%s", clusterFn.ManagementClusterName,
			clusterFn.ProvisionerName, clusterFn.ClusterName)
	} else {
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
	}

	return []*schema.ResourceData{data}, err
}

func readResourceWait(config authctx.TanzuContext, resourceFullName *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName) (resp *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse, err error) {
	responseStatus := dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePHASEUNSPECIFIED

	for responseStatus != dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseERROR {
		resp, err = config.TMCConnection.DataProtectionService.DataProtectionResourceServiceList(resourceFullName)

		if err != nil || resp == nil || resp.DataProtections == nil {
			return nil, err
		}

		dataProtection := resp.DataProtections[0]
		responseStatus = *dataProtection.Status.Phase

		if responseStatus == dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseREADY {
			break
		} else {
			time.Sleep(5 * time.Second)
		}
	}

	if responseStatus == dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseERROR {
		err = errors.Errorf("data protection configurations errored for cluster: %s/%s/%s",
			resourceFullName.ManagementClusterName, resourceFullName.ProvisionerName, resourceFullName.ClusterName)

		return nil, err
	}

	return resp, err
}
