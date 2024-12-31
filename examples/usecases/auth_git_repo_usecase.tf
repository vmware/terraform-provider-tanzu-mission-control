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
resource "tanzu-mission-control_cluster_group" "cluster_group" {
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
    cluster_group = tanzu-mission-control_cluster_group.cluster_group.name // Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

# Create basic Repository Credential
resource "tanzu-mission-control_repository_credential" "cluster_source_secret_username_password" {
  name = "tf-secret" # Required

  scope {
    cluster {
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    data {
      username_password {
        username = "testusername" # Required
        password = "testpassword" # Required
      }
    }
  }
}

# Create cluster level Git Repository with basic authentication
resource "tanzu-mission-control_git_repository" "cluster_git_repository" {
  name = "tf-git-repository-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster {
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
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
    secret_ref         = tanzu-mission-control_repository_credential.cluster_source_secret_username_password.name
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
resource "tanzu-mission-control_kustomization" "cluster_kustomization" {
  name = "tf-kustomization-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster {
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
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
      name      = tanzu-mission-control_git_repository.cluster_git_repository.name           # Required
      namespace = tanzu-mission-control_git_repository.cluster_git_repository.namespace_name # Required
    }
  }
}