// Tanzu Mission Control Provider EKS Cluster Type: AWS EKS clusters.
// Operations supported : Read, Create, Update & Delete

// Read Tanzu Mission Control Provider AWS EKS cluster : fetch cluster details
data "tanzu-mission-control_provider-ekscluster" "tf_eks_cluster" {
  credential_name = "<aws-credential-name>" // Required
  region          = "<aws-region>"          // Required
  name            = "<cluster-name>"        // Required
}

// Create Tanzu Mission Control Provider AWS EKS cluster entry: onboard an EKS cluster
resource "tanzu-mission-control_ekscluster" "tf_eks_cluster" {
  credential_name = "<aws-credential-name>" // Required
  region          = "<aws-region>"          // Required
  name            = "<eks-cluster-name>"    // Required

  ready_wait_timeout = "<time>" // Wait time for cluster operations to finish (default: 30m).

  meta {
    description = "description of the cluster"
    labels      = { "<key>" : "<value>" }
  }

  spec {
    cluster_group = "test-cluster-group" // Default: default
    proxy         = "<proxy>"

    eks_arn = "<eks-cluster-arn>"        // Required, forces new
    agent_name = "<tmc-cluster-name>"    // Required, forces new
  }
}
