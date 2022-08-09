/*
Organization scoped Tanzu Mission Control security policy with strict input recipe.
This policy is applied on a organization with the strict configuration option and is inherited by the cluster groups and clusters.
The scope and input blocks defined can be updated to change the policy's scope and it's recipe respectively.
*/
resource "tanzu-mission-control_security_policy" "organization_scoped_strict_security_policy" {
  name = "tf-sp-test"

  scope {
    organization {
      organization = "tf-create-test"
    }
  }

  spec {
    input {
      strict {
        audit              = false
        disable_native_psp = true
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
          "api-server",
          "agent-gateway"
        ]
      }
    }
  }
}
