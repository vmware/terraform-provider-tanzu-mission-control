---
Title: "Credential Resource"
Description: |-
    Creating the credential resource for different use-cases.
---

# Tanzu Mission Control AWS EKS credential

## Example Usage

```terraform
# Create AWS_EKS credential
resource "tanzu-mission-control_credential" "aws_eks_cred" {
  name = "test-cred-name"

  meta {
    description = "credential"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "MANAGED_K8S_PROVIDER"
    provider   = "AWS_EKS"
    data {
      aws_credential {
        account_id = "account-id"
        iam_role {
          arn    = "arn:aws:iam::4987398738934:role/clusterlifecycle-test.tmc.cloud.vmware.com"
          ext_id = ""
        }
      }
    }
  }
  ready_wait_timeout = "2m" // Wait time for credential create operations to finish (default: 3m).
}
```

# Tanzu Mission Control Azure AKS credential

## Example Usage

```terraform
# Create AZURE_AKS credential resource
resource "tanzu-mission-control_credential" "azure_aks_cred" {
  name = "tf-azure-aks-cred-test"

  meta {
    description = "Azure AKS credential for AKS cluster lifecycle management"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "MANAGED_K8S_PROVIDER"
    provider   = "AZURE_AKS"
    data {
      azure_credential {
        service_principal_with_certificate {
          subscription_id    = "some_subscription_id"
          tenant_id          = "some_tenant_id"
          client_id          = "some_client_id"
          client_certificate = "-----BEGIN PRIVATE KEY-----\nMIvIEFk\nv+FiTAfd5MYtJYjkuU7MVA==\n-----END PRIVATE KEY-----\n-----BEGIN CERTIFICATE-----\nMIICoTCCAYkCAgPoMA0GCSqGSIb3DQEBBQUAMXyNh2KI=\n-----END CERTIFICATE-----\n"
        }
      }
    }
  }
  ready_wait_timeout = "5m" // Wait time for credential create operations to finish (default: 3m).
}
```

# IMAGE REGISTRY credential

## Example Usage

```terraform
# Create IMAGE_REGISTRY credential
resource "tanzu-mission-control_credential" "img_reg_cred" {
  name = "test-cred-name"

  meta {
    description = "credential"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "repository-namespace" : "something"
    }
  }

  spec {
    capability = "IMAGE_REGISTRY"
    provider   = "GENERIC_KEY_VALUE"
    data {
      key_value {
        data = {
          "registry-url" = "somethingnew"
          "ca-cert"      = "ca bundle"
        }
      }
    }
  }
}
```

# Cluster proxy credential

### NOTE:
For proxy credential add the annotation `proxyType : explicit` for explicit proxy, `proxyType : transparent` for transparent proxy. When no such annotation is specified by default it is assumed to be explicit proxy credential.

## Example Usage

```terraform
# Create explicit cluster proxy credential
resource "tanzu-mission-control_credential" "explicit_proxy_cred" {
  name = "explicit_proxy_cred"

  meta {
    description = "explicit proxy credential"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "proxyType" : "explicit",
      "httpProxy" : "http://sfsdf.com:123",
      "httpsProxy" : "http://sfsdf.com:123",
      "noProxyList" : "http://noproxy.com,http://something.com"
    }
  }

  spec {
    capability = "PROXY_CONFIG"
    provider   = "GENERIC_KEY_VALUE"
    data {
      key_value {
        data = {
          "httpUserName"  = "username"
          "httpPassword"  = "password"
          "httpsUserName" = "username"
          "httpsPassword" = "password"
          "proxyCABundle" = "-----BEGIN CERTIFICATE-----\n Encoded string for encryption of data\n ----END CERTIFICATE----" # chain of certificate is supported in CRT format
        }
      }
    }
  }
}

# Create transparent cluster proxy credential
resource "tanzu-mission-control_credential" "transparent_proxy_cred" {
  name = "transparent_proxy_cred"

  meta {
    description = "transparent proxy credential"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "proxyType" : "transparent",
      "noProxyList" : "http://noproxy.com,http://something.com"
    }
  }

  spec {
    capability = "PROXY_CONFIG"
    provider   = "GENERIC_KEY_VALUE"
    data {
      key_value {
        data = {
          "proxyCABundle" = "-----BEGIN CERTIFICATE-----\n Encoded string for encryption of data\n ----END CERTIFICATE----" # chain of certificate is supported in CRT format
        }
      }
    }
  }
}
```

# Credential for Tanzu Mission Control provisioned AWS S3 storage used for data-protection

## Example Usage

```terraform
# Create credential for TMC provisioned AWS S3 storage used for data-protection
resource "tanzu-mission-control_credential" "tmc_provisioned_aws_s3_cred" {
  name = "aws_s3_cred"

  meta {
    description = "TMC provisioned AWS S3 storage"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "DATA_PROTECTION"
    provider   = "AWS_EC2"
    data {
      aws_credential {
        iam_role {
          arn = "arn:aws:iam::4987398738934:role/clusterlifecycle-test.tmc.cloud.vmware.com"
        }
      }
    }
  }
}
```

# Credential for Self provisioned AWS S3 or S3-compatible storage used for data-protection

## Example Usage

```terraform
# Create Self provisioned AWS S3 or S3-compatible credential
resource "tanzu-mission-control_credential" "aws_eks_cred" {
  name = "tf-aws-s3-self-test"

  meta {
    description = "Self provisioned AWS S3 or S3-compatible storage credential for data protection"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "DATA_PROTECTION"
    provider   = "GENERIC_S3"
    data {
      key_value {
        type = "OPAQUE_SECRET_TYPE"
        data = {
          "aws_access_key_id"     = "abcd="
          "aws_secret_access_key" = "xyz=="
        }
      }
    }
  }
  ready_wait_timeout = "5m" // Wait time for credential create operations to finish (default: 3m).
}
```

# Credential for self provisioned storage Azure blob for Data protection

## Example Usage

```terraform
# Create self provisioned storage Azure blob for Data protection.
resource "tanzu-mission-control_credential" "azure_ad_cred" {
  name = "tf-azure-ad-self-dp-test"

  meta {
    description = "Self provisioned storage Azure blob for Data protection"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "DATA_PROTECTION"
    provider   = "AZURE_AD"
    data {
      azure_credential {
        service_principal {
          subscription_id  = "some_subscription_id"
          tenant_id        = "some_tenant_id"
          resource_group   = "dp-backup-rg"
          client_id        = "some_client_id"
          client_secret    = "some_client_id"
          azure_cloud_name = "AzurePublicCloud"
        }
      }
    }
  }
  ready_wait_timeout = "5m" // Wait time for credential create operations to finish (default: 3m).
}
```

# Credential for Tanzu Observability

## Example Usage

```terraform
# Create credential for Tanzu Observability
resource "tanzu-mission-control_credential" "tanzu_observability_cred" {
  name = "tanzu_observability_cred"

  meta {
    description = "TMC integration: tanzu observability"
    labels = {
      "key1" : "value1",
    }
    annotations = {
      "wavefront.url" : "url pointing to your wavefront instance"
    }
  }

  spec {
    capability = "TANZU_OBSERVABILITY"
    provider   = "GENERIC_KEY_VALUE"
    data {
      key_value {
        data = {
          "wavefront.token" = "wavefront api token"
        }
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of this credential

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- `ready_wait_timeout` (String) Wait timeout duration until credential resource reaches VALID state. Accepted timeout duration values like 5s, 5m, or 1h, higher than zero.
- `spec` (Block List, Max: 1) Spec of credential resource (see [below for nested schema](#nestedblock--spec))

### Read-Only

- `id` (String) The ID of this resource.
- `status` (Map of String) Status of credential resource

<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Optional:

- `capability` (String) The Tanzu capability for which the credential shall be used. Value must be in list [DATA_PROTECTION TANZU_OBSERVABILITY TANZU_SERVICE_MESH PROXY_CONFIG MANAGED_K8S_PROVIDER IMAGE_REGISTRY]
- `data` (Block List, Max: 1) Holds credentials sensitive data (see [below for nested schema](#nestedblock--spec--data))
- `provider` (String) The Tanzu provider for which describes credential data type. Value must be in list [PROVIDER_UNSPECIFIED,AWS_EC2,GENERIC_S3,AZURE_AD,AWS_EKS,AZURE_AKS,GENERIC_KEY_VALUE]

<a id="nestedblock--spec--data"></a>
### Nested Schema for `spec.data`

Optional:

- `aws_credential` (Block List, Max: 1) AWS credential data type (see [below for nested schema](#nestedblock--spec--data--aws_credential))
- `azure_credential` (Block List, Max: 1) Azure credential (see [below for nested schema](#nestedblock--spec--data--azure_credential))
- `generic_credential` (String) Generic credential data type used to hold a blob of data represented as string
- `key_value` (Block List, Max: 1) Key Value credential (see [below for nested schema](#nestedblock--spec--data--key_value))

<a id="nestedblock--spec--data--aws_credential"></a>
### Nested Schema for `spec.data.aws_credential`

Optional:

- `account_id` (String) Account ID of the AWS credential
- `generic_credential` (String) Generic credential
- `iam_role` (Block List, Max: 1) AWS IAM role ARN and external ID (see [below for nested schema](#nestedblock--spec--data--aws_credential--iam_role))

<a id="nestedblock--spec--data--aws_credential--iam_role"></a>
### Nested Schema for `spec.data.aws_credential.iam_role`

Optional:

- `arn` (String) AWS IAM role ARN
- `ext_id` (String) An external ID used to assume an AWS IAM role



<a id="nestedblock--spec--data--azure_credential"></a>
### Nested Schema for `spec.data.azure_credential`

Optional:

- `service_principal` (Block List, Max: 1) Azure service principal (see [below for nested schema](#nestedblock--spec--data--azure_credential--service_principal))
- `service_principal_with_certificate` (Block List, Max: 1) Azure service principal (see [below for nested schema](#nestedblock--spec--data--azure_credential--service_principal_with_certificate))

<a id="nestedblock--spec--data--azure_credential--service_principal"></a>
### Nested Schema for `spec.data.azure_credential.service_principal`

Required:

- `client_id` (String) Client ID of the Service Principal
- `resource_group` (String) Resource Group name
- `subscription_id` (String) Subscription ID of the Azure credential
- `tenant_id` (String) Tenant ID of the Azure credential

Optional:

- `azure_cloud_name` (String) Azure Cloud name
- `client_secret` (String) Client Secret of the Service Principal


<a id="nestedblock--spec--data--azure_credential--service_principal_with_certificate"></a>
### Nested Schema for `spec.data.azure_credential.service_principal_with_certificate`

Required:

- `client_certificate` (String) Client certificate of the Service Principal
- `client_id` (String) Client ID of the Service Principal
- `subscription_id` (String) Subscription ID of the Azure credential
- `tenant_id` (String) Tenant ID of the Azure credential

Optional:

- `azure_cloud_name` (String) Azure Cloud name
- `managed_subscriptions` (List of String) IDs of the Azure Subscriptions that the Service Principal can manage



<a id="nestedblock--spec--data--key_value"></a>
### Nested Schema for `spec.data.key_value`

Optional:

- `data` (Map of String) Data secret data in the format of key-value pair
- `type` (String) Type of Secret data, usually mapped to k8s secret type. Supported types: [SECRET_TYPE_UNSPECIFIED,OPAQUE_SECRET_TYPE,DOCKERCONFIGJSON_SECRET_TYPE]
