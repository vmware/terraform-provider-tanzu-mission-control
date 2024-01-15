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
