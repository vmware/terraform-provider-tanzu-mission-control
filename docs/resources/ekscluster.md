---
Title: "EKS Cluster Resource"
Description: |-
    Create an AWS EKS cluster resource managed by Tanzu Mission Control.
---

# EKS Cluster

The `tanzu_mission_control_ekscluster` resource allows you to provision and manage [AWS EKS](https://aws.amazon.com/eks/) through Tanzu Mission Control.
It allows users to connect Tanzu Mission Control to their Amazon Web Services (AWS) account and create, update/upgrade, and delete EKS clusters and node groups (called node pools in Tanzu).

## Provisioning a EKS Cluster

To use the **Tanzu Mission Control** for creating a new cluster, you must first connect your AWS account to Tanzu Mission Control.
For more information, see [connecting an AWS account for EKS cluster lifecycle management][aws-account]
and [create an EKS Cluster][create-cluster].

You must also have the appropriate permissions in Tanzu Mission Control:

- To provision a cluster, you must have `cluster.admin` permissions.
- You must also have `clustergroup.edit` permissions on the cluster group in which you want to put the new cluster.

[aws-account]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-E4627693-7D1A-4914-A9DF-61E49F97FECC.html
[create-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-208B2A5A-AE08-4CE1-9DC0-EB573E4BA4A8.html?hWord=N4IghgNiBcIKIGkDKIC+Q

__Note__: Fields under the [nested Schema for `spec.nodepool`](#nestedblock--spec--nodepool) which are markes as "immutable" can't be changed. To update those fields, you need to create a new nodepool or rename the nodepool (which will have the same effect).

## Example Usage

```terraform
# Create a Tanzu Mission Control AWS EKS cluster entry
resource "tanzu_mission_control_ekscluster" "tf_eks_cluster" {
  credential_name = "eks-test"          // Required
  region          = "us-west-2"         // Required
  name            = "tf2-eks-cluster-2" // Required

  ready_wait_timeout = "30m" // Wait time for cluster operations to finish (default: 30m).

  meta {
    description = "eks test cluster"
    labels      = { "key1" : "value1" }
  }

  spec {
    cluster_group = "test-cluster-group" // Default: default
    #proxy		  = "<proxy>"              // Proxy if used

    config {
      role_arn = "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com" // Required, forces new

      kubernetes_version = "1.23" // Required
      tags               = { "tagkey" : "tagvalue" }

      kubernetes_network_config {
        service_cidr = "10.100.0.0/16" // Forces new
      }

      logging {
        api_server         = false
        audit              = true
        authenticator      = true
        controller_manager = false
        scheduler          = true
      }

      vpc { // Required
        enable_private_access = true
        enable_public_access  = true
        public_access_cidrs = [
          "0.0.0.0/0",
        ]
        security_groups = [ // Forces new
          "sg-0a6768722e9716768",
        ]
        subnet_ids = [ // Forces new
          "subnet-0a184f6302af32a86",
          "subnet-0ed95d5c212ac62a1",
          "subnet-0526ecaecde5b1bf7",
          "subnet-06897e1063cc0cf4e",
        ]
      }
    }

    nodepool {
      info {
        name        = "fist-np"
        description = "tf nodepool description"
      }

      spec {
        role_arn       = "arn:aws:iam::000000000000:role/worker.1234567890123467890.eks.tmc.cloud.vmware.com" // Required

        ami_type       = "AL2_x86_64"
        capacity_type  = "ON_DEMAND"
        root_disk_size = 40 // Default: 20GiB
        tags           = { "nptag" : "nptagvalue9" }
        node_labels    = { "nplabelkey" : "nplabelvalue" }

        subnet_ids = [ // Required
          "subnet-0a184f9301ae39a86",
          "subnet-0b495d7c212fc92a1",
          "subnet-0c86ec9ecde7b9bf7",
          "subnet-06497e6063c209f4d",
        ]

        remote_access {
          ssh_key = "test-key" // Required (if remote access is specified)

          security_groups = [
            "sg-0a6768722e9716768",
          ]
        }

        scaling_config {
          desired_size = 4
          max_size     = 8
          min_size     = 1
        }

        update_config {
          max_unavailable_nodes = "10"
        }

        instance_types = [
          "t3.medium",
          "m3.large"
        ]

      }
    }

    nodepool {
      info {
        name        = "second-np"
        description = "tf nodepool 2 description"
      }

      spec {
        role_arn    = "arn:aws:iam::000000000000:role/worker.1234567890123467890.eks.tmc.cloud.vmware.com" // Required
        tags        = { "nptag" : "nptagvalue7" }
        node_labels = { "nplabelkey" : "nplabelvalue" }

        subnet_ids = [ // Required
          "subnet-0a184f9301ae39a86",
          "subnet-0b495d7c212fc92a1",
          "subnet-0c86ec9ecde7b9bf7",
          "subnet-06497e6063c209f4d",
        ]

        launch_template {
          name    = "vivek"
          version = "7"
        }

        scaling_config {
          desired_size = 4
          max_size     = 8
          min_size     = 1
        }

        update_config {
          max_unavailable_percentage = "12"
        }

        taints {
          effect = "PREFER_NO_SCHEDULE"
          key    = "randomkey"
          value  = "randomvalue"
        }
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credential_name` (String) Name of the AWS Crendential in Tanzu Mission Control
- `name` (String) Name of this cluster
- `region` (String) AWS Region of the this cluster

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- `ready_wait_timeout` (String) Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero
- `spec` (Block List, Max: 1) Spec for the cluster (see [below for nested schema](#nestedblock--spec))

### Read-Only

- `id` (String) The ID of this resource.
- `status` (Map of String) Status of the cluster

<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `config` (Block List, Min: 1, Max: 1) EKS config for the cluster control plane (see [below for nested schema](#nestedblock--spec--config))
- `nodepool` (Block List, Min: 1) Nodepool definitions for the cluster (see [below for nested schema](#nestedblock--spec--nodepool))

Optional:

- `cluster_group` (String) Name of the cluster group to which this cluster belongs
- `proxy` (String) Optional proxy name is the name of the Proxy Config to be used for the cluster

<a id="nestedblock--spec--config"></a>
### Nested Schema for `spec.config`

Required:

- `kubernetes_version` (String) Kubernetes version of the cluster
- `role_arn` (String) ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations
- `vpc` (Block List, Min: 1, Max: 1) VPC config (see [below for nested schema](#nestedblock--spec--config--vpc))

Optional:

- `kubernetes_network_config` (Block List, Max: 1) Kubernetes Network Config (see [below for nested schema](#nestedblock--spec--config--kubernetes_network_config))
- `logging` (Block List, Max: 1) EKS logging configuration (see [below for nested schema](#nestedblock--spec--config--logging))
- `tags` (Map of String) The metadata to apply to the cluster to assist with categorization and organization

<a id="nestedblock--spec--config--vpc"></a>
### Nested Schema for `spec.config.vpc`

Required:

- `subnet_ids` (Set of String) Subnet ids used by the cluster

Optional:

- `enable_private_access` (Boolean) Enable private access on the cluster
- `enable_public_access` (Boolean) Enable public access on the cluster
- `public_access_cidrs` (Set of String) Public access cidrs
- `security_groups` (Set of String) Security groups for the cluster VMs


<a id="nestedblock--spec--config--kubernetes_network_config"></a>
### Nested Schema for `spec.config.kubernetes_network_config`

Required:

- `service_cidr` (String) Service CIDR for Kubernetes services


<a id="nestedblock--spec--config--logging"></a>
### Nested Schema for `spec.config.logging`

Optional:

- `api_server` (Boolean) Enable API server logs
- `audit` (Boolean) Enable audit logs
- `authenticator` (Boolean) Enable authenticator logs
- `controller_manager` (Boolean) Enable controller manager logs
- `scheduler` (Boolean) Enable scheduler logs



<a id="nestedblock--spec--nodepool"></a>
### Nested Schema for `spec.nodepool`

Required:

- `info` (Block List, Min: 1, Max: 1) Info for the nodepool (see [below for nested schema](#nestedblock--spec--nodepool--info))
- `spec` (Block List, Min: 1, Max: 1) Spec for the cluster (see [below for nested schema](#nestedblock--spec--nodepool--spec))

<a id="nestedblock--spec--nodepool--info"></a>
### Nested Schema for `spec.nodepool.info`

Required:

- `name` (String) Name of the nodepool, immutable

Optional:

- `description` (String) Description for the nodepool


<a id="nestedblock--spec--nodepool--spec"></a>
### Nested Schema for `spec.nodepool.spec`

Required:

- `role_arn` (String) ARN of the IAM role that provides permissions for the Kubernetes nodepool to make calls to AWS API operations, immutable
- `subnet_ids` (Set of String) Subnets required for the nodepool

Optional:

- `ami_type` (String) AMI Type, immutable
- `capacity_type` (String) Capacity Type
- `instance_types` (Set of String) Nodepool instance types, immutable
- `launch_template` (Block List, Max: 1) Launch template for the nodepool (see [below for nested schema](#nestedblock--spec--nodepool--spec--launch_template))
- `node_labels` (Map of String) Kubernetes node labels
- `remote_access` (Block List, Max: 1) Remote access to worker nodes, immutable (see [below for nested schema](#nestedblock--spec--nodepool--spec--remote_access))
- `root_disk_size` (Number) Root disk size in GiB, immutable
- `scaling_config` (Block List, Max: 1) Nodepool scaling config (see [below for nested schema](#nestedblock--spec--nodepool--spec--scaling_config))
- `tags` (Map of String) EKS specific tags
- `taints` (Block List) If specified, the node's taints (see [below for nested schema](#nestedblock--spec--nodepool--spec--taints))
- `update_config` (Block List, Max: 1) Update config for the nodepool (see [below for nested schema](#nestedblock--spec--nodepool--spec--update_config))

<a id="nestedblock--spec--nodepool--spec--launch_template"></a>
### Nested Schema for `spec.nodepool.spec.launch_template`

Optional:

- `id` (String) The ID of the launch template
- `name` (String) The name of the launch template
- `version` (String) The version of the launch template to use


<a id="nestedblock--spec--nodepool--spec--remote_access"></a>
### Nested Schema for `spec.nodepool.spec.remote_access`

Optional:

- `security_groups` (Set of String) Security groups for the VMs
- `ssh_key` (String) SSH key for the nodepool VMs


<a id="nestedblock--spec--nodepool--spec--scaling_config"></a>
### Nested Schema for `spec.nodepool.spec.scaling_config`

Optional:

- `desired_size` (Number) Desired size of nodepool
- `max_size` (Number) Maximum size of nodepool
- `min_size` (Number) Minimum size of nodepool


<a id="nestedblock--spec--nodepool--spec--taints"></a>
### Nested Schema for `spec.nodepool.spec.taints`

Optional:

- `effect` (String) Current effect state of the node pool
- `key` (String) The taint key to be applied to a node
- `value` (String) The taint value corresponding to the taint key


<a id="nestedblock--spec--nodepool--spec--update_config"></a>
### Nested Schema for `spec.nodepool.spec.update_config`

Optional:

- `max_unavailable_nodes` (String) Maximum number of nodes unavailable at once during a version update
- `max_unavailable_percentage` (String) Maximum percentage of nodes unavailable during a version update
