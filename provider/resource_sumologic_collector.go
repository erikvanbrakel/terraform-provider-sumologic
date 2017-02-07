package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
	"fmt"
	"strconv"
	"os"

	sumo "github.com/erikvanbrakel/terraform-provider-sumologic/go-sumologic"
)

func resourceSumologicCollector() *schema.Resource {
	return &schema.Resource {
		Create: resourceSumologicCollectorCreate,
		Read: resourceSumologicCollectorRead,
		Delete: resourceSumologicCollectorDelete,

		Schema: map[string]*schema.Schema {
			"name" : {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"category" : {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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

	d.Set("name", collector.Name)
	d.Set("description", collector.Description)
	d.Set("category", collector.Category)

	return nil
}

func resourceSumologicCollectorDelete(d *schema.ResourceData, meta interface{}) error {

	accessId := os.Getenv("SL_ACCESSID")
	accessKey := os.Getenv("SL_ACCESSKEY")
	url := fmt.Sprintf("https://%s:%s@api.eu.sumologic.com/api/v1/collectors/%s", accessId, accessKey, d.Id())

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	return nil
}

func resourceSumologicCollectorCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumo.SumologicClient)

	collector, err := c.CreateCollector(
		"Hosted",
		d.Get("name").(string),
		d.Get("description").(string),
		d.Get("category").(string),
	)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(collector.Id))

	return resourceSumologicCollectorRead(d, meta)
}
