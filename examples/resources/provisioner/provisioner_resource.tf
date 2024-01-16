terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
      version = "1.4.1"
    }
  }
}
# Create provisioner resource
resource "tanzu-mission-control_provisioner" "create_provisioner" {
  name = "demo-test" # Required
  management_cluster = "eks" # Required

  meta {
    description = "Create provisioner through terraform-update"
    labels = {
      "key1" : "value1",
      "key2" : "value2",
    }
  }
}
# Read Tanzu Mission Control provisioner : fetch the given provisioner details
data "tanzu-mission-control_provisioner" "read_provisioner" {
  name = "demo-test-1" # Optional
  management_cluster = "eks" # Required
}

# Read Tanzu Mission Control provisioner : fetch all the provisioner details for the given management cluster
data "tanzu-mission-control_provisioner" "read_provisioner_1" {
  management_cluster = "eks" # Required
}

