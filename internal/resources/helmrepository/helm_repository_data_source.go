/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrepository

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	repositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrepository"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	helmscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrepository/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrepository/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrepository/status"
)

const (
	ResourceName = "tanzu-mission-control_helm_repository"

	nameKey          = "name"
	namespaceNameKey = "namespace_name"
	statusKey        = "status"
)

type dataFromServer struct {
	UID                string
	meta               *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
	clusterScopeStatus *repositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryStatus
	atomicSpec         *repositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySpec
}

func DataSourceHelmRepository() *schema.Resource {
	return &schema.Resource{
		Schema:      helmSchema,
		ReadContext: dataSourceHelmRepositoryRead,
	}
}

var helmSchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the helm repository.",
		Optional:    true,
	},
	namespaceNameKey: {
		Type:        schema.TypeString,
		Description: "Name of Namespace.",
		Optional:    true,
		Default:     "*",
	},
	commonscope.ScopeKey: helmscope.ScopeSchema,
	common.MetaKey:       common.Meta,
	spec.SpecKey:         spec.SpecSchema,
	statusKey:            status.StatusSchema,
}

func dataSourceHelmRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	namespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read repository namespace name")
	}

	scopedFullnameData := helmscope.ConstructScope(d)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control helm repository entry; Scope full name is empty")
	}

	scopedFullnameData.FullnameCluster.NamespaceName = namespaceName

	helmRepoDataFromServer, err := retrieveHelmRepositoryDataFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return
		}

		return diag.FromErr(err)
	}

	// always run
	d.SetId(helmRepoDataFromServer.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(helmRepoDataFromServer.meta)); err != nil {
		return diag.FromErr(err)
	}

	var (
		flattenedSpec   []interface{}
		flattenedStatus interface{}
	)

	flattenedSpec = spec.FlattenSpecForClusterScope(helmRepoDataFromServer.atomicSpec)
	flattenedStatus = status.FlattenStatusForClusterScope(helmRepoDataFromServer.clusterScopeStatus)

	if err := d.Set(spec.SpecKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(statusKey, flattenedStatus); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func retrieveHelmRepositoryDataFromServer(config authctx.TanzuContext, scopedFullnameData *helmscope.ScopedFullname, d *schema.ResourceData) (*dataFromServer, error) {
	var helmRepoDataFromServer = &dataFromServer{}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			resp, err := config.TMCConnection.ClusterHelmRepositoryResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdHelmRepositoryResourceServiceList(&repositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySearchScope{
				ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
				ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
				ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
				NamespaceName:         scopedFullnameData.FullnameCluster.NamespaceName,
			})
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return helmRepoDataFromServer, err
				}

				return helmRepoDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control Helm Repository entry for cluster, name : %s", scopedFullnameData.FullnameCluster.ClusterName)
			}

			if len(resp.Repositories) == 0 {
				return helmRepoDataFromServer, errors.Errorf("No entry found for Tanzu Mission Control Helm Repository for cluster, name : %s", scopedFullnameData.FullnameCluster.ClusterName)
			}

			repo := resp.Repositories[0]

			scopedFullnameData.FullnameCluster = repo.FullName
			helmRepoDataFromServer.UID = repo.Meta.UID
			helmRepoDataFromServer.meta = repo.Meta
			helmRepoDataFromServer.atomicSpec = repo.Spec
			helmRepoDataFromServer.clusterScopeStatus = repo.Status
		}
	case commonscope.UnknownScope:
		return helmRepoDataFromServer, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(helmscope.ScopesAllowed[:], `, `))
	}

	fullName, name, namespace := helmscope.FlattenScope(scopedFullnameData)

	if err := d.Set(nameKey, name); err != nil {
		return helmRepoDataFromServer, err
	}

	if err := d.Set(namespaceNameKey, namespace); err != nil {
		return helmRepoDataFromServer, err
	}

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return helmRepoDataFromServer, err
	}

	return helmRepoDataFromServer, nil
}
