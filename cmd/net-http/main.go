package main

import (
	"log"
	"net/http"

	pkghttp "github.com/dohaelsawy/codescalers/datetimeserver/pkg/net-http"
	pkgsh "github.com/dohaelsawy/codescalers/datetimeserver/pkg/shutdown"
)

func main() {
	server := http.Server{
		Addr:    ":8090",
		Handler: http.HandlerFunc(pkghttp.DateTime),
	}

	// goroutines to make server running while shutdown function wating to inturrupt signal
	go func() {
		log.Printf("Server listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// shuts the server gracefully
	pkgsh.ShutDown(&server)
}
