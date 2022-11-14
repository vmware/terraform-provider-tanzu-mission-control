# Read Tanzu Mission Control credential : fetch credential details
data "tanzu-mission-control_credential" "test_cred" {
  name = "test-credential"
}