/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func ConstructCluster(data *schema.ResourceData) (*models.VmwareTanzuManageV1alpha1AksCluster, error) {
	spec, err := constructAKSClusterSpec(data)
	if err != nil {
		return nil, err
	}

	return &models.VmwareTanzuManageV1alpha1AksCluster{
		FullName: extractClusterFullName(data),
		Meta:     common.ConstructMeta(data),
		Spec:     spec,
	}, nil
}

func constructAKSClusterSpec(data *schema.ResourceData) (*models.VmwareTanzuManageV1alpha1AksclusterSpec, error) {
	specData := extractClusterSpec(data)

	spec := &models.VmwareTanzuManageV1alpha1AksclusterSpec{}
	if v, ok := specData[clusterGroupKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ClusterGroupName, clusterGroupKey)
	}

	if v, ok := specData[proxyNameKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ProxyName, proxyNameKey)
	}

	if v, ok := specData[configKey]; ok {
		configData, _ := v.([]any)
		v, err := constructConfig(configData)

		if err != nil {
			return nil, err
		} else {
			spec.Config = v
		}
	}

	if v, ok := specData[agentNameKey]; ok {
		helper.SetPrimitiveValue(v, &spec.AgentName, agentNameKey)
	}

	if v, ok := specData[resourceIDKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ResourceID, resourceIDKey)
	}

	return spec, nil
}

func extractClusterSpec(data *schema.ResourceData) map[string]any {
	value, ok := data.GetOk(clusterSpecKey)
	if !ok {
		return nil
	}

	dataSpec := value.([]any)
	if len(dataSpec) < 1 {
		return nil
	}

	// Spec schema defines max 1
	return dataSpec[0].(map[string]any)
}

func constructConfig(data []any) (*models.VmwareTanzuManageV1alpha1AksclusterClusterConfig, error) {
	if len(data) < 1 {
		return nil, nil
	}

	// Config schema defines max 1
	configData, _ := data[0].(map[string]any)
	config := &models.VmwareTanzuManageV1alpha1AksclusterClusterConfig{}

	if v, ok := configData[tagsKey]; ok {
		data, _ := v.(map[string]any)
		config.Tags = constructStringMap[string](data)
	}

	if v, ok := configData[accessConfigKey]; ok {
		data, _ := v.([]any)
		config.AccessConfig = constructAccessConfig(data)
	}

	if v, ok := configData[addonsConfigKey]; ok {
		data, _ := v.([]any)
		config.AddonsConfig = constructAddonsConfig(data)
	}

	if v, ok := configData[apiServerAccessConfigKey]; ok {
		data, _ := v.([]any)
		config.APIServerAccessConfig = constructAPIServerAccessConfig(data)
	}

	if v, ok := configData[autoUpgradeConfigKey]; ok {
		data, _ := v.([]any)
		config.AutoUpgradeConfig = constructAutoUpgradeConfig(data)
	}

	if v, ok := configData[diskEncryptionSetKey]; ok {
		helper.SetPrimitiveValue(v, &config.DiskEncryptionSetID, diskEncryptionSetKey)
	}

	if v, ok := configData[kubernetesVersionKey]; ok {
		helper.SetPrimitiveValue(v, &config.Version, kubernetesVersionKey)
	}

	if v, ok := configData[nodeResourceGroupNameKey]; ok {
		helper.SetPrimitiveValue(v, &config.NodeResourceGroupName, nodeResourceGroupNameKey)
	}

	if v, ok := configData[linuxConfigKey]; ok {
		data, _ := v.([]any)
		config.LinuxConfig = constructLinuxConfig(data)
	}

	if v, ok := configData[locationKey]; ok {
		helper.SetPrimitiveValue(v, &config.Location, locationKey)
	}

	if v, ok := configData[networkConfigKey]; ok {
		data, _ := v.([]any)
		config.NetworkConfig = constructNetworkConfig(data)
	}

	if v, ok := configData[skuKey]; ok {
		data, _ := v.([]any)
		config.Sku = constructSku(data)
	}

	if v, ok := configData[storageConfigKey]; ok {
		data, _ := v.([]any)
		config.StorageConfig = constructStorageConfig(data)
	}

	if v, ok := configData[nodeResourceGroupNameKey]; ok {
		helper.SetPrimitiveValue(v, &config.NodeResourceGroupName, nodeResourceGroupNameKey)
	}

	return config, nil
}

func constructSku(data []any) *models.VmwareTanzuManageV1alpha1AksclusterClusterSKU {
	if len(data) < 1 {
		return nil
	}

	// SKU schema defines max 1
	skuData, _ := data[0].(map[string]any)
	sku := &models.VmwareTanzuManageV1alpha1AksclusterClusterSKU{}

	if v, ok := skuData[skuNameKey]; ok {
		name := models.VmwareTanzuManageV1alpha1AksclusterClusterSKUName(v.(string))
		sku.Name = &name
	}

	if v, ok := skuData[skuTierKey]; ok {
		tier := models.VmwareTanzuManageV1alpha1AksclusterTier(v.(string))
		sku.Tier = &tier
	}

	return sku
}

func constructAccessConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterAccessConfig {
	if len(data) < 1 {
		return nil
	}

	// AccessConfig schema defines max 1
	accessConfigData, _ := data[0].(map[string]any)
	accessConfig := &models.VmwareTanzuManageV1alpha1AksclusterAccessConfig{}

	if v, ok := accessConfigData[aadConfigKey]; ok {
		data, _ := v.([]any)
		accessConfig.AadConfig = constructAadConfig(data)
	}

	if v, ok := accessConfigData[disableLocalAccountsKey]; ok {
		helper.SetPrimitiveValue(v, &accessConfig.DisableLocalAccounts, disableLocalAccountsKey)
	}

	if v, ok := accessConfigData[enableRbacKey]; ok {
		helper.SetPrimitiveValue(v, &accessConfig.EnableRbac, enableRbacKey)
	}

	return accessConfig
}

func constructAPIServerAccessConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterAPIServerAccessConfig {
	if len(data) < 1 {
		return nil
	}

	// APIServerConfig schema defines max 1
	apiServerAccessConfigData, _ := data[0].(map[string]any)
	apiServerAccessConfig := &models.VmwareTanzuManageV1alpha1AksclusterAPIServerAccessConfig{}

	if v, ok := apiServerAccessConfigData[authorizedIPRangesKey]; ok {
		apiServerAccessConfig.AuthorizedIPRanges = helper.SetPrimitiveList[string](v, authorizedIPRangesKey)
	}

	if v, ok := apiServerAccessConfigData[enablePrivateClusterKey]; ok {
		helper.SetPrimitiveValue(v, &apiServerAccessConfig.EnablePrivateCluster, enablePrivateClusterKey)
	}

	return apiServerAccessConfig
}

func constructLinuxConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterLinuxConfig {
	if len(data) < 1 {
		return nil
	}

	// LinuxConfig schema defines max 1
	linuxConfigData, _ := data[0].(map[string]any)
	linuxConfig := &models.VmwareTanzuManageV1alpha1AksclusterLinuxConfig{}

	if v, ok := linuxConfigData[adminUserNameKey]; ok {
		helper.SetPrimitiveValue(v, &linuxConfig.AdminUsername, adminUserNameKey)
	}

	if v, ok := linuxConfigData[sshkeysKey]; ok {
		linuxConfig.SSHKeys = helper.SetPrimitiveList[string](v, sshkeysKey)
	}

	return linuxConfig
}

func constructNetworkConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterNetworkConfig {
	if len(data) < 1 {
		return nil
	}

	// NetworkConfig schema defines max 1
	networkConfigData, _ := data[0].(map[string]any)
	networkConfig := &models.VmwareTanzuManageV1alpha1AksclusterNetworkConfig{}

	if v, ok := networkConfigData[networkPluginKey]; ok {
		helper.SetPrimitiveValue(v, &networkConfig.NetworkPlugin, networkPluginKey)
	}

	if v, ok := networkConfigData[networkPluginModeKey]; ok {
		helper.SetPrimitiveValue(v, &networkConfig.NetworkPluginMode, networkPluginModeKey)
	}

	if v, ok := networkConfigData[networkPolicyKey]; ok {
		helper.SetPrimitiveValue(v, &networkConfig.NetworkPolicy, networkPolicyKey)
	}

	if v, ok := networkConfigData[loadBalancerSkuKey]; ok {
		helper.SetPrimitiveValue(v, &networkConfig.LoadBalancerSku, loadBalancerSkuKey)
	}

	if v, ok := networkConfigData[dockerBridgeCidrKey]; ok {
		helper.SetPrimitiveValue(v, &networkConfig.DockerBridgeCidr, dockerBridgeCidrKey)
	}

	if v, ok := networkConfigData[dnsServiceIPKey]; ok {
		helper.SetPrimitiveValue(v, &networkConfig.DNSServiceIP, dnsServiceIPKey)
	}

	if v, ok := networkConfigData[dnsPrefixKey]; ok {
		helper.SetPrimitiveValue(v, &networkConfig.DNSPrefix, dnsPrefixKey)
	}

	if v, ok := networkConfigData[serviceCidrKey]; ok {
		networkConfig.ServiceCidrs = helper.SetPrimitiveList[string](v.([]any), serviceCidrKey)
	}

	if v, ok := networkConfigData[podCidrKey]; ok {
		networkConfig.PodCidrs = helper.SetPrimitiveList[string](v.([]any), podCidrKey)
	}

	return networkConfig
}

func constructStorageConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterStorageConfig {
	if len(data) < 1 {
		return nil
	}

	storageConfigData, _ := data[0].(map[string]any)
	storageConfig := &models.VmwareTanzuManageV1alpha1AksclusterStorageConfig{}

	if v, ok := storageConfigData[enableDiskCsiDriverKey]; ok {
		helper.SetPrimitiveValue(v, &storageConfig.EnableDiskCsiDriver, enableDiskCsiDriverKey)
	}

	if v, ok := storageConfigData[enableFileCsiDriverKey]; ok {
		helper.SetPrimitiveValue(v, &storageConfig.EnableFileCsiDriver, enableFileCsiDriverKey)
	}

	if v, ok := storageConfigData[enableSnapshotControllerKey]; ok {
		helper.SetPrimitiveValue(v, &storageConfig.EnableSnapshotController, enableSnapshotControllerKey)
	}

	return storageConfig
}

func constructAddonsConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterAddonsConfig {
	if len(data) < 1 {
		return nil
	}

	// AddonConfig schema defines max 1
	addonsConfigData, _ := data[0].(map[string]any)
	addonsConfig := &models.VmwareTanzuManageV1alpha1AksclusterAddonsConfig{}

	if v, ok := addonsConfigData[azureKeyvaultSecretsProviderAddonConfigKey]; ok {
		data, _ := v.([]any)
		addonsConfig.AzureKeyvaultSecretsProviderConfig = constructAzureKeyVaultSecretsProviderConfig(data)
	}

	if v, ok := addonsConfigData[azurePolicyAddonConfigKey]; ok {
		data, _ := v.([]any)
		addonsConfig.AzurePolicyConfig = constructAzurePolicyConfig(data)
	}

	if v, ok := addonsConfigData[monitorAddonConfigKey]; ok {
		data, _ := v.([]any)
		addonsConfig.MonitoringConfig = constructMonitoringConfig(data)
	}

	return addonsConfig
}

func constructAadConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterAADConfig {
	if len(data) < 1 {
		return nil
	}

	// AADConfig schema defines max 1
	aadConfigData, _ := data[0].(map[string]any)
	aadConfig := &models.VmwareTanzuManageV1alpha1AksclusterAADConfig{}

	if v, ok := aadConfigData[adminGroupIDsKey]; ok {
		aadConfig.AdminGroupObjectIds = helper.SetPrimitiveList[string](v, adminGroupIDsKey)
	}

	if v, ok := aadConfigData[enableAzureRbacKey]; ok {
		helper.SetPrimitiveValue(v, &aadConfig.EnableAzureRbac, enableAzureRbacKey)
	}

	if v, ok := aadConfigData[managedKey]; ok {
		helper.SetPrimitiveValue(v, &aadConfig.Managed, managedKey)
	}

	if v, ok := aadConfigData[tenantIDKey]; ok {
		helper.SetPrimitiveValue(v, &aadConfig.TenantID, tenantIDKey)
	}

	return aadConfig
}

func constructMonitoringConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterMonitoringAddonConfig {
	if len(data) < 1 {
		return nil
	}

	// MonitoringConfig schema defines max 1
	monitoringConfigData, _ := data[0].(map[string]any)
	monitoringConfig := &models.VmwareTanzuManageV1alpha1AksclusterMonitoringAddonConfig{}

	if v, ok := monitoringConfigData[enableKey]; ok {
		helper.SetPrimitiveValue(v, &monitoringConfig.Enabled, enableKey)
	}

	if v, ok := monitoringConfigData[logAnalyticsWorkspaceIDKey]; ok {
		helper.SetPrimitiveValue(v, &monitoringConfig.LogAnalyticsWorkspaceID, logAnalyticsWorkspaceIDKey)
	}

	return monitoringConfig
}

func constructAzurePolicyConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterAzurePolicyAddonConfig {
	if len(data) < 1 {
		return nil
	}

	// AzurePolicyConfig schema defines max 1
	azurePolicyConfigData, _ := data[0].(map[string]any)
	azurePolicyConfig := &models.VmwareTanzuManageV1alpha1AksclusterAzurePolicyAddonConfig{}

	if v, ok := azurePolicyConfigData[enableKey]; ok {
		helper.SetPrimitiveValue(v, &azurePolicyConfig.Enabled, enableKey)
	}

	return azurePolicyConfig
}

func constructAzureKeyVaultSecretsProviderConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterAzureKeyvaultSecretsProviderAddonConfig {
	if len(data) < 1 {
		return nil
	}

	// AzureKeyVaultSecretsProviderConfig defines max 1
	azureKeyVaultSecretsProviderConfigData, _ := data[0].(map[string]any)
	azureKeyVaultSecretsProviderConfig := &models.VmwareTanzuManageV1alpha1AksclusterAzureKeyvaultSecretsProviderAddonConfig{}

	if v, ok := azureKeyVaultSecretsProviderConfigData[enableSecretsRotationKey]; ok {
		helper.SetPrimitiveValue(v, &azureKeyVaultSecretsProviderConfig.EnableSecretRotation, enableSecretsRotationKey)
	}

	if v, ok := azureKeyVaultSecretsProviderConfigData[enableKey]; ok {
		helper.SetPrimitiveValue(v, &azureKeyVaultSecretsProviderConfig.Enabled, enableKey)
	}

	if v, ok := azureKeyVaultSecretsProviderConfigData[rotationPollIntervalKey]; ok {
		helper.SetPrimitiveValue(v, &azureKeyVaultSecretsProviderConfig.RotationPoolInterval, rotationPollIntervalKey)
	}

	return azureKeyVaultSecretsProviderConfig
}

func constructAutoUpgradeConfig(data []any) *models.VmwareTanzuManageV1alpha1AksclusterAutoUpgradeConfig {
	if len(data) < 1 {
		return nil
	}

	// AutoUpgradeConfig schema defines max 1
	autoUpgradeConfigData, _ := data[0].(map[string]any)
	autoUpgradeConfig := &models.VmwareTanzuManageV1alpha1AksclusterAutoUpgradeConfig{}

	if v, ok := autoUpgradeConfigData[upgradeChannelKey]; ok {
		channel := models.VmwareTanzuManageV1alpha1AksclusterChannel(v.(string))
		autoUpgradeConfig.Channel = &channel
	}

	return autoUpgradeConfig
}

func ToAKSClusterMap(cluster *models.VmwareTanzuManageV1alpha1AksCluster, nodepools []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) any {
	if cluster == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[CredentialNameKey] = cluster.FullName.CredentialName
	data[SubscriptionIDKey] = cluster.FullName.SubscriptionID
	data[ResourceGroupNameKey] = cluster.FullName.ResourceGroupName
	data[NameKey] = cluster.FullName.Name
	data[clusterSpecKey] = toClusterSpecMap(cluster.Spec, nodepools)

	return data
}

func toClusterSpecMap(spec *models.VmwareTanzuManageV1alpha1AksclusterSpec, nodepools []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) []any {
	if spec == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[clusterGroupKey] = spec.ClusterGroupName
	data[proxyNameKey] = spec.ProxyName
	data[configKey] = toConfigMap(spec.Config)
	data[nodepoolKey] = toNodePoolList(nodepools)
	data[agentNameKey] = spec.AgentName
	data[resourceIDKey] = spec.ResourceID

	return []any{data}
}

func toConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterClusterConfig) []any {
	data := make(map[string]any)
	if config == nil {
		return []any{data}
	}

	data[locationKey] = config.Location
	data[kubernetesVersionKey] = config.Version
	data[nodeResourceGroupNameKey] = config.NodeResourceGroupName
	data[diskEncryptionSetKey] = config.DiskEncryptionSetID
	data[tagsKey] = toInterfaceMap(config.Tags)
	data[skuKey] = toSKUMap(config.Sku)
	data[apiServerAccessConfigKey] = toServerAccessConfigMap(config.APIServerAccessConfig)
	data[accessConfigKey] = toAccessConfigMap(config.AccessConfig)
	data[linuxConfigKey] = toLinuxConfigMap(config.LinuxConfig)
	data[networkConfigKey] = toNetworkConfigMap(config.NetworkConfig)
	data[storageConfigKey] = toStorageConfigMap(config.StorageConfig)
	data[addonsConfigKey] = toAddonConfigMap(config.AddonsConfig)
	data[autoUpgradeConfigKey] = toAutoUpgradeConfigMap(config.AutoUpgradeConfig)

	return []any{data}
}

func toSKUMap(sku *models.VmwareTanzuManageV1alpha1AksclusterClusterSKU) []any {
	if sku == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[skuNameKey] = helper.PtrString(sku.Name)
	data[skuTierKey] = helper.PtrString(sku.Tier)

	return []any{data}
}

func toServerAccessConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterAPIServerAccessConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[authorizedIPRangesKey] = toInterfaceArray(config.AuthorizedIPRanges)
	data[enablePrivateClusterKey] = config.EnablePrivateCluster

	return []any{data}
}

func toAccessConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterAccessConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[enableRbacKey] = config.EnableRbac
	data[disableLocalAccountsKey] = config.DisableLocalAccounts
	data[aadConfigKey] = toAADConfigMap(config.AadConfig)

	return []any{data}
}

func toAADConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterAADConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[enableAzureRbacKey] = config.EnableAzureRbac
	data[managedKey] = config.Managed
	data[tenantIDKey] = config.TenantID
	data[adminGroupIDsKey] = toInterfaceArray(config.AdminGroupObjectIds)

	return []any{data}
}

func toLinuxConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterLinuxConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[adminUserNameKey] = config.AdminUsername
	data[sshkeysKey] = toInterfaceArray(config.SSHKeys)

	return []any{data}
}

func toNetworkConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterNetworkConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[loadBalancerSkuKey] = config.LoadBalancerSku
	data[networkPluginKey] = config.NetworkPlugin
	data[networkPluginModeKey] = config.NetworkPluginMode
	data[networkPolicyKey] = config.NetworkPolicy
	data[dnsPrefixKey] = config.DNSPrefix
	data[dnsServiceIPKey] = config.DNSServiceIP
	data[dockerBridgeCidrKey] = config.DockerBridgeCidr
	data[podCidrKey] = toInterfaceArray(config.PodCidrs)
	data[serviceCidrKey] = toInterfaceArray(config.ServiceCidrs)

	return []any{data}
}

func toStorageConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterStorageConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[enableDiskCsiDriverKey] = config.EnableDiskCsiDriver
	data[enableFileCsiDriverKey] = config.EnableFileCsiDriver
	data[enableSnapshotControllerKey] = config.EnableSnapshotController

	return []any{data}
}

func toAddonConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterAddonsConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[azureKeyvaultSecretsProviderAddonConfigKey] = toAzureKeyvaultSecretsProviderConfigMap(config.AzureKeyvaultSecretsProviderConfig)
	data[monitorAddonConfigKey] = toMonitorAddonConfigMap(config.MonitoringConfig)
	data[azurePolicyAddonConfigKey] = toAzurePolicyAddonConfigMap(config.AzurePolicyConfig)

	return []any{data}
}

func toAzureKeyvaultSecretsProviderConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterAzureKeyvaultSecretsProviderAddonConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[enableKey] = config.Enabled
	data[enableSecretsRotationKey] = config.EnableSecretRotation
	data[rotationPollIntervalKey] = config.RotationPoolInterval

	return []any{data}
}

func toMonitorAddonConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterMonitoringAddonConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[enableKey] = config.Enabled
	data[logAnalyticsWorkspaceIDKey] = config.LogAnalyticsWorkspaceID

	return []any{data}
}

func toAzurePolicyAddonConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterAzurePolicyAddonConfig) []any {
	if config == nil {
		return []any{}
	}

	data := make(map[string]any)
	data[enableKey] = config.Enabled

	return []any{data}
}

func toAutoUpgradeConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterAutoUpgradeConfig) []any {
	if config == nil {
		return nil
	}

	data := make(map[string]any)
	data[upgradeChannelKey] = helper.PtrString(config.Channel)

	return []any{data}
}

func toNodePoolList(nodepools []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) []any {
	n := make([]any, 0, len(nodepools))
	for _, v := range nodepools {
		n = append(n, ToNodepoolMap(v))
	}

	return n
}

func toInterfaceArray(vals []string) any {
	if vals == nil {
		return nil
	}

	a := make([]any, 0, len(vals))
	for _, v := range vals {
		a = append(a, v)
	}

	return a
}

func toInterfaceMap(tags map[string]string) any {
	m := make(map[string]any)
	for k, v := range tags {
		m[k] = v
	}

	return m
}

func constructStringMap[T any](data map[string]any) map[string]T {
	if len(data) < 1 {
		return nil
	}

	out := make(map[string]T)

	for k, v := range data {
		var value T

		helper.SetPrimitiveValue(v, &value, valueKey)

		out[k] = value
	}

	return out
}
