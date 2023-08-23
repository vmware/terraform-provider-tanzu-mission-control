---
Title: "IAM Policy Resource"
Description: |-
    Creating the Tanzu Kubernetes IAM policy resource.
---

# IAM Policy

The `tanzu-mission-control_iam_policy` resource allows you to add, update, and delete role bindings to a particular scope for identity and access management through Tanzu Mission Control.

IAM policy (also known as Access Management) allows you to implement role-based access control (RBAC) in Tanzu Mission Control. For more information, see [Access Control.][access-control]

[access-control]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-EB9C6D83-1132-444F-8218-F264E43F25BD.html

## Policy Scope and Inheritance

In the Tanzu Mission Control resource hierarchy, there are three levels and five object types at which you can specify IAM policy resources:
- **organization** - `organization` block under `scope` sub-resource
- **object groups** - `cluster_group` or `workspace` block under `scope` sub-resource
- **Kubernetes objects** - `cluster` or `namespace` block under `scope` sub-resource

In addition to the direct policy defined for a given object, each object has inherited policies described in the parent objects. For example, a cluster has a direct policy and inherited policies from the cluster group and organization to which it is attached.
Similarly, a namespace has a direct policy and inherited policies from the workspace with which it is associated.

**Note:**
The scope parameter is mandatory in the schema and the user needs to add one of the defined scopes to the script for the provider to function.
Only one scope per resource is allowed.

## Managing Access to Your Resources

To use the **Tanzu Mission Control provider** for adding, editing, and removing role bindings, you must define who has access to each resource in your organization using role-based access control.
For more information, see [Managing Access to Resources.][managing-access]

[managing-access]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-CA5A31BC-4D7B-4EDD-A4C8-95BEEC08F7C4.html

## Organization scoped IAM Policy

### Example Usage

```terraform
/*
 Organization scoped Tanzu Mission Control IAM policy.
 This resource is applied on an organization to provision the role bindings on the associated organization.
 The defined scope block can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "organization_scoped_iam_policy" {
  scope {
    organization {
      org_id = "dummy-org-id"
    }
  }

  role_bindings {
    role = "organization.view"
    subjects {
      name = "test-1"
      kind = "USER"
    }
    subjects {
      name = "test-2"
      kind = "GROUP"
    }
  }
  role_bindings {
    role = "organization.edit"
    subjects {
      name = "test-3"
      kind = "USER"
    }
  }
}
```

## Cluster group scoped IAM Policy

### Example Usage

```terraform
/*
 Cluster group scoped Tanzu Mission Control IAM policy.
 This resource is applied on a cluster group to provision the role bindings on the associated cluster group.
 The defined scope block can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "cluster_group_scoped_iam_policy" {
  scope {
    cluster_group {
      name = "default"
    }
  }

  role_bindings {
    role = "clustergroup.admin"
    subjects {
      name = "test"
      kind = "GROUP"
    }
  }
  role_bindings {
    role = "clustergroup.edit"
    subjects {
      name = "test-new"
      kind = "USER"
    }
  }
}
```

## Cluster scoped IAM Policy

### Example Usage

```terraform
/*
 Cluster scoped Tanzu Mission Control IAM policy.
 This resource is applied on a cluster to provision the role bindings on the associated cluster.
 The defined scope block can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "cluster_scoped_iam_policy" {
  scope {
    cluster {
      management_cluster_name = "attached" # Default: attached
      provisioner_name        = "attached" # Default: attached
      name                    = "demo"
    }
  }

  role_bindings {
    role = "cluster.admin"
    subjects {
      name = "test"
      kind = "GROUP"
    }
  }
  role_bindings {
    role = "cluster.edit"
    subjects {
      name = "test-1"
      kind = "USER"
    }
    subjects {
      name = "test-2"
      kind = "USER"
    }
  }
}
```

## Workspace scoped IAM Policy

### Example Usage

```terraform
/*
 Workspace scoped Tanzu Mission Control IAM policy.
 This resource is applied on a workspace to provision the role bindings on the associated workspace.
 The defined scope block can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "workspace_scoped_iam_policy" {
  scope {
    workspace {
      name = "tf-workspace"
    }
  }

  role_bindings {
    role = "workspace.edit"
    subjects {
      name = "test"
      kind = "USER"
    }
  }
}
```

## Workspace scoped IAM Policy using a K8s Service Account

### Example Usage

```terraform
/*
 Workspace scoped Tanzu Mission Control IAM policy using a K8s Service Account
 This resource is applied on a workspace to provision the role bindings on the associated workspace.
 The defined scope block can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "workspace_scoped_iam_policy" {
  scope {
    workspace {
      name = "tf-workspace"
    }
  }
  role_bindings {
    role = "workspace.edit"
    subjects {
      name = "namespace:serviceaccountname"
      kind = "K8S_SERVICEACCOUNT"
    }
  }
}
```

## Namespace scoped IAM Policy

### Example Usage

```terraform
/*
 Namespace scoped Tanzu Mission Control IAM policy.
 This resource is applied on a namespace to provision the role bindings on the associated namespace.
 The defined scope block can be updated to change the access policy's scope.
 */
resource "tanzu-mission-control_iam_policy" "namespace_scoped_iam_policy" {
  scope {
    namespace {
      management_cluster_name = "attached" # Default: attached
      provisioner_name        = "attached" # Default: attached
      cluster_name            = "demo"
      name                    = "tf_namespace"
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

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `role_bindings` (Block List, Min: 1) List of role bindings associated with the policy (see [below for nested schema](#nestedblock--role_bindings))
- `scope` (Block List, Min: 1, Max: 1) Scope of the resource on which the rolebinding has to be added, having one of the valid scopes: organization, cluster_group, cluster, workspace or namespace. (see [below for nested schema](#nestedblock--scope))

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--role_bindings"></a>
### Nested Schema for `role_bindings`

Required:

- `role` (String) Role for this rolebinding: max length for a role is 126 characters.
- `subjects` (Block List, Min: 1) Subject for this rolebinding. (see [below for nested schema](#nestedblock--role_bindings--subjects))

<a id="nestedblock--role_bindings--subjects"></a>
### Nested Schema for `role_bindings.subjects`

Required:

- `kind` (String) Subject type, having one of the subject types: USER or GROUP or K8S_SERVICEACCOUNT
- `name` (String) Subject name: allow max characters for email - 320 characters.



<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `cluster` (Block List, Max: 1) The schema for cluster full name (see [below for nested schema](#nestedblock--scope--cluster))
- `cluster_group` (Block List, Max: 1) The schema for cluster group full name (see [below for nested schema](#nestedblock--scope--cluster_group))
- `namespace` (Block List, Max: 1) The schema for namespace iam policy full name (see [below for nested schema](#nestedblock--scope--namespace))
- `organization` (Block List, Max: 1) The schema for organization iam policy full name (see [below for nested schema](#nestedblock--scope--organization))
- `workspace` (Block List, Max: 1) The schema for workspace iam policy full name (see [below for nested schema](#nestedblock--scope--workspace))

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

- `name` (String) Name of the cluster group


<a id="nestedblock--scope--namespace"></a>
### Nested Schema for `scope.namespace`

Required:

- `cluster_name` (String) Name of Cluster
- `name` (String) Name of the Namespace

Optional:

- `management_cluster_name` (String) Name of ManagementCluster
- `provisioner_name` (String) Name of Provisioner


<a id="nestedblock--scope--organization"></a>
### Nested Schema for `scope.organization`

Required:

- `org_id` (String) ID of the Organization


<a id="nestedblock--scope--workspace"></a>
### Nested Schema for `scope.workspace`

Required:

- `name` (String) Name of the workspace



<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource
