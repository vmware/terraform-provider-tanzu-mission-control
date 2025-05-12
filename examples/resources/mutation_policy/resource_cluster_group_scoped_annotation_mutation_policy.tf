resource "tanzu-mission-control_mutation_policy" "cluster_group_annotation_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      annotation {
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
        scope = "*"
        annotation {
          key   = "test"
          value = "optional"
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values   = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
