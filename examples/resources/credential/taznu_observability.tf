# Create credential for Tanzu Observability
resource "tanzu-mission-control_credential" "tanzu_observability_cred" {
  name = "tanzu_observability_cred"

  meta {
    description = "TMC integration: tanzu observability"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "wavefront.url":"url pointing to your wavefront instance"
    }
  }

  spec {
    capability = "TANZU_OBSERVABILITY"
    provider = "GENERIC_KEY_VALUE"
    data {
      key_value{
        data = {
          "wavefront.token" = "wavefront api token"
        }
      }
    }
  }
}
