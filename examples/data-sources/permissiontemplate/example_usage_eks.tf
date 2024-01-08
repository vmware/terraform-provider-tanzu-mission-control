locals {
  credentials_name = "test-permission-template-eks-tf-43"
  tanzu_capability = "MANAGED_K8S_PROVIDER"
  tanzu_provider   = "AWS_EKS"

  stack_message  = split("\n", aws_cloudformation_stack.crendetials_permission_template.outputs.Message)
  permission_arn = element(local.stack_message, length(local.stack_message) - 1)
}


data "tanzu-mission-control_permission_template" "eks_permissions" {
  credentials_name = local.credentials_name
  tanzu_capability = local.tanzu_capability
  tanzu_provider   = local.tanzu_provider
}


resource "aws_cloudformation_stack" "crendetials_permission_template" {
  name          = local.credentials_name
  parameters    = data.tanzu-mission-control_permission_template.eks_permissions.template_values != null ? data.tanzu-mission-control_permission_template.eks_permissions.template_values : {}
  template_body = base64decode(data.tanzu-mission-control_permission_template.eks_permissions.template)
  capabilities  = ["CAPABILITY_NAMED_IAM"]
}

resource "tanzu-mission-control_credential" "aws_eks_cred" {
  name = local.credentials_name

  spec {
    capability = local.tanzu_capability
    provider   = local.tanzu_provider

    data {
      aws_credential {
        iam_role {
          arn = local.permission_arn
        }
      }
    }
  }
}
