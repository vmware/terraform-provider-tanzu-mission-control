// Create/ Delete/ Update Tanzu Mission Control organization scoped iam policy entry
resource "tanzu-mission-control_iam_policy" "organization_scoped_iam_policy" {
  scope {
    organization {
      org_id = "<organization-id>" // Required
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
