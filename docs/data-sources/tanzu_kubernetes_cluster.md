---
Title: "Tanzu Kubernetes Cluster (Class-based Cluster) Data Source"
Description: |-
   Reading a unified Tanzu Kubernetes Grid cluster.
---

# Tanzu Kubernetes Cluster (Class-based Cluster) Data Source

This data source "tanzu-mission-control_tanzu_kubernetes_cluster" enables users get the details for a TMC managed Tanzu Kubernetes Grid cluster for both Tanzu Kubernetes Grid Vsphere 2.x & Tanzu Kubernetes Grid Service 2.x.


## Example Usage

```terraform
data "tanzu-mission-control_tanzu_kubernetes_cluster" "read_tanzu_cluster" {
  name                    = "tanzu-cluster"
  management_cluster_name = "tanzu-mgmt-cluster"
  provisioner_name        = "tanzu-provisioner"
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `management_cluster_name` (String) Management cluster name
- `name` (String) Cluster name
- `provisioner_name` (String) Cluster provisioner name

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- `spec` (Block List, Max: 1) Spec of tanzu kubernetes cluster. (see [below for nested schema](#nestedblock--spec))
- `timeout_policy` (Block List, Max: 1) Timeout policy for Tanzu Kubernetes cluster. (see [below for nested schema](#nestedblock--timeout_policy))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `topology` (Block List, Min: 1, Max: 1) The cluster topology. (see [below for nested schema](#nestedblock--spec--topology))

Optional:

- `cluster_group_name` (String) Name of the cluster group to which this cluster belongs.
- `image_registry` (String) Name of the image registry configuration to use.
- `proxy_name` (String) Name of the proxy configuration to use.

Read-Only:

- `kubeconfig` (String) Cluster's kubeconfig.
- `tmc_managed` (Boolean) TMC-managed flag indicates if the cluster is managed by tmc.
(Default: False)

<a id="nestedblock--spec--topology"></a>
### Nested Schema for `spec.topology`

Required:

- `cluster_variables` (String) Variables configuration for the cluster.
- `control_plane` (Block List, Min: 1, Max: 1) Control plane specific configuration. (see [below for nested schema](#nestedblock--spec--topology--control_plane))
- `nodepool` (Block List, Min: 1) (Repeatable Block) Node pool definition for the cluster. (see [below for nested schema](#nestedblock--spec--topology--nodepool))
- `version` (String) Kubernetes version of the cluster.

Optional:

- `cluster_class` (String) The name of the cluster class for the cluster.
- `core_addon` (Block List) (Repeatable Block) The core addons. (see [below for nested schema](#nestedblock--spec--topology--core_addon))
- `network` (Block List, Max: 1) Network specific configuration. (see [below for nested schema](#nestedblock--spec--topology--network))

<a id="nestedblock--spec--topology--control_plane"></a>
### Nested Schema for `spec.topology.control_plane`

Required:

- `replicas` (Number) Number of replicas

Optional:

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--spec--topology--control_plane--meta))
- `os_image` (Block List, Max: 1) OS image block (see [below for nested schema](#nestedblock--spec--topology--control_plane--os_image))
- `overrides` (String) Overrides can be used to override cluster level variables.

<a id="nestedblock--spec--topology--control_plane--meta"></a>
### Nested Schema for `spec.topology.control_plane.meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedblock--spec--topology--control_plane--os_image"></a>
### Nested Schema for `spec.topology.control_plane.os_image`

Required:

- `arch` (String) The architecture of the OS image.
- `name` (String) The name of the OS image.
- `version` (String) The version of the OS image.



<a id="nestedblock--spec--topology--nodepool"></a>
### Nested Schema for `spec.topology.nodepool`

Required:

- `name` (String) Name of the node pool.
- `spec` (Block List, Min: 1) Spec for the node pool. (see [below for nested schema](#nestedblock--spec--topology--nodepool--spec))

Optional:

- `description` (String) Description of the node pool.

<a id="nestedblock--spec--topology--nodepool--spec"></a>
### Nested Schema for `spec.topology.nodepool.spec`

Required:

- `replicas` (Number) Number of replicas
- `worker_class` (String) The name of the machine deployment class used to create the node pool.

Optional:

- `failure_domain` (String) The failure domain the machines will be created in.
- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--spec--topology--nodepool--spec--meta))
- `os_image` (Block List, Max: 1) OS image block (see [below for nested schema](#nestedblock--spec--topology--nodepool--spec--os_image))
- `overrides` (String) Overrides can be used to override cluster level variables.

<a id="nestedblock--spec--topology--nodepool--spec--meta"></a>
### Nested Schema for `spec.topology.nodepool.spec.meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `labels` (Map of String) Labels for the resource


<a id="nestedblock--spec--topology--nodepool--spec--os_image"></a>
### Nested Schema for `spec.topology.nodepool.spec.os_image`

Required:

- `arch` (String) The architecture of the OS image.
- `name` (String) The name of the OS image.
- `version` (String) The version of the OS image.




<a id="nestedblock--spec--topology--core_addon"></a>
### Nested Schema for `spec.topology.core_addon`

Required:

- `provider` (String) Provider of core add on
- `type` (String) Type of core add on


<a id="nestedblock--spec--topology--network"></a>
### Nested Schema for `spec.topology.network`

Optional:

- `pod_cidr_blocks` (List of String) Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16.
- `service_cidr_blocks` (List of String) Service CIDR for kubernetes services defaults to 10.96.0.0/12.
- `service_domain` (String) Domain name for services.




<a id="nestedblock--timeout_policy"></a>
### Nested Schema for `timeout_policy`

Optional:

- `fail_on_timeout` (Boolean) Fail on timeout if timeout is reached and cluster is not ready. (Default = true)
- `timeout` (Number) Timeout in minutes for tanzu kubernetes creation process. A value of 0 means that no timeout is set. (Default: 60)
- `wait_for_kubeconfig` (Boolean) Wait for kubeconfig. (Default = true)
