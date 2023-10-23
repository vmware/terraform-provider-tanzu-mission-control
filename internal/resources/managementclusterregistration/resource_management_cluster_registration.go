/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementclusterregistration

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	managementclusterregistrationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/managementclusterregistration"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

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
		Description: "Name of this cluster",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey:   common.Meta,
	specKey:          managementClusterRegistrationSpecSchema,
	attachClusterKey: attachCluster,
	StatusKey: {
		Type:        schema.TypeMap,
		Description: "Status of the cluster",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	waitKey: {
		Type:        schema.TypeString,
		Description: "Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero. Should be set to 0 in case of simple attach cluster where kubeconfig input is not provided.",
		Default:     "default",
		Optional:    true,
	},
}

var attachCluster = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MinItems: 1,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			attachClusterKubeConfigPathKey: {
				Type:        schema.TypeString,
				Description: "Attach cluster KUBECONFIG path",
				ForceNew:    true,
				Optional:    true,
			},
			attachClusterKubeConfigRawKey: {
				Type:        schema.TypeString,
				Description: "Attach cluster KUBECONFIG",
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
			},
			attachClusterDescriptionKey: {
				Type:         schema.TypeString,
				Description:  "Attach cluster description",
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
		},
	},
}
var KubeConfigWayAllowed = [...]string{attachClusterKubeConfigPathKey, attachClusterKubeConfigRawKey}

var managementClusterRegistrationSpecSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MinItems: 1,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			defaultClusterGroupKey: {
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
				Description: "Image registry witch is only allowed for TKGm",
				Optional:    true,
			},
			defaultWorkloadClusterImageRegistryKey: {
				Type:        schema.TypeString,
				Description: "Default workload cluster image registry",
				Optional:    true,
			},
			proxyNameKey: {
				Type:        schema.TypeString,
				Description: "Proxy name",
				Optional:    true,
			},
			defaultWorkloadClusterProxyNameKey: {
				Type:        schema.TypeString,
				Description: "Default workload cluster proxy name",
				Optional:    true,
			},
		},
	},
}

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

	if v, ok := specData[defaultClusterGroupKey]; ok {
		spec.DefaultClusterGroup = v.(string)
	}

	if v, ok := specData[defaultWorkloadClusterImageRegistryKey]; ok {
		spec.DefaultWorkloadClusterImageRegistry = v.(string)
	}

	if v, ok := specData[defaultWorkloadClusterProxyNameKey]; ok {
		spec.DefaultWorkloadClusterProxyName = v.(string)
	}

	if v, ok := specData[imageRegistryKey]; ok {
		spec.ImageRegistry = v.(string)
	}

	if v, ok := specData[kubernetesProviderTypeKey]; ok {
		providerType := clustermodel.VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType(v.(string))
		spec.KubernetesProviderType = &providerType
	}

	if v, ok := specData[proxyNameKey]; ok {
		spec.ProxyName = v.(string)
	}

	return spec
}

func flattenSpec(spec *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[defaultClusterGroupKey] = spec.DefaultClusterGroup

	flattenSpecData[defaultWorkloadClusterImageRegistryKey] = spec.DefaultWorkloadClusterImageRegistry

	flattenSpecData[defaultWorkloadClusterProxyNameKey] = spec.DefaultWorkloadClusterProxyName

	flattenSpecData[imageRegistryKey] = spec.ImageRegistry

	if spec.KubernetesProviderType != nil {
		flattenSpecData[kubernetesProviderTypeKey] = string(*spec.KubernetesProviderType)
	}

	flattenSpecData[proxyNameKey] = spec.ProxyName

	return []interface{}{flattenSpecData}
}
