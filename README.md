# terraform-provider-sumologic
[![Build Status](https://travis-ci.org/erikvanbrakel/terraform-provider-sumologic.svg?branch=master)](https://travis-ci.org/erikvanbrakel/terraform-provider-sumologic)

Terraform provider for https://sumologic.com

## Building the provider

In this section you will see how to build and run terraform-provider-sumologic locally. Please follow the steps below:



## Usage
Download the binary for your platform and architecture from the [releases page](https://github.com/erikvanbrakel/terraform-provider-sumologic/releases). Unpack the zip, and place the `terraform-provider-sumologic` binary in the same directory as `terraform` binary or add a `.terraformrc` file with the provider stanza:

```hcl
providers {
  sumologic = "/PATH/TO/MODULE/ARCH/terraform-provider-sumologic"
}
```

## Usage
See the [documentation][0].

[0]: docs/README.md
