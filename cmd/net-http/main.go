package main

import (
	"log"
	"net/http"

	serverHttp "github.com/codescalersinternships/Datetime-Server-Doha/server/net-http"
	serverShuts "github.com/codescalersinternships/Datetime-Server-Doha/server/shutdown"
)

func main() {
	server := http.Server{
		Addr:    ":8090",
		Handler: http.HandlerFunc(serverHttp.DateTimeHandler),
	}

	go func() {
		log.Printf("Server listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	serverShuts.ShutDown(&server)
}
