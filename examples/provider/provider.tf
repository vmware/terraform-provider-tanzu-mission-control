# Provider configuration for TMC SaaS
provider "tanzu-mission-control" {
  endpoint            = var.endpoint            # optionally use TMC_ENDPOINT env var
  vmw_cloud_api_token = var.vmw_cloud_api_token # optionally use VMW_CLOUD_API_TOKEN env var

  # if you are using dev or different csp endpoint, change the default value below
  # for production environments the vmw_cloud_endpoint is console.tanzu.broadcom.com
  # vmw_cloud_endpoint = "console.tanzu.broadcom.com" or optionally use VMW_CLOUD_ENDPOINT env var
}

# Provider configuration for TMC Self-Managed
provider "tanzu-mission-control" {
  endpoint = var.endpoint # optionally use TMC_ENDPOINT env var

  self_managed {
    oidc_issuer = var.oidc_issuer # optionally use OIDC_ISSUER env var,  Ex: export OIDC_ISSUER=pinniped-supervisor.example.local-dev.tmc.com
    username    = var.username    # optionally use TMC_SM_USERNAME env var
    password    = var.password    # optionally use TMC_SM_PASSWORD env var
  }
  ca_file = var.ca_file # Path to Host's root ca set. The certificates issued by the issuer should be trusted by the host accessing TMC Self-Managed via TMC terraform provider.
}