/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceClusterGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceClusterGroupRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: clusterGroupSchema,
	}
}

func dataSourceClusterGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	clusterGroupName, ok := d.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("unable to read cluster group name")
	}

	fn := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
		Name: clusterGroupName,
	}

	resp, err := config.TMCConnection.ClusterGroupResourceService.ManageV1alpha1ClusterGroupResourceServiceGet(fn)
	if err != nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return
		}

		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group entry, name : %s", clusterGroupName))
	}

	d.SetId(resp.ClusterGroup.Meta.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.ClusterGroup.Meta)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
