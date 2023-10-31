/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrolemodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1IamPermissionResource Resource Types.
//
//   - RESOURCE_UNSPECIFIED: Unknown.
//   - ORGANIZATION: Organization.
//   - MANAGEMENT_CLUSTER: Management cluster.
//   - PROVISIONER: Provisioner.
//   - CLUSTER_GROUP: Cluster group.
//   - CLUSTER: Cluster.
//   - WORKSPACE: Workspace.
//   - NAMESPACE: Namespace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.iam.permission.Resource
type VmwareTanzuManageV1alpha1IamPermissionResource string

const (

	// VmwareTanzuManageV1alpha1IamPermissionResourceRESOURCEUNSPECIFIED captures enum value "RESOURCE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1IamPermissionResourceRESOURCEUNSPECIFIED VmwareTanzuManageV1alpha1IamPermissionResource = "RESOURCE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1IamPermissionResourceORGANIZATION captures enum value "ORGANIZATION".
	VmwareTanzuManageV1alpha1IamPermissionResourceORGANIZATION VmwareTanzuManageV1alpha1IamPermissionResource = "ORGANIZATION"

	// VmwareTanzuManageV1alpha1IamPermissionResourceMANAGEMENTCLUSTER captures enum value "MANAGEMENT_CLUSTER".
	VmwareTanzuManageV1alpha1IamPermissionResourceMANAGEMENTCLUSTER VmwareTanzuManageV1alpha1IamPermissionResource = "MANAGEMENT_CLUSTER"

	// VmwareTanzuManageV1alpha1IamPermissionResourcePROVISIONER captures enum value "PROVISIONER".
	VmwareTanzuManageV1alpha1IamPermissionResourcePROVISIONER VmwareTanzuManageV1alpha1IamPermissionResource = "PROVISIONER"

	// VmwareTanzuManageV1alpha1IamPermissionResourceCLUSTERGROUP captures enum value "CLUSTER_GROUP".
	VmwareTanzuManageV1alpha1IamPermissionResourceCLUSTERGROUP VmwareTanzuManageV1alpha1IamPermissionResource = "CLUSTER_GROUP"

	// VmwareTanzuManageV1alpha1IamPermissionResourceCLUSTER captures enum value "CLUSTER".
	VmwareTanzuManageV1alpha1IamPermissionResourceCLUSTER VmwareTanzuManageV1alpha1IamPermissionResource = "CLUSTER"

	// VmwareTanzuManageV1alpha1IamPermissionResourceWORKSPACE captures enum value "WORKSPACE".
	VmwareTanzuManageV1alpha1IamPermissionResourceWORKSPACE VmwareTanzuManageV1alpha1IamPermissionResource = "WORKSPACE"

	// VmwareTanzuManageV1alpha1IamPermissionResourceNAMESPACE captures enum value "NAMESPACE".
	VmwareTanzuManageV1alpha1IamPermissionResourceNAMESPACE VmwareTanzuManageV1alpha1IamPermissionResource = "NAMESPACE"
)

// for schema.
var vmwareTanzuManageV1alpha1IamPermissionResourceEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1IamPermissionResource

	if err := json.Unmarshal([]byte(`["RESOURCE_UNSPECIFIED","ORGANIZATION","MANAGEMENT_CLUSTER","PROVISIONER","CLUSTER_GROUP","CLUSTER","WORKSPACE","NAMESPACE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1IamPermissionResourceEnum = append(vmwareTanzuManageV1alpha1IamPermissionResourceEnum, v)
	}
}
