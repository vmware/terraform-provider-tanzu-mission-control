# Read Tanzu Mission Control cluster : fetch cluster details
data "tanzu-mission-control_cluster" "ready_only_attach_cluster_view" {
  management_cluster_name = "attached"       # Default: attached
  provisioner_name        = "attached"       # Default: attached
  name                    = "terraform-test" # Required
}

# Read Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster : fetch TKGs cluster details
data "tanzu-mission-control_cluster" "read_tkgs_cluster" {
  management_cluster_name = "test-tkgs"
  provisioner_name        = "test-gc-e2e-demo-ns"
  name                    = "tkgs-workload"
}

# Read Tanzu Mission Control Tanzu Kubernetes Grid vSphere workload cluster : fetch TKG vSphere cluster details
data "tanzu-mission-control_cluster" "read_tkg_vsphere_cluster" {
  management_cluster_name = "tkgm-vsphere"
  provisioner_name        = "default"
  name                    = "tkgm-workload"
}

# Read Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster : fetch TKG AWS cluster details
data "tanzu-mission-control_cluster" "read_tkg_aws_cluster" {
  management_cluster_name = "tkgm-aws"
  provisioner_name        = "default"
  name                    = "tkgm-aws-workload"
}
