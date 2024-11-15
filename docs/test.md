<!--
© Broadcom. All Rights Reserved.
The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
SPDX-License-Identifier: MPL-2.0
-->

<!-- markdownlint-disable first-line-h1 no-inline-html -->

<img src="images/icon-color.svg" alt="VMware Tanzu Mission Control" width="150">

# Testing the Terraform Provider for VMware Tanzu Mission Control

## Configuring Environment Variables

> [!NOTE]
> This section is applicable only for VMware Tanzu Mission Control SaaS offering.

Set the environment variables in your IDE configurations or terminal.

Environment variables that are required to be set universally include:`TMC_ENDPOINT`, `VMW_CLOUD_ENDPOINT`, and `VMW_CLOUD_API_TOKEN`.

Example:

```shell
$ export TMC_ENDPOINT = my-org.tmc.cloud.vmware.com
$ export VMW_CLOUD_ENDPOINT = console.tanzu.broadcom.com
```

Environment variables specific to particular resources:

- **Attach Cluster with Kubeconfig and Namespace Resource**: `KUBECONFIG`.
- **Tanzu Kubernetes Grid Service for vSphere workload cluster**:`MANAGEMENT_CLUSTER`, `PROVISIONER_NAME`, `VERSION`, and `STORAGE_CLASS`.
- **Tanzu Kubernetes Grid workload cluster**:`MANAGEMENT_CLUSTER` and `CONTROL_PLANE_ENDPOINT`.

## Running the Acceptance Tests

You can run the acceptance tests by running:

```sh
$ make testacc
```

If you want to run against a specific set of tests, make use of the build-tags. A build tag name is
equivalent to the corresponding resource name.

By default, running acceptance test without explicitly setting `BUILD_TAGS` runs all the acceptance
test. To specifically run acceptances test for a resources, set `BUILD_TAGS` value to corresponding
resource name.

For example: Run acceptance test for `clustergroup` and `namespace` resource.

```sh
$ export BUILD_TAGS = "clustergroup namespace"
$ make acc-test
```

### Test Provider Changes Locally

Please make use of a unique path as provided in the `Makefile` while building the provider with
changes and use the same path in the source while using the provider to test the local changes.

```shell
terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/dev/tanzu-mission-control"
    }
  }
}

provider "tanzu-mission-control" {
  # Configuration options
}
```

## Debugging the Provider

1. Set the environmental variable `TF_LOG` to one of the log levels `TRACE`, `DEBUG`, `INFO`, `WARN` or `ERROR` to capture the logs.

2. Set the environmental variable `TMC_MODE` to `DEV` to capture more granular logs.

3. Using Visual Studio Code, create the `./.vscode/launch.json` file in the project.

    ```json
    {
          "version": "0.2.0",
          "configurations": [
              {
                  "name": "Debug Terraform Provider",
                  "type": "go",
                  "request": "launch",
                  "mode": "debug",
                  // this assumes your workspace is the root of the repo
                  "program": "${workspaceFolder}",
                  "env": {},
                  "args": [
                      "-debug",
                  ]
              }
          ]
      }
    ```

4. Click on **"Run and Debug"** option in Visual Studio Code, This will open a panel on the left side of the editor. Here, you can see a list of configurations for debugging different languages and tools. Select the item labeled as **"Debug Terraform Provider"**. This will launch the debugger and attach it to your provider process. You can now set breakpoints, inspect variables, and step through your code as usual.

5. Check the **"DEBUG CONSOLE"** tab, there you will find the value of `TF_REATTACH_PROVIDERS`, which is a special environment variable that tells Terraform how to connect to the provider's plugin process. You need to set this variable in your shell before running any Terraform commands. For example, you can use the export command as shown below:

    ```sh
    export TF_REATTACH_PROVIDERS='{"vmware/dev/tanzu-mission-control":{"Protocol":"grpc","ProtocolVersion":5,"Pid":1338,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/r9/h_0mgps9053g3tft7t8xh6rh0000gq/T/plugin2483048401"}}}'

    terraform plan
    ```
