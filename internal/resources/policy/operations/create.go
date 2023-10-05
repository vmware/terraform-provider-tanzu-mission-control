/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyoperations

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
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

func ResourcePolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}, rn string) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	policyName, _ := d.Get(policy.NameKey).(string)
	scopedFullnameData := scope.ConstructScope(d, policyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control cluster %s policy entry; Scope full name is empty", rn)
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

	var UID string

	switch scopedFullnameData.Scope {
	case scope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			policyReq := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest{
				Policy: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicy{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     common.ConstructMeta(d),
					Spec:     policySpec,
				},
			}

			policyResponse, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceCreate(policyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster %s policy entry, name : %s", rn, policyName))
			}

			UID = policyResponse.Policy.Meta.UID
		}
	case scope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			policyReq := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest{
				Policy: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     common.ConstructMeta(d),
					Spec:     policySpec,
				},
			}

			policyResponse, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceCreate(policyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster group %s policy entry, name : %s", rn, policyName))
			}

			UID = policyResponse.Policy.Meta.UID
		}
	case scope.WorkspaceScope:
		if scopedFullnameData.FullnameWorkspace != nil {
			policyReq := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest{
				Policy: &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicy{
					FullName: scopedFullnameData.FullnameWorkspace,
					Meta:     common.ConstructMeta(d),
					Spec:     policySpec,
				},
			}

			policyResponse, err := config.TMCConnection.WorkspacePolicyResourceService.ManageV1alpha1WorkspacePolicyResourceServiceCreate(policyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control workspace %s policy entry, name : %s", rn, policyName))
			}

			UID = policyResponse.Policy.Meta.UID
		}
	case scope.OrganizationScope:
		if scopedFullnameData.FullnameOrganization != nil {
			policyReq := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{
				Policy: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
					FullName: scopedFullnameData.FullnameOrganization,
					Meta:     common.ConstructMeta(d),
					Spec:     policySpec,
				},
			}

			policyResponse, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceCreate(policyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control organization %s policy entry, name : %s", rn, policyName))
			}

			UID = policyResponse.Policy.Meta.UID
		}
	case scope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopeMap[rn], `, `))
	}

	// always run
	d.SetId(UID)

	return ResourcePolicyRead(ctx, d, m, rn)
}
