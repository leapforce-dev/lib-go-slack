package slack

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type ConversationHistoryResponse struct {
	Ok                  bool        `json:"ok"`
	Messages            []Message   `json:"messages"`
	HasMore             bool        `json:"has_more"`
	PinCount            int         `json:"pin_count"`
	ChannelActionsTs    interface{} `json:"channel_actions_ts"`
	ChannelActionsCount int         `json:"channel_actions_count"`
	ResponseMetadata    struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

type Message struct {
	User        string `json:"user,omitempty"`
	Type        string `json:"type"`
	Ts          string `json:"ts"`
	ClientMsgId string `json:"client_msg_id,omitempty"`
	Text        string `json:"text"`
	Team        string `json:"team,omitempty"`
	Blocks      []struct {
		Type     string `json:"type"`
		BlockId  string `json:"block_id"`
		Elements []struct {
			Type     string `json:"type"`
			Elements []struct {
				Type     string `json:"type"`
				Text     string `json:"text,omitempty"`
				Range    string `json:"range,omitempty"`
				Elements []struct {
					Type    string `json:"type"`
					Text    string `json:"text,omitempty"`
					Name    string `json:"name,omitempty"`
					Unicode string `json:"unicode,omitempty"`
					Style   struct {
						Bold bool `json:"bold"`
					} `json:"style,omitempty"`
					Url    string `json:"url,omitempty"`
					UserId string `json:"user_id,omitempty"`
				} `json:"elements,omitempty"`
				Name    string `json:"name,omitempty"`
				Unicode string `json:"unicode,omitempty"`
				Style   struct {
					Bold   bool `json:"bold,omitempty"`
					Italic bool `json:"italic,omitempty"`
				} `json:"style,omitempty"`
				Url       string `json:"url,omitempty"`
				UserId    string `json:"user_id,omitempty"`
				SkinTone  int    `json:"skin_tone,omitempty"`
				ChannelId string `json:"channel_id,omitempty"`
				Unsafe    bool   `json:"unsafe,omitempty"`
			} `json:"elements"`
			Style  string `json:"style,omitempty"`
			Indent int    `json:"indent,omitempty"`
			Border int    `json:"border,omitempty"`
		} `json:"elements"`
	} `json:"blocks"`
	Reactions []struct {
		Name  string   `json:"name"`
		Users []string `json:"users"`
		Count int      `json:"count"`
	} `json:"reactions,omitempty"`
	ThreadTs        string   `json:"thread_ts,omitempty"`
	ReplyCount      int      `json:"reply_count,omitempty"`
	ReplyUsersCount int      `json:"reply_users_count,omitempty"`
	LatestReply     string   `json:"latest_reply,omitempty"`
	ReplyUsers      []string `json:"reply_users,omitempty"`
	IsLocked        bool     `json:"is_locked,omitempty"`
	Subscribed      bool     `json:"subscribed,omitempty"`
	Edited          struct {
		User string `json:"user"`
		Ts   string `json:"ts"`
	} `json:"edited,omitempty"`
	Attachments []struct {
		ImageUrl      string      `json:"image_url,omitempty"`
		ImageWidth    int         `json:"image_width,omitempty"`
		ImageHeight   int         `json:"image_height,omitempty"`
		ImageBytes    int         `json:"image_bytes,omitempty"`
		FromUrl       string      `json:"from_url,omitempty"`
		ServiceIcon   string      `json:"service_icon,omitempty"`
		Id            int         `json:"id,omitempty"`
		OriginalUrl   string      `json:"original_url,omitempty"`
		Fallback      string      `json:"fallback"`
		ServiceName   string      `json:"service_name,omitempty"`
		ServiceUrl    string      `json:"service_url,omitempty"`
		Ts            interface{} `json:"ts,omitempty"`
		AuthorId      string      `json:"author_id,omitempty"`
		ChannelId     string      `json:"channel_id,omitempty"`
		ChannelTeam   string      `json:"channel_team,omitempty"`
		IsMsgUnfurl   bool        `json:"is_msg_unfurl,omitempty"`
		MessageBlocks []struct {
			Team    string `json:"team"`
			Channel string `json:"channel"`
			Ts      string `json:"ts"`
			Message struct {
				Blocks []struct {
					Type     string `json:"type"`
					BlockId  string `json:"block_id"`
					Elements []struct {
						Type     string `json:"type"`
						Elements []struct {
							Type    string `json:"type"`
							Text    string `json:"text,omitempty"`
							Url     string `json:"url,omitempty"`
							Name    string `json:"name,omitempty"`
							Unicode string `json:"unicode,omitempty"`
						} `json:"elements"`
					} `json:"elements"`
				} `json:"blocks"`
			} `json:"message"`
		} `json:"message_blocks,omitempty"`
		Files []struct {
			Id               string `json:"id"`
			Created          int    `json:"created"`
			Timestamp        int    `json:"timestamp"`
			Name             string `json:"name"`
			Title            string `json:"title"`
			Mimetype         string `json:"mimetype"`
			Filetype         string `json:"filetype"`
			PrettyType       string `json:"pretty_type"`
			User             string `json:"user"`
			UserTeam         string `json:"user_team"`
			Editable         bool   `json:"editable"`
			Size             int    `json:"size"`
			Mode             string `json:"mode"`
			IsExternal       bool   `json:"is_external"`
			ExternalType     string `json:"external_type"`
			IsPublic         bool   `json:"is_public"`
			PublicUrlShared  bool   `json:"public_url_shared"`
			DisplayAsBot     bool   `json:"display_as_bot"`
			Username         string `json:"username"`
			UrlPrivate       string `json:"url_private"`
			MediaDisplayType string `json:"media_display_type"`
			ThumbVideo       string `json:"thumb_video"`
			ThumbVideoW      int    `json:"thumb_video_w"`
			ThumbVideoH      int    `json:"thumb_video_h"`
			Permalink        string `json:"permalink"`
			CommentsCount    int    `json:"comments_count"`
			IsStarred        bool   `json:"is_starred"`
			ExternalId       string `json:"external_id"`
			ExternalUrl      string `json:"external_url"`
			HasRichPreview   bool   `json:"has_rich_preview"`
			FileAccess       string `json:"file_access"`
		} `json:"files,omitempty"`
		Color         string   `json:"color,omitempty"`
		IsShare       bool     `json:"is_share,omitempty"`
		Text          string   `json:"text,omitempty"`
		AuthorName    string   `json:"author_name,omitempty"`
		AuthorLink    string   `json:"author_link,omitempty"`
		AuthorIcon    string   `json:"author_icon,omitempty"`
		AuthorSubname string   `json:"author_subname,omitempty"`
		MrkdwnIn      []string `json:"mrkdwn_in,omitempty"`
		Footer        string   `json:"footer,omitempty"`
		/*Blocks        []struct {
			Type     string `json:"type"`
			BlockId  string `json:"block_id"`
			Elements []struct {
				Type     string `json:"type"`
				ActionId string `json:"action_id"`
				Text     struct {
					Type  string `json:"type"`
					Text  string `json:"text"`
					Emoji bool   `json:"emoji"`
				} `json:"text"`
				Style          string `json:"style"`
				Value          string `json:"value"`
				ThirdPartyAuth struct {
					EnableDynamicAuth bool `json:"enable_dynamic_auth"`
				} `json:"third_party_auth"`
			} `json:"elements"`
		} `json:"blocks,omitempty"`*/
		Title           string `json:"title,omitempty"`
		TitleLink       string `json:"title_link,omitempty"`
		FooterIcon      string `json:"footer_icon,omitempty"`
		BotId           string `json:"bot_id,omitempty"`
		BotTeamId       string `json:"bot_team_id,omitempty"`
		AppUnfurlUrl    string `json:"app_unfurl_url,omitempty"`
		IsAppUnfurl     bool   `json:"is_app_unfurl,omitempty"`
		AppId           string `json:"app_id,omitempty"`
		ThumbUrl        string `json:"thumb_url,omitempty"`
		ThumbWidth      int    `json:"thumb_width,omitempty"`
		ThumbHeight     int    `json:"thumb_height,omitempty"`
		VideoHtml       string `json:"video_html,omitempty"`
		VideoHtmlWidth  int    `json:"video_html_width,omitempty"`
		VideoHtmlHeight int    `json:"video_html_height,omitempty"`
	} `json:"attachments,omitempty"`
	Files []struct {
		Id                 string `json:"id"`
		Created            int    `json:"created"`
		Timestamp          int    `json:"timestamp"`
		Name               string `json:"name"`
		Title              string `json:"title"`
		Mimetype           string `json:"mimetype"`
		Filetype           string `json:"filetype"`
		PrettyType         string `json:"pretty_type"`
		User               string `json:"user"`
		UserTeam           string `json:"user_team"`
		Editable           bool   `json:"editable"`
		Size               int    `json:"size"`
		Mode               string `json:"mode"`
		IsExternal         bool   `json:"is_external"`
		ExternalType       string `json:"external_type"`
		IsPublic           bool   `json:"is_public"`
		PublicUrlShared    bool   `json:"public_url_shared"`
		DisplayAsBot       bool   `json:"display_as_bot"`
		Username           string `json:"username"`
		UrlPrivate         string `json:"url_private"`
		UrlPrivateDownload string `json:"url_private_download,omitempty"`
		MediaDisplayType   string `json:"media_display_type"`
		Thumb64            string `json:"thumb_64,omitempty"`
		Thumb80            string `json:"thumb_80,omitempty"`
		Thumb360           string `json:"thumb_360,omitempty"`
		Thumb360W          int    `json:"thumb_360_w,omitempty"`
		Thumb360H          int    `json:"thumb_360_h,omitempty"`
		Thumb160           string `json:"thumb_160,omitempty"`
		OriginalW          int    `json:"original_w,omitempty"`
		OriginalH          int    `json:"original_h,omitempty"`
		ThumbTiny          string `json:"thumb_tiny,omitempty"`
		Permalink          string `json:"permalink"`
		PermalinkPublic    string `json:"permalink_public,omitempty"`
		IsStarred          bool   `json:"is_starred"`
		HasRichPreview     bool   `json:"has_rich_preview"`
		FileAccess         string `json:"file_access"`
		Thumb480           string `json:"thumb_480,omitempty"`
		Thumb480W          int    `json:"thumb_480_w,omitempty"`
		Thumb480H          int    `json:"thumb_480_h,omitempty"`
		Thumb720           string `json:"thumb_720,omitempty"`
		Thumb720W          int    `json:"thumb_720_w,omitempty"`
		Thumb720H          int    `json:"thumb_720_h,omitempty"`
		Thumb800           string `json:"thumb_800,omitempty"`
		Thumb800W          int    `json:"thumb_800_w,omitempty"`
		Thumb800H          int    `json:"thumb_800_h,omitempty"`
		Thumb960           string `json:"thumb_960,omitempty"`
		Thumb960W          int    `json:"thumb_960_w,omitempty"`
		Thumb960H          int    `json:"thumb_960_h,omitempty"`
		Thumb1024          string `json:"thumb_1024,omitempty"`
		Thumb1024W         int    `json:"thumb_1024_w,omitempty"`
		Thumb1024H         int    `json:"thumb_1024_h,omitempty"`
		ThumbVideo         string `json:"thumb_video,omitempty"`
		ThumbVideoW        int    `json:"thumb_video_w,omitempty"`
		ThumbVideoH        int    `json:"thumb_video_h,omitempty"`
		ExternalId         string `json:"external_id,omitempty"`
		ExternalUrl        string `json:"external_url,omitempty"`
		Subtype            string `json:"subtype,omitempty"`
		Transcription      struct {
			Status  string `json:"status"`
			Locale  string `json:"locale"`
			Preview struct {
				Content string `json:"content"`
				HasMore bool   `json:"has_more"`
			} `json:"preview"`
		} `json:"transcription,omitempty"`
		Mp4        string `json:"mp4,omitempty"`
		Vtt        string `json:"vtt,omitempty"`
		Hls        string `json:"hls,omitempty"`
		HlsEmbed   string `json:"hls_embed,omitempty"`
		Mp4Low     string `json:"mp4_low,omitempty"`
		DurationMs int    `json:"duration_ms,omitempty"`
	} `json:"files,omitempty"`
	Upload       bool     `json:"upload,omitempty"`
	DisplayAsBot bool     `json:"display_as_bot,omitempty"`
	Subtype      string   `json:"subtype,omitempty"`
	Username     string   `json:"username,omitempty"`
	BotId        string   `json:"bot_id,omitempty"`
	AppId        string   `json:"app_id,omitempty"`
	TriggerId    string   `json:"trigger_id,omitempty"`
	XFiles       []string `json:"x_files,omitempty"`
	BotProfile   struct {
		Id      string `json:"id"`
		Deleted bool   `json:"deleted"`
		Name    string `json:"name"`
		Updated int    `json:"updated"`
		AppId   string `json:"app_id"`
		Icons   struct {
			Image36 string `json:"image_36"`
			Image48 string `json:"image_48"`
			Image72 string `json:"image_72"`
		} `json:"icons"`
		TeamId string `json:"team_id"`
	} `json:"bot_profile,omitempty"`
	Root struct {
		Text  string `json:"text"`
		Files []struct {
			Id                 string `json:"id"`
			Created            int    `json:"created"`
			Timestamp          int    `json:"timestamp"`
			Name               string `json:"name"`
			Title              string `json:"title"`
			Mimetype           string `json:"mimetype"`
			Filetype           string `json:"filetype"`
			PrettyType         string `json:"pretty_type"`
			User               string `json:"user"`
			UserTeam           string `json:"user_team"`
			Editable           bool   `json:"editable"`
			Size               int    `json:"size"`
			Mode               string `json:"mode"`
			IsExternal         bool   `json:"is_external"`
			ExternalType       string `json:"external_type"`
			IsPublic           bool   `json:"is_public"`
			PublicUrlShared    bool   `json:"public_url_shared"`
			DisplayAsBot       bool   `json:"display_as_bot"`
			Username           string `json:"username"`
			UrlPrivate         string `json:"url_private"`
			UrlPrivateDownload string `json:"url_private_download"`
			MediaDisplayType   string `json:"media_display_type"`
			Thumb64            string `json:"thumb_64"`
			Thumb80            string `json:"thumb_80"`
			Thumb360           string `json:"thumb_360"`
			Thumb360W          int    `json:"thumb_360_w"`
			Thumb360H          int    `json:"thumb_360_h"`
			Thumb480           string `json:"thumb_480"`
			Thumb480W          int    `json:"thumb_480_w"`
			Thumb480H          int    `json:"thumb_480_h"`
			Thumb160           string `json:"thumb_160"`
			Thumb720           string `json:"thumb_720"`
			Thumb720W          int    `json:"thumb_720_w"`
			Thumb720H          int    `json:"thumb_720_h"`
			Thumb800           string `json:"thumb_800"`
			Thumb800W          int    `json:"thumb_800_w"`
			Thumb800H          int    `json:"thumb_800_h"`
			Thumb960           string `json:"thumb_960"`
			Thumb960W          int    `json:"thumb_960_w"`
			Thumb960H          int    `json:"thumb_960_h"`
			Thumb1024          string `json:"thumb_1024"`
			Thumb1024W         int    `json:"thumb_1024_w"`
			Thumb1024H         int    `json:"thumb_1024_h"`
			OriginalW          int    `json:"original_w"`
			OriginalH          int    `json:"original_h"`
			ThumbTiny          string `json:"thumb_tiny"`
			Permalink          string `json:"permalink"`
			PermalinkPublic    string `json:"permalink_public"`
			IsStarred          bool   `json:"is_starred"`
			HasRichPreview     bool   `json:"has_rich_preview"`
			FileAccess         string `json:"file_access"`
		} `json:"files"`
		Upload       bool   `json:"upload"`
		User         string `json:"user"`
		DisplayAsBot bool   `json:"display_as_bot"`
		Blocks       []struct {
			Type     string `json:"type"`
			BlockId  string `json:"block_id"`
			Elements []struct {
				Type     string `json:"type"`
				Elements []struct {
					Type  string `json:"type"`
					Range string `json:"range,omitempty"`
					Text  string `json:"text,omitempty"`
				} `json:"elements"`
			} `json:"elements"`
		} `json:"blocks"`
		Type            string   `json:"type"`
		Ts              string   `json:"ts"`
		ClientMsgId     string   `json:"client_msg_id"`
		ThreadTs        string   `json:"thread_ts"`
		ReplyCount      int      `json:"reply_count"`
		ReplyUsersCount int      `json:"reply_users_count"`
		LatestReply     string   `json:"latest_reply"`
		ReplyUsers      []string `json:"reply_users"`
		IsLocked        bool     `json:"is_locked"`
		Subscribed      bool     `json:"subscribed"`
	} `json:"root,omitempty"`
}

type GetConversationHistoryConfig struct {
	ChannelId string
	Oldest    *string
}

func (service *Service) GetConversationHistory(config *GetConversationHistoryConfig) (*[]Message, *errortools.Error) {
	if config == nil {
		return nil, nil
	}

	var messages []Message
	var cursor = ""

	values := url.Values{}
	values.Set("channel", config.ChannelId)

	for {
		if cursor != "" {
			values.Set("cursor", cursor)
		}
		if config.Oldest != nil {
			values.Set("oldest", *config.Oldest)
		}

		var conversationHistoryResponse ConversationHistoryResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodPost,
			Url:           service.url(fmt.Sprintf("conversations.history?%s", values.Encode())),
			ResponseModel: &conversationHistoryResponse,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		messages = append(messages, conversationHistoryResponse.Messages...)
		cursor = conversationHistoryResponse.ResponseMetadata.NextCursor

		if cursor == "" {
			break
		}

	}

	return &messages, nil
}
