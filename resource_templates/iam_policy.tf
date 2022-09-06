// Create/ Delete/ Update Tanzu Mission Control iam policy entry
resource "tanzu-mission-control_iam_policy" "create_iam_policy" {
  scope {
    cluster {
      management_cluster_name = "<management-cluster>" // Required
      provisioner_name        = "<prov-name>"          // Required
      name                    = "<cluster-name>"       // Required
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
