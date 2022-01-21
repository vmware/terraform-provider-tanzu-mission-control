/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/client/errors"
	workspacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/workspace"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	workspacesName = "name"
	ResourceName   = "tmc_workspace"
)

func ResourceWorkspace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkspaceCreate,
		ReadContext:   dataSourceWorkspaceRead,
		UpdateContext: resourceWorkspaceInPlaceUpdate,
		DeleteContext: resourceWorkspaceDelete,
		Schema:        workspaceSchema,
	}
}

var workspaceSchema = map[string]*schema.Schema{
	workspacesName: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	common.MetaKey: common.Meta,
}

func resourceWorkspaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	workspaceName, _ := d.Get(workspacesName).(string)

	fn := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
		Name: workspaceName,
	}

	err := config.TMCConnection.WorkspaceResourceService.ManageV1alpha1WorkspaceResourceServiceDelete(fn)
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control workspace entry, name : %s", workspaceName))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func resourceWorkspaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var workspaceName, _ = d.Get(workspacesName).(string)

	workspaceRequest := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest{
		Workspace: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceWorkspace{
			FullName: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
				Name: workspaceName,
			},
			Meta: common.ConstructMeta(d),
		},
	}

	workspaceResponse, err := config.TMCConnection.WorkspaceResourceService.ManageV1alpha1WorkspaceResourceServiceCreate(workspaceRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control workspace entry, name : %s", workspaceName))
	}

	d.SetId(workspaceResponse.Workspace.Meta.UID)

	return dataSourceWorkspaceRead(ctx, d, m)
}

func resourceWorkspaceInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	updateRequired := common.HasMetaChanged(d)

	if !updateRequired {
		return diags
	}

	workspaceName, ok := d.Get(workspacesName).(string)
	if !ok {
		return diag.Errorf("unable to read workspace name")
	}

	fn := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
		Name: workspaceName,
	}

	getResp, err := config.TMCConnection.WorkspaceResourceService.ManageV1alpha1WorkspaceResourceServiceGet(fn)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control wrokspace entry, name : %s", workspaceName))
	}

	if updateRequired {
		meta := common.ConstructMeta(d)

		if value, ok := getResp.Workspace.Meta.Labels[common.CreatorLabelKey]; ok {
			meta.Labels[common.CreatorLabelKey] = value
		}

		getResp.Workspace.Meta.Labels = meta.Labels
		getResp.Workspace.Meta.Description = meta.Description
	}

	_, err = config.TMCConnection.WorkspaceResourceService.ManageV1alpha1WorkspaceResourceServiceUpdate(
		&workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest{
			Workspace: getResp.Workspace,
		},
	)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control workspace entry, name : %s", workspaceName))
	}

	return dataSourceWorkspaceRead(ctx, d, m)
}
