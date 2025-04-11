---
Title: "Provisioning of a Tanzu Kubernetes Cluster"
Description: |-
    An example of provisioning Tanzu Kubernetes Grid Service (Class-based Cluster)
---

# Tanzu Kubernetes Cluster (Class-based Cluster)

The `tanzu-mission-control_tanzu_kubernetes_cluster` resource enables you provision cluster class based clusters on Tanzu Kubernetes Grid Vsphere and Tanzu Kubernetes Grid Service cluster.
For an example usage of the new API resource please refer to [Tanzu Kubernetes Cluster Resource (Class-based Cluster)][Tanzu-Kubernetes-Cluster-Resource-(Class-based Cluster)].

[Tanzu-Kubernetes-Cluster-Resource-(Class-based Cluster)]: https://registry.terraform.io/providers/vmware/tanzu-mission-control/latest/docs/resources/tanzu_kubernetes_cluster

**Prerequisites:**

Before creating a Tanzu Kubernetes Grid Service workload cluster in vSphere Supervisor using this Terraform provider we need the following prerequisites.

- Register the Tanzu Kubernetes Grid Service 2.x management cluster in Tanzu Mission Control.
Note that the Tanzu Kubernetes Grid Service management cluster must be **ready** and **healthy**.
Please refer to [registration of a Supervisor Cluster in vSphere Supervisor.][supervisor-cluster-registration]

- Create a provisioner under the management cluster or reuse the existing providers under the management cluster. Please refer to [working with vSphere Namespaces on a Supervisor Cluster.][vsphere-namespaces]

- Ensure the CSP token used in initialising the terraform provider has the right set of permissions to create a workload cluster.

Once you have the `management cluster name` and `provisioner name` from Tanzu Mission Control, we are all set to provision a workload under the chosen management cluster name using the terraform script.

If you want to manage the already existing workload, Refer to [adding a Workload Cluster into Tanzu Mission Control Management.][add-workload-cluster]
For more information about provisioning a workload, please refer to [provision a Cluster in Tanzu Kubernetes Grid Service][provision-cluster-vsphere] or [provision a Cluster in Tanzu Kubernetes Grid vSphere][provision-cluster-service]

[supervisor-cluster-registration]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03.html#GUID-CC6E721E-43BF-4066-AA0A-F744280D6A03
[vsphere-namespaces]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere-supervisor/7-0/vsphere-with-tanzu-configuration-and-management-7-0.html
[add-workload-cluster]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-78908829-CB4E-459F-AA81-BEA415EC9A11.html
[provision-cluster-vsphere]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-0A1AEC6A-3E5C-424F-8EBC-1DDFC14D2688.html
[provision-cluster-service]: https://techdocs.broadcom.com/us/en/vmware-tanzu/standalone-components/tanzu-mission-control/1-4/tanzu-mission-control-documentation/tanzumc-using-GUID-9435CCAC-F90B-4575-9D73-D26315871C8A.html
