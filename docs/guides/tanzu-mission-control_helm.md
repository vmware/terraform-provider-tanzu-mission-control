---
Title: "Helm  Resource"
Description: |-
    Adding Helm resource to a cluster.
---

# Defining Helm

Enable the Helm service when you want to install Helm charts on a cluster or cluster group.
When you enable the Helm service on a cluster or cluster group, you can then create releases in your cluster from Helm charts stored in the Bitnami Helm repository.

# Terraform Resources for Helm

## Helm Feature


The `tanzu-mission-control_helm_feature` resource allows you to enable and disable helm feature for a particular scope through Tanzu Mission Control.

Only after enabling the helm feature User will be able to create the helm release. 

Please refer to [Enable Helm Service on Your Cluster or Cluster Group.][helm-feature]

[helm-feature]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-0927CDC8-A5C1-4FAE-9A7C-8A5D62FDF8D8.html


## Helm Release

The `tanzu-mission-control_helm_release` resource allows you to install, update, get and delete helm chart to a particular scope through Tanzu Mission Control.

Before creating helm charts user needs to enable the helm service on a particular scope cluster or cluster group and for enable user can use `tanzu-mission-control_helm_feature` resource.
The`feature_ref` field of `tanzu-mission-control_helm_release` when specified, ensures clean up of this Terraform resource from the state file by creating a dependency on the Helm feature when the Helm feature is disabled.
To create a Helm release, you must be associated with the cluster.admin or clustergroup.admin role.


# Usecase Example for adding Cluster level helm release from helm repository

```terraform
/*
  NOTE: Creation of cluster level helm release from helm repository
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "cluster_group" {
  name = "demo-cluster-group"
}

# Attach a Tanzu Mission Control cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
  name                    = "demo-cluster" // Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config-path>" // Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = tanzu-mission-control_cluster_group.cluster_group.name // Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

# Create Tanzu Mission Control cluster scope helm feature.
resource "tanzu-mission-control_helm_feature" "cl_helm_feature" {
  scope {
    cluster {
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }
}

# Create Tanzu Mission Control cluster scope helm release.
resource "tanzu-mission-control_helm_release" "cl_helm_release_helm_type" {
  name = "test-helm-release-name" # Required

  namespace_name = "test-namespace-name" # Required

  scope {
    cluster {
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.cg_helm_feature.scope[0].cluster[0].name

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


Please refer to [Install a Helm Chart from a Helm Repository.][helm-release-helmrepo]

[helm-release-helmrepo]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-2602A6A3-1FDA-4270-A76F-047FBD039ADF.html


# Usecase Example for adding Cluster level helm release from git repository

```terraform
/*
  NOTE: Creation of cluster level helm release from git repository
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "cluster_group" {
  name = "demo-cluster-group"
}

# Attach a Tanzu Mission Control cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
  name                    = "demo-cluster" // Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config-path>" // Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = tanzu-mission-control_cluster_group.cluster_group.name // Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

# Create Tanzu Mission Control cluster scope helm feature.
resource "tanzu-mission-control_helm_feature" "cl_helm_feature" {
  scope {
    cluster {
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }
}

# Create Tanzu Mission Control cluster scope helm release.
resource "tanzu-mission-control_helm_release" "cl_helm_release_git_type" {
  name = "test-helm-release-name" # Required

  namespace_name = "test-namespace-name" # Required

  scope {
    cluster {
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.cg_helm_feature.scope[0].cluster[0].name

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


# Usecase Example for adding Cluster level helm release from git repository

```terraform
/*
  NOTE: Creation of cluster group level helm release from git repository
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "cluster_group" {
  name = "demo-cluster-group"
}

# Create Tanzu Mission Control cluster group scope helm feature.
resource "tanzu-mission-control_helm_feature" "cg_helm_feature" {
  scope {
    cluster_group {
      name = tanzu-mission-control_cluster_group.cluster_group.name
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }
}

# Create Tanzu Mission Control cluster group scope helm release.
resource "tanzu-mission-control_helm_release" "cg_helm_release" {
  name = "test-helm-release-name" # Required

  namespace_name = "test-namespace-name" # Required

  scope {
    cluster_group {
      name = tanzu-mission-control_cluster_group.cluster_group.name
    }
  }

  feature_ref = tanzu-mission-control_helm_feature.cg_helm_feature.scope[0].cluster_group[0].name

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



Please refer to [Install a Helm Chart from a Git Repository.][helm-release-gitrepo]

The `tanzu-mission-control_git_repository` resource can be used to add, update, and delete git repository to a particular scope through Tanzu Mission Control.

[helm-release-gitrepo]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-F7F4EFA4-F681-42BC-AFDC-874C43D39CD4.html


# Terraform Data Sources for Helm

## Helm Repository

This `tanzu-mission-control_helm_repository` allows user to read Helm Repository from a cluster through Tanzu Mission Control.

The Available tab on the Catalog page in the Tanzu Mission Control console shows the Helm repositories that are available.


## Cluster scoped Helm Repository

### Example Usage

```terraform
# Read Tanzu Mission Control helm repository : fetch helm repository details
data "tanzu-mission-control_helm_repository" "get_cluster_helm_repo" {
  name = "test-helm-repository_name"

  metadata_name = "test_namespace_name"

  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }
}
```


## Helm Charts

This `tanzu-mission-control_helm_charts` allows you to get list of Helm Charts through Tanzu Mission Control.

The Helm charts tab on the Catalog page in the Tanzu Mission Control console shows the Available Helm charts.


## Organization scoped Helm Charts

### Example Usage

```terraform
# Read Tanzu Mission Control helm charts : fetch helm charts details
data "tanzu-mission-control_helm_charts" "get_cluster_helm_repo" {
  name = "test_name"

  chart_metadata_name = "test_metadata_name"

  repository_name = "test_repository_name"
}
```



Please refer to [Install a Helm Chart from a Helm Repository.][helm-release-helmrepo]

[helm-release-helmrepo]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-2602A6A3-1FDA-4270-A76F-047FBD039ADF.html