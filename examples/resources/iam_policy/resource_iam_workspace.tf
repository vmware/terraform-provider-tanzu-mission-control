/*
 Workspace scoped Tanzu Mission Control IAM policy.
 This resource is applied on a workspace to provision the role bindings on the associated workspace.
 The defined scope block can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "workspace_scoped_iam_policy" {
  scope {
    workspace {
      name = "tf-workspace"
    }
  }

  role_bindings {
    role = "workspace.edit"
    subjects {
      name = "test"
      kind = "USER"
    }
  }
}
