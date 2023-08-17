terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/dev/tanzu-mission-control"
    }
  }
}

terraform {
  backend "local" {
    path = "./terraform.tfstate"
  }
}

provider "tanzu-mission-control" {
}

resource "tanzu-mission-control_akscluster" "demo_AKS_cluster" {
  credential_name = "azure-credential-name"
  subscription_id = "azure-subscription-id"
  resource_group  = "azure-resource-group"
  name            = "azure-cluster-name"
  meta {
    description = "aks test cluster"
    labels      = { "key1" : "value1" }
  }
  spec {
    config {
      location           = "eastus"
      kubernetes_version = "1.24.10"
      network_config {
        dns_prefix = "dns-tf-test"
      }
    }
    nodepool {
      name = "systemnp"
      spec {
        count   = 1
        mode    = "SYSTEM"
        vm_size = "Standard_DS2_v2"
      }
    }
  }
}