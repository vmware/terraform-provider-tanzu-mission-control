/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyoperations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindcustom "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom"
	policykindimage "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image"
	policykindsecurity "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
)

func ResourcePolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}, rn string) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	policyName, ok := d.Get(policy.NameKey).(string)
	if !ok {
		return diag.Errorf("unable to read %s policy name", rn)
	}

	scopedFullnameData := scope.ConstructScope(d, policyName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control %s policy entry; Scope full name is empty", rn)
	}

	UID, meta, spec, err := RetrievePolicyUIDMetaAndSpecFromServer(config, scopedFullnameData, d, policyName, rn)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(meta)); err != nil {
		return diag.FromErr(err)
	}

	var flattenedSpec []interface{}

	switch rn {
	case policykindcustom.ResourceName:
		flattenedSpec = policykindcustom.FlattenSpec(spec)
	case policykindsecurity.ResourceName:
		flattenedSpec = policykindsecurity.FlattenSpec(spec)
	case policykindimage.ResourceName:
		flattenedSpec = policykindimage.FlattenSpec(spec)
	}

	if err := d.Set(policy.SpecKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
