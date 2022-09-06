// Create/ Delete/ Update Tanzu Mission Control namespace scoped iam policy entry
resource "tanzu-mission-control_iam_policy" "namespace_scoped_iam_policy" {
  scope {
    namespace {
      management_cluster_name = "<management-cluster>" // Default: attached
      provisioner_name        = "<prov-name>"          // Default: attached
      cluster_name            = "<cluster-name>"       // Required
      name                    = "<namespace-name>"     // Required
    }
  }

  role_bindings {
    role = "<user-role>"
    subjects {
      name = "<user-identity>"
      kind = "<identities>"
    }
  }
  role_bindings {
    role = "<user-role>"
    subjects {
      name = "<user-identity>"
      kind = "<identities>"
    }
    subjects {
      name = "<user-identity>"
      kind = "<identities>"
    }
  }
}
