---
Title: "Network Policy Resource"
Description: |-
    Creating the Tanzu Kubernetes network policy resource.
---

# Managing Network Communication for Your Clusters

Define how pods communicate using network policies.

Using VMware Tanzu Mission Control, you can create a network policy that defines how pods communicate with each other and other network endpoints, using preconfigured templates called recipes.
By default, Tanzu Mission Control does not impose any such restriction, and allows you to manage network restrictions at the organizational level and at the workspace level.

Tanzu Mission Control implements network policies using Kubernetes native network policies. Each namespace and workspace can be governed by a network policy, and these policies are inherited down through the organizational hierarchy.
Network policies are additive, both inherited and direct network policies are applied and are effective on your namespaces according to Kubernetes rules.

For more information about Kubernetes native network policies, see [Network Policies][network-policies] in the Kuberenetes documentation.
For more information about policy inheritance in Tanzu Mission Control, see [Policy-Driven Cluster Management][policy-driven-cluster-management] in VMware Tanzu Mission Control Concepts.

[network-policies]: https://kubernetes.io/docs/concepts/services-networking/network-policies/
[policy-driven-cluster-management]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-847414C9-EF54-44E5-BA62-C4895160CE1D.html

## Input Recipe

In the Tanzu Mission Control network policy resource, there are currently eight system defined types of network policy recipes that you can use:

- **allow-all**
- **allow-all-to-pods**
- **allow-all-egress**
- **deny-all**
- **deny-all-to-pods**
- **deny-all-egress**
- **custom-egress**
- **custom-ingress**

## Policy Scope and Inheritance

In the Tanzu Mission Control resource hierarchy, there are two levels at which you can specify network policy resource:
- **organization** - `organization` block under `scope` sub-resource
- **object groups** - `workspace` block under `scope` sub-resource

In addition to the direct policy defined for a given object, each object has inherited policies described in the parent objects. For example, a namespace has inherited policies from the workspace and organization to which it is linked.

**Note:**
The scope parameter is mandatory in the schema and the user needs to add one of the defined scopes to the script for the provider to function.
Only one scope per resource is allowed.

## Rules in Network Policies

Some of the network policy recipes allow you to provide a rule that uses a set of criteria to identify the target locations with which to permit or restrict communication, and the port on which they can communicate.
The criteria used in these rules can include the following types:

- IP range (allow and exclude)
- label selector (pods and namespaces)
- port and protocol

You can define multiple criteria of a given type in a single rule, and use these criteria in combination with each other.
The location criteria (IP range and label selector) that you define are specific to the template that you are using.
For the custom-ingress template, you identify sources from which to allow traffic; and for the custom-egress template, you identify destinations to which to allow traffic.

If you do not specify any location criteria, the policy does not restrict traffic by location.
All sources or destinations are allowed. Likewise, if you do not specify any ports, all ports are allowed.

### IP Range Criteria

When you specify a range for allowed IP addresses, traffic is permitted on all IP addresses in that range.
You can also optionally exclude a range of IP addresses within the allowed range.

If you specify multiple IP ranges for a given location, the location must match any one of the criteria.
For example, if you define three allowed IP ranges, traffic is allowed to (or from) locations within any one of the three ranges.

### Label Selector Criteria

If you specify multiple label selectors for a given type, the location must match any one of the criteria to allow traffic.
For example, if you define three pod selectors, traffic is allowed to (or from) pods that have a label matching any one of the three selectors.

If you specify a location using both the pod selector and the namespace selector in a single location definition, then both must be satisfied.

### Port Criteria

The port and protocol fields allow you to specify a port on which to allow traffic, and the protocol that the traffic must use.
You can specify multiple ports, and each one must have a corresponding protocol. The port can be either a numerical or named port.

If you specify multiple ports, the channel must match any one of the criteria to allow traffic.
For example, if you define three ports, traffic is allowed through any one of the three ports.

## Workspace scoped allow-all Network Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control network policy with allow-all input recipe.
This policy is applied to a workspace with the allow-all configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_network_policy" "workspace_scoped_allow-all_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      allow_all {
        from_own_namespace = true
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

## Workspace scoped allow-all-to-pods Network Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control network policy with allow-all-to-pods input recipe.
This policy is applied to a workspace with the allow-all-to-pods configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "workspace_scoped_allow-all-to-pods_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      allow_all_to_pods {
        from_own_namespace = false
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

## Workspace scoped allow-all-egress Network Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control network policy with allow-all-egress input recipe.
This policy is applied to a workspace with the allow-all-egress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "workspace_scoped_allow-all-egress_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      allow_all_egress {}
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

## Workspace scoped deny-all Network Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control network policy with deny-all input recipe.
This policy is applied to a workspace with the deny-all configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "workspace_scoped_deny-all_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      deny_all {}
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

## Workspace scoped deny-all-to-pods Network Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control network policy with deny-all-to-pods input recipe.
This policy is applied to a workspace with the deny-all-to-pods configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "workspace_scoped_deny-all-to-pods_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      deny_all_to_pods {
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

## Workspace scoped deny-all-egress Network Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control network policy with deny-all-egress input recipe.
This policy is applied to a workspace with the deny-all-egress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "workspace_scoped_deny-all-egress_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      deny_all_egress {}
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

## Workspace scoped custom-egress Network Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control network policy with custom-egress input recipe.
This policy is applied to a workspace with the custom-egress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "workspace_scoped_custom-egress_destination-selector_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      custom_egress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_selector {
              namespace_selector = {
                "ns-key-1" = "ns-val-1"
                "ns-key-2" = "ns-val-2"
              }
              pod_selector = {
                "pod-key-1" = "pod-val-1"
                "pod-key-2" = "pod-val-2"
              }
            }
          }
        }
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

resource "tanzu-mission-control_network_policy" "workspace_scoped_custom-egress_destination-ip-block_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      custom_egress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_ip {
              ip_block {
                cidr = "192.168.1.1/24"
                except = [
                  "2001:db9::/64",
                ]
              }
            }
          }
        }
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

## Workspace scoped custom-ingress Network Policy

### Example Usage

```terraform
/*
Workspace scoped Tanzu Mission Control network policy with custom-ingress input recipe.
This policy is applied to a workspace with the custom-ingress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "workspace_scoped_custom-ingress_source-selector_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      custom_ingress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_selector {
              namespace_selector = {
                "ns-key-1" = "ns-val-1"
                "ns-key-2" = "ns-val-2"
              }
              pod_selector = {
                "pod-key-1" = "pod-val-1"
                "pod-key-2" = "pod-val-2"
              }
            }
          }
        }
        to_pod_labels {
          key   = "key-1"
          value = "value-1"
        }
        to_pod_labels {
          key   = "key-2"
          value = "value-2"
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

resource "tanzu-mission-control_network_policy" "workspace_scoped_custom-ingress_source-ip-block_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      custom_ingress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_ip {
              ip_block {
                cidr = "192.168.1.1/24"
                except = [
                  "2001:db9::/64",
                ]
              }
            }
          }
        }
        to_pod_labels {
          key   = "key-1"
          value = "value-1"
        }
        to_pod_labels {
          key   = "key-2"
          value = "value-2"
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


## Organization scoped allow-all Network Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control network policy with allow-all input recipe.
This policy is applied to a organization with the allow-all configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_allow-all_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      allow_all {
        from_own_namespace = false
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

## Organization scoped allow-all-to-pods Network Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control network policy with allow-all-to-pods input recipe.
This policy is applied to a organization with the allow-all-to-pods configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_allow-all-to-pods_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      allow_all_to_pods {
        from_own_namespace = false
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

## Organization scoped allow-all-egress Network Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control network policy with allow-all-egress input recipe.
This policy is applied to a organization with the allow-all-egress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_allow-all-egress_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      allow_all_egress {}
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

## Organization scoped deny-all Network Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control network policy with deny-all input recipe.
This policy is applied to a organization with the deny-all configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_deny-all_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      deny_all {}
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

## Organization scoped deny-all-to-pods Network Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control network policy with deny-all-to-pods input recipe.
This policy is applied to a organization with the deny-all-to-pods configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_deny-all-to-pods_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      deny_all_to_pods {
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

## Organization scoped deny-all-egress Network Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control network policy with deny-all-egress input recipe.
This policy is applied to a organization with the deny-all-egress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_deny-all-egress_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      deny_all_egress {}
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

## Organization scoped custom-egress Network Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control network policy with custom-egress input recipe.
This policy is applied to a organization with the custom-egress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_custom-egress_destination-selector_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom_egress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_selector {
              namespace_selector = {
                "ns-key-1" = "ns-val-1"
                "ns-key-2" = "ns-val-2"
              }
              pod_selector = {
                "pod-key-1" = "pod-val-1"
                "pod-key-2" = "pod-val-2"
              }
            }
          }
        }
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

resource "tanzu-mission-control_network_policy" "organization_scoped_custom-egress_destination-ip-block_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom_egress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_ip {
              ip_block {
                cidr = "192.168.1.1/24"
                except = [
                  "2001:db9::/64",
                ]
              }
            }
          }
        }
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

## Organization scoped custom-ingress Network Policy

### Example Usage

```terraform
/*
Organization scoped Tanzu Mission Control network policy with custom-ingress input recipe.
This policy is applied to a organization with the custom-ingress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_custom-ingress_source-selector_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom_ingress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_selector {
              namespace_selector = {
                "ns-key-1" = "ns-val-1"
                "ns-key-2" = "ns-val-2"
              }
              pod_selector = {
                "pod-key-1" = "pod-val-1"
                "pod-key-2" = "pod-val-2"
              }
            }
          }
        }
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

resource "tanzu-mission-control_network_policy" "organization_scoped_custom-ingress_source-ip-block_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom_ingress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_ip {
              ip_block {
                cidr = "192.168.1.1/24"
                except = [
                  "2001:db9::/64",
                ]
              }
            }
          }
        }
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
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

- `name` (String) Name of the network policy
- `scope` (Block List, Min: 1, Max: 1) Scope for the network policy, having one of the valid scopes: organization, workspace. (see [below for nested schema](#nestedblock--scope))
- `spec` (Block List, Min: 1, Max: 1) Spec for the network policy (see [below for nested schema](#nestedblock--spec))

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

- `input` (Block List, Min: 1, Max: 1) Input for the network policy, having one of the valid recipes: allow-all, allow-all-to-pods, allow-all-egress, deny-all, deny-all-to-pods, deny-all-egress, custom-egress or custom-ingress. (see [below for nested schema](#nestedblock--spec--input))

Optional:

- `namespace_selector` (Block List, Max: 1) Label based Namespace Selector for the policy (see [below for nested schema](#nestedblock--spec--namespace_selector))

<a id="nestedblock--spec--input"></a>
### Nested Schema for `spec.input`

Optional:

- `allow_all` (Block List, Max: 1) The input schema for network policy allow-all recipe version v1 (see [below for nested schema](#nestedblock--spec--input--allow_all))
- `allow_all_egress` (Block List, Max: 1) The input schema for network policy allow-all-egress recipe version v1 (see [below for nested schema](#nestedblock--spec--input--allow_all_egress))
- `allow_all_to_pods` (Block List, Max: 1) The input schema for network policy allow-all-to-pods recipe version v1 (see [below for nested schema](#nestedblock--spec--input--allow_all_to_pods))
- `custom_egress` (Block List, Max: 1) The input schema for network policy custom egress recipe version v1 (see [below for nested schema](#nestedblock--spec--input--custom_egress))
- `custom_ingress` (Block List, Max: 1) The input schema for network policy custom ingress recipe version v1 (see [below for nested schema](#nestedblock--spec--input--custom_ingress))
- `deny_all` (Block List, Max: 1) The input schema for network policy deny-all recipe version v1 (see [below for nested schema](#nestedblock--spec--input--deny_all))
- `deny_all_egress` (Block List, Max: 1) The input schema for network policy deny-all-egress recipe version v1 (see [below for nested schema](#nestedblock--spec--input--deny_all_egress))
- `deny_all_to_pods` (Block List, Max: 1) The input schema for network policy deny-all-to-pods recipe version v1 (see [below for nested schema](#nestedblock--spec--input--deny_all_to_pods))

<a id="nestedblock--spec--input--allow_all"></a>
### Nested Schema for `spec.input.allow_all`

Optional:

- `from_own_namespace` (Boolean) Allow traffic only from own namespace. Allow traffic only from pods in the same namespace as the destination pod.


<a id="nestedblock--spec--input--allow_all_egress"></a>
### Nested Schema for `spec.input.allow_all_egress`


<a id="nestedblock--spec--input--allow_all_to_pods"></a>
### Nested Schema for `spec.input.allow_all_to_pods`

Optional:

- `from_own_namespace` (Boolean) Allow traffic only from own namespace. Allow traffic only from pods in the same namespace as the destination pod.
- `to_pod_labels` (Map of String) Pod Labels on which traffic should be allowed/denied. Use a label selector to identify the pods to which the policy applies.


<a id="nestedblock--spec--input--custom_egress"></a>
### Nested Schema for `spec.input.custom_egress`

Required:

- `rules` (Block List, Min: 1) This specifies list of egress rules to be applied to the selected pods. (see [below for nested schema](#nestedblock--spec--input--custom_egress--rules))

Optional:

- `to_pod_labels` (Map of String) Pod Labels on which traffic should be allowed/denied. Use a label selector to identify the pods to which the policy applies.

<a id="nestedblock--spec--input--custom_egress--rules"></a>
### Nested Schema for `spec.input.custom_egress.rules`

Required:

- `ports` (Block List, Min: 1) List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. Default is this rule matches all ports (traffic not restricted by port). (see [below for nested schema](#nestedblock--spec--input--custom_egress--rules--ports))
- `rule_spec` (Block List, Min: 1) List of destinations for outgoing traffic of pods selected for this rule. Default is the rule matches all destinations (traffic not restricted by destinations). (see [below for nested schema](#nestedblock--spec--input--custom_egress--rules--rule_spec))

<a id="nestedblock--spec--input--custom_egress--rules--ports"></a>
### Nested Schema for `spec.input.custom_egress.rules.ports`

Optional:

- `port` (String) The port on the given protocol. This can either be a numerical or named port on a pod.
- `protocol` (String) The protocol (TCP or UDP) which traffic must match.


<a id="nestedblock--spec--input--custom_egress--rules--rule_spec"></a>
### Nested Schema for `spec.input.custom_egress.rules.rule_spec`

Optional:

- `custom_ip` (Block List) The rule Spec (destination) for IP Block. (see [below for nested schema](#nestedblock--spec--input--custom_egress--rules--rule_spec--custom_ip))
- `custom_selector` (Block List) The rule Spec (destination) for Selectors. (see [below for nested schema](#nestedblock--spec--input--custom_egress--rules--rule_spec--custom_selector))

<a id="nestedblock--spec--input--custom_egress--rules--rule_spec--custom_ip"></a>
### Nested Schema for `spec.input.custom_egress.rules.rule_spec.custom_ip`

Optional:

- `ip_block` (Block List) IPBlock defines policy on a particular IPBlock. If this field is set then neither of the namespaceSelector and PodSelector can be set. (see [below for nested schema](#nestedblock--spec--input--custom_egress--rules--rule_spec--custom_ip--ip_block))

<a id="nestedblock--spec--input--custom_egress--rules--rule_spec--custom_ip--ip_block"></a>
### Nested Schema for `spec.input.custom_egress.rules.rule_spec.custom_ip.ip_block`

Required:

- `cidr` (String) CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64"

Optional:

- `except` (List of String) Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64" Except values will be rejected if they are outside the CIDR range



<a id="nestedblock--spec--input--custom_egress--rules--rule_spec--custom_selector"></a>
### Nested Schema for `spec.input.custom_egress.rules.rule_spec.custom_selector`

Optional:

- `namespace_selector` (Map of String) Use a label selector to identify the namespaces to allow as egress destinations.
- `pod_selector` (Map of String) Use a label selector to identify the pods to allow as egress destinations.





<a id="nestedblock--spec--input--custom_ingress"></a>
### Nested Schema for `spec.input.custom_ingress`

Required:

- `rules` (Block List, Min: 1) This specifies list of ingress rules to be applied to the selected pods. (see [below for nested schema](#nestedblock--spec--input--custom_ingress--rules))

Optional:

- `to_pod_labels` (Map of String) Pod Labels on which traffic should be allowed/denied. Use a label selector to identify the pods to which the policy applies.

<a id="nestedblock--spec--input--custom_ingress--rules"></a>
### Nested Schema for `spec.input.custom_ingress.rules`

Required:

- `ports` (Block List, Min: 1) List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. Default is this rule matches all ports (traffic not restricted by port). (see [below for nested schema](#nestedblock--spec--input--custom_ingress--rules--ports))
- `rule_spec` (Block List, Min: 1) List of sources which should be able to access the pods selected for this rule. Default is the rule matches all sources (traffic not restricted by source). List of items of type V1alpha1CommonPolicySpecNetworkV1CustomIngressRulesRuleSpec0 OR V1alpha1CommonPolicySpecNetworkV1CustomIngressRulesRuleSpec1. (see [below for nested schema](#nestedblock--spec--input--custom_ingress--rules--rule_spec))

<a id="nestedblock--spec--input--custom_ingress--rules--ports"></a>
### Nested Schema for `spec.input.custom_ingress.rules.ports`

Optional:

- `port` (String) The port on the given protocol. This can either be a numerical or named port on a pod.
- `protocol` (String) The protocol (TCP or UDP) which traffic must match.


<a id="nestedblock--spec--input--custom_ingress--rules--rule_spec"></a>
### Nested Schema for `spec.input.custom_ingress.rules.rule_spec`

Optional:

- `custom_ip` (Block List) The rule Spec (source) for IP Block. (see [below for nested schema](#nestedblock--spec--input--custom_ingress--rules--rule_spec--custom_ip))
- `custom_selector` (Block List) The rule Spec (source) for Selectors. (see [below for nested schema](#nestedblock--spec--input--custom_ingress--rules--rule_spec--custom_selector))

<a id="nestedblock--spec--input--custom_ingress--rules--rule_spec--custom_ip"></a>
### Nested Schema for `spec.input.custom_ingress.rules.rule_spec.custom_ip`

Optional:

- `ip_block` (Block List) IPBlock defines policy on a particular IPBlock. If this field is set then neither of the namespaceSelector and PodSelector can be set. (see [below for nested schema](#nestedblock--spec--input--custom_ingress--rules--rule_spec--custom_ip--ip_block))

<a id="nestedblock--spec--input--custom_ingress--rules--rule_spec--custom_ip--ip_block"></a>
### Nested Schema for `spec.input.custom_ingress.rules.rule_spec.custom_ip.ip_block`

Required:

- `cidr` (String) CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64"

Optional:

- `except` (List of String) Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64" Except values will be rejected if they are outside the CIDR range



<a id="nestedblock--spec--input--custom_ingress--rules--rule_spec--custom_selector"></a>
### Nested Schema for `spec.input.custom_ingress.rules.rule_spec.custom_selector`

Optional:

- `namespace_selector` (Map of String) Use a label selector to identify the namespaces to allow as egress destinations.
- `pod_selector` (Map of String) Use a label selector to identify the pods to allow as egress destinations.





<a id="nestedblock--spec--input--deny_all"></a>
### Nested Schema for `spec.input.deny_all`


<a id="nestedblock--spec--input--deny_all_egress"></a>
### Nested Schema for `spec.input.deny_all_egress`


<a id="nestedblock--spec--input--deny_all_to_pods"></a>
### Nested Schema for `spec.input.deny_all_to_pods`

Optional:

- `to_pod_labels` (Map of String) Pod Labels on which traffic should be allowed/denied. Use a label selector to identify the pods to which the policy applies.



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
