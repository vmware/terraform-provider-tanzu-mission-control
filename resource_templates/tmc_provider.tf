// Tanzu Mission Control terraform provider initialization

//[Mac Users] if you are using developer build of tmc_provider,
//please place the binary under : ~/.terraform.d/plugins/vmware/tanzu/Tanzu Mission Control/0.1.1/darwin_amd64/ or run `make build`

terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}

// Basic details needed to configure Tanzu Mission Control provider
provider "tmc" {
  endpoint = "<stack-name>.tmc-dev.cloud.vmware.com" // Required, environment variable: TMC_ENDPOINT
  vmw_cloud_api_token    = "<vmw-cloud-api-token>"   // Required, environment variable: VMW_CLOUD_API_TOKEN
}