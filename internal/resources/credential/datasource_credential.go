/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package credential

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
	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const defaultWaitTimeout = 3 * time.Minute

func DataSourceCredential() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceCredentialRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: credentialSchema,
	}
}

func dataSourceCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	var (
		diags diag.Diagnostics
		resp  *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialGetCredentialResponse
		err   error
	)

	name, _ := d.Get(NameKey).(string)

	getCredentialResourceRetryableFunc := func() (retry bool, err error) {
		resp, err = config.TMCConnection.CredentialResourceService.CredentialResourceServiceGet(constructFullname(d))
		if err != nil || resp == nil {
			if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
				_ = schema.RemoveFromState(d, m)
				return false, nil
			}

			return true, errors.Wrapf(err, "unable to get Tanzu Mission Control credential entry, name : %s", name)
		}

		if resp.Credential.Status == nil || resp.Credential.Status.Phase == nil {
			return true, errors.Errorf("status or phase not found for Tanzu Mission Control credential entry, name: %s", name)
		}

		if !strings.EqualFold(string(*resp.Credential.Status.Phase), string(credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseVALID)) {
			log.Printf("[DEBUG] waiting for creadential(%s) to be in VALID phase", name)
			return true, nil
		}

		return false, nil
	}

	timeoutData := d.Get(waitKey).(string)
	if !helper.IsCreateState(ctx) {
		timeoutData = helper.DoNotRetry
	}

	var timeoutDuration time.Duration

	switch timeoutData {
	case helper.DoNotRetry:
		_, err = getCredentialResourceRetryableFunc()
	default:
		var parseErr error
		timeoutDuration, parseErr = time.ParseDuration(timeoutData)

		if parseErr != nil {
			log.Printf("[INFO] unable to parse the duration value for the key %s. Defaulting to %s minutes"+
				" Please refer to 'https://pkg.go.dev/time#ParseDuration' for providing the right value", waitKey, defaultWaitTimeout)

			timeoutDuration = defaultWaitTimeout
		}

		_, err = helper.RetryUntilTimeout(getCredentialResourceRetryableFunc, 10*time.Second, timeoutDuration)
	}

	if err != nil || resp == nil || resp.Credential == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control credential entry, name : %s", d.Get(NameKey)))
	}

	d.SetId(resp.Credential.Meta.UID)

	status := map[string]interface{}{
		"phase":      resp.Credential.Status.Phase,
		"phase_info": resp.Credential.Status.PhaseInfo,
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Credential.Meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(statusKey, status); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
