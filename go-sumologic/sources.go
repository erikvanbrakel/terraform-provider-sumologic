package sumologic

import (
	"encoding/json"
	"fmt"
)

// Common for all sources
type Source struct {
	Id                         int      `json:"id,omitempty"`
	Type                       string   `json:"sourceType"`
	Name                       string   `json:"name"`
	Description                string   `json:"description,omitempty"`
	Category                   string   `json:"category,omitempty"`
	HostName                   string   `json:"hostName,omitempty"`
	TimeZone                   string   `json:"timeZone,omitempty"`
	AutomaticDateParsing       bool     `json:"automaticDateParsing,omitempty"`
	MultilineProcessingEnabled bool     `json:"multilineProcessingEnabled,omitempty"`
	UseAutolineMatching        bool     `json:"useAutolineMatching,omitempty"`
	ManualPrefixRegexp         string   `json:"manualPrefixRegexp,omitempty"`
	ForceTimeZone              bool     `json:"forceTimeZone,omitempty"`
	DefaultDateFormat          string   `json:"defaultDateFormat,omitempty"`
	Filters                    []string `json:"filters,omitempty"`
	CutoffTimestamp            int64    `json:"cutoffTimestamp,omitempty"`
	CutoffRelativeTime         string   `json:"cutoffRelativeTime,omitempty"`
}

func (s *SumologicClient) DestroySource(sourceId int, collectorId int) error {

	_, err := s.Delete(fmt.Sprintf("collectors/%d/sources/%d", collectorId, sourceId))

	return err
}

// HTTP source specific
type HttpSource struct {
	Source
	MessagePerRequest bool   `json:"messagePerRequest"`
	Url               string `json:"url,omitempty"`
}

func (s *SumologicClient) CreateHttpSource(name string, collectorId int) (int, error) {

	type HttpSourceMessage struct {
		Source HttpSource `json:"source"`
	}

	request := HttpSourceMessage{}

	source := HttpSource{}

	source.Type = "HTTP"
	source.Name = name

	request.Source = source

	urlPath := fmt.Sprintf("collectors/%d/sources", collectorId)
	body, err := s.Post(urlPath, request)

	if err != nil {
		return -1, err
	}

	var response HttpSourceMessage
	err = json.Unmarshal(body, &response)

	if err != nil {
		return -1, err
	}

	newSource := response.Source

	return newSource.Id, nil
}

func (s *SumologicClient) GetHttpSource(collectorId, sourceId int) (*HttpSource, error) {

	urlPath := fmt.Sprintf("collectors/%d/sources/%d", collectorId, sourceId)
	body, _, err := s.Get(urlPath)

	if err != nil {
		return nil, err
	}

	type Response struct {
		Source HttpSource `json:"source"`
	}

	var response Response
	err = json.Unmarshal(body, &response)

	var source = response.Source

	return &source, nil
}

func (s *SumologicClient) UpdateHttpSource(source HttpSource, collectorId int) error {

	url := fmt.Sprintf("collectors/%d/sources/%d", collectorId, source.Id)

	type HttpSourceMessage struct {
		Source HttpSource `json:"source"`
	}

	request := HttpSourceMessage{
		Source: source,
	}

	_, err := s.Put(url, request)

	if err != nil {
		return err
	}

	return nil

}

// Polling source specific
type PollingSource struct {
	Source
	ContentType   string               `json:"contentType"`
	ScanInterval  int                  `json:"scanInterval"`
	Paused        bool                 `json:"paused"`
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

func (s *SumologicClient) CreatePollingSource(name, content_type, category string, scan_interval int, paused bool, collectorId int, auth PollingAuthentication, path PollingPath) (int, error) {

	type PollingSourceMessage struct {
		Source PollingSource `json:"source"`
	}

	request := PollingSourceMessage{}

	request.Source.Type = "Polling"
	request.Source.Name = name
	request.Source.Category = category
	request.Source.ContentType = content_type
	request.Source.ScanInterval = scan_interval
	request.Source.Paused = false
	request.Source.ThirdPartyRef = PollingThirdPartyRef{
		Resources: []PollingResource{
			{ServiceType: content_type, Authentication: auth, Path: path},
		},
	}

	urlPath := fmt.Sprintf("collectors/%d/sources", collectorId)

	body, err := s.Post(urlPath, request)

	if err != nil {
		return -1, err
	}

	var response PollingSourceMessage
	err = json.Unmarshal(body, &response)

	if err != nil {
		return -1, err
	}

	return response.Source.Id, nil
}

func (s *SumologicClient) GetPollingSource(collectorId, sourceId int) (*PollingSource, error) {
	urlPath := fmt.Sprintf("collectors/%d/sources/%d", collectorId, sourceId)
	body, _, err := s.Get(urlPath)

	if err != nil {
		return nil, err
	}

	type PollingSourceResponse struct {
		Source PollingSource `json:"source"`
	}

	var response PollingSourceResponse
	err = json.Unmarshal(body, &response)

	return &response.Source, nil
}

func (s *SumologicClient) UpdatePollingSource(source PollingSource, collectorId int) error {
	url := fmt.Sprintf("collectors/%d/sources/%d", collectorId, source.Id)

	type PollingSourceMessage struct {
		Source PollingSource `json:"source"`
	}

	request := PollingSourceMessage {
		Source: source,
	}

	_, err := s.Put(url, request)

	if err != nil {
		return err
	}

	return nil
}
type CloudsyslogSource struct {
	Source
	Token string `json:"token,omitempty"`
}

func (s *SumologicClient) CreateCloudsyslogSource(name string, collectorId int) (int, error) {

	type CloudsyslogSourceMessage struct {
		Source CloudsyslogSource `json:"source"`
	}

	request := CloudsyslogSourceMessage{}

	source := CloudsyslogSource{}

	source.Type = "Cloudsyslog"
	source.Name = name

	request.Source = source

	urlPath := fmt.Sprintf("collectors/%d/sources", collectorId)
	body, err := s.Post(urlPath, request)

	if err != nil {
		return -1, err
	}

	var response CloudsyslogSourceMessage
	err = json.Unmarshal(body, &response)

	if err != nil {
		return -1, err
	}

	newSource := response.Source

	return newSource.Id, nil
}

func (s *SumologicClient) GetCloudsyslogSource(collectorId, sourceId int) (*CloudsyslogSource, error) {

	urlPath := fmt.Sprintf("collectors/%d/sources/%d", collectorId, sourceId)
	body, _, err := s.Get(urlPath)

	if err != nil {
		return nil, err
	}

	type Response struct {
		Source CloudsyslogSource `json:"source"`
	}

	var response Response
	err = json.Unmarshal(body, &response)

	var source = response.Source

	return &source, nil
}

func (s *SumologicClient) UpdateCloudsyslogSource(source CloudsyslogSource, collectorId int) error {

	url := fmt.Sprintf("collectors/%d/sources/%d", collectorId, source.Id)

	type CloudsyslogSourceMessage struct {
		Source CloudsyslogSource `json:"source"`
	}

	request := CloudsyslogSourceMessage{
		Source: source,
	}

	_, err := s.Put(url, request)

	if err != nil {
		return err
	}

	return nil

}
