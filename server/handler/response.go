package handler

var documentationURL = "https://github.com/kena0ki/github-profile-card"

// Response is common response format.
type Response struct {
	Value         string `json:"value"`
	CardURL       string `json:"card_url"`
	RepositoryURL string `json:"repository_url"`
}

// ErrorResponse is common errror response format.
type ErrorResponse struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentaion_url"`
}
