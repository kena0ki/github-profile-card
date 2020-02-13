package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
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

var dir string
var svgTmplPath string
var sampleHTML string

func init() {
	var err error
	// TODO get root directory
	dir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	svgTmplPath = filepath.Join(dir, "svg/*.tmpl")
	sampleHTMLPath := filepath.Join(dir, "../samples/sample.html")
	println(sampleHTMLPath) // TODO use logger
	dat, err := ioutil.ReadFile(sampleHTMLPath)
	if err == nil {
		sampleHTML = string(dat)
	}
}

// GitHubProfileCard responds GitHub profile card in SVG format.
func GitHubProfileCard(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), env.Timeout*time.Second)
	defer cancel()
	type okObjType struct {
		contentLength int64
		contentType   string
		reader        io.Reader
		extraHeader   map[string]string
	}
	type errObjType struct {
		Message string `json:"message"`
	}
	type respType struct {
		code   int
		okObj  okObjType
		errObj errObjType
	}
	doneCh := make(chan respType)

	go func() {
		userName := strings.Replace(c.Param("user"), ".svg", "", 1)
		logger.Info(userName)
		user, err := github.GetUserData(ctx, userName)
		if err != nil {
			logger.Error(err)
			doneCh <- respType{
				code: http.StatusInternalServerError,
				errObj: errObjType{
					Message: err.Error(),
				},
			}
			return
		}
		avatar, err := github.GetAvatar(ctx, user.AvatarURL)
		if err != nil {
			logger.Error(err)
			doneCh <- respType{
				code: http.StatusInternalServerError,
				errObj: errObjType{
					Message: err.Error(),
				},
			}
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
			logger.Error(err)
			doneCh <- respType{
				code: http.StatusInternalServerError,
				errObj: errObjType{
					Message: err.Error(),
				},
			}
			return
		}
		doneCh <- respType{
			code: http.StatusOK,
			okObj: okObjType{
				contentLength: int64(len(buffer.Bytes())),
				contentType:   "image/svg+xml; charset=UTF-8",
				reader:        buffer,
				extraHeader:   map[string]string{},
			},
		}
		return
	}()

	select {
	case resp := <-doneCh:
		if resp.code == http.StatusOK {
			c.DataFromReader(resp.code, resp.okObj.contentLength,
				resp.okObj.contentType, resp.okObj.reader, resp.okObj.extraHeader)
		} else {
			c.JSON(resp.code, resp.errObj)
		}
	case <-ctx.Done():
		msg := fmt.Sprintf("Processing timed out in %d seconds", env.Timeout)
		logger.Error(msg)
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": msg,
		})
	}
}

// SamplePage responds sample page.
func SamplePage(c *gin.Context) {
	if sampleHTML == "" {
		c.JSON(http.StatusNotFound, "No sample provided")
	}
	// TODO Is this right usage?
	c.Writer.WriteString(sampleHTML)
}
