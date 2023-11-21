resource "tanzu-mission-control_management_cluster" "management_cluster_registration_with_kubeconfig_file_path" {
  name = "tf-registration-test" // Required

  spec {
    cluster_group            = "default" // Required
    kubernetes_provider_type = "VMWARE_TANZU_KUBERNETES_GRID" // Required
  }

  register_management_cluster {
    tkgm_kubeconfig_file      = "<kube-config-path>" // Required
    tkgm_description          = "optional description about the kube-config provided" // Optional
  }

  ready_wait_timeout = "15m" // Optional , default value is 15m
}
