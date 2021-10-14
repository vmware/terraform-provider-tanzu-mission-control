# Read TMC cluster : fetch cluster details
data "tmc_cluster" "ready_only_attach_cluster_view" {
  management_cluster_name = "attached"       # Default: attached
  provisioner_name        = "attached"       # Default: attached
  name                    = "terraform-test" # Required
}

# Read TMC TKG Service Vsphere workload cluster
data "tmc_cluster" "read_tkgs_cluster" {
  management_cluster_name = "test-tkgs"
  provisioner_name        = "test-gc-e2e-demo-ns"
  name                    = "cluster"
}
