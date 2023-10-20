// Read Tanzu Mission Control Azure AKS cluster : fetch cluster details
data "tanzu-mission-control_akscluster" "tf_aks_cluster" {
  credential_name = "test-azure-credential"     // Required
  subscription    = "test-azure-subscription"   // Required
  resource_group  = "test-azure-resource-group" // Required
  name            = "test-aks-cluster"          // Required
}

// Read location of the cluster received as a data source
output "location" {
  value = data.tanzu-mission-control_akscluster.tf_aks_cluster.spec[0].config[0].location
}
