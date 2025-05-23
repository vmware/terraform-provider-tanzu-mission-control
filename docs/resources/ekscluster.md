---
Title: "EKS Cluster Resource"
Description: |-
    Create an AWS EKS cluster resource managed by Tanzu Mission Control.
---

# EKS Cluster

The `tanzu-mission-control_ekscluster` resource allows you to provision and manage [AWS EKS](https://aws.amazon.com/eks/) through Tanzu Mission Control.
It allows users to connect Tanzu Mission Control to their Amazon Web Services (AWS) account and create, update/upgrade, and delete EKS clusters and node groups (called node pools in Tanzu).

## Provisioning a EKS Cluster

To use the **Tanzu Mission Control** for creating a new cluster, you must first connect your AWS account to Tanzu Mission Control.
For more information, see [connecting an AWS account for EKS cluster lifecycle management][aws-account]
and [create an EKS Cluster][create-cluster].

You must also have the appropriate permissions in Tanzu Mission Control:

- To provision a cluster, you must have `cluster.admin` permissions.
- You must also have `clustergroup.edit` permissions on the cluster group in which you want to put the new cluster.

[aws-account]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-E4627693-7D1A-4914-A9DF-61E49F97FECC.html
[create-cluster]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-208B2A5A-AE08-4CE1-9DC0-EB573E4BA4A8.html?hWord=N4IghgNiBcIKIGkDKIC+Q

__Note__: Fields under the [nested Schema for `spec.nodepool`](#nestedblock--spec--nodepool) which are markes as "immutable" can't be changed. To update those fields, you need to create a new node pool or rename the node pool (which will have the same effect).

## Example Usage

```terraform
# Create a Tanzu Mission Control AWS EKS cluster entry
resource "tanzu-mission-control_ekscluster" "tf_eks_cluster" {
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
    #proxy		  = "<proxy>"            // Name of TMC Proxy if outbound connection from EKS cluster is via Proxy

    config {
      role_arn = "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com" // Required, forces new

      kubernetes_version = "1.24" // Required
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

      addons_config { // this whole section is optional
        vpc_cni_config {
          eni_config {
            id = "subnet-0a680171b6330619f" // Required, need not belong to the same VPC as the cluster, subnets provided in vpc_cni_config are expected to be in different AZs
            security_groups = [             //optional, if not provided, the cluster security group will be used
              "sg-00c96ad9d02a22522",
            ]
          }
          eni_config {
            id = "subnet-06feb0bb0451cda79" // Required, need not belong to the same VPC as the cluster, subnets provided in vpc_cni_config are expected to be in different AZs
          }
        }
      }
    }

    nodepool {
      info {
        name        = "fist-np"
        description = "tf nodepool description"
      }

      spec {
        role_arn = "arn:aws:iam::000000000000:role/worker.1234567890123467890.eks.tmc.cloud.vmware.com" // Required

        ami_type       = "CUSTOM"
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

        ami_info {
          ami_id                 = "ami-2qu8409oisdfj0qw"
          override_bootstrap_cmd = "#!/bin/bash\n/etc/eks/bootstrap.sh tf2-eks-cluster-2"
        }

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
          max_unavailable_nodes = "4"
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
          name    = "name-of-pre-existing-launch-template"
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

        // This field is only used for cluster UPDATE, to update the ami release version.
        // Do not use this field for cluster CREATE
        release_version = "1.26.6-20230728"
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credential_name` (String) Name of the AWS Credential in Tanzu Mission Control
- `name` (String) Name of this cluster
- `region` (String) AWS Region of this cluster

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- `ready_wait_timeout` (String) Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero
- `spec` (Block List, Max: 1) Spec for the cluster (see [below for nested schema](#nestedblock--spec))
- `wait_for_kubeconfig` (Boolean) Wait until pinniped extension is ready to provide kubeconfig

### Read-Only

- `id` (String) The ID of this resource.
- `kubeconfig` (String) Kubeconfig for connecting to newly created cluster base64 encoded. This will only be returned if you have elected to wait for kubeconfig.
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

- `addons_config` (Block List, Max: 1) Addons config contains the configuration for all the addons of the cluster, which support customization of addon configuration (see [below for nested schema](#nestedblock--spec--config--addons_config))
- `kubernetes_network_config` (Block List, Max: 1) Kubernetes Network Config (see [below for nested schema](#nestedblock--spec--config--kubernetes_network_config))
- `logging` (Block List, Max: 1) EKS logging configuration (see [below for nested schema](#nestedblock--spec--config--logging))
- `tags` (Map of String) The metadata to apply to the cluster to assist with categorization and organization

<a id="nestedblock--spec--config--vpc"></a>
### Nested Schema for `spec.config.vpc`

Required:

- `subnet_ids` (Set of String) Subnet ids used by the cluster (see [Amazon EKS VPC and subnet requirements and considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html#network-requirements-subnets))

Optional:

- `enable_private_access` (Boolean) Enable Kubernetes API requests within your cluster's VPC (such as node to control plane communication) use the private VPC endpoint (see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html))
- `enable_public_access` (Boolean) Enable cluster API server access from the internet. You can, optionally, limit the CIDR blocks that can access the public endpoint using public_access_cidrs (see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html))
- `public_access_cidrs` (Set of String) Specify which addresses from the internet can communicate to the public endpoint, if public endpoint is enabled (see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html))
- `security_groups` (Set of String) Security groups for the cluster VMs


<a id="nestedblock--spec--config--addons_config"></a>
### Nested Schema for `spec.config.addons_config`

Optional:

- `vpc_cni_config` (Block List, Max: 1) VPC CNI addon config contains the configuration for the VPC CNI addon of the cluster (see [below for nested schema](#nestedblock--spec--config--addons_config--vpc_cni_config))

<a id="nestedblock--spec--config--addons_config--vpc_cni_config"></a>
### Nested Schema for `spec.config.addons_config.vpc_cni_config`

Optional:

- `eni_config` (Block List) ENI config is the VPC CNI Elastic Network Interface config for providing the configuration of subnet and security groups for pods in each AZ.  Subnets need not be in the same VPC as the cluster. The subnets provided across eniConfigs should be in different availability zones. Nodepool subnets need to be in the same AZ as the AZs used in ENIConfig.  (see [below for nested schema](#nestedblock--spec--config--addons_config--vpc_cni_config--eni_config))

<a id="nestedblock--spec--config--addons_config--vpc_cni_config--eni_config"></a>
### Nested Schema for `spec.config.addons_config.vpc_cni_config.eni_config`

Required:

- `id` (String) Subnet Id for the pods running in all Nodes in a given AZ.

Optional:

- `security_groups` (Set of String) List of security group is optional and if not provided default security group created by EKS will be used.




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

- `ami_info` (Block List, Max: 1) AMI info for the nodepool if AMI type is specified as CUSTOM (see [below for nested schema](#nestedblock--spec--nodepool--spec--ami_info))
- `ami_type` (String) AMI type, immutable
- `capacity_type` (String) Capacity Type
- `instance_types` (Set of String) Nodepool instance types, immutable
- `launch_template` (Block List, Max: 1) Launch template for the nodepool (see [below for nested schema](#nestedblock--spec--nodepool--spec--launch_template))
- `node_labels` (Map of String) Kubernetes node labels
- `release_version` (String) AMI release version
- `remote_access` (Block List, Max: 1) Remote access to worker nodes, immutable (see [below for nested schema](#nestedblock--spec--nodepool--spec--remote_access))
- `root_disk_size` (Number) Root disk size in GiB, immutable
- `scaling_config` (Block List, Max: 1) Nodepool scaling config (see [below for nested schema](#nestedblock--spec--nodepool--spec--scaling_config))
- `tags` (Map of String) EKS specific tags
- `taints` (Block List) If specified, the node's taints (see [below for nested schema](#nestedblock--spec--nodepool--spec--taints))
- `update_config` (Block List, Max: 1) Update config for the nodepool (see [below for nested schema](#nestedblock--spec--nodepool--spec--update_config))

<a id="nestedblock--spec--nodepool--spec--ami_info"></a>
### Nested Schema for `spec.nodepool.spec.ami_info`

Optional:

- `ami_id` (String) ID of the AMI to be used
- `override_bootstrap_cmd` (String) Override bootstrap command for the custom AMI


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
- `ssh_key` (String) SSH key allows you to connect to your instances and gather diagnostic information if there are issues.


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
