// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustermodel

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider InfrastructureProvider definition - indicates the cluster infrastructure provider.
/*
  - INFRASTRUCTURE_PROVIDER_UNSPECIFIED: Unspecified infrastructure provider (default).
  - INFRASTRUCTURE_PROVIDER_NONE: No cloud provider (likely bare metal).
  - AWS_EC2: AmazonWeb Services EC2.
  - GCP_GCE: Google Cloud.
  - AZURE_COMPUTE: Azure Compute.
  - VMWARE_VSPHERE: VMWare vSphere.
  - OPENSHIFT: OpenShift.

 swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.InfrastructureProvider
*/
type VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider string

func NewVmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider(value VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider) *VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderINFRASTRUCTUREPROVIDERUNSPECIFIED captures enum value "INFRASTRUCTURE_PROVIDER_UNSPECIFIED".
	VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderINFRASTRUCTUREPROVIDERUNSPECIFIED VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider = "INFRASTRUCTURE_PROVIDER_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderINFRASTRUCTUREPROVIDERNONE captures enum value "INFRASTRUCTURE_PROVIDER_NONE".
	VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderINFRASTRUCTUREPROVIDERNONE VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider = "INFRASTRUCTURE_PROVIDER_NONE"

	// VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderAWSEC2 captures enum value "AWS_EC2".
	VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderAWSEC2 VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider = "AWS_EC2"

	// VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderGCPGCE captures enum value "GCP_GCE".
	VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderGCPGCE VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider = "GCP_GCE"

	// VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderAZURECOMPUTE captures enum value "AZURE_COMPUTE".
	VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderAZURECOMPUTE VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider = "AZURE_COMPUTE"

	// VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderVMWAREVSPHERE captures enum value "VMWARE_VSPHERE".
	VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderVMWAREVSPHERE VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider = "VMWARE_VSPHERE"

	// VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderOPENSHIFT captures enum value "OPENSHIFT".
	VmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderOPENSHIFT VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider = "OPENSHIFT"
)

// for schema.
var vmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider
	if err := json.Unmarshal([]byte(`["INFRASTRUCTURE_PROVIDER_UNSPECIFIED","INFRASTRUCTURE_PROVIDER_NONE","AWS_EC2","GCP_GCE","AZURE_COMPUTE","VMWARE_VSPHERE","OPENSHIFT"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderEnum = append(vmwareTanzuManageV1alpha1CommonClusterInfrastructureProviderEnum, v)
	}
}
