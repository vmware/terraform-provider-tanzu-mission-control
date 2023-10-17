resource "tanzu-mission-control_mutation_policy" "cluster_group_pod_security_mutation_policy" {
  name = "tf-mutation-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      pod_security {
        allow_privilege_escalation {
          condition = "Always"
          value     = true
        }
        capabilities_add {
          operation = "merge"
          values    = ["AUDIT_CONTROL", "AUDIT_WRITE"]
        }
        capabilities_drop {
          operation = "merge"
          values    = ["AUDIT_WRITE"]
        }
        fs_group {
          condition = "Always"
          value     = 0
        }
        privileged {
          condition = "Always"
          value     = true
        }
        read_only_root_filesystem {
          condition = "Always"
          value     = true
        }
        run_as_group {
          condition = "Always"
          value     = 0
        }
        run_as_non_root {
          condition = "Always"
          value     = true
        }
        run_as_user {
          condition = "Always"
          value     = 0
        }
        se_linux_options {
          condition = "Always"
          level     = "test"
          user      = "test"
          role      = "test"
          type      = "test"
        }
        supplemental_groups {
          condition = "Always"
          values    = [0, 1, 2, 3]
        }
      }
    }
    namespace_selector {
      match_expressions = [
        {
          key      = "component"
          operator = "NotIn"
          values   = [
            "api-server",
            "agent-gateway"
          ]
        },
      ]
    }
  }
}
