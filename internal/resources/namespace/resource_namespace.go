/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
	clienterrors "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/errors"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/helper"
	namespacemodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/namespace"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/common"
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
	nameKey: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	managementClusterNameKey: {
		Type:     schema.TypeString,
		Default:  "attached",
		Optional: true,
		ForceNew: true,
	},
	provisionerNameKey: {
		Type:     schema.TypeString,
		Default:  "attached",
		Optional: true,
		ForceNew: true,
	},
	clusterNameKey: {
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

	fullname.ClusterName, _ = d.Get(clusterNameKey).(string)

	fullname.ManagementClusterName, _ = d.Get(managementClusterNameKey).(string)

	fullname.Name, _ = d.Get(nameKey).(string)

	fullname.ProvisionerName, _ = d.Get(provisionerNameKey).(string)

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
		return diag.FromErr(errors.Wrapf(err, "unable to create tanzu TMC namespace entry, name : %s", nameKey))
	}

	d.SetId(namespaceResponse.Namespace.Meta.UID)

	return dataSourceNamespaceRead(ctx, d, m)
}

func resourceNamespaceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	namespaceName, _ := d.Get(nameKey).(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.NamespaceResourceService.ManageV1alpha1NamespaceResourceServiceDelete(constructFullname(d))
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "unable to delete tanzu TMC namespace entry, name : %s", namespaceName))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

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
		return diag.FromErr(errors.Wrapf(err, "unable to get tanzu TMC namespace entry, name : %s", d.Get(clusterNameKey)))
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
		return diag.FromErr(errors.Wrapf(err, "unable to update tanzu TMC namespace entry, name : %s", d.Get(clusterNameKey)))
	}

	return dataSourceNamespaceRead(ctx, d, m)
}
