/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupintegration

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
	clustergroupintegrationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/integration/clustergroup"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	integrationschema "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/integration/schema"
)

var tfModelResourceClusterGroupConverter = &tfModelConverterHelper.TFSchemaModelConverter[*clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationIntegration]{
	TFModelMap: integrationschema.TFModelResourceMap,
}

func ClusterGroupTOIntegrationCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceClusterGroupConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create cluster group tanzu observability integration."))
	}

	model.FullName.Name = integrationschema.TanzuObservabilitySaaSValue
	clusterGroupIntegrationFn := model.FullName
	clusterGroupIntegrationRequest := &clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData{
		Integration: model,
	}

	_, err = config.TMCConnection.IntegrationV2ResourceService.ClusterGroupIntegrationResourceServiceCreate(clusterGroupIntegrationRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create cluster group tanzu observability integration.\nCluster Group Name: %s", clusterGroupIntegrationFn.ClusterGroupName))
	}

	return ClusterGroupTOIntegrationRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func ClusterGroupTOIntegrationUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return diag.FromErr(errors.New("Update of cluster group tanzu observability integration is not supported."))
}

func ClusterGroupTOIntegrationDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceClusterGroupConverter.ConvertTFSchemaToAPIModel(data, []string{integrationschema.ScopeKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete cluster group tanzu observability integration."))
	}

	model.FullName.Name = integrationschema.TanzuObservabilitySaaSValue
	clusterGroupIntegrationFn := model.FullName
	err = config.TMCConnection.IntegrationV2ResourceService.ClusterGroupIntegrationResourceServiceDelete(model.FullName)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete cluster group tanzu observability integration.\nCluster Group Name: %s", clusterGroupIntegrationFn.ClusterGroupName))
	}

	return ClusterGroupTOIntegrationRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func ClusterGroupTOIntegrationRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData

	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceClusterGroupConverter.ConvertTFSchemaToAPIModel(data, []string{integrationschema.ScopeKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read cluster group tanzu observability integration."))
	}

	model.FullName.Name = integrationschema.TanzuObservabilitySaaSValue
	clusterGroupIntegrationFn := model.FullName
	resp, err = readResourceWait(ctx, &config, clusterGroupIntegrationFn)

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

		return diag.FromErr(errors.Wrapf(err, "Couldn't read cluster group tanzu observability integration.\nCluster Group Name: %s", clusterGroupIntegrationFn.ClusterGroupName))
	} else if resp != nil {
		fullNameList := []string{clusterGroupIntegrationFn.ClusterGroupName, clusterGroupIntegrationFn.Name}

		data.SetId(strings.Join(fullNameList, "/"))
	}

	return diags
}

func ClusterGroupTOIntegrationImporter(scopeID string, ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)
	clusterGroupIntegrationFn := &clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationFullName{
		ClusterGroupName: scopeID,
		Name:             integrationschema.TanzuObservabilitySaaSValue,
	}

	resp, err := readResourceWait(ctx, &config, clusterGroupIntegrationFn)

	if err != nil || resp == nil || resp.Integration == nil {
		return nil, errors.Wrapf(err, "Couldn't import cluster group tanzu observability integration.\nCluster Group Name: %s", clusterGroupIntegrationFn.ClusterGroupName)
	}

	err = tfModelResourceClusterGroupConverter.FillTFSchema(resp.Integration, data)

	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't import cluster group tanzu observability integration.\nCluster Group Name: %s", clusterGroupIntegrationFn.ClusterGroupName)
	}

	fullNameList := []string{clusterGroupIntegrationFn.ClusterGroupName, clusterGroupIntegrationFn.Name}

	data.SetId(strings.Join(fullNameList, "/"))

	return []*schema.ResourceData{data}, err
}

func readResourceWait(ctx context.Context, config *authctx.TanzuContext, resourceFullName *clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationFullName) (resp *clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationData, err error) {
	stopStatuses := map[statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhase]bool{
		statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseERROR:   true,
		statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED: true,
	}

	responseStatus := statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhasePHASEUNSPECIFIED
	_, isStopStatus := stopStatuses[responseStatus]
	isCtxCallerSet := helper.IsContextCallerSet(ctx)

	for !isStopStatus {
		if isCtxCallerSet || (!isCtxCallerSet && responseStatus != statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhasePHASEUNSPECIFIED) {
			time.Sleep(5 * time.Second)
		}

		resp, err = config.TMCConnection.IntegrationV2ResourceService.ClusterGroupIntegrationResourceServiceRead(resourceFullName)

		if err != nil || resp == nil || resp.Integration == nil {
			return nil, err
		}

		clusterIntegration := resp.Integration
		responseStatus = *clusterIntegration.Status.Phase
		_, isStopStatus = stopStatuses[responseStatus]
	}

	if responseStatus == statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseERROR {
		clusterGroupIntegrationFn := resp.Integration.FullName
		err = errors.Errorf("Cluster group tanzu observability integration errored.\nCluster Group Name: %s", clusterGroupIntegrationFn.ClusterGroupName)

		return nil, err
	}

	return resp, err
}
