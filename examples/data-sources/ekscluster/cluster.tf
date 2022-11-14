# Read Tanzu Mission AWS EKS Control cluster : fetch cluster details
data "tanzu-mission-control_ekscluster" "tf_eks_cluster" {
  credential_name = "test-aws-cred-name" // Required
  region          = "us-west-2"          // Required
  name            = "test-cluster"       // Required
}
