package main

import (
	"log"
	"net/http"

	pkgin "github.com/dohaelsawy/codescalers/datetimeserver/pkg/gin"
	pkgsh "github.com/dohaelsawy/codescalers/datetimeserver/pkg/shutdown"
	"github.com/gin-gonic/gin"
)

// setup gin engin and call datetime handeler
func setupHandler() *gin.Engine {
	router := gin.Default()

	router.GET("/datetime", pkgin.DateTime)

	return router
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: setupHandler(),
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
