package schema

type TbOwner struct {
	IdOwner            int    `json:"id_owner"`
	IdItem             int64  `json:"id_item"`
	Login              string `json:"login"`
	Id                 int    `json:"id"`
	NodeID             string `json:"node_id"`
	AvatarURL          string `json:"avatar_url"`
	GravatarID         string `json:"gravatar_id"`
	URL                string `json:"url"`
	HTMLURL            string `json:"html_url"`
	FollowersURL       string `json:"followers_url"`
	FollowingURL       string `json:"following_url"`
	GistsURL           string `json:"gists_url"`
	StarredURL         string `json:"starred_url"`
	SubscriptionsURL   string `json:"subscriptions_url"`
	OrganizationsURL   string `json:"organizations_url"`
	ReposURL           string `json:"repos_url"`
	EventsURL          string `json:"events_url"`
	ReceivedEventsURL  string `json:"received_events_url"`
	Type               string `json:"type"`
	SiteAdmin          bool   `json:"site_admin"`
}
