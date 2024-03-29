# Read Tanzu Mission Control provisioner : fetch the given provisioner details
data "tanzu-mission-control_provisioner" "read_provisioner" {
  provisioners {
    name = "test-provisioner" # Optional
    management_cluster = "eks" # Required
  }
}

# Read Tanzu Mission Control provisioner : fetch all the provisioner details for the given management cluster
data "tanzu-mission-control_provisioner" "read_provisioner" {
  provisioners {
    management_cluster = "eks" # Required
  }
}
