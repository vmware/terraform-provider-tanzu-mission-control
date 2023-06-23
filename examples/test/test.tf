terraform {
  required_providers {
    tanzu-mission-control = {
      source = "example.com/quicktest/inss"
    }
  }
}

provider "tanzu-mission-control" {
  endpoint            = "shubbh-can.tmc-dev.cloud.vmware.com"
  vmw_cloud_api_token = "_Kljh0LRkYMTCTTNo4Aw2cHGC2gh95B_7OR0FIhB3MDPjRWzUIfjvtVuAXXd9ZDR"
  vmw_cloud_endpoint  = "console-stg.cloud.vmware.com"
}

resource "tanzu-mission-control_package_install" "ins" {
    name = "test3"

    namespace = "some"

    scope {
        cluster {
            name = "tf-test-2"
        }
    }

    spec {
        package_ref {
            package_metadata_name = "pkg.test.carvel.dev"
            version_selection {
                constraints = "2.0.0"
            }
        }
        
        inline_values = { "bar" : "foo", "some":12 }
    }
}

data "tanzu-mission-control_package" "get_pack" {
  name           = "2.0.0"
  namespace_name = "tanzu-package-repo-global"
  metadata_name  = "pkg.test.carvel.dev"

  scope {
    cluster {
      name = "tf-test-2"
    }
  }
}

# projects.registry.vmware.com/tmc/build-integrations/package/repository/e2e-test-unauth-repo@sha256:87a5f7e0c44523fbc35a9432c657bebce246138bbd0f16d57f5615933ceef632