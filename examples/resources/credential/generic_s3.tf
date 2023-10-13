# Create Self provisioned AWS S3 or S3-compatible credential
resource "tanzu-mission-control_credential" "aws_eks_cred" {
  name = "tf-aws-s3-self-test"

  meta {
    description = "Self provisioned AWS S3 or S3-compatible storage credential for data protection"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "DATA_PROTECTION"
    provider   = "GENERIC_S3"
    data {
      key_value {
        type = "OPAQUE_SECRET_TYPE"
        data = {
          "aws_access_key_id"          = "abcd="
          "aws_secret_access_key"      = "xyz=="
        }
      }
    }
  }
  ready_wait_timeout = "1m" // Wait time for credential create operations to finish (default: 3m).
}
