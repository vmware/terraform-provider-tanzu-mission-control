# Read Tanzu Mission Control helm charts : fetch helm charts details
data "tanzu-mission-control_helm_charts" "get_cluster_helm_repo" {
  name = "test_name"

  chart_metadata_name = "test_metadata_name"

  repository_name = "test_repository_name"
}