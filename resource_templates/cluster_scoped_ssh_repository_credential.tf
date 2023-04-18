# Create Tanzu Mission Control repositorycredential entry
resource "tanzu-mission-control_repository_credential" "test_credential" {
 name = "resposiroy-credential-name"

 scope {
   cluster {
      cluster_name            = "testcluster" # Required
      provisioner_name        = "attached"     # Default: attached
      management_cluster_name = "attached"     # Default: attached
   }
 }

 spec {
   // this block to have any fields specific to cluster group level resource (ex- fanout config) - custom schema validations will be added for these attributes to restrict them to a particular scope
  
   source_secret_type = "SSH"

   data {
    ssh_key {
     ssh_key = "somesshkey"
     known_hosts = "knowshosts"
    }
   }
 }

 // Since this is not asynchronously updated it might not be needed, the terraform resource’s state might be able to convey if the credential has been created or not.
 status {
   credential_phase = "CREATED"
 }
}
