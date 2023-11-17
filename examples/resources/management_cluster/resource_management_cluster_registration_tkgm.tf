resource "tanzu-mission-control_management_cluster" "management_cluster_registration_tkgm" {
  name = "tf-registration-test" // Required

  spec {
    default_cluster_group           = "default" // Required
    kubernetes_provider_type        = "VMWARE_TANZU_KUBERNETES_GRID" // Required
    image_registry                  = "image_registry_value" // Optional - only allowed with TKGm - if supplied this should be the name of a pre configured local image registry configured in TMC to pull images from
    workload_cluster_image_registry = "workload_cluster_image_registry_value" // Optional - only allowed with TKGm - only allowed if image_registry is not empty
    proxy_name                      = "proxy_name_value" // Optional - name of proxy configuration to use which is already configured in TMC
    workload_cluster_proxy_name     = "workload_cluster_proxy_name_value" // Optional - only allowed if proxy_name is not empty
  }
}
