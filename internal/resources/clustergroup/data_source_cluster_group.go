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
	clustergroupmodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/clustergroup"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/common"
)

func DataSourceTMCClusterGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClusterGroupRead,
		Schema:      clusterGroupSchema,
	}
}

func dataSourceClusterGroupRead(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	clusterGroupName, ok := d.Get(clusterGroupName).(string)
	if !ok {
		return diag.Errorf("unable to read cluster group name")
	}

	fn := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
		Name: clusterGroupName,
	}

	resp, err := config.TMCConnection.ClusterGroupResourceService.ManageV1alpha1ClusterGroupResourceServiceGet(fn)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get tanzu TMC cluster group entry, name : %s", clusterGroupName))
	}

	d.SetId(resp.ClusterGroup.Meta.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.ClusterGroup.Meta)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
