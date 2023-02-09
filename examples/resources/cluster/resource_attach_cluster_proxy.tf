# Create Tanzu Mission Control attach cluster entry with proxy
resource "tanzu_mission_control_cluster" "attach_cluster_with_proxy" {
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

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
}
