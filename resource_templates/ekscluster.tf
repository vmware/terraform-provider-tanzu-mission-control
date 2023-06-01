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
      }
    }
  }
}
