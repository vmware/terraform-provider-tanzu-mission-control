---
Title: "Provisioning of a AWS EKS cluster"
Description: |-
    An example of provisioning AWS EKS clusters.
---
# EKS Cluster

The `tanzu-mission-control_ekscluster` resource can directly perform cluster lifecycle management operations on EKS clusters (and associated node groups) including create, update, upgrade, and delete through Tanzu Mission Control.

## Prerequisites

To manage the lifecycle of EKS clusters, you need the following prerequisites.

- Onboard the AWS account onto Tanzu Mission Control under which the EKS Clusters will be provisioned. Please refer [connecting an AWS account for EKS cluster lifecycle management][aws-account] guide for detailed steps. You can also use `tanzu-mission-control_credential` Terraform resource for this purpose. The name of the EKS AWS credential in Tanzu Mission Control will be reffered to as `credential_name` in this guide.

- Create the required VPC with subnets in the desired region. Please refer to Tanzu documentation on how to [create a VPC with Subnets for EKS cluster lifecycle management][tanzu-vpc-guide] and AWS guide for [creating a VPC for your Amazon EKS cluster][aws-vpc-guied].

- Ensure the CSP token used in initialising the terraform provider has the right set of permissions to create a workload cluster.

[tanzu-vpc-guide]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-5708F04E-7EA3-495D-A484-FD6DB7AA8356.html
[aws-vpc-guide]: https://docs.aws.amazon.com/eks/latest/userguide/creating-a-vpc.html

## Provisioning the cluster

You can use the following template as reference to write your own `tanzu-mission-control_ekscluster` resource in the terraform scripts. 

```terraform
// Tanzu Mission Control EKS Cluster Type: AWS EKS clusters.
// Operations supported : Read, Create, Update & Delete

// Read Tanzu Mission Control AWS EKS cluster : fetch cluster details
data "tanzu-mission-control_ekscluster" "tf_eks_cluster" {
  credential_name = "<aws-credential-name>" // Required
  region          = "<aws-region>"          // Required
  name            = "<cluster-name>"        // Required
}

// Create Tanzu Mission Control AWS EKS cluster entry
resource "tanzu-mission-control_ekscluster" "tf_eks_cluster" {
  credential_name = "<aws-credential-name>" // Required
  region          = "<aws-region>"          // Required
  name            = "<cluster-name>"        // Required

  ready_wait_timeout = "<time>" // Wait time for cluster operations to finish (default: 30m).

  meta {
    description = "description of the cluster"
    labels      = { "<key>" : "<value>" }
  }

  spec {
    cluster_group = "test-cluster-group" // Default: default
    proxy         = "<proxy>"

    config {
      role_arn           = "<aws-control-plane-role-arn>" // Required, forces new
      kubernetes_version = "<k8s-version>"                // Required
      tags               = { "<key>" : "<value>" }

      kubernetes_network_config {
        service_cidr = "<service-cidr-block>" // Forces new
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
          "<cidr-blocks>",
        ]
        security_groups = [ // Forces new
          "<security-group-ids>",
        ]
        subnet_ids = [ // Forces new
          "<subnet-ids>",
        ]
      }
    }

    nodepool {
      info {
        name        = "<nodepool-name>" // Required
        description = "description of node pool"
      }

      spec {
        // Refer to nodepool's schema
        role_arn       = "<aws-nodepool-role-arn>" // Required
        ami_type       = "<ami-type>"
        capacity_type  = "<capacity-type>"
        root_disk_size = 40 // In GiB, default: 20GiB
        tags           = { "<key>" : "<value>" }
        node_labels    = { "<key>" : "<value>" }

        subnet_ids = [ // Required
          "<subnet-ids>",
        ]

        ami_info {
          ami_id                 = "<aws-ami-id>"
          override_bootstrap_cmd = "<ami-bootstrap-command>"
        }

        remote_access {
          ssh_key = "<aws-ssh-key-name>" // Required (if remote access is specified)

          security_groups = [
            "<security-group-ids>",
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
          "<instance-types>",
        ]

      }
    }

    nodepool {
      info {
        name        = "<nodepool-name>" // Required
        description = "description of node pool"
      }

      spec {
        role_arn    = "<aws-nodepool-role-arn>" // Required
        tags        = { "<key>" : "<value>" }
        node_labels = { "<key>" : "<value>" }

        subnet_ids = [ // Required
          "<subnet-ids>",
        ]

        launch_template {
          name    = "<launch-template-name>"
          version = "<launch-template-version>"
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
          effect = "<taint-effect>"
          key    = "<taint-key>"
          value  = "<taint-value>"
        }

        // This field is only used for cluster UPDATE, to update the ami release version.
        // Do not use this field for cluster CREATE
        release_version = "<ami_release_version>"
      }
    }
  }
}
```
