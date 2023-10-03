//go:build ignore
// +build ignore

package help

/*
######### Terraform Resource Example #########
resource "tanzu-mission-control_target_location" "demo" {
  name          = "gil-test" //required
  provider_name = "tmc" //required
  #  description   = "<string>" //optional

  spec {
    target_provider = "AWS" //required

    config {
      //required
      aws {
        // either this or azure_config are allowed but not both
        s3_force_path_style = false //default false
        s3_bucket_url       = "https://SOME.BUCKET.URL"
        s3_public_url       = "https://SOME.BUCKET.PUBLIC-URL"
      }
      azure {
        // either this or s3_config are allowed but not both
        resource_group  = "string"
        storage_account = "string"
        subscription_id = "string"
      }
    }

    bucket     = "SOME" //required
    region     = "us-east-1" //required
    credential = {
      //required
      name = "vmware"
    }

    assigned_groups {
      cluster {
        management_cluster_name = "vrabbi-tkg-mgmt"
        provisioner_name        = "default"
        name                    = "aaa-cls-01"
      }

      cluster {
        management_cluster_name = "vrabbi-tkg-mgmt"
        provisioner_name        = "default"
        name                    = "bbb-cls-01"
      }

      cluster_groups = ["default", "henry-tkg-group"]
    }

    ca_cert = "IHaveNoCert" //optional
  }
}
*/
