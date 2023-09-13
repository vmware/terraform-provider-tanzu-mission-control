/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"testing"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/integration"
)

const (
	enableNamespaceExclusionsSpecKey = "enableNamespaceExclusions"
	namespaceExclusionsSpecKey       = "namespaceExclusions"
)

func TestFlattenSpec(t *testing.T) {
	type testCase struct {
		given  *integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec
		expect string
	}

	for tcName, tc := range map[string]testCase{
		"simple": {
			given: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec{
				Configurations: map[string]interface{}{},
			},
			expect: `{}`,
		},
		"full": {
			given: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec{
				Configurations: map[string]interface{}{
					enableNamespaceExclusionsSpecKey: true,
					namespaceExclusionsSpecKey: []map[string]interface{}{
						{
							"match": "namespace-1",
							"type":  "EXACT",
						},
						{
							"match": "kube",
							"type":  "START_WITH",
						},
					},
				},
			},
			expect: `{"enableNamespaceExclusions":true,"namespaceExclusions":[{"match":"namespace-1","type":"EXACT"},{"match":"kube","type":"START_WITH"}]}`,
		},
	} {
		got, ok := flattenSpec(tc.given).([]map[string]interface{})
		if !ok {
			t.Errorf("%s: unexpected type for flattened spec: %T", tcName, got)
		}

		conf, ok := got[0][configurationKey].(string)
		if !ok {
			t.Errorf("%s: missing key %q", tcName, configurationKey)
		}

		if conf != tc.expect {
			t.Errorf("%s: wrong configuration, got %s, expected %s", tcName, conf, tc.expect)
		}
	}
}
