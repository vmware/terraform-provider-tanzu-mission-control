// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package integration

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	integrationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/integration"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/integration"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func ResourceIntegration(alternateClients ...integrationclient.ClientService) *schema.Resource {
	ri := &resourceIntegration{alternateClients: append(alternateClients, &defaultClient{})}

	return &schema.Resource{
		CreateContext: ri.create,
		ReadContext:   ri.read,
		UpdateContext: ri.update,
		DeleteContext: ri.integrationDelete,
		Schema:        integrationSchema,
	}
}

type resourceIntegration struct {
	alternateClients []integrationclient.ClientService
}

func (r *resourceIntegration) create(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	fn := constructFullName(d)

	req := &integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest{
		Integration: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationIntegration{
			FullName: fn,
			Meta:     common.ConstructMeta(d),
			Spec:     constructSpec(d),
		},
	}

	resp, err := r.client(m).ManageV1alpha1ClusterIntegrationResourceServiceCreate(req)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster integration entry, name : %s", fn.Name))
	}

	if resp == nil || resp.Integration == nil {
		return diag.Errorf("invalid nil-response creating resource %v", fn)
	}

	d.SetId(resp.Integration.Meta.UID)

	return append(diags, r.read(ctx, d, m)...)
}

func (r *resourceIntegration) update(context.Context, *schema.ResourceData, interface{}) (diags diag.Diagnostics) {
	return diag.FromErr(errors.New("update of Tanzu Mission Control integration is not supported"))
}

func (r *resourceIntegration) integrationDelete(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	fn := constructFullName(d)

	err := r.client(m).ManageV1alpha1ClusterIntegrationResourceServiceDelete(fn)
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control integration entry, name: %s", fn.Name))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	getIntegrationResourceRetryable := func() (retry bool, err error) {
		if _, err = r.client(m).ManageV1alpha1ClusterIntegrationResourceServiceRead(fn); err == nil {
			return true, errors.New("integration deletion in progress")
		}

		if !clienterrors.IsNotFoundError(err) {
			return true, err
		}

		return false, nil
	}

	if _, err = helper.Retry(getIntegrationResourceRetryable, 10*time.Second, 18); err != nil {
		diag.FromErr(errors.Wrapf(err, "verify %s cluster integration resource clean up", fn.Name))
	}

	return diags
}

func (r *resourceIntegration) client(m interface{}) integrationclient.ClientService {
	switch c := m.(type) {
	case authctx.TanzuContext:
		return c.TMCConnection.IntegrationResourceService
	case integrationclient.ClientService:
		return c
	default:
		if len(r.alternateClients) == 0 {
			return nil
		}

		return r.alternateClients[0]
	}
}
