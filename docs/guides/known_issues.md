# Known Issues in Tanzu Mission Control Provider

## Issue #1
- Using Terraform provider to create or modify ClusterClass based clusters running on VMware vSphere v8.x (or above) using Tanzu Mission Control, results in error and failure.

  Like below example:

```
  Error: Unable to create Tanzu Mission Control cluster entry, name : <cluster-name>: POST request failed with status : 400 Bad Request, response: {"error":"Kubernetes clusters in vSphere 8.0+ must be using ClusterClass, for more information see https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-41DE8F22-56B8-4669-AF3F-E4B4372BDB9E.html","code":9,"message":"Kubernetes clusters in vSphere 8.0+ must be using ClusterClass, for more information see https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-41DE8F22-56B8-4669-AF3F-E4B4372BDB9E.html"}
```


### Workaround
Currently there is no official workaround for using Terraform provider to create ClusterClass based clusters on vSphere 8+ and suggested interim alternate is for using Tanzu Mission Control UI, CLI and/or APIs instead.

### Solution
Native support in the Tanzu Mission Control terraform provider for ClusterClass based K8s clusters running on vSphere 8.x is in development and VMware Tanzu Mission Control will rollout the capability, eliminating the known issue, later in the year.