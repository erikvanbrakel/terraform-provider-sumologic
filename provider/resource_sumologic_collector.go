package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"

	sumo "github.com/erikvanbrakel/terraform-provider-sumologic/go-sumologic"
)

func resourceSumologicCollector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicCollectorCreate,
		Read:   resourceSumologicCollectorRead,
		Delete: resourceSumologicCollectorDelete,
		Update: resourceSumologicCollectorUpdate,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "",
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "",
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "UTC",
			},
		},
	}
}

func resourceSumologicCollectorRead(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*sumo.SumologicClient)

	id, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	collector, err := c.GetCollector(id)

	if err != nil {
		return err
	}

	if collector == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", collector.Name)
	d.Set("description", collector.Description)
	d.Set("category", collector.Category)
	d.Set("timezone", collector.TimeZone)

	return nil
}

func resourceSumologicCollectorDelete(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*sumo.SumologicClient)

	id, _ := strconv.Atoi(d.Id())
	return c.DeleteCollector(id)
}

func resourceSumologicCollectorCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumo.SumologicClient)

	id, err := c.CreateCollector(sumo.Collector{
		CollectorType: "Hosted",
		Name:          d.Get("name").(string),
	})

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(id))

	return resourceSumologicCollectorUpdate(d, meta)
}

func resourceSumologicCollectorUpdate(d *schema.ResourceData, meta interface{}) error {

	collector := resourceToCollector(d)

	c := meta.(*sumo.SumologicClient)
	err := c.UpdateCollector(collector)

	if err != nil {
		return err
	}

	return resourceSumologicCollectorRead(d, meta)
}

func resourceToCollector(d *schema.ResourceData) sumo.Collector {
	id, _ := strconv.Atoi(d.Id())

	return sumo.Collector{
		Id:            id,
		CollectorType: "Hosted",
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Category:      d.Get("category").(string),
		TimeZone:      d.Get("timezone").(string),
	}
}
