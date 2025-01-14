# Create Tanzu Mission Control workspace
resource "tanzu-mission-control_workspace" "workspace" {
  name = "tf-workspace-test"

  meta {
    description = "Create workspace through terraform"
    labels = {
      "key1" : "value1",
      "key2" : "value2"
    }
  }
}
