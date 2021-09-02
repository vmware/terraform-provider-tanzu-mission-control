terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}

// Create cluster group
resource "tmc_cluster_group" "create_cluster_group" {
  name = "tf-cluster-group"
  meta {
    description = "Create cluster group through terraform"
    labels = {
      "key1" : "value1",
      "key2" : "value2"
    }
  }
}

// Create cluster group with minimal information
resource "tmc_cluster_group" "create_cluster_group_min_info" {
  name = "tf-cluster-group-min-info"
}

// Read TMC cluster group
data "tmc_cluster_group" "read_cluster_group" {
  name = "default"
}

// Output cluster group resource
output "cluster_group" {
  value = tmc_cluster_group.create_cluster_group
}

output "cluster_group_min_info" {
  value = tmc_cluster_group.create_cluster_group_min_info
}

output "display_cluster_group" {
  value = data.tmc_cluster_group.read_cluster_group
}
