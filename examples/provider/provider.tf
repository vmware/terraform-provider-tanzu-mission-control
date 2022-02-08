provider "tanzu-mission-control" {
  endpoint            = var.endpoint            # optionally use TMC_ENDPOINT env var
  vmw_cloud_api_token = var.vmw_cloud_api_token # optionally use VMW_CLOUD_API_TOKEN env var

  # if you are using dev or different csp endpoint, change the default value below
  # for production environments the csp_endpoint is console.cloud.vmware.com
  # vmw_cloud_api_endpoint = "console.cloud.vmware.com" or optionally use VMW_CLOUD_ENDPOINT env var
}