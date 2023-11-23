//go:build tointegration
// +build tointegration

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	clusterintegrationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/integration/cluster"
	clustergroupintegrationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/integration/clustergroup"
	integrationschema "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/integration/schema"
	integrationscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/integration/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	CREDENTIALS_NAME_ENV = "TO_CREDENTIALS_NAME"
)

var (
	context = authctx.TanzuContext{
		ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
		Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
		VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		TLSConfig:        &proxy.TLSConfig{},
	}
)

func TestAcceptanceTOIntegrationResource(t *testing.T) {
	err := context.Setup()

	if err != nil {
		t.Error(errors.Wrap(err, "unable to set the context"))
		t.FailNow()
	}

	credentialsName, found := os.LookupEnv(CREDENTIALS_NAME_ENV)

	if !found {
		t.Errorf("Environment variable '%s' must be set with tanzu observability credentials name.", CREDENTIALS_NAME_ENV)
		t.FailNow()
	}

	var (
		tfResourceConfigBuilder = InitResourceTFConfigBuilder(testScopeHelper)
		provider                = initTestProvider(t)
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: tfResourceConfigBuilder.GetClusterTOIntegrationConfig(credentialsName),
				Check: resource.ComposeTestCheckFunc(
					verifyTOIntegrationResourceCreation(provider, ClusterIntegrationResourceFullName, integrationscope.ClusterScopeType),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetClusterGroupTOIntegrationConfig(credentialsName),
				Check: resource.ComposeTestCheckFunc(
					verifyTOIntegrationResourceCreation(provider, ClusterGroupIntegrationResourceFullName, integrationscope.ClusterGroupScopeType),
				),
			},
		},
	},
	)

	t.Log("tanzu observability integration resource acceptance test complete!")
}

func verifyTOIntegrationResourceCreation(
	provider *schema.Provider,
	resourceName string,
	integrationScope integrationscope.SupportedScopes,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("could not found resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource %s", resourceName)
		}

		if integrationScope == integrationscope.ClusterScopeType {
			fn := &clusterintegrationmodels.VmwareTanzuManageV1alpha1ClusterIntegrationFullName{
				ClusterName:           testScopeHelper.Cluster.Name,
				ManagementClusterName: testScopeHelper.Cluster.ManagementClusterName,
				ProvisionerName:       testScopeHelper.Cluster.ProvisionerName,
				Name:                  integrationschema.TanzuObservabilitySaaSValue,
			}

			resp, err := context.TMCConnection.IntegrationV2ResourceService.ClusterIntegrationResourceServiceRead(fn)

			if err != nil {
				return fmt.Errorf("cluster tanzu observability integration not found, resource: %s | err: %s", resourceName, err)
			}

			if resp == nil || resp.Integration == nil {
				return fmt.Errorf("cluster tanzu observability integration resource is empty, resource: %s", resourceName)
			}
		} else {
			fn := &clustergroupintegrationmodels.VmwareTanzuManageV1alpha1ClusterGroupIntegrationFullName{
				ClusterGroupName: testScopeHelper.ClusterGroup.Name,
				Name:             integrationschema.TanzuObservabilitySaaSValue,
			}

			resp, err := context.TMCConnection.IntegrationV2ResourceService.ClusterGroupIntegrationResourceServiceRead(fn)

			if err != nil {
				return fmt.Errorf("cluster group tanzu observability integration not found, resource: %s | err: %s", resourceName, err)
			}

			if resp == nil || resp.Integration == nil {
				return fmt.Errorf("cluster group tanzu observability integration resource is empty, resource: %s", resourceName)
			}
		}

		return nil
	}
}
