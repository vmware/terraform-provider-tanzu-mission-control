package providerekscluster

import (
	"testing"

	"github.com/stretchr/testify/require"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/provider_ekscluster"
)

func TestFlattenCluterSpec(t *testing.T) {
	tests := []struct {
		description    string
		getClusterSpec func() *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec
		expected       []interface{}
	}{
		{
			description: "nil spec",
			getClusterSpec: func() *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec {
				return nil
			},
			expected: []interface{}{},
		},
		{
			description:    "full cluster spec",
			getClusterSpec: getClusterSpec,
			expected: []interface{}{
				map[string]interface{}{
					"cluster_group": "tf-cg",
					"proxy":         "proxy1",
					"agent_name":    "test-tf-cluster",
					"eks_arn":       "arn:aws:eks:us-west-2:999999999999:cluster/adopted-cluster",
				},
			},
		},
		{
			description: "proxy is empty",
			getClusterSpec: func() *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec {
				spec := getClusterSpec()
				spec.ProxyName = ""
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"cluster_group": "tf-cg",
					"agent_name":    "test-tf-cluster",
					"eks_arn":       "arn:aws:eks:us-west-2:999999999999:cluster/adopted-cluster",
				},
			},
		},
		{
			description: "agent name is empty",
			getClusterSpec: func() *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec {
				spec := getClusterSpec()
				spec.AgentName = ""
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"cluster_group": "tf-cg",
					"proxy":         "proxy1",
					"eks_arn":       "arn:aws:eks:us-west-2:999999999999:cluster/adopted-cluster",
				},
			},
		},
		{
			description: "eks arn is empty",
			getClusterSpec: func() *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec {
				spec := getClusterSpec()
				spec.Arn = ""
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"cluster_group": "tf-cg",
					"proxy":         "proxy1",
					"agent_name":    "test-tf-cluster",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			res := flattenClusterSpec(test.getClusterSpec())
			require.Equal(t, test.expected, res)
		})
	}
}

func getClusterSpec() *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec {
	return &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec{
		AgentName:        "test-tf-cluster",
		Arn:              "arn:aws:eks:us-west-2:999999999999:cluster/adopted-cluster",
		ClusterGroupName: "tf-cg",
		ProxyName:        "proxy1",
		TmcManaged:       true,
	}
}
