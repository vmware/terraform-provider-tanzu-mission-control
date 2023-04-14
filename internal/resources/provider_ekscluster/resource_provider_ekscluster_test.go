/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package providerekscluster

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/jarcoal/httpmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/provider_ekscluster"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForProviderEKSClusterResource(t *testing.T) {
	clusterName := "tf-eks-attach-test"
	clusterConfig := map[string][]testhelper.TestAcceptanceOption{
		"CreateProviderEksCluster": {
			testhelper.WithClusterName(clusterName),
			testhelper.WithProviderEKSCluster()},
	}

	// If the flag to execute EKS tests is not found, run this as a unit test by setting up an http intercept for each endpoint
	if val, found := os.LookupEnv(testhelper.EKSMockEnv); !found || val == "" {
		setupHTTPMocks(t, clusterName)
	} else {
		// Environment variables with non default values required for a successful call to MKP
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
			"EKS_ORG_ID",
			"EKS_AWS_ACCOUNT_NUMBER",
			"EKS_CREDENTIAL_NAME",
		}

		// Check if the required environment variables are set
		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	var provider = initTestProvider(t)

	resource.Test(t, resource.TestCase{
		IsUnitTest:        true,
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		Steps: []resource.TestStep{
			{
				Config: testGetResourceClusterDefinition(t, clusterConfig["CreateProviderEksCluster"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["CreateProviderEksCluster"]...),
				),
			},
		},
	})
	t.Log("cluster resource acceptance test complete!")
}

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceName:              ResourceTMCProviderEKSCluster(),
			clustergroup.ResourceName: clustergroup.ResourceClusterGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			ResourceName:              DataSourceTMCProviderEKSCluster(),
			clustergroup.ResourceName: clustergroup.DataSourceClusterGroup(),
		},
		ConfigureContextFunc: testhelper.GetConfigureContextFunc(),
	}
	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}

func testGetResourceClusterDefinition(t *testing.T, opts ...testhelper.TestAcceptanceOption) string {
	templateConfig := testhelper.TestGetDefaultProviderEksAcceptanceConfig()
	for _, option := range opts {
		option(templateConfig)
	}

	if templateConfig.AccTestType != testhelper.CreateProviderEksCluster {
		t.Skipf("unknown test type %v", templateConfig.AccTestType)
	}

	definition, err := testhelper.Parse(templateConfig, templateConfig.TemplateData)
	if err != nil {
		t.Skipf("unable to parse provider eks script: %s", err)
	}

	return definition
}

func checkResourceAttributes(provider *schema.Provider, opts ...testhelper.TestAcceptanceOption) resource.TestCheckFunc {
	testConfig := testhelper.TestGetDefaultProviderEksAcceptanceConfig()
	for _, option := range opts {
		option(testConfig)
	}

	var check = []resource.TestCheckFunc{
		verifyClusterResourceCreation(provider, testhelper.ProviderEksClusterResourceName, testConfig),
		resource.TestCheckResourceAttr(testhelper.ProviderEksClusterResourceName, "name", testConfig.Name),
		resource.TestCheckResourceAttr(testhelper.ProviderEksClusterResourceName, helper.GetFirstElementOf("spec", "cluster_group"), testConfig.ClusterGroupName),
	}

	return resource.ComposeTestCheckFunc(check...)
}

// Function to set up HTTP mocks for the specific eks cluster/nodepool requests anticipated by this test, when not being run against a real TMC stack.
func setupHTTPMocks(t *testing.T, clusterName string) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)

	config := testhelper.TestGetDefaultProviderEksAcceptanceConfig()
	endpoint := os.Getenv("TMC_ENDPOINT")

	// POST Cluster mock setup
	clusterSpec := getMockProviderClusterSpec(config.AWSAccountNumber, clusterName)
	putRequestModel := &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster{
		FullName: &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName{
			Name:           clusterName,
			CredentialName: config.CredentialName,
			Region:         config.Region,
		},
		Spec: &clusterSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: nil,
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}

	reference := objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference{
		Rid: "test_rid",
		UID: "test_uid",
	}
	referenceArray := make([]*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference, 0)
	referenceArray = append(referenceArray, &reference)

	putResponseModel := &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster{
		FullName: &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName{
			OrgID:          config.OrgID,
			Name:           clusterName,
			CredentialName: config.CredentialName,
			Region:         config.Region,
		},
		Spec: &clusterSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: referenceArray,
			UID:              "1886ad24-40bb-4517-9712-af9df737b606",
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}

	putRequest := models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterRequest{
		ProviderEksCluster: putRequestModel,
	}

	putResponse := models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterResponse{
		ProviderEksCluster: putResponseModel,
	}

	// GET Cluster mock setup
	readyStatus := statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE
	readyCondition := statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
		Type:   "ready",
		Status: &readyStatus,
	}

	readyPhase := models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseMANAGED
	getModel := &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster{
		FullName: &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName{
			Name:           clusterName,
			CredentialName: config.CredentialName,
			Region:         config.Region,
		},
		Spec: &clusterSpec,
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			ParentReferences: referenceArray,
			UID:              "1886ad24-40bb-4517-9712-af9df737b606",
			Description:      "resource with description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		Status: &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterStatus{
			Phase: &readyPhase,
			Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"ready": readyCondition,
			},
		},
	}

	getResponse := models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse{
		ProviderEksCluster: getModel,
	}

	// Setup HTTP Responders
	clusterEndpoint := fmt.Sprintf("https://%s/v1alpha1/manage/providereksclusters/%s", endpoint, clusterName)

	httpmock.RegisterResponder("PUT", clusterEndpoint,
		func(r *http.Request) (*http.Response, error) {
			resp, err := testhelper.BodyInspectingResponder(t, putRequest, 200, putResponse, testhelper.MetaUIDIgnoreFunc)(r)

			// update it for delete
			putRequest.ProviderEksCluster.Spec.TmcManaged = false
			httpmock.RegisterResponder("PUT", clusterEndpoint,
				testhelper.BodyInspectingResponder(t, putRequest, 200, putResponse, testhelper.MetaUIDIgnoreFunc))

			return resp, err
		})
	//

	httpmock.RegisterResponder("GET", clusterEndpoint,
		testhelper.BodyInspectingResponder(t, nil, 200, getResponse))
}

func verifyClusterResourceCreation(
	provider *schema.Provider,
	resourceName string,
	testConfig *testhelper.TestAcceptanceConfig,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return errors.New("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return errors.Errorf("not found resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return errors.Errorf("ID not set, resource %s", resourceName)
		}

		config := authctx.TanzuContext{
			ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
			Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
			VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
			TLSConfig:        &proxy.TLSConfig{},
		}

		err := testhelper.GetSetupConfig(&config)
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		fn := &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName{
			Name:           testConfig.Name,
			OrgID:          testConfig.OrgID,
			Region:         testConfig.Region,
			CredentialName: testConfig.CredentialName,
		}

		resp, err := config.TMCConnection.ProviderEKSClusterResourceService.ProviderEksClusterResourceServiceGet(fn)
		if err != nil {
			return errors.Errorf("cluster resource not found: %s", err)
		}

		if resp == nil {
			return errors.Errorf("cluster resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}

func getMockProviderClusterSpec(accountID, clusterName string) models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec {
	return models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec{
		AgentName:        "tf-test-cluster",
		Arn:              fmt.Sprintf("arn:aws:eks:us-west-2:%s:cluster/%s", accountID, clusterName),
		ClusterGroupName: "default",
		TmcManaged:       true,
	}
}
