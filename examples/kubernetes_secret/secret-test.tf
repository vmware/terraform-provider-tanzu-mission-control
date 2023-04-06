terraform {
  required_providers {
    tanzu-mission-control = {
        source = "vmware/dev/tanzu-mission-control"
    }
  }
}

variable "vmw_cloud_api_token" { type= string } 

provider "tanzu-mission-control" {
    endpoint = "shubbhang-stable.tmc-dev.cloud.vmware.com"
    vmw_cloud_api_token = "${var.vmw_cloud_api_token}"
    vmw_cloud_endpoint = "console-stg.cloud.vmware.com"
}

# create cluster group resource

resource "tanzu-mission-control_kubernetes_secret" "create_secret" {
    name = "shubbhang-demo-secret1"

    namespace_name = "default"

    scope {
        cluster {
          cluster_name = "secret-test-demo" 

          provisioner_name = "attached"

          management_cluster_name = "attached"
        }
    }

    export = false

    spec {
        docker_config_json {
            username = "vshubhang3"
            password = "vshubhang1"
            image_registry_url = "some_url2.com"
        }
    }
}

data "tanzu-mission-control_kubernetes_secret" "read_secret" {
    name = tanzu-mission-control_kubernetes_secret.create_secret.name

    namespace_name = tanzu-mission-control_kubernetes_secret.create_secret.namespace_name

    scope {
      cluster {
        cluster_name = tanzu-mission-control_kubernetes_secret.create_secret.scope[0].cluster[0].cluster_name
      }
    }
}