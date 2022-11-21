/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var (
	integrationSchema = map[string]*schema.Schema{
		managementClusterNameKey: {
			Type:        schema.TypeString,
			Description: "Name of the management cluster",
			Default:     attachedValue,
			Optional:    true,
			ForceNew:    true,
		},

		provisionerNameKey: {
			Type:        schema.TypeString,
			Description: "Provisioner of the cluster",
			Default:     attachedValue,
			Optional:    true,
			ForceNew:    true,
		},

		clusterNameKey: {
			Type:        schema.TypeString,
			Description: "Name of this cluster",
			Required:    true,
			ForceNew:    true,
		},

		integrationNameKey: {
			Type:             schema.TypeString,
			Description:      "Name of the Integration; valid options are currently only ['tanzu-service-mesh']",
			Required:         true,
			ForceNew:         true,
			ValidateDiagFunc: validateIntegrationName,
		},

		common.MetaKey: common.Meta,

		specKey: {
			Type:        schema.TypeList,
			Description: "Specification for the Integration",
			Required:    true,
			ForceNew:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        specResource,
		},

		statusKey: {
			Type:        schema.TypeMap,
			Description: "Status of Integration",
			Computed:    true,
			Elem:        computedString,
		},
	}

	specResource = &schema.Resource{
		Schema: map[string]*schema.Schema{
			configurationKey: {
				Type:             schema.TypeString,
				Description:      "Integration specific configurations in JSON format",
				Optional:         true,
				ForceNew:         true,
				ValidateDiagFunc: validateConfiguration,
			},
		},
	}

	computedString = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	// validIntegrationNames are ones allowed by the API.
	// | value must be in list [tanzu-observability-saas tanzu-service-mesh].
	validIntegrationNames = []string{
		tanzuObservabilitySaaSValue,
		tanzuServiceMeshValue,
	}
)

func validateIntegrationName(i interface{}, p cty.Path) diag.Diagnostics {
	var (
		ok bool
		v  string
	)

	if v, ok = i.(string); !ok {
		return diag.Errorf("unexpected type for integration_name: %T (%v)", i, p)
	}

	for _, name := range validIntegrationNames {
		if name == v {
			return nil
		}
	}

	return diag.Errorf("integration_name must be one of: %v", validIntegrationNames)
}

func validateConfiguration(i interface{}, p cty.Path) diag.Diagnostics {
	var (
		ok bool
		v  string
		m  map[string]interface{}
	)

	if v, ok = i.(string); !ok {
		return diag.Errorf("unexpected type for configuration: %T (%v)", i, p)
	}

	if err := json.Unmarshal([]byte(v), &m); err != nil {
		return diag.FromErr(fmt.Errorf("%w: %s", err, v))
	}

	return nil
}
