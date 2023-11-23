/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterintegration

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	clusterintegrationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/integration/cluster"
	integrationschema "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/integration/schema"
)

var tfModelResourceClusterConverter = &tfModelConverterHelper.TFSchemaModelConverter[*clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegration]{
	TFModelMap: integrationschema.TFModelResourceMap,
}

func ClusterTOIntegrationCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceClusterConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create cluster tanzu observability integration."))
	}

	model.FullName.Name = integrationschema.TanzuObservabilitySaaSValue
	clusterIntegrationFn := model.FullName
	clusterIntegrationRequest := &clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData{
		Integration: model,
	}

	_, err = config.TMCConnection.IntegrationV2ResourceService.ClusterIntegrationResourceServiceCreate(clusterIntegrationRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create cluster tanzu observability integration.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName))
	}

	return ClusterTOIntegrationRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func ClusterTOIntegrationUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceClusterConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update cluster tanzu observability integration."))
	}

	model.FullName.Name = integrationschema.TanzuObservabilitySaaSValue
	clusterIntegrationFn := model.FullName
	clusterIntegrationRequest := &clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData{
		Integration: model,
	}

	_, err = config.TMCConnection.IntegrationV2ResourceService.ClusterIntegrationResourceServiceUpdate(clusterIntegrationRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update cluster tanzu observability integration.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName))
	}

	return ClusterTOIntegrationRead(helper.GetContextWithCaller(ctx, helper.UpdateState), data, m)
}

func ClusterTOIntegrationDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceClusterConverter.ConvertTFSchemaToAPIModel(data, []string{integrationschema.ScopeKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete cluster tanzu observability integration."))
	}

	model.FullName.Name = integrationschema.TanzuObservabilitySaaSValue
	clusterIntegrationFn := model.FullName
	err = config.TMCConnection.IntegrationV2ResourceService.ClusterIntegrationResourceServiceDelete(model.FullName)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete cluster tanzu observability integration.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName))
	}

	return ClusterTOIntegrationRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func ClusterTOIntegrationRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData

	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceClusterConverter.ConvertTFSchemaToAPIModel(data, []string{integrationschema.ScopeKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read cluster tanzu observability integration."))
	}

	model.FullName.Name = integrationschema.TanzuObservabilitySaaSValue
	clusterIntegrationFn := model.FullName
	resp, err = readResourceWait(ctx, &config, clusterIntegrationFn)

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

		return diag.FromErr(errors.Wrapf(err, "Couldn't read cluster tanzu observability integration.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName))
	} else if resp != nil {
		// oldSpecData := data.Get(integrationschema.SpecKey).([]interface{})
		// err = tfModelResourceClusterConverter.FillTFSchema(resp.Integration, data)
		//
		// if err != nil {
		// 	return diag.FromErr(errors.Wrapf(err, "Couldn't read cluster tanzu observability integration.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
		// 		clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName))
		// }
		//
		// // This is necessary to prevent state changes due to configuration values difference.
		// newSpecData := getModifiedSpec(oldSpecData, data.Get(integrationschema.SpecKey).([]interface{}))
		//
		// _ = data.Set(integrationschema.SpecKey, newSpecData)

		fullNameList := []string{clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName, clusterIntegrationFn.Name}

		data.SetId(strings.Join(fullNameList, "/"))
	}

	return diags
}

func ClusterTOIntegrationImporter(scopeID string, ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)
	clusterIntegrationID := scopeID
	clusterIntegrationIDParts := strings.Split(clusterIntegrationID, "/")

	if len(clusterIntegrationIDParts) != 3 {
		return nil, errors.New("Cluster tanzu observability integration ID must be comprised of management_cluster_name, provisioner_name and cluster_name - separated by /")
	}

	clusterIntegrationFn := &clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationFullName{
		ManagementClusterName: clusterIntegrationIDParts[0],
		ProvisionerName:       clusterIntegrationIDParts[1],
		ClusterName:           clusterIntegrationIDParts[2],
		Name:                  integrationschema.TanzuObservabilitySaaSValue,
	}

	resp, err := readResourceWait(ctx, &config, clusterIntegrationFn)

	if err != nil || resp == nil || resp.Integration == nil {
		return nil, errors.Wrapf(err, "Couldn't import cluster tanzu observability integration.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName)
	}

	err = tfModelResourceClusterConverter.FillTFSchema(resp.Integration, data)

	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't import cluster tanzu observability integration.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName)
	}

	fullNameList := []string{clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName, clusterIntegrationFn.Name}

	data.SetId(strings.Join(fullNameList, "/"))

	return []*schema.ResourceData{data}, err
}

func readResourceWait(ctx context.Context, config *authctx.TanzuContext, resourceFullName *clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) (resp *clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationData, err error) {
	stopStatuses := map[clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationPhase]bool{
		clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationPhaseERROR: true,
		clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationPhaseREADY: true,
	}

	responseStatus := clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationPhasePHASEUNSPECIFIED
	_, isStopStatus := stopStatuses[responseStatus]
	isCtxCallerSet := helper.IsContextCallerSet(ctx)

	for !isStopStatus {
		if isCtxCallerSet || (!isCtxCallerSet && responseStatus != clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationPhasePHASEUNSPECIFIED) {
			time.Sleep(5 * time.Second)
		}

		resp, err = config.TMCConnection.IntegrationV2ResourceService.ClusterIntegrationResourceServiceRead(resourceFullName)

		if err != nil || resp == nil || resp.Integration == nil {
			return nil, err
		}

		clusterIntegration := resp.Integration
		responseStatus = *clusterIntegration.Status.Phase
		_, isStopStatus = stopStatuses[responseStatus]
	}

	if responseStatus == clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationPhaseERROR {
		clusterIntegrationFn := resp.Integration.FullName
		err = errors.Errorf("Cluster tanzu observability integration errored.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s",
			clusterIntegrationFn.ManagementClusterName, clusterIntegrationFn.ProvisionerName, clusterIntegrationFn.ClusterName)

		return nil, err
	}

	return resp, err
}

// func getModifiedSpec(oldSpec []interface{}, newSpec []interface{}) interface{} {
// 	oldConfigurations := oldSpec[0].(map[string]interface{})[integrationschema.ConfigurationsKey].(string)
// 	newConfigurations := newSpec[0].(map[string]interface{})[integrationschema.ConfigurationsKey].(string)
//
// 	if oldConfigurations == "" && newConfigurations != "" {
// 		newSpec[0].(map[string]interface{})[integrationschema.ConfigurationsKey] = oldConfigurations
// 	} else if oldConfigurations != "" && newConfigurations != "" {
// 		oldConfigurationsJSON := make(map[string]interface{})
// 		newConfigurationsJSON := make(map[string]interface{})
//
// 		_ = json.Unmarshal([]byte(oldConfigurations), &oldConfigurationsJSON)
// 		_ = json.Unmarshal([]byte(newConfigurations), &newConfigurationsJSON)
//
// 		newSpec[0].(map[string]interface{})[integrationschema.ConfigurationsKey] = modifyConfigurations(oldConfigurationsJSON, newConfigurationsJSON)
// 	}
//
// 	return newSpec
// }
//
// func modifyConfigurations(oldConfigValue interface{}, newConfigValue interface{}) interface{} {
// 	switch typedOldConfigValue := oldConfigValue.(type) {
// 	case map[string]interface{}:
// 		modifiedMap := make(map[string]interface{})
//
// 		for k, v := range typedOldConfigValue {
// 			newValue, keyExists := newConfigValue.(map[string]interface{})[k]
//
// 			if keyExists {
// 				modifiedMap[k] = modifyConfigurations(v, newValue)
// 			}
// 		}
//
// 		return modifiedMap
// 	default:
// 		if !helper.IsEmptyInterface(oldConfigValue) && helper.IsEmptyInterface(typedOldConfigValue) {
// 			return oldConfigValue
// 		}
// 	}
//
// 	return newConfigValue
// }
