/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package testing

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
)

func TestPreCheck(t *testing.T) func() {
	return func() {
		for _, env := range []string{authctx.ServerEndpointEnvVar, authctx.CSPTokenEnvVar, authctx.CSPEndpointEnvVar} {
			require.NotEmpty(t, os.Getenv(env))
		}
	}
}

func GetTestProviderFactories(provider *schema.Provider) map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		providerName: func() (*schema.Provider, error) {
			if provider == nil {
				return provider, errors.New("provider cannot be nil")
			}

			return provider, nil
		},
	}
}
