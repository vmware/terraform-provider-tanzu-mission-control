---
Title: "Image Registry Policy Resource"
Description: |-
    Creating the Tanzu Kubernetes image registry policy resource.
---

# Managing Access to Image Registries

Define the registries from which images can be pulled for deployment in your managed namespaces.

Using VMware Tanzu Mission Control, you can make the deployments to namespaces in your clusters more secure by restricting the image registries from which images can be pulled,
as well as the images that can be pulled from a registry.
By default, Tanzu Mission Control does not impose any such restriction,
and allows you to manage image registry restrictions at the organizational level and at the workspace level.

Each namespace and workspace can be protected by an image registry policy that defines the registries from which an image can be pulled,
and these policies are inherited down through the organizational hierarchy.
For more information, see [Policy-Driven Cluster Management][policy-driven-cluster-management] in VMware Tanzu Mission Control Concepts.

[policy-driven-cluster-management]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-847414C9-EF54-44E5-BA62-C4895160CE1D.html

## Input Recipe

In the Tanzu Mission Control image policy resource, there are four system defined types of image policy recipes that you can use:
- **allowed-name-tag** - The Name-Tag allowlist recipe allows you to create rules using an image name or tag name or both.
- **block-latest-tag** - The Block latest tag recipe prevents the use of images that are tagged latest.
- **custom** - The Custom recipe allows you to create rules using multiple factors.
- **require-digest** - The Require Digest recipe prevents the use of images that do not have a digest.

## Policy Scope and Inheritance

In the Tanzu Mission Control resource hierarchy, there are two levels at which you can specify image policy resource:
- **organization** - `organization` block under `scope` sub-resource
- **object groups** - `workspace` block under `scope` sub-resource

In addition to the direct policy defined for a given object, each object has inherited policies described in the parent objects. For example, a namespace has inherited policies from the workspace and organization to which it is linked.

**Note:**
The scope parameter is mandatory in the schema and the user needs to add one of the defined scopes to the script for the provider to function.
Only one scope per resource is allowed.

## Workspace scoped Allowed-name-tag Image Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control image policy with allowed-name-tag input recipe.
This policy is applied to a workspace with the allowed-name-tag configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "workspace_scoped_allowed-name-tag_image_policy" {
  name = "tf-image-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      allowed_name_tag {
        audit = true
        rules {
          imagename = "bar"
          tag {
            negate = true
            value  = "test"
          }
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

## Workspace scoped Block-latest-tag Image Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control image policy with block-latest-tag input recipe.
This policy is applied to a workspace with the block-latest-tag configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "workspace_scoped_block-latest-tag_image_policy" {
  name = "tf-image-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      block_latest_tag {
        audit = false
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

## Workspace scoped Custom Image Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control image policy with custom input recipe.
This policy is applied to a workspace with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "workspace_scoped_custom_image_policy" {
  name = "tf-image-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      custom {
        audit = true
        rules {
          hostname      = "foo"
          imagename     = "bar"
          port          = "80"
          requiredigest = false
          tag {
            negate = false
            value  = "test"
          }
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

## Workspace scoped Require-digest Image Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control image policy with require-digest input recipe.
This policy is applied to a workspace with the require-digest configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "workspace_scoped_require-digest_image_policy" {
  name = "tf-image-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      require_digest {
        audit = false
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

## Organization scoped Allowed-name-tag Image Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control image policy with allowed-name-tag input recipe.
This policy is applied to a organization with the allowed-name-tag configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "organization_scoped_allowed-name-tag_image_policy" {
  name = "tf-image-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      allowed_name_tag {
        audit = true
        rules {
          imagename = "bar"
          tag {
            negate = true
            value  = "test"
          }
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

## Organization scoped Block-latest-tag Image Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control image policy with block-latest-tag input recipe.
This policy is applied to a organization with the block-latest-tag configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "organization_scoped_block-latest-tag_image_policy" {
  name = "tf-image-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      block_latest_tag {
        audit = false
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

## Organization scoped Custom Image Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control image policy with custom input recipe.
This policy is applied to a organization with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "organization_scoped_custom_image_policy" {
  name = "tf-image-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom {
        audit = true
        rules {
          hostname      = "foo"
          imagename     = "bar"
          port          = "80"
          requiredigest = false
          tag {
            negate = false
            value  = "test"
          }
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

## Organization scoped Require-digest Image Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control image policy with require-digest input recipe.
This policy is applied to a organization with the require-digest configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "organization_scoped_require-digest_image_policy" {
  name = "tf-image-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      require_digest {
        audit = false
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

- `name` (String) Name of the image policy
- `scope` (Block List, Min: 1, Max: 1) Scope for the image policy, having one of the valid scopes: organization, workspace. (see [below for nested schema](#nestedblock--scope))
- `spec` (Block List, Min: 1, Max: 1) Spec for the image policy (see [below for nested schema](#nestedblock--spec))

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `organization` (Block List, Max: 1) The schema for organization policy full name (see [below for nested schema](#nestedblock--scope--organization))
- `workspace` (Block List, Max: 1) The schema for workspace policy full name (see [below for nested schema](#nestedblock--scope--workspace))

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

- `input` (Block List, Min: 1, Max: 1) Input for the image policy, having one of the valid recipes: allowed-name-tag, custom, block-latest-tag or require-digest. (see [below for nested schema](#nestedblock--spec--input))

Optional:

- `namespace_selector` (Block List, Max: 1) Label based Namespace Selector for the policy (see [below for nested schema](#nestedblock--spec--namespace_selector))

<a id="nestedblock--spec--input"></a>
### Nested Schema for `spec.input`

Optional:

- `allowed_name_tag` (Block List, Max: 1) The input schema for image policy allowed-name-tag recipe version v1 (see [below for nested schema](#nestedblock--spec--input--allowed_name_tag))
- `block_latest_tag` (Block List, Max: 1) The input schema for image policy block-latest-tag recipe version v1 (see [below for nested schema](#nestedblock--spec--input--block_latest_tag))
- `custom` (Block List, Max: 1) The input schema for image policy custom recipe version v1 (see [below for nested schema](#nestedblock--spec--input--custom))
- `require_digest` (Block List, Max: 1) The input schema for image policy require-digest recipe version v1 (see [below for nested schema](#nestedblock--spec--input--require_digest))

<a id="nestedblock--spec--input--allowed_name_tag"></a>
### Nested Schema for `spec.input.allowed_name_tag`

Required:

- `rules` (Block List, Min: 1) It specifies a list of rules that defines allowed image patterns. (see [below for nested schema](#nestedblock--spec--input--allowed_name_tag--rules))

Optional:

- `audit` (Boolean) Audit (dry-run). Violations will be logged but not denied.

<a id="nestedblock--spec--input--allowed_name_tag--rules"></a>
### Nested Schema for `spec.input.allowed_name_tag.rules`

Optional:

- `imagename` (String) Allowed image names, wildcards are supported(for example: fooservice/*). Empty field is equivalent to *.
- `tag` (Block List, Max: 1) Allowed image tag, wildcards are supported (for example: v1.*). No validation is performed on tag if the field is empty. (see [below for nested schema](#nestedblock--spec--input--allowed_name_tag--rules--tag))

<a id="nestedblock--spec--input--allowed_name_tag--rules--tag"></a>
### Nested Schema for `spec.input.allowed_name_tag.rules.tag`

Optional:

- `negate` (Boolean) The negate flag used to exclude certain tag patterns.
- `value` (String) The value (support wildcard) is used to validate against the tag of the image.




<a id="nestedblock--spec--input--block_latest_tag"></a>
### Nested Schema for `spec.input.block_latest_tag`

Optional:

- `audit` (Boolean) Audit (dry-run). Violations will be logged but not denied.


<a id="nestedblock--spec--input--custom"></a>
### Nested Schema for `spec.input.custom`

Required:

- `rules` (Block List, Min: 1) It specifies a list of rules that defines allowed image patterns. (see [below for nested schema](#nestedblock--spec--input--custom--rules))

Optional:

- `audit` (Boolean) Audit (dry-run). Violations will be logged but not denied.

<a id="nestedblock--spec--input--custom--rules"></a>
### Nested Schema for `spec.input.custom.rules`

Optional:

- `hostname` (String) Allowed image hostnames, wildcards are supported(for example: *.mycompany.com). Empty field is equivalent to *.
- `imagename` (String) Allowed image names, wildcards are supported(for example: fooservice/*). Empty field is equivalent to *.
- `port` (String) Allowed port(if presented) of the image hostname, must associate with valid hostname. Wildcards are supported.
- `requiredigest` (Boolean) The flag used to enforce digest to appear in container images.
- `tag` (Block List, Max: 1) Allowed image tag, wildcards are supported (for example: v1.*). No validation is performed on tag if the field is empty. (see [below for nested schema](#nestedblock--spec--input--custom--rules--tag))

<a id="nestedblock--spec--input--custom--rules--tag"></a>
### Nested Schema for `spec.input.custom.rules.tag`

Optional:

- `negate` (Boolean) The negate flag used to exclude certain tag patterns.
- `value` (String) The value (support wildcard) is used to validate against the tag of the image.




<a id="nestedblock--spec--input--require_digest"></a>
### Nested Schema for `spec.input.require_digest`

Optional:

- `audit` (Boolean) Audit (dry-run). Violations will be logged but not denied.



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
