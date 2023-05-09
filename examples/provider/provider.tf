# Provider configuration for TMC SaaS
provider "tanzu-mission-control" {
  endpoint            = var.endpoint            # optionally use TMC_ENDPOINT env var
  vmw_cloud_api_token = var.vmw_cloud_api_token # optionally use VMW_CLOUD_API_TOKEN env var

  # if you are using dev or different csp endpoint, change the default value below
  # for production environments the vmw_cloud_endpoint is console.cloud.vmware.com
  # vmw_cloud_endpoint = "console.cloud.vmware.com" or optionally use VMW_CLOUD_ENDPOINT env var
}

# Provider configuration for TMC Self-Managed
provider "tanzu-mission-control" {
  endpoint = "example.local-dev.tmc.com"

  self_managed {
    oidc_issuer = "pinniped-supervisor.example.local-dev.tmc.com"
    username = "testuser01@tmcselfmanaged.com"
    password = "dummy_password"
  }
}