resource "tanzu-mission-control_target_location" "demo_aws_self_provisioned" {
  name          = "TARGET_LOCATION_NAME"

  spec {
    target_provider = "AWS"
    credential      = {
      name = "AWS_CREDENTIAL_NAME"
    }

    bucket = "BUCKET_NAME"
    region = "REGION"

    config {
      aws {
        s3_force_path_style = false
        s3_bucket_url       = "AWS_S3_BUCKET_URL"
        s3_public_url       = "AWS_S3_PUBLIC_URL"
      }
    }

    assigned_groups {
      cluster {
        management_cluster_name = "MGMT_CLS_NAME"
        provisioner_name        = "PROVISIONER_NAME"
        name                    = "CLS_NAME"
      }

      cluster {
        management_cluster_name = "MGMT_CLS_NAME"
        provisioner_name        = "PROVISIONER_NAME"
        name                    = "CLS_NAME"
      }

      cluster_groups = ["CLS_GROUP_NAME", "CLS_GROUP_NAME"]
    }
  }
}