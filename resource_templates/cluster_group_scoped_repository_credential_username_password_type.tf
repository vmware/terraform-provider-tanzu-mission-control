// Tanzu Mission Control Repository Credential
// Operations supported : Read, Create, Update & Delete

# Read Tanzu Mission Control cluster group scoped repository credential : fetch cluster group scoped repository credential details
data "tanzu-mission-control_repository_credential" "test_repository_credential" {
  name = "<existing-name>" // Required

  org_id = "<existing-org-id>" // Default: taken from the authentication context

  scope {
    cluster_group {
      name = "<existing-cluster_group_name>" // Required
    }
  }
}


resource "tanzu-mission-control_repository_credential" "test_repository_credential" {
  name = "test" 

  org_id = "test" // optional

  scope {
    cluster_group {
      name = "test"
    }
  }

  spec {
    username_password {
      username = "test"

      password = "test"
    }
  }
}
