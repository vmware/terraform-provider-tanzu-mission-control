/*
Cluster group scoped Tanzu Mission Control security policy with strict input recipe.
This policy is applied on a cluster group with the strict configuration option and is inherited by the clusters.
The scope and input blocks defined can be updated to change the policy's scope and it's recipe respectively.
*/
resource "tanzu-mission-control_security_policy" "cluster_group_scoped_strict_security_policy" {
  name = "tf-sp-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      strict {
        audit              = false
        disable_native_psp = true
      }
    }
  }
}
