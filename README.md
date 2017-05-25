# terraform-provider-sumologic
[![Build Status](https://travis-ci.org/erikvanbrakel/terraform-provider-sumologic.svg?branch=master)](https://travis-ci.org/erikvanbrakel/terraform-provider-sumologic)

Terraform provider for https://sumologic.com

## Building the provider

In this section you will learn how to build and run terraform-provider-sumologic locally. Please follow the steps below:

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.9.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)
- [Sumologic](https://www.sumologic.com/pricing/) Must have an account to be able to get access id and key.

Setting your environment
---------------------
Create the couple environment variables below:
$GOROOT = $HOME/go
$GOPATH = $GOROOT/bin
The GOPATH can be set wherever you want but please read this [topic](https://stackoverflow.com/questions/7970390/what-should-be-the-values-of-gopath-and-goroot) to understand how they work

Clone repository to: `$GOPATH/src/github.com/erikvanbrakel/terraform-provider-sumologic`:

```sh
$ mkdir -p $GOPATH/srcgithub.com/erikvanbrakel; cd $GOPATH/src/github.com/erikvanbrakel/
$ git clone https://github.com/erikvanbrakel/terraform-provider-sumologic.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/erikvanbrakel/terraform-provider-sumologic
$ go build
```

Run the tests

```sh
$ go test -v provider/
```

Generate the provider binary (the binary will be generated in your current directory where you run this command)

```sh
$ go build -o terraform-provider-sumologic
```


## Usage
Download the binary for your platform and architecture from the [releases page](https://github.com/erikvanbrakel/terraform-provider-sumologic/releases). Unpack the zip, and place the `terraform-provider-sumologic` binary in the same directory as `terraform` binary or add a `.terraformrc` file with the provider stanza:

```hcl
providers {
  sumologic = "/PATH/TO/MODULE/ARCH/terraform-provider-sumologic"
}
```

## Documentation
See the [documentation][0].

[0]: docs/README.md
