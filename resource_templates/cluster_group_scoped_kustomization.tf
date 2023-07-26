// Create/ Delete/ Update Tanzu Mission Control cluster group scoped kustomization entry
resource "tanzu-mission-control_kustomization" "test_kustomization" {
  name = "test"

  namespace_name = "test"

  org_id = "test" // optional

  scope {
    cluster_group {
      name = "test"
    }
  }

  spec {
    source {
      namespace = "test" // Namespace of the repository. (can be referenced in Tf - to explicitly call delete on this resource if the dependency is deleted)

      name = "test" // Name of the repository. (can be referenced in Tf - to explicitly call delete on this resource if the dependency is deleted)
    }

    path = "/"

    prune = true

    interval = "5m" // can be "5s" etc.

    target_namespace = "test"
  }
}
