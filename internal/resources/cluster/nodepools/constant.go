/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepools

const (
	clusterNameKey              = "cluster_name"
	nodePoolNameKey             = "name"
	workerNodeCountKey          = "worker_node_count"
	cloudLabelsKey              = "cloud_labels"
	nodeLabelsKey               = "node_labels"
	tkgAWSKey                   = "tkg_aws"
	tkgServiceVsphereKey        = "tkg_service_vsphere"
	tkgVsphereKey               = "tkg_vsphere"
	classKey                    = "class"
	storageClassKey             = "storage_class"
	failureDomainKey            = "failure_domain"
	volumesKey                  = "volumes"
	capacityKey                 = "capacity"
	mountPathKey                = "mount_path"
	volumeNameKey               = "name"
	pvcStorageClassKey          = "pvc_storage_class"
	vmConfigKey                 = "vm_config"
	cpuKey                      = "cpu"
	diskKey                     = "disk_size"
	memoryKey                   = "memory"
	nodepoolAvailabilityZoneKey = "nodepool_availability_zone"
	nodepoolInstanceTypeKey     = "nodepool_instance_type"
	nodePlacementKey            = "node_placement"
	privateSubnetIDKey          = "private_subnet_id"
	nodepoolVersionKey          = "nodepool_version"
	awsAvailabilityZoneKey      = "aws_availability_zone"
	managementClusterNameKey    = "management_cluster_name"
	provisionerNameKey          = "provisioner_name"
	specKey                     = "spec"
	statusKey                   = "status"
	ready                       = "Ready"
	waitKey                     = "ready_wait_timeout"
)
