// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package testing

import "github.com/hashicorp/terraform-plugin-testing/helper/resource"

const (
	MetaTemplate = `meta {
		description = "resource with description"
		labels = {
			"key1" : "value1"
			"key2" : "value2"
		}
	}`
)

func MetaDataSourceAttributeCheck(dataSourceName, resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.description", resourceName, "meta.0.description"),
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key1", resourceName, "meta.0.labels.key1"),
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key2", resourceName, "meta.0.labels.key2"),
		resource.TestCheckResourceAttrSet(dataSourceName, "meta.0.annotations.authoritativeRID"),
		resource.TestCheckResourceAttrSet(dataSourceName, "meta.0.uid"),
	}
}

func MetaResourceAttributeCheck(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "meta.0.description", description),
		resource.TestCheckResourceAttr(resourceName, "meta.0.labels.key1", value1),
		resource.TestCheckResourceAttr(resourceName, "meta.0.labels.key2", value2),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.uid"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.annotations.authoritativeRID"),
	}
}
