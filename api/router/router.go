package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/kena0ki/github-profile-card/api/handler"
	"golang.org/x/time/rate"
)

// InitRouter provide initialized router
func InitRouter() *gin.Engine {
	router := gin.Default()

	limiter := rate.NewLimiter(rate.Limit(1), 50)
	router.Use(func(c *gin.Context) {
		if err := limiter.Wait(c); err != nil {
			logger.Warningf("Too many requests, %v", c.Request)
			c.JSON(http.StatusTooManyRequests, "Too many requests, please wait for a while")
			c.Abort()
		}
		c.Next()
	})

	router.GET("/", handler.SamplePage)
	router.GET("/api/github/:user", handler.GitHubProfileCard)
	router.HEAD("/api/github/:user", handler.GitHubProfileCard) // For Edge, which uses HEAD method for some reason 
	return router
}
