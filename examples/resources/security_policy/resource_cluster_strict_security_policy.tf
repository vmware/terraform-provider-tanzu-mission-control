/*
Cluster scoped Tanzu Mission Control security policy with strict input recipe.
This policy is applied to a cluster with the strict configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_security_policy" "cluster_scoped_strict_security_policy" {
  name = "tf-sp-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      strict {
        audit              = true
        disable_native_psp = false
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "NotIn"
        values   = [
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
