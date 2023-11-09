/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementclusterregistration

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	managementclusterregistrationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/managementclusterregistration"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const defaultWaitTimeout = 15 * time.Minute

func ResourceManagementClusterRegistration() *schema.Resource {
	return &schema.Resource{
		Schema: managementClusterRegistrationSchema,
		//CreateContext: , TODO
		//ReadContext: , TODO
		//UpdateContext: , TODO
		//DeleteContext: , TODO
		Description: "Tanzu Mission Control Management Cluster Registration Resource",
	}
}

var managementClusterRegistrationSchema = map[string]*schema.Schema{
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of this management cluster",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey: common.Meta,
	specKey: {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				clusterGroupKey: {
					Type:         schema.TypeString,
					Description:  "Cluster group name to be used by default for workload clusters",
					Required:     true,
					ValidateFunc: validation.All(validation.StringIsNotEmpty),
				},
				kubernetesProviderTypeKey: {
					Type:         schema.TypeString,
					Description:  "Kubernetes provider type",
					Required:     true,
					ValidateFunc: validation.All(validation.StringIsNotEmpty),
				},
				imageRegistryKey: {
					Type:        schema.TypeString,
					Description: "Image registry which is only allowed for TKGm",
					Optional:    true,
				},
				managedWorkloadClusterImageRegistryKey: {
					Type:        schema.TypeString,
					Description: "Managed workload cluster image registry",
					Optional:    true,
				},
				managementClusterProxyNameKey: {
					Type:        schema.TypeString,
					Description: "Management cluster proxy name",
					Optional:    true,
				},
				managedWorkloadClusterProxyNameKey: {
					Type:        schema.TypeString,
					Description: "Managed workload cluster proxy name",
					Optional:    true,
				},
			},
		},
	},
	registerClusterKey: {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				registerClusterKubeConfigPathForTKGmKey: {
					Type:        schema.TypeString,
					Description: "Register management cluster KUBECONFIG path for only TKGm",
					ForceNew:    true,
					Optional:    true,
				},
				registerClusterKubeConfigRawForTKGmKey: {
					Type:        schema.TypeString,
					Description: "Register management cluster KUBECONFIG for only TKGm",
					Optional:    true,
					ForceNew:    true,
					Sensitive:   true,
				},
				registerClusterDescriptionForTKGmKey: {
					Type:         schema.TypeString,
					Description:  "Register management cluster description for only TKGm",
					Optional:     true,
					ValidateFunc: validation.StringIsNotWhiteSpace,
				},
			},
		},
	},
	StatusKey: {
		Type:        schema.TypeMap,
		Description: "Status of the management cluster",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	waitKey: {
		Type:        schema.TypeString,
		Description: "Wait timeout duration.",
		Default:     defaultWaitTimeout.String(),
		Optional:    true,
	},
}

var KubeConfigWayAllowed = [...]string{registerClusterKubeConfigPathForTKGmKey, registerClusterKubeConfigRawForTKGmKey}

// nolint: unused
func constructSpec(d *schema.ResourceData) (spec *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec) {
	spec = &managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec{}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	if v, ok := specData[clusterGroupKey]; ok {
		spec.DefaultClusterGroup = v.(string)
	}

	if v, ok := specData[managedWorkloadClusterImageRegistryKey]; ok {
		spec.DefaultWorkloadClusterImageRegistry = v.(string)
	}

	if v, ok := specData[managedWorkloadClusterProxyNameKey]; ok {
		spec.DefaultWorkloadClusterProxyName = v.(string)
	}

	if v, ok := specData[imageRegistryKey]; ok {
		spec.ImageRegistry = v.(string)
	}

	if v, ok := specData[kubernetesProviderTypeKey]; ok {
		providerType := clustermodel.VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType(v.(string))
		spec.KubernetesProviderType = &providerType
	}

	if v, ok := specData[managementClusterProxyNameKey]; ok {
		spec.ProxyName = v.(string)
	}

	return spec
}

func flattenSpec(spec *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[clusterGroupKey] = spec.DefaultClusterGroup

	flattenSpecData[managedWorkloadClusterImageRegistryKey] = spec.DefaultWorkloadClusterImageRegistry

	flattenSpecData[managedWorkloadClusterProxyNameKey] = spec.DefaultWorkloadClusterProxyName

	flattenSpecData[imageRegistryKey] = spec.ImageRegistry

	if spec.KubernetesProviderType != nil {
		flattenSpecData[kubernetesProviderTypeKey] = string(*spec.KubernetesProviderType)
	}

	flattenSpecData[managementClusterProxyNameKey] = spec.ProxyName

	return []interface{}{flattenSpecData}
}
