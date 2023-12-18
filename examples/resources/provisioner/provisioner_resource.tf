# Create provisioner resource
resource "tanzu-mission-control_provisioner" "create_provisioner" {
  name = "demo-test" # Required
  management_cluster = "eks" # Required

  meta {
    description = "Create provisioner through terraform"
    labels = {
      "key1" : "value1",
      "key2" : "value2",
    }
  }
}
