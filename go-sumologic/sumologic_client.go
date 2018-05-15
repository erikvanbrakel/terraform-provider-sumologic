package sumologic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SumologicClient struct {
	AccessId    string
	AccessKey   string
	Environment string
	BaseUrl     *url.URL
}

var endpoints map[string]string = map[string]string{
	"us1": "https://api.sumologic.com/api/v1/",
	"us2": "https://api.us2.sumologic.com/api/v1/",
	"eu":  "https://api.eu.sumologic.com/api/v1/",
	"au":  "https://api.au.sumologic.com/api/v1/",
	"de":  "https://api.de.sumologic.com/api/v1/",
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (s *SumologicClient) Post(urlPath string, payload interface{}) ([]byte, error) {
	relativeUrl, _ := url.Parse(urlPath)
	url := s.BaseUrl.ResolveReference(relativeUrl)

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(s.AccessId, s.AccessKey)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	d, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		var errorResponse ErrorResponse
		_ = json.Unmarshal(d, &errorResponse)
		return nil, errors.New(errorResponse.Message)
	}

	return d, nil
}

func (s *SumologicClient) Put(urlPath string, payload interface{}) ([]byte, error) {

	relativeUrl, _ := url.Parse(urlPath)
	url := s.BaseUrl.ResolveReference(relativeUrl)

	_, etag, _ := s.Get(url.String())

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, url.String(), bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("If-Match", etag)

	req.SetBasicAuth(s.AccessId, s.AccessKey)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	d, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		var errorResponse ErrorResponse
		_ = json.Unmarshal(d, &errorResponse)
		return nil, errors.New(errorResponse.Message)
	}

	return d, nil
}

func (s *SumologicClient) Get(urlPath string) ([]byte, string, error) {
	relativeUrl, _ := url.Parse(urlPath)
	url := s.BaseUrl.ResolveReference(relativeUrl)

	req, _ := http.NewRequest(http.MethodGet, url.String(), nil)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(s.AccessId, s.AccessKey)

	resp, _ := http.DefaultClient.Do(req)

	d, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		var errorResponse ErrorResponse
		_ = json.Unmarshal(d, &errorResponse)
		return nil, "", errors.New(errorResponse.Message)
	}

	return d, resp.Header.Get("ETag"), nil
}

func (s *SumologicClient) Delete(urlPath string) ([]byte, error) {
	relativeUrl, _ := url.Parse(urlPath)
	url := s.BaseUrl.ResolveReference(relativeUrl)

	req, _ := http.NewRequest(http.MethodDelete, url.String(), nil)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(s.AccessId, s.AccessKey)

	resp, _ := http.DefaultClient.Do(req)

	d, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return nil, errors.New(string(d))
	}

	return d, nil
}

func NewClient(accessId, accessKey, environment string) (*SumologicClient, error) {
	client := SumologicClient{
		AccessId:    accessId,
		AccessKey:   accessKey,
		Environment: environment,
	}

	client.BaseUrl, _ = url.Parse(endpoints[client.Environment])
	return &client, nil
}
