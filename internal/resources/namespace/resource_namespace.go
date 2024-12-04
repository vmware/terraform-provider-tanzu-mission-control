// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	namespacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/namespace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func ResourceNamespace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNamespaceCreate,
		ReadContext:   dataSourceNamespaceRead,
		UpdateContext: resourceNamespaceInPlaceUpdate,
		DeleteContext: resourceNamespaceDelete,
		Schema:        namespaceSchema,
	}
}

var namespaceSchema = map[string]*schema.Schema{
	NameKey: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	ManagementClusterNameKey: {
		Type:     schema.TypeString,
		Default:  "attached",
		Optional: true,
		ForceNew: true,
	},
	ProvisionerNameKey: {
		Type:     schema.TypeString,
		Default:  "attached",
		Optional: true,
		ForceNew: true,
	},
	ClusterNameKey: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	common.MetaKey: common.Meta,
	specKey:        namespaceSpec,
	statusKey: {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
}

func constructFullname(d *schema.ResourceData) (fullname *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) {
	fullname = &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName{}

	fullname.ClusterName, _ = d.Get(ClusterNameKey).(string)

	fullname.ManagementClusterName, _ = d.Get(ManagementClusterNameKey).(string)

	fullname.Name, _ = d.Get(NameKey).(string)

	fullname.ProvisionerName, _ = d.Get(ProvisionerNameKey).(string)

	return fullname
}

var namespaceSpec = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			workspaceNameKey: {
				Type:     schema.TypeString,
				Default:  workspaceNameDefaultValue,
				Optional: true,
			},
			attachKey: {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
		},
	},
}

func constructSpec(d *schema.ResourceData) (spec *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceSpec) {
	spec = &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceSpec{
		WorkspaceName: workspaceNameDefaultValue,
	}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	if v, ok := specData[workspaceNameKey]; ok {
		spec.WorkspaceName, _ = v.(string)
	}

	if v, ok := specData[attachKey]; ok {
		spec.Attach, _ = v.(bool)
	}

	return spec
}

func flattenSpec(spec *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[workspaceNameKey] = spec.WorkspaceName

	flattenSpecData[attachKey] = spec.Attach

	return []interface{}{flattenSpecData}
}

func resourceNamespaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	namespaceRequest := &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest{
		Namespace: &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceNamespace{
			FullName: constructFullname(d),
			Meta:     common.ConstructMeta(d),
			Spec:     constructSpec(d),
		},
	}

	namespaceResponse, err := config.TMCConnection.NamespaceResourceService.ManageV1alpha1NamespaceResourceServiceCreate(namespaceRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create Tanzu Mission Control namespace entry, name : %s", NameKey))
	}

	d.SetId(namespaceResponse.Namespace.Meta.UID)

	return dataSourceNamespaceRead(ctx, d, m)
}

func resourceNamespaceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	namespaceName, _ := d.Get(NameKey).(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.NamespaceResourceService.ManageV1alpha1NamespaceResourceServiceDelete(constructFullname(d))
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "unable to delete Tanzu Mission Control namespace entry, name : %s", namespaceName))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	_ = schema.RemoveFromState(d, m)

	return diags
}

func resourceNamespaceInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	updateRequired := false

	switch {
	case common.HasMetaChanged(d):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(specKey, workspaceNameKey)):
		updateRequired = true
	}
	// todo: Updating the description field for namespace resource after `OLYMP-23394` is resolved.

	if !updateRequired {
		return diags
	}

	getResp, err := config.TMCConnection.NamespaceResourceService.ManageV1alpha1NamespaceResourceServiceGet(constructFullname(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to get Tanzu Mission Control namespace entry, name : %s", d.Get(ClusterNameKey)))
	}

	if common.HasMetaChanged(d) {
		meta := common.ConstructMeta(d)

		if value, ok := getResp.Namespace.Meta.Labels[common.CreatorLabelKey]; ok {
			meta.Labels[common.CreatorLabelKey] = value
		}

		getResp.Namespace.Meta.Labels = meta.Labels
		getResp.Namespace.Meta.Description = meta.Description
	}

	incomingNamespaceName := d.Get(helper.GetFirstElementOf(specKey, workspaceNameKey))

	if incomingNamespaceName.(string) != "" {
		getResp.Namespace.Spec.WorkspaceName = incomingNamespaceName.(string)
	}

	_, err = config.TMCConnection.NamespaceResourceService.ManageV1alpha1NamespaceResourceServiceUpdate(
		&namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceRequest{
			Namespace: getResp.Namespace,
		},
	)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to update Tanzu Mission Control namespace entry, name : %s", d.Get(ClusterNameKey)))
	}

	return dataSourceNamespaceRead(ctx, d, m)
}
