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
)

func ResourceSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecurityPolicyCreate,
		ReadContext:   resourceSecurityPolicyRead,
		UpdateContext: resourceSecurityPolicyInPlaceUpdate,
		DeleteContext: resourceSecurityPolicyDelete,
		Schema:        securityPolicySchema,
		CustomizeDiff: customdiff.All(
			validateScope,
			validateInput,
			validateSpecLabelSelectorRequirement,
		),
	}
}

var securityPolicySchema = map[string]*schema.Schema{
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of the security policy",
		Required:    true,
		ForceNew:    true,
	},
	scopeKey:       scopeSchema,
	common.MetaKey: common.Meta,
	specKey:        specSchema,
}

func resourceSecurityPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	securityPolicyName, _ := d.Get(nameKey).(string)

	scopedFullnameData := constructScope(d, securityPolicyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to create Tanzu Mission Control cluster security policy entry; Scope full name is empty")
	}

	var UID string

	switch scopedFullnameData.scope {
	case clusterScope:
		if scopedFullnameData.fullnameCluster != nil {
			securityPolicyReq := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest{
				Policy: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicy{
					FullName: scopedFullnameData.fullnameCluster,
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
	case clusterGroupScope:
		if scopedFullnameData.fullnameClusterGroup != nil {
			securityPolicyReq := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest{
				Policy: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy{
					FullName: scopedFullnameData.fullnameClusterGroup,
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
	case organizationScope:
		if scopedFullnameData.fullnameOrganization != nil {
			securityPolicyReq := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{
				Policy: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
					FullName: scopedFullnameData.fullnameOrganization,
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
	case unknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
	}

	// always run
	d.SetId(UID)

	return resourceSecurityPolicyRead(ctx, d, m)
}

func retrieveSecurityPolicyUIDMetaAndSpecFromServer(config authctx.TanzuContext, scopedFullnameData *scopedFullname, d *schema.ResourceData, securityPolicyName string) (string, *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta, *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec, error) {
	var (
		UID  string
		meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
		spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec
	)

	switch scopedFullnameData.scope {
	case clusterScope:
		if scopedFullnameData.fullnameCluster != nil {
			resp, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceGet(scopedFullnameData.fullnameCluster)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName)
			}

			scopedFullnameData = &scopedFullname{
				scope:           clusterScope,
				fullnameCluster: resp.Policy.FullName,
			}

			fullName, name := flattenScope(scopedFullnameData)

			if err := d.Set(nameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(scopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case clusterGroupScope:
		if scopedFullnameData.fullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceGet(scopedFullnameData.fullnameClusterGroup)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName)
			}

			scopedFullnameData = &scopedFullname{
				scope:                clusterGroupScope,
				fullnameClusterGroup: resp.Policy.FullName,
			}

			fullName, name := flattenScope(scopedFullnameData)

			if err := d.Set(nameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(scopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case organizationScope:
		if scopedFullnameData.fullnameOrganization != nil {
			resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(scopedFullnameData.fullnameOrganization)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return "", nil, nil, nil
				}

				return "", nil, nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName)
			}

			scopedFullnameData = &scopedFullname{
				scope:                organizationScope,
				fullnameOrganization: resp.Policy.FullName,
			}

			fullName, name := flattenScope(scopedFullnameData)

			if err := d.Set(nameKey, name); err != nil {
				return "", nil, nil, err
			}

			if err := d.Set(scopeKey, fullName); err != nil {
				return "", nil, nil, err
			}

			UID = resp.Policy.Meta.UID
			meta = resp.Policy.Meta
			spec = resp.Policy.Spec
		}
	case unknownScope:
		return "", nil, nil, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
	}

	return UID, meta, spec, nil
}

func resourceSecurityPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	securityPolicyName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read security policy name")
	}

	scopedFullnameData := constructScope(d, securityPolicyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control cluster security policy entry; Scope full name is empty")
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

	if err := d.Set(specKey, flattenSpec(spec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSecurityPolicyInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	securityPolicyName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read security policy name")
	}

	scopedFullnameData := constructScope(d, securityPolicyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to update Tanzu Mission Control cluster security policy entry; Scope full name is empty")
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

	switch scopedFullnameData.scope {
	case clusterScope:
		if scopedFullnameData.fullnameCluster != nil {
			securityPolicyReq := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest{
				Policy: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicy{
					FullName: scopedFullnameData.fullnameCluster,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceUpdate(securityPolicyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName))
			}
		}
	case clusterGroupScope:
		if scopedFullnameData.fullnameClusterGroup != nil {
			securityPolicyReq := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest{
				Policy: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy{
					FullName: scopedFullnameData.fullnameClusterGroup,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceUpdate(securityPolicyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName))
			}
		}
	case organizationScope:
		if scopedFullnameData.fullnameOrganization != nil {
			securityPolicyReq := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{
				Policy: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
					FullName: scopedFullnameData.fullnameOrganization,
					Meta:     meta,
					Spec:     spec,
				},
			}

			_, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceUpdate(securityPolicyReq)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName))
			}
		}
	case unknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
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
	if !hasSpecChanged(d) {
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

	securityPolicyName, _ := d.Get(nameKey).(string)

	scopedFullnameData := constructScope(d, securityPolicyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to delete Tanzu Mission Control cluster security policy entry; Scope full name is empty")
	}

	switch scopedFullnameData.scope {
	case clusterScope:
		if scopedFullnameData.fullnameCluster != nil {
			err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceDelete(scopedFullnameData.fullnameCluster)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName))
			}
		}
	case clusterGroupScope:
		if scopedFullnameData.fullnameClusterGroup != nil {
			err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceDelete(scopedFullnameData.fullnameClusterGroup)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName))
			}
		}
	case organizationScope:
		if scopedFullnameData.fullnameOrganization != nil {
			err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceDelete(scopedFullnameData.fullnameOrganization)
			if err != nil && !clienterrors.IsNotFoundError(err) {
				return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName))
			}
		}
	case unknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	_ = schema.RemoveFromState(d, m)

	return diags
}
