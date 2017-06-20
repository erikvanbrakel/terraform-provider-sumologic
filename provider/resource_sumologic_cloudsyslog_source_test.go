package provider

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccSumologicCloudsyslogSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicCloudsyslogSourceConfig,
			},
		}})
}

var testAccSumologicCloudsyslogSourceConfig = `

resource "sumologic_collector" "test" {
  name = "MyCollector"
  description = "MyCollectorDesc"
  category = "Cat"
}

resource "sumologic_cloudsyslog_source" "cloudsyslog" {
  name = "test_cloudsyslog"
  category = "source/category"
  collector_id = "${sumologic_collector.test.id}"
}
`
