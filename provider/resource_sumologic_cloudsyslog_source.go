package provider

import (
	sumologic "github.com/erikvanbrakel/terraform-provider-sumologic/go-sumologic"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

func resourceSumologicCloudsyslogSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicCloudsyslogSourceCreate,
		Read:   resourceSumologicCloudsyslogSourceRead,
		Update: resourceSumologicCloudsyslogSourceUpdate,
		Delete: resourceSumologicCloudsyslogSourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"collector_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "",
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSumologicCloudsyslogSourceCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	source := sumologic.CloudsyslogSource{}

	source.Name = d.Get("name").(string)

	id, err := c.CreateCloudsyslogSource(
		d.Get("name").(string),
		d.Get("collector_id").(int),
	)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(id))
	return resourceSumologicCloudsyslogSourceUpdate(d, meta)
}

func resourceSumologicCloudsyslogSourceUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	source := resourceToCloudsyslogSource(d)

	err := c.UpdateCloudsyslogSource(source, d.Get("collector_id").(int))

	if err != nil {
		return err
	}

	return resourceSumologicCloudsyslogSourceRead(d, meta)
}

func resourceToCloudsyslogSource(d *schema.ResourceData) sumologic.CloudsyslogSource {

	id, _ := strconv.Atoi(d.Id())

	source := sumologic.CloudsyslogSource{}
	source.Id = id
	source.Type = "Cloudsyslog"
	source.Category = d.Get("category").(string)
	source.Name = d.Get("name").(string)

	return source
}

func resourceSumologicCloudsyslogSourceRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	id, _ := strconv.Atoi(d.Id())
	source, err := c.GetCloudsyslogSource(d.Get("collector_id").(int), id)

	if err != nil {
		return err
	}

	d.Set("name", source.Name)
	d.Set("token", source.Token)

	return nil
}

func resourceSumologicCloudsyslogSourceDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	id, _ := strconv.Atoi(d.Id())
	collector_id, _ := d.Get("collector_id").(int)

	return c.DestroySource(id, collector_id)
}
