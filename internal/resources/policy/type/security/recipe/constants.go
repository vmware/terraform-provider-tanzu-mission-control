/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

const (
	BaselineKey                  = "baseline"
	CustomKey                    = "custom"
	StrictKey                    = "strict"
	AuditKey                     = "audit"
	DisableNativePspKey          = "disable_native_psp"
	allowPrivilegedContainersKey = "allow_privileged_containers"
	allowPrivilegeEscalationKey  = "allow_privilege_escalation"
	allowHostNamespaceSharingKey = "allow_host_namespace_sharing"
	allowHostNetworkKey          = "allow_host_network"
	readOnlyRootFileSystemKey    = "read_only_root_file_system"
	allowedHostPortRangeKey      = "allowed_host_port_range"
	allowedVolumesKey            = "allowed_volumes"
	runAsUserKey                 = "run_as_user"
	runAsGroupKey                = "run_as_group"
	supplementalGroupsKey        = "supplemental_groups"
	fsGroupKey                   = "fs_group"
	linuxCapabilitiesKey         = "linux_capabilities"
	allowedHostPathsKey          = "allowed_host_paths"
	allowedSELinuxOptionsKey     = "allowed_se_linux_options"
	sysctlsKey                   = "sysctls"
	seccompKey                   = "seccomp"
	minKey                       = "min"
	maxKey                       = "max"
	ruleKey                      = "rule"
	rangesKey                    = "ranges"
	allowedCapabilitiesKey       = "allowed_capabilities"
	requiredDropCapabilitiesKey  = "required_drop_capabilities"
	readOnlyKey                  = "read_only"
	pathPrefixKey                = "path_prefix"
	levelKey                     = "level"
	roleKey                      = "role"
	typeKey                      = "type"
	userKey                      = "user"
	forbiddenSysctlsKey          = "forbidden_sysctls"
	allowedProfilesKey           = "allowed_profiles"
	allowedLocalhostFilesKey     = "allowed_localhost_files"
)
