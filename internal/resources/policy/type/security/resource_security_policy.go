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

	scopedFullname := constructScope(d, securityPolicyName)

	if scopedFullname != nil {
		var UID string

		switch scopedFullname.scope {
		case clusterScope:
			if scopedFullname.fullnameCluster != nil {
				securityPolicyReq := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest{
					Policy: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicy{
						FullName: scopedFullname.fullnameCluster,
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
			if scopedFullname.fullnameClusterGroup != nil {
				securityPolicyReq := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest{
					Policy: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy{
						FullName: scopedFullname.fullnameClusterGroup,
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
			if scopedFullname.fullnameOrganization != nil {
				securityPolicyReq := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{
					Policy: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
						FullName: scopedFullname.fullnameOrganization,
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
			log.Fatalf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
		}

		// always run
		d.SetId(UID)
	}

	return resourceSecurityPolicyRead(ctx, d, m)
}

// nolint: gocognit
func resourceSecurityPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	securityPolicyName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read security policy name")
	}

	scopedFullnameData := constructScope(d, securityPolicyName)

	// nolint: nestif
	if scopedFullnameData != nil {
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
						return
					}

					return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName))
				}

				scopedFullnameData = &scopedFullname{
					scope:           clusterScope,
					fullnameCluster: resp.Policy.FullName,
				}

				fullName, name := flattenScope(scopedFullnameData)

				if err := d.Set(nameKey, name); err != nil {
					return diag.FromErr(err)
				}

				if err := d.Set(scopeKey, fullName); err != nil {
					return diag.FromErr(err)
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
						return
					}

					return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName))
				}

				scopedFullnameData = &scopedFullname{
					scope:                clusterGroupScope,
					fullnameClusterGroup: resp.Policy.FullName,
				}

				fullName, name := flattenScope(scopedFullnameData)

				if err := d.Set(nameKey, name); err != nil {
					return diag.FromErr(err)
				}

				if err := d.Set(scopeKey, fullName); err != nil {
					return diag.FromErr(err)
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
						return
					}

					return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName))
				}

				scopedFullnameData = &scopedFullname{
					scope:                organizationScope,
					fullnameOrganization: resp.Policy.FullName,
				}

				fullName, name := flattenScope(scopedFullnameData)

				if err := d.Set(nameKey, name); err != nil {
					return diag.FromErr(err)
				}

				if err := d.Set(scopeKey, fullName); err != nil {
					return diag.FromErr(err)
				}

				UID = resp.Policy.Meta.UID
				meta = resp.Policy.Meta
				spec = resp.Policy.Spec
			}
		case unknownScope:
			log.Fatalf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
		}

		// always run
		d.SetId(UID)

		if err := d.Set(common.MetaKey, common.FlattenMeta(meta)); err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set(specKey, flattenSpec(spec)); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

// nolint: gocognit
func resourceSecurityPolicyInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var updateAvailable bool

	securityPolicyName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read security policy name")
	}

	scopedFullname := constructScope(d, securityPolicyName)

	// nolint: nestif
	if scopedFullname != nil {
		var (
			meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
			spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec
		)

		switch scopedFullname.scope {
		case clusterScope:
			if scopedFullname.fullnameCluster != nil {
				// Get call to initialise the security policy struct
				getResp, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceGet(scopedFullname.fullnameCluster)
				if err != nil {
					return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName))
				}

				meta = getResp.Policy.Meta
				spec = getResp.Policy.Spec
			}
		case clusterGroupScope:
			if scopedFullname.fullnameClusterGroup != nil {
				// Get call to initialise the security policy struct
				getResp, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceGet(scopedFullname.fullnameClusterGroup)
				if err != nil {
					return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName))
				}

				meta = getResp.Policy.Meta
				spec = getResp.Policy.Spec
			}
		case organizationScope:
			if scopedFullname.fullnameOrganization != nil {
				// Get call to initialise the security policy struct
				getResp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(scopedFullname.fullnameOrganization)
				if err != nil {
					return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName))
				}

				meta = getResp.Policy.Meta
				spec = getResp.Policy.Spec
			}
		case unknownScope:
			log.Fatalf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
		}

		if updateCheckForMeta(d, meta) {
			updateAvailable = true
		}

		if updateCheckForSpec(d, spec) {
			updateAvailable = true
		}

		if updateAvailable {
			switch scopedFullname.scope {
			case clusterScope:
				if scopedFullname.fullnameCluster != nil {
					securityPolicyReq := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest{
						Policy: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyPolicy{
							FullName: scopedFullname.fullnameCluster,
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
				if scopedFullname.fullnameClusterGroup != nil {
					securityPolicyReq := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest{
						Policy: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy{
							FullName: scopedFullname.fullnameClusterGroup,
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
				if scopedFullname.fullnameOrganization != nil {
					securityPolicyReq := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest{
						Policy: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyPolicy{
							FullName: scopedFullname.fullnameOrganization,
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
				log.Fatalf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
			}

			log.Printf("[INFO] security policy update successful")
		}
	}

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

	scopedFullname := constructScope(d, securityPolicyName)

	if scopedFullname != nil {
		switch scopedFullname.scope {
		case clusterScope:
			if scopedFullname.fullnameCluster != nil {
				err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceDelete(scopedFullname.fullnameCluster)
				if err != nil && !clienterrors.IsNotFoundError(err) {
					return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster security policy entry, name : %s", securityPolicyName))
				}
			}
		case clusterGroupScope:
			if scopedFullname.fullnameClusterGroup != nil {
				err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceDelete(scopedFullname.fullnameClusterGroup)
				if err != nil && !clienterrors.IsNotFoundError(err) {
					return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster group security policy entry, name : %s", securityPolicyName))
				}
			}
		case organizationScope:
			if scopedFullname.fullnameOrganization != nil {
				err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceDelete(scopedFullname.fullnameOrganization)
				if err != nil && !clienterrors.IsNotFoundError(err) {
					return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control organization security policy entry, name : %s", securityPolicyName))
				}
			}
		case unknownScope:
			log.Fatalf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
