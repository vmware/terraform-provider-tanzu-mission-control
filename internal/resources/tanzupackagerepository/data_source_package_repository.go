/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/status"
)

func DataSourcePackageRepository() *schema.Resource {
	return &schema.Resource{
		Schema: getDataSourceSchema(),
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataPackageRepositoryRead(context.WithValue(ctx, contextMethodKey{}, DataSourceRead), d, m)
		},
	}
}

func dataPackageRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read package repository name")
	}

	packageRepositoryNamespace, ok := d.Get(NamespaceKey).(string)
	if !ok {
		return diag.Errorf("unable to read package repository name")
	}

	scopedFullnameData := scope.ConstructScope(d, packageRepositoryName, packageRepositoryNamespace)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control package repository entry; Scope full name is empty")
	}

	UID, meta, repoSpec, clusterScopeStatus, err := retrievePackageRepositoryUIDMetaAndSpecFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			if ctx.Value(contextMethodKey{}) == DataSourceRead {
				return diag.FromErr(errors.Wrapf(err, "Tanzu Mission Control package repository entry not found, name : %s", packageRepositoryName))
			}

			_ = schema.RemoveFromState(d, m)

			return diags
		}

		return diag.FromErr(errors.Wrapf(err, "unable to get Tanzu Mission Control package repository entry, name : %s", packageRepositoryName))
	}

	if err := d.Set(disabledKey, clusterScopeStatus.Disabled); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(status.StateKey, status.FlattenStatusForClusterScope(clusterScopeStatus)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(spec.SpecKey, spec.FlattenSpecForClusterScope(repoSpec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
