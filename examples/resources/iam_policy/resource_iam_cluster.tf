/*
 Cluster scoped Tanzu Mission Control IAM policy.
 This resource is applied on a cluster to provision the role bindings on the associated cluster.
 The defined scope block can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "cluster_scoped_iam_policy" {
  scope {
    cluster {
      management_cluster_name = "attached" # Default: attached
      provisioner_name        = "attached" # Default: attached
      name                    = "demo"
    }
  }

  role_bindings {
    role = "cluster.admin"
    subjects {
      name = "test"
      kind = "GROUP"
    }
  }
  role_bindings {
    role = "cluster.edit"
    subjects {
      name = "test-1"
      kind = "USER"
    }
    subjects {
      name = "test-2"
      kind = "USER"
    }
  }
}
