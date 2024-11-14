// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package credential

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
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
	)

	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(d, []string{NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to get Tanzu Mission Control credential"))
	}

	getCredentialResourceRetryableFunc := func() (retry bool, err error) {
		resp, err = config.TMCConnection.CredentialResourceService.CredentialResourceServiceGet(model.FullName)
		if err != nil || resp == nil {
			if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
				_ = schema.RemoveFromState(d, m)
				return false, nil
			}

			return true, errors.Wrapf(err, "unable to get Tanzu Mission Control credential entry, name : %s", model.FullName.Name)
		}

		if resp.Credential.Status == nil || resp.Credential.Status.Phase == nil {
			return true, errors.Errorf("status or phase not found for Tanzu Mission Control credential entry, name: %s", model.FullName.Name)
		}

		if *resp.Credential.Status.Phase != credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseVALID &&
			*resp.Credential.Status.Phase != credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseCREATED {
			log.Printf("[DEBUG] waiting for creadential(%s) to be in VALID phase", model.FullName.Name)
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

	specData := d.Get(specKey)
	err = tfModelResourceConverter.FillTFSchema(resp.Credential, d)

	if len(specData.([]interface{})) > 0 {
		_ = d.Set(specKey, specData)
	}

	if err != nil {
		log.Println(err)

		return diag.Errorf("Unable to get credentials. Schema conversion failed.")
	}

	return diags
}
