resource "tanzu-mission-control_mutation_policy" "cluster_label_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "test"
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
        scope = "*"
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
