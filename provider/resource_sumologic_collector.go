package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strconv"
	"os"
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
	accessId := os.Getenv("SL_ACCESSID")
	accessKey := os.Getenv("SL_ACCESSKEY")
	url := fmt.Sprintf("https://%s:%s@api.eu.sumologic.com/api/v1/collectors/%s", accessId, accessKey, d.Id())

	response, err := http.Get(url)

	if err != nil {
		return err
	}

	body, _ := ioutil.ReadAll(response.Body)
	var resp CollectorRequest
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return err
	}

	d.Set("name", resp.Collector.Name)
	d.Set("description", resp.Collector.Description)
	d.Set("category", resp.Collector.Category)

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
	c := meta.(*SumologicClient)

	req := &CollectorRequest{
		Collector: CollectorDetails{
			CollectorType: "Hosted",
			Name: d.Get("name").(string),
			Description: d.Get("description").(string),
			Category: d.Get("category").(string),
		},
	}

	body, err := c.Post(req)

	if err != nil {
		return err
	}

	var resp CollectorRequest
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(resp.Collector.Id))

	fmt.Print(resp.Collector.Id)

	return resourceSumologicCollectorRead(d, meta)
}
type CollectorDetails struct {
	Id int `json:"ID"`
	CollectorType string `json:"collectorType"`
	Name string `json:"name"`
	Description string `json:"description"`
	Category string `json:"category"`
}

type CollectorRequest struct {
	Collector CollectorDetails `json:"collector"`
}
