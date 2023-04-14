---
Title: "Provider EKS Cluster Resource"
Description: |-
    Onboard an Existing EKS cluster resource.
---

# Provider EKS Cluster 

The `tanzu-mission-control_provider-ekscluster` resource enables you to attach an conformant [AWS EKS][aws-eks] clusters for management through Tanzu Mission Control.
With Tanzu Kubernetes clusters, you can also provision resources to create new workload clusters.

## Onboarding an existing EKS Cluster

To use the **Tanzu Mission Control** for managing an existing EKS cluster, you must first connect your AWS account to TMC.
For more information, please refer [connecting an AWS account for EKS cluster lifecycle management][aws-account].

You must then provide the "clusterlifecycle" role, created while onboarding your AWS account, `cluster-adming` access to the EKS cluster. Please refer to [Enabling IAM principal access to your (EKS) cluster][add-user-role] document for more info.

You must also have the appropriate permissions in TMC:

- To provision a cluster, you must have `cluster.admin` permissions.
- - You must also have `clustergroup.edit` permissions on the cluster group in
    which you want to put the new cluster.

[aws-eks]: https://aws.amazon.com/eks/
[aws-account]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-E4627693-7D1A-4914-A9DF-61E49F97FECC.html
[add-user-role]: https://docs.aws.amazon.com/eks/latest/userguide/add-user-role.html

## Example Usage

```terraform
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
```
