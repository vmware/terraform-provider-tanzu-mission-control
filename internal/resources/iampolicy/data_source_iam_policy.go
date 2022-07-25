/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/client/errors"
	clustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster"
	clustergroupmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	namespacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/namespace"
	organizationmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/organization"
	workspacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

func DataSourceIAMPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIAMPolicyRead,
		Schema:      iamPolicySchema,
	}
}

func dataSourceIAMPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	var diags diag.Diagnostics

	s, err := getScopeType(d)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "unable to get Role Bindings"))
	}

	fnData := constructScopeFullName(s, d)

	switch s {
	case organization:
		fn, _ := fnData.(*organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName)

		resp, err := config.TMCConnection.OrganizationIAMResourceService.ManageV1alpha1OrganizationIAMPolicyGet(fn)
		if err != nil || resp == nil {
			if clienterrors.IsNotFoundError(err) {
				d.SetId("")
				return diags
			}

			return diag.FromErr(errors.Wrapf(err, "unable to get Role Bindings for organization"))
		}
	case clusterGroup:
		fn, _ := fnData.(*clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName)

		resp, err := config.TMCConnection.ClusterGroupIAMResourceService.ManageV1alpha1ClusterGroupIAMPolicyGet(fn)
		if err != nil || resp == nil {
			if clienterrors.IsNotFoundError(err) {
				d.SetId("")
				return diags
			}

			return diag.FromErr(errors.Wrapf(err, "unable to get Role Bindings for cluster group"))
		}
	case cluster:
		fn, _ := fnData.(*clustermodel.VmwareTanzuManageV1alpha1ClusterFullName)

		resp, err := config.TMCConnection.ClusterIAMResourceService.ManageV1alpha1ClusterIAMPolicyGet(fn)
		if err != nil || resp == nil {
			if clienterrors.IsNotFoundError(err) {
				d.SetId("")
				return diags
			}

			return diag.FromErr(errors.Wrapf(err, "unable to get Role Bindings for cluster"))
		}
	case workspace:
		fn, _ := fnData.(*workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName)

		resp, err := config.TMCConnection.WorkspaceIAMResourceService.ManageV1alpha1WorkspaceIAMPolicyGet(fn)
		if err != nil || resp == nil {
			if clienterrors.IsNotFoundError(err) {
				d.SetId("")
				return diags
			}

			return diag.FromErr(errors.Wrapf(err, "unable to get Role Bindings for workspace"))
		}
	case namespace:
		fn, _ := fnData.(*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName)

		resp, err := config.TMCConnection.NamespaceIAMResourceService.ManageV1alpha1ClusterNamespaceIAMPolicyGet(fn)
		if err != nil || resp == nil {
			if clienterrors.IsNotFoundError(err) {
				d.SetId("")
				return diags
			}

			return diag.FromErr(errors.Wrapf(err, "unable to get Role Bindings for namespace"))
		}
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
