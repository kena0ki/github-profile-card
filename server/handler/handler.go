package handler

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	l "github.com/google/logger"
	"github.com/kena0ki/github-profile-card/server/env"
	"github.com/kena0ki/github-profile-card/server/github"

	"github.com/gin-gonic/gin"
)

// GetSVG creates GitHub SVG Card.
func GetSVG(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), env.Timeout*time.Second)
	defer cancel()
	doneCh := make(chan struct{})

	go func() {
		userName := c.Param("user")
		l.Info(userName)
		user, err := github.GetUserData(ctx, userName)
		if err != nil {
			l.Error(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message:          err.Error(),
				DocumentationURL: documentationURL,
			})
			doneCh <- struct{}{}
			return
		}

		dir, err := os.Getwd()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message:          err.Error(),
				DocumentationURL: documentationURL,
			})
			doneCh <- struct{}{}
			return
		}
		pattern := filepath.Join(dir, "svg/*.tmpl")
		tmpl := template.Must(template.ParseGlob(pattern))
		buffer := new(bytes.Buffer)
		err = tmpl.Execute(buffer, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message:          err.Error(),
				DocumentationURL: documentationURL,
			})
			doneCh <- struct{}{}
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
		l.Error(msg)
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": msg,
		})
	}
}
