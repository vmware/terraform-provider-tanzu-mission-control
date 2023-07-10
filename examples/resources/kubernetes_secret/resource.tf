# Create Tanzu Mission Control kubernetes secret with attached set as default value.
resource "tanzu-mission-control_kubernetes_secret" "create_secret" {
  name           = "tf-secret"                # Required
  namespace_name = "tf-secret-namespace-name" # Required 

  scope {
    cluster {
      cluster_name            = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  export = false # Default: false

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    docker_config_json {
      username           = "testusername"         # Required
      password           = "testpassword"         # Required
      image_registry_url = "testimageregistryurl" # Required
    }
  }
}