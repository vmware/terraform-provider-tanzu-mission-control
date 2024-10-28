/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackageinstall

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/status"
)

func DataSourcePackageInstall() *schema.Resource {
	return &schema.Resource{
		Schema: getDataSourceSchema(),
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataPackageInstallRead(context.WithValue(ctx, contextMethodKey{}, DataSourceRead), d, m)
		},
	}
}

func dataPackageInstallRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageInstallName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package install name")
	}

	packageInstallNamespace, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package install namespace name")
	}

	scopedFullnameData, scopesFound := scope.ConstructScope(d, packageInstallName, packageInstallNamespace)

	if len(scopesFound) == 0 {
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v", strings.Join(scope.ScopesAllowed[:], `, `))
	} else if len(scopesFound) > 1 {
		return diag.Errorf("found scopes: %v are not valid: maximum one valid scope type block is allowed", strings.Join(scopesFound, `, `))
	}

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package install entry; Scope full name is empty")
	}

	UID, meta, repoSpec, clusterScopeStatus, err := retrievePackageInstallUIDMetaAndSpecFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			if ctx.Value(contextMethodKey{}) == DataSourceRead {
				return diag.FromErr(errors.Wrapf(err, "Tanzu Mission Control package install entry not found, name : %s", packageInstallName))
			}

			_ = schema.RemoveFromState(d, m)

			return diags
		}

		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control package install entry, name : %s", packageInstallName))
	}

	d.SetId(UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(status.StatusKey, status.FlattenStatusForClusterScope(clusterScopeStatus)); err != nil {
		return diag.FromErr(err)
	}

	// default path
	pathToInlineValues := scopedFullnameData.FullnameCluster.ManagementClusterName + "/" +
		scopedFullnameData.FullnameCluster.ProvisionerName + "/" +
		scopedFullnameData.FullnameCluster.ClusterName + "/" +
		scopedFullnameData.FullnameCluster.NamespaceName + "/" +
		scopedFullnameData.FullnameCluster.Name + ".yaml"

	existingSpec, ok := d.GetOk(policy.SpecKey)
	if !ok {
		if ctx.Value(contextMethodKey{}) != DataSourceRead {
			return diag.FromErr(fmt.Errorf("spec: %v is not valid: minimum one valid spec block is required", existingSpec))
		}
	} else {
		data, _ := existingSpec.([]interface{})

		if len(data) == 0 || data[0] == nil {
			return diag.FromErr(fmt.Errorf("spec data: %v is not valid: minimum one valid spec data block is required", existingSpec))
		}

		specData := data[0].(map[string]interface{})

		v, ok := specData[spec.PathToInlineValuesKey]
		if ok {
			v1, ok := v.(string)
			if !ok {
				return diag.FromErr(fmt.Errorf("type of path_to_inline_values data: %v is not valid", v1))
			}

			pathToInlineValues = v1
		} else {
			pathToInlineValues = ""
		}
	}

	specValue, err := spec.FlattenSpecForClusterScope(repoSpec, pathToInlineValues)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to write inline values to the file: %s", pathToInlineValues))
	}

	if err := d.Set(spec.SpecKey, specValue); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
