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
  name = "tf-kustomizationy-name" # Required

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
