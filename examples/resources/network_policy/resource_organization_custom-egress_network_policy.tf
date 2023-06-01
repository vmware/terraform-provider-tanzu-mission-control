/*
Organization scoped Tanzu Mission Control network policy with custom-egress input recipe.
This policy is applied to a organization with the custom-egress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_custom-egress_destination-selector_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom_egress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_selector {
              namespace_selector = {
                "ns-key-1" = "ns-val-1"
                "ns-key-2" = "ns-val-2"
              }
              pod_selector = {
                "pod-key-1" = "pod-val-1"
                "pod-key-2" = "pod-val-2"
              }
            }
          }
        }
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values = [
          "api-server",
          "agent-gateway"
        ]
      }
      match_expressions {
        key      = "not-a-component"
        operator = "DoesNotExist"
        values   = []
      }
    }
  }
}

resource "tanzu-mission-control_network_policy" "organization_scoped_custom-egress_destination-ip-block_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom_egress {
        rules {
          ports {
            port     = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_ip {
              ip_block {
                cidr = "192.168.1.1/24"
                except = [
                  "2001:db9::/64",
                ]
              }
            }
          }
        }
        to_pod_labels = {
          "key-1" = "value-1"
          "key-2" = "value-2"
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values = [
          "api-server",
          "agent-gateway"
        ]
      }
      match_expressions {
        key      = "not-a-component"
        operator = "DoesNotExist"
        values   = []
      }
    }
  }
}
