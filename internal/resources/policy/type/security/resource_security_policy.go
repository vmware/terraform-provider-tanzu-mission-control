/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package security

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
)

func ResourceSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecurityPolicyCreate,
		ReadContext:   resourceSecurityPolicyRead,
		UpdateContext: resourceSecurityPolicyInPlaceUpdate,
		DeleteContext: resourceSecurityPolicyDelete,
		Schema:        securityPolicySchema,
		CustomizeDiff: customdiff.All(
			policy.ValidateScope,
			validateInput,
			policy.ValidateSpecLabelSelectorRequirement,
		),
	}
}

var securityPolicySchema = map[string]*schema.Schema{
	policy.NameKey: {
		Type:        schema.TypeString,
		Description: "Name of the security policy",
		Required:    true,
		ForceNew:    true,
	},
	policy.ScopeKey: policy.ScopeSchema,
	common.MetaKey:  common.Meta,
	policy.SpecKey:  specSchema,
}

func resourceSecurityPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	securityPolicyName, _ := d.Get(policy.NameKey).(string)

	scopedFullnameData := policy.ConstructScope(d, securityPolicyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control cluster security policy entry; Scope full name is empty")
	}

	var UID string

	switch scopedFullnameData.Scope {
	case policy.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			securityPolicyReq := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest{
				Policy: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicy{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     common.ConstructMeta(d),
					Spec:     constructSpec(d),
				},
			}

			securityPolicyResponse, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceCreate(securityPolicyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName))
			}

			UID = securityPolicyResponse.Policy.Meta.UID
		}
	case policy.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			securityPolicyReq := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest{
				Policy: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     common.ConstructMeta(d),
					Spec:     constructSpec(d),
				},
			}

			securityPolicyResponse, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceCreate(securityPolicyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName))
			}

			UID = securityPolicyResponse.Policy.Meta.UID
		}
	case policy.OrganizationScope:
		if scopedFullnameData.FullnameOrganization != nil {
			securityPolicyReq := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{
				Policy: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
					FullName: scopedFullnameData.FullnameOrganization,
					Meta:     common.ConstructMeta(d),
					Spec:     constructSpec(d),
				},
			}

			securityPolicyResponse, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceCreate(securityPolicyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName))
			}

			UID = securityPolicyResponse.Policy.Meta.UID
		}
	case policy.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policy.ScopesAllowed[:], `, `))
	}

	// always run
	d.SetId(UID)

	return resourceSecurityPolicyRead(ctx, d, m)
}

func retrieveSecurityPolicyUIDMetaAndSpecFromServer(config authctx.TanzuContext, scopedFullnameData *policy.ScopedFullname, d *schema.ResourceData, securityPolicyName string) (string, *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta, *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec, error) {
	var (
		UID  string
		meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
		spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec
	)

	switch scopedFullnameData.Scope {
	case policy.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			resp, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceGet(scopedFullnameData.FullnameCluster)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName)
			}

			scopedFullnameData = &policy.ScopedFullname{
				Scope:           policy.ClusterScope,
				FullnameCluster: resp.Policy.FullName,
			}

			fullName, name := policy.FlattenScope(scopedFullnameData)

			if err := d.Set(policy.NameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(policy.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case policy.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceGet(scopedFullnameData.FullnameClusterGroup)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName)
			}

			scopedFullnameData = &policy.ScopedFullname{
				Scope:                policy.ClusterGroupScope,
				FullnameClusterGroup: resp.Policy.FullName,
			}

			fullName, name := policy.FlattenScope(scopedFullnameData)

			if err := d.Set(policy.NameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(policy.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case policy.OrganizationScope:
		if scopedFullnameData.FullnameOrganization != nil {
			resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(scopedFullnameData.FullnameOrganization)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName)
			}

			scopedFullnameData = &policy.ScopedFullname{
				Scope:                policy.OrganizationScope,
				FullnameOrganization: resp.Policy.FullName,
			}

			fullName, name := policy.FlattenScope(scopedFullnameData)

			if err := d.Set(policy.NameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(policy.ScopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case policy.UnknownScope:
		return "", nil, nil, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policy.ScopesAllowed[:], `, `))
	}

	return UID, meta, spec, nil
}

func resourceSecurityPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	securityPolicyName, ok := d.Get(policy.NameKey).(string)
	if !ok {
		return diag.Errorf("unable to read security policy name")
	}

	scopedFullnameData := policy.ConstructScope(d, securityPolicyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control security policy entry; Scope full name is empty")
	}

	UID, meta, spec, err := retrieveSecurityPolicyUIDMetaAndSpecFromServer(config, scopedFullnameData, d, securityPolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(policy.SpecKey, flattenSpec(spec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSecurityPolicyInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	securityPolicyName, ok := d.Get(policy.NameKey).(string)
	if !ok {
		return diag.Errorf("unable to read security policy name")
	}

	scopedFullnameData := policy.ConstructScope(d, securityPolicyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to update Tanzu Mission Control security policy entry; Scope full name is empty")
	}

	_, meta, spec, err := retrieveSecurityPolicyUIDMetaAndSpecFromServer(config, scopedFullnameData, d, securityPolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	var updateAvailable bool

	if updateCheckForMeta(d, meta) {
		updateAvailable = true
	}

	if updateCheckForSpec(d, spec) {
		updateAvailable = true
	}

	if !updateAvailable {
		log.Printf("[INFO] security policy update is not required")

		return
	}

	switch scopedFullnameData.Scope {
	case policy.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			securityPolicyReq := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest{
				Policy: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicy{
					FullName: scopedFullnameData.FullnameCluster,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceUpdate(securityPolicyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName))
			}
		}
	case policy.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			securityPolicyReq := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest{
				Policy: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy{
					FullName: scopedFullnameData.FullnameClusterGroup,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceUpdate(securityPolicyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName))
			}
		}
	case policy.OrganizationScope:
		if scopedFullnameData.FullnameOrganization != nil {
			securityPolicyReq := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{
				Policy: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
					FullName: scopedFullnameData.FullnameOrganization,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceUpdate(securityPolicyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName))
			}
		}
	case policy.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policy.ScopesAllowed[:], `, `))
	}

	log.Printf("[INFO] security policy update successful")

	return resourceSecurityPolicyRead(ctx, d, m)
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

	log.Printf("[INFO] updating security policy meta data")

	return true
}

func updateCheckForSpec(d *schema.ResourceData, spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec) bool {
	if !policy.HasSpecChanged(d) {
		return false
	}

	policySpec := constructSpec(d)

	spec.Input = policySpec.Input
	spec.NamespaceSelector = policySpec.NamespaceSelector
	spec.Recipe = policySpec.Recipe
	spec.RecipeVersion = policySpec.RecipeVersion

	log.Printf("[INFO] updating security policy spec")

	return true
}

func resourceSecurityPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	securityPolicyName, _ := d.Get(policy.NameKey).(string)

	scopedFullnameData := policy.ConstructScope(d, securityPolicyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to delete Tanzu Mission Control cluster security policy entry; Scope full name is empty")
	}

	switch scopedFullnameData.Scope {
	case policy.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceDelete(scopedFullnameData.FullnameCluster)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName))
			}
		}
	case policy.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceDelete(scopedFullnameData.FullnameClusterGroup)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName))
			}
		}
	case policy.OrganizationScope:
		if scopedFullnameData.FullnameOrganization != nil {
			err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceDelete(scopedFullnameData.FullnameOrganization)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName))
			}
		}
	case policy.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policy.ScopesAllowed[:], `, `))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
