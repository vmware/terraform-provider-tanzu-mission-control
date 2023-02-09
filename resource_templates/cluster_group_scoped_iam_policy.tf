// Create/ Delete/ Update Tanzu Mission Control cluster group scoped iam policy entry
resource "tanzu_mission_control_iam_policy" "cluster_group_scoped_iam_policy" {
  scope {
    cluster_group {
      name = "<cluster-group-name>" // Required
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
