# Read Tanzu Mission Control git repository : fetch cluster group git repository details
resource "tanzu-mission-control_git_repository" "read_cluster_group_git_repository" {
  name = "tf-git-repository-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster_group {
      name = "default" # Required
    }
  }
}
