// Create/ Delete/ Update Tanzu Mission Control workspace scoped iam policy entry
resource "tanzu_mission_control_iam_policy" "workspace_scoped_iam_policy" {
  scope {
    workspace {
      name = "<workspace-name>" // Required
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
