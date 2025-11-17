// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package targetlocation

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	targetlocationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/targetlocation"
)

func DataSourceTargetLocations() *schema.Resource {
	// Unpack resource map to datasource map.
	constructTFModelDataSourceResponseMap()

	return &schema.Resource{
		ReadContext: dataSourceTargetLocationsRead,
		Schema:      backupTargetLocationDataSourceSchema,
	}
}

func dataSourceTargetLocationsRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse

	config := m.(authctx.TanzuContext)

	request, err := tfModelDataSourceRequestConverter.ConvertTFSchemaToAPIModel(data, []string{})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Couldn't read Tanzu Mission Control backup target location."))
	}

	if request.SearchScope.ClusterName == "" {
		request.SearchScope.ProviderName = TMCProviderName
	}

	resp, err = config.TMCConnection.TargetLocationService.TargetLocationResourceServiceList(request)

	switch {
	case err != nil:
		return diag.Errorf("Couldn't list target locations")
	case resp.BackupLocations == nil:
		data.SetId("NO_DATA")
	default:
		err = tfModelDataSourceResponseConverter.FillTFSchema(resp, data)
		if err != nil {
			return diag.FromErr(err)
		}

		if request.SearchScope.ClusterName != "" {
			idKeys := []string{
				request.SearchScope.ManagementClusterName,
				request.SearchScope.ProvisionerName,
				request.SearchScope.ClusterName,
				request.SearchScope.Name,
			}
			data.SetId(fmt.Sprintf("cluster_scope/%s", strings.Join(idKeys, "/")))
		} else {
			idKeys := []string{
				request.SearchScope.ProviderName,
				request.SearchScope.CredentialName,
				request.SearchScope.AssignedGroupName,
				request.SearchScope.Name,
			}
			data.SetId(fmt.Sprintf("provider_scope/%s", strings.Join(idKeys, "/")))
		}
	}

	return diags
}
