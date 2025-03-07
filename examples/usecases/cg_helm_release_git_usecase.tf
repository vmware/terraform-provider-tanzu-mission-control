/*
  NOTE: Creation of cluster group level helm release from git repository
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "cluster_group" {
  name = "demo-cluster-group"
}

# Create Tanzu Mission Control cluster group scope helm feature.
resource "tanzu-mission-control_helm_feature" "cg_helm_feature" {
  scope {
    cluster_group {
      name = tanzu-mission-control_cluster_group.cluster_group.name
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }
}

# Create Tanzu Mission Control cluster group scope helm release.
resource "tanzu-mission-control_helm_release" "cg_helm_release" {
  name = "test-helm-release-name" # Required

  namespace_name = "test-namespace-name" # Required

  scope {
    cluster_group {
      name = tanzu-mission-control_cluster_group.cluster_group.name
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.cg_helm_feature.scope[0].cluster_group[0].name

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