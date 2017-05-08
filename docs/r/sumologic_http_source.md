# sumologic_http_source
Provides a [Sumologic HTTP source][1].

__IMPORTANT:__ The endpoint is stored in plain-text in the state. This could be a potential security issue.

## Example Usage
```hcl
resource "sumologic_http_sourc" "http_source" {
    name              = "HTTP"
    messagePerRequest = true
    category          = "my/source/category"
    collector_id      = "${sumologic_collector.collector.id}"
}

resource "sumologic_collector" "collector" {
    name        = "my-collector"
    description = "Just testing this"
}
```

## Argument reference
The following arguments are supported:
- `name` - (Required) The name of the source. This is required, and has to be unique in the scope of the collector. Changing this will force recreation the source.
- `collector_id` - (Required) The ID of the collector to attach this source to.
- `category` - (Required) The source category this source logs to.
- `messagePerRequest` - (Required) When set to `true`, will create one log message per HTTP request.

## Attributes reference
The following attributes are exported:
- `id` - The internal ID of the source.
- `url` - The HTTP endpoint to use for sending data to this source.

[Back to Index][0]

[0]: ../README.md
[1]: https://help.sumologic.com/Send_Data/Sources/02Sources_for_Hosted_Collectors/HTTP_Source
