package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"hello-world/http_handlers"
)

func main() {
	var (
		port int
		env  string
	)

	flag.IntVar(&port, "port", 3000, "Listing port")
	flag.StringVar(&env, "env", "development", "When development mode, it's ignored to check webhook signature.")
	flag.Parse()

	if err := os.Setenv("ENV", env); err != nil {
		//TODO: Log error
	}

	handler, err := http_handlers.DefaultHandler()
	if err != nil {
		os.Exit(1)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: handler,
	}

	doneCh := make(chan os.Signal, 1)
	signal.Notify(doneCh, os.Interrupt)

	go func() {
		<-doneCh

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			//TODO: Log error
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(http.ErrServerClosed, err) {
			//TODO: Log error
		}
	}
}
