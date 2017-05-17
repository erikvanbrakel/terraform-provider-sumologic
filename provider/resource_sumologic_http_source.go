package provider

import (
	sumologic "github.com/erikvanbrakel/terraform-provider-sumologic/go-sumologic"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

func resourceSumologicHttpSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicHttpSourceCreate,
		Read:   resourceSumologicHttpSourceRead,
		Update: resourceSumologicHttpSourceUpdate,
		Delete: resourceSumologicHttpSourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"messagePerRequest": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default: false,
			},
			"collector_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"category" : {
				Type:	  schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default: "",
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSumologicHttpSourceCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	source := sumologic.HttpSource{}

	source.Name = d.Get("name").(string)

	id, err := c.CreateHttpSource(
		d.Get("name").(string),
		d.Get("collector_id").(int),
	)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(id))
	return resourceSumologicHttpSourceUpdate(d, meta)
}

func resourceSumologicHttpSourceUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	source := resourceToHttpSource(d)

	err := c.UpdateHttpSource(source, d.Get("collector_id").(int))

	if err != nil {
		return err
	}

	return resourceSumologicHttpSourceRead(d, meta)
}

func resourceToHttpSource(d *schema.ResourceData) sumologic.HttpSource {

	id, _ := strconv.Atoi(d.Id())

	source := sumologic.HttpSource{}
	source.Id = id
	source.Type = "HTTP"
	source.Category = d.Get("category").(string)
	source.MessagePerRequest = d.Get("messagePerRequest").(bool)
	source.Name = d.Get("name").(string)

	return source
}

func resourceSumologicHttpSourceRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	id, _ := strconv.Atoi(d.Id())
	source, err := c.GetHttpSource(d.Get("collector_id").(int), id)

	if err != nil {
		return err
	}

	d.Set("name", source.Name)
	d.Set("message_per_request", source.MessagePerRequest)
	d.Set("url", source.Url)

	return nil
}

func resourceSumologicHttpSourceDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	id, _ := strconv.Atoi(d.Id())
	collector_id, _ := d.Get("collector_id").(int)

	return c.DestroySource(id, collector_id)
}
