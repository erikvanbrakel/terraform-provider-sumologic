package sumologic

import (
	"encoding/json"
	"fmt"
)

type HttpSource struct {
	Source struct {
		Type              string `json:"sourceType"`
		Id                int    `json:"id,omitempty"`
		Name              string `json:"name"`
		MessagePerRequest bool   `json:"messagePerRequest"`
		Url               string `json:"url"`
	} `json:"source"`
}

func (s *SumologicClient) CreateHttpSource(name string, messagePerRequest bool, collectorId int) (*HttpSource, error) {
	request := HttpSource{}

	request.Source.Type = "HTTP"
	request.Source.Name = name
	request.Source.MessagePerRequest = messagePerRequest

	urlPath := fmt.Sprintf("collectors/%d/sources", collectorId)
	body, err := s.Post(urlPath, request)

	if err != nil {
		return nil, err
	}
	var response HttpSource
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *SumologicClient) GetHttpSource(collectorId, sourceId int) (*HttpSource, error) {

	urlPath := fmt.Sprintf("collectors/%d/sources/%d", collectorId, sourceId)
	body, err := s.Get(urlPath)

	if err != nil {
		return nil, err
	}
	var response HttpSource
	err = json.Unmarshal(body, &response)

	return &response, nil
}

func (s *SumologicClient) CreatePollingSource(name, content_type, category string, scan_interval int, paused bool, collectorId int, auth PollingAuthentication, path PollingPath) (int, error) {

	request := CreatePollingSource{
		ApiVersion: "v1",
		Source: PollingSource{
			SourceType:   "Polling",
			Name:         name,
			Category:     category,
			ContentType:  content_type,
			ScanInterval: scan_interval,
			Paused:       false,
			ThirdPartyRef: PollingThirdPartyRef{
				Resources: []PollingResource{
					{ServiceType: content_type, Authentication: auth, Path: path},
				},
			},
		},
	}

	urlPath := fmt.Sprintf("collectors/%d/sources", collectorId)

	body, err := s.Post(urlPath, request)

	if err != nil {
		return -1, err
	}

	var response CreatePollingSource
	err = json.Unmarshal(body, &response)

	return response.Source.Id, nil
}

func (s *SumologicClient) GetPollingSource(collectorId, sourceId int) (*PollingSource, error) {
	urlPath := fmt.Sprintf("collectors/%d/sources/%d", collectorId, sourceId)
	body, err := s.Get(urlPath)

	if err != nil {
		return nil, err
	}
	var response GetPollingSource
	err = json.Unmarshal(body, &response)

	return &response.Source, nil
}

type CreatePollingSource struct {
	ApiVersion string        `json:"api.version"`
	Source     PollingSource `json:"source"`
}

type GetPollingSource struct {
	Source PollingSource `json:"source"`
}

type PollingSource struct {
	Id            int                  `json:"id,omitempty"`
	SourceType    string               `json:"sourceType"`
	Name          string               `json:"name"`
	ContentType   string               `json:"contentType"`
	ScanInterval  int                  `json:"scanInterval"`
	Paused        bool                 `json:"paused"`
	Category      string               `json:"category"`
	ThirdPartyRef PollingThirdPartyRef `json:"thirdPartyRef,omitempty"`
}

type PollingThirdPartyRef struct {
	Resources []PollingResource `json:"resources"`
}

type PollingResource struct {
	ServiceType    string                `json:"serviceType"`
	Authentication PollingAuthentication `json:"authentication"`
	Path           PollingPath           `json:"path"`
}

type PollingAuthentication struct {
	Type   string `json:"type"`
	AwsId  string `json:"awsId"`
	AwsKey string `json:"awsKey"`
}

type PollingPath struct {
	Type           string `json:"type"`
	BucketName     string `json:"bucketName"`
	PathExpression string `json:"pathExpression"`
}
