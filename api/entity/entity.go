package entity

// User has user data.
type User struct {
	AvatarURL         string `json:"avatar_url"`
	Login             string `json:"login"`
	Id                uint   `json:"id"`
	NodeId            string `json:"node_id"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	Gistsurl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	PublicRepos       uint   `json:"public_repos"`
	PublicGists       uint   `json:"public_gists"`
	Followers         uint   `json:"followers"`
	Following         uint   `json:"following"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	AvatarURLBase64   string
}

// QueryParam has query parameters
type QueryParam struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

// ProfileCardData has data for profile template
type ProfileCardData struct {
	User User
	QP   QueryParam
}
