package slack

import (
	"encoding/json"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type UsersResponse struct {
	Ok               bool   `json:"ok"`
	Offset           string `json:"offset"`
	Timestamp        string `json:"ts"`
	Members          []User `json:"members"`
	CacheTs          int64  `json:"cache_ts"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

type User struct {
	Id       string `json:"id"`
	TeamId   string `json:"team_id"`
	Name     string `json:"name"`
	Deleted  bool   `json:"deleted"`
	Color    string `json:"color"`
	RealName string `json:"real_name"`
	Tz       string `json:"tz"`
	TzLabel  string `json:"tz_label"`
	TzOffset int64  `json:"tz_offset"`
	Profile  struct {
		Title                   string          `json:"title"`
		Phone                   string          `json:"phone"`
		Skype                   string          `json:"skype"`
		RealName                string          `json:"real_name"`
		RealNameNormalized      string          `json:"real_name_normalized"`
		DisplayName             string          `json:"display_name"`
		DisplayNameNormalized   string          `json:"display_name_normalized"`
		Fields                  json.RawMessage `json:"fields"`
		StatusText              string          `json:"status_text"`
		StatusEmoji             string          `json:"status_emoji"`
		StatusEmojiDisplayInfo  json.RawMessage `json:"status_emoji_display_info"`
		StatusExpiration        int64           `json:"status_expiration"`
		AvatarHash              string          `json:"avatar_hash"`
		ImageOriginal           string          `json:"image_original"`
		IsCustomImage           bool            `json:"is_custom_image"`
		Email                   string          `json:"email"`
		HuddleState             string          `json:"huddle_state"`
		HuddleStateExpirationTs int64           `json:"huddle_state_expiration_ts"`
		FirstName               string          `json:"first_name"`
		LastName                string          `json:"last_name"`
		Image24                 string          `json:"image_24"`
		Image32                 string          `json:"image_32"`
		Image48                 string          `json:"image_48"`
		Image72                 string          `json:"image_72"`
		Image192                string          `json:"image_192"`
		Image512                string          `json:"image_512"`
		Image1024               string          `json:"image_1024"`
		StatusTextCanonical     string          `json:"status_text_canonical"`
		Team                    string          `json:"team"`
	} `json:"profile"`
	IsAdmin                bool   `json:"is_admin"`
	IsOwner                bool   `json:"is_owner"`
	IsPrimaryOwner         bool   `json:"is_primary_owner"`
	IsRestricted           bool   `json:"is_restricted"`
	IsUltraRestricted      bool   `json:"is_ultra_restricted"`
	IsBot                  bool   `json:"is_bot"`
	IsAppUser              bool   `json:"is_app_user"`
	Updated                int64  `json:"updated"`
	IsEmailConfirmed       bool   `json:"is_email_confirmed"`
	WhoCanShareContactCard string `json:"who_can_share_contact_card"`
	Locale                 string `json:"locale"`
}

func (service *Service) GetUsers() (*[]User, *errortools.Error) {
	var usersResponse UsersResponse

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("users.list"),
		ResponseModel: &usersResponse,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &usersResponse.Members, nil
}
