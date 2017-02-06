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