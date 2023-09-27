/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/
package tkc

import (
	"reflect"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	// tkcmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
)

// func clusterSpecEqual(spec1, spec2 *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec) bool {
// 	return spec1.ClusterGroupName == spec2.ClusterGroupName &&
// 		spec1.ProxyName == spec2.ProxyName &&
// 		spec1.ImageRegistry == spec2.ImageRegistry &&
// 		spec1.TmcManaged == spec2.TmcManaged &&
// 		clusterConfigEqual(spec1.Topology, spec2.Topology)
// }

// func clusterConfigEqual(config1, config2 *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology) bool {
// 	if config1 == nil {
// 		return config2 == nil
// 	}

// 	if config2 == nil {
// 		return false
// 	}

// 	return structEqual(config1.KubernetesNetworkConfig, config2.KubernetesNetworkConfig) &&
// 		structEqual(config1.Logging, config2.Logging) &&
// 		config1.RoleArn == config2.RoleArn &&
// 		mapEqual(config1.Tags, config2.Tags) &&
// 		config1.Version == config2.Version &&
// 		clusterVPCConfigEqual(config1.Vpc, config2.Vpc)
// }

func structEqual[T any](a, b *T) bool {
	if a == nil {
		a = new(T)
	}

	if b == nil {
		b = new(T)
	}

	return reflect.DeepEqual(a, b)
}

// mapEqual handles the cases where one map is nil and the other one is empty.
func mapEqual[K comparable, V any](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	return reflect.DeepEqual(a, b)
}

func setEquality(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	els := map[string]bool{}
	for _, s := range s1 {
		els[s] = true
	}

	for _, s := range s2 {
		if !els[s] {
			return false
		}
	}

	return true
}

func constructStringMap(data map[string]interface{}) map[string]string {
	out := make(map[string]string)

	for k, v := range data {
		var value string

		helper.SetPrimitiveValue(v, &value, k)

		out[k] = value
	}

	return out
}

func constructStringList(data []interface{}) []string {
	out := make([]string, 0, len(data))

	for _, v := range data {
		var value string

		helper.SetPrimitiveValue(v, &value, "")

		out = append(out, value)
	}

	return out
}
