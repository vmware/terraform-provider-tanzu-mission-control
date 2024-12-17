// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package gitrepository

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/status"
)

type dataFromServer struct {
	UID                     string
	meta                    *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
	atomicSpec              *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec
	clusterScopeStatus      *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryStatus
	clusterGroupScopeStatus *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryStatus
}

func DataSourceGitRepository() *schema.Resource {
	return &schema.Resource{
		Schema: getGitRepositorySchema(true),
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceGitRepositoryRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
	}
}

func dataSourceGitRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	gitRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read git repository name")
	}

	gitRepositoryNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read git repository namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, gitRepositoryName, gitRepositoryNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control git repository entry; Scope full name is empty")
	}

	gitRepositoryDataFromServer, err := retrieveGitRepositoryDataFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return diags
		}

		return diag.FromErr(err)
	}

	// always run
	d.SetId(gitRepositoryDataFromServer.UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(gitRepositoryDataFromServer.meta)); err != nil {
		return diag.FromErr(err)
	}

	var (
		flattenedSpec   []interface{}
		flattenedStatus interface{}
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		flattenedSpec = spec.FlattenSpecForClusterScope(gitRepositoryDataFromServer.atomicSpec)
		flattenedStatus = status.FlattenStatusForClusterScope(gitRepositoryDataFromServer.clusterScopeStatus)
	case commonscope.ClusterGroupScope:
		clusterGroupScopeSpec := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec{
			AtomicSpec: gitRepositoryDataFromServer.atomicSpec,
		}
		flattenedSpec = spec.FlattenSpecForClusterGroupScope(clusterGroupScopeSpec)
		flattenedStatus = status.FlattenStatusForClusterGroupScope(gitRepositoryDataFromServer.clusterGroupScopeStatus)
	}

	if err := d.Set(spec.SpecKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(statusKey, flattenedStatus); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func retrieveGitRepositoryDataFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (*dataFromServer, error) {
	var gitRepositoryDataFromServer = &dataFromServer{}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			resp, err := config.TMCConnection.ClusterGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceGet(scopedFullnameData.FullnameCluster)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return gitRepositoryDataFromServer, err
				}

				return gitRepositoryDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster git repository entry, name : %s", scopedFullnameData.FullnameCluster.Name)
			}

			scopedFullnameData.FullnameCluster = resp.GitRepository.FullName
			gitRepositoryDataFromServer.UID = resp.GitRepository.Meta.UID
			gitRepositoryDataFromServer.meta = resp.GitRepository.Meta
			gitRepositoryDataFromServer.atomicSpec = resp.GitRepository.Spec
			gitRepositoryDataFromServer.clusterScopeStatus = resp.GitRepository.Status
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceGet(scopedFullnameData.FullnameClusterGroup)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return gitRepositoryDataFromServer, err
				}

				return gitRepositoryDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group git repository entry, name : %s", scopedFullnameData.FullnameClusterGroup.Name)
			}

			scopedFullnameData.FullnameClusterGroup = resp.GitRepository.FullName
			gitRepositoryDataFromServer.UID = resp.GitRepository.Meta.UID
			gitRepositoryDataFromServer.meta = resp.GitRepository.Meta
			gitRepositoryDataFromServer.atomicSpec = resp.GitRepository.Spec.AtomicSpec
			gitRepositoryDataFromServer.clusterGroupScopeStatus = resp.GitRepository.Status
		}
	case commonscope.UnknownScope:
		return gitRepositoryDataFromServer, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	fullName, name, namespace := scope.FlattenScope(scopedFullnameData)

	if err := d.Set(nameKey, name); err != nil {
		return gitRepositoryDataFromServer, err
	}

	if err := d.Set(namespaceNameKey, namespace); err != nil {
		return gitRepositoryDataFromServer, err
	}

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return gitRepositoryDataFromServer, err
	}

	return gitRepositoryDataFromServer, nil
}
