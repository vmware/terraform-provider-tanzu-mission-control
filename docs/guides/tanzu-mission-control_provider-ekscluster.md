---
Title: "Onboarding of a an existing AWS EKS cluster"
Description: |-
    How to onboard an existing AWS EKS Cluster
---
# EKS Cluster

The `tanzu-mission-control_provider-ekscluster` resource can be used to bring an exsiting EKS cluster onto TMC for management.

## Prerequisites

To manage an EKS clusters, you need the following prerequisites.

- Onboard the AWS account onto TMC under which the EKS Clusters will be provisioned. Please refer [connecting an AWS account for EKS cluster lifecycle management][aws-account] guide for detailed steps. You can also use `tanzu-mission-control_credential` Terraform resource for this purpose. The name of the EKS AWS credential in TMC will be reffered to as `credential_name` in this guide.

- You must then provide the "clusterlifecycle" role, created while onboarding your AWS account, `cluster-admin` access to the EKS cluster. Please refer to [Enabling IAM principal access to your (EKS) cluster][add-user-role] document for more info.

- Ensure the CSP token used in initialising the terraform provider has the right set of permissions to create a workload cluster.

[aws-account]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-E4627693-7D1A-4914-A9DF-61E49F97FECC.html
[tanzu-vpc-guide]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-5708F04E-7EA3-495D-A484-FD6DB7AA8356.html
[aws-vpc-guide]: https://docs.aws.amazon.com/eks/latest/userguide/creating-a-vpc.html

## Managing the cluster with TMC

You can use the following template as reference to write your own `tanzu-mission-control_provider-ekscluster` resource in the terraform scripts. 

```terraform
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
```
