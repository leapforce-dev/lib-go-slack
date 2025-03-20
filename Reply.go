package slack

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type RepliesResponse struct {
	Ok               bool      `json:"ok"`
	Messages         []Message `json:"messages"`
	HasMore          bool      `json:"has_more"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

type GetRepliesConfig struct {
	ChannelId string
	Ts        string
	Oldest    *string
}

func (service *Service) GetReplies(config *GetRepliesConfig) (*[]Message, *errortools.Error) {
	if config == nil {
		return nil, nil
	}

	var replies []Message
	var cursor = ""

	values := url.Values{}
	values.Set("channel", config.ChannelId)
	values.Set("ts", config.Ts)

	for {
		if cursor != "" {
			values.Set("cursor", cursor)
		}
		if config.Oldest != nil {
			values.Set("oldest", *config.Oldest)
		}

		var repliesResponse RepliesResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodPost,
			Url:           service.url(fmt.Sprintf("conversations.replies?%s", values.Encode())),
			ResponseModel: &repliesResponse,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		replies = append(replies, repliesResponse.Messages...)
		cursor = repliesResponse.ResponseMetadata.NextCursor

		if cursor == "" {
			break
		}

	}

	return &replies, nil
}
