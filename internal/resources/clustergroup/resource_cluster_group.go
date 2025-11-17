// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func ResourceClusterGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClusterGroupCreate,
		ReadContext:   dataSourceClusterGroupRead,
		UpdateContext: resourceClusterGroupInPlaceUpdate,
		DeleteContext: resourceClusterGroupDelete,
		Schema:        clusterGroupSchema,
	}
}

var clusterGroupSchema = map[string]*schema.Schema{
	NameKey: {
		Type:     schema.TypeString,
		ForceNew: true,
		Required: true,
	},
	common.MetaKey: common.Meta,
}

func resourceClusterGroupInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	updateRequired := common.HasMetaChanged(d)

	if !updateRequired {
		return diags
	}

	clusterGroupName, ok := d.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("unable to read cluster group name")
	}

	fn := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
		Name: clusterGroupName,
	}

	getResp, err := config.TMCConnection.ClusterGroupResourceService.ManageV1alpha1ClusterGroupResourceServiceGet(fn)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get tanzu cluster group entry, name : %s", clusterGroupName))
	}

	if updateRequired {
		meta := common.ConstructMeta(d)

		if value, ok := getResp.ClusterGroup.Meta.Labels[common.CreatorLabelKey]; ok {
			meta.Labels[common.CreatorLabelKey] = value
		}

		getResp.ClusterGroup.Meta.Labels = meta.Labels
		getResp.ClusterGroup.Meta.Description = meta.Description
	}

	_, err = config.TMCConnection.ClusterGroupResourceService.ManageV1alpha1ClusterGroupResourceServiceUpdate(
		&clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupRequest{
			ClusterGroup: getResp.ClusterGroup,
		},
	)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update tanzu TMC cluster group entry, name : %s", clusterGroupName))
	}

	return dataSourceClusterGroupRead(ctx, d, m)
}

func resourceClusterGroupDelete(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	clusterGroupName, _ := d.Get(NameKey).(string)

	fn := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
		Name: clusterGroupName,
	}

	err := config.TMCConnection.ClusterGroupResourceService.ManageV1alpha1ClusterGroupResourceServiceDelete(fn)
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group entry, name : %s", clusterGroupName))
	}

	_ = schema.RemoveFromState(d, m)

	return diags
}

func resourceClusterGroupCreate(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	clusterGroupName, ok := d.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("unable to read cluster group name")
	}

	clusterGroupRequest := &clustergroupmodel.VmwareTanzuManageV1alpha1ClusterGroupRequest{
		ClusterGroup: &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupClusterGroup{
			FullName: &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
				Name: clusterGroupName,
			},
			Meta: common.ConstructMeta(d),
		},
	}

	clusterGroupResponse, err := config.TMCConnection.ClusterGroupResourceService.ManageV1alpha1ClusterGroupResourceServiceCreate(clusterGroupRequest)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster group entry, name : %s", clusterGroupName))
	}

	d.SetId(clusterGroupResponse.ClusterGroup.Meta.UID)

	return diags
}
