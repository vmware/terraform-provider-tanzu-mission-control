# Read Tanzu Mission Control workspace : fetch workspace details
data "tmc_workspace" "read_workspace" {
  name = "default"
}
