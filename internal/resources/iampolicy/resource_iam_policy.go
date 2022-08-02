/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster"
	clustergroupmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	iammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	clusteriammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy/cluster"
	clustergroupiammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy/clustergroup"
	namespaceiammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy/namespace"
	organizationiammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy/organization"
	workspaceiammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy/workspace"
	namespacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/namespace"
	organizationmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/organization"
	workspacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

func ResourceIAMPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIAMPolicyRead,
		CreateContext: resourceIAMPolicyCreate,
		//UpdateContext: resourceIAMPolicyUpdate,
		//DeleteContext: resourceIAMPolicyDelete,
		Schema: iamPolicySchema,
	}
}

var iamPolicySchema = map[string]*schema.Schema{
	scopeKey:       scopeSchema,
	roleBindingKey: roleBinding,
}

var scopeSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Scope of the resource on which the rolebinding has to be added",
	Required:    true,
	ForceNew:    true,
	MinItems:    1,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			orgKey:          orgScope,
			clusterGroupKey: clusterGroupScope,
			clusterKey:      clusterScope,
			workspaceKey:    workspaceScope,
			namespaceKey:    namespaceScope,
		},
	},
}

var orgScope = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Full name of the organization",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			orgIDKey: {
				Type:        schema.TypeString,
				Description: "ID of Organization",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

var clusterGroupScope = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Full name of the cluster group",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterGroupNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster group",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

var clusterScope = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Full name of the cluster",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			managementClusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the management cluster",
				Default:     "attached",
				Optional:    true,
				ForceNew:    true,
			},
			provisionerNameKey: {
				Type:        schema.TypeString,
				Description: "Provisioner of the cluster",
				Default:     "attached",
				Optional:    true,
				ForceNew:    true,
			},
			clusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

var workspaceScope = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Full name of the workspace",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			workspaceNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the workspace",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

var namespaceScope = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Full name of the namespace",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			namespaceNameKey: {
				Type:        schema.TypeString,
				Description: "Name of Namespace",
				Required:    true,
				ForceNew:    true,
			},
			managementClusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of ManagementCluster",
				Default:     "attached",
				Optional:    true,
				ForceNew:    true,
			},
			provisionerNameKey: {
				Type:        schema.TypeString,
				Description: "Name of Provisioner",
				Default:     "attached",
				Optional:    true,
				ForceNew:    true,
			},
			clusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of Cluster",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

var roleBinding = &schema.Schema{
	Type:        schema.TypeList,
	Description: "List of role bindings associated with the policy",
	Required:    true,
	ForceNew:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			roleKey: {
				Type:        schema.TypeString,
				Description: "Role for this rolebinding -max length for role is 126",
				Required:    true,
				ForceNew:    true,
			},
			subjectsKey: {
				Type:        schema.TypeList,
				Description: "Subject of rolebinding",
				Required:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						subjectNameKey: {
							Type:        schema.TypeString,
							Description: "Subject name - allow max characters for email - 320",
							Required:    true,
						},
						subjectKindKey: {
							Type:        schema.TypeString,
							Description: "Subject type",
							Required:    true,
						},
					},
				},
			},
		},
	},
}

func constructRoleBindingList(d *schema.ResourceData) []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding {
	var rbl = make([]*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding, 0)

	value, ok := d.GetOk(roleBindingKey)
	if !ok {
		return rbl
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return rbl
	}

	for _, raw := range data {
		rbData := raw.(map[string]interface{})

		var rb = &iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
			Subjects: make([]*iammodel.VmwareTanzuCoreV1alpha1PolicySubject, 0),
		}

		if v, ok := rbData[roleKey]; ok {
			helper.SetPrimitiveValue(v, &rb.Role, roleKey)
		}

		if v, ok := rbData[subjectsKey]; ok {
			subjects, _ := v.([]interface{})
			for _, sb := range subjects {
				rb.Subjects = append(rb.Subjects, expandSubjects(sb))
			}
		}

		rbl = append(rbl, rb)
	}

	return rbl
}

func constructBindingDeltaList(d *schema.ResourceData, op *iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType) []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta {
	rbList := constructRoleBindingList(d)

	var deltaList []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta

	for _, rb := range rbList {
		for _, sub := range rb.Subjects {
			deltaList = append(
				deltaList,
				&iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta{
					Role:    rb.Role,
					Subject: sub,
					Op:      op,
				},
			)
		}
	}

	return deltaList
}

func expandSubjects(data interface{}) (subject *iammodel.VmwareTanzuCoreV1alpha1PolicySubject) {
	lookUpSubjects, _ := data.(map[string]interface{})
	subject = &iammodel.VmwareTanzuCoreV1alpha1PolicySubject{}

	if v, ok := lookUpSubjects[subjectNameKey]; ok {
		helper.SetPrimitiveValue(v, &subject.Name, subjectNameKey)
	}

	if v, ok := lookUpSubjects[subjectKindKey]; ok {
		subject.Kind = iammodel.NewVmwareTanzuCoreV1alpha1PolicySubjectKind(
			iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKind(v.(string)),
		)
	}

	return subject
}

func flattenRoleBinding(rb *iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding) (data []interface{}) {
	if rb == nil {
		return data
	}

	flattenRoleBinding := make(map[string]interface{})

	flattenRoleBinding[roleKey] = rb.Role

	sbs := make([]interface{}, 0)

	for _, sb := range rb.Subjects {
		sbs = append(sbs, flattenSubjects(sb))
	}

	flattenRoleBinding[subjectsKey] = sbs

	return []interface{}{flattenRoleBinding}
}

func flattenSubjects(subject *iammodel.VmwareTanzuCoreV1alpha1PolicySubject) (data interface{}) {
	flattenSubject := make(map[string]interface{})

	if subject == nil {
		return nil
	}

	flattenSubject[subjectNameKey] = subject.Name
	flattenSubject[subjectKindKey] = string(*subject.Kind)

	return flattenSubject
}

func resourceIAMPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	s, err := getScopeType(d)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "unable to create Role Binding"))
	}

	fnData := constructScopeFullName(s, d)
	if fnData == nil {
		return diag.FromErr(errors.Wrap(err, "unable to create Role Binding, no full name found"))
	}

	blData := constructBindingDeltaList(d, iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD.Pointer())

	switch s {
	case organization:
		fn, _ := fnData.(*organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName)

		iamRequest := &organizationiammodel.VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest{
			FullName:         fn,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.OrganizationIAMResourceService.ManageV1alpha1OrganizationIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for organization"))
		}

		d.SetId(iamResponse.Policy.Meta.UID)
	case clusterGroup:
		fn, _ := fnData.(*clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName)

		iamRequest := &clustergroupiammodel.VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest{
			FullName:         fn,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.ClusterGroupIAMResourceService.ManageV1alpha1ClusterGroupIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for cluster group"))
		}

		d.SetId(iamResponse.Policy.Meta.UID)
	case cluster:
		fn, _ := fnData.(*clustermodel.VmwareTanzuManageV1alpha1ClusterFullName)

		iamRequest := &clusteriammodel.VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest{
			FullName:         fn,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.ClusterIAMResourceService.ManageV1alpha1ClusterIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for cluster"))
		}

		d.SetId(iamResponse.Policy.Meta.UID)
	case workspace:
		fn, _ := fnData.(*workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName)

		iamRequest := &workspaceiammodel.VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest{
			FullName:         fn,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.WorkspaceIAMResourceService.ManageV1alpha1WorkspaceIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for workspace"))
		}

		d.SetId(iamResponse.Policy.Meta.UID)
	case namespace:
		fn, _ := fnData.(*namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName)

		iamRequest := &namespaceiammodel.VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest{
			FullName:         fn,
			BindingDeltaList: blData,
		}

		iamResponse, err := config.TMCConnection.NamespaceIAMResourceService.ManageV1alpha1ClusterNamespaceIAMPolicyPatch(iamRequest)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "unable to create Role Binding for namespace"))
		}

		d.SetId(iamResponse.Policy.Meta.UID)
	default:
		return diag.FromErr(errors.Wrap(err, "unable to create Role Binding, invalid scope defined"))
	}

	return dataSourceIAMPolicyRead(ctx, d, m)
}

//func resourceIAMPolicyDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
//	config := m.(authctx.TanzuContext)
//
//	var diags diag.Diagnostics
//	s, err := getScopeType(d)
//	if err != nil {
//		return diag.FromErr(errors.Wrap(err, "unable to delete Role Binding"))
//	}
//
//	switch s {
//	case organization:
//		err := config.TMCConnection.OrganizationIAMResourceService.ManageV1alpha1NamespaceResourceServiceDelete(constructFullname(d))
//		if err != nil && !clienterrors.IsNotFoundError(err) {
//			return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control namespace entry, name : %s", namespaceName))
//		}
//	}
//
//	// d.SetId("") is automatically called assuming delete returns no errors, but
//	// it is added here for explicitness.
//	d.SetId("")
//
//	return diags
//}
