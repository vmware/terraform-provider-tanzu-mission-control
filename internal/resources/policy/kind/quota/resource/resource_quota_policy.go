// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package quotapolicyresource

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindquota "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

func ResourceQuotaPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: schema.CreateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindquota.ResourceName), policyoperations.WithOperationType(policyoperations.Create))),
		ReadContext:   schema.ReadContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindquota.ResourceName), policyoperations.WithOperationType(policyoperations.Read))),
		UpdateContext: schema.UpdateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindquota.ResourceName), policyoperations.WithOperationType(policyoperations.Update))),
		DeleteContext: schema.DeleteContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindquota.ResourceName), policyoperations.WithOperationType(policyoperations.Delete))),
		Schema:        quotaPolicySchema,
		CustomizeDiff: customdiff.All(
			schema.CustomizeDiffFunc(scope.ValidateScope(policyoperations.ScopeMap[policykindquota.ResourceName])),
			policykindquota.ValidateInput,
			policy.ValidateSpecLabelSelectorRequirement,
		),
	}
}

var quotaPolicySchema = map[string]*schema.Schema{
	policy.NameKey: {
		Type:        schema.TypeString,
		Description: "Name of the namespace quota policy",
		Required:    true,
		ForceNew:    true,
	},
	scope.ScopeKey: scope.ScopeSchema,
	common.MetaKey: common.Meta,
	policy.SpecKey: policykindquota.SpecSchema,
}
