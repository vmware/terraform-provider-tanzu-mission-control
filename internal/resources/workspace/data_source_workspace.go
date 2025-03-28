// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package workspace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceWorkspaceRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: workspaceSchema,
	}
}

func dataSourceWorkspaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	workspaceName, ok := d.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("unable to read workspace name")
	}

	fn := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
		Name: workspaceName,
	}

	resp, err := config.TMCConnection.WorkspaceResourceService.ManageV1alpha1WorkspaceResourceServiceGet(fn)
	if err != nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return
		}

		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control workspace entry, name : %s", workspaceName))
	}

	d.SetId(resp.Workspace.Meta.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Workspace.Meta)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
