package slack

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName string = "Slack"
	apiUrl  string = "https://slack.com/api"
)

type Service struct {
	apiToken    string
	httpService *go_http.Service
}

type ServiceConfig struct {
	ApiToken string
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ApiToken == "" {
		return nil, errortools.ErrorMessage("Service ApiToken not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		apiToken:    serviceConfig.ApiToken,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication header
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("Bearer %s", service.apiToken))
	(*requestConfig).NonDefaultHeaders = &header

	req, res, e := service.httpService.HttpRequest(requestConfig)
	if res.StatusCode == http.StatusTooManyRequests {
		retryAfterSeconds := 60
		retryAfter := res.Header.Get("Retry-After")
		if retryAfter == "" {
			fmt.Println("Retry-After header does not exist")
		} else {
			retryAfterSeconds_, err := strconv.Atoi(retryAfter)
			if err != nil {
				fmt.Printf("Retry-After '%s' is not a valid integer\n", retryAfter)
			} else {
				retryAfterSeconds = retryAfterSeconds_
			}
		}

		fmt.Printf("Sleeping %d seconds\n", retryAfterSeconds)
		time.Sleep(time.Duration(retryAfterSeconds) * time.Second)

		return service.httpRequest(requestConfig)
	}

	return req, res, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.apiToken
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
