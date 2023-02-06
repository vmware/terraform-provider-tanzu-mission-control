/*
Cluster group scoped Tanzu Mission Control namespace quota policy with custom input recipe.
This policy is applied to a cluster group with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_group_scoped_custom_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      custom {
        limits_cpu               = "4"
        limits_memory            = "8Mi"
        persistent_volume_claims = 2
        persistent_volume_claims_per_class = {
          ab : 2
          cd : 4
        }
        requests_cpu     = "2"
        requests_memory  = "4Mi"
        requests_storage = "2G"
        requests_storage_per_class = {
          test : "2G"
          twt : "4G"
        }
        resource_counts = {
          pods : 2
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
