// Create/ Delete/ Update Tanzu Mission Control cluster scoped iam policy entry
resource "tanzu-mission-control_iam_policy" "cluster_scoped_iam_policy" {
  scope {
    cluster {
      management_cluster_name = "<management-cluster>" // Default: attached
      provisioner_name        = "<prov-name>"          // Default: attached
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
