# terraform-provider-sumologic
[![Build Status](https://travis-ci.org/erikvanbrakel/terraform-provider-sumologic.svg?branch=master)](https://travis-ci.org/erikvanbrakel/terraform-provider-sumologic)

Terraform provider for https://sumologic.com

## Usage

### Provider init

Place `terraform-provider-sumologic` in the same directory as `terraform` binary or add a `.terraformrc` file with the provider stanza:

```
providers {
  sumologic = "/PATH/TO/MODULE/ARCH/terraform-provider-sumologic"
}
```

### Provider setup

```
provider "sumologic" {
    access_id = ""
    access_key = ""
    environment = "" # your region (eu, us, us2, etc.)
}
```

## Resources

### sumologic_collector

```
resource "sumologic_collector" "collector" {
    name = ""
    description = ""
    category = ""
}
```

### sumologic_http_source

```
resource "sumologic_http_source" "http" {
    name = "" # string
    messagePerRequest = "" # bool
    collector_id = "${sumologic_collector.collector.id}"
}
```

### sumologic_polling_source

```
resource "sumologic_polling_source" "s3_audit" {
    name = "Amazon S3 Audit"
    category = "aws/s3audit"
    content_type = "AwsS3AuditBucket"
    scan_interval = 1
    paused = false
    collector_id = "${sumologic_collector.collector.id}"

    authentication {
        access_key = "AKIAIOSFODNN7EXAMPLE"
        secret_key = "******"
    }

    path {
        bucket_name = "Bucket1"
        path_expression = "*"
    }
}
```
