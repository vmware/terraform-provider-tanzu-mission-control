//go:build ekscluster
// +build ekscluster

/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/jarcoal/httpmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

// Function to set up HTTP mocks for the specific eks cluster/nodepool requests anticipated by this test, when not being run against a real TMC stack.
func setupHTTPMocks(t *testing.T, clusterName string) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)

	config := testhelper.TestGetDefaultEksAcceptanceConfig()
	endpoint := os.Getenv("TMC_ENDPOINT")

	// POST Cluster mock setup
	clusterSpec, nps := getMockEksClusterSpec(config.AWSAccountNumber, config.CloudFormationTemplateID)
	postRequestModel := &eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster{
		FullName: &eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName{
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

	postResponseModel := &eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster{
		FullName: &eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName{
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

	postRequest := eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest{
		EksCluster: postRequestModel,
	}

	postResponse := eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse{
		EksCluster: postResponseModel,
	}

	// GET Cluster mock setup
	readyStatus := eksmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE
	readyCondition := eksmodel.VmwareTanzuCoreV1alpha1StatusCondition{
		Type:   "ready",
		Status: &readyStatus,
	}

	readyPhase := eksmodel.VmwareTanzuManageV1alpha1EksclusterPhaseREADY
	getModel := &eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster{
		FullName: &eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName{
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
		Status: &eksmodel.VmwareTanzuManageV1alpha1EksclusterStatus{
			Phase: &readyPhase,
			Conditions: map[string]eksmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"ready": readyCondition,
			},
		},
	}

	getResponse := eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse{
		EksCluster: getModel,
	}

	listResponse := eksmodel.VmwareTanzuManageV1alpha1EksclusterListEksClustersResponse{
		EksClusters: []*eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster{getModel},
		TotalCount:  "1",
	}

	// GET Nodepools mock setup
	nodepools := make([]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool, 0)
	nodepoolRequests := make([]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest, 0)
	nodepoolResponses := make([]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse, 0)
	nodepoolReadyPhase := eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseREADY

	for count, nodepool := range nps {
		var nodepoolDescription string
		if count == 0 {
			nodepoolDescription = "tf nodepool description"
		} else {
			nodepoolDescription = fmt.Sprintf("tf nodepool %v description", count+1)
		}

		npObj := eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool{
			FullName: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName{
				CredentialName: config.CredentialName,
				EksClusterName: clusterName,
				Name:           nodepool.Info.Name,
				OrgID:          config.OrgID,
				Region:         config.Region,
			},
			Spec: nodepool.Spec,
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				Description: nodepoolDescription,
			},
		}
		npObjWithStatus := npObj

		npObjWithStatus.Status = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolStatus{
			Phase: &nodepoolReadyPhase,
			Conditions: map[string]eksmodel.VmwareTanzuCoreV1alpha1StatusCondition{
				"ready": readyCondition,
			},
		}

		nodepools = append(nodepools, &npObjWithStatus)

		nodepoolRequests = append(nodepoolRequests, &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest{
			Nodepool: &npObj,
		})

		nodepoolResponses = append(nodepoolResponses, &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse{
			Nodepool: &npObjWithStatus,
		})
	}

	getNodepoolsResponse := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse{
		Nodepools: nodepools,
	}

	// Setup HTTP Responders
	postEndpoint := fmt.Sprintf("https://%s/v1alpha1/eksclusters", endpoint)
	getClusterEndpoint := fmt.Sprintf("https://%s/v1alpha1/eksclusters/%s", endpoint, clusterName)
	listClusterEndpoint := fmt.Sprintf("https://%s/v1alpha1/eksclusters?query=uid%%3D%%22%s%%22", endpoint, postResponseModel.Meta.UID)
	postNodepoolsEndpoint := fmt.Sprintf("https://%s/v1alpha1/eksclusters/%s/nodepools", endpoint, clusterName)
	getClusterNodepoolsEndpoint := fmt.Sprintf("https://%s/v1alpha1/eksclusters/%s/nodepools", endpoint, clusterName)
	deleteEndpoint := getClusterEndpoint

	httpmock.RegisterResponder("POST", postEndpoint,
		bodyInspectingResponder(t, postRequest, 200, postResponse))

	httpmock.RegisterResponder("GET", getClusterEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))

	httpmock.RegisterResponder("GET", listClusterEndpoint,
		bodyInspectingResponder(t, nil, 200, listResponse))

	httpmock.RegisterResponder("POST", postNodepoolsEndpoint,
		nodepoolsBodyInspectingResponder(t, nodepoolRequests, 200, nodepoolResponses))

	httpmock.RegisterResponder("GET", getClusterNodepoolsEndpoint,
		bodyInspectingResponder(t, nil, 200, getNodepoolsResponse))

	httpmock.RegisterResponder("DELETE", deleteEndpoint, changeStateResponder(
		// Set up the get to return 404 after the cluster has been 'deleted'
		func() {
			httpmock.RegisterResponder("GET", getClusterEndpoint,
				httpmock.NewStringResponder(404, "Not found"))
		},
		http.StatusOK,
		nil))
}

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceName:              ResourceTMCEKSCluster(),
			clustergroup.ResourceName: clustergroup.ResourceClusterGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			ResourceName:              DataSourceTMCEKSCluster(),
			clustergroup.ResourceName: clustergroup.DataSourceClusterGroup(),
		},
		ConfigureContextFunc: getConfigureContextFunc(),
	}
	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_EKS_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_EKS_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForMkpClusterResource(t *testing.T) {
	clusterName := "terraform-eks-test"
	clusterConfig := map[string][]testhelper.TestAcceptanceOption{
		"CreateEksCluster": {
			testhelper.WithClusterName(clusterName),
			testhelper.WithEKSCluster()},
	}

	// If the flag to execute EKS tests is not found, run this as a unit test by setting up an http intercept for each endpoint
	if _, found := os.LookupEnv("ENABLE_EKS_ENV_TEST"); !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.cloud.vmware.com")

		log.Println("Setting up the mock endpoints...")
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
			"EKS_CLOUD_FORMATION_TEMPLATE_ID",
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
				Config: testGetResourceClusterDefinition(t, clusterConfig["CreateEksCluster"]...),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, clusterConfig["CreateEksCluster"]...),
				),
			},
			{
				ResourceName:      testhelper.EksClusterResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
	t.Log("cluster resource acceptance test complete!")
}

func testGetResourceClusterDefinition(t *testing.T, opts ...testhelper.TestAcceptanceOption) string {
	templateConfig := testhelper.TestGetDefaultEksAcceptanceConfig()
	for _, option := range opts {
		option(templateConfig)
	}

	switch templateConfig.AccTestType {
	case testhelper.CreateEksCluster:
		if templateConfig.KubernetesVersion == "" {
			t.Skip("KUBERNETES_VERSION env var is not set for TKGs acceptance test")
		}
	default:
		t.Skip("unknown test type")
	}

	definition, err := testhelper.Parse(templateConfig, templateConfig.TemplateData)
	if err != nil {
		t.Skipf("unable to parse cluster script: %s", err)
	}

	return definition
}

func checkResourceAttributes(provider *schema.Provider, opts ...testhelper.TestAcceptanceOption) resource.TestCheckFunc {
	testConfig := testhelper.TestGetDefaultEksAcceptanceConfig()
	for _, option := range opts {
		option(testConfig)
	}

	var check = []resource.TestCheckFunc{
		verifyEKSClusterResourceCreation(provider, testhelper.EksClusterResourceName, testConfig),
		resource.TestCheckResourceAttr(testhelper.EksClusterResourceName, "name", testConfig.Name),
		resource.TestCheckResourceAttr(testhelper.EksClusterResourceName, helper.GetFirstElementOf("spec", "cluster_group"), testConfig.ClusterGroupName),
	}

	return resource.ComposeTestCheckFunc(check...)
}

func verifyEKSClusterResourceCreation(
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

		err := getSetupConfig(&config)
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		fn := &eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName{
			Name:           testConfig.Name,
			OrgID:          testConfig.OrgID,
			Region:         testConfig.Region,
			CredentialName: testConfig.CredentialName,
		}

		resp, err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceGet(fn)
		if err != nil {
			return errors.Errorf("cluster resource not found: %s", err)
		}

		if resp == nil {
			return errors.Errorf("cluster resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}

func getMockEksClusterSpec(accountID string, templateID string) (eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) {
	controlPlaneRoleARN := fmt.Sprintf("arn:aws:iam::%s:role/control-plane.%s.eks.tmc.cloud.vmware.com", accountID, templateID)
	workerRoleArn := fmt.Sprintf("arn:aws:iam::%s:role/worker.%s.eks.tmc.cloud.vmware.com", accountID, templateID)

	return eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec{
			ClusterGroupName: "default",
			Config: &eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig{
				Version: "1.23",
				RoleArn: controlPlaneRoleARN,
				Tags: map[string]string{
					"tmc.cloud.vmware.com/tmc-managed": "true",
					"testclustertag":                   "testclustertagvalue",
					"testingtag":                       "testingtagvalue",
					"testsametag":                      "testsametagval",
				},
				KubernetesNetworkConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig{
					ServiceCidr: "10.100.0.0/16",
				},
				Logging: &eksmodel.VmwareTanzuManageV1alpha1EksclusterLogging{
					APIServer:         false,
					Audit:             true,
					Authenticator:     true,
					ControllerManager: true,
					Scheduler:         true,
				},
				Vpc: &eksmodel.VmwareTanzuManageV1alpha1EksclusterVPCConfig{
					EnablePrivateAccess: true,
					EnablePublicAccess:  true,
					PublicAccessCidrs: []string{
						"0.0.0.0/0",
					},
					SecurityGroups: []string{
						"sg-0a6768722e9716768",
					},
					SubnetIds: []string{
						"subnet-0a184f6302af32a86",
						"subnet-0ed95d5c212ac62a1",
						"subnet-0526ecaecde5b1bf7",
						"subnet-06897e1063cc0cf4e",
					},
				},
			},
		},
		[]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
			{
				Info: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{
					Name:        "first-np",
					Description: "tf nodepool description",
				},
				Spec: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec{
					RoleArn: workerRoleArn,
					AmiType: "CUSTOM",
					AmiInfo: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo{
						AmiID:                "ami-2qu8409oisdfj0qw",
						OverrideBootstrapCmd: "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
					},
					CapacityType: "ON_DEMAND",
					RootDiskSize: 40,
					Tags: map[string]string{
						"testnptag":      "testnptagvalue",
						"testingtag":     "testingnptagvalue",
						"testsametag":    "testsametagval",
						"testclustertag": "testclustertagvalue",
					},
					NodeLabels: map[string]string{
						"testnplabelkey": "testnplabelvalue",
					},
					SubnetIds: []string{
						"subnet-0a184f6302af32a86",
						"subnet-0ed95d5c212ac62a1",
						"subnet-0526ecaecde5b1bf7",
						"subnet-06897e1063cc0cf4e",
					},
					RemoteAccess: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess{
						SSHKey: "anshulc",
						SecurityGroups: []string{
							"sg-0a6768722e9716768",
						},
					},
					ScalingConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig{
						DesiredSize: 4,
						MaxSize:     8,
						MinSize:     1,
					},
					UpdateConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig{
						MaxUnavailableNodes: "2",
					},
					InstanceTypes: []string{
						"t3.medium",
						"m3.large",
					},
					Taints:         make([]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint, 0),
					LaunchTemplate: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{},
				},
			},
			{
				Info: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{
					Name:        "second-np",
					Description: "tf nodepool 2 description",
				},
				Spec: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec{
					RoleArn: workerRoleArn,
					Tags: map[string]string{
						"testnptag":      "testnptagvalue",
						"testingtag":     "testingnptagvalue",
						"testsametag":    "testsametagval",
						"testclustertag": "testclustertagvalue",
					},
					NodeLabels: map[string]string{
						"testnplabelkey": "testnplabelvalue",
					},
					SubnetIds: []string{
						"subnet-0a184f6302af32a86",
						"subnet-0ed95d5c212ac62a1",
						"subnet-0526ecaecde5b1bf7",
						"subnet-06897e1063cc0cf4e",
					},
					LaunchTemplate: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{
						Name:    "PLACE_HOLDER",
						Version: "PLACE_HOLDER",
					},
					ScalingConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig{
						DesiredSize: 4,
						MaxSize:     8,
						MinSize:     1,
					},
					UpdateConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig{
						MaxUnavailablePercentage: "12",
					},
					RemoteAccess: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess{},
					Taints: []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{
						{
							Effect: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							Key:    "randomkey",
							Value:  "randomvalue",
						},
					},
					InstanceTypes: []string{},
					AmiInfo:       &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo{},
				},
			},
		}
}

// nolint: unparam
func bodyInspectingResponder(t *testing.T, expectedContent interface{}, successResponse int, successResponseBody interface{}) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		successFunc := func() (*http.Response, error) {
			return httpmock.NewJsonResponse(successResponse, successResponseBody)
		}

		if expectedContent == nil {
			return successFunc()
		}

		// Compare to expected content.
		expectedBytes, err := json.Marshal(expectedContent)
		if err != nil {
			t.Fail()
			return nil, err
		}

		if r.Body == nil {
			t.Fail()
			return nil, fmt.Errorf("expected body on request")
		}

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fail()
			return nil, err
		}

		// Map of map of strings for comparing subnet equality
		subnetMap := make(map[string]map[string][]string, 0)

		var bodyInterface map[string]interface{}
		if err = json.Unmarshal(bodyBytes, &bodyInterface); err == nil {
			var expectedInterface map[string]interface{}

			err = json.Unmarshal(expectedBytes, &expectedInterface)
			if err != nil {
				return nil, err
			}

			diff := deep.Equal(bodyInterface, expectedInterface)
			if diff == nil {
				return successFunc()
			} else {
				// special check for subnets
				// First, populate all the diffs pertaining to subnets into maps
				for _, diffItem := range diff {
					isSubnetKey := strings.Contains(diffItem, "map[subnetIds]") && strings.Contains(diffItem, ".slice")
					isExpectedTag := strings.Contains(diffItem, "map[tags]") && strings.Contains(diffItem, "tmc.cloud.vmware.com")

					if isSubnetKey {
						segments := strings.Split(diffItem, ":")
						key := strings.Split(segments[0], ".slice")[0]

						// Create map if not present
						if subnetMap[key] == nil {
							subnetMap[key] = make(map[string][]string, 0)
						}

						// Add vals to map
						vals := strings.Split(segments[1], "!=")
						subnetMap[key]["left"] = append(subnetMap[key]["left"], strings.TrimSpace(vals[0]))
						subnetMap[key]["right"] = append(subnetMap[key]["right"], strings.TrimSpace(vals[1]))
					}

					if !(isSubnetKey || isExpectedTag) {
						t.Fail()
						return nil, errors.Errorf("diff identified outside of subnet order and additional VMware tags: %s", diffItem)
					}
				}

				// Then, sort slices and compare
				for _, set := range subnetMap {
					left := set["left"]
					right := set["right"]

					// sort
					sort.Strings(left)
					sort.Strings(right)

					subnetDiff := deep.Equal(left, right)
					if subnetDiff != nil {
						t.Fail()
						return nil, errors.New("subnets did not match")
					}
				}
			}
		} else {
			return nil, err
		}

		return successFunc()
	}
}

func nodepoolsBodyInspectingResponder(t *testing.T, nodepools []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest, successResponse int, successResponseBody []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		successFunc := func(i int) (*http.Response, error) {
			return httpmock.NewJsonResponse(successResponse, successResponseBody[i])
		}

		if r.Body == nil {
			t.Fail()
			return nil, fmt.Errorf("expected body on request")
		}

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fail()
			return nil, err
		}

		// Map of map of strings for comparing subnet equality
		subnetMap := make(map[string]map[string][]string, 0)

		var bodyInterface map[string]interface{}
		if err = json.Unmarshal(bodyBytes, &bodyInterface); err != nil {
			return nil, err
		}

		npName, err := getNodepoolNameFromJSON(bodyInterface)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get nodepool namGe")
		}

		npIdx := slices.IndexFunc(nodepools, func(npr *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest) bool {
			return npr.Nodepool.FullName.Name == npName
		})

		// Compare to expected content.
		expectedBytes, err := json.Marshal(nodepools[npIdx])
		if err != nil {
			t.Fail()
			return nil, err
		}

		var expectedInterface map[string]interface{}

		err = json.Unmarshal(expectedBytes, &expectedInterface)
		if err != nil {
			return nil, err
		}

		diff := deep.Equal(bodyInterface, expectedInterface)
		if diff == nil {
			return successFunc(npIdx)
		}

		// special check for subnets
		// First, populate all the diffs pertaining to subnets into maps
		for _, diffItem := range diff {
			isSubnetKey := strings.Contains(diffItem, "map[subnetIds]") && strings.Contains(diffItem, ".slice")
			isExpectedTag := strings.Contains(diffItem, "map[tags]") && strings.Contains(diffItem, "tmc.cloud.vmware.com")

			if isSubnetKey {
				segments := strings.Split(diffItem, ":")
				key := strings.Split(segments[0], ".slice")[0]

				// Create map if not present
				if subnetMap[key] == nil {
					subnetMap[key] = make(map[string][]string, 0)
				}

				// Add vals to map
				vals := strings.Split(segments[1], "!=")
				subnetMap[key]["left"] = append(subnetMap[key]["left"], strings.TrimSpace(vals[0]))
				subnetMap[key]["right"] = append(subnetMap[key]["right"], strings.TrimSpace(vals[1]))
			}

			if !(isSubnetKey || isExpectedTag) {
				t.Fail()
				return nil, errors.Errorf("diff identified outside of subnet order and additional VMware tags: %s", diffItem)
			}
		}

		// Then, sort slices and compare
		for _, set := range subnetMap {
			left := set["left"]
			right := set["right"]

			// sort
			sort.Strings(left)
			sort.Strings(right)

			subnetDiff := deep.Equal(left, right)
			if subnetDiff != nil {
				t.Fail()
				return nil, errors.New("subnets did not match")
			}
		}

		return successFunc(npIdx)
	}
}

func getNodepoolNameFromJSON(json map[string]interface{}) (string, error) {
	np, ok := json["nodepool"]
	if !ok {
		return "", errors.New("nodepool key not present in json")
	}

	npObj, ok := np.(map[string]interface{})
	if !ok {
		return "", errors.New("nodepool is not an object in json")
	}

	fn, ok := npObj["fullName"]
	if !ok {
		return "", errors.New("fullName key is not present in nodepool object")
	}

	fnObj, ok := fn.(map[string]interface{})
	if !ok {
		return "", errors.New("fullName is not an object in json")
	}

	name, ok := fnObj["name"]
	if !ok {
		return "", errors.New("name key is not present in fullName object")
	}

	nameStr, ok := name.(string)
	if !ok {
		return "", errors.New("name is not a string in json")
	}

	return nameStr, nil
}

// Register a new responder when the given call is made.
func changeStateResponder(registerFunc func(), successResponse int, successResponseBody interface{}) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		registerFunc()
		return httpmock.NewJsonResponse(successResponse, successResponseBody)
	}
}
