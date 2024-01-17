/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package client

import (
	"net/http"
	"runtime"

	aksclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/akscluster"
	aksnodepoolclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/akscluster/nodepool"
	clusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster"
	backupscheduleclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/backupschedule"
	continuousdeliveryclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/continuousdelivery"
	dataprotectionclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/dataprotection"
	gitrepositoryclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/gitrepository"
	helmfeatureclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/helmfeature"
	helmreleaseclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/helmrelease"
	helmrepositoryclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/helmrepository"
	iamclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/iam_policy"
	kustomizationclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/kustomization"
	manifestclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/manifest"
	packageclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/package"
	policyclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/policy"
	sourcesecretclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/sourcesecret"
	clusterclassclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clusterclass"
	clustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup"
	continuousdeliveryclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/continuousdelivery"
	gitrepositoryclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/gitrepository"
	helmfeatureclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/helmfeature"
	helmreleaseclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/helmrelease"
	iamclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/iam_policy"
	secretclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/kubernetessecret"
	secretexportclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/kubernetessecret/secretexport"
	kustomizationclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/kustomization"
	policyclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/policy"
	sourcesecretclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/sourcesecret"
	credentialclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/credential"
	custompolicytemplateclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/custompolicytemplate"
	eksclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/ekscluster"
	eksnodepoolclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/ekscluster/nodepool"
	inspectionsclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/inspections"
	integrationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/integration"
	kubeconfigclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/kubeconfig"
	secretclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/kubernetessecret"
	secretexportclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/kubernetessecret/secretexport"
	managementclusterregistrationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/managementclusterregistration"
	namespaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/namespace"
	iamnamespaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/namespace/iam_policy"
	nodepoolclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/nodepool"
	helmchartsorgclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/organization/helmcharts"
	iamorganizationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/organization/iam_policy"
	policyorganizationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/organization/policy"
	provisionerclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/provisioner"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	recipeclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/recipe"
	tanzukubernetesclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/tanzukubernetescluster"
	tanzupackageclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/tanzupackage"
	pkginstallclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/tanzupackageinstall"
	pkgrepositoryclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/tanzupackagerepository"
	pkgrepoavailabilityclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/tanzupackagerepository/packagerepositoryavailability"
	targetlocationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/targetlocation"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	workspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace"
	iamworkspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace/iam_policy"
	policyworkspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace/policy"
)

// NewHTTPClient creates a new tanzu mission control HTTP client.
func NewHTTPClient(config *proxy.TLSConfig) (*TanzuMissionControl, error) {
	httpClient, err := transport.NewClient(config)
	if err != nil {
		return nil, err
	}

	return newHTTPClient(httpClient), nil
}

// NewTestHTTPClientWithDefaultTransport is intended primarily for testing only, as httpmock requires a default transport object be used in order to intercept and mock traffic.
func NewTestHTTPClientWithDefaultTransport() *TanzuMissionControl {
	return newHTTPClient(transport.NewClientWithDefaultTransport())
}

func newHTTPClient(httpClient *transport.Client) *TanzuMissionControl {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Connection", "keep-alive")
	headers.Set("x-client-name", "tmc-terraform-provider")
	headers.Set("x-client-platform", runtime.GOOS)
	headers.Set("x-client-version", "1.0.2")

	httpClient.AddHeaders(headers)

	return &TanzuMissionControl{
		Client:                                        httpClient,
		ClusterResourceService:                        clusterclient.New(httpClient),
		EKSClusterResourceService:                     eksclusterclient.New(httpClient),
		EKSNodePoolResourceService:                    eksnodepoolclient.New(httpClient),
		AKSClusterResourceService:                     aksclusterclient.New(httpClient),
		AKSNodePoolResourceService:                    aksnodepoolclient.New(httpClient),
		WorkspaceResourceService:                      workspaceclient.New(httpClient),
		NamespaceResourceService:                      namespaceclient.New(httpClient),
		ClusterGroupResourceService:                   clustergroupclient.New(httpClient),
		NodePoolResourceService:                       nodepoolclient.New(httpClient),
		OrganizationIAMResourceService:                iamorganizationclient.New(httpClient),
		ClusterGroupIAMResourceService:                iamclustergroupclient.New(httpClient),
		ClusterIAMResourceService:                     iamclusterclient.New(httpClient),
		WorkspaceIAMResourceService:                   iamworkspaceclient.New(httpClient),
		NamespaceIAMResourceService:                   iamnamespaceclient.New(httpClient),
		ClusterPolicyResourceService:                  policyclusterclient.New(httpClient),
		ClusterGroupPolicyResourceService:             policyclustergroupclient.New(httpClient),
		WorkspacePolicyResourceService:                policyworkspaceclient.New(httpClient),
		OrganizationPolicyResourceService:             policyorganizationclient.New(httpClient),
		CredentialResourceService:                     credentialclient.New(httpClient),
		IntegrationResourceService:                    integrationclient.New(httpClient),
		ClusterContinuousDeliveryResourceService:      continuousdeliveryclusterclient.New(httpClient),
		ClusterGitRepositoryResourceService:           gitrepositoryclusterclient.New(httpClient),
		ClusterKustomizationResourceService:           kustomizationclusterclient.New(httpClient),
		ClusterGroupContinuousDeliveryResourceService: continuousdeliveryclustergroupclient.New(httpClient),
		ClusterGroupGitRepositoryResourceService:      gitrepositoryclustergroupclient.New(httpClient),
		ClusterGroupKustomizationResourceService:      kustomizationclustergroupclient.New(httpClient),
		ClusterSourcesecretResourceService:            sourcesecretclusterclient.New(httpClient),
		ClusterGroupSourcesecretResourceService:       sourcesecretclustergroupclient.New(httpClient),
		SecretResourceService:                         secretclient.New(httpClient),
		SecretExportResourceService:                   secretexportclient.New(httpClient),
		ManifestResourceService:                       manifestclient.New(httpClient),
		ClusterPackageRepositoryService:               pkgrepositoryclusterclient.New(httpClient),
		ClusterPackageRepositoryAvailabilityService:   pkgrepoavailabilityclusterclient.New(httpClient),
		ClusterTanzuPackageService:                    tanzupackageclusterclient.New(httpClient),
		TanzupackageResourceService:                   packageclusterclient.New(httpClient),
		PackageInstallResourceService:                 pkginstallclusterclient.New(httpClient),
		ClusterHelmReleaseResourceService:             helmreleaseclusterclient.New(httpClient),
		ClusterGroupHelmReleaseResourceService:        helmreleaseclustergroupclient.New(httpClient),
		ClusterHelmResourceService:                    helmfeatureclusterclient.New(httpClient),
		ClusterGroupHelmResourceService:               helmfeatureclustergroupclient.New(httpClient),
		ClusterHelmRepositoryResourceService:          helmrepositoryclusterclient.New(httpClient),
		OrganizationHelmChartsResourceService:         helmchartsorgclient.New(httpClient),
		ClusterGroupSecretResourceService:             secretclustergroupclient.New(httpClient),
		ClusterGroupSecretExportResourceService:       secretexportclustergroupclient.New(httpClient),
		KubeConfigResourceService:                     kubeconfigclient.New(httpClient),
		InspectionsResourceService:                    inspectionsclient.New(httpClient),
		BackupScheduleService:                         backupscheduleclient.New(httpClient),
		DataProtectionService:                         dataprotectionclient.New(httpClient),
		TargetLocationService:                         targetlocationclient.New(httpClient),
		ManagementClusterRegistrationResourceService:  managementclusterregistrationclient.New(httpClient),
		ClusterClassResourceService:                   clusterclassclient.New(httpClient),
		TanzuKubernetesClusterResourceService:         tanzukubernetesclusterclient.New(httpClient),
		ProvisionerResourceService:                    provisionerclient.New(httpClient),
		CustomPolicyTemplateResourceService:           custompolicytemplateclient.New(httpClient),
		RecipeResourceService:                         recipeclient.New(httpClient),
	}
}

// TanzuMissionControl is a client for tanzu mission control.
type TanzuMissionControl struct {
	*transport.Client
	ClusterResourceService                        clusterclient.ClientService
	EKSClusterResourceService                     eksclusterclient.ClientService
	EKSNodePoolResourceService                    eksnodepoolclient.ClientService
	AKSClusterResourceService                     aksclusterclient.ClientService
	AKSNodePoolResourceService                    aksnodepoolclient.ClientService
	WorkspaceResourceService                      workspaceclient.ClientService
	NamespaceResourceService                      namespaceclient.ClientService
	ClusterGroupResourceService                   clustergroupclient.ClientService
	NodePoolResourceService                       nodepoolclient.ClientService
	OrganizationIAMResourceService                iamorganizationclient.ClientService
	ClusterGroupIAMResourceService                iamclustergroupclient.ClientService
	ClusterIAMResourceService                     iamclusterclient.ClientService
	WorkspaceIAMResourceService                   iamworkspaceclient.ClientService
	NamespaceIAMResourceService                   iamnamespaceclient.ClientService
	ClusterPolicyResourceService                  policyclusterclient.ClientService
	ClusterGroupPolicyResourceService             policyclustergroupclient.ClientService
	WorkspacePolicyResourceService                policyworkspaceclient.ClientService
	OrganizationPolicyResourceService             policyorganizationclient.ClientService
	CredentialResourceService                     credentialclient.ClientService
	IntegrationResourceService                    integrationclient.ClientService
	ClusterContinuousDeliveryResourceService      continuousdeliveryclusterclient.ClientService
	ClusterGitRepositoryResourceService           gitrepositoryclusterclient.ClientService
	ClusterKustomizationResourceService           kustomizationclusterclient.ClientService
	ClusterGroupContinuousDeliveryResourceService continuousdeliveryclustergroupclient.ClientService
	ClusterGroupGitRepositoryResourceService      gitrepositoryclustergroupclient.ClientService
	ClusterGroupKustomizationResourceService      kustomizationclustergroupclient.ClientService
	ClusterSourcesecretResourceService            sourcesecretclusterclient.ClientService
	ClusterGroupSourcesecretResourceService       sourcesecretclustergroupclient.ClientService
	SecretResourceService                         secretclient.ClientService
	SecretExportResourceService                   secretexportclient.ClientService
	ManifestResourceService                       manifestclient.ClientService
	ClusterPackageRepositoryService               pkgrepositoryclusterclient.ClientService
	ClusterPackageRepositoryAvailabilityService   pkgrepoavailabilityclusterclient.ClientService
	ClusterTanzuPackageService                    tanzupackageclusterclient.ClientService
	TanzupackageResourceService                   packageclusterclient.ClientService
	PackageInstallResourceService                 pkginstallclusterclient.ClientService
	ClusterGroupHelmReleaseResourceService        helmreleaseclustergroupclient.ClientService
	ClusterHelmReleaseResourceService             helmreleaseclusterclient.ClientService
	ClusterHelmResourceService                    helmfeatureclusterclient.ClientService
	ClusterGroupHelmResourceService               helmfeatureclustergroupclient.ClientService
	ClusterHelmRepositoryResourceService          helmrepositoryclusterclient.ClientService
	OrganizationHelmChartsResourceService         helmchartsorgclient.ClientService
	ClusterGroupSecretResourceService             secretclustergroupclient.ClientService
	ClusterGroupSecretExportResourceService       secretexportclustergroupclient.ClientService
	KubeConfigResourceService                     kubeconfigclient.ClientService
	BackupScheduleService                         backupscheduleclient.ClientService
	DataProtectionService                         dataprotectionclient.ClientService
	TargetLocationService                         targetlocationclient.ClientService
	ManagementClusterRegistrationResourceService  managementclusterregistrationclient.ClientService
	ClusterClassResourceService                   clusterclassclient.ClientService
	TanzuKubernetesClusterResourceService         tanzukubernetesclusterclient.ClientService
	ProvisionerResourceService                    provisionerclient.ClientService
	InspectionsResourceService                    inspectionsclient.ClientService
	CustomPolicyTemplateResourceService           custompolicytemplateclient.ClientService
	RecipeResourceService                         recipeclient.ClientService
}
