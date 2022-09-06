/*
 Cluster group scoped IAM policy.
 This resource is applied on a cluster group to provision the role bindings on the associated cluster group.
 The scope block defined can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "cluster_group_scoped_iam_policy" {
  scope {
    cluster_group {
      name = "default"
    }
  }

  role_bindings {
    role = "clustergroup.admin"
    subjects {
      name = "test"
      kind = "GROUP"
    }
  }
  role_bindings {
    role = "clustergroup.edit"
    subjects {
      name = "test-new"
      kind = "USER"
    }
  }
}
