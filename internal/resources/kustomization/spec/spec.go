/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
)

var SpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the Repository.",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			sourceKey: sourceSchema,
			pathKey: {
				Type:        schema.TypeString,
				Description: "Path within the source from which configurations will be applied. The path must exactly match what is in the repository.",
				Required:    true,
			},
			pruneKey: {
				Type:        schema.TypeBool,
				Description: "If true, the workloads will be deleted when the kustomization CR is deleted. When prune is enabled, removing the kustomization will trigger a removal of all kubernetes objects previously applied on all clusters of this cluster group by this kustomization.",
				Optional:    true,
				Default:     false,
			},
			intervalKey: {
				Type:        schema.TypeString,
				Description: "Interval defines the interval at which to reconcile kustomization.",
				Optional:    true,
				Default:     "5m",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					durationInScript, err := time.ParseDuration(old)
					if err != nil {
						return false
					}

					durationInState, err := time.ParseDuration(new)
					if err != nil {
						return false
					}

					return durationInScript.Seconds() == durationInState.Seconds()
				},
			},
			targetNamespaceKey: {
				Type:        schema.TypeString,
				Description: "TargetNamespace sets or overrides the namespaces of resources/kustomization yaml while applying on cluster. Namespace specified here must exist on cluster. It won't be created as a result of specifying here. Enter the name of the namespace you want the kustomization to be synced to. Entering a target namespace removes the need to specify a namespace in your kustomization. If the namespace does not exist in the cluster, syncing the kustomization will fail.",
				Optional:    true,
				Default:     "",
			},
		},
	},
}

var sourceSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Reference to the source from which the configurations will be applied. Please select an existing repository.",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			nameKey: {
				Type:        schema.TypeString,
				Description: "Name of the repository.",
				Required:    true,
			},
			namespaceKey: {
				Type:        schema.TypeString,
				Description: "Namespace of the repository.",
				Required:    true,
			},
		},
	},
}

func expandSource(data []interface{}) (source *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference) {
	if len(data) == 0 || data[0] == nil {
		return source
	}

	sourceData, _ := data[0].(map[string]interface{})

	source = &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{}

	if nameValue, ok := sourceData[nameKey]; ok {
		helper.SetPrimitiveValue(nameValue, &source.Name, nameKey)
	}

	if namespaceValue, ok := sourceData[namespaceKey]; ok {
		helper.SetPrimitiveValue(namespaceValue, &source.Namespace, namespaceKey)
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
