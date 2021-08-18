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

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
	workspacemodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/workspace"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/common"
)

const workspacesName = "workspace_name"

func ResourceWorkspace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkspaceCreate,
		ReadContext:   schema.NoopContext,
		UpdateContext: schema.NoopContext,
		DeleteContext: schema.NoopContext,
		Schema:        workspaceSchema,
	}
}

var workspaceSchema = map[string]*schema.Schema{
	workspacesName: {
		Type:     schema.TypeString,
		Required: true,
	},
	common.MetaKey: common.Meta,
}

func resourceWorkspaceCreate(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var workspaceName = d.Get(workspacesName).(string)

	workspaceRequest := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceRequest{
		Workspace: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceWorkspace{
			FullName: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
				Name: workspaceName,
			},
			Meta: common.ConstructMeta(d),
		},
	}

	workspaceResponse, err := config.TMCConnection.WorkspaceResourceService.ManageV1alpha1WorkspaceResourceServiceCreate(workspaceRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create tanzu TMC workspace entry, name : %s", workspaceName))
	}

	d.SetId(workspaceResponse.Workspace.Meta.UID)

	return diags
}
