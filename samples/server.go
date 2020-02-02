package main // import "github.com/kena0ki/github-profile-card"

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

const port string = ":8081"
const apiPort string = ":8080"
const hostName string = "gpc.znoo.xyz"

func main() {
	logger.SetFlags(log.LstdFlags)
	defer logger.Init("Logger", true, false, ioutil.Discard).Close()

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dat, err := ioutil.ReadFile(dir + "/sample.html")
	if err != nil {
		panic(err)
	}
	html := strings.ReplaceAll(string(dat), hostName, "localhost"+port)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		logger.Info(html)
		c.Writer.WriteString(html)
	})

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		// service connections
		logger.Info("About to start server on " + port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("Listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown: ", err)
	}
	logger.Info("Server exiting")
}
