// Tanzu Mission Control Repository Credential
// Operations supported : Read, Create, Update & Delete

# Read Tanzu Mission Control cluster group scoped repository credential : fetch cluster group scoped repository credential details
data "tanzu-mission-control_repository_credential" "test_repository_credential" {
  name = "<existing-name>" // Required

  org_id = "<existing-org-id>" // Default: taken from the authentication context

  scope {
    cluster {
      management_cluster_name = "<existing-management-cluster>" // Default: attached

      provisioner_name        = "<existing-prov-name>"          // Default: attached

      name                    = "<existing-cluster_name>"       // Required
    }
  }
}


resource "tanzu-mission-control_repository_credential" "test_repository_credential" {
  name = "test" 

  org_id = "test" // optional

  scope {
    cluster {
      management_cluster_name = "test"

      provisioner_name        = "test"

      name                    = "test"
    }
  }

  spec {
    username_password {
      username = "test"

      password = "test"
    }
  }
}
