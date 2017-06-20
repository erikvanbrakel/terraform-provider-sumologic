# terraform-provider-sumologic
[![Build Status](https://travis-ci.org/erikvanbrakel/terraform-provider-sumologic.svg?branch=master)](https://travis-ci.org/erikvanbrakel/terraform-provider-sumologic)

Terraform provider for https://sumologic.com

## Building the provider

In this section you will learn how to build and run terraform-provider-sumologic locally. Please follow the steps below:

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.9.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)
-   [Sumologic](https://www.sumologic.com/pricing/) Must have an account to be able to get access id and key.

Setting your environment
---------------------
Create the couple environment variables below:

$GOROOT = $HOME/go

$GOPATH = $GOROOT/bin

The GOPATH can be set wherever you want but please read this [topic](https://stackoverflow.com/questions/7970390/what-should-be-the-values-of-gopath-and-goroot) to understand how they workasdasd

Clone repository to: `$GOPATH/src/github.com/erikvanbrakel/terraform-provider-sumologic`
Terraform provider for https://sumologic.com

## Installation
Download the binary for your platform and architecture from the [releases page](https://github.com/erikvanbrakel/terraform-provider-sumologic/releases). Unpack the zip, and place the `terraform-provider-sumologic` binary in the same directory as `terraform` binary or add a `.terraformrc` file with the provider stanza:

```hcl
providers {
  sumologic = "/PATH/TO/MODULE/ARCH/terraform-provider-sumologic"
}
```

## Usage
See the [documentation][0].

[0]: docs/README.md
