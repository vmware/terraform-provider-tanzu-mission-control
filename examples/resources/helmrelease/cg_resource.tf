# Create Tanzu Mission Control cluster group scope helm release with attached set as default value.
resource "tanzu-mission-control_helm_release" "create_cg_helm_release" {
  name = "test-helm-release-name" # Required

  namespace = "test-namespace-name" # Required

  scope {
    cluster_group {
      name = "default" # Required
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.create_cg_helm_feature.scope[0].cluster_group[0].name

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    chart_ref {
      git_repository {
        repository_name      = "testgitrepo"
        repository_namespace = "test-gitrepo-namespace"
        chart_path           = "chartpath"
      }
    }

    inline_config = "<inline-config-file-path>"

    target_namespace = "testtargetnamespacename"

    interval = "10m" # Default: 5m
  }
}