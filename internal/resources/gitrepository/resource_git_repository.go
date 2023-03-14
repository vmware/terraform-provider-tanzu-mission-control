/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepository

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/status"
)

func ResourceGitRepository() *schema.Resource {
	return &schema.Resource{
		Schema:        gitRepositorySchema,
		CreateContext: resourceGitRepositoryCreate,
		ReadContext:   dataSourceGitRepositoryRead,
		UpdateContext: resourceGitRepositoryInPlaceUpdate,
		DeleteContext: resourceGitRepositoryDelete,
		CustomizeDiff: schema.CustomizeDiffFunc(commonscope.ValidateScope([]string{commonscope.ClusterKey, commonscope.ClusterGroupKey})),
	}
}

var gitRepositorySchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the Repository.",
		Required:    true,
		ForceNew:    true,
	},
	namespaceNameKey: {
		Type:        schema.TypeString,
		Description: "Name of Namespace.",
		Required:    true,
		ForceNew:    true,
	},
	orgIDKey: {
		Type:        schema.TypeString,
		Description: "ID of Organization.",
		Optional:    true,
		ForceNew:    true,
	},
	commonscope.ScopeKey: scope.ScopeSchema,
	common.MetaKey:       common.Meta,
	spec.SpecKey:         spec.SpecSchema,
	statusKey:            status.StatusSchema,
}

func resourceGitRepositoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	gitRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read git repository name")
	}

	gitRepositoryNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("unable to read git repository namespace name")
	}

	gitRepositoryOrgID, _ := d.Get(orgIDKey).(string)
	scopedFullnameData := scope.ConstructScope(d, gitRepositoryName, gitRepositoryNamespaceName, gitRepositoryOrgID)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control git repository entry; Scope full name is empty")
	}

	var (
		UID  string
		meta = common.ConstructMeta(d)
	)

	err := enableContinuousDelivery(&config, scopedFullnameData, meta)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control git repository entry, name : %s", gitRepositoryName))
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			gitRepositoryReq := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest{
				GitRepository: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
					Spec:     spec.ConstructSpecForClusterScope(d),
				},
			}

			gitRepositoryResponse, err := config.TMCConnection.ClusterGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceCreate(gitRepositoryReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster git repository entry, name : %s", gitRepositoryName))
			}

			UID = gitRepositoryResponse.GitRepository.Meta.UID
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			gitRepositoryReq := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest{
				GitRepository: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
					Spec:     spec.ConstructSpecForClusterGroupScope(d),
				},
			}

			gitRepositoryResponse, err := config.TMCConnection.ClusterGroupGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceCreate(gitRepositoryReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster group git repository entry, name : %s", gitRepositoryName))
			}

			UID = gitRepositoryResponse.GitRepository.Meta.UID
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	// always run
	d.SetId(UID)

	return dataSourceGitRepositoryRead(ctx, d, m)
}

func resourceGitRepositoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	gitRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read git repository name")
	}

	gitRepositoryNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("unable to read git repository namespace name")
	}

	gitRepositoryOrgID, _ := d.Get(orgIDKey).(string)
	scopedFullnameData := scope.ConstructScope(d, gitRepositoryName, gitRepositoryNamespaceName, gitRepositoryOrgID)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to delete Tanzu Mission Control git repository entry; Scope full name is empty")
	}

	err := enableContinuousDelivery(&config, scopedFullnameData, common.ConstructMeta(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control git repository entry, name : %s", gitRepositoryName))
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			err := config.TMCConnection.ClusterGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceDelete(scopedFullnameData.FullnameCluster)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster git repository entry, name : %s", gitRepositoryName))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			err := config.TMCConnection.ClusterGroupGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceDelete(scopedFullnameData.FullnameClusterGroup)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group git repository entry, name : %s", gitRepositoryName))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func resourceGitRepositoryInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	gitRepositoryName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read git repository name")
	}

	gitRepositoryNamespaceName, ok := d.Get(namespaceNameKey).(string)
	if !ok {
		return diag.Errorf("unable to read git repository namespace name")
	}

	gitRepositoryOrgID, _ := d.Get(orgIDKey).(string)
	scopedFullnameData := scope.ConstructScope(d, gitRepositoryName, gitRepositoryNamespaceName, gitRepositoryOrgID)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to update Tanzu Mission Control git repository entry; Scope full name is empty")
	}

	err := enableContinuousDelivery(&config, scopedFullnameData, common.ConstructMeta(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control git repository entry, name : %s", gitRepositoryName))
	}

	// nolint: dogsled
	_, meta, atomicSpec, _, _, err := retrieveGitRepositoryUIDMetaAndSpecFromServer(config, scopedFullnameData, d, gitRepositoryName)
	if err != nil {
		return diag.FromErr(err)
	}

	var updateAvailable bool

	if updateCheckForMeta(d, meta) {
		updateAvailable = true
	}

	if updateCheckForSpec(d, atomicSpec, scopedFullnameData.Scope) {
		updateAvailable = true
	}

	if !updateAvailable {
		log.Printf("[INFO] git repository update is not required")
		return
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			gitRepositoryReq := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest{
				GitRepository: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
					Spec:     atomicSpec,
				},
			}

			_, err := config.TMCConnection.ClusterGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceUpdate(gitRepositoryReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster git repository entry, name : %s", gitRepositoryName))
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			gitRepositoryReq := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest{
				GitRepository: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
					Spec: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec{
						AtomicSpec: atomicSpec,
					},
				},
			}

			_, err := config.TMCConnection.ClusterGroupGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceUpdate(gitRepositoryReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster group git repository entry, name : %s", gitRepositoryName))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	log.Printf("[INFO] git repository update successful")

	return dataSourceGitRepositoryRead(ctx, d, m)
}

func updateCheckForMeta(d *schema.ResourceData, meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta) bool {
	if !common.HasMetaChanged(d) {
		return false
	}

	objectMeta := common.ConstructMeta(d)

	if value, ok := meta.Labels[common.CreatorLabelKey]; ok {
		objectMeta.Labels[common.CreatorLabelKey] = value
	}

	meta.Labels = objectMeta.Labels
	meta.Description = objectMeta.Description

	log.Printf("[INFO] updating git repository meta data")

	return true
}

func updateCheckForSpec(d *schema.ResourceData, atomicSpec *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec, scope commonscope.Scope) bool {
	if !spec.HasSpecChanged(d) {
		return false
	}

	var gitRepositorySpec *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec

	switch scope {
	case commonscope.ClusterScope:
		gitRepositorySpec = spec.ConstructSpecForClusterScope(d)
	case commonscope.ClusterGroupScope:
		clusterGroupScopeSpec := spec.ConstructSpecForClusterGroupScope(d)
		gitRepositorySpec = clusterGroupScopeSpec.AtomicSpec
	}

	atomicSpec.GitImplementation = gitRepositorySpec.GitImplementation
	atomicSpec.Interval = gitRepositorySpec.Interval
	atomicSpec.Ref = gitRepositorySpec.Ref
	atomicSpec.SecretRef = gitRepositorySpec.SecretRef
	atomicSpec.URL = gitRepositorySpec.URL

	log.Printf("[INFO] updating git repository spec")

	return true
}
