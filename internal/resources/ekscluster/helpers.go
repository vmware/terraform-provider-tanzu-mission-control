/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/
package ekscluster

import (
	"reflect"

	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
)

func nodepoolSpecEqual(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec, spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) bool {
	return spec1.AmiType == spec2.AmiType &&
		spec1.CapacityType == spec2.CapacityType &&
		setEquality(spec1.InstanceTypes, spec2.InstanceTypes) &&
		structEqual(spec1.LaunchTemplate, spec2.LaunchTemplate) &&
		reflect.DeepEqual(spec1.NodeLabels, spec2.NodeLabels) &&
		nodepoolRemoteAccessEqual(spec1.RemoteAccess, spec2.RemoteAccess) &&
		spec1.RoleArn == spec2.RoleArn &&
		spec1.RootDiskSize == spec2.RootDiskSize &&
		structEqual(spec1.ScalingConfig, spec2.ScalingConfig) &&
		setEquality(spec1.SubnetIds, spec2.SubnetIds) &&
		reflect.DeepEqual(spec1.Tags, spec2.Tags) &&
		nodepoolTaintsEqual(spec1.Taints, spec2.Taints) &&
		structEqual(spec1.UpdateConfig, spec2.UpdateConfig)
}

func clusterSpecEqual(spec1, spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) bool {
	return spec1.ClusterGroupName == spec2.ClusterGroupName &&
		spec1.ProxyName == spec2.ProxyName &&
		clusterConfigEqual(spec1.Config, spec2.Config)
}

func clusterConfigEqual(config1, config2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig) bool {
	if config1 == nil {
		return config2 == nil
	}

	if config2 == nil {
		return false
	}

	return structEqual(config1.KubernetesNetworkConfig, config2.KubernetesNetworkConfig) &&
		structEqual(config1.Logging, config2.Logging) &&
		config1.RoleArn == config2.RoleArn &&
		reflect.DeepEqual(config1.Tags, config2.Tags) &&
		config1.Version == config2.Version &&
		clusterVPCConfigEqual(config1.Vpc, config2.Vpc)
}

func clusterVPCConfigEqual(vpc1, vpc2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterVPCConfig) bool {
	if vpc1 == nil {
		return vpc2 == nil
	}

	if vpc2 == nil {
		return false
	}

	return vpc1.EnablePrivateAccess == vpc2.EnablePrivateAccess &&
		vpc1.EnablePublicAccess == vpc2.EnablePublicAccess &&
		setEquality(vpc1.PublicAccessCidrs, vpc2.PublicAccessCidrs) &&
		setEquality(vpc1.SecurityGroups, vpc2.SecurityGroups) &&
		setEquality(vpc1.SubnetIds, vpc2.SubnetIds)
}

func structEqual[T any](a, b *T) bool {
	if a == nil {
		a = new(T)
	}

	if b == nil {
		b = new(T)
	}

	return reflect.DeepEqual(a, b)
}

func nodepoolTaintsEqual(taints1, taints2 []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint) bool {
	if len(taints1) != len(taints2) {
		return false
	}

	taints := map[string]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{}
	for _, t := range taints1 {
		taints[t.Key] = t
	}

	for _, t := range taints2 {
		if !reflect.DeepEqual(t, taints[t.Key]) {
			return false
		}
	}

	return true
}

func nodepoolRemoteAccessEqual(ra1, ra2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess) bool {
	if ra1 == nil {
		ra1 = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess{}
	}

	if ra2 == nil {
		ra2 = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess{}
	}

	return ra1.SSHKey == ra2.SSHKey &&
		setEquality(ra1.SecurityGroups, ra2.SecurityGroups)
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
