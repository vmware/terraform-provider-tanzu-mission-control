/*
 Namespace scoped Tanzu Mission Control IAM policy.
 This resource is applied on a namespace to provision the role bindings on the associated namespace.
 The defined scope block can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "namespace_scoped_iam_policy" {
  scope {
    namespace {
      management_cluster_name = "attached" # Default: attached
      provisioner_name        = "attached" # Default: attached
      cluster_name            = "demo"
      name                    = "tf_namespace"
    }
  }

  role_bindings {
    role = "namespace.view"
    subjects {
      name = "test-1"
      kind = "USER"
    }
    subjects {
      name = "test-2"
      kind = "GROUP"
    }
  }
}
