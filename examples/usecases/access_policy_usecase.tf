/*
  NOTE: Creation of IAM policy depends on namespace
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create workspace
resource "tanzu-mission-control_workspace" "create_workspace" {
  name = "demo-workspace"

  meta {
    description = "Create workspace through terraform"
    labels = {
      "key1" : "value1",
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

# Create namespace with attached set as 'true' (need a running cluster)
resource "tanzu-mission-control_namespace" "create_namespace" {
  name                    = "demo-namespace"
  cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name
  management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name
  provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    workspace_name = tanzu-mission-control_workspace.create_workspace.name // Default: default
    attach         = true
  }
}

# Namespace scoped Role Bindings
resource "tanzu-mission-control_iam_policy" "namespace_scoped_iam_policy" {
  scope {
    namespace {
      management_cluster_name = tanzu-mission-control_namespace.create_namespace.management_cluster_name
      provisioner_name        = tanzu-mission-control_namespace.create_namespace.provisioner_name
      cluster_name            = tanzu-mission-control_namespace.create_namespace.cluster_name
      name                    = tanzu-mission-control_namespace.create_namespace.name
    }
  }

  role_bindings {
    role = "namespace.view"
    subjects {
      name = "test-1"
      kind = "USER"
    }
    subjects {
      name = "test-2"
      kind = "GROUP"
    }
  }
}

