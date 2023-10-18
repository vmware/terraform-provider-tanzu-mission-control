---
Title: "Security Policy Resource"
Description: |-
    Creating the Tanzu Kubernetes security policy resource.
---

# Security Policy

The `tanzu-mission-control_security_policy` resource enables you to attach a security policy with an input recipe to a particular scope for management through Tanzu Mission Control.

Security policies allow you to manage the security context in which deployed pods operate in your clusters by imposing constraints that define what pods can do and which resources they can access. For more information, see [Pod Security Management.][pod-security-management]

[pod-security-management]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-6C65B33B-C1EA-465D-B909-3C4F51704C1A.html#GUID-6C65B33B-C1EA-465D-B909-3C4F51704C1A

## Input Recipe

In the Tanzu Mission Control security policy resource, there are three types of security templates that you can use:
- **baseline** - The Baseline template is a preconfigured set of constraints that prevent known privilege escalations but is less stringent than the Strict template to ease the adoption of the security policy for typical containerized workloads. The detailed options defined in this template are displayed on the form in the Tanzu Mission Control console.
- **custom** - The Custom template allows you to specify how to handle the various aspects of pod security for your clusters.
- **strict** - The Strict template is a preconfigured set of constraints that define a tight security context for pods in your clusters. The detailed options described in this template are displayed on the form in the Tanzu Mission Control console.

## Policy Scope and Inheritance

In the Tanzu Mission Control resource hierarchy, there are three levels at which you can specify security policy resources:
- **organization** - `organization` block under `scope` sub-resource
- **object groups** - `cluster_group` block under `scope` sub-resource
- **Kubernetes objects** - `cluster` block under `scope` sub-resource

In addition to the direct policy defined for a given object, each object has inherited policies described in the parent objects. For example, a cluster has a direct policy and inherited policies from the cluster group and organization to which it is attached.

**Note:**
The scope parameter is mandatory in the schema and the user needs to add one of the defined scopes to the script for the provider to function.
Only one scope per resource is allowed.

## Managing Pod Security

To use the **Tanzu Mission Control provider** for creating a security policy for an object, you must be associated with the `.admin` role for that object.
For more information, see [Managing Pod Security.][managing-pod-security]

[managing-pod-security]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-939955FC-17EF-4A84-B686-CAF0BBE018D4.html

## Cluster scoped Baseline Security Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control security policy with baseline input recipe.
This policy is applied to a cluster with the baseline configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_security_policy" "cluster_scoped_baseline_security_policy" {
  name = "tf-sp-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      baseline {
        audit              = true
        disable_native_psp = false
      }
    }

    namespace_selector {
      match_expressions {
        key      = "not-a-component"
        operator = "DoesNotExist"
        values   = []
      }
    }
  }
}
```


## Cluster scoped Custom Security Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control security policy with custom input recipe.
This policy is applied on a cluster with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_security_policy" "cluster_scoped_custom_security_policy" {
  name = "tf-sp-test"

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
        audit                        = true
        disable_native_psp           = false
        allow_privileged_containers  = true
        allow_privilege_escalation   = true
        allow_host_namespace_sharing = true
        allow_host_network           = true
        read_only_root_file_system   = true

        allowed_host_port_range {
          min = 3000
          max = 5000
        }

        allowed_volumes = [
          "configMap",
          "nfs",
          "vsphereVolume"
        ]

        run_as_user {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        run_as_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        supplemental_groups {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        fs_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        linux_capabilities {
          allowed_capabilities = [
            "CHOWN",
            "IPC_LOCK"
          ]
          required_drop_capabilities = [
            "SYS_TIME"
          ]
        }

        allowed_host_paths {
          path_prefix = "p1"
          read_only   = true
        }
        allowed_host_paths {
          path_prefix = "p2"
          read_only   = false
        }
        allowed_host_paths {
          path_prefix = "p3"
          read_only   = true
        }

        allowed_se_linux_options {
          level = "s0"
          role  = "sysadm_r"
          type  = "httpd_sys_content_t"
          user  = "root"
        }

        sysctls {
          forbidden_sysctls = [
            "kernel.msgmax",
            "kernel.sem"
          ]
        }

        seccomp {
          allowed_profiles = [
            "Localhost"
          ]
          allowed_localhost_files = [
            "profiles/audit.json",
            "profiles/violation.json"
          ]
        }
      }
    }
  }
}
```


## Cluster scoped Strict Security Policy

### Example Usage

```terraform
/*
Cluster scoped Tanzu Mission Control security policy with strict input recipe.
This policy is applied to a cluster with the strict configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_security_policy" "cluster_scoped_strict_security_policy" {
  name = "tf-sp-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      strict {
        audit              = true
        disable_native_psp = false
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "NotIn"
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


## Cluster group scoped Baseline Security Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control security policy with baseline input recipe.
This policy is applied to a cluster group with the baseline configuration option and is inherited by the clusters.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_security_policy" "cluster_group_scoped_baseline_security_policy" {
  name = "tf-sp-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
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


## Cluster group scoped Custom Security Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control security policy with a custom input recipe.
This policy is applied to a cluster group with the custom configuration option and is inherited by the clusters.
The defined scope and input blocks can be updated to change the policy's scope and recipe.
*/
resource "tanzu-mission-control_security_policy" "cluster_group_scoped_custom_security_policy" {
  name = "tf-sp-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      custom {
        audit                        = true
        disable_native_psp           = false
        allow_privileged_containers  = true
        allow_privilege_escalation   = true
        allow_host_namespace_sharing = true
        allow_host_network           = true
        read_only_root_file_system   = true

        allowed_host_port_range {
          min = 3000
          max = 5000
        }

        allowed_volumes = [
          "configMap",
          "nfs",
          "vsphereVolume"
        ]

        run_as_user {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        run_as_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        supplemental_groups {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        fs_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        linux_capabilities {
          allowed_capabilities = [
            "CHOWN",
            "IPC_LOCK"
          ]
          required_drop_capabilities = [
            "SYS_TIME"
          ]
        }

        allowed_host_paths {
          path_prefix = "p1"
          read_only   = true
        }
        allowed_host_paths {
          path_prefix = "p2"
          read_only   = false
        }
        allowed_host_paths {
          path_prefix = "p3"
          read_only   = true
        }

        allowed_se_linux_options {
          level = "s0"
          role  = "sysadm_r"
          type  = "httpd_sys_content_t"
          user  = "root"
        }

        sysctls {
          forbidden_sysctls = [
            "kernel.msgmax",
            "kernel.sem"
          ]
        }

        seccomp {
          allowed_profiles = [
            "Localhost"
          ]
          allowed_localhost_files = [
            "profiles/audit.json",
            "profiles/violation.json"
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


## Cluster group scoped Strict Security Policy

### Example Usage

```terraform
/*
Cluster group scoped Tanzu Mission Control security policy with a strict input recipe.
This policy is applied to a cluster group with the strict configuration option and is inherited by the clusters.
The defined scope and input blocks can be updated to change the policy's scope and recipe.
*/
resource "tanzu-mission-control_security_policy" "cluster_group_scoped_strict_security_policy" {
  name = "tf-sp-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      strict {
        audit              = false
        disable_native_psp = true
      }
    }
  }
}
```


## Organization scoped Baseline Security Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control security policy with baseline input recipe.
This policy is applied to an organization with the baseline configuration option and is inherited by the cluster groups and clusters.
The defined scope and input blocks can be updated to change the policy's scope and recipe.
*/
resource "tanzu-mission-control_security_policy" "organization_scoped_baseline_security_policy" {
  name = "tf-sp-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      baseline {
        audit              = false
        disable_native_psp = true
      }
    }
  }
}
```


## Organization scoped Custom Security Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control security policy with a custom input recipe.
This policy is applied to an organization with the custom configuration option and is inherited by the cluster groups and clusters.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_security_policy" "organization_scoped_custom_security_policy" {
  name = "tf-sp-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom {
        audit                        = true
        disable_native_psp           = false
        allow_privileged_containers  = true
        allow_privilege_escalation   = true
        allow_host_namespace_sharing = true
        allow_host_network           = true
        read_only_root_file_system   = true

        allowed_host_port_range {
          min = 3000
          max = 5000
        }

        allowed_volumes = [
          "configMap",
          "nfs",
          "vsphereVolume"
        ]

        run_as_user {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        run_as_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        supplemental_groups {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        fs_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        linux_capabilities {
          allowed_capabilities = [
            "CHOWN",
            "IPC_LOCK"
          ]
          required_drop_capabilities = [
            "SYS_TIME"
          ]
        }

        allowed_host_paths {
          path_prefix = "p1"
          read_only   = true
        }
        allowed_host_paths {
          path_prefix = "p2"
          read_only   = false
        }
        allowed_host_paths {
          path_prefix = "p3"
          read_only   = true
        }

        allowed_se_linux_options {
          level = "s0"
          role  = "sysadm_r"
          type  = "httpd_sys_content_t"
          user  = "root"
        }

        sysctls {
          forbidden_sysctls = [
            "kernel.msgmax",
            "kernel.sem"
          ]
        }

        seccomp {
          allowed_profiles = [
            "Localhost"
          ]
          allowed_localhost_files = [
            "profiles/audit.json",
            "profiles/violation.json"
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


## Organization scoped Strict Security Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control security policy with a strict input recipe.
This policy is applied to an organization with the strict configuration option and is inherited by the cluster groups and clusters.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_security_policy" "organization_scoped_strict_security_policy" {
  name = "tf-sp-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      strict {
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
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the security policy
- `scope` (Block List, Min: 1, Max: 1) Scope for the security policy, having one of the valid scopes: cluster, cluster_group, organization. (see [below for nested schema](#nestedblock--scope))
- `spec` (Block List, Min: 1, Max: 1) Spec for the security policy (see [below for nested schema](#nestedblock--spec))

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

- `input` (Block List, Min: 1, Max: 1) Input for the security policy, having one of the valid recipes: baseline, custom or strict. (see [below for nested schema](#nestedblock--spec--input))

Optional:

- `namespace_selector` (Block List, Max: 1) Label based Namespace Selector for the policy (see [below for nested schema](#nestedblock--spec--namespace_selector))

<a id="nestedblock--spec--input"></a>
### Nested Schema for `spec.input`

Optional:

- `baseline` (Block List, Max: 1) The input schema for security policy baseline recipe version v1 (see [below for nested schema](#nestedblock--spec--input--baseline))
- `custom` (Block List, Max: 1) The input schema for security policy custom recipe version v1 (see [below for nested schema](#nestedblock--spec--input--custom))
- `strict` (Block List, Max: 1) The input schema for security policy strict recipe version v1 (see [below for nested schema](#nestedblock--spec--input--strict))

<a id="nestedblock--spec--input--baseline"></a>
### Nested Schema for `spec.input.baseline`

Optional:

- `audit` (Boolean) Audit (dry-run)
- `disable_native_psp` (Boolean) Disable native pod security policy


<a id="nestedblock--spec--input--custom"></a>
### Nested Schema for `spec.input.custom`

Optional:

- `allow_host_namespace_sharing` (Boolean) Allow host namespace sharing
- `allow_host_network` (Boolean) Allow host network
- `allow_privilege_escalation` (Boolean) Allow privilege escalation
- `allow_privileged_containers` (Boolean) Allow privileged containers
- `allowed_host_paths` (Block List) Allowed host paths (see [below for nested schema](#nestedblock--spec--input--custom--allowed_host_paths))
- `allowed_host_port_range` (Block List, Max: 1) Allowed host port range (see [below for nested schema](#nestedblock--spec--input--custom--allowed_host_port_range))
- `allowed_se_linux_options` (Block List) Allowed selinux options (see [below for nested schema](#nestedblock--spec--input--custom--allowed_se_linux_options))
- `allowed_volumes` (List of String) Allowed volumes
- `audit` (Boolean) Audit (dry-run)
- `disable_native_psp` (Boolean) Disable native pod security policy
- `fs_group` (Block List, Max: 1) fsGroup (see [below for nested schema](#nestedblock--spec--input--custom--fs_group))
- `linux_capabilities` (Block List, Max: 1) Linux capabilities (see [below for nested schema](#nestedblock--spec--input--custom--linux_capabilities))
- `read_only_root_file_system` (Boolean) Read only root file system
- `run_as_group` (Block List, Max: 1) Run as group (see [below for nested schema](#nestedblock--spec--input--custom--run_as_group))
- `run_as_user` (Block List, Max: 1) Run as user (see [below for nested schema](#nestedblock--spec--input--custom--run_as_user))
- `seccomp` (Block List, Max: 1) Seccomp (see [below for nested schema](#nestedblock--spec--input--custom--seccomp))
- `supplemental_groups` (Block List, Max: 1) supplemental groups (see [below for nested schema](#nestedblock--spec--input--custom--supplemental_groups))
- `sysctls` (Block List, Max: 1) Sysctls (see [below for nested schema](#nestedblock--spec--input--custom--sysctls))

<a id="nestedblock--spec--input--custom--allowed_host_paths"></a>
### Nested Schema for `spec.input.custom.allowed_host_paths`

Optional:

- `path_prefix` (String) Path prefix
- `read_only` (Boolean) Read only flag


<a id="nestedblock--spec--input--custom--allowed_host_port_range"></a>
### Nested Schema for `spec.input.custom.allowed_host_port_range`

Optional:

- `max` (Number) Maximum allowed port
- `min` (Number) Minimum allowed port


<a id="nestedblock--spec--input--custom--allowed_se_linux_options"></a>
### Nested Schema for `spec.input.custom.allowed_se_linux_options`

Optional:

- `level` (String) SELinux level
- `role` (String) SELinux role
- `type` (String) SELinux type
- `user` (String) SELinux user


<a id="nestedblock--spec--input--custom--fs_group"></a>
### Nested Schema for `spec.input.custom.fs_group`

Optional:

- `ranges` (Block List) Allowed group id ranges (see [below for nested schema](#nestedblock--spec--input--custom--fs_group--ranges))
- `rule` (String) Rule

<a id="nestedblock--spec--input--custom--fs_group--ranges"></a>
### Nested Schema for `spec.input.custom.fs_group.ranges`

Optional:

- `max` (Number) Maximum group ID
- `min` (Number) Minimum group ID



<a id="nestedblock--spec--input--custom--linux_capabilities"></a>
### Nested Schema for `spec.input.custom.linux_capabilities`

Optional:

- `allowed_capabilities` (List of String) Allowed capabilities
- `required_drop_capabilities` (List of String) Required drop capabilities


<a id="nestedblock--spec--input--custom--run_as_group"></a>
### Nested Schema for `spec.input.custom.run_as_group`

Optional:

- `ranges` (Block List) Allowed group id ranges (see [below for nested schema](#nestedblock--spec--input--custom--run_as_group--ranges))
- `rule` (String) Rule

<a id="nestedblock--spec--input--custom--run_as_group--ranges"></a>
### Nested Schema for `spec.input.custom.run_as_group.ranges`

Optional:

- `max` (Number) Maximum group ID
- `min` (Number) Minimum group ID



<a id="nestedblock--spec--input--custom--run_as_user"></a>
### Nested Schema for `spec.input.custom.run_as_user`

Optional:

- `ranges` (Block List) Allowed user id ranges (see [below for nested schema](#nestedblock--spec--input--custom--run_as_user--ranges))
- `rule` (String) Rule

<a id="nestedblock--spec--input--custom--run_as_user--ranges"></a>
### Nested Schema for `spec.input.custom.run_as_user.ranges`

Optional:

- `max` (Number) Maximum user ID
- `min` (Number) Minimum user ID



<a id="nestedblock--spec--input--custom--seccomp"></a>
### Nested Schema for `spec.input.custom.seccomp`

Optional:

- `allowed_localhost_files` (List of String) Allowed local host files
- `allowed_profiles` (List of String) Allowed profiles


<a id="nestedblock--spec--input--custom--supplemental_groups"></a>
### Nested Schema for `spec.input.custom.supplemental_groups`

Optional:

- `ranges` (Block List) Allowed group id ranges (see [below for nested schema](#nestedblock--spec--input--custom--supplemental_groups--ranges))
- `rule` (String) Rule

<a id="nestedblock--spec--input--custom--supplemental_groups--ranges"></a>
### Nested Schema for `spec.input.custom.supplemental_groups.ranges`

Optional:

- `max` (Number) Maximum group ID
- `min` (Number) Minimum group ID



<a id="nestedblock--spec--input--custom--sysctls"></a>
### Nested Schema for `spec.input.custom.sysctls`

Optional:

- `forbidden_sysctls` (List of String) Forbidden sysctls



<a id="nestedblock--spec--input--strict"></a>
### Nested Schema for `spec.input.strict`

Optional:

- `audit` (Boolean) Audit (dry-run)
- `disable_native_psp` (Boolean) Disable native pod security policy



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
