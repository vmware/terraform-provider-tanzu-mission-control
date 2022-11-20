// Read Tcredential for TMC provisioned AWS S3 storage used for data-protection: fetch details
data "tanzu-mission-control_credential" "test_cred" {
  name = "<credential-name>"
}

// Create/ Delete credential for TMC provisioned AWS S3 storage used for data-protection
resource "tanzu-mission-control_credential" "tmc_provisioned_aws_s3_cred" {
  name = "<credential-name>"

  meta {
    description = "<description of the credential>"
    labels = {
      "key" : "<value>",
    }
  }

  spec {
    capability = "<capability-type>"
    provider   = "<provider>"
    data {
      aws_credential {
        iam_role {
          arn = "<IAM-role-ARN>"
        }
      }
    }
  }
}
