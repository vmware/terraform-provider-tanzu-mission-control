# Create Tanzu Mission Control image registry credential and attach cluster entry with the same credential.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
  name                    = "demo-lir" // Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config-path>" // Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group  = "default" // Default: default
    image_registry = tanzu-mission-control_credential.img_reg_cred.name
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
}

# Create IMAGE_REGISTRY credential
resource "tanzu-mission-control_credential" "img_reg_cred" {
  name = "lir-cred-name"

  meta {
    description = "credential"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "registry-namespace": "tmc"
    }
  }

  spec {
    capability = "IMAGE_REGISTRY"
    provider   = "GENERIC_KEY_VALUE"
    data {
      key_value {
        data = {
          "registry-url" = "dev.registry.tanzu.vmware.com"
        }
      }
    }
  }
}