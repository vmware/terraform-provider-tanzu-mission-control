package policyrecipemutationcommonmodel

const (
	Asterisk   VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope = "*"
	Cluster    VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope = "Cluster"
	Namespaced VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope = "Namespaced"
)

type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope string

func NewVmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope(value VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope) *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope {
	return &value
}
