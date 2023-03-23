/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	repositorycredentialclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/repositorycredential/cluster"
)

var SpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the Repository Credential.",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			secretTypeKey: {
				Type:        schema.TypeString,
				Description: "The type of credential that will be used for the repository. Options are SSH or USERNAME_PASSWORD",
				Required:    true,
			},
			dataKey: dataSchema,
		},
	},
}

var dataSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The SSH key or basic auth credentials used for authenticating to an existing repository",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			usernameKey: {
				Type:        schema.TypeString,
				Description: "The username when using basic auth for authenticating to the repository",
				Required:    false,
			},
			passwordKey: {
				Type:        schema.TypeString,
				Description: "The password when using basic auth for authenticating to the repository",
				Required:    false,
			},
			sshKey: {
				Type:        schema.TypeString,
				Description: "the SSH key used for authenticating to the repository",
				Required:    false,
			},
			knownHostsKey: {
				Type:        schema.TypeString,
				Description: "the known hosts list to use when the source secret type is SSH",
				Required:    false,
			},
		},
	},
}

func expandData(data []interface{}) (source *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType) {
	if len(data) == 0 || data[0] == nil {
		return source
	}

	sourceData, _ := data[0].(map[string]interface{})

	source = &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{}

	if v, ok := sourceData[nameKey]; ok {
		helper.SetPrimitiveValue(v, &source.Name, nameKey)
	}

	if v, ok := sourceData[namespaceKey]; ok {
		helper.SetPrimitiveValue(v, &source.Namespace, namespaceKey)
	}

	return source
}

func flattenSource(source *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference) (data []interface{}) {
	if source == nil {
		return data
	}

	flattenSourceData := make(map[string]interface{})

	flattenSourceData[nameKey] = source.Name
	flattenSourceData[namespaceKey] = source.Namespace

	return []interface{}{flattenSourceData}
}

func HasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(SpecKey, sourceKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, pathKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, pruneKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, intervalKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, targetNamespaceKey)):
		updateRequired = true
	}

	return updateRequired
}
