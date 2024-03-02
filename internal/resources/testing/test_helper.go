/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package testing

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type AcceptanceTestType int

const (
	AttachClusterType AcceptanceTestType = iota
	AttachClusterTypeWithKubeConfig
	AttachClusterTypeWithKubeconfigImageRegistry
	AttachClusterTypeWithKubeconfigProxy
	TkgAWSCluster
	TkgsCluster
	TkgVsphereCluster
	CreateEksCluster
)

const (
	ClusterResource      = "tanzu-mission-control_cluster"
	ClusterResourceVar   = "test_attach_cluster"
	ClusterDataSourceVar = "test_data_attach_cluster"

	// EKS Constants.
	EksClusterResource    = "tanzu-mission-control_ekscluster"
	EksClusterGroup       = "tanzu-mission-control_cluster_group"
	EksClusterResourceVar = "test_create_eks_cluster"
	EksClusterGroupVar    = "test_create_eks_cluster_group"
)

var (
	ClusterResourceName   = fmt.Sprintf("%s.%s", ClusterResource, ClusterResourceVar)
	ClusterDataSourceName = fmt.Sprintf("data.%s.%s", ClusterResource, ClusterDataSourceVar)

	EksClusterResourceName = fmt.Sprintf("%s.%s", EksClusterResource, EksClusterResourceVar)
	EksClusterGroupName    = fmt.Sprintf("data.%s.%s", EksClusterGroup, EksClusterGroupVar)
)

type TestAcceptanceOption func(config *TestAcceptanceConfig)

type TestAcceptanceConfig struct {
	ResourceName          string
	ResourceNameVar       string
	DataSourceNameVar     string
	Name                  string
	KubeConfigPath        string
	Meta                  string
	AccTestType           AcceptanceTestType
	TemplateData          string
	ManagementClusterName string
	ProvisionerName       string
	Version               string
	StorageClass          string
	ControlPlaneEndPoint  string
	// EKS
	KubernetesVersion        string
	Region                   string
	AWSAccountNumber         string
	CloudFormationTemplateID string
	CredentialName           string
	OrgID                    string
	ClusterGroupName         string
	ImageRegistry            string
	Proxy                    string
}

func WithClusterName(name string) TestAcceptanceOption {
	return func(config *TestAcceptanceConfig) {
		config.Name = name
	}
}

func WithTKGmAWSCluster() TestAcceptanceOption {
	return func(config *TestAcceptanceConfig) {
		config.ManagementClusterName = os.Getenv("TKGM_AWS_MANAGEMENT_CLUSTER")
		config.ProvisionerName = os.Getenv("TKGM_AWS_PROVISIONER_NAME")
		config.AccTestType = TkgAWSCluster
		config.TemplateData = testTKGmAWSClusterScript
	}
}

func WithTKGsCluster() TestAcceptanceOption {
	return func(config *TestAcceptanceConfig) {
		config.ManagementClusterName = os.Getenv("TKGS_MANAGEMENT_CLUSTER")
		config.ProvisionerName = os.Getenv("TKGS_PROVISIONER_NAME")
		config.Version = os.Getenv("VERSION")
		config.StorageClass = os.Getenv("STORAGE_CLASS")
		config.AccTestType = TkgsCluster
		config.TemplateData = testTKGsClusterScript
	}
}

func WithTKGmVsphereCluster() TestAcceptanceOption {
	return func(config *TestAcceptanceConfig) {
		config.ManagementClusterName = os.Getenv("TKGM_VSPHERE_MANAGEMENT_CLUSTER")
		config.ProvisionerName = os.Getenv("TKGM_VSPHERE_PROVISIONER_NAME")
		config.AccTestType = TkgVsphereCluster
		config.TemplateData = testTKGmVsphereClusterScript
	}
}

func WithEKSCluster() TestAcceptanceOption {
	return func(config *TestAcceptanceConfig) {
		// Only read environment variables into config if the test is configured to run against a real environment without mocks
		if _, found := os.LookupEnv("ENABLE_EKS_ENV_TEST"); found {
			if val, exists := os.LookupEnv("EKS_ORG_ID"); exists {
				config.OrgID = val
			}

			if val, exists := os.LookupEnv("EKS_AWS_ACCOUNT_NUMBER"); exists {
				config.AWSAccountNumber = val
			}

			if val, exists := os.LookupEnv("EKS_AWS_REGION"); exists {
				config.Region = val
			}

			if val, exists := os.LookupEnv("EKS_CLUSTER_GROUP_NAME"); exists {
				config.ClusterGroupName = val
			}

			if val, exists := os.LookupEnv("EKS_KUBERNETES_VERSION"); exists {
				config.KubernetesVersion = val
			}

			if val, exists := os.LookupEnv("EKS_CREDENTIAL_NAME"); exists {
				config.CredentialName = val
			}

			if val, exists := os.LookupEnv("EKS_CLOUD_FORMATION_TEMPLATE_ID"); exists {
				config.CloudFormationTemplateID = val
			}
		}

		config.AccTestType = CreateEksCluster
		config.TemplateData = testDefaultCreateEksClusterScript
	}
}

func WithKubeConfig() TestAcceptanceOption {
	return func(config *TestAcceptanceConfig) {
		config.KubeConfigPath = os.Getenv("KUBECONFIG")
		config.AccTestType = AttachClusterTypeWithKubeConfig
		config.TemplateData = testAttachClusterWithKubeConfigScript
	}
}

func WithKubeConfigImageRegistry() TestAcceptanceOption {
	return func(config *TestAcceptanceConfig) {
		config.KubeConfigPath = os.Getenv("ATTACH_WITH_IMAGE_REGISTRY_KUBECONFIG")
		config.ImageRegistry = os.Getenv("IMAGE_REGISTRY")
		config.AccTestType = AttachClusterTypeWithKubeconfigImageRegistry
		config.TemplateData = testAttachClusterWithKubeConfigScriptImageRegistry
	}
}

func WithKubeConfigProxy() TestAcceptanceOption {
	return func(config *TestAcceptanceConfig) {
		config.KubeConfigPath = os.Getenv("ATTACH_WITH_PROXY_KUBECONFIG")
		config.Proxy = os.Getenv("PROXY")
		config.AccTestType = AttachClusterTypeWithKubeconfigProxy
		config.TemplateData = testAttachClusterWithKubeConfigScriptProxy
	}
}

func WithDataSourceScript() TestAcceptanceOption {
	return func(config *TestAcceptanceConfig) {
		config.TemplateData = testDataSourceAttachClusterScript
		config.DataSourceNameVar = ClusterDataSourceVar
	}
}

func TestGetDefaultAcceptanceConfig() *TestAcceptanceConfig {
	return &TestAcceptanceConfig{
		ResourceName:          ClusterResource,
		ResourceNameVar:       ClusterResourceVar,
		Meta:                  MetaTemplate,
		AccTestType:           AttachClusterType,
		TemplateData:          testDefaultAttachClusterScript,
		ManagementClusterName: "attached",
		ProvisionerName:       "attached",
	}
}

func TestGetDefaultEksAcceptanceConfig() *TestAcceptanceConfig {
	return &TestAcceptanceConfig{
		ResourceName:             EksClusterResource,
		ResourceNameVar:          EksClusterResourceVar,
		Meta:                     MetaTemplate,
		AccTestType:              CreateEksCluster,
		TemplateData:             testDefaultCreateEksClusterScript,
		OrgID:                    "bc27608b-4809-4cac-9e04-778803963da2",
		AWSAccountNumber:         "919197287370",
		Region:                   "us-west-2",
		ClusterGroupName:         "default",
		KubernetesVersion:        "1.26",
		CredentialName:           "PLACE_HOLDER",
		CloudFormationTemplateID: "PLACE_HOLDER",
	}
}

func Parse(m interface{}, objects string) (string, error) {
	var definitionBytes bytes.Buffer

	t := template.Must(template.New("script").Parse(objects))
	if err := t.Execute(&definitionBytes, m); err != nil {
		return "", err
	}

	return definitionBytes.String(), nil
}
