variable "endpoint" {}
variable "vmw_cloud_api_token" {}
variable "client_auth_cert_file" {}
variable "client_auth_key_file" {}
variable "ca_file" {}
variable "client_auth_cert" {}
variable "client_auth_key" {}
variable "ca_cert" {}

variable "vmw_cloud_endpoint" {
  default = "console.tanzu.broadcom.com"
}

variable "insecure_allow_unverified_ssl" {
  default = false
}
