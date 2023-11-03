//go:build inspections
// +build inspections

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionstests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	inspectionsres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/inspections"
)

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		DataSourcesMap: map[string]*schema.Resource{
			inspectionsres.ResourceNameInspections:       inspectionsres.DataSourceInspections(),
			inspectionsres.ResourceNameInspectionResults: inspectionsres.DataSourceInspectionResults(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}

	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
