---
Title: "Provisioning of a workload cluster"
Description: |-
    An example of provisioning a Tanzu Kubernetes Grid Service and a Tanzu Kubernetes Grid vSphere workload cluster.
---

# Tanzu Kubernetes Grid Service Workload Cluster

Before creating a Tanzu Kubernetes Grid Service workload cluster in vSphere with Tanzu using this Terraform provider we need the following prerequisites.

**Prerequisites:**

- Register the Tanzu Kubernetes Grid Service management cluster in Tanzu Mission Control.
Note that the Tanzu Kubernetes Grid Service management cluster must be **ready** and **healthy**.
Please refer to [registration of a Supervisor Cluster in vSphere with Tanzu.][supervisor-cluster-registration]

- Create a provisioner under the management cluster or reuse the existing providers under the management cluster. Please refer to [working with vSphere Namespaces on a Supervisor Cluster.][vSphere-namespaces]

- Ensure the CSP token used in initialising the terraform provider has the right set of permissions to create a workload cluster.

Once you have the `management cluster name` and `provisioner name` from Tanzu mission control, we are all set to provision a workload under the chosen management cluster name using the terraform script (example below).

If you need to update the cluster, you update the rule definition and then Terraform apply applies it across all the sites.
If you add or remove a site from the list, Terraform also handles creating or removing the rule on the subsequent `terraform apply`operation.

If you want to manage the already existing workload, Refer to [adding a Workload Cluster into Tanzu Mission Control Management.][add-workload-cluster]

[supervisor-cluster-registration]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03.html#GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03
[vSphere-namespaces]: https://docs.vmware.com/en/VMware-vSphere/7.0/vmware-vsphere-with-tanzu/GUID-1544C9FE-0B23-434E-B823-C59EFC2F7309.html
[add-workload-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-78908829-CB4E-459F-AA81-BEA415EC9A11.html

You can provision a Tanzu Kubernetes Grid Service workload cluster in vSphere with Tanzu using this Terraform provider, as shown in the example below.
This example assumes that you have already registered a Tanzu Kubernetes Grid Service management cluster with Tanzu Mission Control.

```terraform
// Tanzu Mission Control Cluster Type: Tanzu Kubernetes Grid Service workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster : fetch cluster details
data "tanzu-mission-control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Required
  provisioner_name        = "<prov-name>"          // Required
  name                    = "<cluster-name>"       // Required
}

# Create Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkgs_workload" {
  management_cluster_name = "<management-cluster>"
  provisioner_name        = "<prov-name>"
  name                    = "<cluster-name>"

  meta {
    labels = { "key" : "test" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
    tkg_service_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "<pods-cidr-blocks>", // Required
            ]
          }
          services {
            cidr_blocks = [
              "<services-cidr-blocks>", // Required
            ]
          }
        }
      }

      distribution {
        version = "<version>" // Required
      }

      topology {
        control_plane {
          class             = "<class>"        // Required
          storage_class     = "<storage-class" // Required
          high_availability = false            // Default: false
        }
        node_pools {
          spec {
            worker_node_count = "<worker-node-count>" // Required
            cloud_label = {
              "<key>" : "<val>"
            }
            node_label = {
              "<key>" : "<val>"
            }
            tkg_service_vsphere {
              class         = "<class>"         // Required
              storage_class = "<storage-class>" // Required
            }
          }
          info {
            name = "<node-pool-name>" // Required
          }
        }
      }
    }
  }
}
```

# Tanzu Kubernetes Grid vSphere Workload Cluster

Before provisioning a Tanzu Kubernetes Grid Vsphere workload cluster using this Terraform provider we need the following prerequisites.

**Prerequisites:**

- Register the Tanzu Kubernetes Grid management cluster in Tanzu Mission Control.
Please refer to [register Management Cluster with Tanzu Mission Control.][management-cluster-registration]

- Create a provisioner under the management cluster or reuse the existing provisioners under the management cluster. Please refer to [create a Provisioner.][create-provisioner]

- Ensure the CSP token used in initialising the terraform provider has the right set of permissions to create a workload cluster.

Once you have the `management cluster name` and `provisioner name` from Tanzu mission control, we are all set to provision a workload under the chosen management cluster name using the terraform script (example below).

If you need to update the cluster, you update the rule definition and then Terraform apply applies it across all the sites.
If you add or remove a site from the list, Terraform also handles creating or removing the rule on the subsequent `terraform apply`operation.

If you want to manage the already existing workload, refer to [adding a Workload Cluster into Tanzu Mission Control Management.][add-workload-cluster]
For more information about provisioning a workload, please refer to [provision a Cluster in vSphere with Tanzu][provision-cluster-vsphere]

[management-cluster-registration]: https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/1.4/vmware-tanzu-kubernetes-grid-14/GUID-mgmt-clusters-register_tmc.html
[create-provisioner]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-BA7124EB-2A6B-46BC-839A-57609871306E.html
[add-workload-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-78908829-CB4E-459F-AA81-BEA415EC9A11.html
[provision-cluster-vsphere]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-0A1AEC6A-3E5C-424F-8EBC-1DDFC14D2688.html

You can provision a Tanzu Kubernetes Grid workload cluster in vSphere using this Terraform provider, as shown in the example below.
This example assumes that you have already registered a Tanzu Kubernetes Grid management cluster with Tanzu Mission Control.

```terraform
// Tanzu Mission Control Cluster Type: Tanzu Kubernetes Grid vSphere workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read Tanzu Mission Control Tanzu Kubernetes Grid vSphere workload cluster : fetch cluster details
data "tanzu-mission-control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Required
  provisioner_name        = "<prov-name>"          // Required
  name                    = "<cluster-name>"       // Required
}

// Create Tanzu Mission Control Tanzu Kubernetes Grid vSphere workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_vsphere_cluster" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  name                    = "<cluster-name>"       // Required

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
    tkg_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "<pods-cidr-blocks>", // Required
            ]
          }

          services {
            cidr_blocks = [
              "<services-cidr-blocks>", // Required
            ]
          }

          control_plane_end_point = "<end-point>" // Required
        }

        security {
          ssh_key = "<ssh-key>" // Required
        }
      }

      distribution {
        version = "<version>" // Required

        workspace {
          datacenter        = "<datacenter>"        // Required
          datastore         = "<datastore>"         // Required
          workspace_network = "<workspace_network>" // Required
          folder            = "<folder>"            // Required
          resource_pool     = "<resource_pool>"     // Required
        }
      }

      topology {
        control_plane {
          vm_config {
            cpu       = "<cpu>"       // Required
            disk_size = "<disk_size>" // Required
            memory    = "<memory>"    // Required
          }

          high_availability = false // Default: false
        }

        node_pools {
          spec {
            worker_node_count = "<worker-node-count>" // Required
            cloud_label = {
              "<key>" : "<val>"
            }
            node_label = {
              "<key>" : "<val>"
            }

            tkg_vsphere {
              vm_config {
                cpu       = "<cpu>"       // Required
                disk_size = "<disk_size>" // Required
                memory    = "<memory>"    // Required
              }
            }
          }

          info {
            name        = "<node-pool-name>" // Required
            description = "<node-pool-description>"
          }
        }
      }
    }
  }
}
```

When you apply these configurations, Terraform creates the workload cluster on the specified management cluster.
If you need to update the cluster, you update the rule definition and then Terraform applies it across all the sites.
If you add / or remove a site from the list, Terraform also handles creating or removing the rule on the subsequent `terraform apply` operation.