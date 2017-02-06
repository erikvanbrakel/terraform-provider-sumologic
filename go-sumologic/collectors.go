package sumologic

import (
	"encoding/json"
	"fmt"
)

func(s *SumologicClient) GetCollector(id int) (*Collector,error) {
	data, err := s.Get(fmt.Sprintf("collectors/%d", id))

	if err != nil {
		return nil, err
	}

	var response CollectorResponse
	err = json.Unmarshal(data, &response)


	return &response.Collector, nil
}

func(s *SumologicClient) DeleteCollector(id int) error {
	_, err := s.Delete(fmt.Sprintf("collectors/%d", id))

	return err
}

func (s *SumologicClient) CreateCollector(collectorType, name, description, category string) (*Collector, error) {

	request := CollectorRequest{
		Collector: Collector{
			CollectorType: collectorType,
			Name:name,
			Description: description,
			Category: category,
		},
	}

	var response CollectorResponse
	responseBody, _ := s.Post("collectors", request)

	err := json.Unmarshal(responseBody, &response)

	if err != nil {
		return nil, err
	}


	return &response.Collector, nil
}


type CollectorRequest struct {
	Collector Collector `json:"collector"`
}

type CollectorResponse struct {
	Collector Collector `json:"collector"`
}

type Collector struct {
	Id int `json:"ID"`
	CollectorType string `json:"collectorType"`
	Name string `json:"name"`
	Description string `json:"description"`
	Category string `json:"category"`
}
