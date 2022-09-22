/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceIntegration() *schema.Resource {
	ri := &resourceIntegration{}

	return &schema.Resource{
		ReadContext: ri.read,
		Schema:      integrationSchema,
	}
}

func (r *resourceIntegration) read(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	fn := constructFullName(d)

	resp, err := r.client(m).ManageV1alpha1ClusterIntegrationResourceServiceRead(fn)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			d.SetId("")
			return
		}

		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control integration entry, name: %s", fn.Name))
	}

	if resp == nil || resp.Integration == nil {
		return diag.Errorf("invalid nil value reading resource %v", fn)
	}

	// Meta
	meta := resp.Integration.Meta
	if meta == nil {
		return diag.Errorf("invalid nil meta value %v", fn)
	}

	d.SetId(meta.UID)

	if err = d.Set(common.MetaKey, common.FlattenMeta(meta)); err != nil {
		return diag.FromErr(err)
	}

	// FullName
	fn = resp.Integration.FullName

	if fn == nil {
		return diag.Errorf("invalid nil full name reading resource %v", fn)
	}

	switch fn.Name {
	case tanzuServiceMeshValue:
	case tanzuObservabilitySaaSValue:
	default:
		return diag.Errorf("Integration Name must be one of %v", validIntegrationNames)
	}

	if err = d.Set(managementClusterNameKey, fn.ManagementClusterName); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set(provisionerNameKey, fn.ProvisionerName); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set(clusterNameKey, fn.ClusterName); err != nil {
		return diag.FromErr(err)
	}

	// Spec
	spec := resp.Integration.Spec
	if spec == nil {
		return diag.Errorf("invalid nil spec reading resource %v", fn)
	}

	if err = d.Set(specKey, flattenSpec(spec)); err != nil {
		return diag.FromErr(err)
	}

	// Status
	if err = d.Set(statusKey, flattenStatus(resp.Integration.Status)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
