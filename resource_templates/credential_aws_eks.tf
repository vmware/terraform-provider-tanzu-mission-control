// Read Tanzu Mission Control AWS_EKS credential: fetch details
data "tanzu-mission-control_credential" "test_cred" {
  name = "<credential-name>"
}

// Create/ Delete Tanzu Mission Control AWS_EKS credential
resource "tanzu-mission-control_credential" "aws_eks_cred" {
  name = "<credential-name>"

  meta {
    description = "<description of the credential>"
    labels = {
     "key" : "<value>" ,
    }
  }

  spec {
    capability = "<capability-type>"
    provider = "<provider>"
    data {
      aws_credential {
        account_id = "<account-id>"
        iam_role{
          arn = "<IAM-role-ARN>"
          ext_id ="external-ID"
        }
      }
    }
  }
}
