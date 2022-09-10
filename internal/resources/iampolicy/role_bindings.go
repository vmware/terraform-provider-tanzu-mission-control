/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

var roleBinding = &schema.Schema{
	Type:        schema.TypeList,
	Description: "List of role bindings associated with the policy",
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			roleKey: {
				Type:         schema.TypeString,
				Description:  "Role for this rolebinding: max length for a role is 126 characters.",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 126),
			},
			subjectsKey: {
				Type:        schema.TypeList,
				Description: "Subject for this rolebinding.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						subjectNameKey: {
							Type:         schema.TypeString,
							Description:  "Subject name: allow max characters for email - 320 characters.",
							Required:     true,
							ValidateFunc: validation.StringLenBetween(1, 320),
						},
						subjectKindKey: {
							Type:         schema.TypeString,
							Description:  "Subject type, having one of the subject types: USER or GROUP",
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"USER", "GROUP"}, false),
						},
					},
				},
			},
		},
	},
}

func constructRoleBindingListFromInterface(value interface{}) []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding {
	var rbl = make([]*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding, 0)

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return rbl
	}

	for _, raw := range data {
		rbData, _ := raw.(map[string]interface{})

		var rb = &iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
			Subjects: make([]*iammodel.VmwareTanzuCoreV1alpha1PolicySubject, 0),
		}

		if v, ok := rbData[roleKey]; ok {
			helper.SetPrimitiveValue(v, &rb.Role, roleKey)
		}

		if v, ok := rbData[subjectsKey]; ok {
			subjects, _ := v.([]interface{})
			for _, sb := range subjects {
				rb.Subjects = append(rb.Subjects, expandSubject(sb))
			}
		}

		rbl = append(rbl, rb)
	}

	return rbl
}

func constructRoleBindingList(d *schema.ResourceData) []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding {
	var rbl = make([]*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding, 0)

	value, ok := d.GetOk(roleBindingsKey)
	if !ok {
		return rbl
	}

	return constructRoleBindingListFromInterface(value)
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

func expandSubject(data interface{}) (subject *iammodel.VmwareTanzuCoreV1alpha1PolicySubject) {
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

func flattenRoleBindingList(rbl []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding) (data []interface{}) {
	if rbl == nil {
		return data
	}

	rbs := make([]interface{}, 0)

	for _, rb := range rbl {
		rbs = append(rbs, flattenRoleBinding(rb))
	}

	return rbs
}

func flattenRoleBinding(rb *iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding) (data interface{}) {
	if rb == nil {
		return data
	}

	flattenRoleBinding := make(map[string]interface{})

	flattenRoleBinding[roleKey] = rb.Role

	sbs := make([]interface{}, 0)

	for _, sb := range rb.Subjects {
		sbs = append(sbs, flattenSubject(sb))
	}

	flattenRoleBinding[subjectsKey] = sbs

	return flattenRoleBinding
}

func flattenSubject(subject *iammodel.VmwareTanzuCoreV1alpha1PolicySubject) (data interface{}) {
	flattenSubject := make(map[string]interface{})

	if subject == nil {
		return nil
	}

	flattenSubject[subjectNameKey] = subject.Name
	flattenSubject[subjectKindKey] = string(*subject.Kind)

	return flattenSubject
}

func validateRoleBindingSubjectDuplicate(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
	value, ok := diff.GetOk(roleBindingsKey)
	if !ok {
		return fmt.Errorf("role binding: %v is not valid: minimum one valid role binding block is required", value)
	}

	data, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("type of role binding block data: %v is not valid", value)
	}

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("role binding data: %v is not valid: minimum one valid role binding block is required", data)
	}

	errStrings := make([]string, 0)

	for _, raw := range data {
		rbData, ok := raw.(map[string]interface{})
		if !ok {
			errStrings = append(errStrings, fmt.Errorf("role binding: %v is not valid: minimum one valid role binding block is required", raw).Error())
		}

		v, ok := rbData[subjectsKey]
		if !ok {
			errStrings = append(errStrings, fmt.Errorf("role binding: %v for role: %v is not valid: minimum one valid subject block is required", rbData, rbData[roleKey]).Error())
		}

		subjects, ok := v.([]interface{})
		if !ok {
			errStrings = append(errStrings, fmt.Errorf("type of subject block data: %v for role: %v is not valid", v, rbData[roleKey]).Error())
		}

		subjectMap := make(map[string]bool)

		for _, sb := range subjects {
			subject := expandSubject(sb)
			if _, ok := subjectMap[subject.Name+string(*subject.Kind)]; !ok {
				subjectMap[subject.Name+string(*subject.Kind)] = true
				continue
			}

			errStrings = append(errStrings, fmt.Errorf("- subject with name: %v and kind: %v is not valid: it is a duplicate for role: %v", subject.Name, subject.Kind, rbData[roleKey]).Error())
		}
	}

	if len(errStrings) != 0 {
		return fmt.Errorf("error(s) in subjects: \n%s", strings.Join(errStrings, "\n"))
	}

	return nil
}
