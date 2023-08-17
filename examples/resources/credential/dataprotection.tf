# Create credential for TMC provisioned AWS S3 storage used for data-protection
resource "tanzu-mission-control_credential" "tmc_provisioned_aws_s3_cred" {
  name = "aws_s3_cred"

  meta {
    description = "TMC provisioned AWS S3 storage"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "DATA_PROTECTION"
    provider   = "AWS_EC2"
    data {
      aws_credential {
        iam_role {
          arn = "arn:aws:iam::4987398738934:role/clusterlifecycle-test.tmc.cloud.vmware.com"
        }
      }
    }
  }
}
