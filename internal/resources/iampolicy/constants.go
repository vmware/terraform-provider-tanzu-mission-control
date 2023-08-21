/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

const (
	ResourceName      = "tanzu-mission-control_iam_policy"
	scopeKey          = "scope"
	clusterKey        = "cluster"
	clusterGroupKey   = "cluster_group"
	namespaceKey      = "namespace"
	workspaceKey      = "workspace"
	organizationKey   = "organization"
	organizationIDKey = "org_id"
	roleBindingsKey   = "role_bindings"
	roleKey           = "role"
	subjectsKey       = "subjects"
	subjectNameKey    = "name"
	subjectKindKey    = "kind"
	createKey         = "create"
	updateKey         = "update"
)

// Allowed scopes.
const (
	unknownScope scope = iota
	organizationScope
	clusterGroupScope
	clusterScope
	workspaceScope
	namespaceScope
)

const roleSubjectDelimiter = ","
