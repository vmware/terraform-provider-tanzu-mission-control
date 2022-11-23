/*
Cluster scoped Tanzu Mission Control security policy with custom input recipe.
This policy is applied on a cluster with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_security_policy" "cluster_scoped_custom_security_policy" {
  name = "tf-sp-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      custom {
        audit                        = true
        disable_native_psp           = false
        allow_privileged_containers  = true
        allow_privilege_escalation   = true
        allow_host_namespace_sharing = true
        allow_host_network           = true
        read_only_root_file_system   = true

        allowed_host_port_range {
          min = 3000
          max = 5000
        }

        allowed_volumes = [
          "configMap",
          "nfs",
          "vsphereVolume"
        ]

        run_as_user {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        run_as_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        supplemental_groups {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        fs_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        linux_capabilities {
          allowed_capabilities = [
            "CHOWN",
            "IPC_LOCK"
          ]
          required_drop_capabilities = [
            "SYS_TIME"
          ]
        }

        allowed_host_paths {
          path_prefix = "p1"
          read_only   = true
        }
        allowed_host_paths {
          path_prefix = "p2"
          read_only   = false
        }
        allowed_host_paths {
          path_prefix = "p3"
          read_only   = true
        }

        allowed_se_linux_options {
          level = "s0"
          role  = "sysadm_r"
          type  = "httpd_sys_content_t"
          user  = "root"
        }

        sysctls {
          forbidden_sysctls = [
            "kernel.msgmax",
            "kernel.sem"
          ]
        }

        seccomp {
          allowed_profiles = [
            "Localhost"
          ]
          allowed_localhost_files = [
            "profiles/audit.json",
            "profiles/violation.json"
          ]
        }
      }
    }
  }
}
