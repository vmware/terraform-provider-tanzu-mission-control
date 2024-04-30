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
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tapeulamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tap/eula"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tap/eula/data"
)

type dataFromServer struct {
	UID  string
	data *tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaData
}

func DataSourceEULA() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceEULAValidate(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: getDataSourceSchema(),
	}
}

func dataSourceEULAValidate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	tapVersion, ok := d.Get(TAPVersionKey).(string)
	if !ok {
		return diag.Errorf("Unable to read TAP version for EULA")
	}

	orgID, ok := d.Get(OrgIDKey).(string)
	if !ok {
		return diag.Errorf("Unable to read organization id for EULA")
	}

	tapEULA := &tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaEula{
		TapVersion: tapVersion,
		OrgID:      orgID,
	}

	eulaDataFromServer, err := retrieveEULADataFromServer(config, tapEULA, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return diags
		}

		return diag.FromErr(err)
	}

	d.SetId(eulaDataFromServer.UID)

	if err := d.Set(dataKey, data.FlattenEULAData(eulaDataFromServer.data)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func retrieveEULADataFromServer(config authctx.TanzuContext, eula *tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaEula, d *schema.ResourceData) (*dataFromServer, error) {
	var eulaDataFromServer = &dataFromServer{}

	resp, err := config.TMCConnection.EulaResourceService.EulaResourceServiceValidate(eula)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			d.SetId("")
			return eulaDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control tap eula entry, version: %s for org: %s", eula.TapVersion, eula.OrgID)
		}

		return eulaDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control tap eula entry, version: %s for org: %s", eula.TapVersion, eula.OrgID)
	}

	eulaDataFromServer.UID = fmt.Sprintf("tap:eula:org:%v:%v", resp.Eula.TapVersion, resp.Eula.OrgID)
	eulaDataFromServer.data = resp.Eula.Data

	if err := d.Set(TAPVersionKey, resp.Eula.TapVersion); err != nil {
		return eulaDataFromServer, err
	}

	orgID, ok := d.Get(OrgIDKey).(string)
	if !ok {
		return eulaDataFromServer, errors.New("Unable to read organization id for EULA")
	}

	if orgID != "" {
		if err := d.Set(OrgIDKey, resp.Eula.OrgID); err != nil {
			return eulaDataFromServer, err
		}
	}

	return eulaDataFromServer, nil
}
