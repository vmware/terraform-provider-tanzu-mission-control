# Create Tanzu Mission Control nodepool entry
resource "tanzu-mission-control_cluster_node_pool" "create_node_pool" {

  management_cluster_name = "tkgs-terraform"
  provisioner_name = "test-gc-e2e-demo-ns"
  cluster_name = "tkgs-test"
  name = "terraform-nodepool"

  spec {
    worker_node_count = "3"
    cloud_labels = {
      "key1" : "val1"
    }
    node_labels = {
      "key2" : "val2"
    }

    tkg_service_vsphere  {
      class = "best-effort-xsmall"
      storage_class = "gc-storage-profile"
    }

  }

}
