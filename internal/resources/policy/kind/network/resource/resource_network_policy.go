/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package networkpolicyresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindnetwork "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

func ResourceNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: schema.CreateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindnetwork.ResourceName), policyoperations.WithOperationType(policyoperations.Create))),
		ReadContext:   schema.ReadContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindnetwork.ResourceName), policyoperations.WithOperationType(policyoperations.Read))),
		UpdateContext: schema.UpdateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindnetwork.ResourceName), policyoperations.WithOperationType(policyoperations.Update))),
		DeleteContext: schema.DeleteContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindnetwork.ResourceName), policyoperations.WithOperationType(policyoperations.Delete))),
		Schema:        networkPolicySchema,
		CustomizeDiff: customdiff.All(
			schema.CustomizeDiffFunc(scope.ValidateScope(policyoperations.ScopeMap[policykindnetwork.ResourceName])),
			policykindnetwork.ValidateInput,
			policy.ValidateSpecLabelSelectorRequirement,
		),
	}
}

var (
	ScopesAllowed = [...]string{scope.OrganizationKey, scope.WorkspaceKey}
	ScopeSchema   = scope.GetScopeSchema(
		scope.WithDescription(fmt.Sprintf("Scope for the network policy, having one of the valid scopes: %v.", strings.Join(ScopesAllowed[:], `, `))),
		scope.WithScopes(ScopesAllowed[:]))
)

var networkPolicySchema = map[string]*schema.Schema{
	policy.NameKey: {
		Type:        schema.TypeString,
		Description: "Name of the network policy",
		Required:    true,
		ForceNew:    true,
	},
	scope.ScopeKey: ScopeSchema,
	common.MetaKey: common.Meta,
	policy.SpecKey: policykindnetwork.SpecSchema,
}
