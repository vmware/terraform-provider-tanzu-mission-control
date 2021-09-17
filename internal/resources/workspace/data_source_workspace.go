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

func DataSourceWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkspaceRead,
		Schema:      workspaceSchema,
	}
}

func dataSourceWorkspaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	workspaceName, ok := d.Get(workspacesName).(string)
	if !ok {
		return diag.Errorf("unable to read workspace name")
	}

	fn := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
		Name: workspaceName,
	}

	resp, err := config.TMCConnection.WorkspaceResourceService.ManageV1alpha1WorkspaceResourceServiceGet(fn)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get tanzu TMC workspace entry, name : %s", workspaceName))
	}

	d.SetId(resp.Workspace.Meta.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Workspace.Meta)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
