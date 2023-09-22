# Create Tanzu Mission Control cluster scope helm release with attached set as default value.
resource "tanzu-mission-control_helm_release" "create_cl_helm_release_gitrepo_type" {
  name = "test-helm-release-name" # Required

  namespace = "test-namespace-name" # Required

  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.create_cg_helm_feature.scope[0].cluster[0].name

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