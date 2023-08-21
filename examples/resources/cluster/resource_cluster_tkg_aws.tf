# Create a Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_aws_cluster" {
  management_cluster_name = "tkgm-aws-terraform" // Default: attached
  provisioner_name        = "default"            // Default: attached
  name                    = "tkgm-aws-workload"  // Required

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" // Default: default
    tkg_aws {
      advanced_configs {
        key   = "AWS_SECURITY_GROUP_BASTION"
        value = "sg-01376425482384"
      }
      settings {
        network {
          cluster {
            pods {
              cidr_blocks = "100.96.0.0/11" // Required
            }

            services {
              cidr_blocks = "100.64.0.0/13" // Required
            }
          }
          provider {
            subnets {
              availability_zone = "us-west-2a"
              cidr_block_subnet = "10.0.1.0/24"
              is_public         = true
            }
            subnets {
              availability_zone = "us-west-2a"
              cidr_block_subnet = "10.0.0.0/24"
            }

            vpc {
              cidr_block_vpc = "10.0.0.0/16"
            }
          }
        }

        security {
          ssh_key = "jumper_ssh_key-sh-1585288-220404-010941" // Required
        }
      }

      distribution {
        os_arch    = "amd64"
        os_name    = "photon"
        os_version = "3"
        region     = "us-west-2"              // Required
        version    = "v1.21.2+vmware.1-tkg.2" // Required
      }

      topology {
        control_plane {
          availability_zones = [
            "us-west-2a",
          ]
          instance_type = "m5.large"
        }

        node_pools {
          spec {
            worker_node_count = "2"
            tkg_aws {
              availability_zone = "us-west-2a"
              instance_type     = "m5.large"
              node_placement {
                aws_availability_zone = "us-west-2a"
              }

              version = "v1.21.2+vmware.1-tkg.2"
            }
          }

          info {
            name = "md-0" // Required
          }
        }
      }
    }
  }
}