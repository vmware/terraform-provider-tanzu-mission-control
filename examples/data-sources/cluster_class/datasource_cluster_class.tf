data "tanzu-mission-control_cluster_class" "demo" {
  name                    = "CLUSTER_CLASS_NAME"
  management_cluster_name = "MGMT_CLS_NAME"
  provisioner_name        = "PROVISIONER_NAME"
}

output "cluster_class_variables_schema" {
  value = data.tanzu-mission-control_cluster_class.demo.variables_schema
}

output "cluster_class_variables_template" {
  value = data.tanzu-mission-control_cluster_class.demo.variables_template
}
