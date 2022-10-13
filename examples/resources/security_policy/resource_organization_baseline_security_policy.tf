/*
Organization scoped Tanzu Mission Control security policy with baseline input recipe.
This policy is applied to an organization with the baseline configuration option and is inherited by the cluster groups and clusters.
The defined scope and input blocks can be updated to change the policy's scope and recipe.
*/
resource "tanzu-mission-control_security_policy" "organization_scoped_baseline_security_policy" {
  name = "tf-sp-test"

  scope {
    organization {
      organization = "tf-create-test"
    }
  }

  spec {
    input {
      baseline {
        audit              = false
        disable_native_psp = true
      }
    }
  }
}
