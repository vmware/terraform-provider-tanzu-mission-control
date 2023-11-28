---
Title: "Mutation Policy Resource"
Description: |-
    Creating the Tanzu Kubernetes mutation policy resource.
---

# Mutation Policy

The `tanzu-mission-control_mutation_policy` resource enables you to attach a mutation policy with an input recipe to a organisation, cluster group, or a cluster for management through Tanzu Mission Control.

## Input Recipe

In the Tanzu Mission Control mutation policy resource, there are three system defined types of mutation templates that you can use:
- **annotation**
- **label**
- **pod-security**

## Policy Scope and Inheritance

In the Tanzu Mission Control resource hierarchy, there are three levels at which you can specify mutation policy resources:
- **organization** - `organization` block under `scope` sub-resource
- **object groups** - `cluster_group` block under `scope` sub-resource
- **Kubernetes objects** - `cluster` block under `scope` sub-resource

**Note:**
The scope parameter is mandatory in the schema and the user needs to add one of the defined scopes to the script for the provider to function.
Only one scope per resource is allowed.

## Target Kubernetes Resources

Label and annotation mutation policy recipes contain a Kubernetes Resource spec that contains `api_groups` and `kind` as sub fields.
These attributes are of the kind `[]string` which the policy API supports. In terraform, while declaring multiple
`api_groups` and `kinds` under one block of `target_kubernetes_resources` is validated by the API but not reflected on the UI.
For UI comparison with Terraform, one must add multiple blocks of `target_kubernetes_resources`, each containing a API Group and a Kind.

Example:

```
target_kubernetes_resources {
  api_groups = [
    "apps",
  ]
  kinds = [
    "Event",
  ]
}
target_kubernetes_resources {
  api_groups = [
    "batch",
  ]
  kinds = [
    "Pod",
  ]
}
```

## Cluster scoped annotation Mutation Policy

### Example Usage

```terraform
resource "tanzu-mission-control_mutation_policy" "cluster_annotation_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "test"
    }
  }
  spec {
    input {
      annotation {
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
        scope = "*"
        annotation {
          key   = "test"
          value = "optional"
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
```

## Cluster scoped label Mutation Policy

### Example Usage

```terraform
resource "tanzu-mission-control_mutation_policy" "cluster_label_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "test"
    }
  }

  spec {
    input {
      label {
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
        scope = "*"
        label {
          key   = "test"
          value = "optional"
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
```

## Cluster scoped pod-security Mutation Policy

### Example Usage

```terraform
resource "tanzu-mission-control_mutation_policy" "cluster_pod_security_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "test"
    }
  }

  spec {
    input {
      pod_security {
        allow_privilege_escalation {
          condition = "Always"
          value     = true
        }
        capabilities_add {
          operation = "merge"
          values    = ["AUDIT_CONTROL", "AUDIT_WRITE"]
        }
        capabilities_drop {
          operation = "merge"
          values    = ["AUDIT_WRITE"]
        }
        fs_group {
          condition = "Always"
          value     = 0
        }
        privileged {
          condition = "Always"
          value     = true
        }
        read_only_root_filesystem {
          condition = "Always"
          value     = true
        }
        run_as_group {
          condition = "Always"
          value     = 0
        }
        run_as_non_root {
          condition = "Always"
          value     = true
        }
        run_as_user {
          condition = "Always"
          value     = 0
        }
        se_linux_options {
          condition = "Always"
          level     = "test"
          user      = "test"
          role      = "test"
          type      = "test"
        }
        supplemental_groups {
          condition = "Always"
          values    = [0, 1, 2, 3]
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
```

## Cluster group scoped annotation Mutation Policy

### Example Usage

```terraform
resource "tanzu-mission-control_mutation_policy" "cluster_group_annotation_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      annotation {
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
        scope = "*"
        annotation {
          key   = "test"
          value = "optional"
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
```

## Cluster group scoped label Mutation Policy

### Example Usage

```terraform
resource "tanzu-mission-control_mutation_policy" "cluster_group_label_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      label {
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
        scope = "*"
        label {
          key   = "test"
          value = "optional"
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
```

## Cluster group scoped pod-security Mutation Policy

### Example Usage

```terraform
resource "tanzu-mission-control_mutation_policy" "cluster_group_pod_security_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      pod_security {
        allow_privilege_escalation {
          condition = "Always"
          value     = true
        }
        capabilities_add {
          operation = "merge"
          values    = ["AUDIT_CONTROL", "AUDIT_WRITE"]
        }
        capabilities_drop {
          operation = "merge"
          values    = ["AUDIT_WRITE"]
        }
        fs_group {
          condition = "Always"
          value     = 0
        }
        privileged {
          condition = "Always"
          value     = true
        }
        read_only_root_filesystem {
          condition = "Always"
          value     = true
        }
        run_as_group {
          condition = "Always"
          value     = 0
        }
        run_as_non_root {
          condition = "Always"
          value     = true
        }
        run_as_user {
          condition = "Always"
          value     = 0
        }
        se_linux_options {
          condition = "Always"
          level     = "test"
          user      = "test"
          role      = "test"
          type      = "test"
        }
        supplemental_groups {
          condition = "Always"
          values    = [0, 1, 2, 3]
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
```

## Organization scoped annotation Mutation Policy

### Example Usage

```terraform
resource "tanzu-mission-control_mutation_policy" "org_annotation_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      annotation {
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
        scope = "*"
        annotation {
          key   = "test"
          value = "optional"
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
```

## Organization scoped label Mutation Policy

### Example Usage

```terraform
resource "tanzu-mission-control_mutation_policy" "org_label_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      label {
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
        scope = "Cluster"
        label {
          key   = "test"
          value = "optional"
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
```

## Organization scoped pod-security Mutation Policy

### Example Usage

```terraform
resource "tanzu-mission-control_mutation_policy" "org_pod_security_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      pod_security {
        allow_privilege_escalation {
          condition = "Always"
          value     = true
        }
        capabilities_add {
          operation = "merge"
          values    = ["AUDIT_CONTROL", "AUDIT_WRITE"]
        }
        capabilities_drop {
          operation = "merge"
          values    = ["AUDIT_WRITE"]
        }
        fs_group {
          condition = "Always"
          value     = 0
        }
        privileged {
          condition = "Always"
          value     = true
        }
        read_only_root_filesystem {
          condition = "Always"
          value     = true
        }
        run_as_group {
          condition = "Always"
          value     = 0
        }
        run_as_non_root {
          condition = "Always"
          value     = true
        }
        run_as_user {
          condition = "Always"
          value     = 0
        }
        se_linux_options {
          condition = "Always"
          level     = "level_test"
          user      = "user_test"
          role      = "role_test"
          type      = "type_test"
        }
        supplemental_groups {
          condition = "Always"
          values    = [0, 1, 2, 3]
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the mutation policy
- `scope` (Block List, Min: 1, Max: 1) Scope for the custom, security, image, network, namespace quota and mutation policy, having one of the valid scopes for custom, security, mutation, and namespace quota policy: cluster, cluster_group or organization and valid scopes for image and network policy: workspace or organization. (see [below for nested schema](#nestedblock--scope))
- `spec` (Block List, Min: 1, Max: 1) Spec for the mutation policy (see [below for nested schema](#nestedblock--spec))

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `cluster` (Block List, Max: 1) The schema for cluster policy full name (see [below for nested schema](#nestedblock--scope--cluster))
- `cluster_group` (Block List, Max: 1) The schema for cluster group policy full name (see [below for nested schema](#nestedblock--scope--cluster_group))
- `organization` (Block List, Max: 1) The schema for organization policy full name (see [below for nested schema](#nestedblock--scope--organization))
- `workspace` (Block List, Max: 1) The schema for workspace policy full name (see [below for nested schema](#nestedblock--scope--workspace))

<a id="nestedblock--scope--cluster"></a>
### Nested Schema for `scope.cluster`

Required:

- `name` (String) Name of this cluster

Optional:

- `management_cluster_name` (String) Name of the management cluster
- `provisioner_name` (String) Provisioner of the cluster


<a id="nestedblock--scope--cluster_group"></a>
### Nested Schema for `scope.cluster_group`

Required:

- `cluster_group` (String) Name of this cluster group


<a id="nestedblock--scope--organization"></a>
### Nested Schema for `scope.organization`

Required:

- `organization` (String) ID of this organization


<a id="nestedblock--scope--workspace"></a>
### Nested Schema for `scope.workspace`

Required:

- `workspace` (String) Name of this workspace



<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `input` (Block List, Min: 1, Max: 1) Input for the mutation policy. (see [below for nested schema](#nestedblock--spec--input))

Optional:

- `namespace_selector` (Block List, Max: 1) Label based Namespace Selector for the policy (see [below for nested schema](#nestedblock--spec--namespace_selector))

<a id="nestedblock--spec--input"></a>
### Nested Schema for `spec.input`

Optional:

- `annotation` (Block List, Max: 1) The input schema for custom policy tmc_block_nodeport_service recipe version v1 (see [below for nested schema](#nestedblock--spec--input--annotation))
- `label` (Block List, Max: 1) The input schema for custom policy tmc_block_nodeport_service recipe version v1 (see [below for nested schema](#nestedblock--spec--input--label))
- `pod_security` (Block List, Max: 1) The pod security schema (see [below for nested schema](#nestedblock--spec--input--pod_security))

<a id="nestedblock--spec--input--annotation"></a>
### Nested Schema for `spec.input.annotation`

Required:

- `target_kubernetes_resources` (Block List, Min: 1) A list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. (see [below for nested schema](#nestedblock--spec--input--annotation--target_kubernetes_resources))

Optional:

- `annotation` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--annotation--annotation))
- `scope` (String) Scope

<a id="nestedblock--spec--input--annotation--target_kubernetes_resources"></a>
### Nested Schema for `spec.input.annotation.target_kubernetes_resources`

Required:

- `api_groups` (List of String) APIGroup is a group containing the resource type.
- `kinds` (List of String) Kind is the name of the object schema (resource type).


<a id="nestedblock--spec--input--annotation--annotation"></a>
### Nested Schema for `spec.input.annotation.annotation`

Required:

- `key` (String)
- `value` (String)



<a id="nestedblock--spec--input--label"></a>
### Nested Schema for `spec.input.label`

Required:

- `target_kubernetes_resources` (Block List, Min: 1) A list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. (see [below for nested schema](#nestedblock--spec--input--label--target_kubernetes_resources))

Optional:

- `label` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--label--label))
- `scope` (String) Scope

<a id="nestedblock--spec--input--label--target_kubernetes_resources"></a>
### Nested Schema for `spec.input.label.target_kubernetes_resources`

Required:

- `api_groups` (List of String) APIGroup is a group containing the resource type.
- `kinds` (List of String) Kind is the name of the object schema (resource type).


<a id="nestedblock--spec--input--label--label"></a>
### Nested Schema for `spec.input.label.label`

Required:

- `key` (String)
- `value` (String)



<a id="nestedblock--spec--input--pod_security"></a>
### Nested Schema for `spec.input.pod_security`

Optional:

- `allow_privilege_escalation` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--pod_security--allow_privilege_escalation))
- `capabilities_add` (Block List, Max: 1) Run as user (see [below for nested schema](#nestedblock--spec--input--pod_security--capabilities_add))
- `capabilities_drop` (Block List, Max: 1) Run as user (see [below for nested schema](#nestedblock--spec--input--pod_security--capabilities_drop))
- `fs_group` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--pod_security--fs_group))
- `privileged` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--pod_security--privileged))
- `read_only_root_filesystem` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--pod_security--read_only_root_filesystem))
- `run_as_group` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--pod_security--run_as_group))
- `run_as_non_root` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--pod_security--run_as_non_root))
- `run_as_user` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--pod_security--run_as_user))
- `se_linux_options` (Block List) Allowed selinux options (see [below for nested schema](#nestedblock--spec--input--pod_security--se_linux_options))
- `supplemental_groups` (Block List, Max: 1) (see [below for nested schema](#nestedblock--spec--input--pod_security--supplemental_groups))

<a id="nestedblock--spec--input--pod_security--allow_privilege_escalation"></a>
### Nested Schema for `spec.input.pod_security.allow_privilege_escalation`

Required:

- `condition` (String)
- `value` (Boolean)


<a id="nestedblock--spec--input--pod_security--capabilities_add"></a>
### Nested Schema for `spec.input.pod_security.capabilities_add`

Required:

- `values` (List of String) Values is an array of string values

Optional:

- `operation` (String) Rule


<a id="nestedblock--spec--input--pod_security--capabilities_drop"></a>
### Nested Schema for `spec.input.pod_security.capabilities_drop`

Required:

- `values` (List of String) Values is an array of string values

Optional:

- `operation` (String) Rule


<a id="nestedblock--spec--input--pod_security--fs_group"></a>
### Nested Schema for `spec.input.pod_security.fs_group`

Required:

- `condition` (String)
- `value` (Number)


<a id="nestedblock--spec--input--pod_security--privileged"></a>
### Nested Schema for `spec.input.pod_security.privileged`

Required:

- `condition` (String)
- `value` (Boolean)


<a id="nestedblock--spec--input--pod_security--read_only_root_filesystem"></a>
### Nested Schema for `spec.input.pod_security.read_only_root_filesystem`

Required:

- `condition` (String)
- `value` (Boolean)


<a id="nestedblock--spec--input--pod_security--run_as_group"></a>
### Nested Schema for `spec.input.pod_security.run_as_group`

Required:

- `condition` (String)
- `value` (Number)


<a id="nestedblock--spec--input--pod_security--run_as_non_root"></a>
### Nested Schema for `spec.input.pod_security.run_as_non_root`

Required:

- `condition` (String)
- `value` (Boolean)


<a id="nestedblock--spec--input--pod_security--run_as_user"></a>
### Nested Schema for `spec.input.pod_security.run_as_user`

Required:

- `condition` (String)
- `value` (Number)


<a id="nestedblock--spec--input--pod_security--se_linux_options"></a>
### Nested Schema for `spec.input.pod_security.se_linux_options`

Optional:

- `condition` (String) SELinux condition
- `level` (String) SELinux level
- `role` (String) SELinux role
- `type` (String) SELinux type
- `user` (String) SELinux user


<a id="nestedblock--spec--input--pod_security--supplemental_groups"></a>
### Nested Schema for `spec.input.pod_security.supplemental_groups`

Required:

- `values` (List of Number)

Optional:

- `condition` (String)




<a id="nestedblock--spec--namespace_selector"></a>
### Nested Schema for `spec.namespace_selector`

Required:

- `match_expressions` (Block List, Min: 1) Match expressions is a list of label selector requirements, the requirements are ANDed (see [below for nested schema](#nestedblock--spec--namespace_selector--match_expressions))

<a id="nestedblock--spec--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.namespace_selector.match_expressions`

Required:

- `values` (List of String) Values is an array of string values

Optional:

- `key` (String) Key is the label key that the selector applies to
- `operator` (String) Operator represents a key's relationship to a set of values




<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource
