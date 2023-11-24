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