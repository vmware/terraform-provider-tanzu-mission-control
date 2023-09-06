# Create AWS_EKS credential
resource "tanzu-mission-control_credential" "aws_eks_cred" {
  name = "test-cred-name"

  meta {
    description = "credential"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "MANAGED_K8S_PROVIDER"
    provider   = "AWS_EKS"
    data {
      aws_credential {
        account_id = "account-id"
        iam_role {
          arn    = "arn:aws:iam::4987398738934:role/clusterlifecycle-test.tmc.cloud.vmware.com"
          ext_id = ""
        }
      }
    }
  }
  ready_wait_timeout = "2m" // Wait time for credential create operations to finish (default: 3m).
}
