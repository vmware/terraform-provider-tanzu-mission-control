# Read Tanzu Mission Control cluster nodepool : fetch cluster nodepool details
data "tanzu-mission-control_cluster_node_pool" "read_node_pool" {
  management_cluster_name = "tkgs-terraform"
  provisioner_name = "test-gc-e2e-demo-ns"
  cluster_name = "tkgs-test"
  name = "default-nodepool"
}
