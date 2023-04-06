/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	secretexportclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster/secretexport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceSecret() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceSecretRead(context.WithValue(ctx, contextMethodKey{}, DataSourceRead), d, m)
		},
		Schema: getDataSourceSchema(),
	}
}

func dataSourceSecretRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	var (
		diags            diag.Diagnostics
		secretResp       *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse
		secretexportResp *secretexportclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretExportResponse
		err              error
	)

	secretName, _ := d.Get(NameKey).(string)

	secretResp, err = config.TMCConnection.SecretResourceService.SecretResourceServiceGet(constructFullname(d))
	if err != nil || secretResp == nil {
		if clienterrors.IsNotFoundError(err) {
			if ctx.Value(contextMethodKey{}) == DataSourceRead {
				return diag.FromErr(errors.Wrapf(err, "Tanzu Mission Control cluster secret entry not found, name : %s", d.Get(NameKey)))
			}

			_ = schema.RemoveFromState(d, m)

			return diags
		}

		return diag.FromErr(errors.Wrapf(err, "unable to get Tanzu Mission Control secret entry, name : %s", secretName))
	}

	var password string

	if _, ok := d.GetOk(specKey); ok {
		password, _ = (d.Get(helper.GetFirstElementOf(specKey, DockerConfigjsonKey, PasswordKey))).(string)
	}

	status := map[string]interface{}{
		SecretPhaseKey: secretResp.Secret.Status.Conditions[Ready].Reason,
	}

	secretexportResp, err = config.TMCConnection.SecretExportResourceService.SecretExportResourceServiceGet(constructFullnameSecetExport(d))

	if d.Get(ExportKey).(bool) {
		if err != nil || secretexportResp == nil {
			switch {
			case clienterrors.IsNotFoundError(err):
				if err := d.Set(ExportKey, false); err != nil {
					return diag.FromErr(err)
				}
			default:
				return diag.FromErr(errors.Wrapf(err, "unable to get Tanzu Mission Control secret export entry, name : %s", secretName))
			}
		} else {
			status[SecretExportPhaseKey] = secretexportResp.SecretExport.Status.Conditions[Ready].Reason
		}
	} else {
		if err == nil && secretexportResp != nil {
			if err := d.Set(ExportKey, true); err != nil {
				return diag.FromErr(err)
			}
			status[SecretExportPhaseKey] = secretexportResp.SecretExport.Status.Conditions[Ready].Reason
		}
	}

	d.SetId(secretResp.Secret.Meta.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(secretResp.Secret.Meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(statusKey, status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(specKey, flattenSpec(secretResp.Secret.Spec, password)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
