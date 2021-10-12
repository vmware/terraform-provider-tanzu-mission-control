# Read TMC cluster : fetch cluster details
data "tmc_cluster" "ready_only_attach_cluster_view" {
  management_cluster_name = "attached"       # Default: attached
  provisioner_name        = "attached"       # Default: attached
  name                    = "terraform-test" # Required
}
