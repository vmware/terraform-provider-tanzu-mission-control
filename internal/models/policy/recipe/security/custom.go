/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
The contents of this file are not auto-generated using swagger CLI as the schema defined for the recipes are not a part of the TMC API models.
The models defined here are used to map the API request and response bodies to and from the terraform provider schema.
*/

package policyrecipesecuritymodel

import (
	"encoding/json"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom Input schema for security policy custom recipe version v1.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.Custom
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom struct {

	// Audit (dry-run).
	Audit bool `json:"audit,omitempty"`

	// Disable native pod security policy.
	DisableNativePsp bool `json:"disableNativePsp,omitempty"`

	// Allow privileged containers.
	AllowPrivilegedContainers bool `json:"allowPrivilegedContainers,omitempty"`

	// Allow privilege escalation.
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation,omitempty"`

	// Allow host namespace sharing.
	AllowHostNamespaceSharing bool `json:"allowHostNamespaceSharing,omitempty"`

	// Allow host network.
	AllowHostNetwork bool `json:"allowHostNetwork,omitempty"`

	// Read only root file system.
	ReadOnlyRootFileSystem bool `json:"readOnlyRootFileSystem,omitempty"`

	// Allowed host port range.
	AllowedHostPortRange *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange `json:"allowedHostPortRange,omitempty"`

	// Allowed volumes.
	AllowedVolumes []*VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume `json:"allowedVolumes,omitempty"`

	// Run as user.
	RunAsUser *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser `json:"runAsUser,omitempty"`

	// Run as group.
	RunAsGroup *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup `json:"runAsGroup,omitempty"`

	// supplemental groups.
	SupplementalGroups *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup `json:"supplementalGroups,omitempty"`

	// fsGroup.
	FsGroup *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup `json:"fsGroup,omitempty"`

	// Linux capabilities.
	LinuxCapabilities *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities `json:"linuxCapabilities,omitempty"`

	// Allowed host paths.
	AllowedHostPaths []*VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath `json:"allowedHostPaths,omitempty"`

	// Allowed selinux options.
	AllowedSELinuxOptions []*VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption `json:"allowedSELinuxOptions,omitempty"`

	// Sysctls.
	Sysctls *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls `json:"sysctls,omitempty"`

	// Seccomp.
	Seccomp *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp `json:"seccomp,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange Range.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.Range
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange struct {

	// Minimum allowed.
	Min int `json:"min"`

	// Maximum allowed.
	Max int `json:"max"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume Allowed volumes.
//
//  - *: A volume type.
//  - configMap: A volume type.
//  - downwardAPI: A volume type.
//  - emptyDir: A volume type.
//  - persistentVolumeClaim: A volume type.
//  - secret: A volume type.
//  - projected: A volume type.
//  - hostPath: A volume type.
//  - flexVolume: A volume type.
//  - awsElasticBlockStore: A volume type.
//  - azureDisk: A volume type.
//  - azureFile: A volume type.
//  - cephfs: A volume type.
//  - cinder: A volume type.
//  - csi: A volume type.
//  - fc: A volume type.
//  - flocker: A volume type.
//  - gcePersistentDisk: A volume type.
//  - gitRepo: A volume type.
//  - glusterfs: A volume type.
//  - iscsi: A volume type.
//  - local: A volume type.
//  - nfs: A volume type.
//  - portworxVolume: A volume type.
//  - quobyte: A volume type.
//  - rbd: A volume type.
//  - scaleIO: A volume type.
//  - storageos: A volume type.
//  - vsphereVolume: A volume type.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.AllowedVolume
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume string

func NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume(value VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume) *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeAll captures enum value "*".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeAll VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "*"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeConfigMap captures enum value "configMap".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeConfigMap VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "configMap"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeDownwardAPI captures enum value "DownwardAPI".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeDownwardAPI VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "downwardAPI"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeEmptyDir captures enum value "emptyDir".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeEmptyDir VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "emptyDir"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumePersistentVolumeClaim captures enum value "persistentVolumeClaim".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumePersistentVolumeClaim VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "persistentVolumeClaim"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeSecret captures enum value "secret".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeSecret VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "secret"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeProjected captures enum value "projected".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeProjected VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "projected"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeHostPath captures enum value "hostPath".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeHostPath VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "hostPath"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeFlexVolume captures enum value "flexVolume".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeFlexVolume VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "flexVolume"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeAwsElasticBlockStore captures enum value "awsElasticBlockStore".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeAwsElasticBlockStore VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "awsElasticBlockStore"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeAzureDisk captures enum value "azureDisk".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeAzureDisk VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "azureDisk"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeAzureFile captures enum value "azureFile".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeAzureFile VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "azureFile"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeCephfs captures enum value "cephfs".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeCephfs VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "cephfs"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeCinder captures enum value "cinder".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeCinder VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "cinder"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeCsi captures enum value "csi".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeCsi VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "csi"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeFc captures enum value "fc".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeFc VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "fc"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeFlocker captures enum value "flocker".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeFlocker VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "flocker"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeGcePersistentDisk captures enum value "gcePersistentDisk".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeGcePersistentDisk VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "gcePersistentDisk"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeGitRepo captures enum value "gitRepo".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeGitRepo VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "gitRepo"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeGlusterfs captures enum value "glusterfs".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeGlusterfs VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "glusterfs"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeIscsi captures enum value "iscsi".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeIscsi VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "iscsi"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeLocal captures enum value "local".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeLocal VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "local"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeNfs captures enum value "nfs".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeNfs VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "nfs"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumePortworxVolume captures enum value "portworxVolume".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumePortworxVolume VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "portworxVolume"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeQuobyte captures enum value "quobyte".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeQuobyte VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "quobyte"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeRbd captures enum value "rbd".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeRbd VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "rbd"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeScaleIO captures enum value "scaleIO".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeScaleIO VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "scaleIO"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeStorageos captures enum value "storageos".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeStorageos VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "storageos"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeVsphereVolume captures enum value "vsphereVolume".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeVsphereVolume VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume = "vsphereVolume"
)

// for schema.
var vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume
	if err := json.Unmarshal([]byte(`["*","configMap","downwardAPI","emptyDir","persistentVolumeClaim","secret","projected","hostPath","flexVolume","awsElasticBlockStore","azureDisk","azureFile","cephfs","cinder","csi","fc","flocker","gcePersistentDisk","gitRepo","glusterfs","iscsi","local","nfs","portworxVolume","quobyte","rbd","scaleIO","storageos","vsphereVolume"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeEnum = append(vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeEnum, v)
	}
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser Run as user.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.RunAsUser
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser struct {

	// Rule.
	Rule VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule `json:"rule"`

	// Allowed user id ranges.
	Ranges []*VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange `json:"ranges,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule Rule.
//
//  - RunAsAny: A rule type.
//  - MustRunAsNonRoot: A rule type.
//  - MustRunAs: A rule type.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.runAsUser.Rule
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule string

func NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule(value VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule) *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleRunAsAny captures enum value "RunAsAny".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleRunAsAny VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule = "RunAsAny"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleMustRunAsNonRoot captures enum value "MustRunAsNonRoot".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleMustRunAsNonRoot VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule = "MustRunAsNonRoot"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleMustRunAs captures enum value "MustRunAs".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleMustRunAs VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule = "MustRunAs"
)

// for schema.
var vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule
	if err := json.Unmarshal([]byte(`["RunAsAny","MustRunAsNonRoot","MustRunAs"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleEnum = append(vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleEnum, v)
	}
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup Run as group.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.RunAsGroup
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup struct {

	// Rule.
	Rule VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule `json:"rule"`

	// Allowed user id ranges.
	Ranges []*VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange `json:"ranges,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule Rule.
//
//  - RunAsAny: A rule type.
//  - MayRunAs: A rule type.
//  - MustRunAs: A rule type.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.runAsGroup.Rule
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule string

func NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule(value VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule) *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny captures enum value "RunAsAny".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule = "RunAsAny"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleMayRunAs captures enum value "MayRunAs".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleMayRunAs VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule = "MayRunAs"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleMustRunAs captures enum value "MustRunAs".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleMustRunAs VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule = "MustRunAs"
)

// for schema.
var vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule
	if err := json.Unmarshal([]byte(`["RunAsAny","MayRunAs","MustRunAs"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleEnum = append(vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleEnum, v)
	}
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities Linux capabilities.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.LinuxCapabilities
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities struct {

	// Allowed capabilities.
	AllowedCapabilities []*VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability `json:"allowedCapabilities,omitempty"`

	// Required drop capabilities.
	RequiredDropCapabilities []*VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability `json:"requiredDropCapabilities,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability Allowed capability.
//
//  - *: An allowed capability type.
//  - AUDIT_CONTROL: An allowed capability type.
//  - AUDIT_READ: An allowed capability type.
//  - AUDIT_WRITE: An allowed capability type.
//  - BLOCK_SUSPEND: An allowed capability type.
//  - CHOWN: An allowed capability type.
//  - DAC_OVERRIDE: An allowed capability type.
//  - DAC_READ_SEARCH: An allowed capability type.
//  - FOWNER: An allowed capability type.
//  - FSETID: An allowed capability type.
//  - IPC_LOCK: An allowed capability type.
//  - IPC_OWNER: An allowed capability type.
//  - KILL: An allowed capability type.
//  - LEASE: An allowed capability type.
//  - LINUX_IMMUTABLE: An allowed capability type.
//  - MAC_ADMIN: An allowed capability type.
//  - MAC_OVERRIDE: An allowed capability type.
//  - MKNOD: An allowed capability type.
//  - NET_ADMIN: An allowed capability type.
//  - NET_BIND_SERVICE: An allowed capability type.
//  - NET_BROADCAST: An allowed capability type.
//  - NET_RAW: An allowed capability type.
//  - SETGID: An allowed capability type.
//  - SETFCAP: An allowed capability type.
//  - SETPCAP: An allowed capability type.
//  - SETUID: An allowed capability type.
//  - SYS_ADMIN: An allowed capability type.
//  - SYS_BOOT: An allowed capability type.
//  - SYS_CHROOT: An allowed capability type.
//  - SYS_MODULE: An allowed capability type.
//  - SYS_NICE: An allowed capability type.
//  - SYS_PACCT: An allowed capability type.
//  - SYS_PTRACE: An allowed capability type.
//  - SYS_RAWIO: An allowed capability type.
//  - SYS_RESOURCE: An allowed capability type.
//  - SYS_TIME: An allowed capability type.
//  - SYS_TTY_CONFIG: An allowed capability type.
//  - SYSLOG: An allowed capability type.
//  - WAKE_ALARM: An allowed capability type.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.linuxCapabilities.AllowedCapability
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability string

func NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability(value VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability) *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability {
	v := value
	return &v
}

// nolint: dupl
const (

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityAll captures enum value "*".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityAll VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "*"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityAuditControl captures enum value "AUDIT_CONTROL".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityAuditControl VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "AUDIT_CONTROL"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityAuditRead captures enum value "AUDIT_READ".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityAuditRead VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "AUDIT_READ"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityAuditWrite captures enum value "AUDIT_WRITE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityAuditWrite VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "AUDIT_WRITE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityBlockSuspend captures enum value "BLOCK_SUSPEND".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityBlockSuspend VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "BLOCK_SUSPEND"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityChown captures enum value "CHOWN".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityChown VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "CHOWN"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityDacOverride captures enum value "DAC_OVERRIDE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityDacOverride VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "DAC_OVERRIDE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityDacReadSearch captures enum value "DAC_READ_SEARCH".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityDacReadSearch VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "DAC_READ_SEARCH"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityFOwner captures enum value "FOWNER".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityFOwner VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "FOWNER"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityFSetID captures enum value "FSETID".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityFSetID VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "FSETID"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityIpcLock captures enum value "IPC_LOCK".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityIpcLock VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "IPC_LOCK"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityIpcOwner captures enum value "IPC_OWNER".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityIpcOwner VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "IPC_OWNER"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityKill captures enum value "KILL".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityKill VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "KILL"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityLease captures enum value "LEASE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityLease VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "LEASE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityLinuxImmutable captures enum value "LINUX_IMMUTABLE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityLinuxImmutable VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "LINUX_IMMUTABLE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityMacAdmin captures enum value "MAC_ADMIN".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityMacAdmin VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "MAC_ADMIN"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityMacOverride captures enum value "MAC_OVERRIDE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityMacOverride VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "MAC_OVERRIDE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityMKNOD captures enum value "MKNOD".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityMKNOD VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "MKNOD"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityNetAdmin captures enum value "NET_ADMIN".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityNetAdmin VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "NET_ADMIN"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityNetBindService captures enum value "NET_BIND_SERVICE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityNetBindService VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "NET_BIND_SERVICE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityNetBroadcast captures enum value "NET_BROADCAST".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityNetBroadcast VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "NET_BROADCAST"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityNetRaw captures enum value "NET_RAW".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityNetRaw VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "NET_RAW"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySetGID captures enum value "SETGID".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySetGID VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SETGID"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySetFCap captures enum value "SETFCAP".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySetFCap VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SETFCAP"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySetPCap captures enum value "SETPCAP".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySetPCap VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SETPCAP"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySetUID captures enum value "SETUID".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySetUID VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SETUID"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysAdmin captures enum value "SYS_ADMIN".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysAdmin VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_ADMIN"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysBoot captures enum value "SYS_BOOT".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysBoot VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_BOOT"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysChroot captures enum value "SYS_CHROOT".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysChroot VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_CHROOT"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysModule captures enum value "SYS_MODULE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysModule VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_MODULE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysNice captures enum value "SYS_NICE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysNice VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_NICE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysPACCT captures enum value "SYS_PACCT".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysPACCT VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_PACCT"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysPTrace captures enum value "SYS_PTRACE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysPTrace VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_PTRACE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysRawIO captures enum value "SYS_RAWIO".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysRawIO VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_RAWIO"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysResource captures enum value "SYS_RESOURCE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysResource VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_RESOURCE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysTime captures enum value "SYS_TIME".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysTime VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_TIME"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysTTYConfig captures enum value "SYS_TTY_CONFIG".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysTTYConfig VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYS_TTY_CONFIG"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysLog captures enum value "SYSLOG".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilitySysLog VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "SYSLOG"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityWakeAlarm captures enum value "WAKE_ALARM".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityWakeAlarm VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability = "WAKE_ALARM"
)

// for schema.
var vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability
	if err := json.Unmarshal([]byte(`["*","AUDIT_CONTROL","AUDIT_READ","AUDIT_WRITE","BLOCK_SUSPEND","CHOWN","DAC_OVERRIDE","DAC_READ_SEARCH","FOWNER","FSETID","IPC_LOCK","IPC_OWNER","KILL","LEASE","LINUX_IMMUTABLE","MAC_ADMIN","MAC_OVERRIDE","MKNOD","NET_ADMIN","NET_BIND_SERVICE","NET_BROADCAST","NET_RAW","SETGID","SETFCAP","SETPCAP","SETUID","SYS_ADMIN","SYS_BOOT","SYS_CHROOT","SYS_MODULE","SYS_NICE","SYS_PACCT","SYS_PTRACE","SYS_RAWIO","SYS_RESOURCE","SYS_TIME","SYS_TTY_CONFIG","SYSLOG","WAKE_ALARM"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityEnum = append(vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityEnum, v)
	}
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability Required drop capability.
//
//  - ALL: A required drop capability type.
//  - AUDIT_CONTROL: A required drop capability type.
//  - AUDIT_READ: A required drop capability type.
//  - AUDIT_WRITE: A required drop capability type.
//  - BLOCK_SUSPEND: A required drop capability type.
//  - CHOWN: A required drop capability type.
//  - DAC_OVERRIDE: A required drop capability type.
//  - DAC_READ_SEARCH: A required drop capability type.
//  - FOWNER: A required drop capability type.
//  - FSETID: A required drop capability type.
//  - IPC_LOCK: A required drop capability type.
//  - IPC_OWNER: A required drop capability type.
//  - KILL: A required drop capability type.
//  - LEASE: A required drop capability type.
//  - LINUX_IMMUTABLE: A required drop capability type.
//  - MAC_ADMIN: A required drop capability type.
//  - MAC_OVERRIDE: A required drop capability type.
//  - MKNOD: A required drop capability type.
//  - NET_ADMIN: A required drop capability type.
//  - NET_BIND_SERVICE: A required drop capability type.
//  - NET_BROADCAST: A required drop capability type.
//  - NET_RAW: A required drop capability type.
//  - SETGID: A required drop capability type.
//  - SETFCAP: A required drop capability type.
//  - SETPCAP: A required drop capability type.
//  - SETUID: A required drop capability type.
//  - SYS_ADMIN: A required drop capability type.
//  - SYS_BOOT: A required drop capability type.
//  - SYS_CHROOT: A required drop capability type.
//  - SYS_MODULE: A required drop capability type.
//  - SYS_NICE: A required drop capability type.
//  - SYS_PACCT: A required drop capability type.
//  - SYS_PTRACE: A required drop capability type.
//  - SYS_RAWIO: A required drop capability type.
//  - SYS_RESOURCE: A required drop capability type.
//  - SYS_TIME: A required drop capability type.
//  - SYS_TTY_CONFIG: A required drop capability type.
//  - SYSLOG: A required drop capability type.
//  - WAKE_ALARM: A required drop capability type.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.linuxCapabilities.RequiredDropCapability
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability string

func NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability(value VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability) *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability {
	v := value
	return &v
}

// nolint: dupl
const (

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityAll captures enum value "ALL".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityAll VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "ALL"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityAuditControl captures enum value "AUDIT_CONTROL".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityAuditControl VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "AUDIT_CONTROL"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityAuditRead captures enum value "AUDIT_READ".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityAuditRead VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "AUDIT_READ"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityAuditWrite captures enum value "AUDIT_WRITE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityAuditWrite VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "AUDIT_WRITE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityBlockSuspend captures enum value "BLOCK_SUSPEND".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityBlockSuspend VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "BLOCK_SUSPEND"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityChown captures enum value "CHOWN".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityChown VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "CHOWN"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityDacOverride captures enum value "DAC_OVERRIDE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityDacOverride VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "DAC_OVERRIDE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityDacReadSearch captures enum value "DAC_READ_SEARCH".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityDacReadSearch VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "DAC_READ_SEARCH"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityFOwner captures enum value "FOWNER".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityFOwner VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "FOWNER"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityFSetID captures enum value "FSETID".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityFSetID VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "FSETID"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityIpcLock captures enum value "IPC_LOCK".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityIpcLock VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "IPC_LOCK"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityIpcOwner captures enum value "IPC_OWNER".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityIpcOwner VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "IPC_OWNER"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityKill captures enum value "KILL".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityKill VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "KILL"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityLease captures enum value "LEASE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityLease VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "LEASE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityLinuxImmutable captures enum value "LINUX_IMMUTABLE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityLinuxImmutable VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "LINUX_IMMUTABLE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityMacAdmin captures enum value "MAC_ADMIN".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityMacAdmin VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "MAC_ADMIN"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityMacOverride captures enum value "MAC_OVERRIDE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityMacOverride VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "MAC_OVERRIDE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityMKNOD captures enum value "MKNOD".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityMKNOD VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "MKNOD"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityNetAdmin captures enum value "NET_ADMIN".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityNetAdmin VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "NET_ADMIN"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityNetBindService captures enum value "NET_BIND_SERVICE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityNetBindService VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "NET_BIND_SERVICE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityNetBroadcast captures enum value "NET_BROADCAST".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityNetBroadcast VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "NET_BROADCAST"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityNetRaw captures enum value "NET_RAW".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityNetRaw VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "NET_RAW"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySetGID captures enum value "SETGID".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySetGID VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SETGID"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySetFCap captures enum value "SETFCAP".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySetFCap VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SETFCAP"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySetPCap captures enum value "SETPCAP".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySetPCap VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SETPCAP"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySetUID captures enum value "SETUID".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySetUID VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SETUID"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysAdmin captures enum value "SYS_ADMIN".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysAdmin VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_ADMIN"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysBoot captures enum value "SYS_BOOT".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysBoot VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_BOOT"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysChroot captures enum value "SYS_CHROOT".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysChroot VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_CHROOT"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysModule captures enum value "SYS_MODULE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysModule VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_MODULE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysNice captures enum value "SYS_NICE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysNice VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_NICE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysPACCT captures enum value "SYS_PACCT".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysPACCT VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_PACCT"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysPTrace captures enum value "SYS_PTRACE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysPTrace VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_PTRACE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysRawIO captures enum value "SYS_RAWIO".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysRawIO VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_RAWIO"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysResource captures enum value "SYS_RESOURCE".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysResource VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_RESOURCE"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysTime captures enum value "SYS_TIME".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysTime VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_TIME"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysTTYConfig captures enum value "SYS_TTY_CONFIG".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysTTYConfig VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYS_TTY_CONFIG"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysLog captures enum value "SYSLOG".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysLog VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "SYSLOG"

	// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityWakeAlarm captures enum value "WAKE_ALARM".
	VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityWakeAlarm VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability = "WAKE_ALARM"
)

// for schema.
var vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability
	if err := json.Unmarshal([]byte(`["ALL","AUDIT_CONTROL","AUDIT_READ","AUDIT_WRITE","BLOCK_SUSPEND","CHOWN","DAC_OVERRIDE","DAC_READ_SEARCH","FOWNER","FSETID","IPC_LOCK","IPC_OWNER","KILL","LEASE","LINUX_IMMUTABLE","MAC_ADMIN","MAC_OVERRIDE","MKNOD","NET_ADMIN","NET_BIND_SERVICE","NET_BROADCAST","NET_RAW","SETGID","SETFCAP","SETPCAP","SETUID","SYS_ADMIN","SYS_BOOT","SYS_CHROOT","SYS_MODULE","SYS_NICE","SYS_PACCT","SYS_PTRACE","SYS_RAWIO","SYS_RESOURCE","SYS_TIME","SYS_TTY_CONFIG","SYSLOG","WAKE_ALARM"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityEnum = append(vmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilityEnum, v)
	}
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath Allowed host path.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.AllowedHostPath
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath struct {

	// Read only flag.
	ReadOnly bool `json:"read_only,omitempty"`

	// Path prefix.
	PathPrefix string `json:"pathPrefix,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption Allowed selinux option.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.AllowedSELinuxOption
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption struct {

	// SELinux level.
	Level string `json:"level,omitempty"`

	// SELinux role.
	Role string `json:"role,omitempty"`

	// SELinux type.
	Type string `json:"type,omitempty"`

	// SELinux user.
	User string `json:"user,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls Sysctls.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.Sysctls
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls struct {

	// Forbidden sysctls.
	ForbiddenSysctls []*string `json:"forbiddenSysctls,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp Seccomp.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.custom.Seccomp
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp struct {

	// Allowed profiles.
	AllowedProfiles []*string `json:"allowedProfiles,omitempty"`

	// Allowed local host files.
	AllowedLocalhostFiles []*string `json:"allowedLocalhostFiles,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
