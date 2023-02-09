# Create Tanzu Mission Control workspace
resource "tanzu_mission_control_workspace" "create_workspace" {
  name = "tf-workspace-test"

  meta {
    description = "Create workspace through terraform"
    labels = {
      "key1" : "value1",
      "key2" : "value2"
    }
  }
}
