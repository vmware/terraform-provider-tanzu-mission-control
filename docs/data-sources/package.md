---
Title: "Package Data Source"
Description: |-
    Read the Package from TMC.
---

# Package Repository

This resource allows you to read Package from a cluster through Tanzu Mission Control.

The Available tab on the Catalog page in the Tanzu Mission Control console shows the packages that are available to be installed, including those that are in the Tanzu Standard package repository and other repositories that you have associated with a cluster.

[package]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-4B45987F-D5A0-4283-8B4E-139F38DCBFD9.html


## Cluster scoped Package

### Example Usage

```terraform
# Read Tanzu Mission Control package : fetch cluster package details
data "tanzu-mission-control_package" "get_cluster_package" {
  name           = "test-package-version" # Required

  metadata_name  = "package-metadata-name" # Required

  scope {
    cluster {
      cluster_name            = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata_name` (String) Metadata name of package.
- `name` (String) Name of the package. It represents version of the Package metadata
- `scope` (Block List, Min: 1, Max: 1) Scope for the data source, having one of the valid scopes: cluster. (see [below for nested schema](#nestedblock--scope))

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))

### Read-Only

- `id` (String) The ID of this resource.
- `namespace_name` (String) Namespae of package.
- `spec` (List of Object) Spec for the Repository. (see [below for nested schema](#nestedatt--spec))

<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `cluster` (Block List, Max: 1) The schema for cluster iam policy full name (see [below for nested schema](#nestedblock--scope--cluster))

<a id="nestedblock--scope--cluster"></a>
### Nested Schema for `scope.cluster`

Required:

- `name` (String) Name of this cluster

Optional:

- `management_cluster_name` (String) Name of the management cluster
- `provisioner_name` (String) Provisioner of the cluster



<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Read-Only:

- `capacity_requirements_description` (String)
- `licenses` (Set of String)
- `release_notes` (String)
- `released_at` (String)
- `repository_name` (String)
- `values_schema` (List of Object) (see [below for nested schema](#nestedobjatt--spec--values_schema))

<a id="nestedobjatt--spec--values_schema"></a>
### Nested Schema for `spec.values_schema`

Read-Only:

- `template` (List of Object) (see [below for nested schema](#nestedobjatt--spec--values_schema--template))

<a id="nestedobjatt--spec--values_schema--template"></a>
### Nested Schema for `spec.values_schema.template`

Read-Only:

- `raw` (String)