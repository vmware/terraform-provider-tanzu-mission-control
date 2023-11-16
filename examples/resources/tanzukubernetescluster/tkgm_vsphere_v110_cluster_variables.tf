locals {
  tkgm_v110_cluster_variables = {
    vcenter = {
      server       = "demo-vc-01.server.demo"
      datacenter   = "/Demo-Datacenter"
      datastore    = "/Demo-Datacenter/datastore/vsanDatastore"
      resourcePool = "/Demo-Datacenter/host/Demo-Cluster/Resources"
      folder       = "/Demo-Datacenter/vm/LABS/folder"
      network      = "/Demo-Datacenter/network/tkg-static-ips"
      cloneMode    = "fullClone"
      template     = "/Demo-Datacenter/vm/Templates/TKG-M/Ubuntu/ubuntu-2004-efi-kube-v1.25.7+vmware.2"
    }

    identityRef = {
      kind = "VSphereClusterIdentity"
      name = "tkg-vc-default"
    }

    user = {
      sshAuthorizedKeys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC4WyexTrHlHayExsx5sv6D/9LV+EsKzK0rekJowoFAk00LGZBCtf5+KMcRN1sVyJNfXZ68uuZCp/FHzjTl9yq4ViZaYNv4MjM4rEtpPfV7nrJeCJnGLZZNhBC5iQ5rnDSIW0ReC5vMERU71twAOHFiA019MSD9LfJE6LA/nmTc7zmYBBLtWuKAPbQBeJNM8UweNvXf6tvH82tkBO/emh4HaaaOyrypPOq8SdAD485I1/5m2dhFX52kUIn6R99AzYXE4sUbqd4wa4lKfQNhxeLYVry7bP/A1jAUWft66GhReKcNox9epyhGCm/icAHcjCCbmMOu0cINeZMPF4BvUQ/q+1j+wGqviNwbOa4Rr2VeO64xFi/ChwPLjajiHdOssBnwXzsXkBRYpWYH/gPLD3Y6W4wrjU958oQcMo2LtJ1pB7MIO9X30mDX0Qn/yTRFHqJ3Tm6ZAN+2pEXWISEEGGg1j6QRZ2TZQKL7rDD1nDQSqU7nH46KsaZ01Y4LU9SXZmk= k8s@tce-admin-vm"]
    }

    aviAPIServerHAProvider = true
    vipNetworkInterface    = "eth0"

    worker = {
      machine = {
        diskGiB   = 50
        memoryMiB = 8192
        numCPUs   = 4
      }
      network = {
        mtu           = 0
        nameservers   = ["10.100.100.100"]
        searchDomains = ["server.demo"]
      }
    }

    controlPlane = {
      machine = {
        diskGiB   = 50
        memoryMiB = 8192
        numCPUs   = 4
      }
      network = {
        mtu           = 0
        nameservers   = ["10.100.100.100"]
        searchDomains = ["server.demo"]
      }
    }

    controlPlaneCertificateRotation = {
      activate   = true
      daysBefore = 90
    }

    network = {
      addressesFromPools = [
        {
          apiGroup = "ipam.cluster.x-k8s.io"
          kind     = "InClusterIPPool"
          name     = "inclusterippool"
        }
      ]
    }
  }

  tkgm_v110_nodepool_a_overrides = {
    nodePoolLabels = [
      {
        key   = "nodepool"
        value = "md-0"
      }
    ]

    worker = {
      machine = {
        diskGiB   = 40
        memoryMiB = 8192
        numCPUs   = 4
      }
      network = {
        nameservers   = ["10.100.100.100"]
        searchDomains = ["server.demo"]
      }
    }
  }
}
