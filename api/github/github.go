package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kena0ki/github-profile-card/api/entity"
	"github.com/kena0ki/github-profile-card/api/env"

	"github.com/pkg/errors"
)

var requestParams string

func init() {
	if env.GithubClientID == "" || env.GithubSecret == "" {
		requestParams = ""
		return
	}
	requestParams = fmt.Sprintf("?client_id=%v&client_secret=%v", env.GithubClientID, env.GithubSecret)
}

// GetUserData request user data.
func GetUserData(ctx context.Context, userName string) (*entity.User, error) {
	uri := fmt.Sprintf("https://api.github.com/users/%v", userName)
	req, err := http.NewRequest("GET", uri+requestParams, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to new request, url: %v", uri))
	}
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to request url: %v, Please make sure the user exists", uri)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	user := &entity.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAvatar request user avatar.
func GetAvatar(ctx context.Context, avatarURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", avatarURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to new request, url: %v", avatarURL))
	}
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to request url: %v, Please make sure the user exists", avatarURL)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
