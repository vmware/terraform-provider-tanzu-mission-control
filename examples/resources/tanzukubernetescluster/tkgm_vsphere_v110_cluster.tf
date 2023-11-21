resource "tanzu-mission-control_tanzu_kubernetes_cluster" "tkgm_cluster" {
  name                    = "CLS_NAME"
  management_cluster_name = "MANAGEMENT_CLS_NAME"
  provisioner_name        = "PROVISIONER_NAME"

  meta {
    description = "description of the cluster"
    labels      = { "key" : "changedvalues" }
  }

  spec {
    cluster_group_name = "default"

    topology {
      version           = "v1.26.5+vmware.2-tkg.1"
      cluster_class     = "tkg-vsphere-default-v1.1.0"
      cluster_variables = jsonencode(local.tkgm_v110_cluster_variables)

      control_plane {
        replicas = 1

        meta {
          labels      = { "key" : "value" }
          annotations = { "key1" : "annotation1" }
        }

        os_image {
          name    = "ubuntu"
          version = "2004"
          arch    = "amd64"
        }
      }

      nodepool {
        name        = "md-0"
        description = "simple small md"

        spec {
          worker_class = "tkg-worker"
          replicas     = 1
          overrides    = jsonencode(local.tkgm_v110_nodepool_a_overrides)

          meta {
            labels      = { "key" : "value" }
            annotations = { "key1" : "annotation1" }
          }

          os_image {
            name    = "ubuntu"
            version = "2004"
            arch    = "amd64"
          }
        }
      }

      network {
        pod_cidr_blocks = [
          "100.96.0.0/11",
        ]
        service_cidr_blocks = [
          "100.64.0.0/13",
        ]
      }

      core_addon {
        type     = "cni"
        provider = "antrea"
      }

      core_addon {
        type     = "cpi"
        provider = "vsphere-cpi"
      }

      core_addon {
        type     = "csi"
        provider = "vsphere-csi"
      }
    }
  }

  timeout_policy {
    timeout             = 60
    wait_for_kubeconfig = true
    fail_on_timeout     = true
  }
}
