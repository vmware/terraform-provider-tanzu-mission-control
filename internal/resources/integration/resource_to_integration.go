/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	integrationschema "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/integration/schema"
	integrationscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/integration/scope"
	clusterintegration "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/integration/scope/cluster"
	clustergroupintegration "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/integration/scope/clustergroup"
)

func ResourceTOIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTOIntegrationCreate,
		ReadContext:   resourceTOIntegrationRead,
		UpdateContext: resourceTOIntegrationUpdate,
		DeleteContext: resourceTOIntegrationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceTOIntegrationImporter,
		},
		Schema: integrationschema.TOIntegrationSchema,
	}
}

var supportedScopes = []string{
	string(integrationscope.ClusterScopeType),
	string(integrationscope.ClusterGroupScopeType),
}

func resourceTOIntegrationCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	scopeData := data.Get(integrationschema.ScopeKey).([]interface{})[0].(map[string]interface{})
	clusterScopeData, _ := scopeData[integrationschema.ClusterScopeKey].([]interface{})

	if len(clusterScopeData) > 0 {
		return clusterintegration.ClusterTOIntegrationCreate(ctx, data, m)
	} else {
		return clustergroupintegration.ClusterGroupTOIntegrationCreate(ctx, data, m)
	}
}

// scopeData := data.Get(integrationschema.ScopeKey).([]interface{})[0].(map[string]interface{})
// clusterScopeData, _ := scopeData[integrationschema.ClusterScopeKey].([]interface{})
//
// if len(clusterScopeData) > 0 {
// 	return clusterintegration.ClusterTOIntegrationUpdate(ctx, data, m)
// } else {
// 	return clustergroupintegration.ClusterGroupTOIntegrationUpdate(ctx, data, m)
// }

func resourceTOIntegrationUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return diag.FromErr(errors.New("Update of tanzu observability integration is not supported."))
}

func resourceTOIntegrationDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	scopeData := data.Get(integrationschema.ScopeKey).([]interface{})[0].(map[string]interface{})
	clusterScopeData, _ := scopeData[integrationschema.ClusterScopeKey].([]interface{})

	if len(clusterScopeData) > 0 {
		return clusterintegration.ClusterTOIntegrationDelete(ctx, data, m)
	} else {
		return clustergroupintegration.ClusterGroupTOIntegrationDelete(ctx, data, m)
	}
}

func resourceTOIntegrationRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	scopeData := data.Get(integrationschema.ScopeKey).([]interface{})[0].(map[string]interface{})
	clusterScopeData, _ := scopeData[integrationschema.ClusterScopeKey].([]interface{})

	if len(clusterScopeData) > 0 {
		return clusterintegration.ClusterTOIntegrationRead(ctx, data, m)
	} else {
		return clustergroupintegration.ClusterGroupTOIntegrationRead(ctx, data, m)
	}
}

func resourceTOIntegrationImporter(ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	integrationFullName := data.Id()
	scopeType, scopeID, foundSep := strings.Cut(integrationFullName, ":")

	if foundSep {
		if scopeID == "" {
			return nil, errors.New("Couldn't import tanzu observability integration because ID is invalid.\nScope ID can't be empty.")
		}

		scope := integrationscope.SupportedScopes(scopeType)

		switch scope {
		case integrationscope.ClusterScopeType:
			return clusterintegration.ClusterTOIntegrationImporter(scopeID, ctx, data, m)
		case integrationscope.ClusterGroupScopeType:
			return clustergroupintegration.ClusterGroupTOIntegrationImporter(scopeID, ctx, data, m)
		default:
			return nil, errors.Errorf("Couldn't import tanzu observability integration because the provided scope is not supported.\nSupported scopes are: %v", supportedScopes)
		}
	} else {
		return nil, errors.New("Couldn't import tanzu observability integration because ID is invalid.\nID must start with a scope type (cluster/cluster_group) followed by a colon and a proper scope id.")
	}
}
