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

or use the following environment variables to set credentials:

- `SL_ACCESSID`: Access ID of the access key pair in Sumologic
- `SL_ACCESSKEY`: Access Key of the access key pair in Sumologic

## Resources

### sumologic\_collector

```
resource "sumologic_collector" "collector" {
    name = ""
    description = ""
    category = ""
}
```

#### Argument reference

- `name`: Name of the HTTP source
- `description`: Long description of the collector
- `category`: Category name of the collector

#### Attributes reference

- `id`: The collector id
- `name`: Name of the HTTP source
- `description`: Long description of the collector
- `category`: Category name of the collector

### sumologic\_http\_source

```
resource "sumologic_http_source" "http" {
    name = "" # string
    messagePerRequest = "" # bool
    collector_id = "${sumologic_collector.collector.id}"
}
```

#### Argument reference

- `name`: Name of the HTTP source
- `messagePerRequest`: When set to true, only a single message will be sent for each HTTP request.
- `collector_id`: Id of the collector terraform resource

#### Attributes reference

- `name`: Name of the HTTP source
- `messagePerRequest`: When set to true, only a single message will be sent for each HTTP request.
- `collector_id`: Id of the collector terraform resource
- `url`: Collector endpoint URL **Public endpoint will be stored in plain text in tfstate**

### sumologic\_polling\_source

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
