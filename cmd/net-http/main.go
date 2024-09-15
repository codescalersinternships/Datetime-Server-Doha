package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pkg "github.com/dohaelsawy/codescalers/datetimeserver/pkg/net-http"
)


func main() {

	// Create a new HTTP server
    srv := &http.Server{
        Addr: ":8090",
		Handler: http.HandlerFunc(pkg.DateTime),
    }


    // Create a channel to receive signals
    signalCh := make(chan os.Signal, 1)
    signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)


    // Start the server in a separate goroutine
    go func() {
        log.Printf("Server listening on %s\n", srv.Addr)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()


    // Wait for a signal to shutdown the server
    sig := <-signalCh
    log.Printf("Received signal: %v\n", sig)

	
    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Shutdown the server gracefully
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server shutdown failed: %v\n", err)
    }

    log.Println("Server shutdown gracefully")
}
