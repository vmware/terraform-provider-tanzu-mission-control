package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType The OS disk type of the nodepool.
//
//   - OS_DISK_TYPE_UNSPECIFIED: Unspecified OS disk type.
//   - EPHEMERAL: Ephemeral OS disks are stored only on the host machine, just like a temporary disk. This provides lower read/write latency, along with faster node scaling and cluster upgrades.
//   - MANAGED: Azure replicates the operating system disk for a virtual machine to Azure storage to avoid data loss should the VM need to be relocated to another host
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.OsDiskType
type VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType string

func NewVmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType(value VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType) *VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType.
func (m VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType) Pointer() *VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeOSDISKTYPEUNSPECIFIED captures enum value "OS_DISK_TYPE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeOSDISKTYPEUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType = "OS_DISK_TYPE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeEPHEMERAL captures enum value "EPHEMERAL".
	VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeEPHEMERAL VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType = "EPHEMERAL"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeMANAGED captures enum value "MANAGED".
	VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeMANAGED VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType = "MANAGED"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType
	if err := json.Unmarshal([]byte(`["OS_DISK_TYPE_UNSPECIFIED","EPHEMERAL","MANAGED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeEnum = append(vmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeEnum, v)
	}
}
