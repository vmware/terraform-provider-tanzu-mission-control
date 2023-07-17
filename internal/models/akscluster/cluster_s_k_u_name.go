package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterClusterSKUName Name options of cluster SKU.
//
//   - NAME_UNSPECIFIED: Unspecified name.
//   - BASIC: Basic option for the AKS control plane.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.ClusterSKU.Name
type VmwareTanzuManageV1alpha1AksclusterClusterSKUName string

func NewVmwareTanzuManageV1alpha1AksclusterClusterSKUName(value VmwareTanzuManageV1alpha1AksclusterClusterSKUName) *VmwareTanzuManageV1alpha1AksclusterClusterSKUName {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterClusterSKUName.
func (m VmwareTanzuManageV1alpha1AksclusterClusterSKUName) Pointer() *VmwareTanzuManageV1alpha1AksclusterClusterSKUName {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterClusterSKUNameNAMEUNSPECIFIED captures enum value "NAME_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterClusterSKUNameNAMEUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterClusterSKUName = "NAME_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterClusterSKUNameBASIC captures enum value "BASIC".
	VmwareTanzuManageV1alpha1AksclusterClusterSKUNameBASIC VmwareTanzuManageV1alpha1AksclusterClusterSKUName = "BASIC"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterClusterSKUNameEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterClusterSKUName
	if err := json.Unmarshal([]byte(`["NAME_UNSPECIFIED","BASIC"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterClusterSKUNameEnum = append(vmwareTanzuManageV1alpha1AksclusterClusterSKUNameEnum, v)
	}
}
