resource "tanzu-mission-control_mutation_policy" "org_label_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      label {
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
        scope = "Cluster"
        label {
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
