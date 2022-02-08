// Tanzu Mission Control terraform provider initialization

terraform {
  required_providers {
    tanzu-mission-control = {
      source  = "vmware/tanzu-mission-control"
      version = "1.0.0"
    }
  }
}

// Basic details needed to configure Tanzu Mission Control provider
provider "tanzu-mission-control" {
  endpoint            = "<stack-name>.tmc.cloud.vmware.com" // Required, environment variable: TMC_ENDPOINT
  vmw_cloud_api_token = "<vmw-cloud-api-token>"             // Required, environment variable: VMW_CLOUD_API_TOKEN
}
