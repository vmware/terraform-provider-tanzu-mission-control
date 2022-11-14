// Read Tanzu Mission Control Image Registry credential: fetch details
data "tanzu-mission-control_credential" "test_cred" {
  name = "<credential-name>"
}

// Create/ Delete Tanzu Mission Control Image Registry credential
resource "tanzu-mission-control_credential" "img_reg_cred" {
  name = "<credential-name>"

  meta {
    description = "<description of the credential>"
    labels = {
     "key" : "<value>" ,
    }
    annotations = {
      "repository-path" : "<path>"
    }
  }

  spec {
    capability = "<capability-type>"
    provider = "<provider>"
    data {
      key_value{
        data  = {
          "registry-url" = "<url>"
          "ca-cert"  = "<ca-cert>"
        }
      }
    }
  }
}
