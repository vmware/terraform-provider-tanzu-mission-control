/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementclusterregistration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	managementclusterregistrationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/managementclusterregistration"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceManagementClusterRegistration() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceClusterRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: managementClusterRegistrationSchema,
	}
}

func dataSourceClusterRead(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	var (
		diags diag.Diagnostics
		resp  *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterCreateManagementClusterResponse
		err   error
	)

	// TODO - get cluster resource

	if err != nil || resp == nil || resp.ManagementCluster == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster entry, name : %s", d.Get(NameKey)))
	}

	d.SetId(resp.ManagementCluster.Meta.UID)

	status := map[string]interface{}{
		"phase":                resp.ManagementCluster.Status.Phase,
		"health":               resp.ManagementCluster.Status.Health,
		"k8s_version":          resp.ManagementCluster.Status.KubeServerVersion,
		"region":               resp.ManagementCluster.Status.Region,
		"k8s_provider_type":    resp.ManagementCluster.Status.KubernetesProvider.Type,
		"k8s_provider_version": resp.ManagementCluster.Status.KubernetesProvider.Version,
		"infra_provider":       resp.ManagementCluster.Status.InfrastructureProvider,
		"registration_url":     resp.ManagementCluster.Status.RegistrationURL,
		"last_update":          resp.ManagementCluster.Status.LastUpdate,
	}

	if err := d.Set(StatusKey, status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.ManagementCluster.Meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(specKey, flattenSpec(resp.ManagementCluster.Spec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
