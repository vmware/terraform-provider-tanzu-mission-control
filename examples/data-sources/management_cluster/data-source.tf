# Read Tanzu Mission Control management cluster registration : fetch management cluster registration details
data "tanzu-mission-control_management_cluster" "read_management_cluster_registration" {
  name   = "default" # Required
  org_id = "<ID of Organization>" # Optional value
}
