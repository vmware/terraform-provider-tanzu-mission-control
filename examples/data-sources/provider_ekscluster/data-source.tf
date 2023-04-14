# Read Tanzu Mission AWS Provider EKS cluster : fetch cluster details
data "tanzu-mission-control_provider-ekscluster" "tf_eks_cluster" {
  credential_name = "test-aws-cred-name" // Required
  region          = "us-west-2"          // Required
  name            = "adoption-test"      // Required
}
