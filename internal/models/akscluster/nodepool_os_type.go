package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolOsType The operation system type options of the nodepool.
//
//   - OS_TYPE_UNSPECIFIED: Unspecified OS type.
//   - LINUX: The linux operation system.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.OsType
type VmwareTanzuManageV1alpha1AksclusterNodepoolOsType string

func NewVmwareTanzuManageV1alpha1AksclusterNodepoolOsType(value VmwareTanzuManageV1alpha1AksclusterNodepoolOsType) *VmwareTanzuManageV1alpha1AksclusterNodepoolOsType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterNodepoolOsType.
func (m VmwareTanzuManageV1alpha1AksclusterNodepoolOsType) Pointer() *VmwareTanzuManageV1alpha1AksclusterNodepoolOsType {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeOSTYPEUNSPECIFIED captures enum value "OS_TYPE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeOSTYPEUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterNodepoolOsType = "OS_TYPE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeLINUX captures enum value "LINUX".
	VmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeLINUX VmwareTanzuManageV1alpha1AksclusterNodepoolOsType = "LINUX"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterNodepoolOsType
	if err := json.Unmarshal([]byte(`["OS_TYPE_UNSPECIFIED","LINUX"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeEnum = append(vmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeEnum, v)
	}
}
