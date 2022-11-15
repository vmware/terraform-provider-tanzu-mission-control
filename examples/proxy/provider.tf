provider "tanzu-mission-control" {
  endpoint            = var.endpoint            # optionally use TMC_ENDPOINT env var
  vmw_cloud_api_token = var.vmw_cloud_api_token # optionally use VMW_CLOUD_API_TOKEN env var

  # if you are using dev or different csp endpoint, change the default value below
  # for production environments the vmw_cloud_endpoint is console.cloud.vmware.com
  # vmw_cloud_endpoint = "console.cloud.vmware.com" or optionally use VMW_CLOUD_ENDPOINT env var

  # the following values shall be only populated when the provider needs to be used behind a proxy.
  # these values will only work if the user provides HTTP_PROXY or HTTPS_PROXY env var
  insecure_allow_unverified_ssl = var.insecure_allow_unverified_ssl # optionally use INSECURE_ALLOW_UNVERIFIED_SSL env var
  client_auth_cert_file         = var.client_auth_cert_file         # optionally use CLIENT_AUTH_CERT_FILE env var
  client_auth_key_file          = var.client_auth_key_file          # optionally use CLIENT_AUTH_KEY_FILE env var
  ca_file                       = var.ca_file                       # optionally use CA_FILE env var
  client_auth_cert              = var.client_auth_cert              # optionally use CLIENT_AUTH_CERT env var
  client_auth_key               = var.client_auth_key               # optionally use CLIENT_AUTH_KEY env var
  ca_cert                       = var.ca_cert                       # optionally use CA_CERT env var
}
