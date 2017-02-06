package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"errors"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider {
		Schema: map[string]*schema.Schema {
			"access_id": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
			},
			"access_key": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",

			},
			"environment": {
				Type: schema.TypeString,
				Optional: true,
				Default: "eu",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sumologic_collector" : resourceSumologicCollector(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := SumologicClient {
		AccessId: d.Get("access_id").(string),
		AccessKey: d.Get("access_key").(string),
		Environment: d.Get("environment").(string),
	}

	if config.AccessId == "" {
		config.AccessId = os.Getenv("SL_ACCESSID")
	}

	if config.AccessKey == "" {
		config.AccessKey = os.Getenv("SL_ACCESSKEY")
	}

	config.BaseUrl = endpoints[config.Environment]
	return &config, nil
}

type SumologicClient struct {
	AccessId string
	AccessKey string
	Environment string
	BaseUrl string
}

var endpoints map[string]string = map[string]string {
	"us1" : "https://api.sumologic.com/api/v1/collectors",
	"us2" : "https://api.us2.sumologic.com/api/v1/collectors",
	"eu" : "https://api.eu.sumologic.com/api/v1/collectors",
	"au" : "https://api.au.sumologic.com/api/v1/collectors",
}

func(s *SumologicClient) Post(payload interface{}) ([]byte, error) {
	url, _ := url.Parse(s.BaseUrl)

	body, _ := json.Marshal(payload)
	req,_ := http.NewRequest (http.MethodPost, url.String(), bytes.NewBuffer(body))
	req.Header.Add("Content-Type","application/json")
	req.SetBasicAuth(s.AccessId, s.AccessKey)

	resp, _ := http.DefaultClient.Do(req)

	d,_ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return nil, errors.New(string(body))
	}

	return d, nil
}
