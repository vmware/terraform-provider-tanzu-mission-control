resource "tanzu-mission-control_cluster_class" "demo-tkg-vsphere-cluster-class" {
    name        = "TKG-vspehre-cluster-class"
    description = ""
    spec {
      tkg_vsphere_v100 {
         vcenter { #required
               server       = "" #required #"pattern": "^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\\-]*[A-Za-z0-9])$",
               datacenter   = "" #required
               resourcePool = "" #required
               folder       = "" #required
               network      = "VMNetwork"
               datastore    = "" #required
               cloneMode    = "fullClone"
               tlsThumbprint = ""
               storagePolicyID = ""
               template = ""
               }
         identityRef { #required
               kind = "" #required
               name = "" #required
         }
         user { #required
               sshAuthorizedKeys = [""]
         }
         aviAPIServerHAProvider = false #required
         apiServerEndpoint = "" #"pattern": "^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\\-]*[A-Za-z0-9])$"
         apiServerPort = integer # >=1
         vipNetworkInterface = "eth0" #required
         controlPlane { #required
            machine {
               customVMXKeys {}
               diskGiB = 40 # >=1
               memoryMiB = 8192 # >=1
               numCPUs = 2 # >=1
            }
            network {
               nameservers = [""] #"pattern": "^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\\-]*[A-Za-z0-9])$",
               searchDomains = [""]
            }
            nodeLabels = [
               {
                  key = ""
                  value = ""
               }
            ]
         }
         worker { #required
            machine {
               customVMXKeys {string}
               diskGiB = 40 # >=1
               memoryMiB = 4096 # >=1
               numCPUs = 2 # >=1
            }
            network {
               nameservers = [""] # "pattern": "^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\\-]*[A-Za-z0-9])$",
               searchDomains = [""]
            }
         }
         nodePoolLabels = [
               {
                  key = "" # "maxLength": 317, "pattern": "^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*\\/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$"
                  value = "" # "maxLength": 63, "pattern": "^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$"
               }
         ]
         additionalFQDN = [""]
         trust {
            additionalTrustedCAs = [
               {
                  data = "" #required
                  name = "" #required
               }
            ]
         }
         controlPlaneKubeletExtraArgs {string}
         controlPlaneCertificateRotation {
            activate = true
            daysBefore = 90 # >=7
         }
         proxy {
            httpProxy = "" #required
            httpsProxy = "" #required
            noProxy = [""]
            systemWide = false
         }
         controlPlaneTaint = true
         pci {
            controlPlane {
               devices = [
                  {
                     deviceId = number #required
                     vendorId = number #required
                  }
               ]
               hardwareVersion = "vmx-15"
            }
            worker {
               devices = [
                  {
                     deviceId = number #required
                     vendorId = number #required
                  }
               ]
               hardwareVersion = "vmx-15"
            }
         }
         imageRepository {
            host = ""
            tlsCertificateValidation = true
         }
         kubeControllerManagerExtraArgs {string}
         customTDNFRepository {
            certificate = ""
         }
         network {
            addressesFromPools = [
               {
                  apiGroup = string #required
                  kind = string #required
                  name = string #required
               }
            ]
            ipv6Primary = false
         }
         kubeVipLoadBalancerProvider = false
         auditLogging {
            enabled = false
         }
         apiServerExtraArgs {string}
         additionalImageRegistries = [
            {
               caCert = string
               host = string #required
               skipTlsVerify = false
            }
         ]
         ntpServers = [string] #"pattern": "^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\\-]*[A-Za-z0-9])$"
         workerKubeletExtraArgs {string}
         podSecurityStandard {
            audit = string # ["", "privileged", "baseline", "restricted"]
            auditVersion = "v1.24"
            deactivated = boolean
            enforce = string # ["", "privileged", "baseline", "restricted"]
            enforceVersion = "v1.24"
            exemptions {
               namespaces = [string]
            }
            warn = string # ["", "privileged", "baseline", "restricted"]
            warnVersion = "v1.24"
         }
         tlsCipherSuites = "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
         etcdExtraArgs {string}
         kubeSchedulerExtraArgs {string}
         eventRateLimitConf = ""
      }
   }
}
 
resource "tanzu-mission-control_cluster_class" "demo-tkg-supervisor-cluster-class" {
    name        = "TKG-supervisor-cluster-class"
    description = ""
    spec {
    vsphere_tanzu_v100 {
                vmClass = "best-effort-medium"
                storageClass = "tkgs-k8s-obj-policy"
                }
                // TODO : Add all variables and mark required ones.
     }
}


/*

*/

// Create Tanzu Mission Control classy workload cluster
resource "tanzu-mission-control_tanzu_kubernetes_cluster" "demo_tanzu_kubernetes_cluster" {
    name                    = "<cls-name>"                         // Required
    management_cluster_name = "<mc-name>"                          // Required
    provisioner_name        = "<attached>"                         // Required
 
    meta {
        description = "description of the cluster"
        labels      = { "key" : "value" }
    }
     
    spec {                                               // Required
        cluster_group_name = "<cluster-group>"           //// Required? default tied to mgmt cluster
        tmc_managed        = <true or false>
        proxy_name         = "<proxy>"                  // Optional
        image_registry     = "<image-registry>"         // Optional
        topology {
            version       = "<kubernetes-version>"      // Required
            cluster_class = "<cluster-class-name>"      // Required
            control_plane {
                replicas = "replicas"                   // Required
                meta {
                    labels      = { "key" : "value" }
                    annotations = {"key1:annotation1", "key2:annotation2"}
                }
                os_image {
                    name    = "name"
                    version = "version"
                    arch    = "arch"
                }
            }
            node_pools = [
                {
                    name        = "<nodepool-name>"                        // Required
                    description = "<description>"
                    spec {
                        class          = "<machine-deployment-class-name>" // Required
                        replicas       = "<replicas>"                      // Required
                        failure_domain = "failure_domain"                    
                        overrides {
                            tkg_vsphere {
                                nodePoolLabels = { "key1" : "value1" }
                                worker {
                                    machine {
                                        disk   = "30"
                                        memory = "4096"
                                        numcpu = "2"
                                    }
                                }  
                            }   
                        }
                        meta {
                            labels = { "key" : "value" }
                            annotations = {"key1:annotation1", "key2:annotation2"}
                        }
                        os_image {
                            name = "name"
                            version = "version"
                            arch = "arch"
                        }
                    }
                }
            ]
            cluster_variables {
                variables = tanzu-mission-control_cluster_class.demo-tkg-vsphere-cluster-class.outputs
            }
            network {      
                pod_cidr_blocks     = [
                                        "<pods-cidr-blocks>",     // Required
                                      ]      
                service_cidr_blocks = [
                                        "<services-cidr-blocks>", // Required
                                      ]
                service_domain      = "service-domain-name"      
            }    
            core_addons = [
                {
                 type = "type"          // Required   // ex:cni
                 provider = "provider"  // Required   // ex:antrea
                }
            ]
        }
    }
}
 
// Create Tanzu Mission Control classy workload cluster
resource "tanzu-mission-control_tanzu_kubernetes_cluster" "demo_tanzu_kubernetes_cluster" {
    name                    = "<cls-name>"                         // Required
    management_cluster_name = "<mc-name>"                          // Required
    provisioner_name        = "<attached>"                         // Required
 
    meta {
        description = "description of the cluster"
        labels      = { "key" : "value" }
    }
     
    spec {                                               // Required
        cluster_group_name = "<cluster-group>"
        tmc_managed        = <true or false>
        proxy_name         = "<proxy>"                  // Optional
        image_registry     = "<image-registry>"         // Optional
        topology {
            version       = "<kubernetes-version>"      // Required
            cluster_class = "<cluster-class-name>"      // Required
            control_plane {
                replicas = "replicas"                   // Required
                meta {
                    labels      = { "key" : "value" }
                    annotations = {"key1:annotation1", "key2:annotation2"}
                }
                os_image {
                    name    = "name"
                    version = "version"
                    arch    = "arch"
                }
            }
            node_pools = [
                {
                    name        = "<nodepool-name>"                        // Required
                    description = "<description>"
                    spec {
                        class          = "<machine-deployment-class-name>" // Required
                        replicas       = "<replicas>"                      // Required
                        failure_domain = "failure_domain"                    
                        overrides {
                            vsphere_tanzu {
                                nodePoolLabels = { "key1" : "value1" }
                                vmClass = "best-effort-medium"
                                storageClass = "tkgs-k8s-obj-policy"
                                nodePoolTaints = [
                                    {
                                        key = ""
                                        value = ""
                                        effect = "NoSchedule"
                                        timeAdded =
                                    }
                                ]
                                nodePoolVolumes = [
                                    {
                                        name = ""
                                        mountPath = ""
                                        storageClass = ""
                                        capacity {
                                            storage = ""
                                        }
                                    }
                                ]
                            }
                        }              
                        meta {
                            labels = { "key" : "value" }
                            annotations = {"key1:annotation1", "key2:annotation2"}
                        }
                        os_image {
                            name = "name"
                            version = "version"
                            arch = "arch"
                        }
                    }
                }
            ]
            cluster_variables {
                variables = tanzu-mission-control_cluster_class.demo-tkg-vsphere-cluster-class.outputs 
            }    
            network {      
                pod_cidr_blocks     = [
                                        "<pods-cidr-blocks>",     // Required
                                      ]      
                service_cidr_blocks = [
                                        "<services-cidr-blocks>", // Required
                                      ]
                service_domain      = "service-domain-name"      
            }    
            core_addons = [
                {
                 type = "type"          // Required   // ex:cni
                 provider = "provider"  // Required   // ex:antrea
                }
            ]
        }
    }
}
