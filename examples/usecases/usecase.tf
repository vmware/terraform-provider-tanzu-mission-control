/*
  NOTE: Creation of attach cluster depends on cluster-group creation
        Similarly, namespace creation depends on attach cluster and workspace creation
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create workspace
resource "tanzu-mission-control_workspace" "create_workspace" {
  name = "demo-workspace"

  meta {
    description = "Create workspace through terraform"
    labels = {
      "key1" : "value1",
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "create_cluster_group" {
  name = "demo-cluster-group"
}

# Attach a Tanzu Mission Control cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
  name                    = "demo-cluster" // Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config path>" // Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name // Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

# Create namespace with attached set as 'true' (need a running cluster)
resource "tanzu-mission-control_namespace" "create_namespace" {
  name                    = "demo-namespace"                                                  // Required
  cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name // Required
  provisioner_name        = "attached"                                                        // Default: attached
  management_cluster_name = "attached"                                                        // Default: attached

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    workspace_name = tanzu-mission-control_workspace.create_workspace.name // Default: default
    attach         = true
  }
}

# Create Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkgs_workload_cluster" {
  management_cluster_name = "tkgs-terraform"
  provisioner_name        = "test-gc-e2e-demo-ns"
  name                    = "tkgs-workload"

  meta {
    labels = { "key" : "test" }
  }

  spec {
    cluster_group = "default"
    tkg_service_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "172.20.0.0/16",
            ]
          }
          services {
            cidr_blocks = [
              "10.96.0.0/16",
            ]
          }
        }
        storage {
          classes = [
            "wcpglobal-storage-profile",
          ]
          default_class = "wcpglobal-storage-profile"
        }
      }

      distribution {
        version = "v1.21.6+vmware.1-tkg.1.b3d708a"
      }

      topology {
        control_plane {
          class             = "guaranteed-xsmall"
          storage_class     = "tkgs-k8s-obj-policy"
          high_availability = false
          volumes {
            capacity          = 4
            mount_path        = "/var/lib/etcd"
            name              = "etcd-0"
            pvc_storage_class = "tkgs-k8s-obj-policy"
          }
          volumes {
            capacity          = 4
            mount_path        = "/var/lib/etcd"
            name              = "etcd-1"
            pvc_storage_class = "tkgs-k8s-obj-policy"
          }
        }
        node_pools {
          spec {
            worker_node_count = "1"
            tkg_service_vsphere {
              class         = "guaranteed-xsmall"
              storage_class = "tkgs-k8s-obj-policy"
              volumes {
                capacity          = 4
                mount_path        = "/var/lib/etcd"
                name              = "etcd-0"
                pvc_storage_class = "tkgs-k8s-obj-policy"
              }
            }
          }
          info {
            name = "default-nodepool"
          }
        }
      }
    }
  }
}

# Create a Tanzu Kubernetes Grid vSphere workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_vsphere_cluster" {
  management_cluster_name = "tkgm-terraform"
  provisioner_name        = "default"
  name                    = "tkgm-workload"

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value1" }
  }

  spec {
    cluster_group = "default" # Default: default
    tkg_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "172.20.0.0/16",
            ]
          }

          services {
            cidr_blocks = [
              "10.96.0.0/16",
            ]
          }

          api_server_port         = 6443
          control_plane_end_point = "10.191.143.100"
        }

        security {
          ssh_key = "default"
        }
      }

      distribution {
        version = "v1.20.5+vmware.2-tkg.1"

        workspace {
          datacenter        = "/dc0"
          datastore         = "/dc0/datastore/local-0"
          workspace_network = "/dc0/network/VM Network"
          folder            = "/dc0/vm"
          resource_pool     = "/dc0/host/cluster0/Resources"
        }
      }

      topology {
        control_plane {
          vm_config {
            cpu       = "2"
            disk_size = "20"
            memory    = "4096"
          }

          high_availability = false
        }

        node_pools {
          spec {
            worker_node_count = "1"

            tkg_vsphere {
              vm_config {
                cpu       = "2"
                disk_size = "40"
                memory    = "4096"
              }
            }
          }

          info {
            name        = "md-0"
            description = "my nodepool"
          }
        }
      }
    }
  }
}

# Create Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_aws_cluster" {
  management_cluster_name = "tkgm-aws" // Default: attached
  provisioner_name        = "default"            // Default: attached
  name                    = "tkgm-aws-workload"  // Required

  meta {
    description = "new description"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" // Default: default
    tkg_aws {
      settings {
        network {
          cluster {
            pods {
              cidr_blocks = "100.96.0.0/11"
            }

            services {
              cidr_blocks = "100.64.0.0/13"
            }
          }
          provider {
            subnets {
              availability_zone = "us-west-2a"
              cidr_block_subnet = "10.0.0.0/24"
            }
            subnets {
              availability_zone = "us-west-2a"
              cidr_block_subnet = "10.0.1.0/24"
              is_public         = true
            }

            vpc {
              cidr_block_vpc = "10.0.0.0/16"
            }
          }
        }

        security {
          ssh_key = "jumper_ssh_key-sh-1643378-220418-062857" // Required
        }
      }

      distribution {
        region  = "us-west-2"              // Required
        version = "v1.21.2+vmware.1-tkg.2" // Required
      }

      topology {
        control_plane {
          availability_zones = [
            "us-west-2a",
          ]
          instance_type = "m5.large"
        }

        node_pools {
          spec {
            worker_node_count = "2"
            tkg_aws {
              availability_zone = "us-west-2a"
              instance_type     = "m5.large"
              node_placement {
                aws_availability_zone = "us-west-2a"
              }

              version = "v1.21.2+vmware.1-tkg.2"
            }
          }

          info {
            name = "md-0" // Required
          }
        }
      }
    }
  }
}

# Create Tanzu Mission Control nodepool entry
resource "tanzu-mission-control_cluster_node_pool" "create_node_pool" {

  management_cluster_name = "tkgs-terraform"
  provisioner_name        = "test-gc-e2e-demo-ns"
  cluster_name            = "tkgs-nodepool"
  name                    = "terraform-nodepool"

  spec {
    worker_node_count = "3"

    tkg_service_vsphere {
      class         = "best-effort-xsmall"
      storage_class = "gc-storage-profile"
    }
  }
}

# Organization scoped Baseline Security Policy
resource "tanzu-mission-control_security_policy" "organization_scoped_baseline_security_policy" {
  scope {
    organization{
      organization = "dummy-org-id"
    }
  }

  spec {
    input {
      baseline {
        audit              = false
        disable_native_psp = true
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
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

# Cluster group scoped Baseline Security Policy
resource "tanzu-mission-control_security_policy" "cluster_group_scoped_baseline_security_policy" {
  scope {
    cluster_group{
      cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name
    }
  }

  spec {
    input {
      baseline {
        audit              = false
        disable_native_psp = true
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
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

# Cluster scoped Baseline Security Policy
resource "tanzu-mission-control_security_policy" "cluster_scoped_baseline_security_policy" {
  scope {
    cluster{
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name
    }
  }

  spec {
    input {
      baseline {
        audit              = false
        disable_native_psp = true
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
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

# Organization scoped Custom Security Policy
resource "tanzu-mission-control_security_policy" "organization_scoped_custom_security_policy" {
  scope {
    organization{
      organization = "dummy-org-id"
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

        allowed_volumes              = [
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
          allowed_capabilities       = [
            "CHOWN",
            "IPC_LOCK"
          ]
          required_drop_capabilities = [
            "SYS_TIME"
          ]
        }

        allowed_host_paths {
          path_prefix = "p1"
          read_only  = true
        }
        allowed_host_paths {
          path_prefix = "p2"
          read_only  = false
        }
        allowed_host_paths {
          path_prefix = "p3"
          read_only  = true
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
          allowed_profiles        = [
            "Localhost"
          ]
          allowed_localhost_files = [
            "profiles/audit.json",
            "profiles/violation.json"
          ]
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
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

# Cluster group scoped Custom Security Policy
resource "tanzu-mission-control_security_policy" "cluster_group_scoped_custom_security_policy" {
  scope {
    cluster_group{
      cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name
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

        allowed_volumes              = [
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
          allowed_capabilities       = [
            "CHOWN",
            "IPC_LOCK"
          ]
          required_drop_capabilities = [
            "SYS_TIME"
          ]
        }

        allowed_host_paths {
          path_prefix = "p1"
          read_only  = true
        }
        allowed_host_paths {
          path_prefix = "p2"
          read_only  = false
        }
        allowed_host_paths {
          path_prefix = "p3"
          read_only  = true
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
          allowed_profiles        = [
            "Localhost"
          ]
          allowed_localhost_files = [
            "profiles/audit.json",
            "profiles/violation.json"
          ]
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
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

# Cluster scoped Custom Security Policy
resource "tanzu-mission-control_security_policy" "cluster_scoped_custom_security_policy" {
  scope {
    cluster{
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name
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

        allowed_volumes              = [
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
          allowed_capabilities       = [
            "CHOWN",
            "IPC_LOCK"
          ]
          required_drop_capabilities = [
            "SYS_TIME"
          ]
        }

        allowed_host_paths {
          path_prefix = "p1"
          read_only  = true
        }
        allowed_host_paths {
          path_prefix = "p2"
          read_only  = false
        }
        allowed_host_paths {
          path_prefix = "p3"
          read_only  = true
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
          allowed_profiles        = [
            "Localhost"
          ]
          allowed_localhost_files = [
            "profiles/audit.json",
            "profiles/violation.json"
          ]
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
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

# Organization scoped Strict Security Policy
resource "tanzu-mission-control_security_policy" "organization_scoped_strict_security_policy" {
  scope {
    organization{
      organization = "dummy-org-id"
    }
  }

  spec {
    input {
      strict {
        audit              = true
        disable_native_psp = false
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
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

# Cluster group scoped Strict Security Policy
resource "tanzu-mission-control_security_policy" "cluster_group_scoped_strict_security_policy" {
  scope {
    cluster_group{
      cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name
    }
  }

  spec {
    input {
      strict {
        audit              = true
        disable_native_psp = false
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
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

# Cluster scoped Strict Security Policy
resource "tanzu-mission-control_security_policy" "cluster_scoped_strict_security_policy" {
  scope {
    cluster{
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name
      name                    = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name
    }
  }

  spec {
    input {
      strict {
        audit              = true
        disable_native_psp = false
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values   = [
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
