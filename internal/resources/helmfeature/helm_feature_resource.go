/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmfeature

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmfeature/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmfeature/status"
)

const (
	ResourceName = "tanzu-mission-control_helm_feature"

	statusKey = "status"
)

func ResourceHelm() *schema.Resource {
	return &schema.Resource{
		Schema:        helmSchema,
		CreateContext: resourceHelmCreate,
		ReadContext:   dataSourceHelmRead,
		DeleteContext: resourceHelmDelete,
		CustomizeDiff: schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey, commonscope.ClusterGroupKey})),
	}
}

var helmSchema = map[string]*schema.Schema{
	commonscope.ScopeKey: scope.ScopeSchema,
	common.MetaKey:       common.Meta,
	statusKey:            status.StatusSchema,
}

func resourceHelmCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	scopedFullnameData := scope.ConstructScope(d)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control helm feature entry; Scope full name is empty")
	}

	var (
		UID  string
		meta = common.ConstructMeta(d)
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			helmReq := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest{
				Helm: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
				},
			}

			helmResponse, err := config.TMCConnection.ClusterHelmResourceService.VmwareTanzuManageV1alpha1ClusterHelmResourceServiceCreate(helmReq)
			if err != nil {
				return diag.FromErr(errors.Wrap(err, "Unable to create Tanzu Mission Control cluster helm feature entry"))
			}

			UID = helmResponse.Helm.Meta.UID
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			helmReq := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest{
				Helm: &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
				},
			}

			helmResponse, err := config.TMCConnection.ClusterGroupHelmResourceService.VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceCreate(helmReq)
			if err != nil {
				return diag.FromErr(errors.Wrap(err, "Unable to create Tanzu Mission Control cluster group helm feature entry"))
			}

			UID = helmResponse.Helm.Meta.UID
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}
	// always run
	d.SetId(UID)

	return dataSourceHelmRead(ctx, d, m)
}

func resourceHelmDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	scopedFullnameData := scope.ConstructScope(d)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to delete Tanzu Mission Control helm feature entry; Scope full name is empty")
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			err := config.TMCConnection.ClusterHelmResourceService.VmwareTanzuManageV1alpha1ClusterHelmResourceServiceDelete(scopedFullnameData.FullnameCluster)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrap(err, "Unable to delete Tanzu Mission Control cluster helm feature entry"))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			err := config.TMCConnection.ClusterGroupHelmResourceService.VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceDelete(scopedFullnameData.FullnameClusterGroup)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrap(err, "Unable to delete Tanzu Mission Control cluster group helm feature entry"))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
