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
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindcustom "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom"
	policykindimage "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image"
	policykindmutation "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/mutation"
	policykindnetwork "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network"
	policykindquota "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota"
	policykindsecurity "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

var ScopeMap = map[string][]string{
	policykindcustom.ResourceName:   {scope.ClusterKey, scope.ClusterGroupKey, scope.OrganizationKey},
	policykindimage.ResourceName:    {scope.WorkspaceKey, scope.OrganizationKey},
	policykindquota.ResourceName:    {scope.ClusterKey, scope.ClusterGroupKey, scope.OrganizationKey},
	policykindsecurity.ResourceName: {scope.ClusterKey, scope.ClusterGroupKey, scope.OrganizationKey},
	policykindnetwork.ResourceName:  {scope.WorkspaceKey, scope.OrganizationKey},
	policykindmutation.ResourceName: {scope.ClusterKey, scope.ClusterGroupKey, scope.OrganizationKey},
}

// nolint: gocognit
func RetrievePolicyUIDMetaAndSpecFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData, policyName, rn string) (string, *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta, *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec, error) {
	var (
		UID  string
		meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
		spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec
	)
	// nolint: dupl
	switch scopedFullnameData.Scope {
	case scope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			resp, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceGet(scopedFullnameData.FullnameCluster)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster %s policy entry, name : %s", rn, policyName)
			}

			scopedFullnameData = &scope.ScopedFullname{
				Scope:           scope.ClusterScope,
				FullnameCluster: resp.Policy.FullName,
			}

			fullName, name := scope.FlattenScope(scopedFullnameData, ScopeMap[rn])

			if err := d.Set(policy.NameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(scope.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case scope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceGet(scopedFullnameData.FullnameClusterGroup)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group %s policy entry, name : %s", rn, policyName)
			}

			scopedFullnameData = &scope.ScopedFullname{
				Scope:                scope.ClusterGroupScope,
				FullnameClusterGroup: resp.Policy.FullName,
			}

			fullName, name := scope.FlattenScope(scopedFullnameData, ScopeMap[rn])

			if err := d.Set(policy.NameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(scope.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case scope.WorkspaceScope:
		if scopedFullnameData.FullnameWorkspace != nil {
			resp, err := config.TMCConnection.WorkspacePolicyResourceService.ManageV1alpha1WorkspacePolicyResourceServiceGet(scopedFullnameData.FullnameWorkspace)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control workspace %s policy entry, name : %s", rn, policyName)
			}

			scopedFullnameData = &scope.ScopedFullname{
				Scope:             scope.WorkspaceScope,
				FullnameWorkspace: resp.Policy.FullName,
			}

			fullName, name := scope.FlattenScope(scopedFullnameData, ScopeMap[rn])

			if err := d.Set(policy.NameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(scope.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case scope.OrganizationScope:
		if scopedFullnameData.FullnameOrganization != nil {
			resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(scopedFullnameData.FullnameOrganization)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control organization %s policy entry, name : %s", rn, policyName)
			}

			scopedFullnameData = &scope.ScopedFullname{
				Scope:                scope.OrganizationScope,
				FullnameOrganization: resp.Policy.FullName,
			}

			fullName, name := scope.FlattenScope(scopedFullnameData, ScopeMap[rn])

			if err := d.Set(policy.NameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(scope.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case scope.UnknownScope:
		return "", nil, nil, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopeMap[rn], `, `))
	}

	return UID, meta, spec, nil
}

type OperationOption func(*OperationConfig)

type Operation string

const (
	Create Operation = "CREATE"
	Read   Operation = "READ"
	Update Operation = "UPDATE"
	Delete Operation = "DELETE"
)

type OperationConfig struct {
	ResourceName  string
	OperationType Operation
}

func WithResourceName(rn string) OperationOption {
	return func(config *OperationConfig) {
		config.ResourceName = rn
	}
}

func WithOperationType(t Operation) OperationOption {
	return func(config *OperationConfig) {
		config.OperationType = t
	}
}

type ResourceOperationType func(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics)

func ResourceOperation(opts ...OperationOption) ResourceOperationType {
	cfg := &OperationConfig{}

	for _, o := range opts {
		o(cfg)
	}

	return func(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
		switch cfg.OperationType {
		case Create:
			diags = ResourcePolicyCreate(ctx, d, m, cfg.ResourceName)
		case Read:
			diags = ResourcePolicyRead(ctx, d, m, cfg.ResourceName)
		case Update:
			diags = ResourcePolicyInPlaceUpdate(ctx, d, m, cfg.ResourceName)
		case Delete:
			diags = ResourcePolicyDelete(ctx, d, m, cfg.ResourceName)
		}

		return diags
	}
}
