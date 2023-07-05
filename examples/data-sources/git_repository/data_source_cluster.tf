# Read Tanzu Mission Control git repository : fetch cluster git repository details
resource "tanzu-mission-control_git_repository" "read_cluster_git_repository" {
  name = "tf-git-repository-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster {
      cluster_name            = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }
}
