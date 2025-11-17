// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package provisioner

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	provisioner "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/provisioner"
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
	provisionerKey: {
		Type:        schema.TypeList,
		Description: "Provisioners info",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
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
			},
		},
	},
}

func dataSourceProvisionerRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	id := make([]string, 0)

	var resp *provisioner.VmwareTanzuManageV1alpha1ManagementclusterProvisionerListprovisionersResponse

	model, err := tfModelDataConverter.ConvertTFSchemaToAPIModel(d, []string{provisionerKey, nameKey, managementClusterNameKey})
	if err != nil || model == nil || model.Provisioners == nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read Tanzu Mission Control provisioner configurations."))
	}

	if model.Provisioners[0].FullName.Name == "" {
		resp, err = config.TMCConnection.ProvisionerResourceService.ProvisionerResourceServiceList(model.Provisioners[0].FullName)
		if err != nil {
			if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
				_ = schema.RemoveFromState(d, m)
				return diags
			}

			return diags
		}
	} else {
		getResp, err := config.TMCConnection.ProvisionerResourceService.ProvisionerResourceServiceGet(model.Provisioners[0].FullName)
		if err != nil {
			if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
				_ = schema.RemoveFromState(d, m)
				return diags
			}

			return diags
		}

		p := &provisioner.VmwareTanzuManageV1alpha1ManagementclusterProvisionerListprovisionersResponse{
			Provisioners: []*provisioner.VmwareTanzuManageV1alpha1ManagementclusterProvisionerProvisioner{
				{
					FullName: &provisioner.VmwareTanzuManageV1alpha1ManagementclusterProvisionerFullName{
						ManagementClusterName: getResp.Provisioner.FullName.ManagementClusterName,
						Name:                  getResp.Provisioner.FullName.Name,
						OrgID:                 getResp.Provisioner.FullName.OrgID,
					},
					Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
						Description:     getResp.Provisioner.Meta.Description,
						Labels:          getResp.Provisioner.Meta.Labels,
						UID:             getResp.Provisioner.Meta.UID,
						ResourceVersion: getResp.Provisioner.Meta.ResourceVersion,
					},
				},
			},
		}

		resp = p
	}

	err = tfModelDataConverter.FillTFSchema(resp, d)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to populate tf schema"))
	}

	for _, prov := range resp.Provisioners {
		id = append(id, prov.Meta.UID)
	}

	d.SetId(strings.Join(id, "_"))

	return diags
}
