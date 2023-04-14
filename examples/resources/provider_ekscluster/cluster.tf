# Adopt an existing AWS EKS cluster onto Tanzu Mission Control
resource "tanzu-mission-control_provider-ekscluster" "existing_eks_cluster" {
  credential_name = "eks-test"          // Required
  region          = "us-west-2"         // Required
  name            = "adopted-cluster"   // Required

  ready_wait_timeout = "30m" // Wait time for cluster operations to finish (default: 30m).

  meta {
    description = "eks test cluster"
    labels      = { "key1" : "value1" }
  }

  spec {
    cluster_group = "test-cluster-group" // Default: default
    proxy    		  = "<proxy>"              // Proxy if used

    eks_arn = "arn:aws:eks:us-west-2:999999999999:cluster/adopted-cluster" // Required, forces new
    agent_name = "tmc-cluster-name" // Required, forces new
  }
}
