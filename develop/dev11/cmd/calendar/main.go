package main

import (
	"context"
	"fmt"
	"github.com/nglmq/calendar/internal/app/config"
	"github.com/nglmq/calendar/internal/app/server"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

func run(ctx context.Context) error {
	r, err := server.NewServer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	port, err := config.ReadJSONConfig("config.json")

	httpServer := &http.Server{
		Addr:    net.JoinHostPort("localhost", port),
		Handler: r,
	}

	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := httpServer.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
