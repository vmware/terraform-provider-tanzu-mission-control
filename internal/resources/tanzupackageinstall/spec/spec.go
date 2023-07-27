/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	packageinstallmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
)

var (
	SpecSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "spec for package install.",
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				PackageRefKey: PackageRefKeySpec,
				RoleBindingScopeKey: {
					Type:        schema.TypeString,
					Description: "Role binding scope for service account which will be used by Package Install.",
					Optional:    true,
					Default:     fmt.Sprint(packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER),
					ValidateFunc: validation.StringInSlice([]string{
						fmt.Sprint(packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeUNSPECIFIED),
						fmt.Sprint(packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER),
						fmt.Sprint(packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeNAMESPACE),
					}, false),
				},
				InlineValuesKey: {
					Type:        schema.TypeMap,
					Description: "Inline values to configure the Package Install.",
					Sensitive:   true,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}

	PackageRefKeySpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Reference to the Package which will be installed.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				PackageMetadataNameKey: {
					Type:        schema.TypeString,
					Description: "Name of the Package Metadata.",
					Required:    true,
					ForceNew:    true,
				},
				VersionSelectionKey: versionSelectionSpec,
			},
		},
	}

	versionSelectionSpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Version Selection of the Package.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				ConstraintsKey: {
					Type:        schema.TypeString,
					Description: "Constraints to select Package. Example: constraints: 'v1.2.3', constraints: '<v1.4.0' etc.",
					Required:    true,
				},
			},
		},
	}
)

func HasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(SpecKey, PackageRefKey, PackageMetadataNameKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, PackageRefKey, VersionSelectionKey, ConstraintsKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, RoleBindingScopeKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, InlineValuesKey)):
		updateRequired = true
	}

	return updateRequired
}
