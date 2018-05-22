package provider

import (
	"testing"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccSumologicHttpSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep {
			{
				Config: testAccSumologicHttpSourceConfig,
			},
		}})
}

var testAccSumologicHttpSourceConfig = `

resource "sumologic_collector" "test" {
  name = "MyCollector"
  description = "MyCollectorDesc"
  category = "Cat"
}

resource "sumologic_http_source" "http" {
  name = "test_http"
  messagePerRequest = false
  category = "source/category"
  collector_id = "${sumologic_collector.test.id}"
}
`
