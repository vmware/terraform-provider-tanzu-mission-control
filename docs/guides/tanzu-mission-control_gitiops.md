---
Title: "Gitops  Resource"
Description: |-
    Adding Gitops resource to a cluster.
---

# Defining Gitops

This feature will allow cluster admins to attach git repositories to their cluster or cluster groups ,
and sync folders in the repo to the cluster. Repositories may or may not require authentication.

The `tanzu-mission-control_git_repository` enables the Continuous Delivery feature on cluster or cluster group level scope.

# Terraform Resources for Gitops

## Git Repository

The `tanzu-mission-control_git_repository` resource allows you to add, update, and delete git repository to a particular scope through Tanzu Mission Control.

Git repositories are used to store kustomizations that will be synced to your cluster.

Please refer to [Add a Git Repository to a Cluster or Cluster Group.][git-repository]

[git-repository]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-26C2D2F3-0E5C-4E56-B875-B7FB003267E4.html

## Kustomization

The `tanzu-mission-control_kustomization` resource allows you to add, update, and delete Kustomization to a particular scope through Tanzu Mission Control.

In Creation of kustomization we must required to create Git Repository first, which we need to referenced in spec of kustomization, Git Repository can be created by using "tanzu-mission-control_git_repository" resource from terraform provider itself.

Please refer to [Add a Kustomization to a Cluster or Cluster Group.][kustomization]

[kustomization]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-99916A6D-5DAF-4A26-88C7-28662F847F2F.html

## Repository Credential

The `tanzu-mission-control_repository_credential` resource allows you to add, update, and delete repository credential to a particular scope through Tanzu Mission Control.

Repository credentials are used to authenticate to Git repositories and must be created before adding your Git repository.

Please refer to [Create a Repository Credential for a Cluster or Cluster Group.][repository-credential]

[repository-credential]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-657661A2-B26E-412A-9A46-7467A44A075A.html

To Create these resources you must be associated with the cluster.admin or clustergroup.admin role.


# Usecase Example for adding Cluster level Git Repository with basic authentication and kustomization

```terraform
/*
  NOTE: Creation of cluster level basic authentication git repository with kustomization
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "create_cluster_group" {
  name = "demo-cluster-group"
}

# Attach a Tanzu Mission Control cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
  name                    = "demo-cluster" // Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config-path>" // Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name // Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

# Create basic Repository Credential
resource "tanzu-mission-control_repository_credential" "create_cluster_source_secret_username_password" {
  name = "tf-secret" # Required

  scope {
    cluster {
      cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    username_password {
      username = "testusername" # Required
      password = "testpassword" # Required
    }
  }
}

# Create cluster level Git Repository with basic authentication
resource "tanzu-mission-control_git_repository" "create_cluster_git_repository" {
  name = "tf-git-repository-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster {
      cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    url                = "testGitRepositoryURL" # Required
    secret_ref         = tanzu-mission-control_repository_credential.create_cluster_source_secret_username_password.name
    interval           = "10m"    # Default: 5m
    git_implementation = "GO_GIT" # Default: GO_GIT
    ref {
      branch = "testBranchName"
      tag    = "testTag"
      semver = "testSemver"
      commit = "testCommit"
    }
  }
}

# Create cluster level Kustomization
resource "tanzu-mission-control_kustomization" "create_cluster_kustomization" {
  name = "tf-kustomization-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster {
      cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    path             = "/" # Required
    prune            = "test"
    interval         = "10m" # Default: 5m
    target_namespace = "testTargetNamespace"
    source {
      name      = tanzu-mission-control_git_repository.create_cluster_git_repository.name           # Required
      namespace = tanzu-mission-control_git_repository.create_cluster_git_repository.namespace_name # Required
    }
  }
}
```


# Usecase Example for adding Cluster level Git Repository and kustomization

```terraform
/*
  NOTE: Creation of cluster level git repository with kustomization
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "create_cluster_group" {
  name = "demo-cluster-group"
}

# Attach a Tanzu Mission Control cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
  name                    = "demo-cluster" // Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config-path>" // Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name // Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

# Create cluster level Git Repository
resource "tanzu-mission-control_git_repository" "create_cluster_git_repository" {
  name = "tf-git-repository-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster {
      cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    url                = "testGitRepositoryURL" # Required
    interval           = "10m"                  # Default: 5m
    git_implementation = "GO_GIT"               # Default: GO_GIT
    ref {
      branch = "testBranchName"
      tag    = "testTag"
      semver = "testSemver"
      commit = "testCommit"
    }
  }
}

# Create cluster level Kustomization
resource "tanzu-mission-control_kustomization" "create_cluster_kustomization" {
  name = "tf-kustomization-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster {
      cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    path             = "/" # Required
    prune            = "test"
    interval         = "10m" # Default: 5m
    target_namespace = "testTargetNamespace"
    source {
      name      = tanzu-mission-control_git_repository.create_cluster_git_repository.name           # Required
      namespace = tanzu-mission-control_git_repository.create_cluster_git_repository.namespace_name # Required
    }
  }
}
```


### Similarly, we can create SSH Key type Repository Credential for authenticated Git Repositories

# Usecase Example for adding Cluster level Repository Credential

```terraform
/*
  NOTE: Creation of cluster level SSH Key type Repository Credential
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "create_cluster_group" {
  name = "demo-cluster-group"
}

# Attach a Tanzu Mission Control cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
  name                    = "demo-cluster" // Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config-path>" // Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name // Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

# Create cluster level Repository Credential
resource "tanzu-mission-control_repository_credential" "create_cluster_source_secret_ssh" {
  name = "tf-secret" # Required

  scope {
    cluster {
      cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    ssh_key {
      identity    = "testidentity"    # Required
      known_hosts = "testknown_hosts" # Required
    }
  }
}
```


# Usecase Example for adding Cluster group level Git Repository and kustomization with SSH Key type authentication

```terraform
/*
  NOTE: Creation of cluster level SSH Key type Repository Credential
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "create_cluster_group" {
  name = "demo-cluster-group"
}

# Create cluster group level SSH key type Repository Credential
resource "tanzu-mission-control_repository_credential" "create_cluster_source_secret_ssh" {
  name = "tf-secret" # Required

  scope {
    cluster_group {
      cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    ssh_key {
      identity    = "testidentity"    # Required
      known_hosts = "testknown_hosts" # Required
    }
  }
}


# Create cluster group level Git Repository with SSH key type authentication
resource "tanzu-mission-control_git_repository" "create_cluster_group_git_repository" {
  name = "tf-git-repository-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster_group {
      cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    url                = "testGitRepositoryURL" # Required
    secret_ref         = tanzu-mission-control_repository_credential.create_cluster_source_secret_ssh.name
    interval           = "10m"    # Default: 5m
    git_implementation = "GO_GIT" # Default: GO_GIT
    ref {
      branch = "testBranchName"
      tag    = "testTag"
      semver = "testSemver"
      commit = "testCommit"
    }
  }
}


# Create cluster group level Kustomization
resource "tanzu-mission-control_kustomization" "create_cluster_group_kustomization" {
  name = "tf-kustomization-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster_group {
      cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    path             = "/" # Required
    prune            = "test"
    interval         = "10m" # Default: 5m
    target_namespace = "testTargetNamespace"
    source {
      name      = tanzu-mission-control_git_repository.create_cluster_group_git_repository.name           # Required
      namespace = tanzu-mission-control_git_repository.create_cluster_group_git_repository.namespace_name # Required
    }
  }
}
```


# Usecase Example for adding Cluster group level Git Repository and kustomization

```terraform
/*
  NOTE: Creation of cluster group level git repository with kustomization
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "create_cluster_group" {
  name = "demo-cluster-group"
}

# Create cluster group level Git Repository
resource "tanzu-mission-control_git_repository" "create_cluster_git_repository" {
  name = "tf-git-repository-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster_group {
      cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    url                = "testGitRepositoryURL" # Required
    interval           = "10m"                  # Default: 5m
    git_implementation = "GO_GIT"               # Default: GO_GIT
    ref {
      branch = "testBranchName"
      tag    = "testTag"
      semver = "testSemver"
      commit = "testCommit"
    }
  }
}

# Create cluster group level Kustomization
resource "tanzu-mission-control_kustomization" "create_cluster_kustomization" {
  name = "tf-kustomization-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster_group {
      cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    path             = "/" # Required
    prune            = "test"
    interval         = "10m" # Default: 5m
    target_namespace = "testTargetNamespace"
    source {
      name      = tanzu-mission-control_git_repository.create_cluster_git_repository.name           # Required
      namespace = tanzu-mission-control_git_repository.create_cluster_git_repository.namespace_name # Required
    }
  }
}
```