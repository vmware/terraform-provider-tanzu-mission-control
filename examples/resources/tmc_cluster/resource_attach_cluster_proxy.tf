# Create Tanzu Mission Control attach cluster entry with proxy
resource "tmc_cluster" "attach_cluster_with_proxy" {
  management_cluster_name = "attached"               # Default: attached
  provisioner_name        = "attached"               # Default: attached
  name                    = "terraform-attach-proxy" # Required

  meta {
    description = "create attach cluster from terraform"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
    proxy         = "proxy-name"
  }

  wait_until_ready = false
}