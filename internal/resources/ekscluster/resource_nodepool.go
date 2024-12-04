// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package ekscluster

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

// nodepoolDefinitionSchema defines the info and nodepool spec for EKS clusters.
//
// Note: ForceNew is not used in any of the elements because this is a part of
// EKS cluster and we don't want to replace full clusters because of Nodepool
// change.
var nodepoolDefinitionSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		infoKey: {
			Type:        schema.TypeList,
			Description: "Info for the nodepool",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					NameKey: {
						Type:        schema.TypeString,
						Description: "Name of the nodepool, immutable",
						Required:    true,
					},
					common.DescriptionKey: {
						Type:        schema.TypeString,
						Description: "Description for the nodepool",
						Optional:    true,
					},
				},
			},
		},
		specKey: nodepoolSpecSchema,
	},
}

var nodepoolSpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			roleArnKey: {
				Type:        schema.TypeString,
				Description: "ARN of the IAM role that provides permissions for the Kubernetes nodepool to make calls to AWS API operations, immutable",
				Required:    true,
			},
			amiTypeKey: {
				Type:        schema.TypeString,
				Description: "AMI type, immutable",
				Optional:    true,
				Computed:    true,
			},
			amiInfoKey: {
				Type:        schema.TypeList,
				Description: "AMI info for the nodepool if AMI type is specified as CUSTOM",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						amiIDKey: {
							Type:        schema.TypeString,
							Description: "ID of the AMI to be used",
							Optional:    true,
						},
						overrideBootstrapCmdKey: {
							Type:        schema.TypeString,
							Description: "Override bootstrap command for the custom AMI",
							Optional:    true,
						},
					},
				},
			},
			capacityTypeKey: {
				Type:        schema.TypeString,
				Description: "Capacity Type",
				Optional:    true,
				Computed:    true,
			},
			rootDiskSizeKey: {
				Type:        schema.TypeInt,
				Description: "Root disk size in GiB, immutable",
				Optional:    true,
				Computed:    true,
			},
			tagsKey: {
				Type:        schema.TypeMap,
				Description: "EKS specific tags",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			nodeLabelsKey: {
				Type:        schema.TypeMap,
				Description: "Kubernetes node labels",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			subnetIdsKey: {
				Type:        schema.TypeSet,
				Description: "Subnets required for the nodepool",
				Required:    true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					MinItems: 2,
				},
			},
			remoteAccessKey: {
				Type:        schema.TypeList,
				Description: "Remote access to worker nodes, immutable",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						sshKeyKey: {
							Type:        schema.TypeString,
							Description: "SSH key allows you to connect to your instances and gather diagnostic information if there are issues.",
							Optional:    true,
						},
						securityGroupsKey: {
							Type:        schema.TypeSet,
							Description: "Security groups for the VMs",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			launchTemplateKey: {
				Type:        schema.TypeList,
				Description: "Launch template for the nodepool",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						idKey: {
							Type:        schema.TypeString,
							Description: "The ID of the launch template",
							Optional:    true,
						},
						nameKey: {
							Type:        schema.TypeString,
							Description: "The name of the launch template",
							Optional:    true,
						},
						versionKey: {
							Type:        schema.TypeString,
							Description: "The version of the launch template to use",
							Optional:    true,
						},
					},
				},
			},
			scalingConfigKey: {
				Type:        schema.TypeList,
				Description: "Nodepool scaling config",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						desiredSizeKey: {
							Type:        schema.TypeInt,
							Description: "Desired size of nodepool",
							Optional:    true,
						},
						maxSizeKey: {
							Type:        schema.TypeInt,
							Description: "Maximum size of nodepool",
							Optional:    true,
						},
						minSizeKey: {
							Type:        schema.TypeInt,
							Description: "Minimum size of nodepool",
							Optional:    true,
						},
					},
				},
			},
			updateConfigKey: {
				Type:        schema.TypeList,
				Description: "Update config for the nodepool",
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						maxUnavailableNodesKey: {
							Type:        schema.TypeString,
							Description: "Maximum number of nodes unavailable at once during a version update",
							Optional:    true,
						},
						maxUnavailablePercentageKey: {
							Type:        schema.TypeString,
							Description: "Maximum percentage of nodes unavailable during a version update",
							Optional:    true,
						},
					},
				},
			},
			taintsKey: {
				Type:        schema.TypeList,
				Description: "If specified, the node's taints",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						effectKey: {
							Type:        schema.TypeString,
							Description: "Current effect state of the node pool",
							Optional:    true,
						},
						keyKey: {
							Type:        schema.TypeString,
							Description: "The taint key to be applied to a node",
							Optional:    true,
						},
						valueKey: {
							Type:        schema.TypeString,
							Description: "The taint value corresponding to the taint key",
							Optional:    true,
						},
					},
				},
			},
			instanceTypesKey: {
				Type:        schema.TypeSet,
				Description: "Nodepool instance types, immutable",
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			releaseVersionKey: {
				Type:        schema.TypeString,
				Description: "AMI release version",
				Optional:    true,
				Computed:    true,
			},
		},
	},
}

func flattenNodePools(arr []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) []interface{} {
	data := make([]interface{}, 0, len(arr))

	if len(arr) == 0 {
		return data
	}

	for _, item := range arr {
		data = append(data, flattenNodePool(item))
	}

	return data
}

func flattenNodePool(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) map[string]interface{} {
	data := make(map[string]interface{})

	if item == nil {
		return data
	}

	if item.Info != nil {
		data[infoKey] = flattenInfo(item.Info)
	}

	if item.Spec != nil {
		data[specKey] = flattenSpec(item.Spec)
	}

	return data
}

func flattenInfo(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[common.DescriptionKey] = item.Description
	data[nameKey] = item.Name

	return []interface{}{data}
}

func flattenSpec(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[amiTypeKey] = item.AmiType

	if item.AmiInfo != nil &&
		(item.AmiInfo.AmiID != "" || item.AmiInfo.OverrideBootstrapCmd != "") {
		data[amiInfoKey] = flattenAmiInfo(item.AmiInfo)
	}

	data[capacityTypeKey] = item.CapacityType

	if len(item.InstanceTypes) > 0 {
		data[instanceTypesKey] = item.InstanceTypes
	}

	if item.LaunchTemplate != nil &&
		(item.LaunchTemplate.ID != "" || item.LaunchTemplate.Name != "" || item.LaunchTemplate.Version != "") {
		data[launchTemplateKey] = flattenLaunchTemplate(item.LaunchTemplate)
	}

	data[nodeLabelsKey] = item.NodeLabels

	if item.RemoteAccess != nil && (item.RemoteAccess.SSHKey != "" || len(item.RemoteAccess.SecurityGroups) > 0) {
		data[remoteAccessKey] = flattenRemoteAccess(item.RemoteAccess)
	}

	data[roleArnKey] = item.RoleArn

	if item.RootDiskSize != 0 {
		data[rootDiskSizeKey] = item.RootDiskSize
	}

	if item.ScalingConfig != nil {
		data[scalingConfigKey] = flattenScalingConfig(item.ScalingConfig)
	}

	if len(item.SubnetIds) > 0 {
		data[subnetIdsKey] = item.SubnetIds
	}

	data[tagsKey] = item.Tags

	if len(item.Taints) > 0 {
		data[taintsKey] = flattenTaints(item.Taints)
	}

	if item.UpdateConfig != nil {
		data[updateConfigKey] = flattenUpdateConfig(item.UpdateConfig)
	}

	if item.ReleaseVersion != "" {
		data[releaseVersionKey] = item.ReleaseVersion
	}

	return []interface{}{data}
}

func flattenAmiInfo(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[amiIDKey] = item.AmiID
	data[overrideBootstrapCmdKey] = item.OverrideBootstrapCmd

	return []interface{}{data}
}

func flattenLaunchTemplate(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[idKey] = item.ID
	data[nameKey] = item.Name
	data[versionKey] = item.Version

	return []interface{}{data}
}

func flattenRemoteAccess(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if len(item.SecurityGroups) > 0 {
		data[securityGroupsKey] = item.SecurityGroups
	}

	data[sshKeyKey] = item.SSHKey

	return []interface{}{data}
}

func flattenScalingConfig(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[desiredSizeKey] = item.DesiredSize
	data[maxSizeKey] = item.MaxSize
	data[minSizeKey] = item.MinSize

	return []interface{}{data}
}

func flattenTaints(arr []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint) []interface{} {
	data := make([]interface{}, 0, len(arr))

	if len(arr) == 0 {
		return data
	}

	for _, item := range arr {
		data = append(data, flattenTaint(item))
	}

	return data
}

func flattenTaint(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint) map[string]interface{} {
	data := make(map[string]interface{})

	if item == nil {
		return data
	}

	data[effectKey] = item.Effect
	data[keyKey] = item.Key
	data[valueKey] = item.Value

	return data
}

func flattenUpdateConfig(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if item.MaxUnavailableNodes != "" {
		data[maxUnavailableNodesKey] = item.MaxUnavailableNodes
	}

	if item.MaxUnavailablePercentage != "" {
		data[maxUnavailablePercentageKey] = item.MaxUnavailablePercentage
	}

	return []interface{}{data}
}

func constructNodepools(nodepoolsDefData []interface{}) []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
	nodepools := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{}

	for _, npDefData := range nodepoolsDefData {
		data, _ := npDefData.(map[string]interface{})
		np := constructNodepoolDef(data)
		nodepools = append(nodepools, np)
	}

	return nodepools
}

func constructNodepoolDef(nodepoolDefData map[string]interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
	definition := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{}

	if v, ok := nodepoolDefData[infoKey]; ok {
		data, _ := v.([]interface{})
		definition.Info = constructNodepoolInfo(data)
	}

	if v, ok := nodepoolDefData[specKey]; ok {
		data, _ := v.([]interface{})
		definition.Spec = constructNodepoolSpec(data)
	}

	return definition
}

func constructNodepoolInfo(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo {
	info := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{}

	if len(data) == 0 || data[0] == nil {
		return info
	}

	nodepoolInfoData, _ := data[0].(map[string]interface{})

	if v, ok := nodepoolInfoData[nameKey]; ok {
		helper.SetPrimitiveValue(v, &info.Name, nameKey)
	}

	if v, ok := nodepoolInfoData[common.DescriptionKey]; ok {
		helper.SetPrimitiveValue(v, &info.Description, common.DescriptionKey)
	}

	return info
}

func constructNodepoolSpec(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
	spec := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec{}

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData, _ := data[0].(map[string]interface{})

	if v, ok := specData[roleArnKey]; ok {
		helper.SetPrimitiveValue(v, &spec.RoleArn, roleArnKey)
	}

	if v, ok := specData[amiTypeKey]; ok {
		helper.SetPrimitiveValue(v, &spec.AmiType, amiTypeKey)
	}

	if v, ok := specData[amiInfoKey]; ok {
		data, _ := v.([]interface{})
		spec.AmiInfo = constructAmiInfo(data)
	}

	if v, ok := specData[capacityTypeKey]; ok {
		helper.SetPrimitiveValue(v, &spec.CapacityType, capacityTypeKey)
	}

	if v, ok := specData[rootDiskSizeKey]; ok {
		helper.SetPrimitiveValue(v, &spec.RootDiskSize, nameKey)
	}

	if v, ok := specData[tagsKey]; ok {
		data, _ := v.(map[string]interface{})
		spec.Tags = constructStringMap(data)
	}

	if v, ok := specData[nodeLabelsKey]; ok {
		data, _ := v.(map[string]interface{})
		spec.NodeLabels = constructStringMap(data)
	}

	if v, ok := specData[subnetIdsKey]; ok {
		if data, ok := v.(*schema.Set); ok {
			spec.SubnetIds = constructStringList(data.List())
		}
	}

	if v, ok := specData[launchTemplateKey]; ok {
		data, _ := v.([]interface{})
		spec.LaunchTemplate = constructLaunchTemplate(data)
	}

	if v, ok := specData[remoteAccessKey]; ok {
		data, _ := v.([]interface{})
		spec.RemoteAccess = constructRemoteAccess(data)
	}

	if v, ok := specData[scalingConfigKey]; ok {
		data, _ := v.([]interface{})
		spec.ScalingConfig = constructScalingConfig(data)
	}

	if v, ok := specData[updateConfigKey]; ok {
		data, _ := v.([]interface{})
		spec.UpdateConfig = constructUpdateConfig(data)
	}

	if v, ok := specData[taintsKey]; ok {
		data, _ := v.([]interface{})
		spec.Taints = constructTaints(data)
	}

	if v, ok := specData[instanceTypesKey]; ok {
		if data, ok := v.(*schema.Set); ok {
			spec.InstanceTypes = constructStringList(data.List())
		}
	}

	if v, ok := specData[releaseVersionKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ReleaseVersion, releaseVersionKey)
	}

	return spec
}

func constructAmiInfo(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo {
	amiInfo := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo{}

	if len(data) == 0 || data[0] == nil {
		return amiInfo
	}

	aiData, _ := data[0].(map[string]interface{})

	if v, ok := aiData[amiIDKey]; ok {
		helper.SetPrimitiveValue(v, &amiInfo.AmiID, amiIDKey)
	}

	if v, ok := aiData[overrideBootstrapCmdKey]; ok {
		helper.SetPrimitiveValue(v, &amiInfo.OverrideBootstrapCmd, overrideBootstrapCmdKey)
	}

	return amiInfo
}

func constructLaunchTemplate(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate {
	lt := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{}

	if len(data) == 0 || data[0] == nil {
		return lt
	}

	ltData, _ := data[0].(map[string]interface{})

	if v, ok := ltData[idKey]; ok {
		helper.SetPrimitiveValue(v, &lt.ID, idKey)
	}

	if v, ok := ltData[nameKey]; ok {
		helper.SetPrimitiveValue(v, &lt.Name, nameKey)
	}

	if v, ok := ltData[versionKey]; ok {
		helper.SetPrimitiveValue(v, &lt.Version, versionKey)
	}

	return lt
}

func constructRemoteAccess(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess {
	ra := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess{}

	if len(data) == 0 || data[0] == nil {
		return ra
	}

	raData, _ := data[0].(map[string]interface{})

	if v, ok := raData[sshKeyKey]; ok {
		helper.SetPrimitiveValue(v, &ra.SSHKey, sshKeyKey)
	}

	if v, ok := raData[securityGroupsKey]; ok {
		if data, ok := v.(*schema.Set); ok {
			ra.SecurityGroups = constructStringList(data.List())
		}
	}

	return ra
}

func constructScalingConfig(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig {
	sc := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig{}

	if len(data) == 0 || data[0] == nil {
		return sc
	}

	scData, _ := data[0].(map[string]interface{})

	if v, ok := scData[desiredSizeKey]; ok {
		helper.SetPrimitiveValue(v, &sc.DesiredSize, desiredSizeKey)
	}

	if v, ok := scData[maxSizeKey]; ok {
		helper.SetPrimitiveValue(v, &sc.MaxSize, maxSizeKey)
	}

	if v, ok := scData[minSizeKey]; ok {
		helper.SetPrimitiveValue(v, &sc.MinSize, minSizeKey)
	}

	return sc
}

func constructUpdateConfig(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig {
	uc := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig{}

	if len(data) == 0 || data[0] == nil {
		return uc
	}

	ucData, _ := data[0].(map[string]interface{})

	if v, ok := ucData[maxUnavailableNodesKey]; ok {
		helper.SetPrimitiveValue(v, &uc.MaxUnavailableNodes, maxUnavailableNodesKey)
	}

	if v, ok := ucData[maxUnavailablePercentageKey]; ok {
		helper.SetPrimitiveValue(v, &uc.MaxUnavailablePercentage, maxUnavailablePercentageKey)
	}

	return uc
}

func constructTaints(taintsData []interface{}) []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint {
	taints := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{}

	for _, data := range taintsData {
		taint := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{}
		tdata, _ := data.(map[string]interface{})

		if v, ok := tdata[effectKey]; ok {
			data, _ := v.(string)
			switch eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(data) {
			case eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectEFFECTUNSPECIFIED:
				taint.Effect = eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectEFFECTUNSPECIFIED.Pointer()
			case eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOEXECUTE:
				taint.Effect = eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOEXECUTE.Pointer()
			case eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOSCHEDULE:
				taint.Effect = eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOSCHEDULE.Pointer()
			case eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE:
				taint.Effect = eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE.Pointer()
			default:
				panic("unknown taint effect")
			}
		}

		if v, ok := tdata[keyKey]; ok {
			helper.SetPrimitiveValue(v, &taint.Key, keyKey)
		}

		if v, ok := tdata[valueKey]; ok {
			helper.SetPrimitiveValue(v, &taint.Value, valueKey)
		}

		taints = append(taints, taint)
	}

	return taints
}

func handleNodepoolDiffs(config authctx.TanzuContext, opsRetryTimeout time.Duration, clusterFn *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName, nodepools []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) error {
	npresp, err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceList(clusterFn)
	if err != nil {
		return errors.Wrapf(err, "failed to list nodepools for cluster: %s", clusterFn)
	}

	npPosMap := nodepoolPosMap(nodepools)
	tmcNps := map[string]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool{}

	npUpdate := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{}
	npCreate := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{}
	npDelete := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName{}

	for _, tmcNp := range npresp.Nodepools {
		tmcNps[tmcNp.FullName.Name] = tmcNp

		if pos, ok := npPosMap[tmcNp.FullName.Name]; ok {
			// np exisits in both TMC and TF
			newNp := nodepools[pos]
			fillTMCSetValues(tmcNp.Spec, newNp.Spec)

			if checkNodepoolUpdate(tmcNp, newNp) {
				npUpdate = append(npUpdate, newNp)
			}
		} else {
			// np exisits in TMC but not in TF
			npDelete = append(npDelete, tmcNp.FullName)
		}
	}

	for _, tfNp := range nodepools {
		if _, ok := tmcNps[tfNp.Info.Name]; !ok {
			npCreate = append(npCreate, tfNp)
		}
	}

	err = handleNodepoolCreates(config, opsRetryTimeout, clusterFn, npCreate)
	if err != nil {
		return errors.Wrap(err, "failed to create nodepools that are not present in TMC")
	}

	err = handleNodepoolUpdates(config, opsRetryTimeout, tmcNps, npUpdate)
	if err != nil {
		return errors.Wrapf(err, "failed to update existing nodepools")
	}

	err = handleNodepoolDeletes(config, opsRetryTimeout, npDelete)
	if err != nil {
		return errors.Wrapf(err, "failed to delete nodepools")
	}

	return nil
}

func handleNodepoolDeletes(config authctx.TanzuContext, opsRetryTimeout time.Duration, npFns []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName) error {
	for _, npFn := range npFns {
		err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceDelete(npFn)
		if err != nil {
			return errors.Wrap(err, "delete api call failed")
		}

		getNodepoolResourceRetryableFn := func() (retry bool, err error) {
			_, err = config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceGet(npFn)
			if err == nil {
				// we don't want to fail deletion if the deletion is not
				// completed within the expected time
				return true, nil
			}

			if !clienterrors.IsNotFoundError(err) {
				return true, err
			}

			return false, nil
		}

		_, err = helper.RetryUntilTimeout(getNodepoolResourceRetryableFn, 10*time.Second, opsRetryTimeout)
		if err != nil {
			return errors.Wrapf(err, "failed to verify EKS nodepool resource(%s) clean up", npFn.Name)
		}
	}

	return nil
}

func handleNodepoolUpdates(config authctx.TanzuContext, opsRetryTimeout time.Duration, tmcNps map[string]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool, nps []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) error {
	for _, np := range nps {
		tmcNp := tmcNps[np.Info.Name]

		req := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest{
			Nodepool: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool{
				FullName: tmcNp.FullName,
				Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
					Annotations:      tmcNp.Meta.Annotations,
					Description:      np.Info.Description,
					Labels:           tmcNp.Meta.Labels,
					ParentReferences: tmcNp.Meta.ParentReferences,
					ResourceVersion:  tmcNp.Meta.ResourceVersion,
					UID:              tmcNp.Meta.UID,
				},
				Spec: np.Spec,
			},
		}

		_, err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceUpdate(req)
		if err != nil {
			return errors.Wrapf(err, "failed to update nodepool %s", np.Info.Name)
		}

		getNodepoolResourceRetryableFn := getWaitForNodepoolReadyFn(config, tmcNp.FullName)

		_, err = helper.RetryUntilTimeout(getNodepoolResourceRetryableFn, 10*time.Second, opsRetryTimeout)
		if err != nil {
			return errors.Wrapf(err, "failed to verify EKS nodepool resource(%s) creation", np.Info.Name)
		}
	}

	return nil
}

func fillTMCSetValues(tmcNpSpec *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec, npSpec *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
	if npSpec.AmiType == "" {
		npSpec.AmiType = tmcNpSpec.AmiType
	}

	if npSpec.CapacityType == "" {
		npSpec.CapacityType = tmcNpSpec.CapacityType
	}

	if npSpec.ReleaseVersion == "" {
		npSpec.ReleaseVersion = tmcNpSpec.ReleaseVersion
	}
}

func handleNodepoolCreates(config authctx.TanzuContext, opsRetryTimeout time.Duration, clusterFn *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName, nps []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) error {
	err := createNodepools(config, clusterFn, nps)
	if err != nil {
		return errors.Wrap(err, "error while creating nodepools")
	}

	for _, np := range nps {
		npFn := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName{
			OrgID:          clusterFn.OrgID,
			CredentialName: clusterFn.CredentialName,
			Region:         clusterFn.Region,
			EksClusterName: clusterFn.Name,
			Name:           np.Info.Name,
		}

		getNodepoolResourceRetryableFn := getWaitForNodepoolReadyFn(config, npFn)

		_, err := helper.RetryUntilTimeout(getNodepoolResourceRetryableFn, 10*time.Second, opsRetryTimeout)
		if err != nil {
			return errors.Wrapf(err, "failed to verify EKS nodepool resource(%s) creation", npFn.Name)
		}
	}

	return nil
}

func createNodepools(config authctx.TanzuContext, clusterFn *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName, nps []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) error {
	for _, np := range nps {
		// Nodepools are created with the default release version, this field is only used for nodepool update
		if np.Spec.ReleaseVersion != "" {
			return errors.New("AMI release version of nodepool is not allowed to be set during Create")
		}

		npFn := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName{
			OrgID:          clusterFn.OrgID,
			CredentialName: clusterFn.CredentialName,
			Region:         clusterFn.Region,
			EksClusterName: clusterFn.Name,
			Name:           np.Info.Name,
		}

		req := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest{
			Nodepool: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool{
				FullName: npFn,
				Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
					Description: np.Info.Description,
				},
				Spec: np.Spec,
			},
		}

		_, err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceCreate(req)
		if err != nil && !clienterrors.IsAlreadyExistsError(err) {
			return errors.Wrapf(err, "failed to create nodepool %s", np.Info.Name)
		}
	}

	return nil
}

func getWaitForNodepoolReadyFn(config authctx.TanzuContext, npFn *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName) func() (retry bool, err error) {
	return func() (retry bool, err error) {
		resp, err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceGet(npFn)
		if err != nil {
			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS nodepoool entry, name : %s", npFn.Name)
		}

		if resp.Nodepool.Status.Phase != nil &&
			*resp.Nodepool.Status.Phase != eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseREADY {
			if c, ok := resp.Nodepool.Status.Conditions[readyCondition]; ok &&
				c.Severity != nil &&
				*c.Severity == eksmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityERROR {
				return false, errors.Errorf("nodepool %s in error state due to %s, %s", npFn.Name, c.Reason, c.Message)
			}

			return true, nil
		}

		return false, nil
	}
}

func checkNodepoolUpdate(oldNp *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool, newNp *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) bool {
	return oldNp.Meta.Description != newNp.Info.Description ||
		!nodepoolSpecEqual(oldNp.Spec, newNp.Spec)
}
