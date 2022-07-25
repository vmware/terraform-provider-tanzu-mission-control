/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	clustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster"
	clustergroupmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	namespacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/namespace"
	organizationmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/organization"
	workspacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

type resourceDataOpt func(map[string][]interface{})

func TestScope(t *testing.T) {
	testCases := []struct {
		name        string
		input       *schema.ResourceData
		expScope    scopeType
		expFullName interface{}
		expErr      error
	}{
		{
			name: "test for organization scope",
			input: getSchemaData(
				t,
				withOrg("dummy"),
			),
			expScope: organization,
			expFullName: &organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName{
				OrgID: "dummy",
			},
		},
		{
			name: "test for cluster group scope",
			input: getSchemaData(
				t,
				withClusterGroup("default"),
			),
			expScope: clusterGroup,
			expFullName: &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
				Name: "default",
			},
		},
		{
			name: "test for cluster scope",
			input: getSchemaData(
				t,
				withCluster("attached", "attached", "dummy"),
			),
			expScope: cluster,
			expFullName: &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
				ManagementClusterName: "attached",
				ProvisionerName:       "attached",
				Name:                  "dummy",
			},
		},
		{
			name: "test for workspace scope",
			input: getSchemaData(
				t,
				withWorkspace("default"),
			),
			expScope: workspace,
			expFullName: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
				Name: "default",
			},
		},
		{
			name: "test for namespace scope",
			input: getSchemaData(
				t,
				withNamespace("attached", "attached", "dummy", "default"),
			),
			expScope: namespace,
			expFullName: &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName{
				ManagementClusterName: "attached",
				ProvisionerName:       "attached",
				ClusterName:           "dummy",
				Name:                  "default",
			},
		},
		{
			name: "error: more than 1 scopes defined",
			input: getSchemaData(
				t,
				withOrg("dummy"),
				withClusterGroup("default"),
			),
			expErr: fmt.Errorf("none or more than one scope types are defined"),
		},
		{
			name:   "error: no scope defined",
			input:  getSchemaData(t),
			expErr: fmt.Errorf("no type defined in scope"),
		},
	}

	for _, testCase := range testCases {
		test := testCase
		t.Run(test.name, func(t *testing.T) {
			actualScope, err := getScopeType(test.input)
			if test.expErr != nil {
				require.EqualError(t, err, test.expErr.Error())
				return
			}

			require.NoError(t, err)
			require.EqualValues(t, test.expScope, actualScope)

			fn := constructScopeFullName(actualScope, test.input)
			require.EqualValues(t, test.expFullName, fn)
		})
	}
}

func getSchemaData(t *testing.T, opts ...resourceDataOpt) *schema.ResourceData {
	d := new(schema.Resource)
	d.Schema = iamPolicySchema

	rd := d.TestResourceData()
	scopeData := map[string][]interface{}{}

	for _, o := range opts {
		o(scopeData)
	}

	err := rd.Set(scopeKey, []interface{}{scopeData})
	require.NoError(t, err)

	return rd
}

func withClusterGroup(val string) resourceDataOpt {
	return func(res map[string][]interface{}) {
		res[clusterGroupKey] = []interface{}{
			map[string]string{
				clusterGroupNameKey: val,
			},
		}
	}
}

func withOrg(val string) resourceDataOpt {
	return func(res map[string][]interface{}) {
		res[orgKey] = []interface{}{
			map[string]string{
				orgIDKey: val,
			},
		}
	}
}

func withWorkspace(val string) resourceDataOpt {
	return func(res map[string][]interface{}) {
		res[workspaceKey] = []interface{}{
			map[string]string{
				workspaceNameKey: val,
			},
		}
	}
}

func withCluster(mc, prov, name string) resourceDataOpt {
	return func(res map[string][]interface{}) {
		res[clusterKey] = []interface{}{
			map[string]string{
				managementClusterNameKey: mc,
				provisionerNameKey:       prov,
				clusterNameKey:           name,
			},
		}
	}
}

func withNamespace(mc, prov, cluster, ns string) resourceDataOpt {
	return func(res map[string][]interface{}) {
		res[namespaceKey] = []interface{}{
			map[string]string{
				managementClusterNameKey: mc,
				provisionerNameKey:       prov,
				clusterNameKey:           cluster,
				namespaceNameKey:         ns,
			},
		}
	}
}

func TestGetRoleBindingSchemaData(t *testing.T) {
	d := new(schema.Resource)
	d.Schema = iamPolicySchema

	rd := d.TestResourceData()

	err := rd.Set(
		roleBindingKey,
		[]interface{}{
			map[string]interface{}{
				roleKey: "dummy-role",
				subjectsKey: []interface{}{
					map[string]interface{}{
						subjectNameKey: "sub",
						subjectKindKey: "kind",
					},
					map[string]interface{}{
						subjectNameKey: "sub2",
						subjectKindKey: "kind2",
					},
				},
			},
			map[string]interface{}{
				roleKey: "dummy-role-2",
				subjectsKey: []interface{}{
					map[string]interface{}{
						subjectNameKey: "sub",
						subjectKindKey: "kind",
					},
					map[string]interface{}{
						subjectNameKey: "sub2",
						subjectKindKey: "GROUP",
					},
				},
			},
		})
	require.NoError(t, err)

	got := constructRoleBindingList(rd)

	fmt.Println(len(got))

	for _, each := range got {
		for _, s := range each.Subjects {
			fmt.Println(each.Role, s.Name, *s.Kind)
		}
	}
}
