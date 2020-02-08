package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/logger"
	"github.com/kena0ki/github-profile-card/api/entity"
	"github.com/kena0ki/github-profile-card/api/env"
	"github.com/kena0ki/github-profile-card/api/github"

	"github.com/gin-gonic/gin"
)

// ErrorResponse is common errror response format.
type ErrorResponse struct {
	Message string `json:"message"`
}

const hostName string = "gpc.znoo.xyz"

var dir string
var svgTmplPath string
var sampleHtml string

func init() {
	var err error
	// TODO get root directory
	dir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	svgTmplPath = filepath.Join(dir, "svg/*.tmpl")
	sampleHtmlPath := filepath.Join(dir, "../samples/sample.html")
	println(sampleHtmlPath) // TODO use logger
	dat, err := ioutil.ReadFile(sampleHtmlPath)
	if err == nil {
		if env.GinMode == "release" {
			sampleHtml = string(dat)
		} else {
			sampleHtml = strings.ReplaceAll(string(dat), hostName, "localhost"+env.Port)
		}
	}
}

func handleError(err error, c *gin.Context, ch chan struct{}) {
	logger.Error(err)
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Message: err.Error(),
	})
	ch <- struct{}{}
	return
}

// GetSVG creates GitHub SVG Card.
func GetSVG(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), env.Timeout*time.Second)
	defer cancel()
	doneCh := make(chan struct{})

	go func() {
		userName := strings.Replace(c.Param("user"), ".svg", "", 1)
		logger.Info(userName)
		user, err := github.GetUserData(ctx, userName)
		if err != nil {
			handleError(err, c, doneCh)
			return
		}
		avatar, err := github.GetAvatar(ctx, user.AvatarURL)
		if err != nil {
			handleError(err, c, doneCh)
			return
		}
		user.AvatarURLBase64 = base64.StdEncoding.EncodeToString(avatar)
		profileDate := entity.ProfileCardData{
			User: *user,
			QP: entity.QueryParam{
				Width:  c.DefaultQuery("width", "400"),
				Height: c.DefaultQuery("height", "110"),
			},
		}
		tmpl := template.Must(template.ParseGlob(svgTmplPath))
		buffer := new(bytes.Buffer)
		err = tmpl.Execute(buffer, profileDate)
		if err != nil {
			handleError(err, c, doneCh)
			return
		}
		c.DataFromReader(http.StatusOK, int64(len(buffer.Bytes())), "image/svg+xml; charset=UTF-8", buffer, map[string]string{})
		doneCh <- struct{}{}
		return
	}()

	select {
	case <-doneCh:
		return
	case <-ctx.Done():
		msg := fmt.Sprintf("Processing timed out in %d seconds", env.Timeout)
		logger.Error(msg)
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": msg,
		})
	}
}

func GetSampleHtml(c *gin.Context) {
	if sampleHtml == "" {
		c.JSON(http.StatusNotFound, "No sample provided")
	}
	// TODO Is this right usage?
	c.Writer.WriteString(sampleHtml)
}
