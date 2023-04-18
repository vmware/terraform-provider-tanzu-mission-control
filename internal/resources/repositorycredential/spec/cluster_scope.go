/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	repositorycredentialclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/repositorycredential/cluster"
)

type CredMapping struct {
	Data map[string]string `json:"data,omitempty"`
}

func ConstructSpecForClusterScope(d *schema.ResourceData) (spec *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec) {
	spec = &repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{}

	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})
	if v, ok := specData[dataKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			auth := v1[0].(map[string]interface{})
			if specData[secretTypeKey] == usernamePasswordKey {

				if v, ok := auth[usernameKey]; ok {
					spec.Data.Data[usernameKey] = v.(string)
				}

				if v, ok := auth[passwordKey]; ok {
					spec.Data.Data[passwordKey] = v.(string)
				}
			}

			// if specData[secretTypeKey] == sshKey {
			// 	if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			// 		auth := v1[0].(map[string]interface{})

			// 		if v, ok := auth[usernameKey]; ok {
			// 			spec.Data.Data[usernameKey] = v.(string)
			// 		}

			// 		if v, ok := auth[passwordKey]; ok {
			// 			spec.Data.Data[passwordKey] = v.(string)
			// 		}
			// 	}
			// }
		}
	}

	return spec
}

func FlattenSpecForClusterScope(spec *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	if spec.Data != nil {
		data := make(map[string]string)
		for key, value := range spec.Data {
			data[key] = string(value)
		}
		flattenSpecData[dataKey] = data
	}

	flattenSpecData[secretTypeKey] = spec.RepositorycredentialType

	return []interface{}{flattenSpecData}
}
