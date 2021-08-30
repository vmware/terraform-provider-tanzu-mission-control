# terraform-provider-tanzu

A custom provider for terraform cli tool to manage TANZU resources

# Manual Installation

## Cloning the Project

First, you will want to clone the repository to
`gitlab.eng.vmware.com/olympus/terraform-provider-tanzu`:

```sh
mkdir -p $GOPATH/src/gitlab.eng.vmware.com/olympus/terraform-provider-tanzu
cd $GOPATH/src/gitlab.eng.vmware.com/olympus/terraform-provider-tanzu
git clone git@gitlab.eng.vmware.com:olympus/terraform-provider-tanzu.git
```

## Building and Installing the Provider

Recommended golang version is go1.14 onwards.
After the clone has been completed, you can enter the provider directory and build the provider.

```sh
cd gitlab.eng.vmware.com/olympus/terraform-provider-tanzu
make
```

After the build is complete, copy the provider executable `terraform-provider-tanzu` into location specified in your provider installation configuration. Make sure to delete provider lock files that might exist in your working directory due to prior provider usage. Run `terraform init`.
For developing, consider using [dev overrides configuration][dev-overrides]. Please note that `terraform init` should not be used with dev overrides.

[dev-overrides]: https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers

## Utilising TMC provider

Please refer to `examples` folder to perform CRUD operations with TMC provider for various resources
