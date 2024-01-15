---
Title: "Data Protection of a Tanzu Kubernetes Cluster"
Description: |-
    An example of using Data Protection Feature for a Tanzu Kubernetes Cluster/Cluster Group
---
# Enable Data Protection

The `tanzu-mission-control_enable_data_protection` resource enables users to activate and set up data protection for a Tanzu Kubernetes Cluster.
Once enabled, users can create instant backups or schedule backups for later.

**NOTE :** To resolve cluster and cluster group data protection conflicts use the below command
``terraform refresh``

For more information regarding data protection, see [Data Protection][data-protection].

[data-protection]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-C16557BC-EB1B-4414-8E63-28AD92E0CAE5.html


# Target Location

The `"tanzu-mission-control_target_location` resource enables users to create and configure target locations for data protection backups.
Once created, a target location can be used to store cluster backups.

**NOTE**: The type of a target location is inherited from the configured credentials type which can be either "TMC Managed" or "Self Managed".

For more information regarding target location, see [Target Location][target-location].

[target-location]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-867683CE-8AF0-4DC7-9121-81AD507EDB3B.html

# Backup Schedule

The `tanzu-mission-control_backup_schedule` resource enables users to create and configure scheduled backups in a cluster/cluster-group.

NOTE : To resolve cluster and cluster group backup schedule conflicts use the below command
``terraform refresh``

Backups can be applied in 3 levels:

* Entire/Full Cluster
* Selected Namespaces
* Resources Selection By Label Selector

For more information regarding scheduled backups, see [Scheduled Backups][backup-schedule].

[backup-schedule]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-89926F80-050A-4F1C-9D04-D56D5F453995.html?hWord=N4IghgNiBcIEZgMYGsCuAHABAZ0QCwFMATVCAkAXyA

## Sample usage of Data Protection

You can use the following template as reference for enabling all stages of data protection feature of Tanzu Mission Control using Terraform (i.e.) Enable data-protection for cluster, Set a Target Location of backup and finally, set a backup schedule for periodic data protection.

```terraform
// Tanzu Mission Control Data Protection Feature

locals {
  cluster_name = "<cluster-name>"
  management_cluster_name = "<management-cluster-name>"
  provisioner_name        = "<provisioner-name>"
}

// Enable Data Protection
resource "tanzu-mission-control_enable_data_protection" "data_protection" {
  scope {
    cluster {
      cluster_name            = local.cluster_name
      management_cluster_name = local.management_cluster_name
      provisioner_name        = local.provisioner_name
    }
  }

  spec {
    disable_restic                       = false // Default: false
    enable_csi_snapshots                 = false // Default: false
    enable_all_api_group_versions_backup = false // Default: false
  }

  deletion_policy {
    delete_backups = false // Default: false
  }
}

// Create Target Location for Scheduled Back Up
// Self managed AWS Target Location
resource "tanzu-mission-control_target_location" "aws_self_provisioned" {
  name          = "<target-location-name>"

  spec {
    target_provider = "AWS"
    credential      = {
      name = "<aws-credential-name?"
    }

    bucket = "<bucket-name>"
    region = "<region>"

    config {
      aws {
        s3_force_path_style = false
        s3_bucket_url       = "<aws-s3-bucket-url>"
        s3_public_url       = "<aws-s3-public-url>"
      }
    }

    assigned_groups {
      cluster {
        name                    = local.cluster_name
        management_cluster_name = local.management_cluster_name
        provisioner_name        = local.provisioner_name
      }

      cluster_groups = ["<cluster-group-name-1>", "<cluster-group-name-2>"]
    }
  }
}

// Create Full Cluster Scheduled Back Up
resource "tanzu-mission-control_backup_schedule" "backup_full_cluster" {
  name                    = "<scheduled-backup-name>"
  scope {
    cluster {
      cluster_name            = local.cluster_name
      management_cluster_name = local.management_cluster_name
      provisioner_name        = local.provisioner_name
    }
  }

  backup_scope = "FULL_CLUSTER"

  spec {
    schedule {
      rate = "0 12 * * 1"
    }

    template {
      backup_ttl          = "2592000s"
      excluded_namespaces = [
        "<namespace-1>",
      ]
      excluded_resources = [
        "<resource-1>",
        "<resource-2>"
      ]

      storage_location = tanzu-mission-control_target_location.aws_self_provisioned.name
    }
  }
}
```