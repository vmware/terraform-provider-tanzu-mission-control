# Terraform VMware Tanzu Mission Control Provider

This is the repository for the Terraform Tanzu Mission Control Provider and can be used with
[VMware Tanzu Mission Control][vmware-tanzu-mission-control].

For general information about Terraform, visit the [official website][hashicorp-terraform] and the [GitHub project page.][terraform-github]

[vmware-tanzu-mission-control]: https://tanzu.vmware.com/mission-control
[hashicorp-terraform]: https://www.terraform.io
[terraform-github]: https://github.com/hashicorp/terraform

# Using the Provider

The latest version of this provider requires Terraform v0.12 or higher to run.

Note that you need to run `terraform init` to fetch the provider before
deploying.

### Controlling the provider version

Note that you can also control the provider version. This requires the use of a
`provider` block in your Terraform configuration if you have not added one
already.

The syntax is as follows:

```hcl
terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
      version = "1.0.0"
    }
  }
}

provider "tanzu-mission-control" {
  # Configuration options
}
```

Version locking uses a pessimistic operator, so this version lock would mean anything within the 1.x namespace, including or after 1.0.0.
[Read more][provider-vc] on provider version control.

[provider-vc]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/index.html

# Manual Installation

## Cloning the Project

First, you will want to clone the repository to
`github.com/vmware/terraform-provider-tanzu-mission-control`:

```sh
mkdir -p $GOPATH/src/github.com/vmware/terraform-provider-tanzu-mission-control
cd $GOPATH/src/github.com/vmware/terraform-provider-tanzu-mission-control
git clone git@github.com:vmware/terraform-provider-tanzu-mission-control.git
```

## Building and Installing the Provider

Recommended golang version is go1.14 onwards.
After the clone has been completed, you can enter the provider directory and build the provider.

```sh
cd github.com/vmware/terraform-provider-tanzu-mission-control
make
```

After the build is complete, copy the provider executable `terraform-provider-tanzu` into location specified in your provider installation configuration. Make sure to delete provider lock files that might exist in your working directory due to prior provider usage. Run `terraform init`.
For developing, consider using [dev overrides configuration][dev-overrides]. Please note that `terraform init` should not be used with dev overrides.

[dev-overrides]: https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers

# Developing the Provider

**NOTE:** Before you start work on a feature, please make sure to check the
[issue tracker][gh-issues] and existing [pull requests][gh-prs] to ensure that
work is not being duplicated. For further clarification, you can also ask in a
new issue.

[gh-issues]: https://github.com/vmware/terraform-provider-tanzu-mission-control/issues
[gh-prs]: https://github.com/vmware/terraform-provider-tanzu-mission-control/pulls

If you wish to work on the provider, you'll first need [Go][go-website]
installed on your machine (version 1.14+ is recommended). You'll also need to
correctly setup a [GOPATH][gopath], as well as adding `$GOPATH/bin` to your
`$PATH`.

[go-website]: https://golang.org/
[gopath]: http://golang.org/doc/code.html#GOPATH

See [Manual Installation](#manual-installation) for details on building the
provider.

# Testing the Provider

## Flattening and Helper Tests
Run the command:
```sh
$ make test
```

## Acceptance Tests
**NOTE:** This block is applicable only for Tanzu Mission Control SaaS offering.
### Configuring Environment Variables:
Set the environment variables in your IDE configurations or Terminal.
Environment variables that are required to be set universally are `TMC_ENDPOINT`, `VMW_CLOUD_ENDPOINT` and `VMW_CLOUD_API_TOKEN`.

Example:

```shell
$ export TMC_ENDPOINT = my-org.tmc.cloud.vmware.com
$ export VMW_CLOUD_ENDPOINT = console.cloud.vmware.com
```

Environment variables specific to particular resources:

- **Attach Cluster with Kubeconfig and Namespace Resource** - `KUBECONFIG`  
- **Tanzu Kubernetes Grid  Service for vSphere workload cluster** - `MANAGEMENT_CLUSTER`, `PROVISIONER_NAME`, `VERSION` and `STORAGE_CLASS`.
- **Tanzu Kubernetes Grid workload cluster** - `MANAGEMENT_CLUSTER` and `CONTROL_PLANE_ENDPOINT`.

### Running the Test:
Run the command:
```sh
$ make acc-test
```

### Test provider changes locally
Please make use of a unique path as provided in the `Makefile` while building the provider with changes 
and kindly use the same path in the source while using the provider to test the local changes.

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

[here]: https://www.terraform.io/internals/debugging
## Debugging Provider
Please set the environmental variable `TF_LOG` to one of the log levels `TRACE`, `DEBUG`, `INFO`, `WARN` or `ERROR` to capture the logs. More details in the link [here].

Set the environmental variable `TMC_MODE` to `DEV` to capture more granular logs.

## Provider Documentation

Tanzu Mission Control Terraform provider documentation is autogenerated using [tfplugindocs][tfplugindocs-link].

Use the tfplugindocs tool to generate documentation for your provider in the format required by the Terraform Registry.
The plugin will read the descriptions and schema of each resource and data source in your provider and generate the relevant Markdown files for you.

[tfplugindocs-link]: https://github.com/hashicorp/terraform-plugin-docs
## Using Tanzu Mission Control Provider

Please refer to `examples` folder to perform CRUD operations with Tanzu Mission Control provider for various resources

# Support

The Tanzu Mission Control Terraform provider is now VMware supported as well as community supported. For bugs and feature requests please open a Github Issue and label it appropriately or contact VMware support.

# License

Copyright © 2015-2022 VMware, Inc. All Rights Reserved.

The Tanzu Mission Control Terraform provider is available under [MPL2.0 license](https://github.com/vmware/terraform-provider-tanzu-mission-control/blob/main/LICENSE).