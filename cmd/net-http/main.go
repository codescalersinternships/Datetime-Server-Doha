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

	go func() {
		log.Printf("Server listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	
	pkgsh.ShutDown(&server)
}
