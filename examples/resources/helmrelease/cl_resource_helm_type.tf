# Create Tanzu Mission Control cluster scope helm release with attached set as default value.
resource "tanzu-mission-control_helm_release" "cl_helm_release_helm_type" {
  name = "test-helm-release-name" # Required

  namespace_name = "test-namespace-name" # Required

  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.cg_helm_feature.scope[0].cluster[0].name

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    chart_ref {
      helm_repository {
        repository_name      = "testgitrepo"
        repository_namespace = "test-helm-namespace"
        chart_name           = "chart-name"
        version              = "test-version"
      }
    }

    inline_config = "<inline-config-file-path>"

    target_namespace = "testtargetnamespacename"

    interval = "10m" # Default: 5m
  }
}