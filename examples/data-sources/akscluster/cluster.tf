# Read Tanzu Mission Control Azure AKS  cluster : fetch cluster details
data "tanzu-mission-control_aks_cluster" "tf_aks_cluster" {
  credential_name = "test-aks-cred-name" // Required
  subscription    = "test-subscription-id"    // Required
  resource_group  = "test-resource-group"  // Required
  name            = "test-cluster-name"    // Required
}