// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policyoperations

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	policyworkspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindcustom "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom"
	policykindimage "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image"
	policykindmutation "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/mutation"
	policykindnetwork "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network"
	policykindquota "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota"
	policykindsecurity "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

func ResourcePolicyInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}, rn string) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	policyName, ok := d.Get(policy.NameKey).(string)

	if !ok {
		return diag.Errorf("unable to read %s policy name", rn)
	}

	scopedFullnameData := scope.ConstructScope(d, policyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to update Tanzu Mission Control %s policy entry; Scope full name is empty", rn)
	}

	_, meta, spec, err := RetrievePolicyUIDMetaAndSpecFromServer(config, scopedFullnameData, d, rn, policyName)
	if err != nil {
		return diag.FromErr(err)
	}

	var updateAvailable bool

	if updateCheckForMeta(d, meta) {
		updateAvailable = true
	}

	if updateCheckForSpec(d, spec, rn) {
		updateAvailable = true
	}

	if !updateAvailable {
		log.Printf("[INFO] %s policy update is not required", rn)

		return
	}

	switch scopedFullnameData.Scope {
	case scope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			policyReq := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest{
				Policy: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicy{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceUpdate(policyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster %s policy entry, name : %s", rn, policyName))
			}
		}
	case scope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			policyReq := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest{
				Policy: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceUpdate(policyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster group %s policy entry, name : %s", rn, policyName))
			}
		}
	case scope.WorkspaceScope:
		if scopedFullnameData.FullnameWorkspace != nil {
			policyReq := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest{
				Policy: &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicy{
					FullName: scopedFullnameData.FullnameWorkspace,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.WorkspacePolicyResourceService.ManageV1alpha1WorkspacePolicyResourceServiceUpdate(policyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control workspace %s policy entry, name : %s", rn, policyName))
			}
		}
	case scope.OrganizationScope:
		if scopedFullnameData.FullnameOrganization != nil {
			policyReq := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{
				Policy: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
					FullName: scopedFullnameData.FullnameOrganization,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceUpdate(policyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control organization %s policy entry, name : %s", rn, policyName))
			}
		}
	case scope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopeMap[rn], `, `))
	}

	log.Printf("[INFO] %s policy update successful", rn)

	return ResourcePolicyRead(ctx, d, m, rn)
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

	log.Printf("[INFO] updating policy meta data")

	return true
}

func updateCheckForSpec(d *schema.ResourceData, spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec, rn string) bool {
	if !hasSpecChanged(d) {
		return false
	}

	var policySpec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec

	switch rn {
	case policykindcustom.ResourceName:
		policySpec = policykindcustom.ConstructSpec(d)
	case policykindsecurity.ResourceName:
		policySpec = policykindsecurity.ConstructSpec(d)
	case policykindimage.ResourceName:
		policySpec = policykindimage.ConstructSpec(d)
	case policykindquota.ResourceName:
		policySpec = policykindquota.ConstructSpec(d)
	case policykindnetwork.ResourceName:
		policySpec = policykindnetwork.ConstructSpec(d)
	case policykindmutation.ResourceName:
		policySpec = policykindmutation.ConstructSpec(d)
	}

	spec.Input = policySpec.Input
	spec.NamespaceSelector = policySpec.NamespaceSelector
	spec.Recipe = policySpec.Recipe
	spec.RecipeVersion = policySpec.RecipeVersion

	log.Printf("[INFO] updating policy spec")

	return true
}

func hasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(policy.SpecKey, policy.InputKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(policy.SpecKey, policy.NamespaceSelectorKey)):
		updateRequired = true
	}

	return updateRequired
}
