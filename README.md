# terraform-provider-tanzu

A custom provider for terraform CLI tool to manage TANZU resources.

// usage/playgo (link to examples)
# Use Cases of TMC Terraform Provider
[use-cases]: https://github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/examples

# Manual Installation

## Cloning the Project

First, you will want to clone the repository to
`github.com/vmware-tanzu/terraform-provider-tanzu-mission-control`:

```sh
mkdir -p $GOPATH/src/github.com/vmware-tanzu/terraform-provider-tanzu-mission-control
cd $GOPATH/src/github.com/vmware-tanzu/terraform-provider-tanzu-mission-control
git clone git@github.com:vmware-tanzu/terraform-provider-tanzu-mission-control.git
```

## Building and Installing the Provider

Recommended golang version is go1.14 onwards.
After the clone has been completed, you can enter the provider directory and build the provider.

```sh
cd github.com/vmware-tanzu/terraform-provider-tanzu-mission-control
make
```

After the build is complete, copy the provider executable `terraform-provider-tanzu` into location specified in your provider installation configuration. Make sure to delete provider lock files that might exist in your working directory due to prior provider usage. Run `terraform init`.
For developing, consider using [dev overrides configuration][dev-overrides]. Please note that `terraform init` should not be used with dev overrides.

[dev-overrides]: https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers

## Utilising TMC provider

Please refer to `examples` folder to perform CRUD operations with TMC provider for various resources
