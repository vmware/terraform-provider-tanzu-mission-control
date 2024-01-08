/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package permissiontemplatetests

import (
	"fmt"

	permissiontemplateres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/permissiontemplate"
)

const (
	dataProtectionPermissionTemplateDataSourceName = "data_protection_permissions"
	eksPermissionTemplateDataSourceName            = "eks_permissions"
)

var (
	dataProtectionPermissionTemplateDataSourceFullName = fmt.Sprintf("data.%s.%s", permissiontemplateres.ResourceName, dataProtectionPermissionTemplateDataSourceName) //nolint:unused
	eksPermissionTemplateDataSourceFullName            = fmt.Sprintf("data.%s.%s", permissiontemplateres.ResourceName, eksPermissionTemplateDataSourceName)            //nolint:unused
)

func GetDataProtectionPermissionTemplateConfig() string {
	return fmt.Sprintf(`	
		data "%s" "%s" {
		  credentials_name = "data-protection-tf-test"
		  tanzu_capability = "DATA_PROTECTION"
		  tanzu_provider   = "AWS_EC2"
		}
		`,
		permissiontemplateres.ResourceName,
		dataProtectionPermissionTemplateDataSourceName,
	)
}

func GetEKSPermissionTemplateConfig() string {
	return fmt.Sprintf(`	
		data "%s" "%s" {
		  credentials_name = "eks-tf-test"
		  tanzu_capability = "MANAGED_K8S_PROVIDER"
		  tanzu_provider   = "AWS_EKS"
		}
		`,
		permissiontemplateres.ResourceName,
		eksPermissionTemplateDataSourceName,
	)
}
