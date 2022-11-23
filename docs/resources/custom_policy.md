---
Title: "Custom Policy Resource"
Description: |-
    Creating the Tanzu Kubernetes custom policy resource.
---

# Custom Policy

The `tanzu-mission-control_custom_policy` resource enables you to attach a custom policy with an input recipe to a particular scope for management through Tanzu Mission Control.
Custom policies allow you to implement additional business rules, using templates that you define, to enforce policies that are not already addressed using the other built-in policy types.
For more information, see [Creating Customized Policies][create-customized-policies] using VMware Tanzu Mission Control.

[create-customized-policies]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-1FF7A1E5-8456-4EF4-A532-9CF31BE88EAA.html

## Input Recipe

In the Tanzu Mission Control custom policy resource, there are six system defined types of custom templates that you can use:
- **tmc-block-nodeport-service**
- **tmc_block_resources**
- **tmc_block_rolebinding_subjects**
- **tmc_external_ips**
- **tmc_https_ingress**
- **tmc_require_labels**

## Policy Scope and Inheritance

In the Tanzu Mission Control resource hierarchy, there are three levels at which you can specify custom policy resources:
- **organization** - `organization` block under `scope` sub-resource
- **object groups** - `cluster_group` block under `scope` sub-resource
- **Kubernetes objects** - `cluster` block under `scope` sub-resource

In addition to the direct policy defined for a given object, each object has inherited policies described in the parent objects. For example, a cluster has a direct policy and inherited policies from the cluster group and organization to which it is attached.

**Note:**
The scope parameter is mandatory in the schema and the user needs to add one of the defined scopes to the script for the provider to function.
Only one scope per resource is allowed.

## Cluster scoped TMC-block-nodeport-service Custom Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control custom policy with tmc-block-nodeport-service input recipe.
This policy is applied to a cluster with the tmc-block-nodeport-service configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_scoped_tmc-block-nodeport-service_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_block_nodeport_service {
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

## Cluster scoped TMC-block-resources Custom Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control custom policy with tmc-block-resources input recipe.
This policy is applied to a cluster with the tmc-block-resources configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_scoped_tmc-block-resources_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
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

## Cluster scoped TMC-block-rolebinding-subjects Custom Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control custom policy with tmc-block-rolebinding-subjects input recipe.
This policy is applied to a cluster with the tmc-block-rolebinding-subjects configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_scoped_tmc-block-rolebinding-subjects_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_block_rolebinding_subjects {
        audit = false
        parameters {
          disallowed_subjects {
            kind = "node"
            name = "subject-1"
          }
        }
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

## Cluster scoped TMC-external-ips Custom Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control custom policy with tmc-external-ips input recipe.
This policy is applied to a cluster with the tmc-external-ips configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_scoped_tmc-external-ips_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_external_ips {
        audit = false
        parameters {
          allowed_ips = [
            "127.0.0.1",
          ]
        }
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

## Cluster scoped TMC-https-ingress Custom Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control custom policy with tmc-https-ingress input recipe.
This policy is applied to a cluster with the tmc-https-ingress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_scoped_tmc-https-ingress_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_https_ingress {
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

## Cluster scoped TMC-require-labels Custom Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control custom policy with tmc-require-labels input recipe.
This policy is applied to a cluster with the tmc-require-labels configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_scoped_tmc-require-labels_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_require_labels {
        audit = false
        parameters {
          labels {
            key   = "test"
            value = "optional"
          }
        }
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

## Cluster group scoped TMC-block-nodeport-service Custom Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control custom policy with tmc-block-nodeport-service input recipe.
This policy is applied to a cluster group with the tmc-block-nodeport-service configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-block-nodeport-service_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_block_nodeport_service {
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

## Cluster group scoped TMC-block-resources Custom Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control custom policy with tmc-block-resources input recipe.
This policy is applied to a cluster group with the tmc-block-resources configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-block-resources_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
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

## Cluster group scoped TMC-block-rolebinding-subjects Custom Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control custom policy with tmc-block-rolebinding-subjects input recipe.
This policy is applied to a cluster group with the tmc-block-rolebinding-subjects configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-block-rolebinding-subjects_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_block_rolebinding_subjects {
        audit = false
        parameters {
          disallowed_subjects {
            kind = "node"
            name = "subject-1"
          }
        }
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

## Cluster group scoped TMC-external-ips Custom Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control custom policy with tmc-external-ips input recipe.
This policy is applied to a cluster group with the tmc-external-ips configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-external-ips_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_external_ips {
        audit = false
        parameters {
          allowed_ips = [
            "127.0.0.1",
          ]
        }
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

## Cluster group scoped TMC-https-ingress Custom Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control custom policy with tmc-https-ingress input recipe.
This policy is applied to a cluster group with the tmc-https-ingress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-https-ingress_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_https_ingress {
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

## Cluster group scoped TMC-require-labels Custom Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control custom policy with tmc-require-labels input recipe.
This policy is applied to a cluster group with the tmc-require-labels configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-require-labels_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_require_labels {
        audit = false
        parameters {
          labels {
            key   = "test"
            value = "optional"
          }
        }
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

## Organization scoped TMC-block-nodeport-service Custom Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control custom policy with tmc-block-nodeport-service input recipe.
This policy is applied to a organization with the tmc-block-nodeport-service configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "organization_scoped_tmc-block-nodeport-service_custom_policy" {
  name = "tf-custom-test"

  scope {
    organization {
      organization = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_block_nodeport_service {
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

## Organization scoped TMC-block-resources Custom Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control custom policy with tmc-block-resources input recipe.
This policy is applied to a organization with the tmc-block-resources configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "organization_scoped_tmc-block-resources_custom_policy" {
  name = "tf-custom-test"

  scope {
    organization {
      organization = "tf-create-test"
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

## Organization scoped TMC-block-rolebinding-subjects Custom Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control custom policy with tmc-block-rolebinding-subjects input recipe.
This policy is applied to a organization with the tmc-block-rolebinding-subjects configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "organization_scoped_tmc-block-rolebinding-subjects_custom_policy" {
  name = "tf-custom-test"

  scope {
    organization {
      organization = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_block_rolebinding_subjects {
        audit = false
        parameters {
          disallowed_subjects {
            kind = "node"
            name = "subject-1"
          }
        }
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

## Organization scoped TMC-external-ips Custom Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control custom policy with tmc-external-ips input recipe.
This policy is applied to a organization with the tmc-external-ips configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "organization_scoped_tmc-external-ips_custom_policy" {
  name = "tf-custom-test"

  scope {
    organization {
      organization = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_external_ips {
        audit = false
        parameters {
          allowed_ips = [
            "127.0.0.1",
          ]
        }
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

## Organization scoped TMC-https-ingress Custom Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control custom policy with tmc-https-ingress input recipe.
This policy is applied to a organization with the tmc-https-ingress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "organization_scoped_tmc-https-ingress_custom_policy" {
  name = "tf-custom-test"

  scope {
    organization {
      organization = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_https_ingress {
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

## Organization scoped TMC-require-labels Custom Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control custom policy with tmc-require-labels input recipe.
This policy is applied to a organization with the tmc-require-labels configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "organization_scoped_tmc-require-labels_custom_policy" {
  name = "tf-custom-test"

  scope {
    organization {
      organization = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_require_labels {
        audit = false
        parameters {
          labels {
            key   = "test"
            value = "optional"
          }
        }
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

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the custom policy
- `scope` (Block List, Min: 1, Max: 1) Scope for the security and custom policy, having one of the valid scopes: cluster, cluster_group or organization. (see [below for nested schema](#nestedblock--scope))
- `spec` (Block List, Min: 1, Max: 1) Spec for the custom policy (see [below for nested schema](#nestedblock--spec))

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



<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `input` (Block List, Min: 1, Max: 1) Input for the custom policy, having one of the valid recipes: tmc_block_nodeport_service, tmc_block_resources, tmc_block_rolebinding_subjects, tmc_external_ips, tmc_https_ingress or tmc_require_labels. (see [below for nested schema](#nestedblock--spec--input))

Optional:

- `namespace_selector` (Block List, Max: 1) Label based Namespace Selector for the policy (see [below for nested schema](#nestedblock--spec--namespace_selector))

<a id="nestedblock--spec--input"></a>
### Nested Schema for `spec.input`

Optional:

- `tmc_block_nodeport_service` (Block List, Max: 1) The input schema for custom policy tmc_block_nodeport_service recipe version v1 (see [below for nested schema](#nestedblock--spec--input--tmc_block_nodeport_service))
- `tmc_block_resources` (Block List, Max: 1) The input schema for custom policy tmc_block_resources recipe version v1 (see [below for nested schema](#nestedblock--spec--input--tmc_block_resources))
- `tmc_block_rolebinding_subjects` (Block List, Max: 1) The input schema for custom policy tmc_block_rolebinding_subjects recipe version v1 (see [below for nested schema](#nestedblock--spec--input--tmc_block_rolebinding_subjects))
- `tmc_external_ips` (Block List, Max: 1) The input schema for custom policy tmc_external_ips recipe version v1 (see [below for nested schema](#nestedblock--spec--input--tmc_external_ips))
- `tmc_https_ingress` (Block List, Max: 1) The input schema for custom policy tmc_https_ingress recipe version v1 (see [below for nested schema](#nestedblock--spec--input--tmc_https_ingress))
- `tmc_require_labels` (Block List, Max: 1) The input schema for custom policy tmc_require_labels recipe version v1 (see [below for nested schema](#nestedblock--spec--input--tmc_require_labels))

<a id="nestedblock--spec--input--tmc_block_nodeport_service"></a>
### Nested Schema for `spec.input.tmc_block_nodeport_service`

Required:

- `target_kubernetes_resources` (Block List, Min: 1) A list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. (see [below for nested schema](#nestedblock--spec--input--tmc_block_nodeport_service--target_kubernetes_resources))

Optional:

- `audit` (Boolean) Audit (dry-run).

<a id="nestedblock--spec--input--tmc_block_nodeport_service--target_kubernetes_resources"></a>
### Nested Schema for `spec.input.tmc_block_nodeport_service.target_kubernetes_resources`

Required:

- `api_groups` (List of String) APIGroup is a group containing the resource type.
- `kinds` (List of String) Kind is the name of the object schema (resource type).



<a id="nestedblock--spec--input--tmc_block_resources"></a>
### Nested Schema for `spec.input.tmc_block_resources`

Required:

- `target_kubernetes_resources` (Block List, Min: 1) A list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. (see [below for nested schema](#nestedblock--spec--input--tmc_block_resources--target_kubernetes_resources))

Optional:

- `audit` (Boolean) Audit (dry-run).

<a id="nestedblock--spec--input--tmc_block_resources--target_kubernetes_resources"></a>
### Nested Schema for `spec.input.tmc_block_resources.target_kubernetes_resources`

Required:

- `api_groups` (List of String) APIGroup is a group containing the resource type.
- `kinds` (List of String) Kind is the name of the object schema (resource type).



<a id="nestedblock--spec--input--tmc_block_rolebinding_subjects"></a>
### Nested Schema for `spec.input.tmc_block_rolebinding_subjects`

Required:

- `parameters` (Block List, Min: 1) Parameters. (see [below for nested schema](#nestedblock--spec--input--tmc_block_rolebinding_subjects--parameters))
- `target_kubernetes_resources` (Block List, Min: 1) A list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. (see [below for nested schema](#nestedblock--spec--input--tmc_block_rolebinding_subjects--target_kubernetes_resources))

Optional:

- `audit` (Boolean) Audit (dry-run).

<a id="nestedblock--spec--input--tmc_block_rolebinding_subjects--parameters"></a>
### Nested Schema for `spec.input.tmc_block_rolebinding_subjects.parameters`

Required:

- `disallowed_subjects` (Block List, Min: 1) Disallowed Subjects. (see [below for nested schema](#nestedblock--spec--input--tmc_block_rolebinding_subjects--parameters--disallowed_subjects))

<a id="nestedblock--spec--input--tmc_block_rolebinding_subjects--parameters--disallowed_subjects"></a>
### Nested Schema for `spec.input.tmc_block_rolebinding_subjects.parameters.disallowed_subjects`

Required:

- `kind` (String) The kind of subject to disallow, can be User/Group/ServiceAccount.
- `name` (String) The name of the subject to disallow.



<a id="nestedblock--spec--input--tmc_block_rolebinding_subjects--target_kubernetes_resources"></a>
### Nested Schema for `spec.input.tmc_block_rolebinding_subjects.target_kubernetes_resources`

Required:

- `api_groups` (List of String) APIGroup is a group containing the resource type.
- `kinds` (List of String) Kind is the name of the object schema (resource type).



<a id="nestedblock--spec--input--tmc_external_ips"></a>
### Nested Schema for `spec.input.tmc_external_ips`

Required:

- `parameters` (Block List, Min: 1) Parameters. (see [below for nested schema](#nestedblock--spec--input--tmc_external_ips--parameters))
- `target_kubernetes_resources` (Block List, Min: 1) A list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. (see [below for nested schema](#nestedblock--spec--input--tmc_external_ips--target_kubernetes_resources))

Optional:

- `audit` (Boolean) Audit (dry-run).

<a id="nestedblock--spec--input--tmc_external_ips--parameters"></a>
### Nested Schema for `spec.input.tmc_external_ips.parameters`

Required:

- `allowed_ips` (List of String) Allowed IPs.


<a id="nestedblock--spec--input--tmc_external_ips--target_kubernetes_resources"></a>
### Nested Schema for `spec.input.tmc_external_ips.target_kubernetes_resources`

Required:

- `api_groups` (List of String) APIGroup is a group containing the resource type.
- `kinds` (List of String) Kind is the name of the object schema (resource type).



<a id="nestedblock--spec--input--tmc_https_ingress"></a>
### Nested Schema for `spec.input.tmc_https_ingress`

Required:

- `target_kubernetes_resources` (Block List, Min: 1) A list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. (see [below for nested schema](#nestedblock--spec--input--tmc_https_ingress--target_kubernetes_resources))

Optional:

- `audit` (Boolean) Audit (dry-run).

<a id="nestedblock--spec--input--tmc_https_ingress--target_kubernetes_resources"></a>
### Nested Schema for `spec.input.tmc_https_ingress.target_kubernetes_resources`

Required:

- `api_groups` (List of String) APIGroup is a group containing the resource type.
- `kinds` (List of String) Kind is the name of the object schema (resource type).



<a id="nestedblock--spec--input--tmc_require_labels"></a>
### Nested Schema for `spec.input.tmc_require_labels`

Required:

- `parameters` (Block List, Min: 1) Parameters. (see [below for nested schema](#nestedblock--spec--input--tmc_require_labels--parameters))
- `target_kubernetes_resources` (Block List, Min: 1) A list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. (see [below for nested schema](#nestedblock--spec--input--tmc_require_labels--target_kubernetes_resources))

Optional:

- `audit` (Boolean) Audit (dry-run).

<a id="nestedblock--spec--input--tmc_require_labels--parameters"></a>
### Nested Schema for `spec.input.tmc_require_labels.parameters`

Required:

- `labels` (Block List, Min: 1) Labels. (see [below for nested schema](#nestedblock--spec--input--tmc_require_labels--parameters--labels))

<a id="nestedblock--spec--input--tmc_require_labels--parameters--labels"></a>
### Nested Schema for `spec.input.tmc_require_labels.parameters.labels`

Required:

- `key` (String) The label key to enforce.

Optional:

- `value` (String) Optional label value to enforce (if left empty, only key will be enforced).



<a id="nestedblock--spec--input--tmc_require_labels--target_kubernetes_resources"></a>
### Nested Schema for `spec.input.tmc_require_labels.target_kubernetes_resources`

Required:

- `api_groups` (List of String) APIGroup is a group containing the resource type.
- `kinds` (List of String) Kind is the name of the object schema (resource type).




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
