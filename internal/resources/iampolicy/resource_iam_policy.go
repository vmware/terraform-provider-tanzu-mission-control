/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

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
	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	clusteriammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy/cluster"
	clustergroupiammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy/clustergroup"
	namespaceiammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy/namespace"
	organizationiammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy/organization"
	workspaceiammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy/workspace"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type (
	contextMethodKey struct{}
)

func ResourceIAMPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceIAMPolicyRead,
		CreateContext: resourceIAMPolicyCreate,
		UpdateContext: resourceIAMPolicyInPlaceUpdate,
		DeleteContext: resourceIAMPolicyDelete,
		Schema:        iamPolicySchema,
		CustomizeDiff: customdiff.All(
			validateScope,
			validateRoleBindingSubjectDuplicate,
		),
	}
}

var iamPolicySchema = map[string]*schema.Schema{
	scopeKey:        scopeSchema,
	common.MetaKey:  common.Meta,
	roleBindingsKey: roleBinding,
}

func resourceIAMPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config, _ := m.(authctx.TanzuContext)

	var UID string

	scopedFullname := constructScope(d)
	blData := constructBindingDeltaList(d, iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD.Pointer())

	if scopedFullname == nil {
		return diag.Errorf("unable to create Role Binding; Scope full name is empty")
	}

	switch scopedFullname.scope {
	case organizationScope:
		if scopedFullname.fullnameOrganization != nil {
			iamRequest := &organizationiammodel.VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest{
				FullName:         scopedFullname.fullnameOrganization,
				BindingDeltaList: blData,
			}

			iamResponse, err := config.TMCConnection.OrganizationIAMResourceService.ManageV1alpha1OrganizationIAMPolicyPatch(iamRequest)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for organization"))
			}

			UID = iamResponse.Policy.Meta.UID
		}
	case clusterGroupScope:
		if scopedFullname.fullnameClusterGroup != nil {
			iamRequest := &clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest{
				FullName:         scopedFullname.fullnameClusterGroup,
				BindingDeltaList: blData,
			}

			iamResponse, err := config.TMCConnection.ClusterGroupIAMResourceService.ManageV1alpha1ClusterGroupIAMPolicyPatch(iamRequest)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for cluster group"))
			}

			UID = iamResponse.Policy.Meta.UID
		}
	case clusterScope:
		if scopedFullname.fullnameCluster != nil {
			iamRequest := &clusteriammodel.VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest{
				FullName:         scopedFullname.fullnameCluster,
				BindingDeltaList: blData,
			}

			iamResponse, err := config.TMCConnection.ClusterIAMResourceService.ManageV1alpha1ClusterIAMPolicyPatch(iamRequest)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for cluster"))
			}

			UID = iamResponse.Policy.Meta.UID
		}
	case workspaceScope:
		if scopedFullname.fullnameWorkspace != nil {
			iamRequest := &workspaceiammodel.VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest{
				FullName:         scopedFullname.fullnameWorkspace,
				BindingDeltaList: blData,
			}

			iamResponse, err := config.TMCConnection.WorkspaceIAMResourceService.ManageV1alpha1WorkspaceIAMPolicyPatch(iamRequest)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for workspace"))
			}

			UID = iamResponse.Policy.Meta.UID
		}
	case namespaceScope:
		if scopedFullname.fullnameNamespace != nil {
			iamRequest := &namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest{
				FullName:         scopedFullname.fullnameNamespace,
				BindingDeltaList: blData,
			}

			iamResponse, err := config.TMCConnection.NamespaceIAMResourceService.ManageV1alpha1ClusterNamespaceIAMPolicyPatch(iamRequest)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for namespace"))
			}

			UID = iamResponse.Policy.Meta.UID
		}
	case unknownScope:
		return diag.Errorf("unable to create Role Binding; No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
	}

	d.SetId(UID)

	return append(
		diags,
		resourceIAMPolicyRead(context.WithValue(ctx, contextMethodKey{}, createKey), d, m)...,
	)
}

func retrieveRoleBindingListFromServer(config authctx.TanzuContext, scopedFullnameData *scopedFullname) (policyList []*iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy, err error) {
	switch scopedFullnameData.scope {
	case organizationScope:
		if scopedFullnameData.fullnameOrganization != nil {
			resp, err := config.TMCConnection.OrganizationIAMResourceService.ManageV1alpha1OrganizationIAMPolicyGet(scopedFullnameData.fullnameOrganization)
			if err != nil || resp == nil {
				return nil, errors.Wrapf(err, "unable to get Role Bindings for organization")
			}

			policyList = resp.PolicyList
		}
	case clusterGroupScope:
		if scopedFullnameData.fullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupIAMResourceService.ManageV1alpha1ClusterGroupIAMPolicyGet(scopedFullnameData.fullnameClusterGroup)
			if err != nil || resp == nil {
				return nil, errors.Wrapf(err, "unable to get Role Bindings for cluster group")
			}

			policyList = resp.PolicyList
		}
	case clusterScope:
		if scopedFullnameData.fullnameCluster != nil {
			resp, err := config.TMCConnection.ClusterIAMResourceService.ManageV1alpha1ClusterIAMPolicyGet(scopedFullnameData.fullnameCluster)
			if err != nil || resp == nil {
				return nil, errors.Wrapf(err, "unable to get Role Bindings for cluster")
			}

			policyList = resp.PolicyList
		}
	case workspaceScope:
		if scopedFullnameData.fullnameWorkspace != nil {
			resp, err := config.TMCConnection.WorkspaceIAMResourceService.ManageV1alpha1WorkspaceIAMPolicyGet(scopedFullnameData.fullnameWorkspace)
			if err != nil || resp == nil {
				return nil, errors.Wrapf(err, "unable to get Role Bindings for workspace")
			}

			policyList = resp.PolicyList
		}
	case namespaceScope:
		if scopedFullnameData.fullnameNamespace != nil {
			resp, err := config.TMCConnection.NamespaceIAMResourceService.ManageV1alpha1ClusterNamespaceIAMPolicyGet(scopedFullnameData.fullnameNamespace)
			if err != nil || resp == nil {
				return nil, errors.Wrapf(err, "unable to get Role Bindings for namespace")
			}

			policyList = resp.PolicyList
		}
	case unknownScope:
		return nil, errors.Wrapf(err, "unable to get Role Binding; No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
	}

	return policyList, err
}

// resourceIAMPolicyRead returns the intersection between binding list in terraform state and TMC server.
func resourceIAMPolicyRead(_ context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config, _ := m.(authctx.TanzuContext)

	var (
		meta         *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
		rbServerList []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding
		calRBList    []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding
	)

	scopedFullnameData := constructScope(d)
	rbStateList := constructRoleBindingList(d)

	if scopedFullnameData == nil {
		return diag.Errorf("unable to get Role Bindings; Scope full name is empty")
	}

	policyList, err := retrieveRoleBindingListFromServer(config, scopedFullnameData)
	if err != nil {
		return diag.FromErr(err)
	}
	// when iam policy resource is empty: no role bindings are existing, this is equivalent to not found condition.
	if len(policyList) == 0 {
		d.SetId("")
		return
	}

	d.SetId(d.State().ID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(meta)); err != nil {
		return diag.FromErr(err)
	}

	// iterate over the policy lists: state and server, and store the intersection of the two
	for _, policy := range policyList {
		if policy.Meta.UID == d.State().ID {
			rbServerList = append(rbServerList, policy.RoleBindings...)
		}
	}
	// nested iteration for preserving order of role binding lists.
	for _, stateRB := range rbStateList {
		for _, serverRB := range rbServerList {
			if stateRB.Role == serverRB.Role {
				calRBList = append(
					calRBList,
					&iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
						Role:     serverRB.Role,
						Subjects: getIntersectionOfSubs(stateRB.Subjects, serverRB.Subjects),
					},
				)
			}
		}
	}

	if err := d.Set(roleBindingsKey, flattenRoleBindingList(calRBList)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func getIntersectionOfSubs(
	state, server []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject,
) []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject {
	var newList []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject
	// nested iteration for preserving order of subjects.
	for _, each := range state {
		for _, sub := range server {
			if sub.Name == each.Name && *sub.Kind == *each.Kind {
				newList = append(newList, each)
			}
		}
	}

	return newList
}

func constructRBOpForUpdate(rbl []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding, subjectIntersect *map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType, setAction iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType) {
	var (
		flipAction    iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType
		toBeFlippedTo iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType
	)

	if setAction == iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD {
		flipAction = iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeDELETE
		toBeFlippedTo = iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeOPTYPEUNSPECIFIED
	}

	for _, rb := range rbl {
		for _, sub := range rb.Subjects {
			key := strings.Join([]string{rb.Role, sub.Name, string(*sub.Kind)}, roleSubjectDelimiter)
			if action, ok := (*subjectIntersect)[key]; ok && action == flipAction {
				(*subjectIntersect)[key] = toBeFlippedTo

				continue
			}

			(*subjectIntersect)[key] = setAction
		}
	}
}

func getRoleBindingListToUpdate(subjectIntersect *map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType) (blData []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta) {
	for k, v := range *subjectIntersect {
		if v != iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeOPTYPEUNSPECIFIED {
			values := strings.Split(k, roleSubjectDelimiter)

			blData = append(blData, &iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta{
				Role: values[0],
				Subject: &iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
					Name: values[1],
					Kind: (*iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKind)(&values[2]),
				},
				Op: v.Pointer(),
			})
		}
	}

	return blData
}

func resourceIAMPolicyInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config, _ := m.(authctx.TanzuContext)

	var (
		UID              string
		blData           []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta
		subjectIntersect = make(map[string]iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType)
	)

	scopedFullname := constructScope(d)
	if scopedFullname == nil {
		return diag.Errorf("unable to update Role Binding; Scope full name is empty")
	}

	policyList, err := retrieveRoleBindingListFromServer(config, scopedFullname)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, policy := range policyList {
		if policy.Meta.UID == d.State().ID {
			constructRBOpForUpdate(policy.RoleBindings, &subjectIntersect, iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeOPTYPEUNSPECIFIED)
		}
	}

	oldValues, newValues := d.GetChange(roleBindingsKey)
	oldRBL := constructRoleBindingListFromInterface(oldValues)
	newRBL := constructRoleBindingListFromInterface(newValues)

	constructRBOpForUpdate(oldRBL, &subjectIntersect, iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeDELETE)
	constructRBOpForUpdate(newRBL, &subjectIntersect, iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD)

	blData = getRoleBindingListToUpdate(&subjectIntersect)

	switch scopedFullname.scope {
	case organizationScope:
		iamRequest := &organizationiammodel.VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest{
			FullName:         scopedFullname.fullnameOrganization,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.OrganizationIAMResourceService.ManageV1alpha1OrganizationIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to update Role Binding for organization"))
		}

		UID = iamResponse.Policy.Meta.UID
	case clusterGroupScope:
		iamRequest := &clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest{
			FullName:         scopedFullname.fullnameClusterGroup,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.ClusterGroupIAMResourceService.ManageV1alpha1ClusterGroupIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to update Role Binding for cluster group"))
		}

		UID = iamResponse.Policy.Meta.UID
	case clusterScope:
		iamRequest := &clusteriammodel.VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest{
			FullName:         scopedFullname.fullnameCluster,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.ClusterIAMResourceService.ManageV1alpha1ClusterIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to update Role Binding for cluster"))
		}

		UID = iamResponse.Policy.Meta.UID
	case workspaceScope:
		iamRequest := &workspaceiammodel.VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest{
			FullName:         scopedFullname.fullnameWorkspace,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.WorkspaceIAMResourceService.ManageV1alpha1WorkspaceIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to update Role Binding for workspace"))
		}

		UID = iamResponse.Policy.Meta.UID
	case namespaceScope:
		iamRequest := &namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest{
			FullName:         scopedFullname.fullnameNamespace,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.NamespaceIAMResourceService.ManageV1alpha1ClusterNamespaceIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to update Role Binding for namespace"))
		}

		UID = iamResponse.Policy.Meta.UID
	case unknownScope:
		return diag.Errorf("unable to update Role Binding; No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
	}

	d.SetId(UID)
	log.Printf("[INFO] role binding update successful")

	return append(
		diags,
		resourceIAMPolicyRead(context.WithValue(ctx, contextMethodKey{}, updateKey), d, m)...,
	)
}

func resourceIAMPolicyDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config, _ := m.(authctx.TanzuContext)

	var diags diag.Diagnostics

	scopedFullname := constructScope(d)
	if scopedFullname == nil {
		return diag.Errorf("unable to delete Role Binding; Scope full name is empty")
	}

	blData := constructBindingDeltaList(d, iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeDELETE.Pointer())

	switch scopedFullname.scope {
	case organizationScope:
		iamRequest := &organizationiammodel.VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest{
			FullName:         scopedFullname.fullnameOrganization,
			BindingDeltaList: blData,
		}

		_, err := config.TMCConnection.OrganizationIAMResourceService.ManageV1alpha1OrganizationIAMPolicyPatch(iamRequest)
		if err != nil && !clienterrors.IsNotFoundError(err) {
			return diag.FromErr(errors.Wrapf(err, "unable to delete Role Binding for organization"))
		}
	case clusterGroupScope:
		iamRequest := &clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest{
			FullName:         scopedFullname.fullnameClusterGroup,
			BindingDeltaList: blData,
		}

		_, err := config.TMCConnection.ClusterGroupIAMResourceService.ManageV1alpha1ClusterGroupIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to delete Role Binding for cluster group"))
		}
	case clusterScope:
		iamRequest := &clusteriammodel.VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest{
			FullName:         scopedFullname.fullnameCluster,
			BindingDeltaList: blData,
		}

		_, err := config.TMCConnection.ClusterIAMResourceService.ManageV1alpha1ClusterIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to delete Role Binding for cluster"))
		}
	case workspaceScope:
		iamRequest := &workspaceiammodel.VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest{
			FullName:         scopedFullname.fullnameWorkspace,
			BindingDeltaList: blData,
		}

		_, err := config.TMCConnection.WorkspaceIAMResourceService.ManageV1alpha1WorkspaceIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to delete Role Binding for workspace"))
		}
	case namespaceScope:
		iamRequest := &namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest{
			FullName:         scopedFullname.fullnameNamespace,
			BindingDeltaList: blData,
		}

		_, err := config.TMCConnection.NamespaceIAMResourceService.ManageV1alpha1ClusterNamespaceIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to delete Role Binding for namespace"))
		}
	case unknownScope:
		return diag.Errorf("unable to delete Role Binding; No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
