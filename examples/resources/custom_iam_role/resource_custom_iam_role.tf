resource "tanzu-mission-control_custom_iam_role" "demo-role" {
  name = "tf-custom-role"

  spec {
    is_deprecated = false

    aggregation_rule {
      cluster_role_selector {
        match_labels = {
          key = "value"
        }
      }

      cluster_role_selector {
        match_expression {
          key      = "aa"
          operator = "Exists"
          values   = ["aa", "bb", "cc"]
        }
      }
    }

    allowed_scopes = [
      "ORGANIZATION",
      "CLUSTER_GROUP",
      "CLUSTER"
    ]

    tanzu_permissions = []

    kubernetes_permissions {
      rule {
        resources  = ["deployments"]
        verbs      = ["get", "list"]
        api_groups = ["*"]
      }

      rule {
        verbs      = ["get", "list"]
        api_groups = ["*"]
        url_paths  = ["/healthz"]
      }
    }
  }
}
