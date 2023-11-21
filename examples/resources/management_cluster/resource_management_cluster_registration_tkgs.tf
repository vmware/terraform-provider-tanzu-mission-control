resource "tanzu-mission-control_management_cluster" "management_cluster_registration_tkgs" {
  name = "tf-registration-test" // Required

  spec {
    default_cluster_group       = "default" // Required
    kubernetes_provider_type    = "VMWARE_TANZU_KUBERNETES_GRID_SERVICE" // Required
    proxy_name                  = "proxy_name_value" // Optional - name of proxy configuration to use which is already configured in TMC
    workload_cluster_proxy_name = "workload_cluster_proxy_name_value" // Optional - only allowed if proxy_name is not empty
  }
}
