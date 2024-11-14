//go:build integration
// +build integration

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package integration

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/integration"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

var (
	client = newTestClient()

	testProvider = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			ResourceName: ResourceIntegration(),
		},
		ConfigureContextFunc: configureTestProvider(client),
	}

	validMinimalModel = &integration.VmwareTanzuManageV1alpha1ClusterIntegrationIntegration{
		FullName: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName{
			Name:                  "tanzu-service-mesh",
			ClusterName:           "test-cluster",
			ManagementClusterName: "attached",
			ProvisionerName:       "attached",
		},
		Spec: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec{},
	}

	validFullModel = &integration.VmwareTanzuManageV1alpha1ClusterIntegrationIntegration{
		FullName: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName{
			Name:                  "tanzu-service-mesh",
			ClusterName:           "test-cluster-2",
			ManagementClusterName: "attached-2",
			ProvisionerName:       "attached-2",
		},
		Spec: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec{
			Configurations: map[string]interface{}{
				enableNamespaceExclusionsSpecKey: true,
				namespaceExclusionsSpecKey:       []map[string]interface{}{},
			},
		},
	}
)

func TestResourceIntegration(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:        false,
		ProviderFactories: testhelper.GetTestProviderFactories(testProvider),
		Steps: []resource.TestStep{
			{
				Config: must(generateResourceManifest(nil)),
			},
			{
				Config:             must(generateResourceManifest(validMinimalModel)),
				ExpectNonEmptyPlan: true,
			},
			{
				Config:             must(generateResourceManifest(validFullModel)),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
