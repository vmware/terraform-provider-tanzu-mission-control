# Read Tanzu Mission Control namespace : fetch namespace details
data "tanzu-mission-control_namespace" "read_namespace" {
  name                    = "tf-namespace" # Required
  cluster_name            = "testcluster"  # Required
  management_cluster_name = "attached"     # Default: attached
  provisioner_name        = "attached"     # Default: attached
}