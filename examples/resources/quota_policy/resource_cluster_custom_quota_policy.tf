/*
Cluster scoped Tanzu Mission Control namespace quota policy with custom input recipe.
This policy is applied to a cluster with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_scoped_custom_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
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
  }
}
