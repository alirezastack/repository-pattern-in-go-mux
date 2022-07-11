package main

import (
	"antoccino/configs"
	"antoccino/helpers"
	"antoccino/routes"
	"context"
	"flag"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()

	// initialize MongoDB connection
	client, cancel := configs.ConnectDB()
	defer cancel()

	repo := &helpers.MongoDBRepository{
		Client: client,
	}

	log.Printf("loading service routes...")
	routes.UserRoute(router, repo)
	log.Printf("all service routes are loaded")

	srv := &http.Server{
		Addr: endpoint,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("ListenAndServe error: %s", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	log.Printf("running server on %s", endpoint)

	// Block until we receive our signal.
	//<-c
	sig := <-c
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
