---
Title: "Tanzu Mission Control Provider"
Description: |-
  The Tanzu Mission Control provider facilitates provisioning of Tanzu Mission Control Resources through Terraform plugin.
---

# Tanzu Mission Control Provider

Manage your Kubernetes clusters in Tanzu Mission Control using this Terraform provider.
The Tanzu Mission Control provider facilitates the provisioning of resources that you can use to manage Tanzu Kubernetes Grid workload clusters (and other conformant Kubernetes clusters) in Tanzu Mission Control.
For more information, please refer [What is Tanzu Mission Control][vmware-tanzu-tmc] in VMware Tanzu Mission Control Concepts.

[vmware-tanzu-tmc]: https://tanzu.vmware.com/mission-control

# Tanzu Mission Control SaaS offering

To use the **Tanzu Mission Control provider** for Terraform, you must have access to Tanzu Mission Control SaaS offering through an VMware Cloud services organization.
Prior to initializing this provider in Terraform, make sure you have the following information:

- The endpoint for your Tanzu Mission Control organization.
- An active API token.

To gather this information, you need to do the following:

1. [Log in to the Tanzu Mission Control console.][login]

    After you have logged in, the domain in the URL displayed in your browser represents the endpoint for your organization.

    For example, the URL for the Clusters page in the Tanzu Mission Control console is something like this:

    `my-org.tmc.cloud.vmware.com/clusters`

    Here, the endpoint is `my-org.tmc.cloud.vmware.com`

[login]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-855A8998-E19A-46AC-A833-12C347486EF7.html

2. In the Tanzu Mission Control console, click on your name in the top right corner, and then click **My Account**.

3. On the My Account page of the VMware Cloud Services console, click the **API Tokens** tab.

From this page, you can generate a new API token, and then copy it to use for the Tanzu Mission Control provider in Terraform.

~> **Note:**
Current version of Tanzu Mission Control provider does not support when API tokens are secured using [multi-factor authentication][mfa-for-api-token].

[mfa-for-api-token]: https://docs.vmware.com/en/VMware-Cloud-services/services/Using-VMware-Cloud-Services/GUID-38D09558-D468-4A21-95BD-581940119FA7.html

# Tanzu Mission Control Self-Managed

The Tanzu Mission Control provider also facilitates the provisioning of resources that you can use to manage Tanzu Kubernetes Grid workload clusters in Tanzu Mission Control Self-Managed.
Similar to SaaS offering, Tanzu Mission Control Self-Managed provides a single point of control that allows you to securely manage the infrastructure and apps for your Kubernetes footprint. However, Tanzu Mission Control Self-Managed runs as a service deployed to a Kubernetes cluster running in your own data center.
For more information, please refer [What is Tanzu Mission Control Self-Managed][vmware-tanzu-tmc-self-managed] in VMware Tanzu Mission Control Concepts.

[vmware-tanzu-tmc-self-managed]: https://tanzu.vmware.com/content/blog/vmware-tanzu-mission-control-self-managed-announcement

To use the **Tanzu Mission Control provider** for Tanzu Mission Control Self-Managed prior to initializing this provider in Terraform, make sure you have the following information of the deployed Tanzu Mission Control Self-Managed instance:

- The endpoint URL of your Tanzu Mission Control Self-Managed deployment.
- IDP credentials. To log in to the Tanzu Mission Control console in a self-managed deployment, you must be a user enrolled in the IDP associated to Tanzu Mission Control Self-Managed. For more information, see the section on setting up authentication in [Preparing your cluster for Tanzu Mission Control Self-Managed][prepapre-cluster-for-tmc-sm].

[prepapre-cluster-for-tmc-sm]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/1.0/tanzumc-sm-install/prepare-cluster.html

-> **Note:**
Tanzu Mission Control Provider supports Tanzu Mission Control Self-Managed from **v1.1.9** onwards.

## Example Usage

```terraform
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
  endpoint = var.endpoint               # optionally use TMC_ENDPOINT env var

  self_managed {
    oidc_issuer   = var.oidc_issuer        # optionally use OIDC_ISSUER env var,  Ex: export OIDC_ISSUER=pinniped-supervisor.example.local-dev.tmc.com
    username      = var.username           # optionally use TMC_SM_USERNAME env var
    password      = var.password           # optionally use TMC_SM_PASSWORD env var
  }
  ca_file = var.ca_file                    # Path to Host's root ca set. The certificates issued by the issuer should be trusted by the host accessing TMC Self-Managed via TMC terraform provider.
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ca_cert` (String, Sensitive)
- `ca_file` (String)
- `client_auth_cert` (String, Sensitive)
- `client_auth_cert_file` (String)
- `client_auth_key` (String, Sensitive)
- `client_auth_key_file` (String)
- `endpoint` (String)
- `insecure_allow_unverified_ssl` (Boolean)
- `self_managed` (Block List, Max: 1) (see [below for nested schema](#nestedblock--self_managed))
- `vmw_cloud_api_token` (String, Sensitive)
- `vmw_cloud_endpoint` (String)

<a id="nestedblock--self_managed"></a>
### Nested Schema for `self_managed`

Optional:

- `oidc_issuer` (String) URL of the OpenID Connect (OIDC) issuer configured with self-managed Taznu mission control instance
- `password` (String, Sensitive) Password for the above mentioned Username field configured in the OIDC
- `username` (String) Username configured in the OIDC
