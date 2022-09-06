/*
 Organization scoped IAM policy.
 This resource is applied on an organization to provision the role bindings on the associated organization.
 The scope block defined can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "organization_scoped_iam_policy" {
  scope {
    organization{
      org_id = "dummy-org-id"
    }
  }

  role_bindings {
    role = "organization.view"
    subjects {
      name = "test-1"
      kind = "USER"
    }
    subjects {
      name = "test-2"
      kind = "GROUP"
    }
  }
  role_bindings {
    role = "organization.edit"
    subjects {
      name = "test-3"
      kind = "USER"
    }
  }
}
