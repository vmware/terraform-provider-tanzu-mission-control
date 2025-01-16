# Create Tanzu Mission Control nodepool entry
resource "tanzu-mission-control_cluster_node_pool" "node_pool" {

  management_cluster_name = "tkgs-terraform"
  provisioner_name        = "test-gc-e2e-demo-ns"
  cluster_name            = "tkgs-test"
  name                    = "terraform-nodepool"
  ready_wait_timeout      = "10m"

  spec {
    worker_node_count = "3"
    cloud_labels = {
      "key1" : "val1"
    }
    node_labels = {
      "key2" : "val2"
    }

    tkg_service_vsphere {
      class         = "best-effort-xsmall"
      storage_class = "tkgs-k8s-obj-policy"
      volumes {
        capacity          = 4
        mount_path        = "/var/lib/etcd"
        name              = "etcd-0"
        pvc_storage_class = "tkgs-k8s-obj-policy"
      }
    }
  }
}
