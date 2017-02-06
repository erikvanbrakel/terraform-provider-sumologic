package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"os"

	sumo "../go-sumologic"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider {
		Schema: map[string]*schema.Schema {
			"access_id": {
				Type: schema.TypeString,
				Optional: true,
				Default: os.Getenv("SL_ACCESSID"),
			},
			"access_key": {
				Type: schema.TypeString,
				Optional: true,
				Default: os.Getenv("SL_ACCESSKEY"),

			},
			"environment": {
				Type: schema.TypeString,
				Optional: true,
				Default: "eu",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sumologic_collector" : resourceSumologicCollector(),
			"sumologic_http_source" : resourceSumologicHttpSource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return sumo.NewClient(
		d.Get("access_id").(string),
		d.Get("access_key").(string),
		d.Get("environment").(string),
	)
}

