# Create Tanzu Mission Control source secret with attached set as default value.
resource "tanzu-mission-control_repository_credential" "create_cluster_group_source_secret_username_password" {
  name = "tf-secret" # Required

  scope {
    cluster_group {
      name = "default" # Required
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    data {
        username_password {
        username = "testusername" # Required
        password = "testpassword" # Required
      }
    }
  }
}
