package provider

import (
	"testing"
	"github.com/hashicorp/terraform/helper/resource"
	"os"
)

func TestAccSumologicPollingSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep {
			{
				Config: testAccSumologicPollingSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sumologic_polling_source.s3_audit", "content_type"),
				),
			},
		}})
}

var accessKey,_ = os.LookupEnv("AWS_ACCESS_KEY_ID")
var secretKey,_ = os.LookupEnv("AWS_SECRET_ACCESS_KEY")

var testAccSumologicPollingSourceConfig = `
resource "sumologic_collector" "AWS" {
    name = "AWS2"
    description = "AWS logs"
    category = "category"
}

resource "sumologic_polling_source" "s3_audit" {
    collector_id = "${sumologic_collector.AWS.id}"
    name = "Amazon S3 Audit"
    category = "some/category"
    content_type = "AwsS3AuditBucket"
    scan_interval = 1000
    paused = false

    authentication {
        access_key = "` + accessKey + `"
        secret_key = "` + secretKey + `"
    }

    path {
        bucket_name = "terraform-sumologic-testing-a39dj4f850f"
        path_expression = "*"
    }
}
`
