/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	clustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster"
	clustergroupmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	namespacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/namespace"
	organizationmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/organization"
	workspacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

type scopeType int

const (
	unknown scopeType = iota
	organization
	clusterGroup
	cluster
	workspace
	namespace
)

var (
	stringToScope = map[string]scopeType{
		orgKey:          organization,
		clusterGroupKey: clusterGroup,
		clusterKey:      cluster,
		workspaceKey:    workspace,
		namespaceKey:    namespace,
	}
)

func (s scopeType) String() string {
	return [...]string{
		"unknown",
		orgKey,
		clusterGroupKey,
		clusterKey,
		workspaceKey,
		namespaceKey,
	}[s]
}

func constructScopeFullName(s scopeType, d *schema.ResourceData) interface{} {
	scopeData, ok := d.GetOk(scopeKey)
	if !ok {
		return nil
	}

	scopeDataList, _ := scopeData.([]interface{})
	if len(scopeDataList) < 1 {
		return nil
	}

	scopeDataMap, _ := scopeDataList[0].(map[string]interface{})

	val, ok := scopeDataMap[s.String()]
	if !ok {
		return nil
	}

	valList := val.([]interface{})
	if len(valList) < 1 {
		return nil
	}

	valMap, _ := valList[0].(map[string]interface{})

	switch s {
	case organization:
		return &organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName{
			OrgID: valMap[orgIDKey].(string),
		}
	case clusterGroup:
		return &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
			Name: valMap[clusterGroupNameKey].(string),
		}
	case cluster:
		return &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
			ManagementClusterName: valMap[managementClusterNameKey].(string),
			ProvisionerName:       valMap[provisionerNameKey].(string),
			Name:                  valMap[clusterNameKey].(string),
		}
	case workspace:
		return &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
			Name: valMap[workspaceNameKey].(string),
		}
	case namespace:
		return &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName{
			ManagementClusterName: valMap[managementClusterNameKey].(string),
			ProvisionerName:       valMap[provisionerNameKey].(string),
			ClusterName:           valMap[clusterNameKey].(string),
			Name:                  valMap[namespaceNameKey].(string),
		}
	default:
		return nil
	}
}

func getScopeType(d *schema.ResourceData) (scopeType, error) {
	scopes := []scopeType{}

	value, ok := d.GetOk(scopeKey)
	if !ok {
		return unknown, fmt.Errorf("no scope defined")
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return unknown, fmt.Errorf("no type defined in scope")
	}

	scopeData := data[0].(map[string]interface{})

	for key, val := range scopeData {
		v := val.([]interface{})
		if len(v) == 0 {
			continue
		}

		scopes = append(scopes, stringToScope[key])
	}

	if len(scopes) != 1 {
		return unknown, fmt.Errorf("none or more than one scope types are defined")
	}

	return scopes[0], nil
}
