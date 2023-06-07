/*
Organization scoped Tanzu Mission Control network policy with allow-all-to-pods input recipe.
This policy is applied to a organization with the allow-all-to-pods configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_allow-all-to-pods_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      allow_all_to_pods {
        from_own_namespace = false
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values = [
          "api-server",
          "agent-gateway"
        ]
      }
      match_expressions {
        key      = "not-a-component"
        operator = "DoesNotExist"
        values   = []
      }
    }
  }
}
