// Create/ Delete/ Update Tanzu Mission Control workspace scoped custom-ingress network policy entry
resource "tanzu-mission-control_network_policy" "workspace_scoped_custom-ingress_source-ip_network_policy" {
  name = "<network-policy-name>"

  scope {
    workspace {
      workspace = "<workspace-name>" // Required
    }
  }

  spec {
    input {
      custom_ingress {
        rules {
          ports {
            port     = "<port-number>"
            protocol = "<TCP/UDP>"
          }
          rule_spec {
            custom_ip {
              ip_block {
                cidr = "<ip-cidr>"
                except = [
                  "<ip-excluded>",
                ]
              }
            }
          }
        }
        to_pod_labels = {
          "<pod-label-key-1>" = "<pod-label-value-1>"
          "<pod-label-key-2>" = "<pod-label-value-2>"
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "<label-selector-requirement-key-1>"
        operator = "<label-selector-requirement-operator>"
        values = [
          "<label-selector-requirement-value-1>",
          "<label-selector-requirement-value-2>"
        ]
      }
      match_expressions {
        key      = "<label-selector-requirement-key-2>"
        operator = "<label-selector-requirement-operator>"
        values   = []
      }
    }
  }
}

resource "tanzu-mission-control_network_policy" "workspace_scoped_custom-ingress_source-selector_network_policy" {
  name = "<network-policy-name>"

  scope {
    workspace {
      workspace = "<workspace-name>" // Required
    }
  }

  spec {
    input {
      custom_ingress {
        rules {
          ports {
            port     = "<port-number>"
            protocol = "<TCP/UDP>"
          }
          rule_spec {
            custom_selector {
              namespace_selector = {
                "<namespace-selector-key-1>" = "<namespace-selector-value-1>"
                "<namespace-selector-key-2>" = "<namespace-selector-value-2>"
              }
              pod_selector = {
                "<pod-selector-key-1>" = "<pod-selector-value-1>"
                "<pod-selector-key-2>" = "<pod-selector-value-2>"
              }
            }
          }
        }
        to_pod_labels = {
          "<pod-label-key-1>" = "<pod-label-value-1>"
          "<pod-label-key-2>" = "<pod-label-value-2>"
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "<label-selector-requirement-key-1>"
        operator = "<label-selector-requirement-operator>"
        values = [
          "<label-selector-requirement-value-1>",
          "<label-selector-requirement-value-2>"
        ]
      }
      match_expressions {
        key      = "<label-selector-requirement-key-2>"
        operator = "<label-selector-requirement-operator>"
        values   = []
      }
    }
  }
}
