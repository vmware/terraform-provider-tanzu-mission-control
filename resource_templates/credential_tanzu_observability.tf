// Read credential for Tanzu Observability: fetch details
data "tanzu-mission-control_credential" "test_cred" {
  name = "<credential-name>"
}

// Create/ Delete credential for Tanzu Observability
resource "tanzu-mission-control_credential" "tanzu_observability_cred" {
  name = "<credential-name>"

  meta {
    description = "<description of the credential>"
    labels = {
     "key" : "<value>" ,
    }
    annotations = {
      "wavefront.url" : "<url of wavefront instance>"
    }
  }

  spec {
    capability = "<capability-type>"
    provider = "<provider>"
    data {
      key_value{
        data  = {
          "wavefront.token" = "<wavefront api token>"
        }
      }
    }
  }
}
