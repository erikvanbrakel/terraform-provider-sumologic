package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	sumologic "github.com/erikvanbrakel/terraform-provider-sumologic/go-sumologic"
	"strconv"
)

func resourceSumologicHttpSource() *schema.Resource {
	return &schema.Resource {
		Create: resourceSumologicHttpSourceCreate,
		Read: resourceSumologicHttpSourceRead,
		Delete: resourceSumologicHttpSourceDelete,

		Schema: map[string]*schema.Schema {
			"name" : {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"messagePerRequest" : {
				Type: schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"collector_id" : {
				Type: schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSumologicHttpSourceCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	response, err := c.CreateHttpSource(
		d.Get("name").(string),
		d.Get("messagePerRequest").(bool),
		d.Get("collector_id").(int),
	)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(response.Source.Id))
	return resourceSumologicHttpSourceRead(d, meta)
}

func resourceSumologicHttpSourceRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	id, _ := strconv.Atoi(d.Id())
	response, _ := c.GetHttpSource(d.Get("collector_id").(int), id)

	d.Set("name", response.Source.Name)
	d.Set("message_per_request", response.Source.MessagePerRequest)

	return nil
}

func resourceSumologicHttpSourceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

