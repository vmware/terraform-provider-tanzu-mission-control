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
resource "tanzu-mission-control_workspace" "workspace" {
  name = "demo-workspace"

  meta {
    description = "Create workspace through terraform"
    labels = {
      "key1" : "value1",
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

# Create namespace with attached set as 'true' (need a running cluster)
resource "tanzu-mission-control_namespace" "namespace" {
  name                    = "demo-namespace"
  cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name
  management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name
  provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    workspace_name = tanzu-mission-control_workspace.workspace.name // Default: default
    attach         = true
  }
}

# Namespace scoped Role Bindings
resource "tanzu-mission-control_iam_policy" "namespace_scoped_iam_policy" {
  scope {
    namespace {
      management_cluster_name = tanzu-mission-control_namespace.namespace.management_cluster_name
      provisioner_name        = tanzu-mission-control_namespace.namespace.provisioner_name
      cluster_name            = tanzu-mission-control_namespace.namespace.cluster_name
      name                    = tanzu-mission-control_namespace.namespace.name
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

## Custom Policy on a Cluster Group

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
resource "tanzu-mission-control_cluster_group" "cluster_group" {
  name = "demo-cluster-group"
}

# Cluster group scoped tmc-block-resources Custom Policy
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-block-resources_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster_group {
      cluster_group = tanzu-mission-control_cluster_group.cluster_group.name
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

## Custom Template and Custom Policy

Template provides a declarative definition of a policy, which can be used to apply custom constraints on managed kubernetes resources.
Custom policy consumes these declared custom templates to enforce specific policies. One must create the [custom template][custom-policy-template] before consuming it in the custom policy.
Please refer to custom policy template and custom policy terraform scripts within examples.

[custom-policy-template]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-F147492B-04FD-4CFD-8D1F-66E36D40D49C.html

## Refer the following example for creating custom policy template and assign it to custom policy

```terraform
/*
  NOTE: Creation of custom policy depends on cluster group and custom policy template.
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group.
resource "tanzu-mission-control_cluster_group" "cluster_group" {
  name = "tf-demo-cluster-group"
}

# Create custom policy template.
resource "tanzu-mission-control_custom_policy_template" "sample_template" {
  name = "tf-custom-template-test"

  spec {
    object_type   = "ConstraintTemplate"
    template_type = "OPAGatekeeper"

    data_inventory {
      kind    = "ConfigMap"
      group   = "admissionregistration.k8s.io"
      version = "v1"
    }

    data_inventory {
      kind    = "Deployment"
      group   = "extensions"
      version = "v1"
    }

    template_manifest = <<YAML
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
  name: tf-custom-template-test
  annotations:
    description: Requires Pods to have readiness and/or liveness probes.
spec:
  crd:
    spec:
      names:
        kind: tf-custom-template-test
      validation:
        openAPIV3Schema:
          properties:
            probes:
              type: array
              items:
                type: string
            probeTypes:
              type: array
              items:
                type: string
  targets:
    - target: admission.k8s.gatekeeper.sh
      rego: |
        package k8srequiredprobes
        probe_type_set = probe_types {
          probe_types := {type | type := input.parameters.probeTypes[_]}
        }
        violation[{"msg": msg}] {
          container := input.review.object.spec.containers[_]
          probe := input.parameters.probes[_]
          probe_is_missing(container, probe)
          msg := get_violation_message(container, input.review, probe)
        }
        probe_is_missing(ctr, probe) = true {
          not ctr[probe]
        }
        probe_is_missing(ctr, probe) = true {
          probe_field_empty(ctr, probe)
        }
        probe_field_empty(ctr, probe) = true {
          probe_fields := {field | ctr[probe][field]}
          diff_fields := probe_type_set - probe_fields
          count(diff_fields) == count(probe_type_set)
        }
        get_violation_message(container, review, probe) = msg {
          msg := sprintf("Container <%v> in your <%v> <%v> has no <%v>", [container.name, review.kind.kind, review.object.metadata.name, probe])
        }
YAML
  }
}


# Cluster group scoped custom template assigned Custom Policy
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_custom_template_assigned_custom_policy" {
  name = "tf-custom-template-policy-test"

  scope {
    cluster_group {
      cluster_group = tanzu-mission-control_cluster_group.cluster_group.name
    }
  }

  spec {
    input {
      custom {
        template_name = tanzu-mission-control_custom_policy_template.sample_template.name
        audit         = false

        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Deployment"
          ]
        }

        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "StatefulSet",
          ]
        }
      }
    }
  }
}
```
