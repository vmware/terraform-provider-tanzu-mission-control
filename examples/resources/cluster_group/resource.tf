# Create Tanzu Mission Control cluster group
resource "tanzu-mission-control_cluster_group" "cluster_group" {
  name = "tf-cluster-group"
  meta {
    description = "Create cluster group through terraform"
    labels = {
      "key1" : "value1",
      "key2" : "value2"
    }
  }
}

# Create cluster group with minimal information
resource "tanzu-mission-control_cluster_group" "cluster_group_min_info" {
  name = "tf-cluster-group-min-info"
}
