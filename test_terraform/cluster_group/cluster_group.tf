terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}

// Create cluster group
resource "tmc_cluster_group" "cluster_group_create" {
    name               = "tf-cluster-group"
    meta  {
        description    = "Create cluster group through terraform"
        labels         = {
            "key1" : "value1",
            "key2" : "value2"
        }
    }
}

// Create cluster group with minimal information
resource "tmc_cluster_group" "cluster_group_create_min_info" {
    name = "tf-cluster-group-min-info"
}

// Output cluster group resource
output "cluster_group" {
    value = tmc_cluster_group.cluster_group_create
}

output "cluster_group_min_info" {
    value = tmc_cluster_group.cluster_group_create_min_info
}
