// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package sourcesecret

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	sourcesecretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
	sourcesecretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret/spec"
)

func DataSourceSourcesecret() *schema.Resource {
	return &schema.Resource{
		Schema: getDataSourceSchema(),
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return resourceSourcesecretRead(context.WithValue(ctx, contextMethodKey{}, DataSourceRead), d, m)
		},
	}
}

func resourceSourcesecretRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	sourcesecretName, ok := d.Get(nameKey).(string)
	if !ok {
		return diag.Errorf("unable to read source secret name")
	}

	scopedFullnameData, scopesFound := scope.ConstructScope(d, sourcesecretName)

	if len(scopesFound) == 0 {
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v", strings.Join(scope.CredentialTypesAllowed[:], `, `))
	} else if len(scopesFound) > 1 {
		return diag.Errorf("found scopes: %v are not valid: maximum one valid scope type block is allowed", strings.Join(scopesFound, `, `))
	}

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control source secret entry; Scope full name is empty")
	}

	UID, meta, atomicSpec, err := retrieveSourcesecretUIDMetaAndSpecFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			if ctx.Value(contextMethodKey{}) == DataSourceRead {
				return diag.FromErr(errors.Wrapf(err, "Tanzu Mission Control source secret entry not found, name : %s", d.Get(nameKey)))
			}

			_ = schema.RemoveFromState(d, m)

			return diags
		}

		return diag.FromErr(err)
	}

	// always run
	d.SetId(UID)

	if err := d.Set(common.MetaKey, common.FlattenMeta(meta)); err != nil {
		return diag.FromErr(err)
	}

	var (
		flattenedSpec []interface{}
		specTypeData  string
	)

	switch *atomicSpec.SourceSecretType {
	case *sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeUSERNAMEPASSWORD):
		if _, ok := d.GetOk(spec.SpecKey); ok {
			specTypeData, _ = (d.Get(helper.GetFirstElementOf(spec.SpecKey, spec.DataKey, spec.UsernamePasswordKey, spec.PasswordKey))).(string)
		}
	case *sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeSSH):
		if _, ok := d.GetOk(spec.SpecKey); ok {
			specTypeData, _ = (d.Get(helper.GetFirstElementOf(spec.SpecKey, spec.DataKey, spec.SSHKey, spec.IdentityKey))).(string)
		}
	default:
		specTypeData = ""
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		flattenedSpec = spec.FlattenSpecForClusterScope(atomicSpec, specTypeData)
	case commonscope.ClusterGroupScope:
		clusterGroupScopeSpec := &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec{
			AtomicSpec: atomicSpec,
		}
		flattenedSpec = spec.FlattenSpecForClusterGroupScope(clusterGroupScopeSpec, specTypeData)
	}

	if err := d.Set(spec.SpecKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
