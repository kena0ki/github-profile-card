package env

import (
	"os"
)

var (
	// Port - API port.
	Port string
	// LogLevel - logging level.
	LogLevel string
	// GithubClientID - github client id.
	GithubClientID string
	// GithubSecret - github client id.
	GithubSecret string
)

const (
	// Timeout second of request timeout.
	Timeout = 30
)

func init() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = ":8080"
	}
	Port = port
	LogLevel = os.Getenv("LOG_LEVEL")
	GithubClientID = os.Getenv("GITHUB_CLIENT_ID")
	GithubSecret = os.Getenv("GITHUB_SECRET")
}
