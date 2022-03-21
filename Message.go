package slack

import (
	"net/http"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type MessageWrite struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type MessageRead struct {
	Ok        bool   `json:"ok"`
	Channel   string `json:"channel"`
	Timestamp string `json:"ts"`
	Message   struct {
		BotId      string `json:"bot_id"`
		Type       string `json:"type"`
		Text       string `json:"text"`
		User       string `json:"user"`
		Timestamp  string `json:"ts"`
		Team       string `json:"team"`
		BotProfile struct {
			Id    string `json:"id"`
			AppId string `json:"app_id"`
			Name  string `json:"name"`
			Icons struct {
				Image36 string `json:"image_36"`
				Image48 string `json:"image_48"`
				Image72 string `json:"image_72"`
			} `json:"icons"`
			Deleted bool   `json:"deleted"`
			Updated int64  `json:"updated"`
			TeamId  string `json:"team_id"`
		} `json:"bot_profile"`
	} `json:"message"`
	Warning          string `json:"warning"`
	ResponseMetadata struct {
		Warnings []string `json:"warnings"`
	} `json:"response_metadata"`
}

// GetEmployees returns all employees
//
func (service *Service) WriteMessage(channelId string, message string) (*MessageRead, *errortools.Error) {
	messageWrite := MessageWrite{
		Channel: channelId,
		Text:    message,
	}

	messageRead := MessageRead{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("chat.postMessage"),
		BodyModel:     messageWrite,
		ResponseModel: &messageRead,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}
	if !messageRead.Ok {
		return nil, errortools.ErrorMessagef("response returned ok = false, warning: %s, warnings: %s", messageRead.Warning, strings.Join(messageRead.ResponseMetadata.Warnings, ","))
	}

	return &messageRead, nil
}
