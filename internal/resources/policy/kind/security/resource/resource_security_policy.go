/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package securitypolicyresource

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindsecurity "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

func ResourceSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: schema.CreateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindsecurity.ResourceName), policyoperations.WithOperationType(policyoperations.Create))),
		ReadContext:   schema.ReadContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindsecurity.ResourceName), policyoperations.WithOperationType(policyoperations.Read))),
		UpdateContext: schema.UpdateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindsecurity.ResourceName), policyoperations.WithOperationType(policyoperations.Update))),
		DeleteContext: schema.DeleteContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindsecurity.ResourceName), policyoperations.WithOperationType(policyoperations.Delete))),
		Schema:        securityPolicySchema,
		CustomizeDiff: customdiff.All(
			schema.CustomizeDiffFunc(scope.ValidateScope(policyoperations.ScopeMap[policykindsecurity.ResourceName])),
			policykindsecurity.ValidateInput,
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
	scope.ScopeKey: scope.ScopeSchema,
	common.MetaKey: common.Meta,
	policy.SpecKey: policykindsecurity.SpecSchema,
}
