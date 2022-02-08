# Create Tanzu Mission Control attach cluster entry
resource "tmc_cluster" "attach_cluster_without_apply" {
  management_cluster_name = "attached"         # Default: attached
  provisioner_name        = "attached"         # Default: attached
  name                    = "terraform-attach" # Required

  meta {
    description = "create attach cluster from terraform"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
  }

  # The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}