/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tapeula

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tapeulamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tap/eula"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tap/eula/data"
)

func ResourceEULA() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEULAAccept,
		// The API is idempotent, so update implementation is same as create.
		UpdateContext: resourceEULAAccept,
		ReadContext:   dataSourceEULAValidate,
		DeleteContext: resourceEULADelete,
		Schema:        getResourceSchema(),
	}
}

func getResourceSchema() map[string]*schema.Schema {
	return getTAPEULASchema(false)
}

func getDataSourceSchema() map[string]*schema.Schema {
	return getTAPEULASchema(true)
}

func getTAPEULASchema(isDataSource bool) map[string]*schema.Schema {
	var eulaSchema = map[string]*schema.Schema{
		TAPVersionKey: {
			Type:        schema.TypeString,
			Description: "Version of TAP solution.",
			Required:    true,
			ValidateFunc: validation.All(
				validation.StringLenBetween(1, 126),
				validation.StringIsNotEmpty,
				validation.StringIsNotWhiteSpace,
			),
		},
		OrgIDKey: {
			Type:        schema.TypeString,
			Description: "ID of Organization.",
			Optional:    true,
			ValidateFunc: validation.All(
				validation.StringIsNotEmpty,
				validation.StringIsNotWhiteSpace,
			),
		},
	}

	innerMap := map[string]*schema.Schema{
		dataKey: data.EULAData,
	}

	for key, value := range innerMap {
		if isDataSource {
			eulaSchema[key] = helper.UpdateDataSourceSchema(value)
		} else {
			eulaSchema[key] = value
		}
	}

	return eulaSchema
}

func resourceEULAAccept(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	tapVersion, ok := d.Get(TAPVersionKey).(string)
	if !ok {
		return diag.Errorf("Unable to read TAP version for EULA")
	}

	orgID, ok := d.Get(OrgIDKey).(string)
	if !ok {
		return diag.Errorf("Unable to read organization id for EULA")
	}

	tapEULAReq := &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaRequest{
		Eula: &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaEula{
			TapVersion: tapVersion,
			OrgID:      orgID,
		},
	}

	resp, err := config.TMCConnection.EulaResourceService.EulaResourceServiceAccept(tapEULAReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control tap eula entry, version: %s for org: %s", tapEULAReq.Eula.TapVersion, tapEULAReq.Eula.OrgID))
	}

	UID := fmt.Sprintf("tap:eula:org:%v:%v", resp.Eula.TapVersion, resp.Eula.OrgID)

	d.SetId(UID)

	return dataSourceEULAValidate(ctx, d, m)
}

func resourceEULADelete(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	// TAP EULA deletion is not allowed from TMC backend.
	_ = schema.RemoveFromState(d, m)

	return diags
}
