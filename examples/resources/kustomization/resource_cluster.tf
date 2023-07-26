# Create Tanzu Mission Control kustomization with attached set as default value.
resource "tanzu-mission-control_kustomization" "create_cluster_kustomization" {
  name = "tf-kustomization-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster {
      cluster_name            = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    path = "testPath" # Required
    prune = "testPrune"
    interval = "10m" # Default: 5m
    target_namespace = "testTargetNamespace"
    source {
        name = "testGitRepositoryName" # Required
        namespace = "testGitRepositoryNamespace" # Required
    }
  }
}
