---
Title: "Defining policies on different Scopes"
Description: |-
    Adding different policy resources to different scopes on which they are defined.
---

# Defining Security policy

Security policies allow you to manage the security context in which deployed pods operate in your clusters by imposing constraints that define what pods can do and which resources they can access.
The `tanzu-mission-control_security_policy` resource enables you to attach a security policy with an input recipe to a organisation, cluster group, or a cluster for management through Tanzu Mission Control.

## Security Policy on Cluster

For defining a security policy on a cluster, one can use dependency in the terraform script by defining a cluster resource (attach, workload clusters, EKS) and referencing the same in the policy resource.

For provisioning of a cluster, refer to the `tanzu-mission-control_cluster` in guides.

For cluster resource, one can again reference the cluster group name from the cluster group resource, based on the use case.

You could define a `security policy` for a cluster with the following config:

```terraform
/*
  NOTE: Creation of security policy depends on cluster attach
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
    kubeconfig_file = "<kube-config path>" // Required
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

# Cluster scoped Baseline Security Policy
resource "tanzu-mission-control_security_policy" "cluster_scoped_baseline_security_policy" {
  scope {
    cluster {
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name
    }
  }

  spec {
    input {
      baseline {
        audit              = false
        disable_native_psp = true
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values = [
          "api-server",
          "agent-gateway"
        ]
      }
      match_expressions {
        key      = "not-a-component"
        operator = "DoesNotExist"
        values   = []
      }
    }
  }
}
```


Similarly, one can define other policies such as custom, namespace quota and access policies using the above referencing hierarchy.
Also, the scope of the policy can directly be a cluster group, organisation, workspace, or a namespace.

Follow the below examples for reference.

## Access Policy on a Namespace

In the following example, there are multiple dependencies shown.

- Cluster dependency on cluster group
- Namespace dependency on cluster and workspace
- IAM policy dependency on namespace

```terraform
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
    kubeconfig_file = "<kube-config path>" // Required
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
```

## Custom Policy on a CLuster Group

```terraform
/*
  NOTE: Creation of custom policy depends on cluster group
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

# Cluster group scoped tmc-block-resources Custom Policy
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-block-resources_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster_group {
      cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name
    }
  }

  spec {
    input {
      tmc_block_resources {
        audit = false
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values = [
          "api-server",
          "agent-gateway"
        ]
      }
      match_expressions {
        key      = "not-a-component"
        operator = "DoesNotExist"
        values   = []
      }
    }
  }
}
```

