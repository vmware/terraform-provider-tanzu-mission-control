/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackage

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	tanzupackagemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
	tanzupakageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackage"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package/spec"
)

func DataSourceTanzuPackage() *schema.Resource {
	return &schema.Resource{
		Schema:      packageSchema,
		ReadContext: resourcePackageRead,
	}
}

var packageSchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the package. It represents version of the Package metadata",
		Required:    true,
	},
	namespaceKey: {
		Type:        schema.TypeString,
		Description: "Namespae of package.",
		Computed:    true,
	},
	metadataNameKey: {
		Type:        schema.TypeString,
		Description: "Metadata name of package.",
		Required:    true,
	},
	commonscope.ScopeKey: scope.ScopeSchema,
	common.MetaKey:       common.Meta,
	spec.SpecKey:         spec.SpecSchema,
}

func resourcePackageRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	packageName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package name")
	}

	metadataName, ok := d.Get(metadataNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read package metadata name")
	}

	scopedFullnameData, scopesFound := scope.ConstructScope(d)

	scopedFullnameData.FullnameCluster.MetadataName = metadataName
	scopedFullnameData.FullnameCluster.Name = packageName

	if len(scopesFound) == 0 {
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v", strings.Join(scope.ScopeAllowed[:], `, `))
	} else if len(scopesFound) > 1 {
		return diag.Errorf("found scopes: %v are not valid: maximum one valid scope type block is allowed", strings.Join(scopesFound, `, `))
	}

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control package entry; Scope full name is empty")
	}

	globalNs, err := GetGlobalNamespace(config, &tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageSearchScope{
		ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
		ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
		ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	scopedFullnameData.FullnameCluster.NamespaceName = globalNs

	UID, meta, atomicSpec, err := retrieveSourcesecretUIDMetaAndSpecFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			_ = schema.RemoveFromState(d, m)

			return diag.FromErr(errors.Wrapf(err, "Tanzu Mission Control package entry not found, name : %s", packageName))
		}

		return diag.FromErr(err)
	}

	// always run
	d.SetId(UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(meta)); err != nil {
		return diag.FromErr(err)
	}

	flattenedSpec := spec.FlattenSpecForClusterScope(atomicSpec)

	if err := d.Set(spec.SpecKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func retrieveSourcesecretUIDMetaAndSpecFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (
	string,
	*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta,
	*tanzupackagemodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec,
	error) {
	var (
		UID  string
		meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
		spec *tanzupackagemodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			resp, err := config.TMCConnection.TanzupackageResourceService.ManageV1alpha1ClusterPackageResourceServiceGet(scopedFullnameData.FullnameCluster)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					return "", nil, nil, err
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster package entry, name : %s", scopedFullnameData.FullnameCluster.Name)
			}

			scopedFullnameData = &scope.ScopedFullname{
				Scope:           commonscope.ClusterScope,
				FullnameCluster: resp.Package.FullName,
			}

			fullName := scope.FlattenScope(scopedFullnameData)

			if err := d.Set(nameKey, scopedFullnameData.FullnameCluster.Name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(metadataNameKey, scopedFullnameData.FullnameCluster.MetadataName); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(namespaceKey, scopedFullnameData.FullnameCluster.NamespaceName); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Package.Meta.UID
			meta = resp.Package.Meta
			spec = resp.Package.Spec
		}
	case commonscope.UnknownScope:
		return "", nil, nil, errors.Errorf("No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopeAllowed[:], `, `))
	}

	return UID, meta, spec, nil
}

func GetGlobalNamespace(config authctx.TanzuContext, searchscope *tanzupakageclustermodel.VmwareTanzuManageV1alpha1ClusterTanzupackageSearchScope) (string, error) {
	response, err := config.TMCConnection.ClusterTanzuPackageService.TanzuPackageResourceServiceList(searchscope)
	if err != nil {
		return "", err
	}

	if len(response.TanzuPackages) == 0 {
		return "", fmt.Errorf("cluster not found")
	}

	globalNs := (response.TanzuPackages[0]).Status.PackageRepositoryGlobalNamespace

	if globalNs == "" {
		return "", fmt.Errorf("global namespace not set for cluster")
	}

	return globalNs, nil
}
