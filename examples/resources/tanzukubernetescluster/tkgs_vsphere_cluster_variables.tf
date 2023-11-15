locals {
  tkgs_cluster_variables = {
    "controlPlaneCertificateRotation" : {
      "activate" : true,
      "daysBefore" : 30
    },
    "defaultStorageClass" : "k8s-storage-policy-vsan",
    "defaultVolumeSnapshotClass" : "volumesnapshotclass-delete",
    "nodePoolLabels" : [

    ],
    "nodePoolVolumes" : [
      {
        "capacity" : {
          "storage" : "20G"
        },
        "mountPath" : "/var/lib/containerd",
        "name" : "containerd",
        "storageClass" : "k8s-storage-policy-vsan"
      },
      {
        "capacity" : {
          "storage" : "20G"
        },
        "mountPath" : "/var/lib/kubelet",
        "name" : "kubelet",
        "storageClass" : "k8s-storage-policy-vsan"
      }
    ],
    "ntp" : "172.16.20.10",
    "storageClass" : "k8s-storage-policy-vsan",
    "storageClasses" : [
      "k8s-storage-policy-vsan"
    ],
    "vmClass" : "best-effort-medium"
  }

  tkgs_nodepool_a_overrides = {
    "nodePoolLabels" : [
      {
        "key" : "sample-worker-label",
        "value" : "value"
      }
    ],
    "storageClass" : "k8s-storage-policy-vsan",
    "vmClass" : "best-effort-medium"
  }
}
