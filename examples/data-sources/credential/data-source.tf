# Read Tanzu Mission Control credential : fetch credential details
data "tanzu_mission_control_credential" "test_cred" {
  name = "test-credential"
}
