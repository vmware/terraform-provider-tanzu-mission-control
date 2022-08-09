/*
Cluster scoped Tanzu Mission Control security policy with baseline input recipe.
This policy is applied on a cluster with the baseline configuration option.
The scope and input blocks defined can be updated to change the policy's scope and it's recipe respectively.
*/
resource "tanzu-mission-control_security_policy" "cluster_scoped_baseline_security_policy" {
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
      baseline {
        audit              = true
        disable_native_psp = false
      }
    }

    namespace_selector {
      match_expressions {
        key      = "not-a-component"
        operator = "DoesNotExist"
        values   = []
      }
    }
  }
}
