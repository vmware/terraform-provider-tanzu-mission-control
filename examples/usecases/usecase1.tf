/*
  NOTE: Creation of attach cluster depends on cluster-group creation
        Similarly, namespace creation depends on attach cluster and workspace creation
*/

terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}

// Create workspace
resource "tmc_workspace" "create_workspace" {
  name = "demo-workspace"
  meta {
    description = "Create workspace through terraform"
    labels = {
      "key1" : "value1",
    }
  }
}

// Create cluster group
resource "tmc_cluster_group" "create_cluster_group" {
  name = "demo-cluster-group"
}

// Create TMC attach cluster with k8s cluster kubeconfig provided
// The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tmc_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
  name                    = "demo-cluster" // Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config path>" // Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = tmc_cluster_group.create_cluster_group.name // Default: default
  }

  wait_until_ready = true // Default: false, when set resource waits until 3 min for the cluster to become ready

  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

// Create namespace with attached set as 'true' (need a running cluster)
resource "tmc_namespace" "create_namespace" {
  name                    = "demo-namespace"                                // Required
  cluster_name            = tmc_cluster.attach_cluster_with_kubeconfig.name // Required
  provisioner_name        = "attached"                                      // Default: attached
  management_cluster_name = "attached"                                      // Default: attached

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    workspace_name = tmc_workspace.create_workspace.name // Default: default
    attach         = true
  }
}

# Create TMC TKG Service Vsphere workload cluster entry
resource "tmc_cluster" "create_tkgs_workload" {
  management_cluster_name = "test-tkgs"
  provisioner_name        = "test-gc-e2e-demo-ns"
  name                    = "cluster"

  meta {
    labels      = { "key" : "test"}
  }

  spec {
    cluster_group = "default"
    tkg_service_vsphere {
      settings  {
        network  {
          pods  {
            cidr_blocks = [
              "172.20.0.0/16",
            ]
          }
          services  {
            cidr_blocks = [
              "10.96.0.0/16",
            ]
          }
        }
      }

      distribution  {
        version = "v1.20.8+vmware.1-tkg.2"
      }

      topology  {
        control_plane  {
          class = "best-effort-xsmall"
          storage_class = "wcpglobal-storage-profile"
          high_availability = false
        }
        node_pools  {
          spec  {
            worker_node_count = "1"
            tkg_service_vsphere  {
              class = "best-effort-xsmall"
              storage_class = "wcpglobal-storage-profile"
            }
          }
          info {
            name = "default-nodepool"
          }
        }
      }
    }
  }
}

output "cluster_group" {
  value = tmc_cluster_group.create_cluster_group
}

output "workspace" {
  value = tmc_workspace.create_workspace
}

output "namespace" {
  value = tmc_namespace.create_namespace
}

output "attach_output" {
  value = tmc_cluster.attach_cluster_with_kubeconfig
}

output "display_cluster" {
  value = tmc_cluster.create_tkgs_workload
}
