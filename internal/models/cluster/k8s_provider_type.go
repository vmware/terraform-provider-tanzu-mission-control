// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustermodel

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType KubernetesProviderType definition - indicates the k8s provider Type.
/*
  - KUBERNETES_PROVIDER_UNSPECIFIED: Unspecified k8s provider (default).
  - VMWARE_TANZU_KUBERNETES_GRID: VMware Tanzu Kubernetes Grid.
  - VMWARE_TANZU_KUBERNETES_GRID_SERVICE: VMware Tanzu Kubernetes Grid Service (Guest Cluster Management).
  - VMWARE_TANZU_KUBERNETES_GRID_HOSTED: VMware Tanzu Kubernetes Grid hosted in TMC.

 swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.KubernetesProviderType
*/
type VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType string

func NewVmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType(value VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType) *VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeKUBERNETESPROVIDERUNSPECIFIED captures enum value "KUBERNETES_PROVIDER_UNSPECIFIED".
	VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeKUBERNETESPROVIDERUNSPECIFIED VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType = "KUBERNETES_PROVIDER_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRID captures enum value "VMWARE_TANZU_KUBERNETES_GRID".
	VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRID VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType = "VMWARE_TANZU_KUBERNETES_GRID"

	// VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRIDSERVICE captures enum value "VMWARE_TANZU_KUBERNETES_GRID_SERVICE".
	VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRIDSERVICE VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType = "VMWARE_TANZU_KUBERNETES_GRID_SERVICE"

	// VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRIDHOSTED captures enum value "VMWARE_TANZU_KUBERNETES_GRID_HOSTED".
	VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRIDHOSTED VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType = "VMWARE_TANZU_KUBERNETES_GRID_HOSTED"
)

// for schema.
var vmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType
	if err := json.Unmarshal([]byte(`["KUBERNETES_PROVIDER_UNSPECIFIED","VMWARE_TANZU_KUBERNETES_GRID","VMWARE_TANZU_KUBERNETES_GRID_SERVICE","VMWARE_TANZU_KUBERNETES_GRID_HOSTED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeEnum = append(vmwareTanzuManageV1alpha1CommonClusterKubernetesProviderTypeEnum, v)
	}
}
