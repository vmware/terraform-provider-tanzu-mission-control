/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/integration"
)

func constructFullName(d *schema.ResourceData) (fn *integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) {
	fn = &integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName{}

	fn.ManagementClusterName, _ = d.Get(managementClusterNameKey).(string)
	fn.ProvisionerName, _ = d.Get(provisionerNameKey).(string)
	fn.ClusterName, _ = d.Get(clusterNameKey).(string)
	fn.Name, _ = d.Get(integrationNameKey).(string)

	return fn
}

func constructSpec(d *schema.ResourceData) (spec *integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec) { // nolint:unused
	spec = &integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec{
		Configurations: map[string]interface{}{},
	}

	configs, ok := d.Get(specKey).([]interface{})
	if !ok || len(configs) < 1 {
		return spec
	}

	item, ok := configs[0].(map[string]interface{})
	if !ok {
		return spec
	}

	v, ok := item[configurationKey].(string)
	if !ok {
		return spec
	}

	var m map[string]interface{}
	if err := json.Unmarshal([]byte(v), &m); err == nil {
		spec.Configurations = m
	}

	return spec
}

func flattenSpec(spec *integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec) interface{} {
	flattened := map[string]interface{}{}

	if spec != nil && spec.Configurations != nil {
		flattened[configurationKey] = toJSON(spec.Configurations)
	}

	return []map[string]interface{}{flattened}
}

func flattenStatus(status *integration.VmwareTanzuManageV1alpha1ClusterIntegrationStatus) interface{} {
	return map[string]interface{}{
		"cluster_view_url": status.ClusterViewURL,
		"phase":            status.Phase,
		"tmcAdapter":       status.TmcAdapter,
		"conditions":       toJSON(status.Conditions),
		"operator":         toJSON(status.Operator),
		"workload":         toJSON(status.Workload),
	}
}

func toJSON(data any) string {
	v, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf(`{"error": %q}`, err)
	}

	return string(v)
}
