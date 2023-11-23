/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

type SupportedScopes string

const (
	ClusterScopeType      SupportedScopes = "cluster"
	ClusterGroupScopeType SupportedScopes = "cluster_group"
)
