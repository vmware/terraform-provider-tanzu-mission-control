/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	secretmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/scope"
)

func ResourceSecret() *schema.Resource {
	return &schema.Resource{
		Schema: secretSchema,
	}
}

var secretSchema = map[string]*schema.Schema{
	NameKey: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	NamespaceNameKey: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	OrgIDKey: {
		Type: schema.TypeString,
	},
	scope.ScopeKey: scope.ScopeSchema,
	specKey:        secretSpec,
	ExportKey: {
		Type:     schema.TypeBool,
		Required: true,
	},
	statusKey: {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
}

// nolint: unused
type dockerConfigJSON struct {
	Auths map[string]map[string]interface{} `json:"auths,omitempty"`
}

var secretSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the kubernetes secret",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			DockerConfigjsonKey: {
				Type:        schema.TypeList,
				Description: "SecretType definition - SECRET_TYPE_DOCKERCONFIGJSON, Kubernetes secrets type.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						UsernameKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - indicates the kubernetes username.",
						},
						PasswordKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - indicates the kubernetes password.",
						},
						ImageRegistryURLKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - indicates the kubernetes image registry url.",
						},
					},
				},
			},
		},
	},
}

// nolint: unused
func constructSpec(d *schema.ResourceData) (spec *secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec) {
	spec = &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	spec.SecretType = secretmodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEDOCKERCONFIGJSON)

	if v, ok := specData[DockerConfigjsonKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			auth := v1[0].(map[string]interface{})

			var serverURL, username, password string

			if v, ok := auth[ImageRegistryURLKey]; ok {
				serverURL = v.(string)
			}

			if v, ok := auth[UsernameKey]; ok {
				username = v.(string)
			}

			if v, ok := auth[PasswordKey]; ok {
				password = v.(string)
			}

			secretSpecData, err := getEncodedSpecData(serverURL, username, password)
			if err != nil {
				return nil
			}

			spec.Data = map[string]strfmt.Base64{
				dockerConfigJSONSecretDataKey: secretSpecData,
			}
		}
	}

	return spec
}

// nolint: unused
func flattenSpec(spec *secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	secretData, ok := spec.Data[dockerConfigJSONSecretDataKey]
	if ok {
		return nil
	}

	dockerConfigJSON, err := getDecodedSpecData(secretData)
	if err != nil {
		return data
	}

	var dockerConfigJSONData = make(map[string]interface{})

	for serverURL, creds := range dockerConfigJSON.Auths {
		for attribute, value := range creds {
			dockerConfigJSONData[ImageRegistryURLKey] = serverURL

			if attribute == UsernameKey {
				stringValue, ok := value.(string)
				if ok {
					return data
				}

				dockerConfigJSONData[UsernameKey] = stringValue
			}
		}
	}

	flattenSpecData[DockerConfigjsonKey] = dockerConfigJSONData

	return []interface{}{flattenSpecData}
}

// nolint: unused
func getEncodedSpecData(serverURL, username, password string) (strfmt.Base64, error) {
	var secretspecdata strfmt.Base64

	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	err := secretspecdata.Scan(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"auths":{"%s":{"username":"%s","password":"%s","auth":"%s"}}}`, serverURL, username, password, auth))))
	if err != nil {
		return nil, err
	}

	return secretspecdata, nil
}

// nolint: unused
func getDecodedSpecData(data strfmt.Base64) (*dockerConfigJSON, error) {
	rawData, err := base64.StdEncoding.DecodeString(data.String())
	if err != nil {
		return nil, err
	}

	dockerConfigJSON := &dockerConfigJSON{}

	err = json.Unmarshal(rawData, dockerConfigJSON)
	if err != nil {
		return nil, err
	}

	return dockerConfigJSON, nil
}
