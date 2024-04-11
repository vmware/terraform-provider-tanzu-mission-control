/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package data

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tapeulamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tap/eula"
)

var EULAData = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Data for the TAP EULA",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			acceptedKey: {
				Type:        schema.TypeBool,
				Description: "Identifies whether this user has accepted the EULA terms.",
				Computed:    true,
			},
			EulaURLKey: {
				Type:        schema.TypeString,
				Description: "URL at which this end user license agreement can be found.",
				Computed:    true,
			},
			releasedAtKey: {
				Type:        schema.TypeString,
				Description: "Time when this EULA version was released.",
				Computed:    true,
			},
			userKey: {
				Type:        schema.TypeString,
				Description: "User email identifier.",
				Computed:    true,
			},
		},
	},
}

func FlattenEULAData(eulaData *tapeulamodel.VmwareTanzuManageV1alpha1TanzupackageTapEulaData) (data interface{}) {
	if eulaData == nil {
		return data
	}

	flattenEULAData := make(map[string]interface{})

	flattenEULAData[acceptedKey] = eulaData.Accepted
	flattenEULAData[EulaURLKey] = eulaData.EulaURL
	flattenEULAData[releasedAtKey] = eulaData.ReleasedAt.String()
	flattenEULAData[userKey] = eulaData.User

	return []interface{}{flattenEULAData}
}
