 data "tanzu-mission-control_inspections" "demo" {
   management_cluster_name = "MGMT_CLS_NAME"
   provisioner_name        = "PROVISIONER_NAME"
   cluster_name            = "CLS_NAME"
 }

 output "inspections" {
   value = data.tanzu-mission-control_inspections.demo.inspections
 }
