resource "tanzu-mission-control_management_cluster" "management_cluster_registration_minimal_tkgm" {
  name = "tf-registration-test" // Required

  spec {
    default_cluster_group    = "default" // Required
    kubernetes_provider_type = "VMWARE_TANZU_KUBERNETES_GRID" // Required
  }
}
