# Create Tanzu Mission Control kubernetes secret with attached set as default value.
# Example for creating the dockerconfigjson secret
resource "tanzu-mission-control_kubernetes_secret" "create_dockerconfigjson_secret" {
  name           = "tf-secret"                # Required
  namespace_name = "tf-secret-namespace-name" # Required 

  scope {
    cluster_group {
      name = "default" # Required
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

# Example for creating the opaque secret
resource "tanzu-mission-control_kubernetes_secret" "create_opaque_secret" {
  name           = "tf-secret"                # Required
  namespace_name = "tf-secret-namespace-name" # Required 

  scope {
    cluster_group {
      name = "default" # Required
    }
  }

  export = false # Default: false

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    opaque = {
      "key1" : "value1"
      "key2" : "value2"
    }
  }
}
