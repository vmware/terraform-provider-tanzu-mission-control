/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policy

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

func MetaResourceAttributeCheck(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.uid"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.resource_version"),
	}
}
