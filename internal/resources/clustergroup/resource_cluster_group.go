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

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
	clienterrors "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/errors"
	clustergroupmodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/clustergroup"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/common"
)

const clusterGroupName = "name"

func ResourceClusterGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClusterGroupCreate,
		ReadContext:   dataSourceClusterGroupRead,
		UpdateContext: schema.NoopContext,
		DeleteContext: resourceClusterGroupDelete,
		Schema:        clusterGroupSchema,
	}
}

var clusterGroupSchema = map[string]*schema.Schema{
	clusterGroupName: {
		Type:     schema.TypeString,
		ForceNew: true,
		Required: true,
	},
	common.MetaKey: common.Meta,
}

func resourceClusterGroupDelete(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	clusterGroupName, _ := d.Get(clusterGroupName).(string)

	fn := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
		Name: clusterGroupName,
	}

	err := config.TMCConnection.ClusterGroupResourceService.ManageV1alpha1ClusterGroupResourceServiceDelete(fn)
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete tanzu TMC cluster group entry, name : %s", clusterGroupName))
	}

	d.SetId("")

	return diags
}

func resourceClusterGroupCreate(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	clusterGroupName, ok := d.Get(clusterGroupName).(string)
	if !ok {
		return diag.Errorf("unable to read cluster group name")
	}

	clusterGroupRequest := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupRequest{
		ClusterGroup: &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupClusterGroup{
			FullName: &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
				Name: clusterGroupName,
			},
			Meta: common.ConstructMeta(d),
		},
	}

	clusterGroupResponse, err := config.TMCConnection.ClusterGroupResourceService.ManageV1alpha1ClusterGroupResourceServiceCreate(clusterGroupRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create tanzu TMC cluster group entry, name : %s", clusterGroupName))
	}

	d.SetId(clusterGroupResponse.ClusterGroup.Meta.UID)

	return diags
}
