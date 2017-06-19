package provider

import (
	"github.com/erikvanbrakel/terraform-provider-sumologic/go-sumologic"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
	"log"
)

func resourceSumologicPollingSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicPollingSourceCreate,
		Read:   resourceSumologicPollingSourceRead,
		Update: resourceSumologicPollingSourceUpdate,
		Delete: resourceSumologicPollingSourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"category": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scan_interval": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},
			"paused": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: false,
			},
			"collector_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"authentication": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"secret_key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"path": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"path_expression": {
							Type:     schema.TypeString,
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

	sourceId, err := c.CreatePollingSource(
		d.Get("name").(string),
		d.Get("content_type").(string),
		d.Get("category").(string),
		d.Get("scan_interval").(int),
		d.Get("paused").(bool),
		d.Get("collector_id").(int),
		getAuthentication(d),
		getPathSettings(d),
	)

	if err != nil {
		return err
	}

	id := strconv.Itoa(sourceId)

	d.SetId(id)

	return resourceSumologicPollingSourceRead(d, meta)
}

func resourceSumologicPollingSourceRead(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*sumologic.SumologicClient)

	id, err := strconv.Atoi(d.Id())
	collector_id := d.Get("collector_id").(int)

	source, err := c.GetPollingSource(collector_id, id)

	if err != nil {
		return err
	}

	pollingResources := source.ThirdPartyRef.Resources
	path := getThirdyPartyPathAttributes(pollingResources)

	if err := d.Set("path", path); err != nil {
		return err
	}

	d.Set("name", source.Name)
	d.Set("content_type", source.ContentType)
	d.Set("category", source.Category)
	d.Set("scan_interval", source.ScanInterval)
	d.Set("paused", source.Paused)

	return nil
}

func resourceSumologicPollingSourceDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	id, _ := strconv.Atoi(d.Id())
	collector_id, _ := d.Get("collector_id").(int)

	return c.DestroySource(id, collector_id)
}

func resourceSumologicPollingSourceUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*sumologic.SumologicClient)

	source := resourceToPollingSource(d)

	err := c.UpdatePollingSource(source, d.Get("collector_id").(int))

	if err != nil {
		return err
	}

	return resourceSumologicPollingSourceRead(d, meta)
}

func resourceToPollingSource(d *schema.ResourceData) sumologic.PollingSource {

	id, _ := strconv.Atoi(d.Id())
	source := sumologic.PollingSource{}
	pollingResource := sumologic.PollingResource{}

	source.Id = id
	source.Type = "Polling"
	source.Category = d.Get("category").(string)
	source.Paused = d.Get("paused").(bool)
	source.Name = d.Get("name").(string)
	source.ScanInterval = d.Get("scan_interval").(int)
	source.ContentType = d.Get("content_type").(string)

	pollingResource.ServiceType    = "AwsS3AuditBucket"
	pollingResource.Authentication = getAuthentication(d)
	pollingResource.Path           = getPathSettings(d)

	source.ThirdPartyRef.Resources = append(source.ThirdPartyRef.Resources, pollingResource)

	return source
}

func getThirdyPartyPathAttributes(pollingResource []sumologic.PollingResource) []map[string]interface{} {

	var s []map[string]interface{}
	for _, t := range pollingResource {
		mapping := map[string]interface{}{
			"bucket_name":        t.Path.BucketName,
			"path_expression":    t.Path.PathExpression,
		}
		s = append(s, mapping)
	}

	return s
}

func getAuthentication(d *schema.ResourceData) sumologic.PollingAuthentication {

	auths := d.Get("authentication").([]interface{})
	authSettings := sumologic.PollingAuthentication{}

	if len(auths) > 0 {
		auth := auths[0].(map[string]interface{})
		authSettings.Type = "S3BucketAuthentication"
		authSettings.AwsId = auth["access_key"].(string)
		authSettings.AwsKey = auth["secret_key"].(string)
	}

	return authSettings
}

func getPathSettings(d *schema.ResourceData) sumologic.PollingPath {
	pathSettings := sumologic.PollingPath{}
	paths := d.Get("path").([]interface{})

	if len(paths) > 0 {
		path := paths[0].(map[string]interface{})
		pathSettings.Type = "S3BucketPathExpression"
		pathSettings.BucketName = path["bucket_name"].(string)
		pathSettings.PathExpression = path["path_expression"].(string)
	}

	return pathSettings
}
