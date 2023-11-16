/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dataprotection2 "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/backupschedule"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/integration"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/nodepools"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/credential"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/ekscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmcharts"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmfeature"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrelease"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmrepository"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/iampolicy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/namespace"
	tanzupackage "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package"
	tanzupackages "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/packages"
	custompolicy "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom"
	custompolicyresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom/resource"
	imagepolicy "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image"
	imagepolicyresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image/resource"
	mutationpolicy "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/mutation"
	mutationpolicyresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/mutation/resource"
	networkpolicy "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network"
	networkpolicyresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network/resource"
	quotapolicy "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota"
	quotapolicyresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota/resource"
	securitypolicy "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security"
	securitypolicyresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security/resource"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/sourcesecret"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall"
	packagerepository "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/targetlocation"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

// Provider for Tanzu Mission Control resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			cluster.ResourceName:             cluster.ResourceTMCCluster(),
			ekscluster.ResourceName:          ekscluster.ResourceTMCEKSCluster(),
			akscluster.ResourceName:          akscluster.ResourceTMCAKSCluster(),
			workspace.ResourceName:           workspace.ResourceWorkspace(),
			namespace.ResourceName:           namespace.ResourceNamespace(),
			clustergroup.ResourceName:        clustergroup.ResourceClusterGroup(),
			nodepools.ResourceName:           nodepools.ResourceNodePool(),
			iampolicy.ResourceName:           iampolicy.ResourceIAMPolicy(),
			custompolicy.ResourceName:        custompolicyresource.ResourceCustomPolicy(),
			securitypolicy.ResourceName:      securitypolicyresource.ResourceSecurityPolicy(),
			imagepolicy.ResourceName:         imagepolicyresource.ResourceImagePolicy(),
			quotapolicy.ResourceName:         quotapolicyresource.ResourceQuotaPolicy(),
			networkpolicy.ResourceName:       networkpolicyresource.ResourceNetworkPolicy(),
			credential.ResourceName:          credential.ResourceCredential(),
			integration.ResourceName:         integration.ResourceIntegration(),
			gitrepository.ResourceName:       gitrepository.ResourceGitRepository(),
			kustomization.ResourceName:       kustomization.ResourceKustomization(),
			sourcesecret.ResourceName:        sourcesecret.ResourceSourceSecret(),
			packagerepository.ResourceName:   packagerepository.ResourcePackageRepository(),
			tanzupackageinstall.ResourceName: tanzupackageinstall.ResourcePackageInstall(),
			kubernetessecret.ResourceName:    kubernetessecret.ResourceSecret(),
			mutationpolicy.ResourceName:      mutationpolicyresource.ResourceMutationPolicy(),
			helmrelease.ResourceName:         helmrelease.ResourceHelmRelease(),
			helmfeature.ResourceName:         helmfeature.ResourceHelm(),
			backupschedule.ResourceName:      backupschedule.ResourceBackupSchedule(),
			dataprotection2.ResourceName:     dataprotection2.ResourceEnableDataProtection(),
			targetlocation.ResourceName:      targetlocation.ResourceTargetLocation(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			cluster.ResourceName:             cluster.DataSourceTMCCluster(),
			ekscluster.ResourceName:          ekscluster.DataSourceTMCEKSCluster(),
			akscluster.ResourceName:          akscluster.DataSourceTMCAKSCluster(),
			workspace.ResourceName:           workspace.DataSourceWorkspace(),
			namespace.ResourceName:           namespace.DataSourceNamespace(),
			clustergroup.ResourceName:        clustergroup.DataSourceClusterGroup(),
			nodepools.ResourceName:           nodepools.DataSourceClusterNodePool(),
			credential.ResourceName:          credential.DataSourceCredential(),
			integration.ResourceName:         integration.DataSourceIntegration(),
			gitrepository.ResourceName:       gitrepository.DataSourceGitRepository(),
			sourcesecret.ResourceName:        sourcesecret.DataSourceSourcesecret(),
			packagerepository.ResourceName:   packagerepository.DataSourcePackageRepository(),
			tanzupackage.ResourceName:        tanzupackage.DataSourceTanzuPackage(),
			tanzupackages.ResourceName:       tanzupackages.DataSourceTanzuPackages(),
			tanzupackageinstall.ResourceName: tanzupackageinstall.DataSourcePackageInstall(),
			kubernetessecret.ResourceName:    kubernetessecret.DataSourceSecret(),
			helmfeature.ResourceName:         helmfeature.DataSourceHelm(),
			helmcharts.ResourceName:          helmcharts.DataSourceHelmCharts(),
			helmrepository.ResourceName:      helmrepository.DataSourceHelmRepository(),
			backupschedule.ResourceName:      backupschedule.DataSourceBackupSchedule(),
			targetlocation.ResourceName:      targetlocation.DataSourceTargetLocations(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}
}
