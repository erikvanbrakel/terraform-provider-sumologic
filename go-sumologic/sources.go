package sumologic

import (
	"fmt"
	"encoding/json"
)

type HttpSource struct {
	Source struct {
		Type string `json:"sourceType"`
		Id int `json:"id,omitempty"`
		Name string `json:"name"`
		MessagePerRequest bool `json:"messagePerRequest"`
	} `json:"source"`
}

func (s *SumologicClient) CreateHttpSource(name string, messagePerRequest bool, collectorId int) (*HttpSource, error) {
	request := HttpSource {}

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

func (s *SumologicClient) CreatePollingSource(name, content_type string, scan_interval int, paused bool, collectorId int) (*PollingSource, error) {

	request := CreatePollingSource{
		ApiVersion: "v1",
		Source: PollingSource{
			SourceType: "Polling",
			Name: name,
			ContentType: content_type,
			ScanInterval: scan_interval,
			Paused: false,
		},
	}

	urlPath := fmt.Sprintf("collectors/%d/sources", collectorId)

	body, err := s.Post(urlPath, request)

	if err != nil {
		return nil, err
	}

	var response CreatePollingSource
	err = json.Unmarshal(body, &response)

	return &response.Source, nil
}

type CreatePollingSource struct {
	ApiVersion string `json:"api.version"`
	Source PollingSource `json:"source"`
}

type GetPollingSource struct {
	Source PollingSource `json:"source"`
}

type PollingSource struct {
	SourceType string `json:"sourceType"`
	Name string `json:"name"`
	ContentType string `json:"contentType"`
	ScanInterval int `json:"scanInterval"`
	Paused bool `json:"paused"`
	ThirdPartyRef []PollingThirdPartyRef `json:"thirdPartyRef,omitempty"`
}

type PollingThirdPartyRef struct {
	Authentication PollingAuthentication
	Path PollingPath
}

type PollingAuthentication struct {
	Type string
	AwsId string
	AwsKey string
}

type PollingPath struct {
	Type string
	BucketName string
	PathExpression string
}