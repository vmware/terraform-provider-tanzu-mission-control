// Create/ Delete/ Update Tanzu Mission Control organization scoped custom security policy entry
resource "tanzu-mission-control_security_policy" "organization_scoped_custom_security_policy" {
  scope {
    organization{
      organization = "<organization-id>" // Required
    }
  }

  spec {
    input {
      custom {
        audit                        = false // Default: false
        disable_native_psp           = false // Default: false
        allow_privileged_containers  = false // Default: false
        allow_privilege_escalation   = false // Default: false
        allow_host_namespace_sharing = false // Default: false
        allow_host_network           = false // Default: false
        read_only_root_file_system   = false // Default: false

        allowed_host_port_range {
          min = 0 // Default: 0
          max = 65535 // Default: 65535
        }

        allowed_volumes              = [
          "<allowed-volume-1>",
          "<allowed-volume-2>",
          "<allowed-volume-3>"
        ] // Default: ["*"]

        run_as_user {
          rule = "<run-as-user-rule>" // Default: "RunAsAny"

          ranges {
            min = 0 // Default: 0
            max = 65535 // Default: 65535
          }
        }

        run_as_group {
          rule = "<run-as-group-rule>" // Default: "RunAsAny"

          ranges {
            min = 0 // Default: 0
            max = 65535 // Default: 65535
          }
        }

        supplemental_groups {
          rule = "<supplemental-groups-rule>" // Default: "RunAsAny"

          ranges {
            min = 0 // Default: 0
            max = 65535 // Default: 65535
          }
        }

        fs_group {
          rule = "<fs-group-rule>" // Default: "RunAsAny"

          ranges {
            min = 0 // Default: 0
            max = 65535 // Default: 65535
          }
        }

        linux_capabilities {
          allowed_capabilities       = [
            "<allowed-linux-capability-1>",
            "<allowed-linux-capability-2>",
          ] // Default: ["*"]
          required_drop_capabilities = [
            "<required-drop-linux-capability>",
          ] // Default: []
        }

        allowed_host_paths {
          path_prefix = "<allowed-host-path-prefix-1>" // Default: ""
          read_only  = false // Default: false
        }
        allowed_host_paths {
          path_prefix = "<allowed-host-path-prefix-2>" // Default: ""
          read_only  = false // Default: false
        }
        allowed_host_paths {
          path_prefix = "<allowed-host-path-prefix-3>" // Default: ""
          read_only  = false // Default: false
        }

        allowed_se_linux_options {
          level = "<allowed-se-linux-option-level>" // Default: ""
          role  = "<allowed-se-linux-option-role>" // Default: ""
          type  = "<allowed-se-linux-option-type>" // Default: ""
          user  = "<allowed-se-linux-option-user>" // Default: ""
        }

        sysctls {
          forbidden_sysctls = [
            "<forbidden-sysctls-1>",
            "<forbidden-sysctls-1>"
          ] // Default: []
        }

        seccomp {
          allowed_profiles        = [
            "<seccomp-allowed-profile>"
          ] // Default: ["*"]
          allowed_localhost_files = [
            "<seccomp-allowed-localhost-file-1>",
            "<seccomp-allowed-localhost-file-2>"
          ] // Default: ["*"]
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "<label-selector-requirement-key-1>"
        operator = "<label-selector-requirement-operator>"
        values   = [
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
