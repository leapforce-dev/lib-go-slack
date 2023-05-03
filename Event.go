package slack

type EventWrapper struct {
	Type           string `json:"type"`
	Token          string `json:"token"`
	Challenge      string `json:"challenge"`
	TeamId         string `json:"team_id"`
	ApiAppId       string `json:"api_app_id"`
	Event          Event  `json:"event"`
	EventContext   string `json:"event_context"`
	EventId        string `json:"event_id"`
	EventTime      int    `json:"event_time"`
	Authorizations []struct {
		EnterpriseId        string `json:"enterprise_id"`
		TeamId              string `json:"team_id"`
		UserId              string `json:"user_id"`
		IsBot               bool   `json:"is_bot"`
		IsEnterpriseInstall bool   `json:"is_enterprise_install"`
	} `json:"authorizations"`
	IsExtSharedChannel  bool   `json:"is_ext_shared_channel"`
	ContextTeamId       string `json:"context_team_id"`
	ContextEnterpriseId string `json:"context_enterprise_id"`
}

type Event struct {
	ClientMsgId string `json:"client_msg_id"`
	Type        string `json:"type"`
	Text        string `json:"text"`
	User        string `json:"user"`
	Item        struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
		Ts      string `json:"ts"`
	} `json:"item"`
	Reaction string `json:"reaction"`
	ItemUser string `json:"item_user"`
	Ts       string `json:"ts"`
	Blocks   []struct {
		Type     string `json:"type"`
		BlockId  string `json:"block_id"`
		Elements []struct {
			Type     string `json:"type"`
			Elements []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"elements"`
		} `json:"elements"`
	} `json:"blocks"`
	Team        string `json:"team"`
	Channel     string `json:"channel"`
	EventTs     string `json:"event_ts"`
	ChannelType string `json:"channel_type"`
}
