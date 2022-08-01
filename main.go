package main

import (
	"antoccino/configs"
	"antoccino/helpers"
	"antoccino/routes"
	"antoccino/store"
	"context"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var wait time.Duration
	var debug bool

	flag.DurationVar(
		&wait,
		"graceful-timeout",
		configs.GracefulTimeout(),
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m",
	)

	flag.BoolVar(
		&debug,
		"debug",
		false,
		"sets log level to debug",
	)

	flag.Parse()

	zerolog.TimeFieldFormat = time.RFC3339
	// Default log level is info, unless debug flag is set to true
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Debug().Float64("ServerTimeout", wait.Seconds()).Msgf("Graceful-timeout is set to %s", wait)

	endpoint := configs.ServiceAddress()
	router := gin.New()

	router.Use(
		// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
		// By default gin.DefaultWriter = os.Stdout
		gin.LoggerWithFormatter(helpers.LogFormatter),

		// Recovery middleware recovers from any panics and writes a 500 if there was one.
		gin.CustomRecovery(helpers.GinCustomRecovery),
	)

	mongoStore := store.NewMongoDBStore()

	log.Info().Msg("loading service routes...")
	routes.UserRoute(router, mongoStore)
	log.Info().Msg("all service routes are loaded")

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
			log.Info().Msgf("ListenAndServe error: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT ()
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Info().Msgf("running server on %s", endpoint)

	// Block until we receive our signal
	sig := <-quit
	log.Info().Msgf("received shutdown signal: %+v", sig)

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Error().Msgf("server shutdown with error: %s\n", err)
		os.Exit(1)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info().Msg("server shutdown gracefully ;)")
	os.Exit(0)

}
