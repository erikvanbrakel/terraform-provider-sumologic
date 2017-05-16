package sumologic

import (
	"encoding/json"
	"fmt"
)

func (s *SumologicClient) GetCollector(id int) (*Collector, error) {
	data, _, err := s.Get(fmt.Sprintf("collectors/%d", id))

	if err != nil {
		return nil, err
	}

	var response CollectorResponse
	err = json.Unmarshal(data, &response)

	return &response.Collector, nil
}

func (s *SumologicClient) DeleteCollector(id int) error {
	_, err := s.Delete(fmt.Sprintf("collectors/%d", id))

	return err
}

func (s *SumologicClient) CreateCollector(collector Collector) (int, error) {

	request := CollectorRequest{
		Collector: collector,
	}

	var response CollectorResponse

	responseBody, err := s.Post("collectors", request)

	if err != nil {
		return -1, err
	}

	err = json.Unmarshal(responseBody, &response)

	if err != nil {
		return -1, err
	}

	return response.Collector.Id, nil
}

func (s *SumologicClient) UpdateCollector(collector Collector) error {
	url := fmt.Sprintf("collectors/%d", collector.Id)

	request := CollectorRequest{
		Collector: collector,
	}

	_, err := s.Put(url, request)

	if err != nil {
		return err
	}

	return nil
}

type CollectorRequest struct {
	Collector Collector `json:"collector"`
}

type CollectorResponse struct {
	Collector Collector `json:"collector"`
}

type Collector struct {
	Id            int    `json:"id,omitempty"`
	CollectorType string `json:"collectorType,omitempty"`
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	Category      string `json:"category,omitempty"`
	TimeZone      string `json:"timeZone,omitempty"`
}
