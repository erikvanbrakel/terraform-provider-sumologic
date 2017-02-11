package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/erikvanbrakel/terraform-provider-sumologic/go-sumologic"
)

func resourceSumologicPollingSource() *schema.Resource {
	return &schema.Resource {
		Create: resourceSumologicPollingSourceCreate,
		Read: resourceSumologicPollingSourceRead,
		Delete: resourceSumologicPollingSourceDelete,

		Schema: map[string]*schema.Schema {
			"name" : {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content_type" : {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scan_interval" : {
				Type: schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"paused" : {
				Type: schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"collector_id" : {
				Type: schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"authentication" : {
				Type: schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema {
						"access_key" : {
							Type: schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"secret_key" : {
							Type: schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"path" : {
				Type: schema.TypeList,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket_name": {
							Type: schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"path_expression": {
							Type: schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceSumologicPollingSourceCreate(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*sumologic.SumologicClient)

	c.CreatePollingSource(
		d.Get("name").(string),
		d.Get("content_type").(string),
		d.Get("scan_interval").(int),
		d.Get("paused").(bool),
		d.Get("collector_id").(int),
	)

	return resourceSumologicPollingSourceRead(d, meta)
}

func resourceSumologicPollingSourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSumologicPollingSourceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}