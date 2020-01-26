package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	l "github.com/google/logger"
	"github.com/kena0ki/github-profile-card/server/handler"
	"golang.org/x/time/rate"
)

// InitRouter provide initialized router
func InitRouter() *gin.Engine {
	router := gin.Default()

	limiter := rate.NewLimiter(rate.Limit(1), 50)
	router.Use(func(c *gin.Context) {
		if err := limiter.Wait(c); err != nil {
			l.Warningf("To many requests, %v", c.Request)
			c.JSON(http.StatusTooManyRequests, "To many requests, Please wait for a while")
			c.Abort()
		}
		c.Next()
	})

	router.GET("/api/github/:user", handler.GetSVG)
	return router
}
