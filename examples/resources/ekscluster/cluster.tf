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
