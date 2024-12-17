// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package managementcluster

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	managementclusterregistrationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/managementcluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const defaultWaitTimeout = 15 * time.Minute

func DataSourceManagementClusterRegistration() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceClusterRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: managementClusterRegistrationSchema,
	}
}

func dataSourceClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	var (
		diags diag.Diagnostics
		resp  *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterCreateManagementClusterResponse
		err   error
	)

	getRegistrationResourceRetryableFn := func() (retry bool, err error) {
		resp, err = config.TMCConnection.ManagementClusterRegistrationResourceService.ManagementClusterResourceServiceGet(constructFullname(d))
		if err != nil {
			if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
				_ = schema.RemoveFromState(d, m)

				if helper.IsRefreshState(ctx) {
					return false, nil
				}

				return true, nil
			}

			// refresh auth bearer token if it expired
			authctx.RefreshUserAuthContext(&config, clienterrors.IsUnauthorizedError, err)

			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control management cluster registration entry, name : %s", d.Get(NameKey))
		}

		d.SetId(resp.ManagementCluster.Meta.UID)

		if resp.ManagementCluster.Status == nil || resp.ManagementCluster.Status.Phase == nil {
			return true, errors.Wrapf(err, "Status or Phase not found for Tanzu Mission Control management cluster registration entry, name : %s", d.Get(NameKey))
		}

		if !strings.EqualFold(string(managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterPhaseREADY), string(*resp.ManagementCluster.Status.Phase)) {
			log.Printf("[DEBUG] waiting for management cluster registration(%s) to be in %v phase, present phase:%v", constructFullname(d).ToString(), managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterPhaseREADY, *resp.ManagementCluster.Status.Phase)
			return true, nil
		}

		return false, nil
	}

	timeoutData := d.Get(waitKey).(string)

	if helper.IsDataRead(ctx) {
		timeoutData = helper.DoNotRetry
	}

	if _, ok := d.GetOk(registerClusterKey); !ok {
		timeoutData = helper.DoNotRetry
	}

	var timeoutDuration time.Duration

	switch timeoutData {
	case helper.DoNotRetry:
		_, err = getRegistrationResourceRetryableFn()
	default:
		var parseErr error
		timeoutDuration, parseErr = time.ParseDuration(timeoutData)

		if parseErr != nil {
			log.Printf("[INFO] unable to prase the duration value for the key %s. Defaulting to 15 minutes(15m)"+
				" Please refer to 'https://pkg.go.dev/time#ParseDuration' for providing the right value", waitKey)

			timeoutDuration = defaultWaitTimeout
		}

		_, err = helper.RetryUntilTimeout(getRegistrationResourceRetryableFn, 10*time.Second, timeoutDuration)
	}

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
		"last_update":          resp.ManagementCluster.Status.LastUpdate.String(),
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

	if err := d.Set(NameKey, resp.ManagementCluster.FullName.Name); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
