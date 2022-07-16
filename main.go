package main

import (
	"antoccino/configs"
	"antoccino/helpers"
	"antoccino/routes"
	"antoccino/store"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var wait time.Duration

	flag.DurationVar(&wait, "graceful-timeout", configs.GracefulTimeout(), "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	log.Printf("Graceful-timeout is set to %s", wait)

	endpoint := configs.ServiceAddress()
	router := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// custom log format
		return fmt.Sprintf("%s - [%s] %s %s %s %d %s %s %s",
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			log.Printf("An internal server error occurred: %s", err)
			helpers.ReturnResponse(c, errors.New(err), http.StatusInternalServerError)
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	mongoStore := store.NewMongoDBStore()

	log.Printf("loading service routes...")
	routes.UserRoute(router, mongoStore)
	log.Printf("all service routes are loaded")

	srv := &http.Server{
		Addr: endpoint,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gin/gonic in.
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("ListenAndServe error: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT ()
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("running server on %s", endpoint)

	// Block until we receive our signal
	sig := <-quit
	log.Printf("received shutdown signal: %+v", sig)

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("server shutdown with error: %s\n", err)
		os.Exit(1)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("server shutdown gracefully ;)")
	os.Exit(0)

}
