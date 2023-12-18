/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provisioner

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceProvisioner() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceProvisionerRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: provisionerListSchema,
	}
}

var provisionerListSchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the provisioner",
		Optional:    true,
	},
	managementClusterNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the management cluster",
		Required:    true,
		ForceNew:    true,
	},
	orgIDKey: {
		Type:        schema.TypeString,
		Description: "ID of the organization",
		Optional:    true,
	},
	common.MetaKey: common.Meta,
}

func dataSourceProvisionerRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(d, []string{nameKey, managementClusterNameKey})
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read Tanzu Mission Control provisioner configurations."))
	}

	resp, err := config.TMCConnection.ProvisionerResourceService.ProvisionerResourceServiceList(model.FullName)
	if err != nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return
		}
	}

	for i := range resp.Provisioners {
		d.SetId(resp.Provisioners[i].Meta.UID)

		if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Provisioners[i].Meta)); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}
