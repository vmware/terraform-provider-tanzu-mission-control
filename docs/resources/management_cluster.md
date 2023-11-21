---
Title: "Management Cluster Registration Resource"
Description: |-
  Creating management cluster registration resource.
---

# Management Cluster Registration

Manage a cluster registration using this Terraform module.

Registering a management cluster enables you to use VMware Tanzu Mission Control to manage cluster lifecycle on various
infrastructure platforms.

### TKGm flow options
- Registration link is provided after management cluster registration resource has been created.
- When kubeconfig as input is provided then provider will finalize the registration of the resource.

### TKGs flow options
- Registration link is provided after management cluster registration resource has been created.

For creating management cluster registration resource, you must have `managementcluster.admin` permissions in Tanzu Mission Control.
For more information, see [Register a Management Cluster with Tanzu Mission Control.][registration]

[registration]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-EB507AAF-5F4F-400F-9623-BA611233E0BD.html

## Create registration for Tanzu Kubernetes Grid management cluster

Registration output contains registration URL which could be applied according to following
tutorial [Complete the Registration][grid-registration]

[grid-registration]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03.html

### Example Usage

```terraform
resource "tanzu-mission-control_management_cluster" "management_cluster_registration_minimal_tkgm" {
  name = "tf-registration-test" // Required

  spec {
    cluster_group            = "default" // Required
    kubernetes_provider_type = "VMWARE_TANZU_KUBERNETES_GRID" // Required
  }
}
```

## Create registration for Tanzu Kubernetes Grid management cluster with image registry and proxy

Registration output contains registration URL which could be applied according to following
tutorial [Complete the Registration][grid-registration]

[grid-registration]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03.html

### Example Usage

```terraform
resource "tanzu-mission-control_management_cluster" "management_cluster_registration_tkgm" {
  name = "tf-registration-test" // Required

  spec {
    cluster_group                           = "default" // Required
    kubernetes_provider_type                = "VMWARE_TANZU_KUBERNETES_GRID" // Required
    image_registry                          = "image_registry_value"// Optional - only allowed with TKGm - if supplied this should be the name of a pre configured local image registry configured in TMC to pull images from
    managed_workload_cluster_image_registry = "workload_cluster_image_registry_value"// Optional - only allowed with TKGm - only allowed if image_registry is not empty
    management_proxy_name                   = "proxy_name_value"// Optional - name of proxy configuration to use which is already configured in TMC
    managed_workload_cluster_proxy_name     = "workload_cluster_proxy_name_value"// Optional - only allowed if proxy_name is not empty
  }
}
```

## Register Tanzu Kubernetes Grid management cluster with provided kubeconfig file path

### Example Usage

```terraform
# The provider will apply deployment link manifests on to the provided k8s kubeconfig.
resource "tanzu-mission-control_management_cluster" "management_cluster_registration_with_kubeconfig_file_path" {
  name = "tf-registration-test" // Required

  spec {
    cluster_group    = "default" // Required
    kubernetes_provider_type = "VMWARE_TANZU_KUBERNETES_GRID" // Required
  }

  register_management_cluster {
    tkgm_kubeconfig_file = "<kube-config-path>" // Required
    tkgm_description          = "optional description about the kube-config provided" // Optional
  }

  ready_wait_timeout = "15m" // Optional , default value is 15m
}
```

## Register Tanzu Kubernetes Grid management cluster with provided kubeconfig

### Example Usage

```terraform
# The provider will apply deployment link manifests on to the provided k8s kubeconfig.
resource "tanzu-mission-control_management_cluster" "management_cluster_registration_with_kubeconfig_raw_input" {
  name = "tf-registration-test" // Required

  spec {
    cluster_group            = "default" // Required
    kubernetes_provider_type = "VMWARE_TANZU_KUBERNETES_GRID" // Required
  }

  register_management_cluster {
    tkgm_kubeconfig_raw = var.kubeconfig // Required
    tkgm_description    = "optional description about the kube-config provided" // Optional
  }

  ready_wait_timeout = "15m" // Optional , default value is 15m
}

variable "kubeconfig" {
  default = <<EOF
<config>
EOF
}
```

## Register vSphere with Tanzu management cluster

Registration output contains registration URL which could be applied according to following
tutorial [Complete the Registration][vpshere-registration]

[vpshere-registration]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-D85335D2-1430-4662-ABF6-722B7C6276FA.html

### Example Usage

```terraform
resource "tanzu-mission-control_management_cluster" "management_cluster_registration_minimal_tkgs" {
  name = "tf-registration-test" // Required

  spec {
    cluster_group            = "default" // Required
    kubernetes_provider_type = "VMWARE_TANZU_KUBERNETES_GRID_SERVICE" // Required
  }
}
```

## Register vSphere with Tanzu management cluster with image registry and proxy

Registration output contains registration URL which could be applied according to following
tutorial [Complete the Registration][vpshere-registration]

[vpshere-registration]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-D85335D2-1430-4662-ABF6-722B7C6276FA.html

### Example Usage

```terraform
resource "tanzu-mission-control_management_cluster" "management_cluster_registration_tkgs" {
  name = "tf-registration-test" // Required

  spec {
    cluster_group                       = "default" // Required
    kubernetes_provider_type            = "VMWARE_TANZU_KUBERNETES_GRID_SERVICE" // Required
    management_proxy_name               = "proxy_name_value"// Optional - name of proxy configuration to use which is already configured in TMC
    managed_workload_cluster_proxy_name = "workload_cluster_proxy_name_value"// Optional - only allowed if proxy_name is not empty
  }
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- `name` (String) Name of this management cluster
- `spec` (Block List, Min: 1, Max: 1) spec for management cluster registration. (
  see [below for nested schema](#nestedblock--spec))

### Optional

- `org_id` (String) ID of Organization.
- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- `register_management_cluster` (Block List, Max: 1) (
  see [below for nested schema](#nestedblock--register_management_cluster))

### Read-Only

- `id` (String) The ID of this resource.
- `status` (Map of String) Status of the cluster
- `ready_wait_timeout` (String) Wait timeout duration.

<a id="nestedblock--spec"></a>

### Nested Schema for `spec`

Required:

- `cluster_group` (String) Cluster group name to be used by default for workload clusters
- `kubernetes_provider_type` (String) Kubernetes provider type

Optional:

- `image_registry` (String) Image registry which is only allowed for TKGm
- `managed_workload_cluster_image_registry` (String) Managed workload cluster image registry
- `management_proxy_name` (String) Management cluster proxy name
- `managed_workload_cluster_proxy_name` (String) Managed workload cluster proxy name

<a id="nestedblock--meta"></a>

### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource

<a id="nestedblock--register_management_cluster"></a>

### Nested Schema for `register_management_cluster`

Optional:

- `tkgm_kubeconfig_file` (String) Register management cluster KUBECONFIG path for only TKGm
- `tkgm_kubeconfig_raw` (String) Register management cluster KUBECONFIG for only TKGm
- `tkgm_description` (String) Register management cluster description for only TKGm
