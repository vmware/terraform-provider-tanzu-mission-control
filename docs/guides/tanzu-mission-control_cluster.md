---
Title: "Provisioning of a workload cluster"
Description: |-
    An example of provisioning Tanzu Kubernetes Grid Service, Tanzu Kubernetes Grid vSphere and Tanzu Kubernetes Grid AWS workload clusters.
---
# Cluster

The `tanzu-mission-control_cluster` resource enables you to attach conformant Kubernetes clusters for management through Tanzu Mission Control.
With Tanzu Kubernetes clusters, you can also provision resources to create new workload clusters.

## Tanzu Kubernetes Grid Service Workload Cluster

Before creating a Tanzu Kubernetes Grid Service workload cluster in vSphere Supervisor using this Terraform provider we need the following prerequisites.

~> **Note:**
Current version of `tanzu-mission-control_cluster` resource in Tanzu Mission Control provider supports creation of deafult nodepool and updates to some fields specifically `worker_node_count, class, storage_class` in case of Tanzu Kubernetes Grid Service workload cluster and updation of `worker_node_count` in case Tanzu Kubernetes Grid vSphere Workload Cluster of default nodepool. Deletion of the default nodepool is not yet supported via this resource and will be added in the upcoming releases.
All other nodepools except the default nodepool can be managed via `tanzu-mission-control_cluster_node_pool` resource.

**Prerequisites:**

- Register the Tanzu Kubernetes Grid Service management cluster in Tanzu Mission Control.
Note that the Tanzu Kubernetes Grid Service management cluster must be **ready** and **healthy**.
Please refer to [registration of a Supervisor Cluster in vSphere Supervisor.][supervisor-cluster-registration]

- Create a provisioner under the management cluster or reuse the existing providers under the management cluster. Please refer to [working with vSphere Namespaces on a Supervisor Cluster.][vsphere-namespaces]

- Ensure the CSP token used in initialising the terraform provider has the right set of permissions to create a workload cluster.

Once you have the `management cluster name` and `provisioner name` from Tanzu Mission Control, we are all set to provision a workload under the chosen management cluster name using the terraform script (example below).

If you need to update the cluster, you update the rule definition and then Terraform apply applies it across all the sites.
If you add or remove a site from the list, Terraform also handles creating or removing the rule on the subsequent `terraform apply`operation.

If you want to manage the already existing workload, Refer to [adding a Workload Cluster into Tanzu Mission Control Management.][add-workload-cluster]
For more information about provisioning a workload, please refer to [provision a Cluster in vSphere Supervisor][provision-cluster-vsphere]

[supervisor-cluster-registration]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03.html#GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03
[vsphere-namespaces]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere-supervisor/7-0/vsphere-with-tanzu-configuration-and-management-7-0.html
[add-workload-cluster]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-78908829-CB4E-459F-AA81-BEA415EC9A11.html
[provision-cluster-vsphere]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-0A1AEC6A-3E5C-424F-8EBC-1DDFC14D2688.html

You can provision a Tanzu Kubernetes Grid Service workload cluster in vSphere Supervisor using this Terraform provider, as shown in the example below.
This example assumes that you have already registered a Tanzu Kubernetes Grid Service management cluster with Tanzu Mission Control.

```terraform
// Tanzu Mission Control Cluster Type: Tanzu Kubernetes Grid Service workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster : fetch cluster details for already present TKGs cluster
data "tanzu-mission-control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<existing-management-cluster>" // Required
  provisioner_name        = "<existing-prov-name>"          // Required
  name                    = "<existing-cluster-name>"       // Required
}

# Create Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster entry
resource "tanzu-mission-control_cluster" "tkgs_workload" {
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
              class          = "<class>"         // Required
              storage_class  = "<storage-class>" // Required
              failure_domain = "<failure-domain>"
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

Once you have the `management cluster name` and `provisioner name` from Tanzu Mission Control, we are all set to provision a workload under the chosen management cluster name using the terraform script (example below).

If you need to update the cluster, you update the rule definition and then Terraform apply applies it across all the sites.
If you add or remove a site from the list, Terraform also handles creating or removing the rule on the subsequent `terraform apply`operation.

If you want to manage the already existing workload, refer to [adding a Workload Cluster into Tanzu Mission Control Management.][add-workload-cluster]
For more information about provisioning a workload, please refer to [provision a workload cluster in vSphere][provision-cluster-vsphere]

[management-cluster-registration]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-kubernetes-grid/2-5/tkg/mgmt-deploy-post-deploy.html
[create-provisioner]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-BA7124EB-2A6B-46BC-839A-57609871306E.html
[add-workload-cluster]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-78908829-CB4E-459F-AA81-BEA415EC9A11.html
[provision-cluster-vsphere]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-9435CCAC-F90B-4575-9D73-D26315871C8A.html

You can provision a Tanzu Kubernetes Grid workload cluster on vSphere using this Terraform provider, as shown in the example below.
This example assumes that you have already registered a Tanzu Kubernetes Grid management cluster with Tanzu Mission Control.

```terraform
// Tanzu Mission Control Cluster Type: Tanzu Kubernetes Grid vSphere workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read Tanzu Mission Control Tanzu Kubernetes Grid vSphere workload cluster : fetch cluster details for already present TKG vSphere cluster
data "tanzu-mission-control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<existing-management-cluster>" // Required
  provisioner_name        = "<existing-prov-name>"          // Required
  name                    = "<existing-cluster-name>"       // Required
}

// Create Tanzu Mission Control Tanzu Kubernetes Grid vSphere workload cluster entry
resource "tanzu-mission-control_cluster" "tkg_vsphere_cluster" {
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
      advanced_configs {
        key   = "<key>"
        value = "<value>"
      }
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
          control_plane_end_point = "<end-point>" // Optional
        }

        security {
          ssh_key = "<ssh-key>" // Required
        }
      }

      distribution {
        os_arch    = "<os-arch>"
        os_name    = "<os-name>"
        os_version = "<os-version>"
        version    = "<version>" // Required

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

Once you have the `management cluster name` and `provisioner name` from Tanzu Mission Control, we are all set to provision a workload under the chosen management cluster name using the terraform script (example below).

If you need to update the cluster, you update the rule definition and then Terraform apply applies it across all the sites.
If you add or remove a site from the list, Terraform also handles creating or removing the rule on the subsequent `terraform apply`operation.

If you want to manage the already existing workload, refer to [adding a Workload Cluster into Tanzu Mission Control Management.][add-workload-cluster]
For more information about provisioning a workload, please refer to [provision a workload cluster in AWS][provision-cluster-aws]

[management-cluster-registration]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-kubernetes-grid/2-5/tkg/mgmt-deploy-post-deploy.html
[create-provisioner]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-C778E447-DDBB-49FC-B0B2-A8012AC56B0E.html
[add-workload-cluster]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-78908829-CB4E-459F-AA81-BEA415EC9A11.html
[provision-cluster-aws]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-E5F9242B-CBB1-4142-B089-3E16EED102F4.html

You can provision a Tanzu Kubernetes Grid workload cluster on AWS using this Terraform provider, as shown in the example below.
This example assumes that you have already registered a Tanzu Kubernetes Grid management cluster with Tanzu Mission Control.

```terraform
// Tanzu Mission Control Cluster Type: Tanzu Kubernetes Grid AWS workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster : fetch cluster details for already present TKG AWS cluster
data "tanzu-mission-control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<existing-management-cluster>" // Required
  provisioner_name        = "<existing-prov-name>"          // Required
  name                    = "<existing-cluster-name>"       // Required
}

// Create Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster entry
resource "tanzu-mission-control_cluster" "tkg_aws_cluster" {
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
      advanced_configs {
        key   = "<key>"
        value = "<value>"
      }
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
        os_arch    = "<os-arch>"
        os_name    = "<os-name>"
        os_version = "<os-version>"
        region     = "<region>"  // Required
        version    = "<version>" // Required
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
