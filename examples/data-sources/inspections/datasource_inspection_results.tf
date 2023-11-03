data "tanzu-mission-control_inspection_results" "demo" {
  management_cluster_name = "MGMT_CLS_NAME"
  provisioner_name        = "PROVISIONER_NAME"
  cluster_name            = "CLS_NAME"
  name                    = "INSPECTION_NAME"
}

output "inspection_report" {
  value = jsondecode(data.tanzu-mission-control_inspection_results.demo.status.report)
}
