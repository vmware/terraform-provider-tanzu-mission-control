# Create IMAGE_REGISTRY credential
resource "tanzu-mission-control_credential" "img_reg_cred" {
  name = "test-cred-name"

  meta {
    description = "credential"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "repository-namespace" : "something"
    }
  }

  spec {
    capability = "IMAGE_REGISTRY"
    provider = "GENERIC_KEY_VALUE"
    data {
      key_value{
        data  = {
          "registry-url" = "somethingnew"
          "ca-cert"  = "ca bundle"
        }
      }
    }
  }
}
