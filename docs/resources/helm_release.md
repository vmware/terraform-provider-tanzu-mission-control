---
Title: "Helm Release Resource"
Description: |-
    Creating the Helm Release resource.
---

# Helm Release

The `tanzu-mission-control_helm_release` resource allows you to install, update, get and delete helm chart to a particular scope through Tanzu Mission Control.

Before creating helm charts user needs to enable the helm service on a particular scope cluster or cluster group and for enable user can use `tanzu-mission-control_helm_feature` resource.
The`feature_ref` field of `tanzu-mission-control_helm_release` when specified, ensures clean up of this Terraform resource from the state file by creating a dependency on the Helm feature when the Helm feature is disabled.
To add a helm charts, you must be associated with the cluster.admin or clustergroup.admin role.

## Helm Release Scope

In the Tanzu Mission Control resource hierarchy, there are two levels at which you can specify Helm Feature resources:
- **object groups** - `cluster_group` block under `scope` sub-resource
- **Kubernetes objects** - `cluster` block under `scope` sub-resource

**Note:**
The scope parameter is mandatory in the schema and the user needs to add one of the defined scopes to the script for the provider to function.
Only one scope per resource is allowed.

### Install a Helm Chart from a Git Repository

The Helm service must already be enabled to be able to install Helm releases on a cluster or cluster group.
[helm-release]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-F7F4EFA4-F681-42BC-AFDC-874C43D39CD4.html

## Cluster group scoped Helm Release using Git Repository

### Example Usage

```terraform
# Create Tanzu Mission Control cluster group scope helm release with attached set as default value.
resource "tanzu-mission-control_helm_release" "create_cg_helm_release" {
  name = "test-helm-release-name" # Required

  namespace = "test-namespace-name" # Required

  scope {
    cluster_group {
      name = "default" # Required
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.create_cg_helm_feature.scope[0].cluster_group[0].name

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    chart_ref {
      git_repository {
        repository_name      = "testgitrepo"
        repository_namespace = "test-gitrepo-namespace"
        chart_path           = "chartpath"
      }
    }

    inline_config = "<inline-config-file-path>"

    target_namespace = "testtargetnamespacename"

    interval = "10m" # Default: 5m
  }
}
```

## Cluster scoped Helm Release using Git Repository

### Example Usage

```terraform
# Create Tanzu Mission Control cluster scope helm release with attached set as default value.
resource "tanzu-mission-control_helm_release" "create_cl_helm_release_gitrepo_type" {
  name = "test-helm-release-name" # Required

  namespace = "test-namespace-name" # Required

  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.create_cg_helm_feature.scope[0].cluster[0].name

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    chart_ref {
      git_repository {
        repository_name      = "testgitrepo"
        repository_namespace = "test-gitrepo-namespace"
        chart_path           = "chartpath"
      }
    }

    inline_config = "<inline-config-file-path>"

    target_namespace = "testtargetnamespacename"

    interval = "10m" # Default: 5m
  }
}
```

### Install a Helm Chart from a Helm Repository

The Helm service must already be enabled to be able to install Helm releases on a cluster.
[helm-release]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-2602A6A3-1FDA-4270-A76F-047FBD039ADF.html

## Cluster scoped Helm Release using Helm Repository

### Example Usage

```terraform
# Create Tanzu Mission Control cluster scope helm release with attached set as default value.
resource "tanzu-mission-control_helm_release" "create_cl_helm_release_helm_type" {
  name = "test-helm-release-name" # Required

  namespace = "test-namespace-name" # Required

  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.create_cg_helm_feature.scope[0].cluster[0].name

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    chart_ref {
      helm_repository {
        repository_name      = "testgitrepo"
        repository_namespace = "test-helm-namespace"
        chart_name           = "chart-name"
        version              = "test-version"
      }
    }

    inline_config = "<inline-config-file-path>"

    target_namespace = "testtargetnamespacename"

    interval = "10m" # Default: 5m
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the Repository.
- `namespace_name` (String) Name of Namespace.
- `scope` (Block List, Min: 1, Max: 1) Scope for the Helm release, having one of the valid scopes: cluster, cluster_group. (see [below for nested schema](#nestedblock--scope))
- `spec` (Block List, Min: 1, Max: 1) Spec for the Repository. (see [below for nested schema](#nestedblock--spec))

### Optional

- `feature_ref` (String) when specified, ensures clean up of this Terraform resource from the state file by creating a dependency on the Helm feature when the Helm feature is disabled
- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))

### Read-Only

- `id` (String) The ID of this resource.
- `status` (List of Object) status for helm release. (see [below for nested schema](#nestedatt--status))

<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `cluster` (Block List, Max: 1) The schema for cluster full name (see [below for nested schema](#nestedblock--scope--cluster))
- `cluster_group` (Block List, Max: 1) The schema for cluster group full name (see [below for nested schema](#nestedblock--scope--cluster_group))

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



<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `chart_ref` (Block List, Min: 1, Max: 1) Reference to the chart which will be installed. (see [below for nested schema](#nestedblock--spec--chart_ref))

Optional:

- `inline_config` (String) File to read inline values from (in yaml format).User need to specify the file path for inline config
- `interval` (String) Interval at which to reconcile the Helm release. This is the interval at which Tanzu Mission Control will attempt to reconcile changes in the helm release to the cluster. A sync interval of 0 would result in no future syncs. If no value is entered, a default interval of 5 minutes will be applied as `5m`.
- `target_namespace` (String) TargetNamespace sets or overrides the namespaces of resources yaml while applying on cluster.

<a id="nestedblock--spec--chart_ref"></a>
### Nested Schema for `spec.chart_ref`

Optional:

- `git_repository` (Block List, Max: 1) Git repository type spec. (see [below for nested schema](#nestedblock--spec--chart_ref--git_repository))
- `helm_repository` (Block List, Max: 1) Helm repository type Spec. (see [below for nested schema](#nestedblock--spec--chart_ref--helm_repository))

<a id="nestedblock--spec--chart_ref--git_repository"></a>
### Nested Schema for `spec.chart_ref.git_repository`

Required:

- `chart_path` (String) Path of the chart in the git repository.
- `repository_name` (String) Name of the Git repository.
- `repository_namespace` (String) Namespace Name for the Git repository.


<a id="nestedblock--spec--chart_ref--helm_repository"></a>
### Nested Schema for `spec.chart_ref.helm_repository`

Required:

- `chart_name` (String) Name of the chart in the helm repository.
- `repository_name` (String) Name of the Helm repository.
- `repository_namespace` (String) Namespace Name for the Helm repository.
- `version` (String) Chart version, applicable for helm repository type.




<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedatt--status"></a>
### Nested Schema for `status`

Read-Only:

- `generated_resources` (List of Object) Kuberenetes RBAC resources and service account created on the cluster by TMC for helm release. (see [below for nested schema](#nestedobjatt--status--generated_resources))
- `phase` (String) Phase of the Cluster Group helm release application on member Clusters.

<a id="nestedobjatt--status--generated_resources"></a>
### Nested Schema for `status.generated_resources`

Read-Only:

- `cluster_role_name` (String) Name of the cluster role used for helm release.
- `role_binding_name` (String) Name of the role binding used for helm release.
- `service_account_name` (String) Name of the service account used for helm release.