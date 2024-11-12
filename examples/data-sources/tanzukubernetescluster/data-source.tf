data "tanzu-mission-control_tanzu_kubernetes_cluster" "read_tanzu_cluster" {
  name                    = "tanzu-cluster"
  management_cluster_name = "tanzu-mgmt-cluster"
  provisioner_name        = "tanzu-provisioner"
}
