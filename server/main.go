package main // import "github.com/kena0ki/github-profile-card"

import (
	"context"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	l "github.com/google/logger"
	"github.com/kena0ki/github-profile-card/server/env"
	"github.com/kena0ki/github-profile-card/server/router"
)

func main() {
	l.SetFlags(log.LstdFlags)
	defer l.Init("Logger", true, false, ioutil.Discard).Close()

	r := router.InitRouter()

	srv := &http.Server{
		Addr:    env.Port,
		Handler: r,
	}

	go func() {
		// service connections
		l.Info("About to start server on " + env.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Error("Listen: %s\n", err)
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
	l.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Error("Server Shutdown: ", err)
	}
	l.Info("Server exiting")
}
