// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackages

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	pakageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
	tanzupakageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackage"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	packagehelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package/scope"
	pack "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/packages/packageslist"
)

func DataSourceTanzuPackages() *schema.Resource {
	return &schema.Resource{
		Schema:      packagesSchema,
		ReadContext: resourcePackagesRead,
	}
}

var packagesSchema = map[string]*schema.Schema{
	namespaceKey: {
		Type:        schema.TypeString,
		Description: "Namespae of package.",
		Computed:    true,
	},
	metadataNameKey: {
		Type:        schema.TypeString,
		Description: "Metadata name of package.",
		Default:     "*",
		Optional:    true,
	},
	commonscope.ScopeKey: scope.ScopeSchema,
	pack.PackagesKey:     pack.PackageSchema,
}

func resourcePackagesRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	metadataName, ok := d.Get(metadataNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package metadata name")
	}

	scopedFullnameData, scopesFound := scope.ConstructScope(d)

	scopedFullnameData.FullnameCluster.MetadataName = metadataName

	if len(scopesFound) == 0 {
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v", strings.Join(scope.ScopeAllowed[:], `, `))
	} else if len(scopesFound) > 1 {
		return diag.Errorf("found scopes: %v are not valid: maximum one valid scope type block is allowed", strings.Join(scopesFound, `, `))
	}

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control package entry; Scope full name is empty")
	}

	globalNs, err := packagehelper.GetGlobalNamespace(config, &tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageSearchScope{
		ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
		ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
		ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	scopedFullnameData.FullnameCluster.NamespaceName = globalNs

	resp, err := config.TMCConnection.TanzupackageResourceService.ManageV1alpha1ClusterPackageResourceServiceList(&pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSearchScope{
		ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
		ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
		ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
		MetadataName:          metadataName,
		NamespaceName:         globalNs,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if len(resp.Packages) == 0 {
		return diag.Errorf("Tanzu Mission Control no package entry found with meatadata name : %s", metadataName)
	}

	d.SetId(resp.Packages[0].Meta.UID)

	flattenedSpec := pack.FlattenSpecForClusterScope(resp)

	if err := d.Set(metadataNameKey, metadataName); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(pack.PackagesKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(namespaceKey, globalNs); err != nil {
		return diag.FromErr(err)
	}

	fullName := scope.FlattenScope(scopedFullnameData)

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
