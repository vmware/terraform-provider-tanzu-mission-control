---
Title: "Package Install data source"
Description: |-
    Creating the Package Install data source.
---

# Package Install

This data source allows you to read created package install on a cluster through Tanzu Mission Control.

To install an available package on a cluster, you must be associated with the .admin role on that cluster.

Use the Catalog page of the Tanzu Mission Control console to install a package from your repository to your Kubernetes cluster.

The Available tab on the Catalog page in the Tanzu Mission Control console shows the packages that are available to be installed, including those that are in the Tanzu Standard package repository and other repositories that you have associated with a cluster.

[package-install]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-E0168103-7A6F-4C07-8768-19D9B1EB4EFA.html


## Cluster scoped Package install

### Example Usage

```terraform
# Read Tanzu Mission Control package install with attached set as default value.
data "tanzu-mission-control_package_install" "read_package_install" {
  name = "test-pakage-install-name" # Required

  namespace = "test-namespace-name" # Required

  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the package install resource.
- `namespace` (String) Name of Namespace where package install will be created.
- `scope` (Block List, Min: 1, Max: 1) Scope for the package install, having one of the valid scopes: cluster. (see [below for nested schema](#nestedblock--scope))

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))

### Read-Only

- `id` (String) The ID of this resource.
- `spec` (List of Object) spec for package install. (see [below for nested schema](#nestedatt--spec))
- `status` (List of Object) status for package install. (see [below for nested schema](#nestedatt--status))

<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `cluster` (Block List, Max: 1) The schema for cluster full name (see [below for nested schema](#nestedblock--scope--cluster))

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

- `inline_values` (Map of String) Deprecated, use `path_to_inline_values` instead. Inline values to configure the Package Install.
- `package_ref` (List of Object) (see [below for nested schema](#nestedobjatt--spec--package_ref))
- `path_to_inline_values` (String) File to read inline values from (in yaml format). User needs to specify the file path for inline values
- `role_binding_scope` (String)

<a id="nestedobjatt--spec--package_ref"></a>
### Nested Schema for `spec.package_ref`

Read-Only:

- `package_metadata_name` (String)
- `version_selection` (List of Object) (see [below for nested schema](#nestedobjatt--spec--package_ref--version_selection))

<a id="nestedobjatt--spec--package_ref--version_selection"></a>
### Nested Schema for `spec.package_ref.version_selection`

Read-Only:

- `constraints` (String)




<a id="nestedatt--status"></a>
### Nested Schema for `status`

Read-Only:

- `generated_resources` (List of Object) (see [below for nested schema](#nestedobjatt--status--generated_resources))
- `managed` (Boolean)
- `package_install_phase` (String)
- `referred_by` (List of String)
- `resolved_version` (String)

<a id="nestedobjatt--status--generated_resources"></a>
### Nested Schema for `status.generated_resources`

Read-Only:

- `cluster_role_name` (String)
- `role_binding_name` (String)
- `service_account_name` (String)
