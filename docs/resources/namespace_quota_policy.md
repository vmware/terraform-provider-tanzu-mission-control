---
Title: "Namespace Quota Policy Resource"
Description: |-
    Creating the Tanzu Kubernetes namespace quota policy resource.
---

# Namespace Quota Policy

The `tanzu-mission-control_namespace_quota_policy` resource enables you to attach a namespace quota policy with an input recipe to a particular scope for management through Tanzu Mission Control.

Namespace quota policies allow you to constrain the resources used in your clusters, as aggregate quantities across specified namespaces, using preconfigured and custom templates.
For more information, see [Managing Resource Consumption in Your Clusters][managing-resource-consumption] in using VMware Tanzu Mission Control.

[managing-resource-consumption]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-1905352C-856F-4D06-BB86-426F90486C32.html

## Input Recipe

In the Tanzu Mission Control namespace quota policy resource, there are four set of recipe templates that you can use:
- **small** - The small template is a preconfigured set of resource limits with constraints as CPU requests = 0.5 vCPU, Memory requests = 512 MB, CPU limits = 1 vCPU, Memory limits = 2 GB.
- **medium** - The medium template is a preconfigured set of resource limits with constraints as CPU requests = 1 vCPU, Memory requests = 1 GB, CPU limits = 2 vCPU, Memory limits = 4 GB.
- **large** - The large template is a preconfigured set of resource limits with constraints as CPU requests = 2 vCPU, Memory requests = 2 GB, CPU limits = 4 vCPU, Memory limits = 8 GB.
- **custom** - The custom template allows you to specify the quantity limits of various resource types.

## Policy Scope and Inheritance

In the Tanzu Mission Control resource hierarchy, there are three levels at which you can specify quota policy resources:
- **organization** - `organization` block under `scope` sub-resource
- **object groups** - `cluster_group` block under `scope` sub-resource
- **Kubernetes objects** - `cluster` block under `scope` sub-resource

In addition to the direct policy defined for a given object, each object has inherited policies described in the parent objects. For example, a cluster has a direct policy and inherited policies from the cluster group and organization to which it is attached.

**Note:**
The scope parameter is mandatory in the schema and the user needs to add one of the defined scopes to the script for the provider to function.
Only one scope per resource is allowed.

## Cluster scoped Small Namespace Quota Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control namespace quota policy with small input recipe.
This policy is applied to a cluster with the small configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_scoped_small_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      small {} // Pre-defined parameters for Small Namespace quota Policy: CPU requests = 0.5 vCPU, Memory requests = 512 MB, CPU limits = 1 vCPU, Memory limits = 2 GB
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


## Cluster scoped Medium Namespace Quota Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control namespace quota policy with medium input recipe.
This policy is applied to a cluster with the medium configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_scoped_medium_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      medium {} // Pre-defined parameters for Medium Namespace quota Policy: CPU requests = 1 vCPU, Memory requests = 1 GB, CPU limits = 2 vCPU, Memory limits = 4 GB
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


## Cluster scoped Large Namespace Quota Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control namespace quota policy with large input recipe.
This policy is applied to a cluster with the large configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_scoped_large_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      large {} // Pre-defined parameters for Large Namespace quota Policy: CPU requests = 2 vCPU, Memory requests = 2 GB, CPU limits = 4 vCPU, Memory limits = 8 GB
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

## Cluster scoped Custom Namespace Quota Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control namespace quota policy with custom input recipe.
This policy is applied to a cluster with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_scoped_custom_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      custom {
        limits_cpu               = "4"
        limits_memory            = "8Mi"
        persistent_volume_claims = 2
        persistent_volume_claims_per_class = {
          ab : 2
          cd : 4
        }
        requests_cpu     = "2"
        requests_memory  = "4Mi"
        requests_storage = "2G"
        requests_storage_per_class = {
          test : "2G"
          twt : "4G"
        }
        resource_counts = {
          pods : 2
        }
      }
    }
  }
}
```


## Cluster group scoped Small Namespace Quota Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control namespace quota policy with small input recipe.
This policy is applied to a cluster group with the small configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_group_scoped_small_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      small {} // Pre-defined parameters for Small Namespace quota Policy: CPU requests = 0.5 vCPU, Memory requests = 512 MB, CPU limits = 1 vCPU, Memory limits = 2 GB
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

## Cluster group scoped Medium Namespace Quota Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control namespace quota policy with medium input recipe.
This policy is applied to a cluster group with the medium configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_group_scoped_medium_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      medium {} // Pre-defined parameters for Medium Namespace quota Policy: CPU requests = 1 vCPU, Memory requests = 1 GB, CPU limits = 2 vCPU, Memory limits = 4 GB
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

## Cluster group scoped Large Namespace Quota Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control namespace quota policy with large input recipe.
This policy is applied to a cluster group with the large configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_group_scoped_large_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      large {} // Pre-defined parameters for Large Namespace quota Policy: CPU requests = 2 vCPU, Memory requests = 2 GB, CPU limits = 4 vCPU, Memory limits = 8 GB
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


## Cluster group scoped Custom Namespace Quota Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control namespace quota policy with custom input recipe.
This policy is applied to a cluster group with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_group_scoped_custom_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      custom {
        limits_cpu               = "4"
        limits_memory            = "8Mi"
        persistent_volume_claims = 2
        persistent_volume_claims_per_class = {
          ab : 2
          cd : 4
        }
        requests_cpu     = "2"
        requests_memory  = "4Mi"
        requests_storage = "2G"
        requests_storage_per_class = {
          test : "2G"
          twt : "4G"
        }
        resource_counts = {
          pods : 2
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


## Organization scoped Small Namespace Quota Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control namespace quota policy with small input recipe.
This policy is applied to a organization with the small configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "organization_scoped_small_quota_policy" {
  name = "tf-qt-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      small {} // Pre-defined parameters for Small Namespace quota Policy: CPU requests = 0.5 vCPU, Memory requests = 512 MB, CPU limits = 1 vCPU, Memory limits = 2 GB
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


## Organization scoped Medium Namespace Quota Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control namespace quota policy with medium input recipe.
This policy is applied to a organization with the medium configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "organization_scoped_medium_quota_policy" {
  name = "tf-qt-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      medium {} // Pre-defined parameters for Medium Namespace quota Policy: CPU requests = 1 vCPU, Memory requests = 1 GB, CPU limits = 2 vCPU, Memory limits = 4 GB
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

## Organization scoped Large Namespace Quota Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control namespace quota policy with large input recipe.
This policy is applied to a organization with the large configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "organization_scoped_large_quota_policy" {
  name = "tf-qt-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      large {} // Pre-defined parameters for Large Namespace quota Policy: CPU requests = 2 vCPU, Memory requests = 2 GB, CPU limits = 4 vCPU, Memory limits = 8 GB
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

## Organization scoped Custom Namespace Quota Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control namespace quota policy with custom input recipe.
This policy is applied to a organization with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "organization_scoped_custom_quota_policy" {
  name = "tf-qt-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom {
        limits_cpu               = "4"
        limits_memory            = "8Mi"
        persistent_volume_claims = 2
        persistent_volume_claims_per_class = {
          ab : 2
          cd : 4
        }
        requests_cpu     = "2"
        requests_memory  = "4Mi"
        requests_storage = "2G"
        requests_storage_per_class = {
          test : "2G"
          twt : "4G"
        }
        resource_counts = {
          pods : 2
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

- `name` (String) Name of the namespace quota policy
- `scope` (Block List, Min: 1, Max: 1) Scope for the quota policy, having one of the valid scopes: cluster, cluster_group, organization. (see [below for nested schema](#nestedblock--scope))
- `spec` (Block List, Min: 1, Max: 1) Spec for the namespace namespace quota policy (see [below for nested schema](#nestedblock--spec))

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

- `input` (Block List, Min: 1, Max: 1) Input for the namespace quota policy, having one of the valid recipes: small, medium, large or custom. (see [below for nested schema](#nestedblock--spec--input))

Optional:

- `namespace_selector` (Block List, Max: 1) Label based Namespace Selector for the policy (see [below for nested schema](#nestedblock--spec--namespace_selector))

<a id="nestedblock--spec--input"></a>
### Nested Schema for `spec.input`

Optional:

- `custom` (Block List, Max: 1) The input schema for namespace quota policy custom recipe version v1 (see [below for nested schema](#nestedblock--spec--input--custom))
- `large` (Block List, Max: 1) The input schema for namespace quota policy large recipe version v1 (see [below for nested schema](#nestedblock--spec--input--large))
- `medium` (Block List, Max: 1) The input schema for namespace quota policy medium recipe version v1 (see [below for nested schema](#nestedblock--spec--input--medium))
- `small` (Block List, Max: 1) The input schema for namespace quota policy small recipe version v1 (see [below for nested schema](#nestedblock--spec--input--small))

<a id="nestedblock--spec--input--custom"></a>
### Nested Schema for `spec.input.custom`

Optional:

- `limits_cpu` (String) The sum of CPU limits across all pods in a non-terminal state cannot exceed this value
- `limits_memory` (String) The sum of memory limits across all pods in a non-terminal state cannot exceed this value
- `persistent_volume_claims` (Number) The total number of PersistentVolumeClaims that can exist in a namespace
- `persistent_volume_claims_per_class` (Map of Number) Across all persistent volume claims associated with each storage class, the total number of persistent volume claims that can exist in the namespace
- `requests_cpu` (String) The sum of CPU requests across all pods in a non-terminal state cannot exceed this value
- `requests_memory` (String) The sum of memory requests across all pods in a non-terminal state cannot exceed this value
- `requests_storage` (String) The sum of storage requests across all persistent volume claims cannot exceed this value
- `requests_storage_per_class` (Map of String) Across all persistent volume claims associated with each storage class, the sum of storage requests cannot exceed this value
- `resource_counts` (Map of Number) The total number of Services of the given type that can exist in a namespace


<a id="nestedblock--spec--input--large"></a>
### Nested Schema for `spec.input.large`


<a id="nestedblock--spec--input--medium"></a>
### Nested Schema for `spec.input.medium`


<a id="nestedblock--spec--input--small"></a>
### Nested Schema for `spec.input.small`



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