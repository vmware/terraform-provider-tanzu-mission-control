/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicyresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindcustom "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

func ResourceCustomPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: schema.CreateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindcustom.ResourceName), policyoperations.WithOperationType(policyoperations.Create))),
		ReadContext:   schema.ReadContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindcustom.ResourceName), policyoperations.WithOperationType(policyoperations.Read))),
		UpdateContext: schema.UpdateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindcustom.ResourceName), policyoperations.WithOperationType(policyoperations.Update))),
		DeleteContext: schema.DeleteContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindcustom.ResourceName), policyoperations.WithOperationType(policyoperations.Delete))),
		Schema:        customPolicySchema,
		CustomizeDiff: customdiff.All(
			schema.CustomizeDiffFunc(scope.ValidateScope(policyoperations.ScopeMap[policykindcustom.ResourceName])),
			policykindcustom.ValidateInput,
			policy.ValidateSpecLabelSelectorRequirement,
		),
	}
}

var (
	ScopesAllowed = [...]string{scope.ClusterKey, scope.ClusterGroupKey, scope.OrganizationKey}
	ScopeSchema   = scope.GetScopeSchema(
		scope.WithDescription(fmt.Sprintf("Scope for the custom policy, having one of the valid scopes: %v.", strings.Join(ScopesAllowed[:], `, `))),
		scope.WithScopes(ScopesAllowed[:]))
)

var customPolicySchema = map[string]*schema.Schema{
	policy.NameKey: {
		Type:        schema.TypeString,
		Description: "Name of the custom policy",
		Required:    true,
		ForceNew:    true,
	},
	scope.ScopeKey: ScopeSchema,
	common.MetaKey: common.Meta,
	policy.SpecKey: policykindcustom.SpecSchema,
}
