// Tanzu Mission Control Cluster Type: Tanzu Kubernetes Grid AWS workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster : fetch cluster details for already present TKG AWS cluster
data "tanzu-mission-control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<existing-management-cluster>" // Required
  provisioner_name        = "<existing-prov-name>"          // Required
  name                    = "<existing-cluster-name>"       // Required
}

// Create Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_aws_cluster" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  name                    = "<cluster-name>"       // Required

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
    tkg_aws {
      advanced_configs {
        key   = "<key>"
        value = "<value>"
      }
      settings {
        network {
          cluster {
            pods {
              cidr_blocks = "<pods-cidr-blocks>" // Required
            }

            services {
              cidr_blocks = "<services-cidr-blocks>" // Required
            }
            api_server_port = api-server-port-default-value
          }
          provider {
            subnets {
              availability_zone = "<availability-zone>"
              cidr_block        = "<subnets-cidr-blocks>"
              is_public         = false
            }
            subnets {
              availability_zone = "<availability-zone>"
              cidr_block        = "<subnets-cidr-blocks>"
              is_public         = true
            }
            vpc {
              cidr_block = "<vpc-cidr-blocks>"
            }
          }
        }

        security {
          ssh_key = "<ssh-key>" // Required
        }
      }

      distribution {
        os_arch    = "<os-arch>"
        os_name    = "<os-name>"
        os_version = "<os-version>"
        region     = "<region>"  // Required
        version    = "<version>" // Required
      }

      topology {
        control_plane {
          availability_zones = [
            "<availability-zone>",
          ]
          instance_type = "<instance-type>"
        }

        node_pools {
          spec {
            worker_node_count = "<worker-node-count>"
            tkg_aws {
              availability_zone = "<availability-zone>"
              instance_type     = "<instance-type>"
              node_placement {
                aws_availability_zone = "<availability_zone>"
              }
              version = "<version>"
            }
          }

          info {
            name        = "<node-pool-name>" // Required
            description = "<node-pool-description>"
          }
        }
      }
    }
  }
}