/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package imagepolicyresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindimage "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

func ResourceImagePolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: schema.CreateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindimage.ResourceName), policyoperations.WithOperationType(policyoperations.Create))),
		ReadContext:   schema.ReadContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindimage.ResourceName), policyoperations.WithOperationType(policyoperations.Read))),
		UpdateContext: schema.UpdateContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindimage.ResourceName), policyoperations.WithOperationType(policyoperations.Update))),
		DeleteContext: schema.DeleteContextFunc(policyoperations.ResourceOperation(policyoperations.WithResourceName(policykindimage.ResourceName), policyoperations.WithOperationType(policyoperations.Delete))),
		Schema:        imagePolicySchema,
		CustomizeDiff: customdiff.All(
			schema.CustomizeDiffFunc(scope.ValidateScope(policyoperations.ScopeMap[policykindimage.ResourceName])),
			policykindimage.ValidateInput,
			policy.ValidateSpecLabelSelectorRequirement,
		),
	}
}

var (
	ScopesAllowed = [...]string{scope.OrganizationKey, scope.WorkspaceKey}
	ScopeSchema   = scope.GetScopeSchema(
		scope.WithDescription(fmt.Sprintf("Scope for the image policy, having one of the valid scopes: %v.", strings.Join(ScopesAllowed[:], `, `))),
		scope.WithScopes(ScopesAllowed[:]))
)

var imagePolicySchema = map[string]*schema.Schema{
	policy.NameKey: {
		Type:        schema.TypeString,
		Description: "Name of the image policy",
		Required:    true,
		ForceNew:    true,
	},
	scope.ScopeKey: ScopeSchema,
	common.MetaKey: common.Meta,
	policy.SpecKey: policykindimage.SpecSchema,
}
