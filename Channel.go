package slack

import (
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type ChannelsResponse struct {
	Ok               bool      `json:"ok"`
	Channels         []Channel `json:"channels"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

type Channel struct {
	Id                      string          `json:"id"`
	Name                    string          `json:"name"`
	IsChannel               bool            `json:"is_channel"`
	IsGroup                 bool            `json:"is_group"`
	IsIm                    bool            `json:"is_im"`
	IsMpim                  bool            `json:"is_mpim"`
	IsPrivate               bool            `json:"is_private"`
	Created                 int64           `json:"created"`
	IsArchived              bool            `json:"is_archived"`
	IsGeneral               bool            `json:"is_general"`
	Unlinked                int64           `json:"unlinked"`
	NameNormalized          string          `json:"name_normalized"`
	IsShared                bool            `json:"is_shared"`
	IsOrgShared             bool            `json:"is_org_shared"`
	IsPendingExtShared      bool            `json:"is_pending_ext_shared"`
	PendingShared           json.RawMessage `json:"pending_shared"`
	ContextTeamId           string          `json:"context_team_id"`
	Updated                 int64           `json:"updated"`
	ParentConversation      string          `json:"parent_conversation"`
	Creator                 string          `json:"creator"`
	IsExtShared             bool            `json:"is_ext_shared"`
	SharedTeamIds           []string        `json:"shared_team_ids"`
	PendingConnectedTeamIds []string        `json:"pending_connected_team_ids"`
	IsMember                bool            `json:"is_member"`
	Topic                   struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int64  `json:"last_set"`
	} `json:"topic"`
	Purpose struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int64  `json:"last_set"`
	} `json:"purpose"`
	PreviousNames []string `json:"previous_names"`
	NumMembers    int64    `json:"num_members"`
}

func (service *Service) GetChannels() (*[]Channel, *errortools.Error) {
	var channels []Channel
	var cursor = ""

	for {
		var channelsResponse ChannelsResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("conversations.list?cursor=%s", cursor)),
			ResponseModel: &channelsResponse,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		channels = append(channels, channelsResponse.Channels...)
		cursor = channelsResponse.ResponseMetadata.NextCursor

		if cursor == "" {
			break
		}

	}

	return &channels, nil
}
