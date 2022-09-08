---
Title: "Provisioning of a workload cluster"
Description: |-
    An example of provisioning Tanzu Kubernetes Grid Service, Tanzu Kubernetes Grid vSphere and Tanzu Kubernetes Grid AWS workload clusters.
---
# Cluster

The `tanzu-mission-control_cluster` resource enables you to attach conformant Kubernetes clusters for management through Tanzu Mission Control.
With Tanzu Kubernetes clusters, you can also provision resources to create new workload clusters.

## Tanzu Kubernetes Grid Service Workload Cluster

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
For more information about provisioning a workload, please refer to [provision a Cluster in vSphere with Tanzu][provision-cluster-vsphere]

[supervisor-cluster-registration]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03.html#GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03
[vSphere-namespaces]: https://docs.vmware.com/en/VMware-vSphere/7.0/vmware-vsphere-with-tanzu/GUID-1544C9FE-0B23-434E-B823-C59EFC2F7309.html
[add-workload-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-78908829-CB4E-459F-AA81-BEA415EC9A11.html
[provision-cluster-vsphere]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-0A1AEC6A-3E5C-424F-8EBC-1DDFC14D2688.html

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
        storage {
          classes = [
            "<storage-classes>",
          ]
          default_class = "<default-storage-class>"
        }
      }

      distribution {
        version = "<version>" // Required
      }

      topology {
        control_plane {
          class             = "<class>"         // Required
          storage_class     = "<storage-class>" // Required
          high_availability = false             // Default: false
          volumes {
            capacity          = volume-capacity
            mount_path        = "<mount-path>"
            name              = "<volume-name>"
            pvc_storage_class = "<storage-class>"
          }
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
              volumes {
                capacity          = volume-capacity
                mount_path        = "<mount-path>"
                name              = "<volume-name>"
                pvc_storage_class = "<storage-class>"
              }
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

## Tanzu Kubernetes Grid vSphere Workload Cluster

Before provisioning a Tanzu Kubernetes Grid vSphere workload cluster using this Terraform provider we need the following prerequisites.

**Prerequisites:**

- Register the Tanzu Kubernetes Grid management cluster in Tanzu Mission Control.
Please refer to [register Management Cluster with Tanzu Mission Control.][management-cluster-registration]

- Create a provisioner under the management cluster or reuse the existing provisioners under the management cluster. Please refer to [create a Provisioner.][create-provisioner]

- Ensure the CSP token used in initialising the terraform provider has the right set of permissions to create a workload cluster.

Once you have the `management cluster name` and `provisioner name` from Tanzu mission control, we are all set to provision a workload under the chosen management cluster name using the terraform script (example below).

If you need to update the cluster, you update the rule definition and then Terraform apply applies it across all the sites.
If you add or remove a site from the list, Terraform also handles creating or removing the rule on the subsequent `terraform apply`operation.

If you want to manage the already existing workload, refer to [adding a Workload Cluster into Tanzu Mission Control Management.][add-workload-cluster]
For more information about provisioning a workload, please refer to [provision a workload cluster in vSphere][provision-cluster-vsphere]

[management-cluster-registration]: https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/1.4/vmware-tanzu-kubernetes-grid-14/GUID-mgmt-clusters-register_tmc.html
[create-provisioner]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-BA7124EB-2A6B-46BC-839A-57609871306E.html
[add-workload-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-78908829-CB4E-459F-AA81-BEA415EC9A11.html
[provision-cluster-vsphere]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-9435CCAC-F90B-4575-9D73-D26315871C8A.html

You can provision a Tanzu Kubernetes Grid workload cluster on vSphere using this Terraform provider, as shown in the example below.
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

          api_server_port         = api-server-port-default-value
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

## Tanzu Kubernetes Grid AWS Workload Cluster

Before provisioning a Tanzu Kubernetes Grid AWS workload cluster using this Terraform provider we need the following prerequisites.

**Prerequisites:**

- Register the Tanzu Kubernetes Grid management cluster in Tanzu Mission Control.
Please refer to [register Management Cluster with Tanzu Mission Control.][management-cluster-registration]

- Create a provisioner under the management cluster or reuse the existing provisioners under the management cluster. Please refer to [create a Provisioner.][create-provisioner]

- Ensure the CSP token used in initialising the terraform provider has the right set of permissions to create a workload cluster.

Once you have the `management cluster name` and `provisioner name` from Tanzu mission control, we are all set to provision a workload under the chosen management cluster name using the terraform script (example below).

If you need to update the cluster, you update the rule definition and then Terraform apply applies it across all the sites.
If you add or remove a site from the list, Terraform also handles creating or removing the rule on the subsequent `terraform apply`operation.

If you want to manage the already existing workload, refer to [adding a Workload Cluster into Tanzu Mission Control Management.][add-workload-cluster]
For more information about provisioning a workload, please refer to [provision a workload cluster in AWS][provision-cluster-aws]

[management-cluster-registration]: https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/1.4/vmware-tanzu-kubernetes-grid-14/GUID-mgmt-clusters-register_tmc.html
[create-provisioner]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-BA7124EB-2A6B-46BC-839A-57609871306E.html
[add-workload-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-78908829-CB4E-459F-AA81-BEA415EC9A11.html
[provision-cluster-aws]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-E5F9242B-CBB1-4142-B089-3E16EED102F4.html

You can provision a Tanzu Kubernetes Grid workload cluster on AWS using this Terraform provider, as shown in the example below.
This example assumes that you have already registered a Tanzu Kubernetes Grid management cluster with Tanzu Mission Control.

```terraform
// Tanzu Mission Control Cluster Type: Tanzu Kubernetes Grid AWS workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster : fetch cluster details
data "tanzu-mission-control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Required
  provisioner_name        = "<prov-name>"          // Required
  name                    = "<cluster-name>"       // Required
}

// Create Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_aws_cluster" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  name                    = "<cluster-name>"       // Required

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
    tkg_aws {
      settings {
        network {
          cluster {
            pods {
              cidr_blocks = "<pods-cidr-blocks>" // Required
            }

            services {
              cidr_blocks = "<services-cidr-blocks>" // Required
            }
            api_server_port = api-server-port-default-value
          }
          provider {
            subnets {
              availability_zone = "<availability-zone>"
              cidr_block        = "<subnets-cidr-blocks>"
              is_public         = false
            }
            subnets {
              availability_zone = "<availability-zone>"
              cidr_block        = "<subnets-cidr-blocks>"
              is_public         = true
            }
            vpc {
              cidr_block = "<vpc-cidr-blocks>"
            }
          }
        }

        security {
          ssh_key = "<ssh-key>" // Required
        }
      }

      distribution {
        region  = "<region>"  // Required
        version = "<version>" // Required
      }

      topology {
        control_plane {
          availability_zones = [
            "<availability-zone>",
          ]
          instance_type = "<instance-type>"
        }

        node_pools {
          spec {
            worker_node_count = "<worker-node-count>"
            tkg_aws {
              availability_zone = "<availability-zone>"
              instance_type     = "<instance-type>"
              node_placement {
                aws_availability_zone = "<availability_zone>"
              }
              version = "<version>"
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

**Note:**
For the AWS workload cluster resource, kindly follow the ordering of the subnet blocks as described in the example above.

When you apply these configurations, Terraform creates the workload cluster on the specified management cluster.
If you need to update the cluster, you update the rule definition and then Terraform applies it across all the sites.
If you add / or remove a site from the list, Terraform also handles creating or removing the rule on the subsequent `terraform apply` operation.